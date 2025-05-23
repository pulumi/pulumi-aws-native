// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver
{
    public static class GetResolverDnssecConfig
    {
        /// <summary>
        /// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
        /// </summary>
        public static Task<GetResolverDnssecConfigResult> InvokeAsync(GetResolverDnssecConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResolverDnssecConfigResult>("aws-native:route53resolver:getResolverDnssecConfig", args ?? new GetResolverDnssecConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
        /// </summary>
        public static Output<GetResolverDnssecConfigResult> Invoke(GetResolverDnssecConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResolverDnssecConfigResult>("aws-native:route53resolver:getResolverDnssecConfig", args ?? new GetResolverDnssecConfigInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
        /// </summary>
        public static Output<GetResolverDnssecConfigResult> Invoke(GetResolverDnssecConfigInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetResolverDnssecConfigResult>("aws-native:route53resolver:getResolverDnssecConfig", args ?? new GetResolverDnssecConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResolverDnssecConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetResolverDnssecConfigArgs()
        {
        }
        public static new GetResolverDnssecConfigArgs Empty => new GetResolverDnssecConfigArgs();
    }

    public sealed class GetResolverDnssecConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetResolverDnssecConfigInvokeArgs()
        {
        }
        public static new GetResolverDnssecConfigInvokeArgs Empty => new GetResolverDnssecConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetResolverDnssecConfigResult
    {
        /// <summary>
        /// Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// AccountId
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.ResolverDnssecConfigValidationStatus? ValidationStatus;

        [OutputConstructor]
        private GetResolverDnssecConfigResult(
            string? id,

            string? ownerId,

            Pulumi.AwsNative.Route53Resolver.ResolverDnssecConfigValidationStatus? validationStatus)
        {
            Id = id;
            OwnerId = ownerId;
            ValidationStatus = validationStatus;
        }
    }
}
