// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTotalAggregationFunctionArgs : global::Pulumi.ResourceArgs
    {
        [Input("simpleTotalAggregationFunction")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardSimpleTotalAggregationFunction>? SimpleTotalAggregationFunction { get; set; }

        public DashboardTotalAggregationFunctionArgs()
        {
        }
        public static new DashboardTotalAggregationFunctionArgs Empty => new DashboardTotalAggregationFunctionArgs();
    }
}