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
    /// The customizations that you specified for the distribution tenant for geographic restrictions.
    /// </summary>
    [OutputType]
    public sealed class DistributionTenantGeoRestrictionCustomization
    {
        /// <summary>
        /// The locations for geographic restrictions.
        /// </summary>
        public readonly ImmutableArray<string> Locations;
        /// <summary>
        /// The method that you want to use to restrict distribution of your content by country:
        ///   +  ``none``: No geographic restriction is enabled, meaning access to content is not restricted by client geo location.
        ///   +  ``blacklist``: The ``Location`` elements specify the countries in which you don't want CloudFront to distribute your content.
        ///   +  ``whitelist``: The ``Location`` elements specify the countries in which you want CloudFront to distribute your content.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudFront.DistributionTenantGeoRestrictionCustomizationRestrictionType? RestrictionType;

        [OutputConstructor]
        private DistributionTenantGeoRestrictionCustomization(
            ImmutableArray<string> locations,

            Pulumi.AwsNative.CloudFront.DistributionTenantGeoRestrictionCustomizationRestrictionType? restrictionType)
        {
            Locations = locations;
            RestrictionType = restrictionType;
        }
    }
}
