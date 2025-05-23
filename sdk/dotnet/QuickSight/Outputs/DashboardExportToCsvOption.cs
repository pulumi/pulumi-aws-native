// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Export to .csv option.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DashboardExportToCsvOption
    {
        /// <summary>
        /// Availability status.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardBehavior? AvailabilityStatus;

        [OutputConstructor]
        private DashboardExportToCsvOption(Pulumi.AwsNative.QuickSight.DashboardBehavior? availabilityStatus)
        {
            AvailabilityStatus = availabilityStatus;
        }
    }
}
