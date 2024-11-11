The Custom Resource Emulator allows you to use AWS CloudFormation Custom Resources directly in your Pulumi programs. It provides a way to invoke AWS Lambda functions that implement custom provisioning logic following the CloudFormation Custom Resource protocol.

> **Note**: Currently, only Lambda-backed Custom Resources are supported. SNS-backed Custom Resources are not supported at this time.

## Example Usage

```typescript
import * as aws from "@pulumi/aws-native";

const bucket = new aws.s3.Bucket('custom-resource-emulator');

// Create a Custom Resource that invokes a Lambda function
const cr = new aws.cloudformation.CustomResourceEmulator('cr', {
    bucketName: bucket.id,
    bucketKeyPrefix: 'custom-resource-emulator',
    customResourceProperties: {
        hello: "world"
    },
    serviceToken: "arn:aws:lambda:us-west-2:123456789012:function:my-custom-resource",
    resourceType: 'Custom::MyResource',
}, { customTimeouts: { create: '5m', update: '5m', delete: '5m' } });

// Access the response data
export const customResourceData = customResource.data;
```

A full example of creating a CloudFormation Custom Resource Lambda function and using it in Pulumi can be found [here](https://github.com/pulumi/pulumi-aws-native/tree/master/examples/cfn-custom-resource).

## About CloudFormation Custom Resources

CloudFormation Custom Resources allow you to write custom provisioning logic for resources that aren't directly available as AWS CloudFormation resource types. Common use cases include:

- Implementing complex provisioning logic
- Performing custom validations or transformations
- Integrating with third-party services
- Implementing organization-specific infrastructure patterns

For more information about CloudFormation Custom Resources, see [Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html) in the AWS CloudFormation User Guide.

## Permissions

The IAM principal used by your Pulumi program must have the following permissions:

1. `lambda:InvokeFunction` on the Lambda function specified in `serviceToken`
2. S3 permissions on the bucket specified in `bucketName`:
   - `s3:PutObject`
   - `s3:GetObject`
   - `s3:HeadObject`

## Lambda Function Requirements

The Lambda function specified in `serviceToken` must implement the CloudFormation Custom Resource lifecycle.
For detailed information about implementing Lambda-backed Custom Resources, see [AWS Lambda-backed Custom Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources-lambda.html) in the AWS CloudFormation User Guide.

## Timeouts

Custom Resources have a default timeout of 60 minutes, matching the CloudFormation timeout for custom resource operations. You can customize it using the [`customTimeouts`](https://www.pulumi.com/docs/iac/concepts/options/customtimeouts/) resource option.
