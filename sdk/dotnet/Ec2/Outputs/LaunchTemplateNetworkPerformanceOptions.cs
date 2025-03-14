// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateNetworkPerformanceOptions
    {
        /// <summary>
        /// Specifies the performance options of your instance or sets it to default.
        /// </summary>
        public readonly string? BandwidthWeighting;

        [OutputConstructor]
        private LaunchTemplateNetworkPerformanceOptions(string? bandwidthWeighting)
        {
            BandwidthWeighting = bandwidthWeighting;
        }
    }
}
