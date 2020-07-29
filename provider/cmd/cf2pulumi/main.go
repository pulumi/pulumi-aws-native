package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v2/codegen/go"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v2/codegen/python"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [target] [template path]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Valid targets are dotnet, go, nodejs, and python.\n")
		os.Exit(-1)
	}

	target, templatePath := os.Args[1], os.Args[2]

	template, err := RenderFile(templatePath)
	if err != nil {
		log.Fatalf("failed to render template: %v", err)
	}
	formatBody(template)
	programText := fmt.Sprintf("%v", template)

	parser := syntax.NewParser()
	if err := parser.ParseFile(strings.NewReader(programText), "prorgram.pp"); err != nil {
		log.Fatalf("failed to parse IR: %v", err)
	}
	if parser.Diagnostics.HasErrors() {
		fmt.Print(programText)
		parser.NewDiagnosticWriter(os.Stderr, 0, true).WriteDiagnostics(parser.Diagnostics)
		os.Exit(-1)
	}

	program, diags, err := hcl2.BindProgram(parser.Files)
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
