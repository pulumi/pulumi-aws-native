# Pulumi AWS Native Provider (preview)

The Pulumi AWS Native Provider enables you to build, deploy, and manage [any AWS resource that's supported by the AWS Cloud Control API](https://github.com/pulumi/pulumi-aws-native/blob/master/provider/cmd/pulumi-gen-aws-native/supported-types.txt).
With AWS Native, you get same-day access to all new AWS resources and all new properties on existing resources supported by the Cloud Control API.
You can use AWS Native from a Pulumi program written in any Pulumi language: C#, Go, JavaScript/TypeScript, and Python.
You'll need to [install and configure the Pulumi CLI](https://pulumi.com/docs/get-started/install) if you haven't already.

---
**NOTE**

AWS Native is in public preview.
Not all AWS resources are currently available in the AWS Native Provider. It covers all resources that are currently supported by the AWS Cloud Control API and AWS CloudFormation Registry. [List of supported resources](https://github.com/pulumi/pulumi-aws-native/blob/master/provider/cmd/pulumi-gen-aws-native/supported-types.txt).

For new projects, we recommend using AWS Native and [AWS Classic](https://github.com/pulumi/pulumi-aws) side-by-side so you can get the speed and correctness benefits of AWS Native where possible.
For existing projects, AWS Classic remains fully supported; at this time, we recommend waiting to migrate existing projects to AWS Native.

---

## Configuring credentials

To learn how to configure credentials refer to the [AWS configuration options](https://www.pulumi.com/registry/packages/aws-native/installation-configuration/#configuration-options).

## Building

### Dependencies

- Go 1.20
- NodeJS 10.X.X or later
- Yarn 1.22 or later
- Python 3.6 or later
- .NET Core 3.1
- Gradel 7
- Pulumi CLI and language plugins
- pulumictl

You can quickly launch a shell environment with all the required dependencies using
[devbox](https://www.jetpack.io/devbox/):
```
# Install devbox if needed
$ which devbox || curl -fsSL https://get.jetpack.io/devbox | bash
$ devbox shell
```

Alternatively, you can develop in a preconfigured container environment using 
[an editor or service that supports the devcontainer standard](https://containers.dev/supporting#editors)
such as [VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) or [Github Codespaces](https://codespaces.new/pulumi/pulumi-aws-native). Please note that building this project can be fairly memory intensive, if you are having trouble building in a container, please ensure you have at least 12GB of memory available for the container.

### Building locally

Run the following commands to install Go modules, generate all SDKs, and build the provider:

```
$ make ensure
$ make build
```

Add the `bin` folder to your `$PATH` or copy the `bin/pulumi-resource-aws-native` file to another location in your `$PATH`.

### Running tests

To run unittests, use:

```
$ make test_provider
```

### Running an example

Navigate to the ECS example and run Pulumi:

```
$ cd ./examples/ecs
$ yarn link @pulumi/aws-native
$ pulumi config set aws:region us-west-2
$ pulumi config set aws-native:region us-west-2
$ pulumi up
``` 

### Local Development

#### Additional Build Targets

`make build` can be a bit slow as it rebuilds the sdks for every language; 
you can use `make provider` or `make codegen` to just rebuild the provider plugin or codegen binaries

#### Debugging / Logging

Oftentimes, it can be informative to investigate the precise requests this provider makes to upstream AWS APIs. By default, the Pulumi CLI writes all of its logs to files rather than stdout or stderr (though this can be overridden with the `--logtostderr` flag). This works to our benefit, however, as the AWS SDK used in this provider writes to stderr by default. To view a trace of all HTTP requests and responses between this provider and AWS APIs, run the Pulumi CLI with the following arguments:

```shell
pulumi -v 9 --logflow [command]
```

this will correctly set verbosity to the level that the provider expects to log these requests (via `-v 9`), as well as flowing that verbosity setting down from the Pulumi CLI to the provider itself (via `--logflow`).
