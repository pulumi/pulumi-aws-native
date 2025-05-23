// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisWaterfallChartOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This option determines the total bar label of a waterfall visual.
        /// </summary>
        [Input("totalBarLabel")]
        public Input<string>? TotalBarLabel { get; set; }

        public AnalysisWaterfallChartOptionsArgs()
        {
        }
        public static new AnalysisWaterfallChartOptionsArgs Empty => new AnalysisWaterfallChartOptionsArgs();
    }
}
