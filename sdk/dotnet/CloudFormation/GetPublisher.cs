// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetPublisher
    {
        /// <summary>
        /// Register as a publisher in the CloudFormation Registry.
        /// </summary>
        public static Task<GetPublisherResult> InvokeAsync(GetPublisherArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublisherResult>("aws-native:cloudformation:getPublisher", args ?? new GetPublisherArgs(), options.WithDefaults());

        /// <summary>
        /// Register as a publisher in the CloudFormation Registry.
        /// </summary>
        public static Output<GetPublisherResult> Invoke(GetPublisherInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublisherResult>("aws-native:cloudformation:getPublisher", args ?? new GetPublisherInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Register as a publisher in the CloudFormation Registry.
        /// </summary>
        public static Output<GetPublisherResult> Invoke(GetPublisherInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublisherResult>("aws-native:cloudformation:getPublisher", args ?? new GetPublisherInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublisherArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
        /// </summary>
        [Input("publisherId", required: true)]
        public string PublisherId { get; set; } = null!;

        public GetPublisherArgs()
        {
        }
        public static new GetPublisherArgs Empty => new GetPublisherArgs();
    }

    public sealed class GetPublisherInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
        /// </summary>
        [Input("publisherId", required: true)]
        public Input<string> PublisherId { get; set; } = null!;

        public GetPublisherInvokeArgs()
        {
        }
        public static new GetPublisherInvokeArgs Empty => new GetPublisherInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublisherResult
    {
        /// <summary>
        /// The type of account used as the identity provider when registering this publisher with CloudFormation.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFormation.PublisherIdentityProvider? IdentityProvider;
        /// <summary>
        /// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
        /// </summary>
        public readonly string? PublisherId;
        /// <summary>
        /// The URL to the publisher's profile with the identity provider.
        /// </summary>
        public readonly string? PublisherProfile;
        /// <summary>
        /// Whether the publisher is verified.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFormation.PublisherStatus? PublisherStatus;

        [OutputConstructor]
        private GetPublisherResult(
            Pulumi.AwsNative.CloudFormation.PublisherIdentityProvider? identityProvider,

            string? publisherId,

            string? publisherProfile,

            Pulumi.AwsNative.CloudFormation.PublisherStatus? publisherStatus)
        {
            IdentityProvider = identityProvider;
            PublisherId = publisherId;
            PublisherProfile = publisherProfile;
            PublisherStatus = publisherStatus;
        }
    }
}
