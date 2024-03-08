// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as awsNative from "@pulumi/aws-native";
import * as aws from "@pulumi/aws";

// export const layout = vpc.subnets;
export const lambdaStore = new awsNative.s3.Bucket("ls", {
  accessControl: "Private",
  versioningConfiguration: {
    status: "Enabled",
  },
  lifecycleConfiguration: {
    rules: [
      {
        id: "keepLastX",
        status: "Enabled",
        noncurrentVersionExpiration: {
          // keep last 10, even if a service isn't touched for a while
          newerNoncurrentVersions: 10,
          // keep anything from last 30 days
          noncurrentDays: 30,
        },
      },
    ],
  },
});

// file is a placeholder, overwritten by external deploys in the lambda
export const graphqlPublicZip = new aws.s3.BucketObject(
  "lambda-source",
  {
    bucket: lambdaStore.id,
    key: "samplefn",
    source: new pulumi.asset.AssetArchive({
      "index.js": new pulumi.asset.StringAsset(
        `module.exports = () => { console.log("placeholder"); return Promise.resolve(); }`
      ),
    }),
  },
  { ignoreChanges: ["source"] }
);

// file is a placeholder, overwritten by external deploys in the lambda
export const layerCode = new aws.s3.BucketObject(
  "layer-source",
  {
    bucket: lambdaStore.id,
    key: "samplefn",
    source: new pulumi.asset.AssetArchive({
      "index.js": new pulumi.asset.StringAsset(
        `module.exports = () => { console.log("placeholder"); return Promise.resolve(); }`
      ),
    }),
  },
  { ignoreChanges: ["source"] }
);

const graphqlPublicRole = new aws.iam.Role("lambda-role", {
  assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal(
    aws.iam.Principals.LambdaPrincipal
  ),
  managedPolicyArns: [
    /** Can execute lambdas and send logs to cloudwatch */
    aws.iam.ManagedPolicy.AWSLambdaBasicExecutionRole,
  ],
});

export const layer = new awsNative.lambda.LayerVersion("layer", {
  content: {
    s3Bucket: lambdaStore.id,
    s3Key: layerCode.key,
    s3ObjectVersion: layerCode.versionId,
  },
});

export const graphqlPublic = new awsNative.lambda.Function("sample", {
  functionName: "fnname",
  role: graphqlPublicRole.arn,
  memorySize: 128,
  architectures: ["arm64"],
  runtime: "nodejs16.x",
  handler: "index.handler",
  code: {
    s3Bucket: lambdaStore.id,
    s3Key: graphqlPublicZip.key,
    s3ObjectVersion: graphqlPublicZip.versionId,
  },
  layers: [layer.layerVersionArn],
});
