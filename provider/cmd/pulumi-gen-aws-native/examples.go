// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl/v2"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/cf2pulumi"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
	pschema "github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/pkg/v3/codegen/pcl"
	"github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateExamples(pkgSpec *schema.PackageSpec, metadata *metadata.CloudAPIMetadata, languages []string) error {
	// Find all snippets in the AWS CloudFormation Docs repo.
	folder := path.Join(".", "aws-cloudformation-user-guide", "doc_source")
	examples, err := findAllExamples(folder)
	if err != nil {
		return err
	}

	// Cache to speed up code generation.
	hcl2Cache := pcl.Cache(pcl.NewPackageCache())
	pkg, err := schema.ImportSpec(*pkgSpec, nil, schema.ValidationOptions{AllowDanglingReferences: true})
	if err != nil {
		return err
	}
	loaderOption := pcl.Loader(pschema.InMemoryPackageLoader(map[string]*schema.Package{
		"aws-native": pkg,
	}))

	// Render examples to SDK languages.
	total := 0
	examplesRenderData := map[string][]exampleRenderData{}
	for _, yaml := range examples {
		example, err := generateExample(yaml, metadata, languages, hcl2Cache, loaderOption)
		if err != nil {
			// Skip all snippets that don't produce valid examples.
			continue
		}
		if strings.HasPrefix(example.ResourceType, "aws-native:wafv2:WebACL") {
			// TODO: These examples are non-deterministic - figure out why.
			continue
		}
		var existing []exampleRenderData
		if other, ok := examplesRenderData[example.ResourceType]; ok {
			existing = other
		}
		examplesRenderData[example.ResourceType] = append(existing, exampleRenderData{
			ExampleDescription:       "Example",
			LanguageToExampleProgram: example.LanguageToExampleProgram,
		})
		total += 1
	}
	fmt.Printf("Number of examples: %v\n", total)

	// Write examples to the schema docs.
	for resourceType, data := range examplesRenderData {
		err = renderExampleToSchema(pkgSpec, resourceType, data)
		if err != nil {
			return errors.Wrapf(err, "rendering %s %+v", resourceType, data)
		}
	}
	return nil
}

func generateExample(yaml string, metadata *metadata.CloudAPIMetadata, languages []string, bindOpts ...pcl.BindOption) (*resourceExample, error) {
	body, diagnostics, err := cf2pulumi.RenderText(yaml, metadata)
	if err != nil {
		return nil, errors.Wrapf(err, "rendering YAML")
	}
	parser := syntax.NewParser()
	if diagnostics.HasErrors() {
		buf := new(bytes.Buffer)
		_ = parser.NewDiagnosticWriter(buf, 0, false).WriteDiagnostics(parser.Diagnostics)
		return nil, errors.Errorf("parser diagnostic errors: %s", buf)
	}

	cf2pulumi.FormatBody(body)
	text := fmt.Sprintf("%v", body)

	if err := parser.ParseFile(strings.NewReader(text), "program.pp"); err != nil {
		return nil, errors.Wrapf(err, "parsing IR")
	}
	if parser.Diagnostics.HasErrors() {
		buf := new(bytes.Buffer)
		_ = parser.NewDiagnosticWriter(buf, 0, false).WriteDiagnostics(parser.Diagnostics)
		return nil, errors.Errorf("parser diagnostic errors: %s", buf)
	}

	resourceType := ""

	perLanguage := languageToExampleProgram{}
	for _, target := range languages {

		program, diags, err := pcl.BindProgram(parser.Files, bindOpts...)
		if err != nil {
			return nil, errors.Wrapf(err, "binding program")
		}
		if diags.HasErrors() {
			buf := new(bytes.Buffer)
			_ = program.NewDiagnosticWriter(buf, 0, false).WriteDiagnostics(diags)
			return nil, errors.Errorf("bind diagnostic errors: %s", buf)
		}

		if resourceType == "" {
			for _, node := range program.Nodes {
				if res, ok := node.(*pcl.Resource); ok {
					resourceType = res.Token
					break
				}
			}
		}
		if resourceType == "" {
			return nil, errors.New("no resource node found")
		}

		var files map[string][]byte
		switch target {
		case "dotnet":
			files, diags, err = dotnet.GenerateProgram(program)
		case "go":
			files, diags, err = recoverableProgramGen(program, gogen.GenerateProgram)
		case "nodejs":
			files, diags, err = nodejs.GenerateProgram(program)
		case "python":
			files, diags, err = python.GenerateProgram(program)
		}
		if err != nil {
			if target == "go" {
				continue
			}
			return nil, errors.Wrapf(err, "generating program")
		}
		if diags.HasErrors() {
			if target == "go" {
				continue
			}
			buf := new(bytes.Buffer)
			_ = program.NewDiagnosticWriter(buf, 0, true).WriteDiagnostics(diags)
			return nil, errors.Errorf("generate diagnostic errors: %s", buf)
		}

		var sb strings.Builder
		for _, f := range files {
			sb.WriteString(string(f))
		}
		perLanguage[language(target)] = programText(sb.String())
	}
	return &resourceExample{
		ResourceType:             resourceType,
		LanguageToExampleProgram: perLanguage,
	}, nil
}

func findAllExamples(folder string) ([]string, error) {
	var fileNames, result []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(path, "aws-resource-") {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, fileName := range fileNames {
		examples, err := findExamples(fileName)
		if err != nil {
			return nil, err
		}
		result = append(result, examples...)
	}
	return result, nil
}

func findExamples(fileName string) ([]string, error) {
	docFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "opening file")
	}
	defer func() {
		_ = docFile.Close()
	}()

	var result []string
	var buf strings.Builder
	snippet := false

	scanner := bufio.NewScanner(docFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "```") {
			if snippet && buf.Len() > 0 {
				ex := naming.SanitizeCfnString(buf.String())
				result = append(result, ex)
			}
			snippet = !snippet
			continue
		}
		if snippet {
			buf.WriteString(line)
			buf.WriteRune('\n')
		} else {
			buf.Reset()
		}
	}
	return result, scanner.Err()
}

type programGenFn func(*pcl.Program) (map[string][]byte, hcl.Diagnostics, error)

func recoverableProgramGen(program *pcl.Program, fn programGenFn) (files map[string][]byte, d hcl.Diagnostics, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered during generation: %v", r)
		}
	}()

	return fn(program)
}

type programText string
type language string

type languageToExampleProgram map[language]programText
type exampleRenderData struct {
	ExampleDescription       string
	LanguageToExampleProgram languageToExampleProgram
}
type resourceExample struct {
	ResourceType             string
	LanguageToExampleProgram languageToExampleProgram
}

func renderExampleToSchema(pkgSpec *schema.PackageSpec, resourceName string,
	examplesRenderData []exampleRenderData) error {
	const tmpl = `

{{"{{% examples %}}"}}
## Example Usage
{{- range . }}
{{ "{{% example %}}" }}
### {{ .ExampleDescription }}

{{- range $lang, $example := .LanguageToExampleProgram }}
{{ beginLanguage $lang }}
{{ $example }}
{{ endLanguage }}
{{ end }}
{{"{{% /example %}}"}}
{{- end }}
{{"{{% /examples %}}"}}
`
	res, ok := pkgSpec.Resources[resourceName]
	if !ok {
		return fmt.Errorf("missing resource from schema: %s", resourceName)
	}

	t, err := template.New("examples").Funcs(template.FuncMap{
		"beginLanguage": func(lang interface{}) string {
			l := fmt.Sprintf("%s", lang)
			switch l {
			case "nodejs":
				l = "typescript"
			case "dotnet":
				l = "csharp"
			}
			return fmt.Sprintf("```%s", l)
		},
		"endLanguage": func() string {
			return "```"
		},
	}).Parse(tmpl)
	if err != nil {
		return err
	}
	b := strings.Builder{}
	if err = t.Execute(&b, examplesRenderData); err != nil {
		return err
	}
	res.Description += b.String()
	pkgSpec.Resources[resourceName] = res
	return nil
}
