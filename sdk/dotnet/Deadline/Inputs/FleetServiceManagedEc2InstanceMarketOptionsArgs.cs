// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline.Inputs
{

    public sealed class FleetServiceManagedEc2InstanceMarketOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Deadline.FleetEc2MarketType> Type { get; set; } = null!;

        public FleetServiceManagedEc2InstanceMarketOptionsArgs()
        {
        }
        public static new FleetServiceManagedEc2InstanceMarketOptionsArgs Empty => new FleetServiceManagedEc2InstanceMarketOptionsArgs();
    }
}