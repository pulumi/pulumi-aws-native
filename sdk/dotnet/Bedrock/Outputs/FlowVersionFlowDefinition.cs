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
    /// Flow definition
    /// </summary>
    [OutputType]
    public sealed class FlowVersionFlowDefinition
    {
        /// <summary>
        /// List of connections
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowVersionFlowConnection> Connections;
        /// <summary>
        /// List of nodes in a flow
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowVersionFlowNode> Nodes;

        [OutputConstructor]
        private FlowVersionFlowDefinition(
            ImmutableArray<Outputs.FlowVersionFlowConnection> connections,

            ImmutableArray<Outputs.FlowVersionFlowNode> nodes)
        {
            Connections = connections;
            Nodes = nodes;
        }
    }
}