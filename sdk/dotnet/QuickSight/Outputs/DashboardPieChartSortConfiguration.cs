// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardPieChartSortConfiguration
    {
        public readonly Outputs.DashboardItemsLimitConfiguration? CategoryItemsLimit;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> CategorySort;
        public readonly Outputs.DashboardItemsLimitConfiguration? SmallMultiplesLimitConfiguration;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> SmallMultiplesSort;

        [OutputConstructor]
        private DashboardPieChartSortConfiguration(
            Outputs.DashboardItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.DashboardFieldSortOptions> categorySort,

            Outputs.DashboardItemsLimitConfiguration? smallMultiplesLimitConfiguration,

            ImmutableArray<Outputs.DashboardFieldSortOptions> smallMultiplesSort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
            SmallMultiplesLimitConfiguration = smallMultiplesLimitConfiguration;
            SmallMultiplesSort = smallMultiplesSort;
        }
    }
}