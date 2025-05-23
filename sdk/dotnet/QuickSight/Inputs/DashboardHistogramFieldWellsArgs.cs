// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardHistogramFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field well configuration of a histogram.
        /// </summary>
        [Input("histogramAggregatedFieldWells")]
        public Input<Inputs.DashboardHistogramAggregatedFieldWellsArgs>? HistogramAggregatedFieldWells { get; set; }

        public DashboardHistogramFieldWellsArgs()
        {
        }
        public static new DashboardHistogramFieldWellsArgs Empty => new DashboardHistogramFieldWellsArgs();
    }
}
