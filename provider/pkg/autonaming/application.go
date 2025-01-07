// Copyright 2024, Pulumi Corporation.

package autonaming

import (
	"fmt"
	"math"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// ProviderAutoNamingConfig contains autonaming parameters configured for the provider.
type ProviderAutoNamingConfig struct {
	AutoTrim              bool `json:"autoTrim"`
	RandomSuffixMinLength int  `json:"randomSuffixMinLength"`
}

// EngineAutoNamingConfig contains autonaming parameters passed to the provider from the engine.
type EngineAutoNamingConfig struct {
	RandomSeed     []byte
	AutonamingMode *EngineAutonamingMode
	ProposedName   string
}

// EngineAutonamingMode is the mode of autonaming to apply to the resource.
type EngineAutonamingMode int32

const (
	EngineAutonamingModePropose EngineAutonamingMode = iota
	EngineAutonamingModeEnforce
	EngineAutonamingModeDisable
)

func ApplyAutoNaming(
	spec *metadata.AutoNamingSpec,
	urn resource.URN,
	engineConfig EngineAutoNamingConfig,
	providerConfig *ProviderAutoNamingConfig,
	olds,
	news resource.PropertyMap,
) error {
	if spec == nil {
		return nil
	}
	// Auto-name fields if not already specified
	val, err := getDefaultName(urn, engineConfig, providerConfig, spec, olds, news)
	if err != nil {
		return err
	}
	if val != nil {
		news[resource.PropertyKey(spec.SdkName)] = *val
	}
	return nil
}

// getDefaultName retrieves either the explicitly specified name in inputs,
// or the equivalent in the old values. If neither is specified, it generates
// a random name for a resource's autonameable fields
// based on its URN name, It ensures the name meets the length constraints, if known.
// Defaults to the name followed by 7 random hex characters separated by a '-'.
func getDefaultName(
	urn resource.URN,
	engineConfig EngineAutoNamingConfig,
	providerConfig *ProviderAutoNamingConfig,
	propertySpec *metadata.AutoNamingSpec,
	olds,
	news resource.PropertyMap,
) (*resource.PropertyValue, error) {
	sdkName := propertySpec.SdkName

	// Prefer explicitly specified name
	if v, ok := news[resource.PropertyKey(sdkName)]; ok {
		return &v, nil
	}

	// Fallback to previous name if specified/set.
	if v, ok := olds[resource.PropertyKey(sdkName)]; ok {
		return &v, nil
	}

	// Generate naming trivia for the resource.
	namingTriviaApplies, namingTrivia, err := CheckNamingTrivia(sdkName, news, propertySpec.TriviaSpec)
	if err != nil {
		return nil, err
	}

	if engineConfig.AutonamingMode != nil {
		// Engine autonaming conflicts with provider autonaming. If both are specified, return an error.
		if providerConfig != nil {
			return nil, fmt.Errorf("pulumi:autonaming conflicts with provider autonaming configuration, please specify only one")
		}

		switch *engineConfig.AutonamingMode {
		case EngineAutonamingModeDisable:
			return nil, nil
		case EngineAutonamingModeEnforce:
			v := resource.NewStringProperty(engineConfig.ProposedName)
			return &v, nil
		case EngineAutonamingModePropose:
			proposedName := engineConfig.ProposedName

			// Apply naming trivia to the generated name.
			if namingTriviaApplies {
				proposedName = ApplyTrivia(namingTrivia, proposedName)
			}

			// Validate the proposed name against the length constraints.
			if propertySpec.MaxLength > 0 && len(proposedName) > propertySpec.MaxLength {
				return nil, fmt.Errorf("proposed name %q exceeds max length of %d", proposedName, propertySpec.MaxLength)
			}
			if propertySpec.MinLength > 0 && len(proposedName) < propertySpec.MinLength {
				return nil, fmt.Errorf("proposed name %q is shorter than min length of %d", proposedName, propertySpec.MinLength)
			}

			v := resource.NewStringProperty(proposedName)
			return &v, nil
		}
	}

	var autoTrim bool
	// resource.NewUniqueName does not allow for a random suffix shorter than 1.
	randomSuffixMinLength := 1
	if providerConfig != nil {
		autoTrim = providerConfig.AutoTrim
		if providerConfig.RandomSuffixMinLength != 0 {
			randomSuffixMinLength = providerConfig.RandomSuffixMinLength
		}
	}

	// Generate random name that fits the length constraints.
	name := urn.Name()
	prefix := name + "-"
	randLength := 7 // Default value
	if randomSuffixMinLength > randLength {
		randLength = randomSuffixMinLength
	}
	if len(prefix)+namingTrivia.Length()+randLength < propertySpec.MinLength {
		randLength = propertySpec.MinLength - len(prefix) - namingTrivia.Length()
	}

	maxLength := 0
	if propertySpec.MaxLength > 0 {
		left := propertySpec.MaxLength - len(prefix) - namingTrivia.Length()

		if left <= 0 && autoTrim {
			autoTrimMaxLength := propertySpec.MaxLength - namingTrivia.Length() - randomSuffixMinLength
			if autoTrimMaxLength <= 0 {
				return nil, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d"+
					" with required suffix length %[4]d. Please provide a value for %[1]q",
					sdkName, prefix, propertySpec.MaxLength, randomSuffixMinLength)
			}
			prefix = trimName(prefix, autoTrimMaxLength)
			randLength = randomSuffixMinLength
			left = randomSuffixMinLength
		}

		if left <= 0 {
			if namingTrivia.Length() > 0 {
				return nil, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d"+
					" with required suffix %[4]q. Please provide a value for %[1]q",
					sdkName, prefix, propertySpec.MaxLength, namingTrivia.Suffix)
			} else {
				return nil, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d. Please provide a value for %[1]q",
					sdkName, prefix, propertySpec.MaxLength)
			}
		}
		if left < randLength {
			randLength = left
		}
		maxLength = len(prefix) + left
	}

	// Resource name is URN name + "-" + random suffix.
	random, err := resource.NewUniqueName(engineConfig.RandomSeed, prefix, randLength, maxLength, nil)
	if err != nil {
		return nil, err
	}

	// Apply naming trivia to the generated name.
	if namingTriviaApplies {
		random = ApplyTrivia(namingTrivia, random)
	}

	v := resource.NewStringProperty(random)
	return &v, nil
}

// trimName will trim the prefix to fit within the max length constraint.
// It will cut out part of the middle, keeping the beginning and the end of the string.
// This is so that really long generated names can still be unique. For example, if the
// user creates a resource name by appending the parent onto the child, you could end up
// with names like:
// - "topParent-middleParent-bottonParent-child1"
// - "topParent-middleParent-bottonParent-child2"
//
// If the max length is 30, the trimmed generated name for both would be something like
// "topParent-middleParent-bottonP" and you would not be able to tell them apart.
//
// By trimming from the middle you would end up with something like this, preserving
// the uniqueness of the generated names:
// - "topParent-middlonParent-child1"
// - "topParent-middlonParent-child2"
func trimName(name string, maxLength int) string {
	floorHalf := math.Floor(float64(maxLength) / 2)
	half := int(floorHalf)
	left := maxLength - half
	return name[0:half] + name[len(name)-left:]
}
