package schema

import (
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/naming"
)

// Represents a property listed in `readOnlyProperties`, `writeOnlyProperties`,
// or `createOnlyProperties` in the Cloud Control schema.
type ResourceProperty string

// ToSdkName converts the property name to an SDK-compatible name.
// This will also handle converting all nested properties.
//   - Nested object properties are delimited with `/`
//   - Nested array properties are delimited with `/*/`
//
// e.g.
//   - `Foo/Bar/Baz` => `foo/bar/baz`
//   - `Foo/*/BarBaz` => `foo/*/barBaz`
func (r ResourceProperty) ToSdkName() string {
	arrayProps := strings.Split(string(r), "/*/")
	for i, p := range arrayProps {
		nested := strings.Split(p, "/")
		for idx, n := range nested {
			nested[idx] = naming.ToSdkName(n)
		}
		arrayProps[i] = strings.Join(nested, "/")
	}
	return strings.Join(arrayProps, "/*/")
}
