// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector
{
    /// <summary>
    /// A resource schema for an EventType in Amazon Fraud Detector.
    /// </summary>
    [AwsNativeResourceType("aws-native:frauddetector:EventType")]
    public partial class EventType : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the event type.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The time when the event type was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The description of the event type.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The event type entity types.
        /// </summary>
        [Output("entityTypes")]
        public Output<ImmutableArray<Outputs.EventTypeEntityType>> EntityTypes { get; private set; } = null!;

        /// <summary>
        /// The event type event variables.
        /// </summary>
        [Output("eventVariables")]
        public Output<ImmutableArray<Outputs.EventTypeEventVariable>> EventVariables { get; private set; } = null!;

        /// <summary>
        /// The event type labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.EventTypeLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// The time when the event type was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// The name for the event type
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventType(string name, EventTypeArgs args, CustomResourceOptions? options = null)
            : base("aws-native:frauddetector:EventType", name, args ?? new EventTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:frauddetector:EventType", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventType(name, id, options);
        }
    }

    public sealed class EventTypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the event type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("entityTypes", required: true)]
        private InputList<Inputs.EventTypeEntityTypeArgs>? _entityTypes;

        /// <summary>
        /// The event type entity types.
        /// </summary>
        public InputList<Inputs.EventTypeEntityTypeArgs> EntityTypes
        {
            get => _entityTypes ?? (_entityTypes = new InputList<Inputs.EventTypeEntityTypeArgs>());
            set => _entityTypes = value;
        }

        [Input("eventVariables", required: true)]
        private InputList<Inputs.EventTypeEventVariableArgs>? _eventVariables;

        /// <summary>
        /// The event type event variables.
        /// </summary>
        public InputList<Inputs.EventTypeEventVariableArgs> EventVariables
        {
            get => _eventVariables ?? (_eventVariables = new InputList<Inputs.EventTypeEventVariableArgs>());
            set => _eventVariables = value;
        }

        [Input("labels", required: true)]
        private InputList<Inputs.EventTypeLabelArgs>? _labels;

        /// <summary>
        /// The event type labels.
        /// </summary>
        public InputList<Inputs.EventTypeLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.EventTypeLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name for the event type
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Tags associated with this event type.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public EventTypeArgs()
        {
        }
        public static new EventTypeArgs Empty => new EventTypeArgs();
    }
}
