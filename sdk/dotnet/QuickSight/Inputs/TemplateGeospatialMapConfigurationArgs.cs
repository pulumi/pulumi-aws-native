// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateGeospatialMapConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldWells")]
        public Input<Inputs.TemplateGeospatialMapFieldWellsArgs>? FieldWells { get; set; }

        [Input("legend")]
        public Input<Inputs.TemplateLegendOptionsArgs>? Legend { get; set; }

        [Input("mapStyleOptions")]
        public Input<Inputs.TemplateGeospatialMapStyleOptionsArgs>? MapStyleOptions { get; set; }

        [Input("pointStyleOptions")]
        public Input<Inputs.TemplateGeospatialPointStyleOptionsArgs>? PointStyleOptions { get; set; }

        [Input("tooltip")]
        public Input<Inputs.TemplateTooltipOptionsArgs>? Tooltip { get; set; }

        [Input("visualPalette")]
        public Input<Inputs.TemplateVisualPaletteArgs>? VisualPalette { get; set; }

        [Input("windowOptions")]
        public Input<Inputs.TemplateGeospatialWindowOptionsArgs>? WindowOptions { get; set; }

        public TemplateGeospatialMapConfigurationArgs()
        {
        }
        public static new TemplateGeospatialMapConfigurationArgs Empty => new TemplateGeospatialMapConfigurationArgs();
    }
}