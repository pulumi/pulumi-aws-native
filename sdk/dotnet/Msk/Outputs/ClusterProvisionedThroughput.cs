// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterProvisionedThroughput
    {
        /// <summary>
        /// Provisioned throughput is on or off.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second.
        /// </summary>
        public readonly int? VolumeThroughput;

        [OutputConstructor]
        private ClusterProvisionedThroughput(
            bool? enabled,

            int? volumeThroughput)
        {
            Enabled = enabled;
            VolumeThroughput = volumeThroughput;
        }
    }
}
