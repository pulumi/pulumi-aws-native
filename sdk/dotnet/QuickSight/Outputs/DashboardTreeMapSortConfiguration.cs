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
    public sealed class DashboardTreeMapSortConfiguration
    {
        public readonly Outputs.DashboardItemsLimitConfiguration? TreeMapGroupItemsLimitConfiguration;
        public readonly ImmutableArray<Outputs.DashboardFieldSortOptions> TreeMapSort;

        [OutputConstructor]
        private DashboardTreeMapSortConfiguration(
            Outputs.DashboardItemsLimitConfiguration? treeMapGroupItemsLimitConfiguration,

            ImmutableArray<Outputs.DashboardFieldSortOptions> treeMapSort)
        {
            TreeMapGroupItemsLimitConfiguration = treeMapGroupItemsLimitConfiguration;
            TreeMapSort = treeMapSort;
        }
    }
}