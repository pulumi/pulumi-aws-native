// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetLocationSmb
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::LocationSMB.
        /// </summary>
        public static Task<GetLocationSmbResult> InvokeAsync(GetLocationSmbArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocationSmbResult>("aws-native:datasync:getLocationSmb", args ?? new GetLocationSmbArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::LocationSMB.
        /// </summary>
        public static Output<GetLocationSmbResult> Invoke(GetLocationSmbInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocationSmbResult>("aws-native:datasync:getLocationSmb", args ?? new GetLocationSmbInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationSmbArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public string LocationArn { get; set; } = null!;

        public GetLocationSmbArgs()
        {
        }
        public static new GetLocationSmbArgs Empty => new GetLocationSmbArgs();
    }

    public sealed class GetLocationSmbInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        public GetLocationSmbInvokeArgs()
        {
        }
        public static new GetLocationSmbInvokeArgs Empty => new GetLocationSmbInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocationSmbResult
    {
        /// <summary>
        /// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        /// </summary>
        public readonly ImmutableArray<string> AgentArns;
        /// <summary>
        /// The name of the Windows domain that the SMB server belongs to.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        public readonly string? LocationArn;
        /// <summary>
        /// The URL of the SMB location that was described.
        /// </summary>
        public readonly string? LocationUri;
        public readonly Outputs.LocationSmbMountOptions? MountOptions;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationSmbTag> Tags;
        /// <summary>
        /// The user who can mount the share, has the permissions to access files and folders in the SMB share.
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private GetLocationSmbResult(
            ImmutableArray<string> agentArns,

            string? domain,

            string? locationArn,

            string? locationUri,

            Outputs.LocationSmbMountOptions? mountOptions,

            ImmutableArray<Outputs.LocationSmbTag> tags,

            string? user)
        {
            AgentArns = agentArns;
            Domain = domain;
            LocationArn = locationArn;
            LocationUri = locationUri;
            MountOptions = mountOptions;
            Tags = tags;
            User = user;
        }
    }
}