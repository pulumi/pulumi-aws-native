// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver
{
    public static class GetResolverEndpoint
    {
        /// <summary>
        /// Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
        /// </summary>
        public static Task<GetResolverEndpointResult> InvokeAsync(GetResolverEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResolverEndpointResult>("aws-native:route53resolver:getResolverEndpoint", args ?? new GetResolverEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Route53Resolver::ResolverEndpoint
        /// </summary>
        public static Output<GetResolverEndpointResult> Invoke(GetResolverEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResolverEndpointResult>("aws-native:route53resolver:getResolverEndpoint", args ?? new GetResolverEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResolverEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("resolverEndpointId", required: true)]
        public string ResolverEndpointId { get; set; } = null!;

        public GetResolverEndpointArgs()
        {
        }
        public static new GetResolverEndpointArgs Empty => new GetResolverEndpointArgs();
    }

    public sealed class GetResolverEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("resolverEndpointId", required: true)]
        public Input<string> ResolverEndpointId { get; set; } = null!;

        public GetResolverEndpointInvokeArgs()
        {
        }
        public static new GetResolverEndpointInvokeArgs Empty => new GetResolverEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetResolverEndpointResult
    {
        public readonly string? Arn;
        public readonly string? HostVPCId;
        public readonly string? IpAddressCount;
        public readonly ImmutableArray<Outputs.ResolverEndpointIpAddressRequest> IpAddresses;
        public readonly string? Name;
        public readonly string? ResolverEndpointId;
        public readonly string? ResolverEndpointType;
        public readonly ImmutableArray<Outputs.ResolverEndpointTag> Tags;

        [OutputConstructor]
        private GetResolverEndpointResult(
            string? arn,

            string? hostVPCId,

            string? ipAddressCount,

            ImmutableArray<Outputs.ResolverEndpointIpAddressRequest> ipAddresses,

            string? name,

            string? resolverEndpointId,

            string? resolverEndpointType,

            ImmutableArray<Outputs.ResolverEndpointTag> tags)
        {
            Arn = arn;
            HostVPCId = hostVPCId;
            IpAddressCount = ipAddressCount;
            IpAddresses = ipAddresses;
            Name = name;
            ResolverEndpointId = resolverEndpointId;
            ResolverEndpointType = resolverEndpointType;
            Tags = tags;
        }
    }
}