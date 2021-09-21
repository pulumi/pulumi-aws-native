// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"bytes"
	"compress/gzip"
	ctx "context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	jsschema "github.com/lestrrat-go/jsschema"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <languages> <schema-file> <root-dir>"
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

	languages, schemaFolder, version := args[0], args[1], args[2]
	genDir := filepath.Join(".", "provider", "cmd", "pulumi-gen-aws-native")

	if languages == "discovery" {
		if err := writeSupportedResourceTypes(genDir); err != nil {
			panic(err)
		}
		return
	}

	supportedTypes := readSupportedResourceTypes(genDir)
	jsonSchemas := readJsonSchemas(schemaFolder)

	fullSpec, _, err := schema.GatherPackage(supportedTypes, jsonSchemas, true)
	if err != nil {
		panic(fmt.Sprintf("error generating schema: %v", err))
	}
	fullSpec.Version = version

	for _, language := range strings.Split(languages, ",") {
		fmt.Printf("Generating %s...\n", language)
		switch language {
		case "nodejs", "python", "dotnet", "go":
			dir := filepath.Join(".", "sdk", language)
			err = emitPackage(fullSpec, language, dir)
		case "schema":
			cf2pulumiDir := filepath.Join(".", "provider", "cmd", "cf2pulumi")
			writePulumiSchema(*fullSpec, cf2pulumiDir, "schema-full.json", true)

			supportedSpec, meta, err := schema.GatherPackage(supportedTypes, jsonSchemas, false)
			if err != nil {
				panic(fmt.Sprintf("error generating schema: %v", err))
			}

			fmt.Println("Generating examples...")
			err = generateExamples(supportedSpec, []string{"nodejs", "python", "dotnet", "go"})
			if err != nil {
				panic(fmt.Sprintf("error generating examples: %v", err))
			}
			providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")
			writePulumiSchema(*supportedSpec, providerDir, "schema.json",true)

			// Also, emit the resource metadata for the provider.
			if err = writeMetadata(meta, providerDir, "main", true); err != nil {
				break
			}
		default:
			panic(fmt.Sprintf("Unrecognized language '%s'", language))
		}
	}
}

func readJsonSchemas(schemaDir string) (res []jsschema.Schema) {
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

func readJsonSchema(schemaPath string) jsschema.Schema {
	s, err := jsschema.ReadFile(schemaPath)
	if err != nil {
		panic(errors.Wrapf(err, schemaPath))
	}

	return *s
}

const supportedResourcesFile = "supported-types.txt"

func readSupportedResourceTypes(outDir string) []string {
	path := filepath.Join(outDir, supportedResourcesFile)
	bytes, err := ioutil.ReadFile(path)
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

	var result []string
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
				result = append(result, *r.TypeName)
			}
			nextToken = out.NextToken
			if nextToken == nil {
				break
			}
		}
	}
	sort.Strings(result)

	val := strings.Join(result, "\n")
	return emitFile(outDir, supportedResourcesFile, []byte(val))
}

func generate(ppkg *pschema.Package, language string) (map[string][]byte, error) {
	toolDescription := "the Pulumi SDK Generator"
	extraFiles := map[string][]byte{}
	switch language {
	case "nodejs":
		return nodejsgen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "python":
		return pythongen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "go":
		return gogen.GeneratePackage(toolDescription, ppkg)
	case "dotnet":
		return dotnetgen.GeneratePackage(toolDescription, ppkg, extraFiles)
	}

	return nil, errors.Errorf("unknown language '%s'", language)
}

// emitPackage emits an entire package pack into the configured output directory with the configured settings.
func emitPackage(pkgSpec *pschema.PackageSpec, language, outDir string) error {
	ppkg, err := pschema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return errors.Wrap(err, "reading schema")
	}

	files, err := generate(ppkg, language)
	if err != nil {
		return errors.Wrapf(err, "generating %s package", language)
	}

	for f, contents := range files {
		if err := emitFile(outDir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
		}
	}

	return nil
}
func writePulumiSchema(pkgSpec pschema.PackageSpec, outdir, jsonFileName string, emitJSON bool) error {
	compressedSchema := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedSchema)
	err := json.NewEncoder(compressedWriter).Encode(pkgSpec)
	if err != nil {
		panic(errors.Wrap(err, "marshaling schema"))
	}
	if err = compressedWriter.Close(); err != nil {
		panic(err)
	}

	err = emitFile(outdir, "schema.go", []byte(fmt.Sprintf(`package main
var pulumiSchema = %#v
`, compressedSchema.Bytes())))
	if err != nil {
		panic(errors.Wrap(err, "saving metadata"))
	}

	if emitJSON {
		pkgSpec.Version = ""
		schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
		if err != nil {
			panic(errors.Wrap(err, "marshaling Pulumi schema"))
		}
		return emitFile(outdir, jsonFileName, schemaJSON)
	}

	return nil
}

func writeMetadata(metadata *schema.CloudAPIMetadata, outDir string, goPackageName string, emitJSON bool) error {
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

	err = emitFile(outDir, "metadata.go", []byte(fmt.Sprintf(`package %s
var cloudApiResources = %#v
`, goPackageName, compressedMeta.Bytes())))
	if err != nil {
		return err
	}

	if emitJSON {
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
