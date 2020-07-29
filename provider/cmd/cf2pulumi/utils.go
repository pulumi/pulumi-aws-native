package main

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/goccy/go-yaml/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/zclconf/go-cty/cty"
)

// camel replaces the first contiguous string of upper case runes in the given string with its lower-case equivalent.
func camel(s string) string {
	c, sz := utf8.DecodeRuneInString(s)
	if sz == 0 || unicode.IsLower(c) {
		return s
	}

	// The first rune is not lowercase. Iterate until we find a rune that is.
	var word []rune
	for {
		s = s[sz:]

		n, nsz := utf8.DecodeRuneInString(s)
		if nsz == 0 {
			word = append(word, unicode.ToLower(c))
			return string(word)
		}
		if unicode.IsLower(n) {
			if len(word) == 0 {
				c = unicode.ToLower(c)
			}
			word = append(word, c)
			return string(word) + s
		}
		c, sz, word = n, nsz, append(word, unicode.ToLower(c))
	}
}

// isLegalIdentifierStart returns true if it is legal for c to be the first character of an HCL2 identifier.
func isLegalIdentifierStart(c rune) bool {
	return c == '$' || c == '_' ||
		unicode.In(c, unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl)
}

// isLegalIdentifierPart returns true if it is legal for c to be part of an HCL2 identifier.
func isLegalIdentifierPart(c rune) bool {
	return isLegalIdentifierStart(c) || unicode.In(c, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc)
}

// makeLegalIdentifier deletes characters that are not allowed in HCL2 identifiers with underscores. No attempt is
// made to ensure that the result is unique.
func makeLegalIdentifier(name string) string {
	var builder strings.Builder
	for i, c := range name {
		if isLegalIdentifierPart(c) {
			if i == 0 && !isLegalIdentifierStart(c) {
				builder.WriteRune('_')
			}
			builder.WriteRune(c)
		}
	}
	if builder.Len() == 0 {
		return "x"
	}
	return builder.String()
}

// mapValues returns a slice of MappingValueNodes if the provided node is a MappingNode or MappingValueNode.
func mapValues(v ast.Node) ([]*ast.MappingValueNode, bool) {
	switch v := v.(type) {
	case *ast.MappingNode:
		return v.Values, true
	case *ast.MappingValueNode:
		return []*ast.MappingValueNode{v}, true
	default:
		return nil, false
	}
}

// keyString returns the string value of the given MappingValueNode's Key.
func keyString(v *ast.MappingValueNode) string {
	return v.Key.(*ast.StringNode).Value
}

// valueAt returns the value in the slice of MappingValueNodes with the given Key, if any.
func valueAt(values []*ast.MappingValueNode, key string) (ast.Node, bool) {
	for _, f := range values {
		if keyString(f) == key {
			return f.Value, true
		}
	}
	return nil, false
}

// plainLit returns an unquoted string literal expression.
func plainLit(v string) *model.LiteralValueExpression {
	return &model.LiteralValueExpression{Value: cty.StringVal(v)}
}

// quotedLit returns a quoted string literal expression.
func quotedLit(v string) *model.TemplateExpression {
	return &model.TemplateExpression{Parts: []model.Expression{plainLit(v)}}
}

// getStringValue returns the literal string value of the given expression, if any.
func getStringValue(v model.Expression) (string, bool) {
	switch v := v.(type) {
	case *model.LiteralValueExpression:
		if v.Value.Type() == cty.String {
			return v.Value.AsString(), true
		}
	case *model.TemplateExpression:
		if len(v.Parts) == 1 {
			if lit, ok := v.Parts[0].(*model.LiteralValueExpression); ok {
				if lit.Value.Type() == cty.String {
					return lit.Value.AsString(), true
				}
			}
		}
	}
	return "", false
}

// objectConsItem returns a new ObjectConsItem with the given key and value.
func objectConsItem(key string, value model.Expression) model.ObjectConsItem {
	var keyExpr model.Expression = plainLit(key)
	if !hclsyntax.ValidIdentifier(key) {
		keyExpr = quotedLit(key)
	}
	return model.ObjectConsItem{
		Key:   keyExpr,
		Value: value,
	}
}

// invoke returns a new call to `invoke` with the given token and inputs. The inputs are combined into an
// ObjectConsExpression.
func invoke(token string, inputs ...model.ObjectConsItem) *model.FunctionCallExpression {
	args := []model.Expression{quotedLit(token)}
	if len(inputs) != 0 {
		args = append(args, &model.ObjectConsExpression{Items: inputs})
	}
	return &model.FunctionCallExpression{
		Name: "invoke",
		Args: args,
	}
}

// relativeTraversal returns a new RelativeTraversalExpression that accesses the given attribute of the source
// expression.
func relativeTraversal(source model.Expression, attr string) *model.RelativeTraversalExpression {
	return &model.RelativeTraversalExpression{
		Source:    source,
		Traversal: hcl.Traversal{hcl.TraverseAttr{Name: attr}},
		Parts:     []model.Traversable{model.DynamicType, model.DynamicType},
	}
}

// resourceToken returns the Pulumi token for the given CloudFormation resource type.
func resourceToken(typ string) string {
	components := strings.Split(typ, "::")
	if len(components) != 3 {
		return typ
	}
	moduleName, resourceName := components[1], components[2]

	// Override the name of the Config module.
	if moduleName == "Config" {
		moduleName = "Configuration"
	}
	return "cloudformation:" + moduleName + ":" + resourceName
}
