// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// Rule that controls how a fleet is scaled. Scaling policies are uniquely identified by the combination of name and fleet ID.
    /// </summary>
    [OutputType]
    public sealed class ContainerFleetScalingPolicy
    {
        /// <summary>
        /// Comparison operator to use when measuring a metric against the threshold value.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyComparisonOperator? ComparisonOperator;
        /// <summary>
        /// Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.
        /// </summary>
        public readonly int? EvaluationPeriods;
        /// <summary>
        /// Name of the Amazon GameLift-defined metric that is used to trigger a scaling adjustment.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyMetricName MetricName;
        /// <summary>
        /// A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of scaling policy to create. For a target-based policy, set the parameter MetricName to 'PercentAvailableGameSessions' and specify a TargetConfiguration. For a rule-based policy set the following parameters: MetricName, ComparisonOperator, Threshold, EvaluationPeriods, ScalingAdjustmentType, and ScalingAdjustment.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyPolicyType? PolicyType;
        /// <summary>
        /// Amount of adjustment to make, based on the scaling adjustment type.
        /// </summary>
        public readonly int? ScalingAdjustment;
        /// <summary>
        /// The type of adjustment to make to a fleet's instance count.
        /// </summary>
        public readonly Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyScalingAdjustmentType? ScalingAdjustmentType;
        /// <summary>
        /// An object that contains settings for a target-based scaling policy.
        /// </summary>
        public readonly Outputs.ContainerFleetTargetConfiguration? TargetConfiguration;
        /// <summary>
        /// Metric value used to trigger a scaling event.
        /// </summary>
        public readonly double? Threshold;

        [OutputConstructor]
        private ContainerFleetScalingPolicy(
            Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyComparisonOperator? comparisonOperator,

            int? evaluationPeriods,

            Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyMetricName metricName,

            string name,

            Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyPolicyType? policyType,

            int? scalingAdjustment,

            Pulumi.AwsNative.GameLift.ContainerFleetScalingPolicyScalingAdjustmentType? scalingAdjustmentType,

            Outputs.ContainerFleetTargetConfiguration? targetConfiguration,

            double? threshold)
        {
            ComparisonOperator = comparisonOperator;
            EvaluationPeriods = evaluationPeriods;
            MetricName = metricName;
            Name = name;
            PolicyType = policyType;
            ScalingAdjustment = scalingAdjustment;
            ScalingAdjustmentType = scalingAdjustmentType;
            TargetConfiguration = targetConfiguration;
            Threshold = threshold;
        }
    }
}
