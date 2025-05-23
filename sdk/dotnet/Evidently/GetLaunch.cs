// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently
{
    public static class GetLaunch
    {
        /// <summary>
        /// Resource Type definition for AWS::Evidently::Launch.
        /// </summary>
        public static Task<GetLaunchResult> InvokeAsync(GetLaunchArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLaunchResult>("aws-native:evidently:getLaunch", args ?? new GetLaunchArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Evidently::Launch.
        /// </summary>
        public static Output<GetLaunchResult> Invoke(GetLaunchInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLaunchResult>("aws-native:evidently:getLaunch", args ?? new GetLaunchInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Evidently::Launch.
        /// </summary>
        public static Output<GetLaunchResult> Invoke(GetLaunchInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLaunchResult>("aws-native:evidently:getLaunch", args ?? new GetLaunchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLaunchArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetLaunchArgs()
        {
        }
        public static new GetLaunchArgs Empty => new GetLaunchArgs();
    }

    public sealed class GetLaunchInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetLaunchInvokeArgs()
        {
        }
        public static new GetLaunchInvokeArgs Empty => new GetLaunchInvokeArgs();
    }


    [OutputType]
    public sealed class GetLaunchResult
    {
        /// <summary>
        /// The ARN of the launch. For example, `arn:aws:evidently:us-west-2:0123455678912:project/myProject/launch/myLaunch`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// An optional description for the launch.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Start or Stop Launch Launch. Default is not started.
        /// </summary>
        public readonly Outputs.LaunchExecutionStatusObject? ExecutionStatus;
        /// <summary>
        /// An array of structures that contains the feature and variations that are to be used for the launch. You can up to five launch groups in a launch.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchGroupObject> Groups;
        /// <summary>
        /// An array of structures that define the metrics that will be used to monitor the launch performance. You can have up to three metric monitors in the array.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchMetricDefinitionObject> MetricMonitors;
        /// <summary>
        /// When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served. This randomization ID is a combination of the entity ID and `randomizationSalt` . If you omit `randomizationSalt` , Evidently uses the launch name as the `randomizationsSalt` .
        /// </summary>
        public readonly string? RandomizationSalt;
        /// <summary>
        /// An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchStepConfig> ScheduledSplitsConfig;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetLaunchResult(
            string? arn,

            string? description,

            Outputs.LaunchExecutionStatusObject? executionStatus,

            ImmutableArray<Outputs.LaunchGroupObject> groups,

            ImmutableArray<Outputs.LaunchMetricDefinitionObject> metricMonitors,

            string? randomizationSalt,

            ImmutableArray<Outputs.LaunchStepConfig> scheduledSplitsConfig,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Description = description;
            ExecutionStatus = executionStatus;
            Groups = groups;
            MetricMonitors = metricMonitors;
            RandomizationSalt = randomizationSalt;
            ScheduledSplitsConfig = scheduledSplitsConfig;
            Tags = tags;
        }
    }
}
