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
    public sealed class AnalysisSankeyDiagramChartConfiguration
    {
        /// <summary>
        /// The data label configuration of a sankey diagram.
        /// </summary>
        public readonly Outputs.AnalysisDataLabelOptions? DataLabels;
        /// <summary>
        /// The field well configuration of a sankey diagram.
        /// </summary>
        public readonly Outputs.AnalysisSankeyDiagramFieldWells? FieldWells;
        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        public readonly Outputs.AnalysisVisualInteractionOptions? Interactions;
        /// <summary>
        /// The sort configuration of a sankey diagram.
        /// </summary>
        public readonly Outputs.AnalysisSankeyDiagramSortConfiguration? SortConfiguration;

        [OutputConstructor]
        private AnalysisSankeyDiagramChartConfiguration(
            Outputs.AnalysisDataLabelOptions? dataLabels,

            Outputs.AnalysisSankeyDiagramFieldWells? fieldWells,

            Outputs.AnalysisVisualInteractionOptions? interactions,

            Outputs.AnalysisSankeyDiagramSortConfiguration? sortConfiguration)
        {
            DataLabels = dataLabels;
            FieldWells = fieldWells;
            Interactions = interactions;
            SortConfiguration = sortConfiguration;
        }
    }
}
