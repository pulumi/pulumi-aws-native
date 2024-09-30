[![Deploy](https://get.pulumi.com/new/button.svg)](https://app.pulumi.com/new?template=https://github.com/pulumi/pulumi-aws-native/blob/master/examples/aws-native-ts-stepfunctions/README.md)

# AWS Step Functions

A basic example that demonstrates using AWS Step Functions with a Lambda function using the AWS Cloud Control provider.

```
# Create and configure a new stack
$ pulumi stack init stepfunctions-dev
$ export AWS_REGION=us-east-2

# Install dependencies
$ npm install

# Preview and run the deployment
$ pulumi up

# Start execution using the AWS CLI (or from the console at https://console.aws.amazon.com/states)
$ aws stepfunctions start-sync-execution --state-machine-arn $(pulumi stack output stateMachineArn)

# Remove the app
$ pulumi destroy
```
