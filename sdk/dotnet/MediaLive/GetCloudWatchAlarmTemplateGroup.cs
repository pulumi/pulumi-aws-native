// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive
{
    public static class GetCloudWatchAlarmTemplateGroup
    {
        /// <summary>
        /// Definition of AWS::MediaLive::CloudWatchAlarmTemplateGroup Resource Type
        /// </summary>
        public static Task<GetCloudWatchAlarmTemplateGroupResult> InvokeAsync(GetCloudWatchAlarmTemplateGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudWatchAlarmTemplateGroupResult>("aws-native:medialive:getCloudWatchAlarmTemplateGroup", args ?? new GetCloudWatchAlarmTemplateGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::MediaLive::CloudWatchAlarmTemplateGroup Resource Type
        /// </summary>
        public static Output<GetCloudWatchAlarmTemplateGroupResult> Invoke(GetCloudWatchAlarmTemplateGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudWatchAlarmTemplateGroupResult>("aws-native:medialive:getCloudWatchAlarmTemplateGroup", args ?? new GetCloudWatchAlarmTemplateGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudWatchAlarmTemplateGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("identifier", required: true)]
        public string Identifier { get; set; } = null!;

        public GetCloudWatchAlarmTemplateGroupArgs()
        {
        }
        public static new GetCloudWatchAlarmTemplateGroupArgs Empty => new GetCloudWatchAlarmTemplateGroupArgs();
    }

    public sealed class GetCloudWatchAlarmTemplateGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        public GetCloudWatchAlarmTemplateGroupInvokeArgs()
        {
        }
        public static new GetCloudWatchAlarmTemplateGroupInvokeArgs Empty => new GetCloudWatchAlarmTemplateGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudWatchAlarmTemplateGroupResult
    {
        /// <summary>
        /// A cloudwatch alarm template group's ARN (Amazon Resource Name)
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The date and time of resource creation.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A resource's optional description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
        /// </summary>
        public readonly string? Id;
        public readonly string? Identifier;
        /// <summary>
        /// The date and time of latest resource modification.
        /// </summary>
        public readonly string? ModifiedAt;

        [OutputConstructor]
        private GetCloudWatchAlarmTemplateGroupResult(
            string? arn,

            string? createdAt,

            string? description,

            string? id,

            string? identifier,

            string? modifiedAt)
        {
            Arn = arn;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Identifier = identifier;
            ModifiedAt = modifiedAt;
        }
    }
}