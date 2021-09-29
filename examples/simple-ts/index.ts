// Copyright 2016-2021, Pulumi Corporation.

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
