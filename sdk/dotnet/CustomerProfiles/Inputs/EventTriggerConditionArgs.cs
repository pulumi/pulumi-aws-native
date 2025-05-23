// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    /// <summary>
    /// Specifies the circumstances under which the event should trigger the destination.
    /// </summary>
    public sealed class EventTriggerConditionArgs : global::Pulumi.ResourceArgs
    {
        [Input("eventTriggerDimensions", required: true)]
        private InputList<Inputs.EventTriggerDimensionArgs>? _eventTriggerDimensions;
        public InputList<Inputs.EventTriggerDimensionArgs> EventTriggerDimensions
        {
            get => _eventTriggerDimensions ?? (_eventTriggerDimensions = new InputList<Inputs.EventTriggerDimensionArgs>());
            set => _eventTriggerDimensions = value;
        }

        [Input("logicalOperator", required: true)]
        public Input<Pulumi.AwsNative.CustomerProfiles.EventTriggerLogicalOperator> LogicalOperator { get; set; } = null!;

        public EventTriggerConditionArgs()
        {
        }
        public static new EventTriggerConditionArgs Empty => new EventTriggerConditionArgs();
    }
}
