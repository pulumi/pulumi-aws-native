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

// getDefaultName retrieves either the explicitly specified name in inputs,
// or the equivalent in the old values. If neither is specified, it generates
// a random name for a resource's autonameable fields
// based on its URN name, It ensures the name meets the length constraints, if known.
// Defaults to the name followed by 7 random hex characters separated by a '-'.
func getDefaultName(
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

	maxLength := 0
	if autoNamingSpec.MaxLength > 0 {
		left := autoNamingSpec.MaxLength - len(prefix)

		if left <= 0 {
			return resource.PropertyValue{}, fmt.Errorf("failed to auto-generate value for %[1]q."+
				" Prefix: %[2]q is too large to fix max length constraint of %[3]d. Please provide a value for %[1]q",
				sdkName, prefix, autoNamingSpec.MaxLength)
		}
		if left < randLength {
			randLength = left
		}
		maxLength = len(prefix) + left
	}

	// Resource name is URN name + "-" + random suffix.
	random, err := resource.NewUniqueHex(prefix, randLength, maxLength)
	if err != nil {
		return resource.PropertyValue{}, err
	}
	return resource.NewStringProperty(random), nil
}
