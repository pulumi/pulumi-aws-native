// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationAutoScaling.Outputs
{

    /// <summary>
    /// Describes the dimension of a metric.
    /// </summary>
    [OutputType]
    public sealed class ScalingPolicyPredictiveScalingMetricDimension
    {
        /// <summary>
        /// The name of the dimension.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value of the dimension.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ScalingPolicyPredictiveScalingMetricDimension(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
