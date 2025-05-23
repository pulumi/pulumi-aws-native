// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetLocalGatewayRouteTable
    {
        /// <summary>
        /// Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
        /// </summary>
        public static Task<GetLocalGatewayRouteTableResult> InvokeAsync(GetLocalGatewayRouteTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalGatewayRouteTableResult>("aws-native:ec2:getLocalGatewayRouteTable", args ?? new GetLocalGatewayRouteTableArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
        /// </summary>
        public static Output<GetLocalGatewayRouteTableResult> Invoke(GetLocalGatewayRouteTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalGatewayRouteTableResult>("aws-native:ec2:getLocalGatewayRouteTable", args ?? new GetLocalGatewayRouteTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.
        /// </summary>
        public static Output<GetLocalGatewayRouteTableResult> Invoke(GetLocalGatewayRouteTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalGatewayRouteTableResult>("aws-native:ec2:getLocalGatewayRouteTable", args ?? new GetLocalGatewayRouteTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalGatewayRouteTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the local gateway route table.
        /// </summary>
        [Input("localGatewayRouteTableId", required: true)]
        public string LocalGatewayRouteTableId { get; set; } = null!;

        public GetLocalGatewayRouteTableArgs()
        {
        }
        public static new GetLocalGatewayRouteTableArgs Empty => new GetLocalGatewayRouteTableArgs();
    }

    public sealed class GetLocalGatewayRouteTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the local gateway route table.
        /// </summary>
        [Input("localGatewayRouteTableId", required: true)]
        public Input<string> LocalGatewayRouteTableId { get; set; } = null!;

        public GetLocalGatewayRouteTableInvokeArgs()
        {
        }
        public static new GetLocalGatewayRouteTableInvokeArgs Empty => new GetLocalGatewayRouteTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalGatewayRouteTableResult
    {
        /// <summary>
        /// The ARN of the local gateway route table.
        /// </summary>
        public readonly string? LocalGatewayRouteTableArn;
        /// <summary>
        /// The ID of the local gateway route table.
        /// </summary>
        public readonly string? LocalGatewayRouteTableId;
        /// <summary>
        /// The ARN of the outpost.
        /// </summary>
        public readonly string? OutpostArn;
        /// <summary>
        /// The owner of the local gateway route table.
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// The state of the local gateway route table.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The tags for the local gateway route table.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetLocalGatewayRouteTableResult(
            string? localGatewayRouteTableArn,

            string? localGatewayRouteTableId,

            string? outpostArn,

            string? ownerId,

            string? state,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            LocalGatewayRouteTableArn = localGatewayRouteTableArn;
            LocalGatewayRouteTableId = localGatewayRouteTableId;
            OutpostArn = outpostArn;
            OwnerId = ownerId;
            State = state;
            Tags = tags;
        }
    }
}
