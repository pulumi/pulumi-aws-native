// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector
{
    /// <summary>
    /// An label for fraud detector.
    /// </summary>
    [AwsNativeResourceType("aws-native:frauddetector:Label")]
    public partial class Label : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The label ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the label was created.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The label description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the label was last updated.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Tags associated with this label.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LabelTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Label resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Label(string name, LabelArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:frauddetector:Label", name, args ?? new LabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Label(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:frauddetector:Label", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Label resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Label Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Label(name, id, options);
        }
    }

    public sealed class LabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The label description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the label.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.LabelTagArgs>? _tags;

        /// <summary>
        /// Tags associated with this label.
        /// </summary>
        public InputList<Inputs.LabelTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LabelTagArgs>());
            set => _tags = value;
        }

        public LabelArgs()
        {
        }
        public static new LabelArgs Empty => new LabelArgs();
    }
}