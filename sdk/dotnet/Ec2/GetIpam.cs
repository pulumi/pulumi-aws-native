// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetIpam
    {
        /// <summary>
        /// Resource Schema of AWS::EC2::IPAM Type
        /// </summary>
        public static Task<GetIpamResult> InvokeAsync(GetIpamArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpamResult>("aws-native:ec2:getIpam", args ?? new GetIpamArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Schema of AWS::EC2::IPAM Type
        /// </summary>
        public static Output<GetIpamResult> Invoke(GetIpamInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpamResult>("aws-native:ec2:getIpam", args ?? new GetIpamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the IPAM.
        /// </summary>
        [Input("ipamId", required: true)]
        public string IpamId { get; set; } = null!;

        public GetIpamArgs()
        {
        }
        public static new GetIpamArgs Empty => new GetIpamArgs();
    }

    public sealed class GetIpamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the IPAM.
        /// </summary>
        [Input("ipamId", required: true)]
        public Input<string> IpamId { get; set; } = null!;

        public GetIpamInvokeArgs()
        {
        }
        public static new GetIpamInvokeArgs Empty => new GetIpamInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpamResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IPAM.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The Id of the default association to the default resource discovery, created with this IPAM.
        /// </summary>
        public readonly string? DefaultResourceDiscoveryAssociationId;
        /// <summary>
        /// The Id of the default resource discovery, created with this IPAM.
        /// </summary>
        public readonly string? DefaultResourceDiscoveryId;
        public readonly string? Description;
        /// <summary>
        /// Id of the IPAM.
        /// </summary>
        public readonly string? IpamId;
        /// <summary>
        /// The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
        /// </summary>
        public readonly ImmutableArray<Outputs.IpamOperatingRegion> OperatingRegions;
        /// <summary>
        /// The Id of the default scope for publicly routable IP space, created with this IPAM.
        /// </summary>
        public readonly string? PrivateDefaultScopeId;
        /// <summary>
        /// The Id of the default scope for publicly routable IP space, created with this IPAM.
        /// </summary>
        public readonly string? PublicDefaultScopeId;
        /// <summary>
        /// The count of resource discoveries associated with this IPAM.
        /// </summary>
        public readonly int? ResourceDiscoveryAssociationCount;
        /// <summary>
        /// The number of scopes that currently exist in this IPAM.
        /// </summary>
        public readonly int? ScopeCount;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpamTag> Tags;
        /// <summary>
        /// The tier of the IPAM.
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.IpamTier? Tier;

        [OutputConstructor]
        private GetIpamResult(
            string? arn,

            string? defaultResourceDiscoveryAssociationId,

            string? defaultResourceDiscoveryId,

            string? description,

            string? ipamId,

            ImmutableArray<Outputs.IpamOperatingRegion> operatingRegions,

            string? privateDefaultScopeId,

            string? publicDefaultScopeId,

            int? resourceDiscoveryAssociationCount,

            int? scopeCount,

            ImmutableArray<Outputs.IpamTag> tags,

            Pulumi.AwsNative.Ec2.IpamTier? tier)
        {
            Arn = arn;
            DefaultResourceDiscoveryAssociationId = defaultResourceDiscoveryAssociationId;
            DefaultResourceDiscoveryId = defaultResourceDiscoveryId;
            Description = description;
            IpamId = ipamId;
            OperatingRegions = operatingRegions;
            PrivateDefaultScopeId = privateDefaultScopeId;
            PublicDefaultScopeId = publicDefaultScopeId;
            ResourceDiscoveryAssociationCount = resourceDiscoveryAssociationCount;
            ScopeCount = scopeCount;
            Tags = tags;
            Tier = tier;
        }
    }
}