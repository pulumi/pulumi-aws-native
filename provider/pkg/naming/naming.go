// Copyright 2016-2021, Pulumi Corporation.

package naming

import (
	"strings"
)

// ToSdkName converts a Cloud Formation property or attribute name to the lowerCamelCase convention that
// is used in Pulumi schema's properties.
func ToSdkName(s string) string {
	if s == "" {
		return s
	}
	s = LowerAcronyms(s)
	if r := rune(s[0]); r >= 'A' && r <= 'Z' {
		s = strings.ToLower(string(r)) + s[1:]
	}
	return s
}

// ToCfnName converts a lowerCamelCase schema property name to the Cloud Formation property or attribute name
// either by looking up in the table if present or converting to PascalCase.
func ToCfnName(s string, lookupTable map[string]string) string {
	if name, ok := lookupTable[s]; ok {
		return name
	}
	return ToPascalCase(s)
}

// ToPascalCase converts a lowerCamelCase schema property name to PascalCase .
func ToPascalCase(s string) string {
	if r := rune(s[0]); r >= 'a' && r <= 'z' {
		s = strings.ToUpper(string(r)) + s[1:]
	}
	return s
}

// HasUppercaseAcronym checks if a CamelCase string contains an Uppercase acronym
// by looking for runs of capitals longer than 2
func HasUppercaseAcronym(s string) bool {
	startIndex, endIndex := firstUppercaseAcronym(s)
	if startIndex == -1 {
		return false
	}
	// ignore single character "acronyms" since these will not be modified by lowerAcronyms
	// Note: we've defined uppercase to be ASCII [A-Z] so index math is safe here
	if endIndex-startIndex > 1 {
		return true
	}
	return HasUppercaseAcronym(s[endIndex:])
}

// lowers the trailing chars of any uppercase acronyms
func LowerAcronyms(s string) string {
	startIndex, endIndex := firstUppercaseAcronym(s)
	if startIndex == -1 {
		return s
	}

	// Note: we've defined uppercase to be ASCII [A-Z] so index math is safe here
	startIndex = startIndex + 1 // don't lower the first char of the run

	return s[:startIndex] + strings.ToLower(s[startIndex:endIndex]) + LowerAcronyms(s[endIndex:])
}

// Returns the indices of the first Uppercase acronym in the string
func firstUppercaseAcronym(s string) (int, int) {
	startIndex, endIndex := findFirstRunOfUppercase(s, 2)
	if startIndex == -1 {
		return startIndex, endIndex
	}

	// Treat the last uppercase char in a run as part of the next word UNLESS:
	// - we're at the end of the string
	// - the acronym is followed by a single lowercase 's' (eg. as in "ARNs")
	// Note: we've defined uppercase to be ASCII [A-Z] so index math is safe here
	if !(endIndex == len(s) || startsWithIsolatedLowercaseS(s[endIndex:])) {
		endIndex = endIndex - 1
	}

	return startIndex, endIndex
}

func startsWithIsolatedLowercaseS(s string) bool {
	switch len(s) {
	case 0:
		return false
	case 1:
		return rune(s[0]) == 's'
	default:
		return rune(s[0]) == 's' && isUpperAcronymChar(rune(s[1]))
	}
}

// returns the indicies of the first run of at least minLength uppercase characters encountered in str
func findFirstRunOfUppercase(s string, minLength int) (int, int) {
	startIndex := -1
	for i, char := range s {
		if startIndex == -1 { // looking for first uppercase char
			if isUpperAcronymChar(char) {
				startIndex = i
			}
		} else { // in a run, looking for non-uppercase char
			if !isUpperAcronymChar(char) {
				// Note: we've defined uppercase to be ASCII [0-9A-Z] so index math is safe here
				if i-startIndex >= minLength {
					return startIndex, i
				} else {
					startIndex = -1
				}
			}
		}
	}

	if startIndex == -1 {
		return -1, -1
	}

	return startIndex, len(s)
}

func isUpperAcronymChar(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
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
