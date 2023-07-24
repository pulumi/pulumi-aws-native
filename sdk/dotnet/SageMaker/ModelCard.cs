// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// Resource Type definition for AWS::SageMaker::ModelCard.
    /// </summary>
    [AwsNativeResourceType("aws-native:sagemaker:ModelCard")]
    public partial class ModelCard : global::Pulumi.CustomResource
    {
        [Output("content")]
        public Output<Outputs.ModelCardContent> Content { get; private set; } = null!;

        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        [Output("createdBy")]
        public Output<Outputs.ModelCardUserContext?> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// The date and time the model card was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        [Output("lastModifiedBy")]
        public Output<Outputs.ModelCardUserContext?> LastModifiedBy { get; private set; } = null!;

        /// <summary>
        /// The date and time the model card was last modified.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the successfully created model card.
        /// </summary>
        [Output("modelCardArn")]
        public Output<string> ModelCardArn { get; private set; } = null!;

        /// <summary>
        /// The unique name of the model card.
        /// </summary>
        [Output("modelCardName")]
        public Output<string> ModelCardName { get; private set; } = null!;

        /// <summary>
        /// The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
        /// </summary>
        [Output("modelCardProcessingStatus")]
        public Output<Pulumi.AwsNative.SageMaker.ModelCardProcessingStatus> ModelCardProcessingStatus { get; private set; } = null!;

        /// <summary>
        /// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
        /// </summary>
        [Output("modelCardStatus")]
        public Output<Pulumi.AwsNative.SageMaker.ModelCardStatus> ModelCardStatus { get; private set; } = null!;

        /// <summary>
        /// A version of the model card.
        /// </summary>
        [Output("modelCardVersion")]
        public Output<int> ModelCardVersion { get; private set; } = null!;

        [Output("securityConfig")]
        public Output<Outputs.ModelCardSecurityConfig?> SecurityConfig { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs used to manage metadata for model cards.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ModelCardTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ModelCard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModelCard(string name, ModelCardArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:ModelCard", name, args ?? new ModelCardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModelCard(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:ModelCard", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ModelCard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModelCard Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModelCard(name, id, options);
        }
    }

    public sealed class ModelCardArgs : global::Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<Inputs.ModelCardContentArgs> Content { get; set; } = null!;

        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        [Input("createdBy")]
        public Input<Inputs.ModelCardUserContextArgs>? CreatedBy { get; set; }

        /// <summary>
        /// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
        /// </summary>
        [Input("lastModifiedBy")]
        public Input<Inputs.ModelCardUserContextArgs>? LastModifiedBy { get; set; }

        /// <summary>
        /// The unique name of the model card.
        /// </summary>
        [Input("modelCardName")]
        public Input<string>? ModelCardName { get; set; }

        /// <summary>
        /// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
        /// </summary>
        [Input("modelCardStatus", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.ModelCardStatus> ModelCardStatus { get; set; } = null!;

        [Input("securityConfig")]
        public Input<Inputs.ModelCardSecurityConfigArgs>? SecurityConfig { get; set; }

        [Input("tags")]
        private InputList<Inputs.ModelCardTagArgs>? _tags;

        /// <summary>
        /// Key-value pairs used to manage metadata for model cards.
        /// </summary>
        public InputList<Inputs.ModelCardTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ModelCardTagArgs>());
            set => _tags = value;
        }

        public ModelCardArgs()
        {
        }
        public static new ModelCardArgs Empty => new ModelCardArgs();
    }
}