# AWS CloudFormation Custom Resource Emulator

This example demonstrates how to use AWS CloudFormation Custom Resources directly in your Pulumi programs using the Custom Resource Emulator. It shows how to invoke AWS Lambda functions that implement custom provisioning logic following the CloudFormation Custom Resource protocol.

## Overview

The Custom Resource Emulator allows you to seamlessly integrate existing CloudFormation Custom Resources into your Pulumi infrastructure code. This example illustrates the conversion of a CloudFormation Custom Resource implementation to its Pulumi equivalent, highlighting the similarities and differences between both approaches.

> **Note**: Currently, only Lambda-backed Custom Resources are supported. SNS-backed Custom Resources are not supported at this time.

## Components

The example consists of two main parts:

1. **Lambda Function**: A simple AWS Lambda function that implements the Custom Resource protocol to retrieve EC2 AMIs.
2. **Custom Resource**: The Pulumi resource that invokes the Lambda function using the Custom Resource Emulator.

## Code Structure

```
.
├── index.ts                # Pulumi program
├── ami-lookup.js           # The Lambda Function Code
├── Pulumi.yaml             # Pulumi project file
└── package.json            # Node.js package file
```

## Implementation Details

The example demonstrates:
- How to create a Lambda function that implements the Custom Resource protocol
- How to use the Custom Resource Emulator to invoke the Lambda function
- How to pass properties to the Custom Resource
- How to retrieve outputs from the Custom Resource


## Additional Resources

- [Pulumi AWS Native Provider Documentation](https://www.pulumi.com/docs/reference/pkg/aws-native/)
- [CloudFormation Custom Resource Documentation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html)
- [Custom Resource Protocol Specification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/crpg-ref.html)
