// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class SpotFleetBaselinePerformanceFactorsRequest
    {
        /// <summary>
        /// The CPU performance to consider, using an instance family as the baseline reference.
        /// </summary>
        public readonly Outputs.SpotFleetCpuPerformanceFactorRequest? Cpu;

        [OutputConstructor]
        private SpotFleetBaselinePerformanceFactorsRequest(Outputs.SpotFleetCpuPerformanceFactorRequest? cpu)
        {
            Cpu = cpu;
        }
    }
}
