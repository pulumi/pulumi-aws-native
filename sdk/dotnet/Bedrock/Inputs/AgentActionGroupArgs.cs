// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Contains the information of an Agent Action Group
    /// </summary>
    public sealed class AgentActionGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionGroupExecutor")]
        public InputUnion<Inputs.AgentActionGroupExecutor0PropertiesArgs, Inputs.AgentActionGroupExecutor1PropertiesArgs>? ActionGroupExecutor { get; set; }

        /// <summary>
        /// Name of the action group
        /// </summary>
        [Input("actionGroupName", required: true)]
        public Input<string> ActionGroupName { get; set; } = null!;

        [Input("actionGroupState")]
        public Input<Pulumi.AwsNative.Bedrock.AgentActionGroupState>? ActionGroupState { get; set; }

        [Input("apiSchema")]
        public InputUnion<Inputs.AgentApiSchema0PropertiesArgs, Inputs.AgentApiSchema1PropertiesArgs>? ApiSchema { get; set; }

        /// <summary>
        /// Description of action group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("functionSchema")]
        public Input<Inputs.AgentFunctionSchemaArgs>? FunctionSchema { get; set; }

        [Input("parentActionGroupSignature")]
        public Input<Pulumi.AwsNative.Bedrock.AgentActionGroupSignature>? ParentActionGroupSignature { get; set; }

        /// <summary>
        /// Specifies whether to allow deleting action group while it is in use.
        /// </summary>
        [Input("skipResourceInUseCheckOnDelete")]
        public Input<bool>? SkipResourceInUseCheckOnDelete { get; set; }

        public AgentActionGroupArgs()
        {
        }
        public static new AgentActionGroupArgs Empty => new AgentActionGroupArgs();
    }
}