// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The transport parameters associated with an incoming media stream.
    /// </summary>
    public sealed class FlowInputConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port that the flow listens on for an incoming media stream.
        /// </summary>
        [Input("inputPort", required: true)]
        public Input<int> InputPort { get; set; } = null!;

        /// <summary>
        /// The VPC interface where the media stream comes in from.
        /// </summary>
        [Input("interface", required: true)]
        public Input<Inputs.FlowInterfaceArgs> Interface { get; set; } = null!;

        public FlowInputConfigurationArgs()
        {
        }
        public static new FlowInputConfigurationArgs Empty => new FlowInputConfigurationArgs();
    }
}