package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudformation/provider/pkg/schema"
	dotnetgen "github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v2/codegen/python"
	pschema "github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <target> <swagger-or-schema-file> <root-pulumi-kubernetes-dir>"
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

	target, inputFile, baseDir := args[0], args[1], args[2]

	outdir := filepath.Join(baseDir, "sdk", target)
	providerDir := filepath.Join(baseDir, "provider", "cmd", "pulumi-resource-cloudformation")
	//providerPkgDir := filepath.Join(baseDir, "provider", "pkg", "provider")

	switch target {
	case "nodejs":
		writeNodeJSSDK(readPulumiSchema(inputFile, version), outdir)
	case "python":
		writePythonSDK(readPulumiSchema(inputFile, version), outdir)
	case "dotnet":
		writeDotnetSDK(readPulumiSchema(inputFile, version), outdir)
	case "go":
		writeGoSDK(readPulumiSchema(inputFile, version), outdir)
		//	case "provider":
		//		moduleMap := buildModuleMap(readCFNSchema(inputFile))
		//		writeProvider(moduleMap, providerPkgDir)
	case "schema":
		pkgSpec := gatherPackage(readCFNSchema(inputFile))
		writePulumiSchema(pkgSpec, providerDir)
	default:
		panic(fmt.Sprintf("Unrecognized target '%s'", target))
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

func readPulumiSchema(schemaPath, version string) *pschema.Package {
	// Read in, decode, and import the schema.
	schemaBytes, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		panic(err)
	}

	var pkgSpec pschema.PackageSpec
	if err = json.Unmarshal(schemaBytes, &pkgSpec); err != nil {
		panic(err)
	}
	pkgSpec.Version = version

	pkg, err := pschema.ImportSpec(pkgSpec, nil)
	if err != nil {
		panic(err)
	}
	return pkg
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
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		panic(errors.Wrap(err, "marshaling Pulumi schema"))
	}
	mustWriteFile(outdir, "schema.json", schemaJSON)
}

//func writeProvider(moduleMap map[string]string, outdir string) {
//	mapSource := fmt.Sprintf(`package provider
//
//var moduleMap = %#v
//`, moduleMap)
//	formatted, err := format.Source([]byte(mapSource))
//	if err != nil {
//		panic(errors.Wrap(err, "formatting module map"))
//	}
//	mustWriteFile(outdir, "moduleMap.go", formatted)
//}

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
