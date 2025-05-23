// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisHistogramFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field well configuration of a histogram.
        /// </summary>
        [Input("histogramAggregatedFieldWells")]
        public Input<Inputs.AnalysisHistogramAggregatedFieldWellsArgs>? HistogramAggregatedFieldWells { get; set; }

        public AnalysisHistogramFieldWellsArgs()
        {
        }
        public static new AnalysisHistogramFieldWellsArgs Empty => new AnalysisHistogramFieldWellsArgs();
    }
}
