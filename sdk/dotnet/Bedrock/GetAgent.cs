// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    public static class GetAgent
    {
        /// <summary>
        /// Definition of AWS::Bedrock::Agent Resource Type
        /// </summary>
        public static Task<GetAgentResult> InvokeAsync(GetAgentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentResult>("aws-native:bedrock:getAgent", args ?? new GetAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Bedrock::Agent Resource Type
        /// </summary>
        public static Output<GetAgentResult> Invoke(GetAgentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentResult>("aws-native:bedrock:getAgent", args ?? new GetAgentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Bedrock::Agent Resource Type
        /// </summary>
        public static Output<GetAgentResult> Invoke(GetAgentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentResult>("aws-native:bedrock:getAgent", args ?? new GetAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        public GetAgentArgs()
        {
        }
        public static new GetAgentArgs Empty => new GetAgentArgs();
    }

    public sealed class GetAgentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        public GetAgentInvokeArgs()
        {
        }
        public static new GetAgentInvokeArgs Empty => new GetAgentInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentResult
    {
        /// <summary>
        /// List of ActionGroups
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentActionGroup> ActionGroups;
        /// <summary>
        /// Arn representation of the Agent.
        /// </summary>
        public readonly string? AgentArn;
        /// <summary>
        /// The agent's collaboration settings.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.AgentCollaboration? AgentCollaboration;
        /// <summary>
        /// List of Agent Collaborators
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentCollaborator> AgentCollaborators;
        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        public readonly string? AgentId;
        /// <summary>
        /// Name for a resource.
        /// </summary>
        public readonly string? AgentName;
        /// <summary>
        /// ARN of a IAM role.
        /// </summary>
        public readonly string? AgentResourceRoleArn;
        /// <summary>
        /// The status of the agent and whether it is ready for use. The following statuses are possible:
        /// 
        /// - CREATING – The agent is being created.
        /// - PREPARING – The agent is being prepared.
        /// - PREPARED – The agent is prepared and ready to be invoked.
        /// - NOT_PREPARED – The agent has been created but not yet prepared.
        /// - FAILED – The agent API operation failed.
        /// - UPDATING – The agent is being updated.
        /// - DELETING – The agent is being deleted.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.AgentStatus? AgentStatus;
        /// <summary>
        /// Draft Agent Version.
        /// </summary>
        public readonly string? AgentVersion;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Contains custom orchestration configurations for the agent.
        /// </summary>
        public readonly Outputs.AgentCustomOrchestration? CustomOrchestration;
        /// <summary>
        /// A KMS key ARN
        /// </summary>
        public readonly string? CustomerEncryptionKeyArn;
        /// <summary>
        /// Description of the Resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Failure Reasons for Error.
        /// </summary>
        public readonly ImmutableArray<string> FailureReasons;
        /// <summary>
        /// The foundation model used for orchestration by the agent.
        /// </summary>
        public readonly string? FoundationModel;
        /// <summary>
        /// Details about the guardrail associated with the agent.
        /// </summary>
        public readonly Outputs.AgentGuardrailConfiguration? GuardrailConfiguration;
        /// <summary>
        /// Max Session Time.
        /// </summary>
        public readonly double? IdleSessionTtlInSeconds;
        /// <summary>
        /// Instruction for the agent.
        /// </summary>
        public readonly string? Instruction;
        /// <summary>
        /// List of Agent Knowledge Bases
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentKnowledgeBase> KnowledgeBases;
        /// <summary>
        /// Contains memory configuration for the agent.
        /// </summary>
        public readonly Outputs.AgentMemoryConfiguration? MemoryConfiguration;
        /// <summary>
        /// Specifies the orchestration strategy for the agent.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.AgentOrchestrationType? OrchestrationType;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? PreparedAt;
        /// <summary>
        /// Contains configurations to override prompt templates in different parts of an agent sequence. For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html) .
        /// </summary>
        public readonly Outputs.AgentPromptOverrideConfiguration? PromptOverrideConfiguration;
        /// <summary>
        /// The recommended actions users can take to resolve an error in failureReasons.
        /// </summary>
        public readonly ImmutableArray<string> RecommendedActions;
        /// <summary>
        /// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
        /// 
        /// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
        /// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:
        /// 
        /// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
        /// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
        /// </summary>
        public readonly ImmutableDictionary<string, string>? TestAliasTags;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetAgentResult(
            ImmutableArray<Outputs.AgentActionGroup> actionGroups,

            string? agentArn,

            Pulumi.AwsNative.Bedrock.AgentCollaboration? agentCollaboration,

            ImmutableArray<Outputs.AgentCollaborator> agentCollaborators,

            string? agentId,

            string? agentName,

            string? agentResourceRoleArn,

            Pulumi.AwsNative.Bedrock.AgentStatus? agentStatus,

            string? agentVersion,

            string? createdAt,

            Outputs.AgentCustomOrchestration? customOrchestration,

            string? customerEncryptionKeyArn,

            string? description,

            ImmutableArray<string> failureReasons,

            string? foundationModel,

            Outputs.AgentGuardrailConfiguration? guardrailConfiguration,

            double? idleSessionTtlInSeconds,

            string? instruction,

            ImmutableArray<Outputs.AgentKnowledgeBase> knowledgeBases,

            Outputs.AgentMemoryConfiguration? memoryConfiguration,

            Pulumi.AwsNative.Bedrock.AgentOrchestrationType? orchestrationType,

            string? preparedAt,

            Outputs.AgentPromptOverrideConfiguration? promptOverrideConfiguration,

            ImmutableArray<string> recommendedActions,

            ImmutableDictionary<string, string>? tags,

            ImmutableDictionary<string, string>? testAliasTags,

            string? updatedAt)
        {
            ActionGroups = actionGroups;
            AgentArn = agentArn;
            AgentCollaboration = agentCollaboration;
            AgentCollaborators = agentCollaborators;
            AgentId = agentId;
            AgentName = agentName;
            AgentResourceRoleArn = agentResourceRoleArn;
            AgentStatus = agentStatus;
            AgentVersion = agentVersion;
            CreatedAt = createdAt;
            CustomOrchestration = customOrchestration;
            CustomerEncryptionKeyArn = customerEncryptionKeyArn;
            Description = description;
            FailureReasons = failureReasons;
            FoundationModel = foundationModel;
            GuardrailConfiguration = guardrailConfiguration;
            IdleSessionTtlInSeconds = idleSessionTtlInSeconds;
            Instruction = instruction;
            KnowledgeBases = knowledgeBases;
            MemoryConfiguration = memoryConfiguration;
            OrchestrationType = orchestrationType;
            PreparedAt = preparedAt;
            PromptOverrideConfiguration = promptOverrideConfiguration;
            RecommendedActions = recommendedActions;
            Tags = tags;
            TestAliasTags = testAliasTags;
            UpdatedAt = updatedAt;
        }
    }
}
