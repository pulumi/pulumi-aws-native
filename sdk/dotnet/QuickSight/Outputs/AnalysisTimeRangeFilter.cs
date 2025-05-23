// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisTimeRangeFilter
    {
        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        public readonly Outputs.AnalysisColumnIdentifier Column;
        /// <summary>
        /// The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.
        /// </summary>
        public readonly Outputs.AnalysisDefaultFilterControlConfiguration? DefaultFilterControlConfiguration;
        /// <summary>
        /// The exclude period of the time range filter.
        /// </summary>
        public readonly Outputs.AnalysisExcludePeriodConfiguration? ExcludePeriodConfiguration;
        /// <summary>
        /// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
        /// </summary>
        public readonly string FilterId;
        /// <summary>
        /// Determines whether the maximum value in the filter value range should be included in the filtered results.
        /// </summary>
        public readonly bool? IncludeMaximum;
        /// <summary>
        /// Determines whether the minimum value in the filter value range should be included in the filtered results.
        /// </summary>
        public readonly bool? IncludeMinimum;
        /// <summary>
        /// This option determines how null values should be treated when filtering data.
        /// 
        /// - `ALL_VALUES` : Include null values in filtered results.
        /// - `NULLS_ONLY` : Only include null values in filtered results.
        /// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisFilterNullOption NullOption;
        /// <summary>
        /// The maximum value for the filter value range.
        /// </summary>
        public readonly Outputs.AnalysisTimeRangeFilterValue? RangeMaximumValue;
        /// <summary>
        /// The minimum value for the filter value range.
        /// </summary>
        public readonly Outputs.AnalysisTimeRangeFilterValue? RangeMinimumValue;
        /// <summary>
        /// The level of time precision that is used to aggregate `DateTime` values.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity? TimeGranularity;

        [OutputConstructor]
        private AnalysisTimeRangeFilter(
            Outputs.AnalysisColumnIdentifier column,

            Outputs.AnalysisDefaultFilterControlConfiguration? defaultFilterControlConfiguration,

            Outputs.AnalysisExcludePeriodConfiguration? excludePeriodConfiguration,

            string filterId,

            bool? includeMaximum,

            bool? includeMinimum,

            Pulumi.AwsNative.QuickSight.AnalysisFilterNullOption nullOption,

            Outputs.AnalysisTimeRangeFilterValue? rangeMaximumValue,

            Outputs.AnalysisTimeRangeFilterValue? rangeMinimumValue,

            Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity? timeGranularity)
        {
            Column = column;
            DefaultFilterControlConfiguration = defaultFilterControlConfiguration;
            ExcludePeriodConfiguration = excludePeriodConfiguration;
            FilterId = filterId;
            IncludeMaximum = includeMaximum;
            IncludeMinimum = includeMinimum;
            NullOption = nullOption;
            RangeMaximumValue = rangeMaximumValue;
            RangeMinimumValue = rangeMinimumValue;
            TimeGranularity = timeGranularity;
        }
    }
}
