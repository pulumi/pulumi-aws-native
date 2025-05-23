// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class Ec2FleetBaselinePerformanceFactorsRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU performance to consider, using an instance family as the baseline reference.
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.Ec2FleetCpuPerformanceFactorRequestArgs>? Cpu { get; set; }

        public Ec2FleetBaselinePerformanceFactorsRequestArgs()
        {
        }
        public static new Ec2FleetBaselinePerformanceFactorsRequestArgs Empty => new Ec2FleetBaselinePerformanceFactorsRequestArgs();
    }
}
