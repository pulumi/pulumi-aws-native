// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    public sealed class AutoScalingGroupBaselinePerformanceFactorsRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU performance to consider, using an instance family as the baseline reference.
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.AutoScalingGroupCpuPerformanceFactorRequestArgs>? Cpu { get; set; }

        public AutoScalingGroupBaselinePerformanceFactorsRequestArgs()
        {
        }
        public static new AutoScalingGroupBaselinePerformanceFactorsRequestArgs Empty => new AutoScalingGroupBaselinePerformanceFactorsRequestArgs();
    }
}
