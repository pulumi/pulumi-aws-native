// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateHistogramConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("binOptions")]
        public Input<Inputs.TemplateHistogramBinOptionsArgs>? BinOptions { get; set; }

        [Input("dataLabels")]
        public Input<Inputs.TemplateDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplateHistogramFieldWellsArgs>? FieldWells { get; set; }

        [Input("tooltip")]
        public Input<Inputs.TemplateTooltipOptionsArgs>? Tooltip { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.TemplateVisualPaletteArgs>? VisualPalette { get; set; }

        [Input("xAxisDisplayOptions")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? XAxisDisplayOptions { get; set; }

        [Input("xAxisLabelOptions")]
        public Input<Inputs.TemplateChartAxisLabelOptionsArgs>? XAxisLabelOptions { get; set; }

        [Input("yAxisDisplayOptions")]
        public Input<Inputs.TemplateAxisDisplayOptionsArgs>? YAxisDisplayOptions { get; set; }

        public TemplateHistogramConfigurationArgs()
        {
        }
        public static new TemplateHistogramConfigurationArgs Empty => new TemplateHistogramConfigurationArgs();
    }
}