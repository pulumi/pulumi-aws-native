// Copyright 2016-2021, Pulumi Corporation.

import * as awsClassic from "@pulumi/aws";
import * as aws from "@pulumi/aws-native";

const logGroup = new aws.s3.Bucket("bucket", {
    tags: [{
        key: "foo",
        value: "buzz", // <-- this value has changed
    }]
});

const lambdaRole = new awsClassic.iam.Role("lambdaRole", {
    assumeRolePolicy: awsClassic.iam.assumeRolePolicyForPrincipal({ Service: "lambda.amazonaws.com" }),
    inlinePolicies: [{
        name: "lambdaAccess", policy: JSON.stringify({
            Version: "2012-10-17",
            Statement: [{
                Effect: "Allow",
                Action: [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents",
                ],
                Resource: "arn:aws:logs:*:*:*",
            }],
        })
    }]
});

const func = new aws.lambda.Function("function", {
    runtime: "nodejs16.x",
    role: lambdaRole.arn,
    handler: "index.handler",
    code: {
        zipFile: `
exports.handler =  async function(event, context) {
  console.log("Running")
  return "bar" // <-- this value has changed
}
`,
    },
});
