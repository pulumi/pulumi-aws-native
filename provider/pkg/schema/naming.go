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

// ToUpperCamel converts a string to UpperCamelCase.
func ToUpperCamel(s string) string {
	return toCamelInitCase(s, true)
}

func toCamelInitCase(s string, initCase bool) string {
	if s == strings.ToUpper(s) {
		// lowercase the UPPER_SNAKE_CASE
		s = strings.ToLower(s)
	}

	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' || v == '.' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}
