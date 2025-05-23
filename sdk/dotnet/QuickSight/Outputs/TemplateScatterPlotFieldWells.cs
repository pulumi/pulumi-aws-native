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
    public sealed class TemplateScatterPlotFieldWells
    {
        /// <summary>
        /// The aggregated field wells of a scatter plot. The x and y-axes of scatter plots with aggregated field wells are aggregated by category, label, or both.
        /// </summary>
        public readonly Outputs.TemplateScatterPlotCategoricallyAggregatedFieldWells? ScatterPlotCategoricallyAggregatedFieldWells;
        /// <summary>
        /// The unaggregated field wells of a scatter plot. The x and y-axes of these scatter plots are unaggregated.
        /// </summary>
        public readonly Outputs.TemplateScatterPlotUnaggregatedFieldWells? ScatterPlotUnaggregatedFieldWells;

        [OutputConstructor]
        private TemplateScatterPlotFieldWells(
            Outputs.TemplateScatterPlotCategoricallyAggregatedFieldWells? scatterPlotCategoricallyAggregatedFieldWells,

            Outputs.TemplateScatterPlotUnaggregatedFieldWells? scatterPlotUnaggregatedFieldWells)
        {
            ScatterPlotCategoricallyAggregatedFieldWells = scatterPlotCategoricallyAggregatedFieldWells;
            ScatterPlotUnaggregatedFieldWells = scatterPlotUnaggregatedFieldWells;
        }
    }
}
