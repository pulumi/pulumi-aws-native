// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardKpiSortConfiguration
    {
        /// <summary>
        /// The sort configuration of the trend group fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> TrendGroupSort;

        [OutputConstructor]
        private DashboardKpiSortConfiguration(ImmutableArray<Outputs.DashboardFieldSortOptions> trendGroupSort)
        {
            TrendGroupSort = trendGroupSort;
        }
    }
}
