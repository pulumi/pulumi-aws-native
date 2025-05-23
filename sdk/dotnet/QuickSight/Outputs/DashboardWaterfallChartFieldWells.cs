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
    public sealed class DashboardWaterfallChartFieldWells
    {
        /// <summary>
        /// The field well configuration of a waterfall visual.
        /// </summary>
        public readonly Outputs.DashboardWaterfallChartAggregatedFieldWells? WaterfallChartAggregatedFieldWells;

        [OutputConstructor]
        private DashboardWaterfallChartFieldWells(Outputs.DashboardWaterfallChartAggregatedFieldWells? waterfallChartAggregatedFieldWells)
        {
            WaterfallChartAggregatedFieldWells = waterfallChartAggregatedFieldWells;
        }
    }
}
