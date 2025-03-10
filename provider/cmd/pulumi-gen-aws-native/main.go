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
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"slices"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/pkg/errors"
	jsschema "github.com/pulumi/jsschema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// ensureDir ensures that a target directory exists (like `mkdir -p`), returning a non-nil error if any problem occurs.
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o700)
}

type schemaUrls []string

// CloudFormationAutoNameOverlay represents the overlay file `aws-cfn-autoname-overlay.json`.
// Some resources have name requirements that are not defined yet in the CloudFormation schema.
// This file allows us to define those requirements as an overlay. It will never overwrite existing
// values in the schema, so as they add these requirements to the schema, it will overwrite what we have.
type CloudFormationAutoNameOverlay struct {
	// Resources is a map of CloudFormation resource types to their auto-naming properties.
	Resources map[string]AutoNamingOverlay `json:"Resources"`
}

// AutoNamingOverlay is a map of property names to their auto-naming properties.
type AutoNamingOverlay map[string]AutoNamingProps

// AutoNamingProps is the auto-naming properties for a property.
// These properties will overwrite the auto-naming properties in the schema only
// if the schema property does not already have a value set.
type AutoNamingProps struct {
	MinLength int `json:"minLength"`
	MaxLength int `json:"maxLength"`
}

func (s *schemaUrls) String() string {
	return fmt.Sprint(*s)
}

func (s *schemaUrls) Set(value string) error {
	if len(*s) > 0 {
		return errors.New("--schema-urls flag already set")
	}

	urls := strings.Split(value, ",")
	for _, u := range urls {
		*s = append(*s, u)
	}
	return nil
}

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s [OPTIONS] <operation>"
		_, err := fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		contract.IgnoreError(err)
		flag.PrintDefaults()
	}

	var version, operation, schemaFolder, docsUrl, metadataFolder string
	var jsonSchemaUrls schemaUrls
	flag.StringVar(&version, "version", "", "the provider version to record in the generated code")
	flag.StringVar(&schemaFolder, "schema-folder", "", "The folder containing the CloudFormation schema files")
	flag.StringVar(&docsUrl, "docs-url", "", "The URL to download the CloudFormation docs")
	flag.StringVar(&metadataFolder, "metadata-folder", "", "The folder containing metadata files needed for schema generation")
	flag.Var(&jsonSchemaUrls, "schema-urls", "A comma delimited list of CloudFormation schema urls")

	flag.Parse()
	operation = flag.Arg(0)
	fmt.Printf("%s: version: %s, schema-folder: %s, docs-url: %s, schema-urls: %s, metadata-folder: %s\n", operation, version, schemaFolder, docsUrl, jsonSchemaUrls, metadataFolder)
	if version == "" && operation == "" && (jsonSchemaUrls == nil && docsUrl == "") {
		flag.Usage()
		return
	}

	genDir := filepath.Join(".", "provider", "cmd", "pulumi-gen-aws-native")
	semanticsDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")
	autonameOverlay := filepath.Join(".", "aws-cfn-autoname-overlay.json")
	autoNameOverlay, err := readAutonamingOverlay(autonameOverlay)
	if err != nil {
		fmt.Println("Error reading autonaming overlay file: ", err)
		return
	}
	refDBFile := filepath.Join(".", "meta", "ref-db.json")

	switch operation {
	case "docs":
		if docsUrl == "" {
			fmt.Println("Error: docs operation requires additional --docs-url argument")
			flag.Usage()
			return
		}

		if err := downloadCloudFormationDocs(docsUrl, filepath.Join(".", "aws-cloudformation-docs")); err != nil {
			fatalf("error download CloudFormation docs: %v", err)
		}

	case "discovery":
		if jsonSchemaUrls == nil {
			fmt.Println("Error: discovery operation requires additional --schema-urls argument")
			flag.Usage()
			return
		}

		if err := writeSupportedResourceTypes(genDir); err != nil {
			fatalf("error writing supported resource types: %v", err)
		}
		if err := downloadCloudFormationSchemas(jsonSchemaUrls, filepath.Join(".", schemaFolder), genDir); err != nil {
			fatalf("error downloading CloudFormation schemas: %v", err)
		}

	case "schema":
		supportedTypes := readSupportedResourceTypes(genDir)
		jsonSchemas := readJsonSchemas(schemaFolder, autoNameOverlay)
		docsTypes, err := schema.ReadCloudFormationDocsFile(filepath.Join(".", "aws-cloudformation-docs", "CloudFormationDocumentation.json"))
		if err != nil {
			fatalf("error reading CloudFormation docs file: %v", err)
		}
		semanticsDocument, err := schema.GatherSemantics(semanticsDir)
		if err != nil {
			fatalf("error gathering semantics: %v", err)
		}

		regions, err := readRegions(metadataFolder)
		if err != nil {
			fatalf("error reading regions: %v", err)
		}

		refDB := new(refdb.RefDB)
		if err := refDB.LoadJSONFile(refDBFile); err != nil {
			fatalf("error reading %q", refDBFile)
		}

		packageSpec, meta, reports, err := schema.GatherPackage(supportedTypes, jsonSchemas, false, &semanticsDocument, docsTypes, regions, refDB)
		if err != nil {
			fatalf("error generating schema: %v", err)
		}

		fmt.Println("Generating examples...")
		err = generateExamples(packageSpec, meta, []string{"nodejs", "python", "dotnet", "go"})
		if err != nil {
			fatalf("error generating examples: %v", err)
		}
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")
		if err = writePulumiSchema(*packageSpec, providerDir, true /* includeUncompressed */); err != nil {
			fatalf("error writing schema: %v", err)
		}

		cf2pulumiDir := filepath.Join(".", "provider", "cmd", "cf2pulumi")
		if err = writePulumiSchema(*packageSpec, cf2pulumiDir, false /* includeUncompressed */); err != nil {
			fatalf("error writing schema: %v", err)
		}
		if err = writePulumiSchema(*packageSpec, "bin", false /* includeUncompressed */); err != nil {
			fatalf("error writing schema: %v", err)
		}

		// Emit the resource metadata for cf2pulumi.
		if err = writeMetadata(meta, cf2pulumiDir, false /* includeUncompressed */); err != nil {
			fatalf("error writing metadata: %v", err)
		}
		// Also, emit the resource metadata for the provider.
		if err = writeMetadata(meta, providerDir, true /* includeUncompressed */); err != nil {
			fatalf("error writing metadata: %v", err)
		}
		// Also write to bin folder to be picked up as an asset for testing.
		if err = writeMetadata(meta, "bin", true /* includeUncompressed */); err != nil {
			fatalf("error writing metadata: %v", err)
		}
		// Emit the reports for the provider.
		if err := reports.WriteToDirectory("reports"); err != nil {
			fatalf("error writing reports: %v", err)
		}

	default:
		fatalf("Unrecognized language '%s'", operation)
	}
}

// Fail loud and clear.
func fatalf(format string, a ...any) {
	barrier := strings.Repeat("=", 110)
	log.Fatalf(fmt.Sprintf("schema generation failed\n%s\n%s\n%s\n", barrier, fmt.Sprintf(format, a...), barrier))
}

func readJsonSchemas(schemaDir string, overlay map[string]AutoNamingOverlay) (res []*jsschema.Schema) {
	var fileNames []string
	root := filepath.Join(".", schemaDir)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		fatalf("error reading JSON schemas: %v", err)
	}

	sort.Strings(fileNames)
	for _, fileName := range fileNames {
		res = append(res, readJsonSchema(fileName, overlay))
	}
	return
}

func readJsonSchema(schemaPath string, overlay map[string]AutoNamingOverlay) *jsschema.Schema {
	s, err := jsschema.ReadFile(schemaPath)
	if err != nil {
		fatalf("%v", errors.Wrapf(err, schemaPath))
	}
	mergeAutoNaming(s, overlay)

	return s
}

func mergeAutoNaming(schema *jsschema.Schema, overlay map[string]AutoNamingOverlay) {
	if cfTypeName, ok := schema.Extras["typeName"].(string); ok {
		if typeOverlay, ok := overlay[cfTypeName]; ok {
			for propertyName, autoName := range typeOverlay {
				if p, ok := schema.Properties[propertyName]; ok {
					if p.MaxLength.Val == 0 {
						p.MaxLength.Val = autoName.MaxLength
					}
					if p.MinLength.Val == 0 {
						p.MinLength.Val = autoName.MinLength
					}
				}
			}
		}
	}
}

const regionsFile = "regions.json"

func readRegions(metadataFolder string) ([]schema.RegionInfo, error) {
	path := filepath.Join(metadataFolder, regionsFile)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var regions []schema.RegionInfo
	err = json.Unmarshal(bytes, &regions)
	if err != nil {
		return nil, err
	}
	return regions, nil
}

const supportedResourcesFile = "supported-types.txt"
const deprecatedResourcesFile = "deprecated-types.txt"

func readSupportedResourceTypes(outDir string) []string {
	path := filepath.Join(outDir, supportedResourcesFile)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.Trim(string(bytes), "\n"), "\n")
}

func readDeprecatedResourceTypes(outdir string) []string {
	path := filepath.Join(outdir, deprecatedResourcesFile)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	types := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	fileNames := []string{}
	for _, typ := range types {
		name := strings.ReplaceAll(typ, "::", "-")
		fileNames = append(fileNames, fmt.Sprintf("%s.json", strings.ToLower(name)))
	}
	return fileNames
}

func readAutonamingOverlay(file string) (map[string]AutoNamingOverlay, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var overlay CloudFormationAutoNameOverlay
	err = json.Unmarshal(bytes, &overlay)
	if err != nil {
		return nil, err
	}
	return overlay.Resources, nil
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

	for _, t := range excludedTypes() {
		supported.Delete(t)
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
	supportedContent := strings.Join(stillSupported.SortedValues(), "\n") + "\n"
	return emitFile(outDir, supportedResourcesFile, []byte(supportedContent))
}

func cleanDir(dir string, perm os.FileMode, keep ...string) error {
	if keep == nil {
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}

		return os.MkdirAll(dir, perm)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if slices.Contains(keep, entry.Name()) {
			continue
		}
		err := os.RemoveAll(filepath.Join(dir, entry.Name()))
		if err != nil {
			return err
		}
	}
	return nil
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

func downloadCloudFormationSchemas(urls []string, outDir string, genDir string) error {
	// start with an empty directory
	err := cleanDir(outDir, 0755, readDeprecatedResourceTypes(genDir)...)
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
	if err := ensureDir(filepath.Dir(p)); err != nil {
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
