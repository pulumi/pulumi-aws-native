// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Inputs
{

    /// <summary>
    /// Describes the scaling metric.
    /// </summary>
    public sealed class ScalingPolicyPredictiveScalingMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.ScalingPolicyPredictiveScalingMetricDimensionArgs>? _dimensions;

        /// <summary>
        /// Describes the dimensions of the metric.
        /// </summary>
        public InputList<Inputs.ScalingPolicyPredictiveScalingMetricDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.ScalingPolicyPredictiveScalingMetricDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The name of the metric.
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        /// <summary>
        /// The namespace of the metric.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public ScalingPolicyPredictiveScalingMetricArgs()
        {
        }
        public static new ScalingPolicyPredictiveScalingMetricArgs Empty => new ScalingPolicyPredictiveScalingMetricArgs();
    }
}
