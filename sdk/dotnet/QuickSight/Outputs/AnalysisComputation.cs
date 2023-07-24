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
    public sealed class AnalysisComputation
    {
        public readonly Outputs.AnalysisForecastComputation? Forecast;
        public readonly Outputs.AnalysisGrowthRateComputation? GrowthRate;
        public readonly Outputs.AnalysisMaximumMinimumComputation? MaximumMinimum;
        public readonly Outputs.AnalysisMetricComparisonComputation? MetricComparison;
        public readonly Outputs.AnalysisPeriodOverPeriodComputation? PeriodOverPeriod;
        public readonly Outputs.AnalysisPeriodToDateComputation? PeriodToDate;
        public readonly Outputs.AnalysisTopBottomMoversComputation? TopBottomMovers;
        public readonly Outputs.AnalysisTopBottomRankedComputation? TopBottomRanked;
        public readonly Outputs.AnalysisTotalAggregationComputation? TotalAggregation;
        public readonly Outputs.AnalysisUniqueValuesComputation? UniqueValues;

        [OutputConstructor]
        private AnalysisComputation(
            Outputs.AnalysisForecastComputation? forecast,

            Outputs.AnalysisGrowthRateComputation? growthRate,

            Outputs.AnalysisMaximumMinimumComputation? maximumMinimum,

            Outputs.AnalysisMetricComparisonComputation? metricComparison,

            Outputs.AnalysisPeriodOverPeriodComputation? periodOverPeriod,

            Outputs.AnalysisPeriodToDateComputation? periodToDate,

            Outputs.AnalysisTopBottomMoversComputation? topBottomMovers,

            Outputs.AnalysisTopBottomRankedComputation? topBottomRanked,

            Outputs.AnalysisTotalAggregationComputation? totalAggregation,

            Outputs.AnalysisUniqueValuesComputation? uniqueValues)
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