// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetCustomerGateway
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::CustomerGateway
        /// </summary>
        public static Task<GetCustomerGatewayResult> InvokeAsync(GetCustomerGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomerGatewayResult>("aws-native:ec2:getCustomerGateway", args ?? new GetCustomerGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::CustomerGateway
        /// </summary>
        public static Output<GetCustomerGatewayResult> Invoke(GetCustomerGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomerGatewayResult>("aws-native:ec2:getCustomerGateway", args ?? new GetCustomerGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomerGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public string CustomerGatewayId { get; set; } = null!;

        public GetCustomerGatewayArgs()
        {
        }
        public static new GetCustomerGatewayArgs Empty => new GetCustomerGatewayArgs();
    }

    public sealed class GetCustomerGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        public GetCustomerGatewayInvokeArgs()
        {
        }
        public static new GetCustomerGatewayInvokeArgs Empty => new GetCustomerGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomerGatewayResult
    {
        /// <summary>
        /// CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.
        /// </summary>
        public readonly string? CustomerGatewayId;
        /// <summary>
        /// One or more tags for the customer gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomerGatewayTag> Tags;

        [OutputConstructor]
        private GetCustomerGatewayResult(
            string? customerGatewayId,

            ImmutableArray<Outputs.CustomerGatewayTag> tags)
        {
            CustomerGatewayId = customerGatewayId;
            Tags = tags;
        }
    }
}