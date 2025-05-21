// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/cf2pulumi"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/provider"
	pschema "github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/version"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

//go:embed schema.json.gz
var schemaBytes []byte

//go:embed metadata.json.gz
var metadataBytes []byte

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [target] [template path]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Valid targets are dotnet, go, nodejs, and python.\n")
		os.Exit(-1)
	}

	target, templatePath := os.Args[1], os.Args[2]

	pkgSpec, err := loadSchema()
	if err != nil {
		log.Fatalf("failed to load schema: %v", err)
	}

	metadata, err := provider.LoadMetadata(metadataBytes)
	if err != nil {
		log.Fatalf("failed to load metadata: %v", err)
	}
	template, diags, err := cf2pulumi.RenderFile(templatePath, metadata)
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}
	if len(diags) > 0 {
		syntax.NewDiagnosticWriter(os.Stderr, nil, 0, true).WriteDiagnostics(diags)
	}

	cf2pulumi.FormatBody(template)
	programText := fmt.Sprintf("%v", template)

	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programText), "program.pp"); err != nil {
		log.Fatalf("failed to parse IR: %v", err)
	}
	if parser.Diagnostics.HasErrors() {
		fmt.Print(programText)
		parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		os.Exit(-1)
	}

	hcl2Cache := pcl.Cache(pcl.NewPackageCache())
	pkg, err := schema.ImportSpec(*pkgSpec, nil, schema.ValidationOptions{AllowDanglingReferences: true})
	if err != nil {
		log.Fatalf("failed to parse import the spec: %v", err)
	}
	loaderOption := pcl.Loader(pschema.InMemoryPackageLoader(map[string]*schema.Package{
		"aws-native": pkg,
	}))

	program, diags, err := pcl.BindProgram(parser.Files, hcl2Cache, loaderOption)
	if err != nil {
		log.Fatalf("failed to bind program: %v", err)
	}
	if diags.HasErrors() {
		fmt.Print(programText)
		program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		os.Exit(-1)
	}

	var files map[string][]byte
	switch target {
	case "dotnet":
		files, diags, err = dotnet.GenerateProgram(program)
	case "go":
		files, diags, err = gogen.GenerateProgram(program)
	case "nodejs":
		files, diags, err = nodejs.GenerateProgram(program)
	case "python":
		files, diags, err = python.GenerateProgram(program)
	case "pulumi":
		files = map[string][]byte{"program.pp": []byte(programText)}
	}
	if err != nil {
		log.Fatalf("failed to generate program: %v", err)
	}
	if diags.HasErrors() {
		program.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(diags)
		os.Exit(-1)
	}

	for _, f := range files {
		fmt.Print(string(f))
	}
}

// loadSchema loads the serialized/compressed schema generated during generation from schema.go.
func loadSchema() (*schema.PackageSpec, error) {
	var pkgSpec schema.PackageSpec
	uncompressed, err := gzip.NewReader(bytes.NewReader(schemaBytes))
	if err != nil {
		return nil, fmt.Errorf("loading schema: %w", err)
	}

	if err = json.NewDecoder(uncompressed).Decode(&pkgSpec); err != nil {
		return nil, fmt.Errorf("deserializing schema: %w", err)
	}
	if err = uncompressed.Close(); err != nil {
		return nil, fmt.Errorf("closing uncompress stream for schema: %w", err)
	}
	// embed version because go codegen is particularly sensitive to this.
	if pkgSpec.Version == "" {
		pkgSpec.Version = version.Version
	}
	return &pkgSpec, nil
}
