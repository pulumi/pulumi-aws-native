// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Creates an evaluation form for the specified CON instance.
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:EvaluationForm")]
    public partial class EvaluationForm : global::Pulumi.CustomResource
    {
        [Output("autoEvaluationConfiguration")]
        public Output<Outputs.EvaluationFormAutoEvaluationConfiguration?> AutoEvaluationConfiguration { get; private set; } = null!;

        /// <summary>
        /// The description of the evaluation form.
        ///  *Length Constraints*: Minimum length of 0. Maximum length of 1024.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the evaluation form.
        /// </summary>
        [Output("evaluationFormArn")]
        public Output<string> EvaluationFormArn { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// Items that are part of the evaluation form. The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
        ///  *Minimum size*: 1
        ///  *Maximum size*: 100
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Outputs.EvaluationFormBaseItem>> Items { get; private set; } = null!;

        /// <summary>
        /// A scoring strategy of the evaluation form.
        /// </summary>
        [Output("scoringStrategy")]
        public Output<Outputs.EvaluationFormScoringStrategy?> ScoringStrategy { get; private set; } = null!;

        /// <summary>
        /// The status of the evaluation form.
        ///  *Allowed values*: ``DRAFT`` | ``ACTIVE``
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Connect.EvaluationFormStatus> Status { get; private set; } = null!;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// A title of the evaluation form.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a EvaluationForm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EvaluationForm(string name, EvaluationFormArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:EvaluationForm", name, args ?? new EvaluationFormArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EvaluationForm(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:EvaluationForm", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EvaluationForm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EvaluationForm Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EvaluationForm(name, id, options);
        }
    }

    public sealed class EvaluationFormArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoEvaluationConfiguration")]
        public Input<Inputs.EvaluationFormAutoEvaluationConfigurationArgs>? AutoEvaluationConfiguration { get; set; }

        /// <summary>
        /// The description of the evaluation form.
        ///  *Length Constraints*: Minimum length of 0. Maximum length of 1024.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        [Input("items", required: true)]
        private InputList<Inputs.EvaluationFormBaseItemArgs>? _items;

        /// <summary>
        /// Items that are part of the evaluation form. The total number of sections and questions must not exceed 100 each. Questions must be contained in a section.
        ///  *Minimum size*: 1
        ///  *Maximum size*: 100
        /// </summary>
        public InputList<Inputs.EvaluationFormBaseItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.EvaluationFormBaseItemArgs>());
            set => _items = value;
        }

        /// <summary>
        /// A scoring strategy of the evaluation form.
        /// </summary>
        [Input("scoringStrategy")]
        public Input<Inputs.EvaluationFormScoringStrategyArgs>? ScoringStrategy { get; set; }

        /// <summary>
        /// The status of the evaluation form.
        ///  *Allowed values*: ``DRAFT`` | ``ACTIVE``
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AwsNative.Connect.EvaluationFormStatus> Status { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// A title of the evaluation form.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public EvaluationFormArgs()
        {
        }
        public static new EvaluationFormArgs Empty => new EvaluationFormArgs();
    }
}
