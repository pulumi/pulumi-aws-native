// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPieChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The label options of the group/color that is displayed in a pie chart.
        /// </summary>
        [Input("categoryLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? CategoryLabelOptions { get; set; }

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
        /// The options that determine the shape of the chart. This option determines whether the chart is a pie chart or a donut chart.
        /// </summary>
        [Input("donutOptions")]
        public Input<Inputs.DashboardDonutOptionsArgs>? DonutOptions { get; set; }

        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        [Input("fieldWells")]
        public Input<Inputs.DashboardPieChartFieldWellsArgs>? FieldWells { get; set; }

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
        /// The small multiples setup for the visual.
        /// </summary>
        [Input("smallMultiplesOptions")]
        public Input<Inputs.DashboardSmallMultiplesOptionsArgs>? SmallMultiplesOptions { get; set; }

        /// <summary>
        /// The sort configuration of a pie chart.
        /// </summary>
        [Input("sortConfiguration")]
        public Input<Inputs.DashboardPieChartSortConfigurationArgs>? SortConfiguration { get; set; }

        /// <summary>
        /// The tooltip display setup of the visual.
        /// </summary>
        [Input("tooltip")]
        public Input<Inputs.DashboardTooltipOptionsArgs>? Tooltip { get; set; }

        /// <summary>
        /// The label options for the value that is displayed in a pie chart.
        /// </summary>
        [Input("valueLabelOptions")]
        public Input<Inputs.DashboardChartAxisLabelOptionsArgs>? ValueLabelOptions { get; set; }

        /// <summary>
        /// The palette (chart color) display setup of the visual.
        /// </summary>
        [Input("visualPalette")]
        public Input<Inputs.DashboardVisualPaletteArgs>? VisualPalette { get; set; }

        public DashboardPieChartConfigurationArgs()
        {
        }
        public static new DashboardPieChartConfigurationArgs Empty => new DashboardPieChartConfigurationArgs();
    }
}
