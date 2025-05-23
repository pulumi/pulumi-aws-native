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
    public sealed class TemplateLineChartSeriesSettings
    {
        /// <summary>
        /// Line styles options for a line series in `LineChartVisual` .
        /// </summary>
        public readonly Outputs.TemplateLineChartLineStyleSettings? LineStyleSettings;
        /// <summary>
        /// Marker styles options for a line series in `LineChartVisual` .
        /// </summary>
        public readonly Outputs.TemplateLineChartMarkerStyleSettings? MarkerStyleSettings;

        [OutputConstructor]
        private TemplateLineChartSeriesSettings(
            Outputs.TemplateLineChartLineStyleSettings? lineStyleSettings,

            Outputs.TemplateLineChartMarkerStyleSettings? markerStyleSettings)
        {
            LineStyleSettings = lineStyleSettings;
            MarkerStyleSettings = markerStyleSettings;
        }
    }
}
