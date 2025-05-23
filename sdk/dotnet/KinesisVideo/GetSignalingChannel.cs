// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisVideo
{
    public static class GetSignalingChannel
    {
        /// <summary>
        /// Resource Type Definition for AWS::KinesisVideo::SignalingChannel
        /// </summary>
        public static Task<GetSignalingChannelResult> InvokeAsync(GetSignalingChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSignalingChannelResult>("aws-native:kinesisvideo:getSignalingChannel", args ?? new GetSignalingChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type Definition for AWS::KinesisVideo::SignalingChannel
        /// </summary>
        public static Output<GetSignalingChannelResult> Invoke(GetSignalingChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSignalingChannelResult>("aws-native:kinesisvideo:getSignalingChannel", args ?? new GetSignalingChannelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type Definition for AWS::KinesisVideo::SignalingChannel
        /// </summary>
        public static Output<GetSignalingChannelResult> Invoke(GetSignalingChannelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSignalingChannelResult>("aws-native:kinesisvideo:getSignalingChannel", args ?? new GetSignalingChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSignalingChannelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kinesis Video Signaling Channel.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetSignalingChannelArgs()
        {
        }
        public static new GetSignalingChannelArgs Empty => new GetSignalingChannelArgs();
    }

    public sealed class GetSignalingChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kinesis Video Signaling Channel.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetSignalingChannelInvokeArgs()
        {
        }
        public static new GetSignalingChannelInvokeArgs Empty => new GetSignalingChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetSignalingChannelResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The period of time a signaling channel retains undelivered messages before they are discarded.
        /// </summary>
        public readonly int? MessageTtlSeconds;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisVideo.SignalingChannelType? Type;

        [OutputConstructor]
        private GetSignalingChannelResult(
            string? arn,

            int? messageTtlSeconds,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Pulumi.AwsNative.KinesisVideo.SignalingChannelType? type)
        {
            Arn = arn;
            MessageTtlSeconds = messageTtlSeconds;
            Tags = tags;
            Type = type;
        }
    }
}
