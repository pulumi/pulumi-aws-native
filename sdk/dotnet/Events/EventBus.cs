// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events
{
    /// <summary>
    /// Resource type definition for AWS::Events::EventBus
    /// </summary>
    [AwsNativeResourceType("aws-native:events:EventBus")]
    public partial class EventBus : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the event bus.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Dead Letter Queue for the event bus.
        /// </summary>
        [Output("deadLetterConfig")]
        public Output<Outputs.DeadLetterConfigProperties?> DeadLetterConfig { get; private set; } = null!;

        /// <summary>
        /// The description of the event bus.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
        /// </summary>
        [Output("eventSourceName")]
        public Output<string?> EventSourceName { get; private set; } = null!;

        /// <summary>
        /// Kms Key Identifier used to encrypt events at rest in the event bus.
        /// </summary>
        [Output("kmsKeyIdentifier")]
        public Output<string?> KmsKeyIdentifier { get; private set; } = null!;

        /// <summary>
        /// The name of the event bus.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A JSON string that describes the permission policy statement for the event bus.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
        /// </summary>
        [Output("policy")]
        public Output<object?> Policy { get; private set; } = null!;

        /// <summary>
        /// Any tags assigned to the event bus.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EventBus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventBus(string name, EventBusArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:events:EventBus", name, args ?? new EventBusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventBus(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:events:EventBus", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EventBus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventBus Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventBus(name, id, options);
        }
    }

    public sealed class EventBusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dead Letter Queue for the event bus.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.DeadLetterConfigPropertiesArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// The description of the event bus.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
        /// </summary>
        [Input("eventSourceName")]
        public Input<string>? EventSourceName { get; set; }

        /// <summary>
        /// Kms Key Identifier used to encrypt events at rest in the event bus.
        /// </summary>
        [Input("kmsKeyIdentifier")]
        public Input<string>? KmsKeyIdentifier { get; set; }

        /// <summary>
        /// The name of the event bus.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A JSON string that describes the permission policy statement for the event bus.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
        /// </summary>
        [Input("policy")]
        public Input<object>? Policy { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Any tags assigned to the event bus.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public EventBusArgs()
        {
        }
        public static new EventBusArgs Empty => new EventBusArgs();
    }
}
