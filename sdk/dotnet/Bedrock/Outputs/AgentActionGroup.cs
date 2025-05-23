// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lambda function containing the business logic that is carried out upon invoking the action or the custom control method for handling the information elicited from the user.
        /// </summary>
        public readonly Union<Outputs.AgentActionGroupExecutor0Properties, Outputs.AgentActionGroupExecutor1Properties>? ActionGroupExecutor;
        /// <summary>
        /// Name of the action group
        /// </summary>
        public readonly string ActionGroupName;
        /// <summary>
        /// Specifies whether the action group is available for the agent to invoke or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.
        /// </summary>
        public readonly Pulumi.AwsNative.Bedrock.AgentActionGroupState? ActionGroupState;
        /// <summary>
        /// Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema. For more information, see [Action group OpenAPI schemas](https://docs.aws.amazon.com//bedrock/latest/userguide/agents-api-schema.html) .
        /// </summary>
        public readonly Union<Outputs.AgentApiSchema0Properties, Outputs.AgentApiSchema1Properties>? ApiSchema;
        /// <summary>
        /// Description of action group
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.
        /// </summary>
        public readonly Outputs.AgentFunctionSchema? FunctionSchema;
        /// <summary>
        /// If this field is set as `AMAZON.UserInput` , the agent can request the user for additional information when trying to complete a task. The `description` , `apiSchema` , and `actionGroupExecutor` fields must be blank for this action group.
        /// 
        /// During orchestration, if the agent determines that it needs to invoke an API in an action group, but doesn't have enough information to complete the API request, it will invoke this action group instead and return an [Observation](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html) reprompting the user for more information.
        /// </summary>
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
