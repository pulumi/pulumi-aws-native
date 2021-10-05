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
	awslogs "github.com/pulumi/pulumi-aws-native/sdk/go/aws/logs"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		name, err := random.NewRandomString(ctx, "name", &random.RandomStringArgs{
			Length:  pulumi.Int(8),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		logGroup, err := awslogs.NewLogGroup(ctx, "log-test", &awslogs.LogGroupArgs{
			LogGroupName:    name.Result,
			RetentionInDays: pulumi.Int(90),
		})
		if err != nil {
			return err
		}

		ctx.Export("arn", logGroup.Arn)
		return nil
	})
}
