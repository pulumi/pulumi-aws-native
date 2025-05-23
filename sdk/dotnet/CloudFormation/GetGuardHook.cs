// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetGuardHook
    {
        /// <summary>
        /// This is a CloudFormation resource for activating the first-party AWS::Hooks::GuardHook.
        /// </summary>
        public static Task<GetGuardHookResult> InvokeAsync(GetGuardHookArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGuardHookResult>("aws-native:cloudformation:getGuardHook", args ?? new GetGuardHookArgs(), options.WithDefaults());

        /// <summary>
        /// This is a CloudFormation resource for activating the first-party AWS::Hooks::GuardHook.
        /// </summary>
        public static Output<GetGuardHookResult> Invoke(GetGuardHookInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuardHookResult>("aws-native:cloudformation:getGuardHook", args ?? new GetGuardHookInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This is a CloudFormation resource for activating the first-party AWS::Hooks::GuardHook.
        /// </summary>
        public static Output<GetGuardHookResult> Invoke(GetGuardHookInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuardHookResult>("aws-native:cloudformation:getGuardHook", args ?? new GetGuardHookInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGuardHookArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the activated hook
        /// </summary>
        [Input("hookArn", required: true)]
        public string HookArn { get; set; } = null!;

        public GetGuardHookArgs()
        {
        }
        public static new GetGuardHookArgs Empty => new GetGuardHookArgs();
    }

    public sealed class GetGuardHookInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the activated hook
        /// </summary>
        [Input("hookArn", required: true)]
        public Input<string> HookArn { get; set; } = null!;

        public GetGuardHookInvokeArgs()
        {
        }
        public static new GetGuardHookInvokeArgs Empty => new GetGuardHookInvokeArgs();
    }


    [OutputType]
    public sealed class GetGuardHookResult
    {
        /// <summary>
        /// Attribute to specify CloudFormation behavior on hook failure.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFormation.GuardHookFailureMode? FailureMode;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the activated hook
        /// </summary>
        public readonly string? HookArn;
        /// <summary>
        /// Attribute to specify which stacks this hook applies to or should get invoked for
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFormation.GuardHookHookStatus? HookStatus;
        /// <summary>
        /// S3 Bucket where the guard validate report will be uploaded to
        /// </summary>
        public readonly string? LogBucket;
        /// <summary>
        /// Specifies the S3 location of your input parameters.
        /// </summary>
        public readonly Outputs.OptionsProperties? Options;
        /// <summary>
        /// Specifies the S3 location of your Guard rules.
        /// </summary>
        public readonly Outputs.GuardHookS3Location? RuleLocation;
        /// <summary>
        /// Filters to allow hooks to target specific stack attributes
        /// </summary>
        public readonly Outputs.StackFiltersProperties? StackFilters;
        /// <summary>
        /// Attribute to specify which targets should invoke the hook
        /// </summary>
        public readonly Union<Outputs.TargetFilters0Properties, Outputs.TargetFilters1Properties>? TargetFilters;
        /// <summary>
        /// Which operations should this Hook run against? Resource changes, stacks or change sets.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.CloudFormation.GuardHookTargetOperation> TargetOperations;

        [OutputConstructor]
        private GetGuardHookResult(
            Pulumi.AwsNative.CloudFormation.GuardHookFailureMode? failureMode,

            string? hookArn,

            Pulumi.AwsNative.CloudFormation.GuardHookHookStatus? hookStatus,

            string? logBucket,

            Outputs.OptionsProperties? options,

            Outputs.GuardHookS3Location? ruleLocation,

            Outputs.StackFiltersProperties? stackFilters,

            Union<Outputs.TargetFilters0Properties, Outputs.TargetFilters1Properties>? targetFilters,

            ImmutableArray<Pulumi.AwsNative.CloudFormation.GuardHookTargetOperation> targetOperations)
        {
            FailureMode = failureMode;
            HookArn = hookArn;
            HookStatus = hookStatus;
            LogBucket = logBucket;
            Options = options;
            RuleLocation = ruleLocation;
            StackFilters = stackFilters;
            TargetFilters = targetFilters;
            TargetOperations = targetOperations;
        }
    }
}
