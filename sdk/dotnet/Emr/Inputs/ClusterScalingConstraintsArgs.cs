// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Inputs
{

    public sealed class ClusterScalingConstraintsArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxCapacity", required: true)]
        public Input<int> MaxCapacity { get; set; } = null!;

        [Input("minCapacity", required: true)]
        public Input<int> MinCapacity { get; set; } = null!;

        public ClusterScalingConstraintsArgs()
        {
        }
        public static new ClusterScalingConstraintsArgs Empty => new ClusterScalingConstraintsArgs();
    }
}