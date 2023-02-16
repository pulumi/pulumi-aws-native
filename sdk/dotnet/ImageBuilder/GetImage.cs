// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder
{
    public static class GetImage
    {
        /// <summary>
        /// Resource schema for AWS::ImageBuilder::Image
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("aws-native:imagebuilder:getImage", args ?? new GetImageArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::ImageBuilder::Image
        /// </summary>
        public static Output<GetImageResult> Invoke(GetImageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageResult>("aws-native:imagebuilder:getImage", args ?? new GetImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetImageArgs()
        {
        }
        public static new GetImageArgs Empty => new GetImageArgs();
    }

    public sealed class GetImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetImageInvokeArgs()
        {
        }
        public static new GetImageInvokeArgs Empty => new GetImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The AMI ID of the EC2 AMI in current region.
        /// </summary>
        public readonly string? ImageId;
        /// <summary>
        /// URI for containers created in current Region with default ECR image tag
        /// </summary>
        public readonly string? ImageUri;
        /// <summary>
        /// The name of the image.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetImageResult(
            string? arn,

            string? imageId,

            string? imageUri,

            string? name)
        {
            Arn = arn;
            ImageId = imageId;
            ImageUri = imageUri;
            Name = name;
        }
    }
}