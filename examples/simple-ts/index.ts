/*
 * Copyright 2016-2021, Pulumi Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as aws from "@pulumi/aws-native";
import * as random from "@pulumi/random";

const name = new random.RandomString("name", {
    length: 8,
    special: false,
    upper: false,
});

const logGroup = new aws.logs.LogGroup("log-test", {
    logGroupName: name.result,
    retentionInDays: 90,
});

export const arn = logGroup.arn;

new aws.wafv2.WebACL("acl", { 
    scope: aws.types.enums.wafv2.WebACLScope.Regional, 
    defaultAction: { block: {} }, 
    rules: [{ 
        name: "AWS-AWSManagedRulesCommonRuleSet", 
        priority: 1, 
        statement: { 
            managedRuleGroupStatement: { 
                vendorName: "AWS", 
                name: "AWSManagedRulesCommonRuleSet" 
            } 
        }, 
        overrideAction: { none: {} }, 
        visibilityConfig: { 
            sampledRequestsEnabled: true, 
            cloudWatchMetricsEnabled: true, 
            metricName: "AWSManagedRulesCommonRuleSet" 
        } 
    }], 
    visibilityConfig: { 
        cloudWatchMetricsEnabled: true, 
        metricName: "frontend", 
        sampledRequestsEnabled: true 
    },
});
