// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisRadarChartSeriesSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The area style settings of a radar chart.
        /// </summary>
        [Input("areaStyleSettings")]
        public Input<Inputs.AnalysisRadarChartAreaStyleSettingsArgs>? AreaStyleSettings { get; set; }

        public AnalysisRadarChartSeriesSettingsArgs()
        {
        }
        public static new AnalysisRadarChartSeriesSettingsArgs Empty => new AnalysisRadarChartSeriesSettingsArgs();
    }
}
