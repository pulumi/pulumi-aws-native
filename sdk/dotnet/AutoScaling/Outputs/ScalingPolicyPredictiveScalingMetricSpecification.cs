// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class ScalingPolicyPredictiveScalingMetricSpecification
    {
        public readonly Outputs.ScalingPolicyPredictiveScalingCustomizedCapacityMetric? CustomizedCapacityMetricSpecification;
        public readonly Outputs.ScalingPolicyPredictiveScalingCustomizedLoadMetric? CustomizedLoadMetricSpecification;
        public readonly Outputs.ScalingPolicyPredictiveScalingCustomizedScalingMetric? CustomizedScalingMetricSpecification;
        public readonly Outputs.ScalingPolicyPredictiveScalingPredefinedLoadMetric? PredefinedLoadMetricSpecification;
        public readonly Outputs.ScalingPolicyPredictiveScalingPredefinedMetricPair? PredefinedMetricPairSpecification;
        public readonly Outputs.ScalingPolicyPredictiveScalingPredefinedScalingMetric? PredefinedScalingMetricSpecification;
        public readonly double TargetValue;

        [OutputConstructor]
        private ScalingPolicyPredictiveScalingMetricSpecification(
            Outputs.ScalingPolicyPredictiveScalingCustomizedCapacityMetric? customizedCapacityMetricSpecification,

            Outputs.ScalingPolicyPredictiveScalingCustomizedLoadMetric? customizedLoadMetricSpecification,

            Outputs.ScalingPolicyPredictiveScalingCustomizedScalingMetric? customizedScalingMetricSpecification,

            Outputs.ScalingPolicyPredictiveScalingPredefinedLoadMetric? predefinedLoadMetricSpecification,

            Outputs.ScalingPolicyPredictiveScalingPredefinedMetricPair? predefinedMetricPairSpecification,

            Outputs.ScalingPolicyPredictiveScalingPredefinedScalingMetric? predefinedScalingMetricSpecification,

            double targetValue)
        {
            CustomizedCapacityMetricSpecification = customizedCapacityMetricSpecification;
            CustomizedLoadMetricSpecification = customizedLoadMetricSpecification;
            CustomizedScalingMetricSpecification = customizedScalingMetricSpecification;
            PredefinedLoadMetricSpecification = predefinedLoadMetricSpecification;
            PredefinedMetricPairSpecification = predefinedMetricPairSpecification;
            PredefinedScalingMetricSpecification = predefinedScalingMetricSpecification;
            TargetValue = targetValue;
        }
    }
}