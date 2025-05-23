// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetAgent
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::Agent.
        /// </summary>
        public static Task<GetAgentResult> InvokeAsync(GetAgentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAgentResult>("aws-native:datasync:getAgent", args ?? new GetAgentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::Agent.
        /// </summary>
        public static Output<GetAgentResult> Invoke(GetAgentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentResult>("aws-native:datasync:getAgent", args ?? new GetAgentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::Agent.
        /// </summary>
        public static Output<GetAgentResult> Invoke(GetAgentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAgentResult>("aws-native:datasync:getAgent", args ?? new GetAgentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DataSync Agent ARN.
        /// </summary>
        [Input("agentArn", required: true)]
        public string AgentArn { get; set; } = null!;

        public GetAgentArgs()
        {
        }
        public static new GetAgentArgs Empty => new GetAgentArgs();
    }

    public sealed class GetAgentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The DataSync Agent ARN.
        /// </summary>
        [Input("agentArn", required: true)]
        public Input<string> AgentArn { get; set; } = null!;

        public GetAgentInvokeArgs()
        {
        }
        public static new GetAgentInvokeArgs Empty => new GetAgentInvokeArgs();
    }


    [OutputType]
    public sealed class GetAgentResult
    {
        /// <summary>
        /// The DataSync Agent ARN.
        /// </summary>
        public readonly string? AgentArn;
        /// <summary>
        /// The name configured for the agent. Text reference used to identify the agent in the console.
        /// </summary>
        public readonly string? AgentName;
        /// <summary>
        /// The service endpoints that the agent will connect to.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.AgentEndpointType? EndpointType;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAgentResult(
            string? agentArn,

            string? agentName,

            Pulumi.AwsNative.DataSync.AgentEndpointType? endpointType,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AgentArn = agentArn;
            AgentName = agentName;
            EndpointType = endpointType;
            Tags = tags;
        }
    }
}
