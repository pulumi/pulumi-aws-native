// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// A Verified Access Trust Provider.
    /// </summary>
    [OutputType]
    public sealed class VerifiedAccessInstanceVerifiedAccessTrustProvider
    {
        /// <summary>
        /// The description of trust provider.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The type of device-based trust provider.
        /// </summary>
        public readonly string? DeviceTrustProviderType;
        /// <summary>
        /// The type of trust provider (user- or device-based).
        /// </summary>
        public readonly string? TrustProviderType;
        /// <summary>
        /// The type of user-based trust provider.
        /// </summary>
        public readonly string? UserTrustProviderType;
        /// <summary>
        /// The ID of the trust provider.
        /// </summary>
        public readonly string? VerifiedAccessTrustProviderId;

        [OutputConstructor]
        private VerifiedAccessInstanceVerifiedAccessTrustProvider(
            string? description,

            string? deviceTrustProviderType,

            string? trustProviderType,

            string? userTrustProviderType,

            string? verifiedAccessTrustProviderId)
        {
            Description = description;
            DeviceTrustProviderType = deviceTrustProviderType;
            TrustProviderType = trustProviderType;
            UserTrustProviderType = userTrustProviderType;
            VerifiedAccessTrustProviderId = verifiedAccessTrustProviderId;
        }
    }
}
