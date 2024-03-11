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

package main

import (
	awsNative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	ssm "github.com/pulumi/pulumi-aws-native/sdk/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsNativeConfig := config.New(ctx, "aws-native")
		awsNativeProvider, err := awsNative.NewProvider(ctx, "aws-native", &awsNative.ProviderArgs{
			Region: pulumi.String(awsNativeConfig.Require("region")),
			DefaultTags: awsNative.ProviderDefaultTagsArgs{
				Tags: pulumi.StringMap{
					"defaultTag": pulumi.String("defaultTagValue"),
				},
			},
		})
		if err != nil {
			return err
		}
		logGroup, err := ssm.NewParameter(ctx, "log-test", &ssm.ParameterArgs{
			Type:  ssm.ParameterTypeString,
			Value: pulumi.String("test"),
			Tags: pulumi.StringMap{
				"localTag": pulumi.String("localTagValue"),
			},
		}, pulumi.Provider(awsNativeProvider))
		if err != nil {
			return err
		}

		ctx.Export("tags", logGroup.Tags)
		return nil
	})
}
