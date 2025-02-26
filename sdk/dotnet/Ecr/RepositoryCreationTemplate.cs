// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecr
{
    /// <summary>
    /// AWS::ECR::RepositoryCreationTemplate is used to create repository with configuration from a pre-defined template.
    /// </summary>
    [AwsNativeResourceType("aws-native:ecr:RepositoryCreationTemplate")]
    public partial class RepositoryCreationTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
        /// </summary>
        [Output("appliedFor")]
        public Output<ImmutableArray<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateAppliedForItem>> AppliedFor { get; private set; } = null!;

        /// <summary>
        /// Create timestamp of the template.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The ARN of the role to be assumed by ECR. This role must be in the same account as the registry that you are configuring.
        /// </summary>
        [Output("customRoleArn")]
        public Output<string?> CustomRoleArn { get; private set; } = null!;

        /// <summary>
        /// The description of the template.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The encryption configuration associated with the repository creation template.
        /// </summary>
        [Output("encryptionConfiguration")]
        public Output<Outputs.RepositoryCreationTemplateEncryptionConfiguration?> EncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// The image tag mutability setting for the repository.
        /// </summary>
        [Output("imageTagMutability")]
        public Output<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateImageTagMutability?> ImageTagMutability { get; private set; } = null!;

        /// <summary>
        /// The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
        /// </summary>
        [Output("lifecyclePolicy")]
        public Output<string?> LifecyclePolicy { get; private set; } = null!;

        /// <summary>
        /// The prefix use to match the repository name and apply the template.
        /// </summary>
        [Output("prefix")]
        public Output<string> Prefix { get; private set; } = null!;

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
        /// </summary>
        [Output("repositoryPolicy")]
        public Output<string?> RepositoryPolicy { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("resourceTags")]
        public Output<ImmutableArray<Outputs.RepositoryCreationTemplateTag>> ResourceTags { get; private set; } = null!;

        /// <summary>
        /// Update timestamp of the template.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryCreationTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryCreationTemplate(string name, RepositoryCreationTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ecr:RepositoryCreationTemplate", name, args ?? new RepositoryCreationTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryCreationTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecr:RepositoryCreationTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "prefix",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RepositoryCreationTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryCreationTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RepositoryCreationTemplate(name, id, options);
        }
    }

    public sealed class RepositoryCreationTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("appliedFor", required: true)]
        private InputList<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateAppliedForItem>? _appliedFor;

        /// <summary>
        /// A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.
        /// </summary>
        public InputList<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateAppliedForItem> AppliedFor
        {
            get => _appliedFor ?? (_appliedFor = new InputList<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateAppliedForItem>());
            set => _appliedFor = value;
        }

        /// <summary>
        /// The ARN of the role to be assumed by ECR. This role must be in the same account as the registry that you are configuring.
        /// </summary>
        [Input("customRoleArn")]
        public Input<string>? CustomRoleArn { get; set; }

        /// <summary>
        /// The description of the template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The encryption configuration associated with the repository creation template.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.RepositoryCreationTemplateEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// The image tag mutability setting for the repository.
        /// </summary>
        [Input("imageTagMutability")]
        public Input<Pulumi.AwsNative.Ecr.RepositoryCreationTemplateImageTagMutability>? ImageTagMutability { get; set; }

        /// <summary>
        /// The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
        /// </summary>
        [Input("lifecyclePolicy")]
        public Input<string>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The prefix use to match the repository name and apply the template.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html
        /// </summary>
        [Input("repositoryPolicy")]
        public Input<string>? RepositoryPolicy { get; set; }

        [Input("resourceTags")]
        private InputList<Inputs.RepositoryCreationTemplateTagArgs>? _resourceTags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.RepositoryCreationTemplateTagArgs> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputList<Inputs.RepositoryCreationTemplateTagArgs>());
            set => _resourceTags = value;
        }

        public RepositoryCreationTemplateArgs()
        {
        }
        public static new RepositoryCreationTemplateArgs Empty => new RepositoryCreationTemplateArgs();
    }
}
