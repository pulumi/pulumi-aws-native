// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisColumnTooltipItem
    {
        public readonly Outputs.AnalysisAggregationFunction? Aggregation;
        public readonly Outputs.AnalysisColumnIdentifier Column;
        public readonly string? Label;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisColumnTooltipItem(
            Outputs.AnalysisAggregationFunction? aggregation,

            Outputs.AnalysisColumnIdentifier column,

            string? label,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            Aggregation = aggregation;
            Column = column;
            Label = label;
            Visibility = visibility;
        }
    }
}