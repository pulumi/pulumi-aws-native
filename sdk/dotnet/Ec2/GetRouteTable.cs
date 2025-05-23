// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetRouteTable
    {
        /// <summary>
        /// Specifies a route table for the specified VPC. After you create a route table, you can add routes and associate the table with a subnet.
        ///  For more information, see [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*.
        /// </summary>
        public static Task<GetRouteTableResult> InvokeAsync(GetRouteTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteTableResult>("aws-native:ec2:getRouteTable", args ?? new GetRouteTableArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a route table for the specified VPC. After you create a route table, you can add routes and associate the table with a subnet.
        ///  For more information, see [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*.
        /// </summary>
        public static Output<GetRouteTableResult> Invoke(GetRouteTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteTableResult>("aws-native:ec2:getRouteTable", args ?? new GetRouteTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a route table for the specified VPC. After you create a route table, you can add routes and associate the table with a subnet.
        ///  For more information, see [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the *Amazon VPC User Guide*.
        /// </summary>
        public static Output<GetRouteTableResult> Invoke(GetRouteTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteTableResult>("aws-native:ec2:getRouteTable", args ?? new GetRouteTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the route table.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        public GetRouteTableArgs()
        {
        }
        public static new GetRouteTableArgs Empty => new GetRouteTableArgs();
    }

    public sealed class GetRouteTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the route table.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public GetRouteTableInvokeArgs()
        {
        }
        public static new GetRouteTableInvokeArgs Empty => new GetRouteTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteTableResult
    {
        /// <summary>
        /// The ID of the route table.
        /// </summary>
        public readonly string? RouteTableId;
        /// <summary>
        /// Any tags assigned to the route table.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetRouteTableResult(
            string? routeTableId,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            RouteTableId = routeTableId;
            Tags = tags;
        }
    }
}
