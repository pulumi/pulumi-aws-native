// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	ctx "context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/pkg/errors"
	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <operation> <schema-folder> <version> (<schema urls>) <docs-url>"
		_, err := fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		contract.IgnoreError(err)
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", "", "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		flag.Usage()
		return
	}

	operation, schemaFolder, version := args[0], args[1], args[2]
	genDir := filepath.Join(".", "provider", "cmd", "pulumi-gen-aws-native")
	semanticsDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")

	switch operation {
	case "discovery":
		if len(args) < 5 {
			fmt.Println("Error: discovery operation requires additional schema urls argument")
			flag.Usage()
			return
		}
		jsonSchemaUrls := strings.Split(args[3], ",")

		if err := writeSupportedResourceTypes(genDir); err != nil {
			panic(fmt.Errorf("error writing supported resource types: %v", err))
		}
		if err := downloadCloudFormationSchemas(jsonSchemaUrls, filepath.Join(".", schemaFolder)); err != nil {
			panic(fmt.Errorf("error downloading CloudFormation schemas: %v", err))
		}

		docsUrl := args[4]
		if err := downloadCloudFormationDocs(docsUrl, filepath.Join(".", "aws-cloudformation-docs")); err != nil {
			panic(fmt.Errorf("error download CloudFormation docs: %v", err))
		}

	case "schema":
		supportedTypes := readSupportedResourceTypes(genDir)
		jsonSchemas := readJsonSchemas(schemaFolder)
		docsTypes, err := schema.ReadCloudFormationDocsFile(filepath.Join(".", "aws-cloudformation-docs", "CloudFormationDocumentation.json"))
		if err != nil {
			panic(fmt.Errorf("error reading CloudFormation docs file: %v", err))
		}
		semanticsDocument, err := schema.GatherSemantics(semanticsDir)
		if err != nil {
			panic(fmt.Errorf("error gathering semantics: %v", err))
		}

		packageSpec, meta, reports, err := schema.GatherPackage(supportedTypes, jsonSchemas, false, &semanticsDocument, docsTypes)
		if err != nil {
			panic(fmt.Errorf("error generating schema: %v", err))
		}

		fmt.Println("Generating examples...")
		err = generateExamples(packageSpec, meta, []string{"nodejs", "python", "dotnet", "go"})
		if err != nil {
			panic(fmt.Errorf("error generating examples: %v", err))
		}
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")
		if err = writePulumiSchema(*packageSpec, providerDir, true /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing schema: %v", err))
		}

		cf2pulumiDir := filepath.Join(".", "provider", "cmd", "cf2pulumi")
		if err = writePulumiSchema(*packageSpec, cf2pulumiDir, false /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing schema: %v", err))
		}
		if err = writePulumiSchema(*packageSpec, "bin", false /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing schema: %v", err))
		}

		// Emit the resource metadata for cf2pulumi.
		if err = writeMetadata(meta, cf2pulumiDir, false /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing metadata: %v", err))
		}
		// Also, emit the resource metadata for the provider.
		if err = writeMetadata(meta, providerDir, true /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing metadata: %v", err))
		}
		// Also write to bin folder to be picked up as an asset for testing.
		if err = writeMetadata(meta, "bin", true /* includeUncompressed */); err != nil {
			panic(fmt.Errorf("error writing metadata: %v", err))
		}
		// Emit the reports for the provider.
		if err := reports.WriteToDirectory("reports"); err != nil {
			panic(fmt.Errorf("error writing reports: %v", err))
		}

	default:
		panic(fmt.Sprintf("Unrecognized language '%s'", operation))
	}
}

func readJsonSchemas(schemaDir string) (res []*jsschema.Schema) {
	var fileNames []string
	root := filepath.Join(".", schemaDir)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	sort.Strings(fileNames)
	for _, fileName := range fileNames {
		res = append(res, readJsonSchema(fileName))
	}
	return
}

func readJsonSchema(schemaPath string) *jsschema.Schema {
	s, err := jsschema.ReadFile(schemaPath)
	if err != nil {
		panic(errors.Wrapf(err, schemaPath))
	}

	return s
}

const supportedResourcesFile = "supported-types.txt"
const deprecatedResourcesFile = "deprecated-types.txt"

func readSupportedResourceTypes(outDir string) []string {
	path := filepath.Join(outDir, supportedResourcesFile)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}

func writeSupportedResourceTypes(outDir string) error {
	cfg, err := config.LoadDefaultConfig(ctx.Background())
	cfg.Region = "us-west-2"
	if err != nil {
		return errors.Wrapf(err, "could not load AWS config")
	}

	cfn := cloudformation.NewFromConfig(cfg)

	supported := codegen.NewStringSet()
	for _, provisioningType := range []types.ProvisioningType{types.ProvisioningTypeFullyMutable, types.ProvisioningTypeImmutable} {
		var nextToken *string
		for {
			out, err := cfn.ListTypes(ctx.Background(), &cloudformation.ListTypesInput{
				Visibility:       "PUBLIC",
				ProvisioningType: provisioningType,
				NextToken:        nextToken,
			})
			if err != nil {
				return errors.Wrapf(err, "could not list types")
			}

			for _, r := range out.TypeSummaries {
				supported.Add(*r.TypeName)
			}
			nextToken = out.NextToken
			if nextToken == nil {
				break
			}
		}
	}

	// Load existing supported resources, check which are no longer supported,
	// but don't remove support automatically - maintain previously supported list.
	previouslySupported := codegen.NewStringSet(readSupportedResourceTypes(outDir)...)
	noLongerSupported := previouslySupported.Subtract(supported)
	unsupportedContent := strings.Join(noLongerSupported.SortedValues(), "\n")
	err = emitFile(outDir, deprecatedResourcesFile, []byte(unsupportedContent))
	if err != nil {
		return err
	}

	stillSupported := supported.Union(previouslySupported)
	supportedContent := strings.Join(stillSupported.SortedValues(), "\n")
	return emitFile(outDir, supportedResourcesFile, []byte(supportedContent))
}

func cleanDir(dir string, perm os.FileMode) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return os.MkdirAll(dir, perm)
}

func downloadCloudFormationDocs(url, outDir string) error {
	// start with an empty directory
	err := cleanDir(outDir, 0755)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	outPath := filepath.Join(outDir, "CloudFormationDocumentation.json")
	if err := os.WriteFile(outPath, body, 0644); err != nil {
		return err
	}
	return nil
}

func downloadCloudFormationSchemas(urls []string, outDir string) error {
	// start with an empty directory
	err := cleanDir(outDir, 0755)
	if err != nil {
		return err
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
		if err != nil {
			return err
		}

		// Read all the files from zip archive
		for _, f := range zipReader.File {
			outPath := filepath.Join(outDir, f.Name)

			outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			rc, err := f.Open()
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			outFile.Close()
			rc.Close()
		}
	}

	return nil
}

func writePulumiSchema(pkgSpec pschema.PackageSpec, outdir string, includeUncompressed bool) error {
	compressedSchema, err := schema.CompressSchema(pkgSpec)
	if err != nil {
		return errors.Wrap(err, "failed to compress schema")
	}
	err = emitFile(outdir, "schema.json.gz", []byte(compressedSchema))
	if err != nil {
		panic(errors.Wrap(err, "saving metadata"))
	}

	if includeUncompressed {
		pkgSpec.Version = ""
		schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
		if err != nil {
			panic(errors.Wrap(err, "marshaling Pulumi schema"))
		}
		return emitFile(outdir, "schema.json", schemaJSON)
	}

	return nil
}

func writeMetadata(metadata *metadata.CloudAPIMetadata, outDir string, includeUncompressed bool) error {
	compressedMeta := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedMeta)
	err := json.NewEncoder(compressedWriter).Encode(metadata)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	if err = compressedWriter.Close(); err != nil {
		return err
	}

	formatted, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	err = emitFile(outDir, "metadata.json.gz", []byte(compressedMeta.Bytes()))
	if err != nil {
		return err
	}

	if includeUncompressed {
		err = emitFile(outDir, "metadata.json", formatted)
		if err != nil {
			return err
		}
	}
	return nil
}

// emitFile creates a file in a given directory and writes the byte contents to it.
func emitFile(outDir, relPath string, contents []byte) error {
	p := filepath.Join(outDir, relPath)
	if err := tools.EnsureDir(filepath.Dir(p)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}
