// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    public static class GetAgentAlias
    {
        /// <summary>
        /// Definition of AWS::Bedrock::AgentAlias Resource Type
        /// </summary>
        public static Task<GetAgentAliasResult> InvokeAsync(GetAgentAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentAliasResult>("aws-native:bedrock:getAgentAlias", args ?? new GetAgentAliasArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Bedrock::AgentAlias Resource Type
        /// </summary>
        public static Output<GetAgentAliasResult> Invoke(GetAgentAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentAliasResult>("aws-native:bedrock:getAgentAlias", args ?? new GetAgentAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentAliasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id for an Agent Alias generated at the server side.
        /// </summary>
        [Input("agentAliasId", required: true)]
        public string AgentAliasId { get; set; } = null!;

        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        public GetAgentAliasArgs()
        {
        }
        public static new GetAgentAliasArgs Empty => new GetAgentAliasArgs();
    }

    public sealed class GetAgentAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id for an Agent Alias generated at the server side.
        /// </summary>
        [Input("agentAliasId", required: true)]
        public Input<string> AgentAliasId { get; set; } = null!;

        /// <summary>
        /// Identifier for a resource.
        /// </summary>
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        public GetAgentAliasInvokeArgs()
        {
        }
        public static new GetAgentAliasInvokeArgs Empty => new GetAgentAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentAliasResult
    {
        /// <summary>
        /// Arn representation of the Agent Alias.
        /// </summary>
        public readonly string? AgentAliasArn;
        /// <summary>
        /// The list of history events for an alias for an Agent.
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentAliasHistoryEvent> AgentAliasHistoryEvents;
        /// <summary>
        /// Id for an Agent Alias generated at the server side.
        /// </summary>
        public readonly string? AgentAliasId;
        /// <summary>
        /// Name for a resource.
        /// </summary>
        public readonly string? AgentAliasName;
        public readonly Pulumi.AwsNative.Bedrock.AgentAliasStatus? AgentAliasStatus;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Description of the Resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Routing configuration for an Agent alias.
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentAliasRoutingConfigurationListItem> RoutingConfiguration;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private GetAgentAliasResult(
            string? agentAliasArn,

            ImmutableArray<Outputs.AgentAliasHistoryEvent> agentAliasHistoryEvents,

            string? agentAliasId,

            string? agentAliasName,

            Pulumi.AwsNative.Bedrock.AgentAliasStatus? agentAliasStatus,

            string? createdAt,

            string? description,

            ImmutableArray<Outputs.AgentAliasRoutingConfigurationListItem> routingConfiguration,

            ImmutableDictionary<string, string>? tags,

            string? updatedAt)
        {
            AgentAliasArn = agentAliasArn;
            AgentAliasHistoryEvents = agentAliasHistoryEvents;
            AgentAliasId = agentAliasId;
            AgentAliasName = agentAliasName;
            AgentAliasStatus = agentAliasStatus;
            CreatedAt = createdAt;
            Description = description;
            RoutingConfiguration = routingConfiguration;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}