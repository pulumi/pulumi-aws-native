// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as awsnative from "@pulumi/aws-native";

const config = new pulumi.Config();
const roleToAssumeARN = config.require("roleToAssumeARN");

const provider = new awsnative.Provider("privileged", {
    assumeRole: {
        roleArn: roleToAssumeARN,
        sessionName: "PulumiSession",
        externalId: "PulumiApplication",
    },
    region: aws.config.requireRegion(),
});

const name = new random.RandomString("name", {
    length: 8,
    special: false,
    upper: false,
});

const logGroup = new awsnative.logs.LogGroup("log-test", {
    logGroupName: name.result,
    retentionInDays: 90,
}, {provider});

export const arn = logGroup.arn;
