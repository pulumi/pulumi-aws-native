// Copyright 2016-2023, Pulumi Corporation.

package autonaming

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type NamingPredicate string

const (
	NamingPredicateEquals             NamingPredicate = "equal"
	NamingPredicateNotEqual           NamingPredicate = "notEqual"
	NamingPredicateContains           NamingPredicate = "contains"
	NamingPredicateNotContains        NamingPredicate = "notContains"
	NamingPredicateStartsWith         NamingPredicate = "startsWith"
	NamingPredicateNotStartsWith      NamingPredicate = "notStartsWith"
	NamingPredicateEndsWith           NamingPredicate = "endsWith"
	NamingPredicateNotEndsWith        NamingPredicate = "notEndsWith"
	NamingPredicateLessThan           NamingPredicate = "lessThan"
	NamingPredicateLessThanOrEqual    NamingPredicate = "lessThanOrEqual"
	NamingPredicateGreaterThan        NamingPredicate = "greaterThan"
	NamingPredicateGreaterThanOrEqual NamingPredicate = "greaterThanOrEqual"
	NamingPredicateIn                 NamingPredicate = "in"
	NamingPredicateNotIn              NamingPredicate = "notIn"
	NamingPredicateExists             NamingPredicate = "exists"
	NamingPredicateNotExists          NamingPredicate = "notExists"
)

type NamingCondition struct {
	Predicate NamingPredicate        `json:"predicate,omitempty"`
	Value     resource.PropertyValue `json:"value,omitempty"`
}

type NamingRule struct {
	Field     resource.PropertyKey `json:"field,omitempty"`
	Condition *NamingCondition     `json:"condition,omitempty"`
}

type SemanticsSpecDocument struct {
	Resources map[string]SemanticsSpec `json:"resources,omitempty"`
}

type SemanticsSpec struct {
	CfType           string                      `json:"cf,omitempty"`
	NamingTriviaSpec map[string]NamingTriviaSpec `json:"namingTriviaSpec,omitempty"`
}

type NamingTriviaSpec struct {
	Rule   *NamingRule   `json:"rule,omitempty"`
	Trivia *NamingTrivia `json:"trivia,omitempty"`
}

// NamingTrivia is a struct which contains any derived information
// about a resource's name that is used to generate a random name.
type NamingTrivia struct {
	Suffix string `json:"suffix,omitempty"`
}

func CheckTriviaCondition(props resource.PropertyMap, rule *NamingRule) (bool, error) {
	switch rule.Condition.Predicate {
	case NamingPredicateEquals:
		return props[rule.Field].DeepEqualsIncludeUnknowns(rule.Condition.Value), nil
	default:
		return false, fmt.Errorf("unsupported naming trivia predicate: %s", rule.Condition.Predicate)
	}
}

// Checks naming trivia for a resource, returning whether the naming trivia applies and the naming trivia.
func CheckNamingTrivia(sdkName string, props resource.PropertyMap, spec *NamingTriviaSpec) (bool, *NamingTrivia, error) {
	if spec == nil {
		return false, nil, nil
	}
	namingTriviaApplies, err := CheckTriviaCondition(props, spec.Rule)
	if err != nil {
		return false, nil, err
	}
	namingTrivia := *spec.Trivia

	return namingTriviaApplies, &namingTrivia, nil
}

// Extracts naming trivia from a resource name, returning the name and the trivia.
func (n NamingTrivia) extractNamingTrivia(sdkName string, name string) string {
	canonicalName := name
	if len(name) > len(n.Suffix) && name[len(name)-len(n.Suffix):] == n.Suffix {
		canonicalName = name[:len(name)-len(n.Suffix)]
	}
	return canonicalName
}

// Applies structured trivia data to an existing resource name.
func (n NamingTrivia) ApplyTrivia(sdkName, name string) string {
	canonicalName := n.extractNamingTrivia(sdkName, name)
	return fmt.Sprintf("%s%s", canonicalName, n.Suffix)
}

func (trivia *NamingTrivia) Length() int {
	if trivia == nil {
		return 0
	}
	return len(trivia.Suffix)
}
