// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// Rule that controls how a fleet is scaled. Scaling policies are uniquely identified by the combination of name and fleet ID.
    /// </summary>
    public sealed class FleetScalingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comparison operator to use when measuring a metric against the threshold value.
        /// </summary>
        [Input("comparisonOperator")]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyComparisonOperator>? ComparisonOperator { get; set; }

        /// <summary>
        /// Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.
        /// </summary>
        [Input("evaluationPeriods")]
        public Input<int>? EvaluationPeriods { get; set; }

        /// <summary>
        /// The fleet location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the Amazon GameLift-defined metric that is used to trigger a scaling adjustment.
        /// </summary>
        [Input("metricName", required: true)]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyMetricName> MetricName { get; set; } = null!;

        /// <summary>
        /// A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The type of scaling policy to create. For a target-based policy, set the parameter MetricName to 'PercentAvailableGameSessions' and specify a TargetConfiguration. For a rule-based policy set the following parameters: MetricName, ComparisonOperator, Threshold, EvaluationPeriods, ScalingAdjustmentType, and ScalingAdjustment.
        /// </summary>
        [Input("policyType")]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyPolicyType>? PolicyType { get; set; }

        /// <summary>
        /// Amount of adjustment to make, based on the scaling adjustment type.
        /// </summary>
        [Input("scalingAdjustment")]
        public Input<int>? ScalingAdjustment { get; set; }

        /// <summary>
        /// The type of adjustment to make to a fleet's instance count.
        /// </summary>
        [Input("scalingAdjustmentType")]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyScalingAdjustmentType>? ScalingAdjustmentType { get; set; }

        /// <summary>
        /// Current status of the scaling policy. The scaling policy can be in force only when in an ACTIVE status. Scaling policies can be suspended for individual fleets. If the policy is suspended for a fleet, the policy status does not change.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyStatus>? Status { get; set; }

        /// <summary>
        /// An object that contains settings for a target-based scaling policy.
        /// </summary>
        [Input("targetConfiguration")]
        public Input<Inputs.FleetTargetConfigurationArgs>? TargetConfiguration { get; set; }

        /// <summary>
        /// Metric value used to trigger a scaling event.
        /// </summary>
        [Input("threshold")]
        public Input<double>? Threshold { get; set; }

        /// <summary>
        /// The current status of the fleet's scaling policies in a requested fleet location. The status PENDING_UPDATE indicates that an update was requested for the fleet but has not yet been completed for the location.
        /// </summary>
        [Input("updateStatus")]
        public Input<Pulumi.AwsNative.GameLift.FleetScalingPolicyUpdateStatus>? UpdateStatus { get; set; }

        public FleetScalingPolicyArgs()
        {
        }
        public static new FleetScalingPolicyArgs Empty => new FleetScalingPolicyArgs();
    }
}
