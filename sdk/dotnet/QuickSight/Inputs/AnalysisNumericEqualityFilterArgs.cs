// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisNumericEqualityFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The aggregation function of the filter.
        /// </summary>
        [Input("aggregationFunction")]
        public Input<Inputs.AnalysisAggregationFunctionArgs>? AggregationFunction { get; set; }

        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.AnalysisColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.
        /// </summary>
        [Input("defaultFilterControlConfiguration")]
        public Input<Inputs.AnalysisDefaultFilterControlConfigurationArgs>? DefaultFilterControlConfiguration { get; set; }

        /// <summary>
        /// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
        /// </summary>
        [Input("filterId", required: true)]
        public Input<string> FilterId { get; set; } = null!;

        /// <summary>
        /// The match operator that is used to determine if a filter should be applied.
        /// </summary>
        [Input("matchOperator", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisNumericEqualityMatchOperator> MatchOperator { get; set; } = null!;

        /// <summary>
        /// This option determines how null values should be treated when filtering data.
        /// 
        /// - `ALL_VALUES` : Include null values in filtered results.
        /// - `NULLS_ONLY` : Only include null values in filtered results.
        /// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
        /// </summary>
        [Input("nullOption", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisFilterNullOption> NullOption { get; set; } = null!;

        /// <summary>
        /// The parameter whose value should be used for the filter value.
        /// </summary>
        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        /// <summary>
        /// Select all of the values. Null is not the assigned value of select all.
        /// 
        /// - `FILTER_ALL_VALUES`
        /// </summary>
        [Input("selectAllOptions")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisNumericFilterSelectAllOptions>? SelectAllOptions { get; set; }

        /// <summary>
        /// The input value.
        /// </summary>
        [Input("value")]
        public Input<double>? Value { get; set; }

        public AnalysisNumericEqualityFilterArgs()
        {
        }
        public static new AnalysisNumericEqualityFilterArgs Empty => new AnalysisNumericEqualityFilterArgs();
    }
}
