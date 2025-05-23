// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice
{
    public static class GetResourceConfiguration
    {
        /// <summary>
        /// VpcLattice ResourceConfiguration CFN resource
        /// </summary>
        public static Task<GetResourceConfigurationResult> InvokeAsync(GetResourceConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceConfigurationResult>("aws-native:vpclattice:getResourceConfiguration", args ?? new GetResourceConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// VpcLattice ResourceConfiguration CFN resource
        /// </summary>
        public static Output<GetResourceConfigurationResult> Invoke(GetResourceConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceConfigurationResult>("aws-native:vpclattice:getResourceConfiguration", args ?? new GetResourceConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// VpcLattice ResourceConfiguration CFN resource
        /// </summary>
        public static Output<GetResourceConfigurationResult> Invoke(GetResourceConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceConfigurationResult>("aws-native:vpclattice:getResourceConfiguration", args ?? new GetResourceConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource configuration.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetResourceConfigurationArgs()
        {
        }
        public static new GetResourceConfigurationArgs Empty => new GetResourceConfigurationArgs();
    }

    public sealed class GetResourceConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource configuration.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetResourceConfigurationInvokeArgs()
        {
        }
        public static new GetResourceConfigurationInvokeArgs Empty => new GetResourceConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceConfigurationResult
    {
        /// <summary>
        /// Specifies whether the resource configuration can be associated with a sharable service network.
        /// </summary>
        public readonly bool? AllowAssociationToSharableServiceNetwork;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource configuration.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The ID of the resource configuration.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource configuration.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (SINGLE, GROUP, CHILD) The TCP port ranges that a consumer can use to access a resource configuration (for example: 1-65535). You can separate port ranges using commas (for example: 1,2,22-30).
        /// </summary>
        public readonly ImmutableArray<string> PortRanges;
        /// <summary>
        /// Identifies the resource configuration in one of the following ways:
        /// 
        /// - *Amazon Resource Name (ARN)* - Supported resource-types that are provisioned by AWS services, such as RDS databases, can be identified by their ARN.
        /// - *Domain name* - Any domain name that is publicly resolvable.
        /// - *IP address* - For IPv4 and IPv6, only IP addresses in the VPC are supported.
        /// </summary>
        public readonly object? ResourceConfigurationDefinition;
        /// <summary>
        /// The tags for the resource configuration.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetResourceConfigurationResult(
            bool? allowAssociationToSharableServiceNetwork,

            string? arn,

            string? id,

            string? name,

            ImmutableArray<string> portRanges,

            object? resourceConfigurationDefinition,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AllowAssociationToSharableServiceNetwork = allowAssociationToSharableServiceNetwork;
            Arn = arn;
            Id = id;
            Name = name;
            PortRanges = portRanges;
            ResourceConfigurationDefinition = resourceConfigurationDefinition;
            Tags = tags;
        }
    }
}
