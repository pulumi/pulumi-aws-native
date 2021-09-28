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

// Note: These types must match the types defined in the Go SDK (sdk/aws/go/pulumiTypes.go). Copying the types avoids
// having the provider depend on the SDK.

// The configuration for a Provider to assume a role.
type ProviderAssumeRole struct {
	// Number of seconds to restrict the assume role session duration.
	DurationSeconds *int `json:"durationSeconds"`
	// External identifier to use when assuming the role.
	ExternalId *string `json:"externalId"`
	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	Policy *string `json:"policy"`
	// Set of Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the role.
	PolicyArns []string `json:"policyArns"`
	// Amazon Resource Name (ARN) of the IAM Role to assume.
	RoleArn *string `json:"roleArn"`
	// Session name to use when assuming the role.
	SessionName *string `json:"sessionName"`
	// Map of assume role session tags.
	Tags map[string]string `json:"tags"`
	// A list of keys for session tags that you want to set as transitive. If you set a tag key as transitive, the corresponding key and value passes to subsequent sessions in a role chain.
	TransitiveTagKeys []string `json:"transitiveTagKeys"`
}
