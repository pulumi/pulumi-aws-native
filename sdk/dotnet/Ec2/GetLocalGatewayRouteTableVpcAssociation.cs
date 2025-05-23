// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetLocalGatewayRouteTableVpcAssociation
    {
        /// <summary>
        /// Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
        /// </summary>
        public static Task<GetLocalGatewayRouteTableVpcAssociationResult> InvokeAsync(GetLocalGatewayRouteTableVpcAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalGatewayRouteTableVpcAssociationResult>("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", args ?? new GetLocalGatewayRouteTableVpcAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
        /// </summary>
        public static Output<GetLocalGatewayRouteTableVpcAssociationResult> Invoke(GetLocalGatewayRouteTableVpcAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalGatewayRouteTableVpcAssociationResult>("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", args ?? new GetLocalGatewayRouteTableVpcAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
        /// </summary>
        public static Output<GetLocalGatewayRouteTableVpcAssociationResult> Invoke(GetLocalGatewayRouteTableVpcAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalGatewayRouteTableVpcAssociationResult>("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", args ?? new GetLocalGatewayRouteTableVpcAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalGatewayRouteTableVpcAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the association.
        /// </summary>
        [Input("localGatewayRouteTableVpcAssociationId", required: true)]
        public string LocalGatewayRouteTableVpcAssociationId { get; set; } = null!;

        public GetLocalGatewayRouteTableVpcAssociationArgs()
        {
        }
        public static new GetLocalGatewayRouteTableVpcAssociationArgs Empty => new GetLocalGatewayRouteTableVpcAssociationArgs();
    }

    public sealed class GetLocalGatewayRouteTableVpcAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the association.
        /// </summary>
        [Input("localGatewayRouteTableVpcAssociationId", required: true)]
        public Input<string> LocalGatewayRouteTableVpcAssociationId { get; set; } = null!;

        public GetLocalGatewayRouteTableVpcAssociationInvokeArgs()
        {
        }
        public static new GetLocalGatewayRouteTableVpcAssociationInvokeArgs Empty => new GetLocalGatewayRouteTableVpcAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalGatewayRouteTableVpcAssociationResult
    {
        /// <summary>
        /// The ID of the local gateway.
        /// </summary>
        public readonly string? LocalGatewayId;
        /// <summary>
        /// The ID of the association.
        /// </summary>
        public readonly string? LocalGatewayRouteTableVpcAssociationId;
        /// <summary>
        /// The state of the association.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The tags for the association.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetLocalGatewayRouteTableVpcAssociationResult(
            string? localGatewayId,

            string? localGatewayRouteTableVpcAssociationId,

            string? state,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            LocalGatewayId = localGatewayId;
            LocalGatewayRouteTableVpcAssociationId = localGatewayRouteTableVpcAssociationId;
            State = state;
            Tags = tags;
        }
    }
}
