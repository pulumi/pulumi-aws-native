// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    public static class GetOidcProvider
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::OIDCProvider
        /// </summary>
        public static Task<GetOidcProviderResult> InvokeAsync(GetOidcProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOidcProviderResult>("aws-native:iam:getOidcProvider", args ?? new GetOidcProviderArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::OIDCProvider
        /// </summary>
        public static Output<GetOidcProviderResult> Invoke(GetOidcProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcProviderResult>("aws-native:iam:getOidcProvider", args ?? new GetOidcProviderInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::OIDCProvider
        /// </summary>
        public static Output<GetOidcProviderResult> Invoke(GetOidcProviderInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOidcProviderResult>("aws-native:iam:getOidcProvider", args ?? new GetOidcProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOidcProviderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetOidcProviderArgs()
        {
        }
        public static new GetOidcProviderArgs Empty => new GetOidcProviderArgs();
    }

    public sealed class GetOidcProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetOidcProviderInvokeArgs()
        {
        }
        public static new GetOidcProviderInvokeArgs Empty => new GetOidcProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetOidcProviderResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the OIDC provider
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
        /// </summary>
        public readonly ImmutableArray<string> ClientIdList;
        /// <summary>
        /// A list of tags that are attached to the specified IAM OIDC provider. The returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide* .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) .
        /// 
        /// This property is optional. If it is not included, IAM will retrieve and use the top intermediate certificate authority (CA) thumbprint of the OpenID Connect identity provider server certificate.
        /// </summary>
        public readonly ImmutableArray<string> ThumbprintList;

        [OutputConstructor]
        private GetOidcProviderResult(
            string? arn,

            ImmutableArray<string> clientIdList,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            ImmutableArray<string> thumbprintList)
        {
            Arn = arn;
            ClientIdList = clientIdList;
            Tags = tags;
            ThumbprintList = thumbprintList;
        }
    }
}
