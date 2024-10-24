// Copyright 2024, Pulumi Corporation.

package autonaming

import (
	"fmt"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func ApplyAutoNaming(
	spec *metadata.AutoNamingSpec,
	urn resource.URN,
	randomSeed []byte,
	olds,
	news resource.PropertyMap,
) error {
	if spec == nil {
		return nil
	}
	// Auto-name fields if not already specified
	val, err := getDefaultName(randomSeed, urn, spec, olds, news)
	if err != nil {
		return err
	}
	news[resource.PropertyKey(spec.SdkName)] = val
	return nil
}

// getDefaultName retrieves either the explicitly specified name in inputs,
// or the equivalent in the old values. If neither is specified, it generates
// a random name for a resource's autonameable fields
// based on its URN name, It ensures the name meets the length constraints, if known.
// Defaults to the name followed by 7 random hex characters separated by a '-'.
func getDefaultName(
	randomSeed []byte,
	urn resource.URN,
	autoNamingSpec *metadata.AutoNamingSpec,
	olds,
	news resource.PropertyMap,
) (resource.PropertyValue, error) {
	sdkName := autoNamingSpec.SdkName

	// Prefer explicitly specified name
	if v, ok := news[resource.PropertyKey(sdkName)]; ok {
		return v, nil
	}

	// Fallback to previous name if specified/set.
	if v, ok := olds[resource.PropertyKey(sdkName)]; ok {
		return v, nil
	}

	// Generate naming trivia for the resource.
	namingTriviaApplies, namingTrivia, err := CheckNamingTrivia(sdkName, news, autoNamingSpec.TriviaSpec)
	if err != nil {
		return resource.PropertyValue{}, err
	}

	// Generate random name that fits the length constraints.
	name := urn.Name()
	prefix := name + "-"
	randLength := 7
	if len(prefix)+namingTrivia.Length()+randLength < autoNamingSpec.MinLength {
		randLength = autoNamingSpec.MinLength - len(prefix) - namingTrivia.Length()
	}

	maxLength := 0
	if autoNamingSpec.MaxLength > 0 {
		left := autoNamingSpec.MaxLength - len(prefix) - namingTrivia.Length()

		if left <= 0 {
			if namingTrivia.Length() > 0 {
				return resource.PropertyValue{}, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d"+
					" with required suffix %[4]q. Please provide a value for %[1]q",
					sdkName, prefix, autoNamingSpec.MaxLength, namingTrivia.Suffix)
			} else {
				return resource.PropertyValue{}, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d. Please provide a value for %[1]q",
					sdkName, prefix, autoNamingSpec.MaxLength)
			}
		}
		if left < randLength {
			randLength = left
		}
		maxLength = len(prefix) + left
	}

	// Resource name is URN name + "-" + random suffix.
	random, err := resource.NewUniqueName(randomSeed, prefix, randLength, maxLength, nil)
	if err != nil {
		return resource.PropertyValue{}, err
	}

	// Apply naming trivia to the generated name.
	if namingTriviaApplies {
		random = ApplyTrivia(namingTrivia, random)
	}

	return resource.NewStringProperty(random), nil
}
