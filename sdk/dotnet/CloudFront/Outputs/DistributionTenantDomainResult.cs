// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// The details about the domain result.
    /// </summary>
    [OutputType]
    public sealed class DistributionTenantDomainResult
    {
        /// <summary>
        /// The specified domain.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// Whether the domain is active or inactive.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFront.DistributionTenantDomainResultStatus? Status;

        [OutputConstructor]
        private DistributionTenantDomainResult(
            string? domain,

            Pulumi.AwsNative.CloudFront.DistributionTenantDomainResultStatus? status)
        {
            Domain = domain;
            Status = status;
        }
    }
}
