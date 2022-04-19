// Copyright 2016-2022, Pulumi Corporation.
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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPartition(t *testing.T) {
	cases := []struct {
		region      string
		partitionID string
	}{
		{
			region:      "us-west-2",
			partitionID: "aws",
		},
		{
			region:      "us-east-1",
			partitionID: "aws",
		},
		{
			region:      "cn-north-1",
			partitionID: "aws-cn",
		},
		{
			region:      "us-gov-west-1",
			partitionID: "aws-us-gov",
		},
		{
			region:      "us-iso-east-1",
			partitionID: "aws-iso",
		},
		{
			region:      "us-isob-east-1",
			partitionID: "aws-iso-b",
		},
	}
	for _, c := range cases {
		p := getPartition(c.region)
		assert.Equal(t, c.partitionID, p.ID)
	}
}
