// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage
{
    public static class GetOriginEndpoint
    {
        /// <summary>
        /// Resource schema for AWS::MediaPackage::OriginEndpoint
        /// </summary>
        public static Task<GetOriginEndpointResult> InvokeAsync(GetOriginEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOriginEndpointResult>("aws-native:mediapackage:getOriginEndpoint", args ?? new GetOriginEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaPackage::OriginEndpoint
        /// </summary>
        public static Output<GetOriginEndpointResult> Invoke(GetOriginEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOriginEndpointResult>("aws-native:mediapackage:getOriginEndpoint", args ?? new GetOriginEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOriginEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the OriginEndpoint.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetOriginEndpointArgs()
        {
        }
        public static new GetOriginEndpointArgs Empty => new GetOriginEndpointArgs();
    }

    public sealed class GetOriginEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the OriginEndpoint.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetOriginEndpointInvokeArgs()
        {
        }
        public static new GetOriginEndpointInvokeArgs Empty => new GetOriginEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetOriginEndpointResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
        /// </summary>
        public readonly string? Arn;
        public readonly Outputs.OriginEndpointAuthorization? Authorization;
        /// <summary>
        /// The ID of the Channel the OriginEndpoint is associated with.
        /// </summary>
        public readonly string? ChannelId;
        public readonly Outputs.OriginEndpointCmafPackage? CmafPackage;
        public readonly Outputs.OriginEndpointDashPackage? DashPackage;
        /// <summary>
        /// A short text description of the OriginEndpoint.
        /// </summary>
        public readonly string? Description;
        public readonly Outputs.OriginEndpointHlsPackage? HlsPackage;
        /// <summary>
        /// A short string appended to the end of the OriginEndpoint URL.
        /// </summary>
        public readonly string? ManifestName;
        public readonly Outputs.OriginEndpointMssPackage? MssPackage;
        /// <summary>
        /// Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination
        /// </summary>
        public readonly Pulumi.AwsNative.MediaPackage.OriginEndpointOrigination? Origination;
        /// <summary>
        /// Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint.
        /// </summary>
        public readonly int? StartoverWindowSeconds;
        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        public readonly ImmutableArray<Outputs.OriginEndpointTag> Tags;
        /// <summary>
        /// Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint.
        /// </summary>
        public readonly int? TimeDelaySeconds;
        /// <summary>
        /// The URL of the packaged OriginEndpoint for consumption.
        /// </summary>
        public readonly string? Url;
        /// <summary>
        /// A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.
        /// </summary>
        public readonly ImmutableArray<string> Whitelist;

        [OutputConstructor]
        private GetOriginEndpointResult(
            string? arn,

            Outputs.OriginEndpointAuthorization? authorization,

            string? channelId,

            Outputs.OriginEndpointCmafPackage? cmafPackage,

            Outputs.OriginEndpointDashPackage? dashPackage,

            string? description,

            Outputs.OriginEndpointHlsPackage? hlsPackage,

            string? manifestName,

            Outputs.OriginEndpointMssPackage? mssPackage,

            Pulumi.AwsNative.MediaPackage.OriginEndpointOrigination? origination,

            int? startoverWindowSeconds,

            ImmutableArray<Outputs.OriginEndpointTag> tags,

            int? timeDelaySeconds,

            string? url,

            ImmutableArray<string> whitelist)
        {
            Arn = arn;
            Authorization = authorization;
            ChannelId = channelId;
            CmafPackage = cmafPackage;
            DashPackage = dashPackage;
            Description = description;
            HlsPackage = hlsPackage;
            ManifestName = manifestName;
            MssPackage = mssPackage;
            Origination = origination;
            StartoverWindowSeconds = startoverWindowSeconds;
            Tags = tags;
            TimeDelaySeconds = timeDelaySeconds;
            Url = url;
            Whitelist = whitelist;
        }
    }
}