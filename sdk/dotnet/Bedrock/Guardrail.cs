// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    /// <summary>
    /// Definition of AWS::Bedrock::Guardrail Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:bedrock:Guardrail")]
    public partial class Guardrail : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Messaging for when violations are detected in text
        /// </summary>
        [Output("blockedInputMessaging")]
        public Output<string> BlockedInputMessaging { get; private set; } = null!;

        /// <summary>
        /// Messaging for when violations are detected in text
        /// </summary>
        [Output("blockedOutputsMessaging")]
        public Output<string> BlockedOutputsMessaging { get; private set; } = null!;

        /// <summary>
        /// The content filter policies to configure for the guardrail.
        /// </summary>
        [Output("contentPolicyConfig")]
        public Output<Outputs.GuardrailContentPolicyConfig?> ContentPolicyConfig { get; private set; } = null!;

        [Output("contextualGroundingPolicyConfig")]
        public Output<Outputs.GuardrailContextualGroundingPolicyConfig?> ContextualGroundingPolicyConfig { get; private set; } = null!;

        /// <summary>
        /// Time Stamp
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("crossRegionConfig")]
        public Output<Outputs.GuardrailCrossRegionConfig?> CrossRegionConfig { get; private set; } = null!;

        /// <summary>
        /// Description of the guardrail or its version
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of failure recommendations
        /// </summary>
        [Output("failureRecommendations")]
        public Output<ImmutableArray<string>> FailureRecommendations { get; private set; } = null!;

        /// <summary>
        /// Arn representation for the guardrail
        /// </summary>
        [Output("guardrailArn")]
        public Output<string> GuardrailArn { get; private set; } = null!;

        /// <summary>
        /// Unique id for the guardrail
        /// </summary>
        [Output("guardrailId")]
        public Output<string> GuardrailId { get; private set; } = null!;

        /// <summary>
        /// The KMS key with which the guardrail was encrypted at rest
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// Name of the guardrail
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The sensitive information policy to configure for the guardrail.
        /// </summary>
        [Output("sensitiveInformationPolicyConfig")]
        public Output<Outputs.GuardrailSensitiveInformationPolicyConfig?> SensitiveInformationPolicyConfig { get; private set; } = null!;

        /// <summary>
        /// The status of the guardrail.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Bedrock.GuardrailStatus> Status { get; private set; } = null!;

        /// <summary>
        /// List of status reasons
        /// </summary>
        [Output("statusReasons")]
        public Output<ImmutableArray<string>> StatusReasons { get; private set; } = null!;

        /// <summary>
        /// List of Tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The topic policies to configure for the guardrail.
        /// </summary>
        [Output("topicPolicyConfig")]
        public Output<Outputs.GuardrailTopicPolicyConfig?> TopicPolicyConfig { get; private set; } = null!;

        /// <summary>
        /// Time Stamp
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Guardrail version
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The word policy you configure for the guardrail.
        /// </summary>
        [Output("wordPolicyConfig")]
        public Output<Outputs.GuardrailWordPolicyConfig?> WordPolicyConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Guardrail resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Guardrail(string name, GuardrailArgs args, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:Guardrail", name, args ?? new GuardrailArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Guardrail(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:Guardrail", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Guardrail resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Guardrail Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Guardrail(name, id, options);
        }
    }

    public sealed class GuardrailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Messaging for when violations are detected in text
        /// </summary>
        [Input("blockedInputMessaging", required: true)]
        public Input<string> BlockedInputMessaging { get; set; } = null!;

        /// <summary>
        /// Messaging for when violations are detected in text
        /// </summary>
        [Input("blockedOutputsMessaging", required: true)]
        public Input<string> BlockedOutputsMessaging { get; set; } = null!;

        /// <summary>
        /// The content filter policies to configure for the guardrail.
        /// </summary>
        [Input("contentPolicyConfig")]
        public Input<Inputs.GuardrailContentPolicyConfigArgs>? ContentPolicyConfig { get; set; }

        [Input("contextualGroundingPolicyConfig")]
        public Input<Inputs.GuardrailContextualGroundingPolicyConfigArgs>? ContextualGroundingPolicyConfig { get; set; }

        [Input("crossRegionConfig")]
        public Input<Inputs.GuardrailCrossRegionConfigArgs>? CrossRegionConfig { get; set; }

        /// <summary>
        /// Description of the guardrail or its version
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The KMS key with which the guardrail was encrypted at rest
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Name of the guardrail
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The sensitive information policy to configure for the guardrail.
        /// </summary>
        [Input("sensitiveInformationPolicyConfig")]
        public Input<Inputs.GuardrailSensitiveInformationPolicyConfigArgs>? SensitiveInformationPolicyConfig { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// List of Tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The topic policies to configure for the guardrail.
        /// </summary>
        [Input("topicPolicyConfig")]
        public Input<Inputs.GuardrailTopicPolicyConfigArgs>? TopicPolicyConfig { get; set; }

        /// <summary>
        /// The word policy you configure for the guardrail.
        /// </summary>
        [Input("wordPolicyConfig")]
        public Input<Inputs.GuardrailWordPolicyConfigArgs>? WordPolicyConfig { get; set; }

        public GuardrailArgs()
        {
        }
        public static new GuardrailArgs Empty => new GuardrailArgs();
    }
}
