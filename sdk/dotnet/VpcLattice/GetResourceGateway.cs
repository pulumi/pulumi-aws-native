// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice
{
    public static class GetResourceGateway
    {
        /// <summary>
        /// Creates a resource gateway for a service. 
        /// </summary>
        public static Task<GetResourceGatewayResult> InvokeAsync(GetResourceGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceGatewayResult>("aws-native:vpclattice:getResourceGateway", args ?? new GetResourceGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a resource gateway for a service. 
        /// </summary>
        public static Output<GetResourceGatewayResult> Invoke(GetResourceGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceGatewayResult>("aws-native:vpclattice:getResourceGateway", args ?? new GetResourceGatewayInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Creates a resource gateway for a service. 
        /// </summary>
        public static Output<GetResourceGatewayResult> Invoke(GetResourceGatewayInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceGatewayResult>("aws-native:vpclattice:getResourceGateway", args ?? new GetResourceGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceGatewayArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource gateway.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetResourceGatewayArgs()
        {
        }
        public static new GetResourceGatewayArgs Empty => new GetResourceGatewayArgs();
    }

    public sealed class GetResourceGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource gateway.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetResourceGatewayInvokeArgs()
        {
        }
        public static new GetResourceGatewayInvokeArgs Empty => new GetResourceGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourceGatewayResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the resource gateway.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The ID of the resource gateway.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The ID of one or more security groups to associate with the endpoint network interface.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The tags for the resource gateway.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetResourceGatewayResult(
            string? arn,

            string? id,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Id = id;
            SecurityGroupIds = securityGroupIds;
            Tags = tags;
        }
    }
}
