// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetVerifiedAccessTrustProvider
    {
        /// <summary>
        /// The AWS::EC2::VerifiedAccessTrustProvider type describes a verified access trust provider
        /// </summary>
        public static Task<GetVerifiedAccessTrustProviderResult> InvokeAsync(GetVerifiedAccessTrustProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVerifiedAccessTrustProviderResult>("aws-native:ec2:getVerifiedAccessTrustProvider", args ?? new GetVerifiedAccessTrustProviderArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::VerifiedAccessTrustProvider type describes a verified access trust provider
        /// </summary>
        public static Output<GetVerifiedAccessTrustProviderResult> Invoke(GetVerifiedAccessTrustProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifiedAccessTrustProviderResult>("aws-native:ec2:getVerifiedAccessTrustProvider", args ?? new GetVerifiedAccessTrustProviderInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::VerifiedAccessTrustProvider type describes a verified access trust provider
        /// </summary>
        public static Output<GetVerifiedAccessTrustProviderResult> Invoke(GetVerifiedAccessTrustProviderInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifiedAccessTrustProviderResult>("aws-native:ec2:getVerifiedAccessTrustProvider", args ?? new GetVerifiedAccessTrustProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVerifiedAccessTrustProviderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Amazon Web Services Verified Access trust provider.
        /// </summary>
        [Input("verifiedAccessTrustProviderId", required: true)]
        public string VerifiedAccessTrustProviderId { get; set; } = null!;

        public GetVerifiedAccessTrustProviderArgs()
        {
        }
        public static new GetVerifiedAccessTrustProviderArgs Empty => new GetVerifiedAccessTrustProviderArgs();
    }

    public sealed class GetVerifiedAccessTrustProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Amazon Web Services Verified Access trust provider.
        /// </summary>
        [Input("verifiedAccessTrustProviderId", required: true)]
        public Input<string> VerifiedAccessTrustProviderId { get; set; } = null!;

        public GetVerifiedAccessTrustProviderInvokeArgs()
        {
        }
        public static new GetVerifiedAccessTrustProviderInvokeArgs Empty => new GetVerifiedAccessTrustProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetVerifiedAccessTrustProviderResult
    {
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// A description for the Amazon Web Services Verified Access trust provider.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The last updated time.
        /// </summary>
        public readonly string? LastUpdatedTime;
        /// <summary>
        /// The OpenID Connect (OIDC) options.
        /// </summary>
        public readonly Outputs.VerifiedAccessTrustProviderNativeApplicationOidcOptions? NativeApplicationOidcOptions;
        /// <summary>
        /// The options for an OpenID Connect-compatible user-identity trust provider.
        /// </summary>
        public readonly Outputs.VerifiedAccessTrustProviderOidcOptions? OidcOptions;
        /// <summary>
        /// The configuration options for customer provided KMS encryption.
        /// </summary>
        public readonly Outputs.SseSpecificationProperties? SseSpecification;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The ID of the Amazon Web Services Verified Access trust provider.
        /// </summary>
        public readonly string? VerifiedAccessTrustProviderId;

        [OutputConstructor]
        private GetVerifiedAccessTrustProviderResult(
            string? creationTime,

            string? description,

            string? lastUpdatedTime,

            Outputs.VerifiedAccessTrustProviderNativeApplicationOidcOptions? nativeApplicationOidcOptions,

            Outputs.VerifiedAccessTrustProviderOidcOptions? oidcOptions,

            Outputs.SseSpecificationProperties? sseSpecification,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? verifiedAccessTrustProviderId)
        {
            CreationTime = creationTime;
            Description = description;
            LastUpdatedTime = lastUpdatedTime;
            NativeApplicationOidcOptions = nativeApplicationOidcOptions;
            OidcOptions = oidcOptions;
            SseSpecification = sseSpecification;
            Tags = tags;
            VerifiedAccessTrustProviderId = verifiedAccessTrustProviderId;
        }
    }
}
