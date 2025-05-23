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
    /// Specifies the ``actions`` to be performed when the ``condition`` evaluates to TRUE.
    /// </summary>
    public sealed class DetectorModelEventArgs : global::Pulumi.ResourceArgs
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
        /// Optional. The Boolean expression that, when TRUE, causes the ``actions`` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// The name of the event.
        /// </summary>
        [Input("eventName", required: true)]
        public Input<string> EventName { get; set; } = null!;

        public DetectorModelEventArgs()
        {
        }
        public static new DetectorModelEventArgs Empty => new DetectorModelEventArgs();
    }
}
