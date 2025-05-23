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
    public sealed class AnalysisForecastComputation
    {
        /// <summary>
        /// The ID for a computation.
        /// </summary>
        public readonly string ComputationId;
        /// <summary>
        /// The custom seasonality value setup of a forecast computation.
        /// </summary>
        public readonly double? CustomSeasonalityValue;
        /// <summary>
        /// The lower boundary setup of a forecast computation.
        /// </summary>
        public readonly double? LowerBoundary;
        /// <summary>
        /// The name of a computation.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The periods backward setup of a forecast computation.
        /// </summary>
        public readonly double? PeriodsBackward;
        /// <summary>
        /// The periods forward setup of a forecast computation.
        /// </summary>
        public readonly double? PeriodsForward;
        /// <summary>
        /// The prediction interval setup of a forecast computation.
        /// </summary>
        public readonly double? PredictionInterval;
        /// <summary>
        /// The seasonality setup of a forecast computation. Choose one of the following options:
        /// 
        /// - `AUTOMATIC`
        /// - `CUSTOM` : Checks the custom seasonality value.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisForecastComputationSeasonality? Seasonality;
        /// <summary>
        /// The time field that is used in a computation.
        /// </summary>
        public readonly Outputs.AnalysisDimensionField? Time;
        /// <summary>
        /// The upper boundary setup of a forecast computation.
        /// </summary>
        public readonly double? UpperBoundary;
        /// <summary>
        /// The value field that is used in a computation.
        /// </summary>
        public readonly Outputs.AnalysisMeasureField? Value;

        [OutputConstructor]
        private AnalysisForecastComputation(
            string computationId,

            double? customSeasonalityValue,

            double? lowerBoundary,

            string? name,

            double? periodsBackward,

            double? periodsForward,

            double? predictionInterval,

            Pulumi.AwsNative.QuickSight.AnalysisForecastComputationSeasonality? seasonality,

            Outputs.AnalysisDimensionField? time,

            double? upperBoundary,

            Outputs.AnalysisMeasureField? value)
        {
            ComputationId = computationId;
            CustomSeasonalityValue = customSeasonalityValue;
            LowerBoundary = lowerBoundary;
            Name = name;
            PeriodsBackward = periodsBackward;
            PeriodsForward = periodsForward;
            PredictionInterval = predictionInterval;
            Seasonality = seasonality;
            Time = time;
            UpperBoundary = upperBoundary;
            Value = value;
        }
    }
}
