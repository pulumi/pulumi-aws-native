// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Node configuration in a flow
    /// </summary>
    [OutputType]
    public sealed class FlowVersionFlowNodeConfiguration7Properties
    {
        public readonly Outputs.FlowVersionAgentFlowNodeConfiguration Agent;

        [OutputConstructor]
        private FlowVersionFlowNodeConfiguration7Properties(Outputs.FlowVersionAgentFlowNodeConfiguration agent)
        {
            Agent = agent;
        }
    }
}