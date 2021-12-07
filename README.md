# Pulumi AWS Native Provider (preview)

The Pulumi AWS Native Provider enables you to build, deploy, and manage [any AWS resource that's supported by the AWS Cloud Control API](https://github.com/pulumi/pulumi-aws-native/blob/master/provider/cmd/pulumi-gen-aws-native/supported-types.txt).
With AWS Native, you get same-day access to all new AWS resources and all new properties on existing resources supported by the Cloud Control API.
You can use AWS Native from a Pulumi program written in any Pulumi language: C#, Go, JavaScript/TypeScript, and Python.
You'll need to [install and configure the Pulumi CLI](https://pulumi.com/docs/get-started/install) if you haven't already.

---
**NOTE**

AWS Native is in public preview.
Not all AWS resources are currently available in the AWS Native Provider. It covers all resources that are currently supported by the AWS Cloud Control API and AWS CloudFormation Registry. You can see the [list of supported resources here](https://github.com/pulumi/pulumi-aws-native/blob/master/provider/cmd/pulumi-gen-aws-native/supported-types.txt).

For new projects, we recommend using AWS Native and [AWS Classic](https://github.com/pulumi/pulumi-aws) side-by-side so you can get the speed and correctness benefits of AWS Native where possible.
For existing projects, AWS Classic remains fully supported; at this time, we recommend waiting to migrate existing projects to AWS Native.

---

## Configuring credentials

Credentials configuration is compatible with the classic AWS provider.

Please refer to [this quickstart guide](
https://www.pulumi.com/docs/intro/cloud-providers/aws-native/setup/) for possible configuration options.

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
$ cd ./examples/ecs
$ yarn link @pulumi/aws-native
$ pulumi config set aws:region us-west-2
$ pulumi config set aws-native:region us-west-2
$ pulumi up
``` 
