// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetGcmChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::GCMChannel
        /// </summary>
        public static Task<GetGcmChannelResult> InvokeAsync(GetGcmChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGcmChannelResult>("aws-native:pinpoint:getGcmChannel", args ?? new GetGcmChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::GCMChannel
        /// </summary>
        public static Output<GetGcmChannelResult> Invoke(GetGcmChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGcmChannelResult>("aws-native:pinpoint:getGcmChannel", args ?? new GetGcmChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGcmChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetGcmChannelArgs()
        {
        }
        public static new GetGcmChannelArgs Empty => new GetGcmChannelArgs();
    }

    public sealed class GetGcmChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetGcmChannelInvokeArgs()
        {
        }
        public static new GetGcmChannelInvokeArgs Empty => new GetGcmChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetGcmChannelResult
    {
        public readonly string? ApiKey;
        public readonly string? DefaultAuthenticationMethod;
        public readonly bool? Enabled;
        public readonly string? Id;
        public readonly string? ServiceJson;

        [OutputConstructor]
        private GetGcmChannelResult(
            string? apiKey,

            string? defaultAuthenticationMethod,

            bool? enabled,

            string? id,

            string? serviceJson)
        {
            ApiKey = apiKey;
            DefaultAuthenticationMethod = defaultAuthenticationMethod;
            Enabled = enabled;
            Id = id;
            ServiceJson = serviceJson;
        }
    }
}