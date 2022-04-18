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
	"regexp"
)

type partition struct {
	ID          string
	URLSuffix   string
	RegionRegex *regexp.Regexp
}

var partitions = []partition{
	{
		ID:          "aws",
		URLSuffix:   "amazonaws.com",
		RegionRegex: regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	},
	{
		ID:          "aws-cn",
		URLSuffix:   "amazonaws.com.cn",
		RegionRegex: regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	},
	{
		ID:          "aws-us-gov",
		URLSuffix:   "amazonaws.com",
		RegionRegex: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
	},
	{
		ID:          "aws-iso",
		URLSuffix:   "c2s.ic.gov",
		RegionRegex: regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	},
	{
		ID:          "aws-iso-b",
		URLSuffix:   "sc2s.sgov.gov",
		RegionRegex: regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	},
}

func getPartition(region string) partition {
	for _, p := range partitions {
		if p.RegionRegex.MatchString(region) {
			return p
		}
	}
	return partitions[0]
}
