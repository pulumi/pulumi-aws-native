// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Outputs
{

    [OutputType]
    public sealed class FleetServiceManagedEc2FleetConfiguration
    {
        public readonly Outputs.FleetServiceManagedEc2InstanceCapabilities InstanceCapabilities;
        public readonly Outputs.FleetServiceManagedEc2InstanceMarketOptions InstanceMarketOptions;

        [OutputConstructor]
        private FleetServiceManagedEc2FleetConfiguration(
            Outputs.FleetServiceManagedEc2InstanceCapabilities instanceCapabilities,

            Outputs.FleetServiceManagedEc2InstanceMarketOptions instanceMarketOptions)
        {
            InstanceCapabilities = instanceCapabilities;
            InstanceMarketOptions = instanceMarketOptions;
        }
    }
}