// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    /// <summary>
    /// Definition of AWS::Bedrock::Agent Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:bedrock:Agent")]
    public partial class Agent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of ActionGroups
        /// </summary>
        [Output("actionGroups")]
        public Output<ImmutableArray<Outputs.AgentActionGroup>> ActionGroups { get; private set; } = null!;

        /// <summary>
        /// Arn representation of the Agent.
        /// </summary>
        [Output("agentArn")]
        public Output<string> AgentArn { get; private set; } = null!;

        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Output("agentId")]
        public Output<string> AgentId { get; private set; } = null!;

        /// <summary>
        /// Name for a resource.
        /// </summary>
        [Output("agentName")]
        public Output<string> AgentName { get; private set; } = null!;

        /// <summary>
        /// ARN of a IAM role.
        /// </summary>
        [Output("agentResourceRoleArn")]
        public Output<string?> AgentResourceRoleArn { get; private set; } = null!;

        [Output("agentStatus")]
        public Output<Pulumi.AwsNative.Bedrock.AgentStatus> AgentStatus { get; private set; } = null!;

        /// <summary>
        /// Draft Agent Version.
        /// </summary>
        [Output("agentVersion")]
        public Output<string> AgentVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to automatically prepare after creating or updating the agent.
        /// </summary>
        [Output("autoPrepare")]
        public Output<bool?> AutoPrepare { get; private set; } = null!;

        /// <summary>
        /// Time Stamp.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A KMS key ARN
        /// </summary>
        [Output("customerEncryptionKeyArn")]
        public Output<string?> CustomerEncryptionKeyArn { get; private set; } = null!;

        /// <summary>
        /// Description of the Resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Failure Reasons for Error.
        /// </summary>
        [Output("failureReasons")]
        public Output<ImmutableArray<string>> FailureReasons { get; private set; } = null!;

        /// <summary>
        /// ARN or name of a Bedrock model.
        /// </summary>
        [Output("foundationModel")]
        public Output<string?> FoundationModel { get; private set; } = null!;

        /// <summary>
        /// Max Session Time.
        /// </summary>
        [Output("idleSessionTtlInSeconds")]
        public Output<double?> IdleSessionTtlInSeconds { get; private set; } = null!;

        /// <summary>
        /// Instruction for the agent.
        /// </summary>
        [Output("instruction")]
        public Output<string?> Instruction { get; private set; } = null!;

        /// <summary>
        /// List of Agent Knowledge Bases
        /// </summary>
        [Output("knowledgeBases")]
        public Output<ImmutableArray<Outputs.AgentKnowledgeBase>> KnowledgeBases { get; private set; } = null!;

        /// <summary>
        /// Time Stamp.
        /// </summary>
        [Output("preparedAt")]
        public Output<string> PreparedAt { get; private set; } = null!;

        [Output("promptOverrideConfiguration")]
        public Output<Outputs.AgentPromptOverrideConfiguration?> PromptOverrideConfiguration { get; private set; } = null!;

        /// <summary>
        /// The recommended actions users can take to resolve an error in failureReasons.
        /// </summary>
        [Output("recommendedActions")]
        public Output<ImmutableArray<string>> RecommendedActions { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to allow deleting agent while it is in use.
        /// </summary>
        [Output("skipResourceInUseCheckOnDelete")]
        public Output<bool?> SkipResourceInUseCheckOnDelete { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Time Stamp.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Agent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Agent(string name, AgentArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:Agent", name, args ?? new AgentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Agent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:Agent", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Agent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Agent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Agent(name, id, options);
        }
    }

    public sealed class AgentArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionGroups")]
        private InputList<Inputs.AgentActionGroupArgs>? _actionGroups;

        /// <summary>
        /// List of ActionGroups
        /// </summary>
        public InputList<Inputs.AgentActionGroupArgs> ActionGroups
        {
            get => _actionGroups ?? (_actionGroups = new InputList<Inputs.AgentActionGroupArgs>());
            set => _actionGroups = value;
        }

        /// <summary>
        /// Name for a resource.
        /// </summary>
        [Input("agentName")]
        public Input<string>? AgentName { get; set; }

        /// <summary>
        /// ARN of a IAM role.
        /// </summary>
        [Input("agentResourceRoleArn")]
        public Input<string>? AgentResourceRoleArn { get; set; }

        /// <summary>
        /// Specifies whether to automatically prepare after creating or updating the agent.
        /// </summary>
        [Input("autoPrepare")]
        public Input<bool>? AutoPrepare { get; set; }

        /// <summary>
        /// A KMS key ARN
        /// </summary>
        [Input("customerEncryptionKeyArn")]
        public Input<string>? CustomerEncryptionKeyArn { get; set; }

        /// <summary>
        /// Description of the Resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ARN or name of a Bedrock model.
        /// </summary>
        [Input("foundationModel")]
        public Input<string>? FoundationModel { get; set; }

        /// <summary>
        /// Max Session Time.
        /// </summary>
        [Input("idleSessionTtlInSeconds")]
        public Input<double>? IdleSessionTtlInSeconds { get; set; }

        /// <summary>
        /// Instruction for the agent.
        /// </summary>
        [Input("instruction")]
        public Input<string>? Instruction { get; set; }

        [Input("knowledgeBases")]
        private InputList<Inputs.AgentKnowledgeBaseArgs>? _knowledgeBases;

        /// <summary>
        /// List of Agent Knowledge Bases
        /// </summary>
        public InputList<Inputs.AgentKnowledgeBaseArgs> KnowledgeBases
        {
            get => _knowledgeBases ?? (_knowledgeBases = new InputList<Inputs.AgentKnowledgeBaseArgs>());
            set => _knowledgeBases = value;
        }

        [Input("promptOverrideConfiguration")]
        public Input<Inputs.AgentPromptOverrideConfigurationArgs>? PromptOverrideConfiguration { get; set; }

        /// <summary>
        /// Specifies whether to allow deleting agent while it is in use.
        /// </summary>
        [Input("skipResourceInUseCheckOnDelete")]
        public Input<bool>? SkipResourceInUseCheckOnDelete { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AgentArgs()
        {
        }
        public static new AgentArgs Empty => new AgentArgs();
    }
}