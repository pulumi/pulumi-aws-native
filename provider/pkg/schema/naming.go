// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"strings"
)

// ToSdkName converts a Cloud Formation property or attribute name to the lowerCamelCase convention that
// is used in Pulumi schema's properties.
func ToSdkName(s string) string {
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return s
}

// ToCfnName converts a lowerCamelCase schema property name to the Cloud Formation property or attribute name
// in PascalCase.
func ToCfnName(s string) string {
	if r := rune(s[0]); r >= 'a' && r <= 'z' {
		s = strings.ToUpper(string(r)) + s[1:]
	}
	return s
}
