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
    /// Describes a scaling metric for a predictive scaling policy. 
    ///  When returned in the output of ``DescribePolicies``, it indicates that a predictive scaling policy uses individually specified load and scaling metrics instead of a metric pair.
    /// </summary>
    public sealed class ScalingPolicyPredictiveScalingPredefinedScalingMetricArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The metric type.
        /// </summary>
        [Input("predefinedMetricType", required: true)]
        public Input<string> PredefinedMetricType { get; set; } = null!;

        /// <summary>
        /// A label that uniquely identifies a specific target group from which to determine the average request count.
        /// </summary>
        [Input("resourceLabel")]
        public Input<string>? ResourceLabel { get; set; }

        public ScalingPolicyPredictiveScalingPredefinedScalingMetricArgs()
        {
        }
        public static new ScalingPolicyPredictiveScalingPredefinedScalingMetricArgs Empty => new ScalingPolicyPredictiveScalingPredefinedScalingMetricArgs();
    }
}
