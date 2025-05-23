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
    public sealed class TemplateComputation
    {
        /// <summary>
        /// The forecast computation configuration.
        /// </summary>
        public readonly Outputs.TemplateForecastComputation? Forecast;
        /// <summary>
        /// The growth rate computation configuration.
        /// </summary>
        public readonly Outputs.TemplateGrowthRateComputation? GrowthRate;
        /// <summary>
        /// The maximum and minimum computation configuration.
        /// </summary>
        public readonly Outputs.TemplateMaximumMinimumComputation? MaximumMinimum;
        /// <summary>
        /// The metric comparison computation configuration.
        /// </summary>
        public readonly Outputs.TemplateMetricComparisonComputation? MetricComparison;
        /// <summary>
        /// The period over period computation configuration.
        /// </summary>
        public readonly Outputs.TemplatePeriodOverPeriodComputation? PeriodOverPeriod;
        /// <summary>
        /// The period to `DataSetIdentifier` computation configuration.
        /// </summary>
        public readonly Outputs.TemplatePeriodToDateComputation? PeriodToDate;
        /// <summary>
        /// The top movers and bottom movers computation configuration.
        /// </summary>
        public readonly Outputs.TemplateTopBottomMoversComputation? TopBottomMovers;
        /// <summary>
        /// The top ranked and bottom ranked computation configuration.
        /// </summary>
        public readonly Outputs.TemplateTopBottomRankedComputation? TopBottomRanked;
        /// <summary>
        /// The total aggregation computation configuration.
        /// </summary>
        public readonly Outputs.TemplateTotalAggregationComputation? TotalAggregation;
        /// <summary>
        /// The unique values computation configuration.
        /// </summary>
        public readonly Outputs.TemplateUniqueValuesComputation? UniqueValues;

        [OutputConstructor]
        private TemplateComputation(
            Outputs.TemplateForecastComputation? forecast,

            Outputs.TemplateGrowthRateComputation? growthRate,

            Outputs.TemplateMaximumMinimumComputation? maximumMinimum,

            Outputs.TemplateMetricComparisonComputation? metricComparison,

            Outputs.TemplatePeriodOverPeriodComputation? periodOverPeriod,

            Outputs.TemplatePeriodToDateComputation? periodToDate,

            Outputs.TemplateTopBottomMoversComputation? topBottomMovers,

            Outputs.TemplateTopBottomRankedComputation? topBottomRanked,

            Outputs.TemplateTotalAggregationComputation? totalAggregation,

            Outputs.TemplateUniqueValuesComputation? uniqueValues)
        {
            Forecast = forecast;
            GrowthRate = growthRate;
            MaximumMinimum = maximumMinimum;
            MetricComparison = metricComparison;
            PeriodOverPeriod = periodOverPeriod;
            PeriodToDate = periodToDate;
            TopBottomMovers = topBottomMovers;
            TopBottomRanked = topBottomRanked;
            TotalAggregation = totalAggregation;
            UniqueValues = uniqueValues;
        }
    }
}
