// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.XRay
{
    public static class GetGroup
    {
        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("aws-native:xray:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws-native:xray:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("aws-native:xray:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        [Input("groupArn", required: true)]
        public string GroupArn { get; set; } = null!;

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        [Input("groupArn", required: true)]
        public Input<string> GroupArn { get; set; } = null!;

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The filter expression defining criteria by which to group traces.
        /// </summary>
        public readonly string? FilterExpression;
        /// <summary>
        /// The ARN of the group that was generated on creation.
        /// </summary>
        public readonly string? GroupArn;
        /// <summary>
        /// The case-sensitive name of the new group. Names must be unique.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// The structure containing configurations related to insights.
        /// 
        /// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
        /// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
        /// </summary>
        public readonly Outputs.GroupInsightsConfiguration? InsightsConfiguration;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetGroupResult(
            string? filterExpression,

            string? groupArn,

            string? groupName,

            Outputs.GroupInsightsConfiguration? insightsConfiguration,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            FilterExpression = filterExpression;
            GroupArn = groupArn;
            GroupName = groupName;
            InsightsConfiguration = insightsConfiguration;
            Tags = tags;
        }
    }
}
