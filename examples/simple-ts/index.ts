// Copyright 2016-2021, Pulumi Corporation.

import * as aws from "@pulumipreview/aws-native";
import * as random from "@pulumi/random";

const name = new random.RandomString("name", {
    length: 8,
    special: false,
    upper: false,
});

const logGroup = new aws.logs.LogGroup("logtest", {
    logGroupName: name.result,
    retentionInDays: 90,
});

export const arn = logGroup.arn;
