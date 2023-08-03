// Copyright 2016-2021, Pulumi Corporation.

import * as aws from "@pulumi/aws-native";

const logGroup = new aws.logs.LogGroup("log-test", {
    retentionInDays: 90,
});

export const arn = logGroup.arn;

new aws.wafv2.WebAcl("acl", {
    scope: aws.types.enums.wafv2.WebAclScope.Regional,
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
