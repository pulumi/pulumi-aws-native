// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardBarChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the arrangement of the bars. The orientation and arrangement of bars determine the type of bar that is used in the visual.
        /// </summary>
        [Input("barsArrangement")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardBarsArrangement>? BarsArrangement { get; set; }

        /// <summary>
        /// The label display options (grid line, range, scale, axis step) for bar chart category.
        /// </summary>
        [Input("categoryAxis")]
        public Input<Inputs.DashboardAxisDisplayOptionsArgs>? CategoryAxis { get; set; }

        /// <summary>
        /// The label options (label text, label visibility and sort icon visibility) for a bar chart.
        /// </summary>
        [Input("categoryLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? CategoryLabelOptions { get; set; }

        /// <summary>
        /// The label options (label text, label visibility and sort icon visibility) for a color that is used in a bar chart.
        /// </summary>
        [Input("colorLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? ColorLabelOptions { get; set; }

        [Input("contributionAnalysisDefaults")]
        private InputList<Inputs.DashboardContributionAnalysisDefaultArgs>? _contributionAnalysisDefaults;

        /// <summary>
        /// The contribution analysis (anomaly configuration) setup of the visual.
        /// </summary>
        public InputList<Inputs.DashboardContributionAnalysisDefaultArgs> ContributionAnalysisDefaults
        {
            get => _contributionAnalysisDefaults ?? (_contributionAnalysisDefaults = new InputList<Inputs.DashboardContributionAnalysisDefaultArgs>());
            set => _contributionAnalysisDefaults = value;
        }

        /// <summary>
        /// The options that determine if visual data labels are displayed.
        /// </summary>
        [Input("dataLabels")]
        public Input<Inputs.DashboardDataLabelOptionsArgs>? DataLabels { get; set; }

        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        [Input("fieldWells")]
        public Input<Inputs.DashboardBarChartFieldWellsArgs>? FieldWells { get; set; }

        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        [Input("interactions")]
        public Input<Inputs.DashboardVisualInteractionOptionsArgs>? Interactions { get; set; }

        /// <summary>
        /// The legend display setup of the visual.
        /// </summary>
        [Input("legend")]
        public Input<Inputs.DashboardLegendOptionsArgs>? Legend { get; set; }

        /// <summary>
        /// The orientation of the bars in a bar chart visual. There are two valid values in this structure:
        /// 
        /// - `HORIZONTAL` : Used for charts that have horizontal bars. Visuals that use this value are horizontal bar charts, horizontal stacked bar charts, and horizontal stacked 100% bar charts.
        /// - `VERTICAL` : Used for charts that have vertical bars. Visuals that use this value are vertical bar charts, vertical stacked bar charts, and vertical stacked 100% bar charts.
        /// </summary>
        [Input("orientation")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardBarChartOrientation>? Orientation { get; set; }

        [Input("referenceLines")]
        private InputList<Inputs.DashboardReferenceLineArgs>? _referenceLines;

        /// <summary>
        /// The reference line setup of the visual.
        /// </summary>
        public InputList<Inputs.DashboardReferenceLineArgs> ReferenceLines
        {
            get => _referenceLines ?? (_referenceLines = new InputList<Inputs.DashboardReferenceLineArgs>());
            set => _referenceLines = value;
        }

        /// <summary>
        /// The small multiples setup for the visual.
        /// </summary>
        [Input("smallMultiplesOptions")]
        public Input<Inputs.DashboardSmallMultiplesOptionsArgs>? SmallMultiplesOptions { get; set; }

        /// <summary>
        /// The sort configuration of a `BarChartVisual` .
        /// </summary>
        [Input("sortConfiguration")]
        public Input<Inputs.DashboardBarChartSortConfigurationArgs>? SortConfiguration { get; set; }

        /// <summary>
        /// The tooltip display setup of the visual.
        /// </summary>
        [Input("tooltip")]
        public Input<Inputs.DashboardTooltipOptionsArgs>? Tooltip { get; set; }

        /// <summary>
        /// The label display options (grid line, range, scale, axis step) for a bar chart value.
        /// </summary>
        [Input("valueAxis")]
        public Input<Inputs.DashboardAxisDisplayOptionsArgs>? ValueAxis { get; set; }

        /// <summary>
        /// The label options (label text, label visibility and sort icon visibility) for a bar chart value.
        /// </summary>
        [Input("valueLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? ValueLabelOptions { get; set; }

        /// <summary>
        /// The palette (chart color) display setup of the visual.
        /// </summary>
        [Input("visualPalette")]
        public Input<Inputs.DashboardVisualPaletteArgs>? VisualPalette { get; set; }

        public DashboardBarChartConfigurationArgs()
        {
        }
        public static new DashboardBarChartConfigurationArgs Empty => new DashboardBarChartConfigurationArgs();
    }
}
