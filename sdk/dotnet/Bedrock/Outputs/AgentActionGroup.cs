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
    /// Contains the information of an Agent Action Group
    /// </summary>
    [OutputType]
    public sealed class AgentActionGroup
    {
        public readonly Union<Outputs.AgentActionGroupExecutor0Properties, Outputs.AgentActionGroupExecutor1Properties>? ActionGroupExecutor;
        /// <summary>
        /// Name of the action group
        /// </summary>
        public readonly string ActionGroupName;
        public readonly Pulumi.AwsNative.Bedrock.AgentActionGroupState? ActionGroupState;
        public readonly Union<Outputs.AgentApiSchema0Properties, Outputs.AgentApiSchema1Properties>? ApiSchema;
        /// <summary>
        /// Description of action group
        /// </summary>
        public readonly string? Description;
        public readonly Outputs.AgentFunctionSchema? FunctionSchema;
        public readonly Pulumi.AwsNative.Bedrock.AgentActionGroupSignature? ParentActionGroupSignature;
        /// <summary>
        /// Specifies whether to allow deleting action group while it is in use.
        /// </summary>
        public readonly bool? SkipResourceInUseCheckOnDelete;

        [OutputConstructor]
        private AgentActionGroup(
            Union<Outputs.AgentActionGroupExecutor0Properties, Outputs.AgentActionGroupExecutor1Properties>? actionGroupExecutor,

            string actionGroupName,

            Pulumi.AwsNative.Bedrock.AgentActionGroupState? actionGroupState,

            Union<Outputs.AgentApiSchema0Properties, Outputs.AgentApiSchema1Properties>? apiSchema,

            string? description,

            Outputs.AgentFunctionSchema? functionSchema,

            Pulumi.AwsNative.Bedrock.AgentActionGroupSignature? parentActionGroupSignature,

            bool? skipResourceInUseCheckOnDelete)
        {
            ActionGroupExecutor = actionGroupExecutor;
            ActionGroupName = actionGroupName;
            ActionGroupState = actionGroupState;
            ApiSchema = apiSchema;
            Description = description;
            FunctionSchema = functionSchema;
            ParentActionGroupSignature = parentActionGroupSignature;
            SkipResourceInUseCheckOnDelete = skipResourceInUseCheckOnDelete;
        }
    }
}