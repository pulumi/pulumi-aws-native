// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetTransitGatewayVpcAttachment
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment
        /// </summary>
        public static Task<GetTransitGatewayVpcAttachmentResult> InvokeAsync(GetTransitGatewayVpcAttachmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTransitGatewayVpcAttachmentResult>("aws-native:ec2:getTransitGatewayVpcAttachment", args ?? new GetTransitGatewayVpcAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment
        /// </summary>
        public static Output<GetTransitGatewayVpcAttachmentResult> Invoke(GetTransitGatewayVpcAttachmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTransitGatewayVpcAttachmentResult>("aws-native:ec2:getTransitGatewayVpcAttachment", args ?? new GetTransitGatewayVpcAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitGatewayVpcAttachmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTransitGatewayVpcAttachmentArgs()
        {
        }
        public static new GetTransitGatewayVpcAttachmentArgs Empty => new GetTransitGatewayVpcAttachmentArgs();
    }

    public sealed class GetTransitGatewayVpcAttachmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTransitGatewayVpcAttachmentInvokeArgs()
        {
        }
        public static new GetTransitGatewayVpcAttachmentInvokeArgs Empty => new GetTransitGatewayVpcAttachmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetTransitGatewayVpcAttachmentResult
    {
        public readonly ImmutableArray<string> AddSubnetIds;
        public readonly string? Id;
        /// <summary>
        /// The options for the transit gateway vpc attachment.
        /// </summary>
        public readonly Outputs.OptionsProperties? Options;
        public readonly ImmutableArray<string> RemoveSubnetIds;
        public readonly ImmutableArray<Outputs.TransitGatewayVpcAttachmentTag> Tags;

        [OutputConstructor]
        private GetTransitGatewayVpcAttachmentResult(
            ImmutableArray<string> addSubnetIds,

            string? id,

            Outputs.OptionsProperties? options,

            ImmutableArray<string> removeSubnetIds,

            ImmutableArray<Outputs.TransitGatewayVpcAttachmentTag> tags)
        {
            AddSubnetIds = addSubnetIds;
            Id = id;
            Options = options;
            RemoveSubnetIds = removeSubnetIds;
            Tags = tags;
        }
    }
}