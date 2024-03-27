// Copyright 2016-2023, Pulumi Corporation.

package autonaming

import (
	"fmt"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func CheckTriviaCondition(props resource.PropertyMap, rule *metadata.NamingRule) (bool, error) {
	switch rule.Condition.Predicate {
	case metadata.NamingPredicateEquals:
		return props[rule.Field].DeepEqualsIncludeUnknowns(rule.Condition.Value), nil
	default:
		return false, fmt.Errorf("unsupported naming trivia predicate: %s", rule.Condition.Predicate)
	}
}

// Checks naming trivia for a resource, returning whether the naming trivia applies and the naming trivia.
func CheckNamingTrivia(sdkName string, props resource.PropertyMap, spec *metadata.NamingTriviaSpec) (bool, *metadata.NamingTrivia, error) {
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
func extractNamingTrivia(trivia *metadata.NamingTrivia, name string) string {
	canonicalName := name
	if len(name) > len(trivia.Suffix) && name[len(name)-len(trivia.Suffix):] == trivia.Suffix {
		canonicalName = name[:len(name)-len(trivia.Suffix)]
	}
	return canonicalName
}

// Applies structured trivia data to an existing resource name.
func ApplyTrivia(trivia *metadata.NamingTrivia, name string) string {
	canonicalName := extractNamingTrivia(trivia, name)
	return fmt.Sprintf("%s%s", canonicalName, trivia.Suffix)
}
