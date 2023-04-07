// Copyright 2016-2023, Pulumi Corporation.

import * as aws from "@pulumi/aws-native";

for (let i = 0; i < 100; i++) {
    new aws.iam.Role(`role-${i}`, {
        assumeRolePolicyDocument: {
            Version: "2012-10-17",
            Statement: [{
                Effect: "Allow",
                Principal: {
                    Service: ["grafana.amazonaws.com"],
                },
                Action: ["sts:AssumeRole"],
            }],
        },
    });
}
