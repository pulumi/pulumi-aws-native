// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify
{
    public static class GetBranch
    {
        /// <summary>
        /// The AWS::Amplify::Branch resource creates a new branch within an app.
        /// </summary>
        public static Task<GetBranchResult> InvokeAsync(GetBranchArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBranchResult>("aws-native:amplify:getBranch", args ?? new GetBranchArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Amplify::Branch resource creates a new branch within an app.
        /// </summary>
        public static Output<GetBranchResult> Invoke(GetBranchInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBranchResult>("aws-native:amplify:getBranch", args ?? new GetBranchInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::Amplify::Branch resource creates a new branch within an app.
        /// </summary>
        public static Output<GetBranchResult> Invoke(GetBranchInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetBranchResult>("aws-native:amplify:getBranch", args ?? new GetBranchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBranchArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN for a branch, part of an Amplify App.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetBranchArgs()
        {
        }
        public static new GetBranchArgs Empty => new GetBranchArgs();
    }

    public sealed class GetBranchInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ARN for a branch, part of an Amplify App.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetBranchInvokeArgs()
        {
        }
        public static new GetBranchInvokeArgs Empty => new GetBranchInvokeArgs();
    }


    [OutputType]
    public sealed class GetBranchResult
    {
        /// <summary>
        /// ARN for a branch, part of an Amplify App.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The backend for a `Branch` of an Amplify app. Use for a backend created from an AWS CloudFormation stack.
        /// 
        /// This field is available to Amplify Gen 2 apps only. When you deploy an application with Amplify Gen 2, you provision the app's backend infrastructure using Typescript code.
        /// </summary>
        public readonly Outputs.BranchBackend? Backend;
        /// <summary>
        /// The build specification (build spec) for the branch.
        /// </summary>
        public readonly string? BuildSpec;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to assign to a branch of an SSR app. The SSR Compute role allows the Amplify Hosting compute service to securely access specific AWS resources based on the role's permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-SSR-compute-role.html) in the *Amplify User Guide* .
        /// </summary>
        public readonly string? ComputeRoleArn;
        /// <summary>
        /// The description for the branch that is part of an Amplify app.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Enables auto building for the branch.
        /// </summary>
        public readonly bool? EnableAutoBuild;
        /// <summary>
        /// Enables performance mode for the branch.
        /// 
        /// Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out.
        /// </summary>
        public readonly bool? EnablePerformanceMode;
        /// <summary>
        /// Specifies whether Amplify Hosting creates a preview for each pull request that is made for this branch. If this property is enabled, Amplify deploys your app to a unique preview URL after each pull request is opened. Development and QA teams can use this preview to test the pull request before it's merged into a production or integration branch.
        /// 
        /// To provide backend support for your preview, Amplify automatically provisions a temporary backend environment that it deletes when the pull request is closed. If you want to specify a dedicated backend environment for your previews, use the `PullRequestEnvironmentName` property.
        /// 
        /// For more information, see [Web Previews](https://docs.aws.amazon.com/amplify/latest/userguide/pr-previews.html) in the *AWS Amplify Hosting User Guide* .
        /// </summary>
        public readonly bool? EnablePullRequestPreview;
        /// <summary>
        /// Specifies whether the skew protection feature is enabled for the branch.
        /// 
        /// Deployment skew protection is available to Amplify applications to eliminate version skew issues between client and servers in web applications. When you apply skew protection to a branch, you can ensure that your clients always interact with the correct version of server-side assets, regardless of when a deployment occurs. For more information about skew protection, see [Skew protection for Amplify deployments](https://docs.aws.amazon.com/amplify/latest/userguide/skew-protection.html) in the *Amplify User Guide* .
        /// </summary>
        public readonly bool? EnableSkewProtection;
        /// <summary>
        /// The environment variables for the branch.
        /// </summary>
        public readonly ImmutableArray<Outputs.BranchEnvironmentVariable> EnvironmentVariables;
        /// <summary>
        /// The framework for the branch.
        /// </summary>
        public readonly string? Framework;
        /// <summary>
        /// If pull request previews are enabled for this branch, you can use this property to specify a dedicated backend environment for your previews. For example, you could specify an environment named `prod` , `test` , or `dev` that you initialized with the Amplify CLI and mapped to this branch.
        /// 
        /// To enable pull request previews, set the `EnablePullRequestPreview` property to `true` .
        /// 
        /// If you don't specify an environment, Amplify Hosting provides backend support for each preview by automatically provisioning a temporary backend environment. Amplify Hosting deletes this environment when the pull request is closed.
        /// 
        /// For more information about creating backend environments, see [Feature Branch Deployments and Team Workflows](https://docs.aws.amazon.com/amplify/latest/userguide/multi-environments.html) in the *AWS Amplify Hosting User Guide* .
        /// </summary>
        public readonly string? PullRequestEnvironmentName;
        /// <summary>
        /// Describes the current stage for the branch.
        /// </summary>
        public readonly Pulumi.AwsNative.Amplify.BranchStage? Stage;
        /// <summary>
        /// The tag for the branch.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetBranchResult(
            string? arn,

            Outputs.BranchBackend? backend,

            string? buildSpec,

            string? computeRoleArn,

            string? description,

            bool? enableAutoBuild,

            bool? enablePerformanceMode,

            bool? enablePullRequestPreview,

            bool? enableSkewProtection,

            ImmutableArray<Outputs.BranchEnvironmentVariable> environmentVariables,

            string? framework,

            string? pullRequestEnvironmentName,

            Pulumi.AwsNative.Amplify.BranchStage? stage,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Backend = backend;
            BuildSpec = buildSpec;
            ComputeRoleArn = computeRoleArn;
            Description = description;
            EnableAutoBuild = enableAutoBuild;
            EnablePerformanceMode = enablePerformanceMode;
            EnablePullRequestPreview = enablePullRequestPreview;
            EnableSkewProtection = enableSkewProtection;
            EnvironmentVariables = environmentVariables;
            Framework = framework;
            PullRequestEnvironmentName = pullRequestEnvironmentName;
            Stage = stage;
            Tags = tags;
        }
    }
}
