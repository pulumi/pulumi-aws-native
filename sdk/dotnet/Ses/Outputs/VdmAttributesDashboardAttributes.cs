// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    /// <summary>
    /// Preferences regarding the Dashboard feature.
    /// </summary>
    [OutputType]
    public sealed class VdmAttributesDashboardAttributes
    {
        /// <summary>
        /// Whether emails sent from this account have engagement tracking enabled.
        /// </summary>
        public readonly string? EngagementMetrics;

        [OutputConstructor]
        private VdmAttributesDashboardAttributes(string? engagementMetrics)
        {
            EngagementMetrics = engagementMetrics;
        }
    }
}
