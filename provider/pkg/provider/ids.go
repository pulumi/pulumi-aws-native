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
	"github.com/pulumi/pulumi-aws-native/provider/pkg/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// setDefaultName generates a random name for a resource's autonameable fields
// based on its URN name, It ensures the name meets the length constraints, if known.
// Defaults to the name followed by 7 random hex characters separated by a '-'.
func setDefaultName(
	urn resource.URN,
	resourceInfo schema.CloudAPIResource,
	olds,
	news resource.PropertyMap,
) bool {
	var mutated bool

	for sdkName, autoNamingSpec := range resourceInfo.AutoNamingSpecs {
		if _, ok := olds[resource.PropertyKey(sdkName)]; ok {
			continue
		}

		name := urn.Name().String()
		prefix := name+"-"
		randLength := 7
		if len(prefix) + randLength < autoNamingSpec.MinLength {
			randLength = autoNamingSpec.MinLength
		}

		// Resource name is URN name + "-" + random suffix.
		random, err := resource.NewUniqueHex(prefix, randLength, autoNamingSpec.MaxLength)
		contract.AssertNoError(err)
		news[resource.PropertyKey(sdkName)] = resource.NewStringProperty(random)
		mutated = true
	}
	return mutated
}
