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
    public sealed class TemplateHeatMapSortConfiguration
    {
        public readonly Outputs.TemplateItemsLimitConfiguration? HeatMapColumnItemsLimitConfiguration;
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> HeatMapColumnSort;
        public readonly Outputs.TemplateItemsLimitConfiguration? HeatMapRowItemsLimitConfiguration;
        public readonly ImmutableArray<Outputs.TemplateFieldSortOptions> HeatMapRowSort;

        [OutputConstructor]
        private TemplateHeatMapSortConfiguration(
            Outputs.TemplateItemsLimitConfiguration? heatMapColumnItemsLimitConfiguration,

            ImmutableArray<Outputs.TemplateFieldSortOptions> heatMapColumnSort,

            Outputs.TemplateItemsLimitConfiguration? heatMapRowItemsLimitConfiguration,

            ImmutableArray<Outputs.TemplateFieldSortOptions> heatMapRowSort)
        {
            HeatMapColumnItemsLimitConfiguration = heatMapColumnItemsLimitConfiguration;
            HeatMapColumnSort = heatMapColumnSort;
            HeatMapRowItemsLimitConfiguration = heatMapRowItemsLimitConfiguration;
            HeatMapRowSort = heatMapRowSort;
        }
    }
}