// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Inputs
{

    public sealed class ClusterPlacementTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        public ClusterPlacementTypeArgs()
        {
        }
        public static new ClusterPlacementTypeArgs Empty => new ClusterPlacementTypeArgs();
    }
}