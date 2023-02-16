// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetVoiceChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::VoiceChannel
        /// </summary>
        public static Task<GetVoiceChannelResult> InvokeAsync(GetVoiceChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVoiceChannelResult>("aws-native:pinpoint:getVoiceChannel", args ?? new GetVoiceChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::VoiceChannel
        /// </summary>
        public static Output<GetVoiceChannelResult> Invoke(GetVoiceChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVoiceChannelResult>("aws-native:pinpoint:getVoiceChannel", args ?? new GetVoiceChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVoiceChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVoiceChannelArgs()
        {
        }
        public static new GetVoiceChannelArgs Empty => new GetVoiceChannelArgs();
    }

    public sealed class GetVoiceChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVoiceChannelInvokeArgs()
        {
        }
        public static new GetVoiceChannelInvokeArgs Empty => new GetVoiceChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetVoiceChannelResult
    {
        public readonly bool? Enabled;
        public readonly string? Id;

        [OutputConstructor]
        private GetVoiceChannelResult(
            bool? enabled,

            string? id)
        {
            Enabled = enabled;
            Id = id;
        }
    }
}