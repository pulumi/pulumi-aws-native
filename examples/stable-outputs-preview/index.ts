import * as pulumi from '@pulumi/pulumi';
import * as path from 'path';
import * as ccapi from "@pulumi/aws-native";
import * as aws from '@pulumi/aws';
import { zipDirectory } from './zip';

const config = new pulumi.Config();
const desc = config.get('lambdaDescription') ?? 'test lambda';
const role = new ccapi.iam.Role('lambda-role', {
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


const bucket = new ccapi.s3.Bucket('bucket');
const object = new aws.s3.BucketObjectv2(
  'object',
  {
    source: zipDirectory(path.join(__dirname, 'app'), 'handler.zip'),
    bucket: bucket.bucketName.apply(bucket => bucket!),
    key: 'handler.zip',
  },
);
const handler = new ccapi.lambda.Function('my-function', {
  code: {
    s3Bucket: bucket.bucketName.apply(bucket => bucket!),
    s3Key: object.key,
  },
  description: desc,
  runtime: 'nodejs20.x',
  handler: 'index.handler',
  role: role.arn,
});

new ccapi.lambda.Permission('chall-permission', {
  functionName: handler.arn,
  action: 'lambda:InvokeFunction',
  principal: 's3.amazonaws.com',
});
