// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    public static class GetCloudFrontOriginAccessIdentity
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity
        /// </summary>
        public static Task<GetCloudFrontOriginAccessIdentityResult> InvokeAsync(GetCloudFrontOriginAccessIdentityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudFrontOriginAccessIdentityResult>("aws-native:cloudfront:getCloudFrontOriginAccessIdentity", args ?? new GetCloudFrontOriginAccessIdentityArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity
        /// </summary>
        public static Output<GetCloudFrontOriginAccessIdentityResult> Invoke(GetCloudFrontOriginAccessIdentityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudFrontOriginAccessIdentityResult>("aws-native:cloudfront:getCloudFrontOriginAccessIdentity", args ?? new GetCloudFrontOriginAccessIdentityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudFrontOriginAccessIdentityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the origin access identity, for example, `E74FTE3AJFJ256A` .
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCloudFrontOriginAccessIdentityArgs()
        {
        }
        public static new GetCloudFrontOriginAccessIdentityArgs Empty => new GetCloudFrontOriginAccessIdentityArgs();
    }

    public sealed class GetCloudFrontOriginAccessIdentityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID for the origin access identity, for example, `E74FTE3AJFJ256A` .
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCloudFrontOriginAccessIdentityInvokeArgs()
        {
        }
        public static new GetCloudFrontOriginAccessIdentityInvokeArgs Empty => new GetCloudFrontOriginAccessIdentityInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudFrontOriginAccessIdentityResult
    {
        /// <summary>
        /// The current configuration information for the identity.
        /// </summary>
        public readonly Outputs.CloudFrontOriginAccessIdentityConfig? CloudFrontOriginAccessIdentityConfig;
        /// <summary>
        /// The ID for the origin access identity, for example, `E74FTE3AJFJ256A` .
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The Amazon S3 canonical user ID for the origin access identity, used when giving the origin access identity read permission to an object in Amazon S3. For example: `b970b42360b81c8ddbd79d2f5df0069ba9033c8a79655752abe380cd6d63ba8bcf23384d568fcf89fc49700b5e11a0fd` .
        /// </summary>
        public readonly string? S3CanonicalUserId;

        [OutputConstructor]
        private GetCloudFrontOriginAccessIdentityResult(
            Outputs.CloudFrontOriginAccessIdentityConfig? cloudFrontOriginAccessIdentityConfig,

            string? id,

            string? s3CanonicalUserId)
        {
            CloudFrontOriginAccessIdentityConfig = cloudFrontOriginAccessIdentityConfig;
            Id = id;
            S3CanonicalUserId = s3CanonicalUserId;
        }
    }
}
