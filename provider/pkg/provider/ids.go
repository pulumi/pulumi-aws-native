// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var NamingTriviaPropMap map[string]NamingTrivia = map[string]NamingTrivia{
	"queueName": {
		Suffix: ".fifo",
	},
}

// NamingTrivia is a struct which contains any derived information
// about a resource's name that is used to generate a random name.
type NamingTrivia struct {
	Suffix string
}

// Builds naming trivia for a resource based on its properties.
func createNamingTrivia(sdkName string, props resource.PropertyMap) NamingTrivia {
	var namingTrivia NamingTrivia
	switch sdkName {
	case "queueName":
		if fifoQueue, ok := props["fifoQueue"]; ok && fifoQueue.BoolValue() {
			namingTrivia = NamingTriviaPropMap[sdkName]
		}
	}
	return namingTrivia
}

// Extracts naming trivia from a resource name, returning the name and the trivia.
func extractNamingTrivia(sdkName string, name string) (string, NamingTrivia) {
	canonicalName := name
	namingTrivia := NamingTriviaPropMap[sdkName]
	if len(name) > len(namingTrivia.Suffix) && name[len(name)-len(namingTrivia.Suffix):] == namingTrivia.Suffix {
		canonicalName = name[:len(name)-len(namingTrivia.Suffix)]
	}
	return canonicalName, namingTrivia
}

// Applies structured trivia data to an existing resource name.
func (n NamingTrivia) applyTrivia(sdkName, name string) string {
	canonicalName, _ := extractNamingTrivia(sdkName, name)
	return fmt.Sprintf("%s%s", canonicalName, n.Suffix)
}

// getDefaultName retrieves either the explicitly specified name in inputs,
// or the equivalent in the old values. If neither is specified, it generates
// a random name for a resource's autonameable fields
// based on its URN name, It ensures the name meets the length constraints, if known.
// Defaults to the name followed by 7 random hex characters separated by a '-'.
func getDefaultName(
	randomSeed []byte,
	urn resource.URN,
	autoNamingSpec *schema.AutoNamingSpec,
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

	// Generate random name that fits the length constraints.
	name := urn.Name().String()
	prefix := name + "-"
	randLength := 7
	if len(prefix)+randLength < autoNamingSpec.MinLength {
		randLength = autoNamingSpec.MinLength - len(prefix)
	}

	// Generate naming trivia for the resource.
	namingTrivia := createNamingTrivia(sdkName, news)

	maxLength := 0
	if autoNamingSpec.MaxLength > 0 {
		left := autoNamingSpec.MaxLength - len(prefix) - len(namingTrivia.Suffix)

		if left <= 0 {
			if len(namingTrivia.Suffix) > 0 {
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
	random = namingTrivia.applyTrivia(sdkName, random)

	return resource.NewStringProperty(random), nil
}
