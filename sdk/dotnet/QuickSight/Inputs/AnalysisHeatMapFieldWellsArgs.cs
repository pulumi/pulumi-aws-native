// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisHeatMapFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregated field wells of a heat map.
        /// </summary>
        [Input("heatMapAggregatedFieldWells")]
        public Input<Inputs.AnalysisHeatMapAggregatedFieldWellsArgs>? HeatMapAggregatedFieldWells { get; set; }

        public AnalysisHeatMapFieldWellsArgs()
        {
        }
        public static new AnalysisHeatMapFieldWellsArgs Empty => new AnalysisHeatMapFieldWellsArgs();
    }
}
