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
    /// Internal mixin for flow node
    /// </summary>
    [OutputType]
    public sealed class FlowVersionFlowNode
    {
        /// <summary>
        /// Contains configurations for the node.
        /// </summary>
        public readonly object? Configuration;
        /// <summary>
        /// List of node inputs in a flow
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowVersionFlowNodeInput> Inputs;
        /// <summary>
        /// Name of a node in a flow
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of node outputs in a flow
        /// </summary>
        public readonly ImmutableArray<Outputs.FlowVersionFlowNodeOutput> Outputs;
        /// <summary>
        /// The type of node. This value must match the name of the key that you provide in the configuration you provide in the `FlowNodeConfiguration` field.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.FlowVersionFlowNodeType Type;

        [OutputConstructor]
        private FlowVersionFlowNode(
            object? configuration,

            ImmutableArray<Outputs.FlowVersionFlowNodeInput> inputs,

            string name,

            ImmutableArray<Outputs.FlowVersionFlowNodeOutput> outputs,

            Pulumi.AwsNative.Bedrock.FlowVersionFlowNodeType type)
        {
            Configuration = configuration;
            Inputs = inputs;
            Name = name;
            Outputs = outputs;
            Type = type;
        }
    }
}