// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFunnelChartFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field well configuration of a `FunnelChartVisual` .
        /// </summary>
        [Input("funnelChartAggregatedFieldWells")]
        public Input<Inputs.DashboardFunnelChartAggregatedFieldWellsArgs>? FunnelChartAggregatedFieldWells { get; set; }

        public DashboardFunnelChartFieldWellsArgs()
        {
        }
        public static new DashboardFunnelChartFieldWellsArgs Empty => new DashboardFunnelChartFieldWellsArgs();
    }
}
