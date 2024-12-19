// Copyright 2024, Pulumi Corporation.

package autonaming

import (
	"fmt"
	"math"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type AutoNamingConfig struct {
	AutoTrim              bool `json:"autoTrim"`
	RandomSuffixMinLength int  `json:"randomSuffixMinLength"`
}

// EngineAutonamingConfiguration contains autonaming parameters passed to the provider from the engine.
type EngineAutonamingConfiguration struct {
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
	engineAutonaming EngineAutonamingConfiguration,
	olds,
	news resource.PropertyMap,
	config *AutoNamingConfig,
) error {
	if spec == nil {
		return nil
	}
	// Auto-name fields if not already specified
	val, err := getDefaultName(urn, engineAutonaming, spec, olds, news, config)
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
	engineAutonaming EngineAutonamingConfiguration,
	autoNamingSpec *metadata.AutoNamingSpec,
	olds,
	news resource.PropertyMap,
	config *AutoNamingConfig,
) (*resource.PropertyValue, error) {
	sdkName := autoNamingSpec.SdkName

	// Prefer explicitly specified name
	if v, ok := news[resource.PropertyKey(sdkName)]; ok {
		return &v, nil
	}

	// Fallback to previous name if specified/set.
	if v, ok := olds[resource.PropertyKey(sdkName)]; ok {
		return &v, nil
	}

	// Generate naming trivia for the resource.
	namingTriviaApplies, namingTrivia, err := CheckNamingTrivia(sdkName, news, autoNamingSpec.TriviaSpec)
	if err != nil {
		return nil, err
	}

	if engineAutonaming.AutonamingMode != nil {
		//panic("yes")
		switch *engineAutonaming.AutonamingMode {
		case EngineAutonamingModeDisable:
			return nil, nil
		case EngineAutonamingModeEnforce:
			v := resource.NewStringProperty(engineAutonaming.ProposedName)
			return &v, nil
		case EngineAutonamingModePropose:
			proposedName := engineAutonaming.ProposedName

			// Apply naming trivia to the generated name.
			if namingTriviaApplies {
				proposedName = ApplyTrivia(namingTrivia, proposedName)
			}

			// Validate the proposed name against the length constraints.
			if autoNamingSpec.MaxLength > 0 && len(proposedName) > autoNamingSpec.MaxLength {
				return nil, fmt.Errorf("proposed name %q exceeds max length of %d", proposedName, autoNamingSpec.MaxLength)
			}
			if autoNamingSpec.MinLength > 0 && len(proposedName) < autoNamingSpec.MinLength {
				return nil, fmt.Errorf("proposed name %q is shorter than min length of %d", proposedName, autoNamingSpec.MinLength)
			}

			v := resource.NewStringProperty(proposedName)
			return &v, nil
		}
	}

	var autoTrim bool
	// resource.NewUniqueName does not allow for a random suffix shorter than 1.
	randomSuffixMinLength := 1
	if config != nil {
		autoTrim = config.AutoTrim
		if config.RandomSuffixMinLength != 0 {
			randomSuffixMinLength = config.RandomSuffixMinLength
		}
	}

	// Generate random name that fits the length constraints.
	name := urn.Name()
	prefix := name + "-"
	randLength := 7 // Default value
	if randomSuffixMinLength > randLength {
		randLength = randomSuffixMinLength
	}
	if len(prefix)+namingTrivia.Length()+randLength < autoNamingSpec.MinLength {
		randLength = autoNamingSpec.MinLength - len(prefix) - namingTrivia.Length()
	}

	maxLength := 0
	if autoNamingSpec.MaxLength > 0 {
		left := autoNamingSpec.MaxLength - len(prefix) - namingTrivia.Length()

		if left <= 0 && autoTrim {
			autoTrimMaxLength := autoNamingSpec.MaxLength - namingTrivia.Length() - randomSuffixMinLength
			if autoTrimMaxLength <= 0 {
				return nil, fmt.Errorf("failed to auto-generate value for %[1]q."+
					" Prefix: %[2]q is too large to fix max length constraint of %[3]d"+
					" with required suffix length %[4]d. Please provide a value for %[1]q",
					sdkName, prefix, autoNamingSpec.MaxLength, randomSuffixMinLength)
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
					sdkName, prefix, autoNamingSpec.MaxLength, namingTrivia.Suffix)
			} else {
				return nil, fmt.Errorf("failed to auto-generate value for %[1]q."+
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
	random, err := resource.NewUniqueName(engineAutonaming.RandomSeed, prefix, randLength, maxLength, nil)
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
