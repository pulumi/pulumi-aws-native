// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GlobalAccelerator
{
    public static class GetCrossAccountAttachment
    {
        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment
        /// </summary>
        public static Task<GetCrossAccountAttachmentResult> InvokeAsync(GetCrossAccountAttachmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCrossAccountAttachmentResult>("aws-native:globalaccelerator:getCrossAccountAttachment", args ?? new GetCrossAccountAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment
        /// </summary>
        public static Output<GetCrossAccountAttachmentResult> Invoke(GetCrossAccountAttachmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCrossAccountAttachmentResult>("aws-native:globalaccelerator:getCrossAccountAttachment", args ?? new GetCrossAccountAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCrossAccountAttachmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the attachment.
        /// </summary>
        [Input("attachmentArn", required: true)]
        public string AttachmentArn { get; set; } = null!;

        public GetCrossAccountAttachmentArgs()
        {
        }
        public static new GetCrossAccountAttachmentArgs Empty => new GetCrossAccountAttachmentArgs();
    }

    public sealed class GetCrossAccountAttachmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the attachment.
        /// </summary>
        [Input("attachmentArn", required: true)]
        public Input<string> AttachmentArn { get; set; } = null!;

        public GetCrossAccountAttachmentInvokeArgs()
        {
        }
        public static new GetCrossAccountAttachmentInvokeArgs Empty => new GetCrossAccountAttachmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetCrossAccountAttachmentResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the attachment.
        /// </summary>
        public readonly string? AttachmentArn;
        /// <summary>
        /// The Friendly identifier of the attachment.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Principals to share the resources with.
        /// </summary>
        public readonly ImmutableArray<string> Principals;
        /// <summary>
        /// Resources shared using the attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrossAccountAttachmentResource> Resources;
        /// <summary>
        /// Add tags for a cross-account attachment.
        /// 
        /// For more information, see [Tagging in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetCrossAccountAttachmentResult(
            string? attachmentArn,

            string? name,

            ImmutableArray<string> principals,

            ImmutableArray<Outputs.CrossAccountAttachmentResource> resources,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AttachmentArn = attachmentArn;
            Name = name;
            Principals = principals;
            Resources = resources;
            Tags = tags;
        }
    }
}