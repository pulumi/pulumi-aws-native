// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleStepFunctionsActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID. Step Functions automatically creates a unique name for each state machine execution if one is not provided.
        /// </summary>
        [Input("executionNamePrefix")]
        public Input<string>? ExecutionNamePrefix { get; set; }

        /// <summary>
        /// The ARN of the role that grants IoT permission to start execution of a state machine ("Action":"states:StartExecution").
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The name of the Step Functions state machine whose execution will be started.
        /// </summary>
        [Input("stateMachineName", required: true)]
        public Input<string> StateMachineName { get; set; } = null!;

        public TopicRuleStepFunctionsActionArgs()
        {
        }
        public static new TopicRuleStepFunctionsActionArgs Empty => new TopicRuleStepFunctionsActionArgs();
    }
}
