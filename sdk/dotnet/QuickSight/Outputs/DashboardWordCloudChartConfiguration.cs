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
    public sealed class DashboardWordCloudChartConfiguration
    {
        public readonly Outputs.DashboardChartAxisLabelOptions? CategoryLabelOptions;
        public readonly Outputs.DashboardWordCloudFieldWells? FieldWells;
        public readonly Outputs.DashboardWordCloudSortConfiguration? SortConfiguration;
        public readonly Outputs.DashboardWordCloudOptions? WordCloudOptions;

        [OutputConstructor]
        private DashboardWordCloudChartConfiguration(
            Outputs.DashboardChartAxisLabelOptions? categoryLabelOptions,

            Outputs.DashboardWordCloudFieldWells? fieldWells,

            Outputs.DashboardWordCloudSortConfiguration? sortConfiguration,

            Outputs.DashboardWordCloudOptions? wordCloudOptions)
        {
            CategoryLabelOptions = categoryLabelOptions;
            FieldWells = fieldWells;
            SortConfiguration = sortConfiguration;
            WordCloudOptions = wordCloudOptions;
        }
    }
}