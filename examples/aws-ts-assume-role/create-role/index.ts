// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const unprivilegedUsername = config.require("unprivilegedUsername");

const unprivilegedUser = new aws.iam.User("unprivileged-user", {
    name: unprivilegedUsername,
});

const unprivilegedUserCreds = new aws.iam.AccessKey("unprivileged-user-key", {
    user: unprivilegedUser.name,
}, 
// additional_secret_outputs specify properties that must be encrypted as secrets
// https://www.pulumi.com/docs/intro/concepts/resources/#additionalsecretoutputs
{ additionalSecretOutputs: ["secret"] });

const allowLogGroupManagementRole = new aws.iam.Role("allow-loggroup-management", {
    description: "Allow management of LogGroups",
    assumeRolePolicy: unprivilegedUser.arn.apply(arn => {
        return aws.iam.assumeRolePolicyForPrincipal({AWS: arn});
    }),
});

const policy = new aws.iam.RolePolicy("allow-loggroup-management-policy", {
    role: allowLogGroupManagementRole,
    policy: {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:DeleteLogGroup",
                    "logs:DeleteLogStream",
                    "logs:DescribeLogGroups",
                    "logs:DescribeLogStreams",
                    "logs:PutLogEvents",
                    "logs:PutRetentionPolicy",
                ],
                "Resource": [
                    "arn:aws:logs:*:*:*"
                ]
            },
            {
                "Effect": "Allow",
                "Action": [
                    "cloudformation:CreateResource",
                    "cloudformation:DeleteResource",
                    "cloudformation:GetResource",
                    "cloudformation:GetResourceRequestStatus",
                ],
                "Resource": [
                    "*"
                ]
            }
        ]
    },
}, {parent: allowLogGroupManagementRole});

export const roleArn = allowLogGroupManagementRole.arn;
export const accessKeyId = unprivilegedUserCreds.id;
export const secretAccessKey = unprivilegedUserCreds.secret;
