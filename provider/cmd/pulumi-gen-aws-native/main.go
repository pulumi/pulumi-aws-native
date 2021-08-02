package main

import (
	ctx "context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
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

	languages, inputFile, version := args[0], args[1], args[2]
	genDir := filepath.Join(".", "provider", "cmd", "pulumi-gen-aws-native")

	if languages == "discovery" {
		if err := writeSupportedResourceTypes(genDir); err != nil {
			panic(err)
		}
		return
	}

	pkgSpec := gatherPackage(readCFNSchema(inputFile), readSupportedResourceTypes(genDir))
	pkgSpec.Version = version
	ppkg, err := pschema.ImportSpec(pkgSpec, nil)
	if err != nil {
		panic(fmt.Sprintf("error importing schema: %v", err))
	}

	for _, language := range strings.Split(languages, ",") {
		outdir := path.Join(".", "sdk", language)
		providerDir := filepath.Join(".", "provider", "cmd", "pulumi-resource-aws-native")

		switch language {
		case "nodejs":
			writeNodeJSSDK(ppkg, outdir)
		case "python":
			writePythonSDK(ppkg, outdir)
		case "dotnet":
			writeDotnetSDK(ppkg, outdir)
		case "go":
			writeGoSDK(ppkg, outdir)
		case "schema":
			writePulumiSchema(pkgSpec, providerDir)
		default:
			panic(fmt.Sprintf("Unrecognized language '%s'", language))
		}
	}
}

func readCFNSchema(schemaPath string) schema.CloudFormationSchema {
	// Read in, decode, and import the schema.
	schemaBytes, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		panic(err)
	}

	var sch schema.CloudFormationSchema
	if err = json.Unmarshal(schemaBytes, &sch); err != nil {
		panic(err)
	}

	return sch
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
	for _, provisioningType := range []types.ProvisioningType {types.ProvisioningTypeFullyMutable, types.ProvisioningTypeImmutable} {
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
	mustWriteFile(outDir, supportedResourcesFile, []byte(val))
	return nil
}

func writeNodeJSSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := nodejsgen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writePythonSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := pythongen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writeDotnetSDK(pkg *pschema.Package, outdir string) {
	overlays := map[string][]byte{}
	files, err := dotnetgen.GeneratePackage("pulumigen", pkg, overlays)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writeGoSDK(pkg *pschema.Package, outdir string) {
	files, err := gogen.GeneratePackage("pulumigen", pkg)
	if err != nil {
		panic(err)
	}
	mustWriteFiles(outdir, files)
}

func writePulumiSchema(pkgSpec pschema.PackageSpec, outdir string) {
	pkgSpec.Version = ""
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		panic(errors.Wrap(err, "marshaling Pulumi schema"))
	}
	mustWriteFile(outdir, "schema.json", schemaJSON)
}

func mustWriteFiles(rootDir string, files map[string][]byte) {
	for filename, contents := range files {
		mustWriteFile(rootDir, filename, contents)
	}
}

func mustWriteFile(rootDir, filename string, contents []byte) {
	outPath := filepath.Join(rootDir, filename)

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		panic(err)
	}
	err := ioutil.WriteFile(outPath, contents, 0600)
	if err != nil {
		panic(err)
	}
}
