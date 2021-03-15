// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"strings"
)

// ToPropertyName converts a Cloud Formation property or attribute name to the lowerCamelCase convention that
// is used in Pulumi schema's properties.
func ToPropertyName(s string) string {
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return s
}
