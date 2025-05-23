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
    public sealed class DashboardComboChartSortConfiguration
    {
        /// <summary>
        /// The item limit configuration for the category field well of a combo chart.
        /// </summary>
        public readonly Outputs.DashboardItemsLimitConfiguration? CategoryItemsLimit;
        /// <summary>
        /// The sort configuration of the category field well in a combo chart.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> CategorySort;
        /// <summary>
        /// The item limit configuration of the color field well in a combo chart.
        /// </summary>
        public readonly Outputs.DashboardItemsLimitConfiguration? ColorItemsLimit;
        /// <summary>
        /// The sort configuration of the color field well in a combo chart.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> ColorSort;

        [OutputConstructor]
        private DashboardComboChartSortConfiguration(
            Outputs.DashboardItemsLimitConfiguration? categoryItemsLimit,

            ImmutableArray<Outputs.DashboardFieldSortOptions> categorySort,

            Outputs.DashboardItemsLimitConfiguration? colorItemsLimit,

            ImmutableArray<Outputs.DashboardFieldSortOptions> colorSort)
        {
            CategoryItemsLimit = categoryItemsLimit;
            CategorySort = categorySort;
            ColorItemsLimit = colorItemsLimit;
            ColorSort = colorSort;
        }
    }
}
