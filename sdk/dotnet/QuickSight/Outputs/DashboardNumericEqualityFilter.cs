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
    public sealed class DashboardNumericEqualityFilter
    {
        public readonly Outputs.DashboardAggregationFunction? AggregationFunction;
        public readonly Outputs.DashboardColumnIdentifier Column;
        public readonly string FilterId;
        public readonly Pulumi.AwsNative.QuickSight.DashboardNumericEqualityMatchOperator MatchOperator;
        public readonly Pulumi.AwsNative.QuickSight.DashboardFilterNullOption NullOption;
        public readonly string? ParameterName;
        public readonly Pulumi.AwsNative.QuickSight.DashboardNumericFilterSelectAllOptions? SelectAllOptions;
        public readonly double? Value;

        [OutputConstructor]
        private DashboardNumericEqualityFilter(
            Outputs.DashboardAggregationFunction? aggregationFunction,

            Outputs.DashboardColumnIdentifier column,

            string filterId,

            Pulumi.AwsNative.QuickSight.DashboardNumericEqualityMatchOperator matchOperator,

            Pulumi.AwsNative.QuickSight.DashboardFilterNullOption nullOption,

            string? parameterName,

            Pulumi.AwsNative.QuickSight.DashboardNumericFilterSelectAllOptions? selectAllOptions,

            double? value)
        {
            AggregationFunction = aggregationFunction;
            Column = column;
            FilterId = filterId;
            MatchOperator = matchOperator;
            NullOption = nullOption;
            ParameterName = parameterName;
            SelectAllOptions = selectAllOptions;
            Value = value;
        }
    }
}