// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Information that defines a state of a detector.
    /// </summary>
    public sealed class DetectorModelStateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When entering this state, perform these ``actions`` if the ``condition`` is TRUE.
        /// </summary>
        [Input("onEnter")]
        public Input<Inputs.DetectorModelOnEnterArgs>? OnEnter { get; set; }

        /// <summary>
        /// When exiting this state, perform these ``actions`` if the specified ``condition`` is ``TRUE``.
        /// </summary>
        [Input("onExit")]
        public Input<Inputs.DetectorModelOnExitArgs>? OnExit { get; set; }

        /// <summary>
        /// When an input is received and the ``condition`` is TRUE, perform the specified ``actions``.
        /// </summary>
        [Input("onInput")]
        public Input<Inputs.DetectorModelOnInputArgs>? OnInput { get; set; }

        /// <summary>
        /// The name of the state.
        /// </summary>
        [Input("stateName", required: true)]
        public Input<string> StateName { get; set; } = null!;

        public DetectorModelStateArgs()
        {
        }
        public static new DetectorModelStateArgs Empty => new DetectorModelStateArgs();
    }
}
