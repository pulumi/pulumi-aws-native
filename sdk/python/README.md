# Native AWS Pulumi Provider (preview)

The native AWS Provider for Pulumi lets you use AWS resources in your cloud programs.
This provider uses the AWS Cloud API and therefore provides full access to AWS CloudFormation resources.

The provider is currently in private preview.

To use this package, please [install the Pulumi CLI first](https://pulumi.com/).

## Configuring credentials

Credentials configuration is compatible with the classic AWS provider.

Please refer to [this quickstart guide](
https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/) for possible configuration options.

## Building

### Dependencies

- Go 1.16
- NodeJS 10.X.X or later
- Python 3.6 or later
- .NET Core 3.1

Please refer to [Contributing to Pulumi](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for installation
guidance.

### Building locally

Run the following commands to install Go modules, generate all SDKs, and build the provider:

```
$ make ensure
$ make build
```

Add the `bin` folder to your `$PATH` or copy the `bin/pulumi-resource-aws-native` file to another location in your `$PATH`.

### Running an example

Navigate to the ECS example and run Pulumi:

```
$ cd ./exampes/ecs
$ yarn link @pulumi/aws-native
$ pulumi config set aws:region us-west-2
$ pulumi config set aws-native:region us-west-2
$ pulumi up
``` 
