// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";

const bucket = new aws.s3.Bucket("bucket", {
  tags: [
    {
      key: "foo",
      value: "bar",
    },
    {
      key: "secretfoo",
      value: pulumi.secret("secretbar"),
    },
  ],
});

const lambdaRole = new aws.iam.Role("lambdaRole", {
  assumeRolePolicyDocument: {
    Version: "2012-10-17",
    Statement: [
      {
        Sid: "AllowAssumeRole",
        Effect: "Allow",
        Principal: {
          Service: "lambda.amazonaws.com",
        },
        Action: "sts:AssumeRole",
      },
    ],
  },
  policies: [
    {
      policyName: "lambdaAccess",
      policyDocument: JSON.stringify({
        Version: "2012-10-17",
        Statement: [
          {
            Effect: "Allow",
            Action: [
              "logs:CreateLogGroup",
              "logs:CreateLogStream",
              "logs:PutLogEvents",
            ],
            Resource: "arn:aws:logs:*:*:*",
          },
        ],
      }),
    },
  ],
});

const func = new aws.lambda.Function("function", {
  runtime: "nodejs16.x",
  role: lambdaRole.arn,
  handler: "index.handler",
  code: {
    zipFile: `
exports.handler =  async function(event, context) {
  console.log("Running")
  return "foo"
}
`,
  },
});
