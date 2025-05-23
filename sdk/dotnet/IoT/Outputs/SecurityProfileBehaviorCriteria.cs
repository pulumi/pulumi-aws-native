// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// The criteria by which the behavior is determined to be normal.
    /// </summary>
    [OutputType]
    public sealed class SecurityProfileBehaviorCriteria
    {
        /// <summary>
        /// The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).
        /// </summary>
        public readonly Pulumi.AwsNative.IoT.SecurityProfileBehaviorCriteriaComparisonOperator? ComparisonOperator;
        /// <summary>
        /// If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.
        /// </summary>
        public readonly int? ConsecutiveDatapointsToAlarm;
        /// <summary>
        /// If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.
        /// </summary>
        public readonly int? ConsecutiveDatapointsToClear;
        /// <summary>
        /// Use this to specify the time duration over which the behavior is evaluated.
        /// </summary>
        public readonly int? DurationSeconds;
        /// <summary>
        /// The confidence level of the detection model.
        /// </summary>
        public readonly Outputs.SecurityProfileMachineLearningDetectionConfig? MlDetectionConfig;
        /// <summary>
        /// A statistical ranking (percentile)that indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.
        /// </summary>
        public readonly Outputs.SecurityProfileStatisticalThreshold? StatisticalThreshold;
        /// <summary>
        /// The value to be compared with the `metric` .
        /// </summary>
        public readonly Outputs.SecurityProfileMetricValue? Value;

        [OutputConstructor]
        private SecurityProfileBehaviorCriteria(
            Pulumi.AwsNative.IoT.SecurityProfileBehaviorCriteriaComparisonOperator? comparisonOperator,

            int? consecutiveDatapointsToAlarm,

            int? consecutiveDatapointsToClear,

            int? durationSeconds,

            Outputs.SecurityProfileMachineLearningDetectionConfig? mlDetectionConfig,

            Outputs.SecurityProfileStatisticalThreshold? statisticalThreshold,

            Outputs.SecurityProfileMetricValue? value)
        {
            ComparisonOperator = comparisonOperator;
            ConsecutiveDatapointsToAlarm = consecutiveDatapointsToAlarm;
            ConsecutiveDatapointsToClear = consecutiveDatapointsToClear;
            DurationSeconds = durationSeconds;
            MlDetectionConfig = mlDetectionConfig;
            StatisticalThreshold = statisticalThreshold;
            Value = value;
        }
    }
}
