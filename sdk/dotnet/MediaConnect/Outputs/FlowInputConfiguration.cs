// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// The transport parameters associated with an incoming media stream.
    /// </summary>
    [OutputType]
    public sealed class FlowInputConfiguration
    {
        /// <summary>
        /// The port that the flow listens on for an incoming media stream.
        /// </summary>
        public readonly int InputPort;
        /// <summary>
        /// The VPC interface where the media stream comes in from.
        /// </summary>
        public readonly Outputs.FlowInterface Interface;

        [OutputConstructor]
        private FlowInputConfiguration(
            int inputPort,

            Outputs.FlowInterface @interface)
        {
            InputPort = inputPort;
            Interface = @interface;
        }
    }
}
