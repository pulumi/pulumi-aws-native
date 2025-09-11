// Copyright 2016-2024, Pulumi Corporation.

import * as pulumi from '@pulumi/pulumi';
import * as aws from "@pulumi/aws-native";
import * as awsClassic from "@pulumi/aws";

const amiRegion = new pulumi.Config().require("amiRegion");

// Create an IAM role for the Lambda function
const lambdaRole = new awsClassic.iam.Role("lambdaRole", {
    assumeRolePolicy: awsClassic.iam.assumeRolePolicyForPrincipal({ Service: "lambda.amazonaws.com" }),
});

const policy = new awsClassic.iam.Policy("lambdaPolicy", {
    policy: {
        Version: "2012-10-17",
        Statement: [{
            Action: "ec2:DescribeImages",
            Effect: "Allow",
            Resource: "*",
        }],
    },
});

const rpa1 = new awsClassic.iam.RolePolicyAttachment("lambdaRolePolicyAttachment1", {
    role: lambdaRole.name,
    policyArn: policy.arn,
});

const rpa2 = new awsClassic.iam.RolePolicyAttachment("lambdaRolePolicyAttachment2", {
    role: lambdaRole.name,
    policyArn: awsClassic.iam.ManagedPolicy.AWSLambdaBasicExecutionRole,
});

const bucket = new awsClassic.s3.Bucket('custom-resource-emulator', {
    forceDestroy: true,
});

const handlerCode = new awsClassic.s3.BucketObjectv2("handler-code", {
    bucket: bucket.bucket,
    key: "handlerCode",
    source: new pulumi.asset.AssetArchive({
        "index.js": new pulumi.asset.FileAsset("ami-lookup.js"),
    })
})

// Create the Lambda function for the custom resource
const lambdaFunction = new awsClassic.lambda.Function("ami-lookup-custom-resource", {
    runtime: awsClassic.types.enums.lambda.Runtime.NodeJS20dX,
    s3Bucket: bucket.bucket,
    s3Key: handlerCode.key,
    handler: "index.handler",
    role: lambdaRole.arn,
    memorySize: 128,
    timeout: 30,
}, { dependsOn: [rpa1, rpa2] });

const cfnCustomResource = new aws.cloudformation.CustomResourceEmulator('emulator', {
    bucketName: bucket.id,
    bucketKeyPrefix: 'custom-resource-emulator',
    customResourceProperties: {
        Region: amiRegion,
        Architecture: 'HVM64',
    },
    serviceToken: lambdaFunction.arn,
    resourceType: 'Custom::MyResource',
}, { customTimeouts: { create: '5m', update: '5m', delete: '5m' } });

const cloudformationStack = new awsClassic.cloudformation.Stack('stack', {
    templateBody: pulumi.interpolate`{
  "AWSTemplateFormatVersion" : "2010-09-09",

  "Description" : "AWS CloudFormation AMI Look Up Sample Template: Demonstrates how to dynamically specify an AMI ID. This template provisions an EC2 instance with an AMI ID that is based on the instance's type and region. **WARNING** This template creates an Amazon EC2 instance. You will be billed for the AWS resources used if you create a stack from this template.",

  "Resources" : {
    "AMIInfo": {
      "Type": "Custom::AMIInfo",
      "Properties": {
        "ServiceToken": "${lambdaFunction.arn}",
        "ServiceTimeout": 300,
        "Region": "${amiRegion}",
        "Architecture": "HVM64"
      }
    }
  },

  "Outputs" : {
    "AMIID" : {
      "Description": "The Amazon EC2 instance AMI ID.",
      "Value" : { "Fn::GetAtt": [ "AMIInfo", "Id" ] }
    }
  }
}
`
});

export const cloudformationAmiId = cloudformationStack.outputs['AMIID'];
export const emulatorAmiId = cfnCustomResource.data['Id'];
