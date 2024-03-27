// Copyright 2016-2023, Pulumi Corporation.

package metadata

import (
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

func (trivia *NamingTrivia) Length() int {
	if trivia == nil {
		return 0
	}
	return len(trivia.Suffix)
}
