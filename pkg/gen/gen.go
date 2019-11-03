//go:generate go run assets_generate.go

package gen

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudformation/pkg/schema"
	"github.com/pulumi/pulumi/pkg/util/contract"
)

type nestedMember struct {
	Name     string
	Optional bool
	Type     string
}

type nestedTypeDef struct {
	Name    string
	Members []*nestedMember
}

type resourceDef struct {
	Name               string
	AttributesType     string
	PropertiesOptional bool
	PropertiesType     string
	Token              string
}

type moduleDef struct {
	Namespace   string
	NestedTypes map[string]*nestedTypeDef
	Resources   map[string]*resourceDef
}

type packageDef struct {
	schema schema.CloudFormationSchema

	Modules   map[string]*moduleDef
	RootTypes map[string]*nestedTypeDef
}

func primitiveTypeName(primitiveType string, isInput bool) string {
	var name string
	switch primitiveType {
	case "String", "Json", "Timestamp":
		name = "string"
	case "Long", "Integer", "Double":
		name = "number"
	case "Boolean":
		name = "boolean"
	default:
		contract.Failf("unexpected primitive type '%v'", primitiveType)
	}
	if isInput {
		return fmt.Sprintf("pulumi.Input<%s>", name)
	}
	return name
}

func (p *packageDef) getModule(name string) *moduleDef {
	mod, has := p.Modules[name]
	if !has {
		mod = &moduleDef{
			Namespace:   strings.ToLower(name),
			NestedTypes: make(map[string]*nestedTypeDef),
			Resources:   make(map[string]*resourceDef),
		}
		p.Modules[name] = mod
	}
	return mod
}

func (p *packageDef) nestedTypeName(resourceType, typ string, isInput bool) string {
	propertyTypeName := resourceType + "." + typ

	var name string
	_, ok := p.schema.PropertyTypes[propertyTypeName]
	if ok {
		name = strings.Split(resourceType, "::")[2] + typ
	} else {
		_, ok = p.schema.PropertyTypes[typ]
		contract.Assertf(ok, "missing property type '%v' in scope %v", typ, resourceType)
		name = typ
	}
	if isInput {
		name += "Properties"
	} else {
		name += "Attributes"
	}
	return name
}

func (p *packageDef) itemTypeName(resourceType, itemType, primitiveItemType string, isInput bool) string {
	if primitiveItemType != "" {
		return primitiveTypeName(primitiveItemType, isInput)
	}
	name := p.nestedTypeName(resourceType, itemType, isInput)
	if isInput {
		return fmt.Sprintf("pulumi.Input<%s>", name)
	}
	return name
}

func (p *packageDef) propertyTypeName(resourceType, itemType, primitiveItemType, primitiveType, typ string, isInput bool) string {
	if primitiveType != "" {
		return primitiveTypeName(primitiveType, isInput)
	}

	var name string
	switch typ {
	case "List":
		name = fmt.Sprintf("%s[]", p.itemTypeName(resourceType, itemType, primitiveItemType, isInput))
	case "Map":
		name = fmt.Sprintf("{[key: string]: %s}", p.itemTypeName(resourceType, itemType, primitiveItemType, isInput))
	case "":
		contract.Failf("unexpected empty property type")
	default:
		name = p.nestedTypeName(resourceType, typ, isInput)
	}
	if isInput {
		return fmt.Sprintf("pulumi.Input<%s>", name)
	}
	return name
}

func (p *packageDef) gatherNestedType(name, moduleName, resourceType string, spec schema.PropertyTypeSpec, isInput bool) *nestedTypeDef {
	nt := &nestedTypeDef{Name: name}

	memberNames := make([]string, 0, len(spec.Properties))
	for n := range spec.Properties {
		memberNames = append(memberNames, n)
	}
	sort.Strings(memberNames)

	for _, n := range memberNames {
		spec := spec.Properties[n]

		if strings.Contains(n, ".") {
			n = fmt.Sprintf(`"%s"`, n)
		}
		nt.Members = append(nt.Members, &nestedMember{
			Name:     n,
			Optional: !spec.Required,
			Type:     p.propertyTypeName(resourceType, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type, isInput),
		})
	}

	if moduleName == "" {
		p.RootTypes[name] = nt
	} else {
		p.getModule(moduleName).NestedTypes[name] = nt
	}
	return nt
}

func (p *packageDef) gatherPropertyType(name string, spec schema.PropertyTypeSpec, isInput bool) *nestedTypeDef {
	components := strings.Split(name, ".")

	var resourceType, moduleName string
	if len(components) == 1 {
		name = components[0]
	} else {
		resourceType = components[0]
		name = components[1]

		resourceTypeComponents := strings.Split(resourceType, "::")
		resourceName := resourceTypeComponents[2]

		moduleName, name = resourceTypeComponents[1], resourceName+name
	}
	if isInput {
		name += "Properties"
	} else {
		name += "Attributes"
	}

	return p.gatherNestedType(name, moduleName, resourceType, spec, isInput)
}

func (p *packageDef) gatherAttributesType(resourceType string, attributes map[string]schema.AttributeSpec) *nestedTypeDef {
	components := strings.Split(resourceType, "::")
	moduleName, resourceName := components[1], components[2]

	nt := &nestedTypeDef{Name: resourceName + "Attributes"}

	memberNames := make([]string, 0, len(attributes))
	for n := range attributes {
		memberNames = append(memberNames, n)
	}
	sort.Strings(memberNames)

	for _, n := range memberNames {
		spec := attributes[n]
		if strings.Contains(n, ".") {
			n = fmt.Sprintf(`"%s"`, n)
		}
		nt.Members = append(nt.Members, &nestedMember{
			Name: n,
			Type: p.propertyTypeName(resourceType, spec.ItemType, spec.PrimitiveItemType, spec.PrimitiveType, spec.Type, false),
		})
	}

	p.getModule(moduleName).NestedTypes[nt.Name] = nt
	return nt
}

func (p *packageDef) gatherResourceType(resourceType string, spec schema.ResourceSpec) {
	components := strings.Split(resourceType, "::")
	contract.Assertf(len(components) == 3, "unexpected resource type '%v'", resourceType)

	moduleName, name := components[1], components[2]
	propertiesOptional := true
	for _, spec := range spec.Properties {
		if spec.Required {
			propertiesOptional = false
			break
		}
	}

	propertiesType := p.gatherNestedType(name+"Properties", moduleName, resourceType, spec.PropertyTypeSpec, true)
	attributesType := p.gatherAttributesType(resourceType, spec.Attributes)

	rt := &resourceDef{
		Name:               name,
		AttributesType:     attributesType.Name,
		PropertiesOptional: propertiesOptional,
		PropertiesType:     propertiesType.Name,
		Token:              fmt.Sprintf("cloudformation:%s:%s", moduleName, name),
	}

	p.getModule(moduleName).Resources[name] = rt
}

func gatherPackage(schema schema.CloudFormationSchema) *packageDef {
	p := &packageDef{
		schema:    schema,
		Modules:   make(map[string]*moduleDef),
		RootTypes: make(map[string]*nestedTypeDef),
	}

	// Gather nested property and attribute types
	propertyTypeNames := make([]string, 0, len(schema.PropertyTypes))
	for n := range schema.PropertyTypes {
		propertyTypeNames = append(propertyTypeNames, n)
	}
	sort.Strings(propertyTypeNames)

	for _, n := range propertyTypeNames {
		spec := schema.PropertyTypes[n]
		p.gatherPropertyType(n, spec, true)
		p.gatherPropertyType(n, spec, false)
	}

	// Gather resource types.
	resourceTypeNames := make([]string, 0, len(schema.ResourceTypes))
	for n := range schema.ResourceTypes {
		resourceTypeNames = append(resourceTypeNames, n)
	}
	sort.Strings(resourceTypeNames)

	for _, n := range resourceTypeNames {
		spec := schema.ResourceTypes[n]
		p.gatherResourceType(n, spec)
	}

	return p
}

func GenerateSDK(w io.Writer, schema schema.CloudFormationSchema) error {
	templateFile, err := templates.Open("cloudformation.ts")
	if err != nil {
		return errors.Wrap(err, "could not open template file")
	}
	defer contract.IgnoreClose(templateFile)

	templateText, err := ioutil.ReadAll(templateFile)
	if err != nil {
		return errors.Wrap(err, "could not load template file")
	}

	t, err := template.New("cloudformation.ts").Parse(string(templateText))
	if err != nil {
		return errors.Wrap(err, "could not parse template file")
	}

	return t.Execute(w, gatherPackage(schema))
}
