// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Inputs
{

    /// <summary>
    /// Describes the dimension of a metric.
    /// </summary>
    public sealed class ScalingPolicyTargetTrackingMetricDimensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the dimension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The value of the dimension.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ScalingPolicyTargetTrackingMetricDimensionArgs()
        {
        }
        public static new ScalingPolicyTargetTrackingMetricDimensionArgs Empty => new ScalingPolicyTargetTrackingMetricDimensionArgs();
    }
}