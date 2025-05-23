// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisRadarChartAreaStyleSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The visibility settings of a radar chart.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisRadarChartAreaStyleSettingsArgs()
        {
        }
        public static new AnalysisRadarChartAreaStyleSettingsArgs Empty => new AnalysisRadarChartAreaStyleSettingsArgs();
    }
}
