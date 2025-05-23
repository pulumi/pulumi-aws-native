// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetAgentStatus
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::AgentStatus
        /// </summary>
        public static Task<GetAgentStatusResult> InvokeAsync(GetAgentStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentStatusResult>("aws-native:connect:getAgentStatus", args ?? new GetAgentStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::AgentStatus
        /// </summary>
        public static Output<GetAgentStatusResult> Invoke(GetAgentStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentStatusResult>("aws-native:connect:getAgentStatus", args ?? new GetAgentStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::AgentStatus
        /// </summary>
        public static Output<GetAgentStatusResult> Invoke(GetAgentStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentStatusResult>("aws-native:connect:getAgentStatus", args ?? new GetAgentStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the agent status.
        /// </summary>
        [Input("agentStatusArn", required: true)]
        public string AgentStatusArn { get; set; } = null!;

        public GetAgentStatusArgs()
        {
        }
        public static new GetAgentStatusArgs Empty => new GetAgentStatusArgs();
    }

    public sealed class GetAgentStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the agent status.
        /// </summary>
        [Input("agentStatusArn", required: true)]
        public Input<string> AgentStatusArn { get; set; } = null!;

        public GetAgentStatusInvokeArgs()
        {
        }
        public static new GetAgentStatusInvokeArgs Empty => new GetAgentStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentStatusResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the agent status.
        /// </summary>
        public readonly string? AgentStatusArn;
        /// <summary>
        /// The description of the status.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The display order of the status.
        /// </summary>
        public readonly int? DisplayOrder;
        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        public readonly string? InstanceArn;
        /// <summary>
        /// Last modified region.
        /// </summary>
        public readonly string? LastModifiedRegion;
        /// <summary>
        /// Last modified time.
        /// </summary>
        public readonly double? LastModifiedTime;
        /// <summary>
        /// The name of the status.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// A number indicating the reset order of the agent status.
        /// </summary>
        public readonly bool? ResetOrderNumber;
        /// <summary>
        /// The state of the status.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.AgentStatusState? State;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The type of agent status.
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.AgentStatusType? Type;

        [OutputConstructor]
        private GetAgentStatusResult(
            string? agentStatusArn,

            string? description,

            int? displayOrder,

            string? instanceArn,

            string? lastModifiedRegion,

            double? lastModifiedTime,

            string? name,

            bool? resetOrderNumber,

            Pulumi.AwsNative.Connect.AgentStatusState? state,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Pulumi.AwsNative.Connect.AgentStatusType? type)
        {
            AgentStatusArn = agentStatusArn;
            Description = description;
            DisplayOrder = displayOrder;
            InstanceArn = instanceArn;
            LastModifiedRegion = lastModifiedRegion;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            ResetOrderNumber = resetOrderNumber;
            State = state;
            Tags = tags;
            Type = type;
        }
    }
}
