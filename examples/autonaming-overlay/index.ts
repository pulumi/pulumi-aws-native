import * as pulumi from '@pulumi/pulumi';
import * as aws from "@pulumi/aws-native";

const config = new pulumi.Config();
const name = config.require('roleName');
const role = new aws.iam.Role(name, {
  assumeRolePolicyDocument: {
    Version: "2012-10-17",
    Statement: [
      {
        Effect: "Allow",
        Principal: {
          Service: "lambda.amazonaws.com",
        },
        Action: "sts:AssumeRole",
      },
    ],
  },
});

export const roleName = role.roleName;
