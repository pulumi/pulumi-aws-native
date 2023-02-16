// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Specifies the `actions `performed and the next `state` entered when a `condition` evaluates to `TRUE`.
    /// </summary>
    public sealed class DetectorModelTransitionEventArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.DetectorModelActionArgs>? _actions;

        /// <summary>
        /// The actions to be performed.
        /// </summary>
        public InputList<Inputs.DetectorModelActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.DetectorModelActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// A Boolean expression that when `TRUE` causes the `actions` to be performed and the `nextState` to be entered.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// The name of the event.
        /// </summary>
        [Input("eventName", required: true)]
        public Input<string> EventName { get; set; } = null!;

        /// <summary>
        /// The next state to enter.
        /// </summary>
        [Input("nextState", required: true)]
        public Input<string> NextState { get; set; } = null!;

        public DetectorModelTransitionEventArgs()
        {
        }
        public static new DetectorModelTransitionEventArgs Empty => new DetectorModelTransitionEventArgs();
    }
}