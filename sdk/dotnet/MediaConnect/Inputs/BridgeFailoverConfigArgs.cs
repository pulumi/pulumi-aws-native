// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The settings for source failover.
    /// </summary>
    public sealed class BridgeFailoverConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of failover you choose for this flow. FAILOVER allows switching between different streams.
        /// </summary>
        [Input("failoverMode", required: true)]
        public Input<Pulumi.AwsNative.MediaConnect.BridgeFailoverModeEnum> FailoverMode { get; set; } = null!;

        /// <summary>
        /// The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.
        /// </summary>
        [Input("sourcePriority")]
        public Input<Inputs.BridgeSourcePriorityArgs>? SourcePriority { get; set; }

        [Input("state")]
        public Input<Pulumi.AwsNative.MediaConnect.BridgeFailoverConfigStateEnum>? State { get; set; }

        public BridgeFailoverConfigArgs()
        {
        }
        public static new BridgeFailoverConfigArgs Empty => new BridgeFailoverConfigArgs();
    }
}