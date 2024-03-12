// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws-native";

const logGroup = new aws.s3.Bucket("bucket", {
  tags: [
    {
      key: "foo",
      value: "buzz", // <-- this value has changed
    },
    {
      key: "secretfoo",
      value: pulumi.secret("secretbuzz"), // <-- this value has changed
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
  // Remove policies to avoid deletion issue: https://github.com/pulumi/pulumi-aws-native/issues/416
});

const func = new aws.lambda.Function("function", {
  runtime: "nodejs16.x",
  role: lambdaRole.arn,
  handler: "index.handler",
  code: {
    zipFile: `
  exports.handler =  async function(event, context) {
    console.log("Running")
    return "bar"  // <-- this value has changed
  }
  `,
  },
});
