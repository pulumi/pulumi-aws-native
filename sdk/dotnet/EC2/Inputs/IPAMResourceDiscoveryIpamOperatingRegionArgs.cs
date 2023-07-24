// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// The regions IPAM Resource Discovery is enabled for. Allows for monitoring.
    /// </summary>
    public sealed class IPAMResourceDiscoveryIpamOperatingRegionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the region.
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        public IPAMResourceDiscoveryIpamOperatingRegionArgs()
        {
        }
        public static new IPAMResourceDiscoveryIpamOperatingRegionArgs Empty => new IPAMResourceDiscoveryIpamOperatingRegionArgs();
    }
}