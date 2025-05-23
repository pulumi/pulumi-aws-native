// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetVerifiedAccessInstance
    {
        /// <summary>
        /// The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.
        /// </summary>
        public static Task<GetVerifiedAccessInstanceResult> InvokeAsync(GetVerifiedAccessInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVerifiedAccessInstanceResult>("aws-native:ec2:getVerifiedAccessInstance", args ?? new GetVerifiedAccessInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.
        /// </summary>
        public static Output<GetVerifiedAccessInstanceResult> Invoke(GetVerifiedAccessInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifiedAccessInstanceResult>("aws-native:ec2:getVerifiedAccessInstance", args ?? new GetVerifiedAccessInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.
        /// </summary>
        public static Output<GetVerifiedAccessInstanceResult> Invoke(GetVerifiedAccessInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVerifiedAccessInstanceResult>("aws-native:ec2:getVerifiedAccessInstance", args ?? new GetVerifiedAccessInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVerifiedAccessInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the AWS Verified Access instance.
        /// </summary>
        [Input("verifiedAccessInstanceId", required: true)]
        public string VerifiedAccessInstanceId { get; set; } = null!;

        public GetVerifiedAccessInstanceArgs()
        {
        }
        public static new GetVerifiedAccessInstanceArgs Empty => new GetVerifiedAccessInstanceArgs();
    }

    public sealed class GetVerifiedAccessInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the AWS Verified Access instance.
        /// </summary>
        [Input("verifiedAccessInstanceId", required: true)]
        public Input<string> VerifiedAccessInstanceId { get; set; } = null!;

        public GetVerifiedAccessInstanceInvokeArgs()
        {
        }
        public static new GetVerifiedAccessInstanceInvokeArgs Empty => new GetVerifiedAccessInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetVerifiedAccessInstanceResult
    {
        /// <summary>
        /// Introduce CidrEndpointsCustomSubDomain property to represent the domain (say, ava.my-company.com)
        /// </summary>
        public readonly string? CidrEndpointsCustomSubDomain;
        /// <summary>
        /// Property to represent the name servers assoicated with the domain that AVA manages (say, ['ns1.amazonaws.com', 'ns2.amazonaws.com', 'ns3.amazonaws.com', 'ns4.amazonaws.com']).
        /// </summary>
        public readonly ImmutableArray<string> CidrEndpointsCustomSubDomainNameServers;
        /// <summary>
        /// Time this Verified Access Instance was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// A description for the AWS Verified Access instance.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Indicates whether FIPS is enabled
        /// </summary>
        public readonly bool? FipsEnabled;
        /// <summary>
        /// Time this Verified Access Instance was last updated.
        /// </summary>
        public readonly string? LastUpdatedTime;
        /// <summary>
        /// The configuration options for AWS Verified Access instances.
        /// </summary>
        public readonly Outputs.VerifiedAccessInstanceVerifiedAccessLogs? LoggingConfigurations;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The ID of the AWS Verified Access instance.
        /// </summary>
        public readonly string? VerifiedAccessInstanceId;
        /// <summary>
        /// The IDs of the AWS Verified Access trust providers.
        /// </summary>
        public readonly ImmutableArray<string> VerifiedAccessTrustProviderIds;
        /// <summary>
        /// AWS Verified Access trust providers.
        /// </summary>
        public readonly ImmutableArray<Outputs.VerifiedAccessInstanceVerifiedAccessTrustProvider> VerifiedAccessTrustProviders;

        [OutputConstructor]
        private GetVerifiedAccessInstanceResult(
            string? cidrEndpointsCustomSubDomain,

            ImmutableArray<string> cidrEndpointsCustomSubDomainNameServers,

            string? creationTime,

            string? description,

            bool? fipsEnabled,

            string? lastUpdatedTime,

            Outputs.VerifiedAccessInstanceVerifiedAccessLogs? loggingConfigurations,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? verifiedAccessInstanceId,

            ImmutableArray<string> verifiedAccessTrustProviderIds,

            ImmutableArray<Outputs.VerifiedAccessInstanceVerifiedAccessTrustProvider> verifiedAccessTrustProviders)
        {
            CidrEndpointsCustomSubDomain = cidrEndpointsCustomSubDomain;
            CidrEndpointsCustomSubDomainNameServers = cidrEndpointsCustomSubDomainNameServers;
            CreationTime = creationTime;
            Description = description;
            FipsEnabled = fipsEnabled;
            LastUpdatedTime = lastUpdatedTime;
            LoggingConfigurations = loggingConfigurations;
            Tags = tags;
            VerifiedAccessInstanceId = verifiedAccessInstanceId;
            VerifiedAccessTrustProviderIds = verifiedAccessTrustProviderIds;
            VerifiedAccessTrustProviders = verifiedAccessTrustProviders;
        }
    }
}
