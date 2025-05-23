// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// You can enable a restart policy for each container defined in your task definition, to overcome transient failures faster and maintain task availability. When you enable a restart policy for a container, Amazon ECS can restart the container if it exits, without needing to replace the task. For more information, see [Restart individual containers in Amazon ECS tasks with container restart policies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-restart-policy.html) in the *Amazon Elastic Container Service Developer Guide*.
    /// </summary>
    public sealed class TaskDefinitionRestartPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether a restart policy is enabled for the container.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("ignoredExitCodes")]
        private InputList<int>? _ignoredExitCodes;

        /// <summary>
        /// A list of exit codes that Amazon ECS will ignore and not attempt a restart on. You can specify a maximum of 50 container exit codes. By default, Amazon ECS does not ignore any exit codes.
        /// </summary>
        public InputList<int> IgnoredExitCodes
        {
            get => _ignoredExitCodes ?? (_ignoredExitCodes = new InputList<int>());
            set => _ignoredExitCodes = value;
        }

        /// <summary>
        /// A period of time (in seconds) that the container must run for before a restart can be attempted. A container can be restarted only once every ``restartAttemptPeriod`` seconds. If a container isn't able to run for this time period and exits early, it will not be restarted. You can set a minimum ``restartAttemptPeriod`` of 60 seconds and a maximum ``restartAttemptPeriod`` of 1800 seconds. By default, a container must run for 300 seconds before it can be restarted.
        /// </summary>
        [Input("restartAttemptPeriod")]
        public Input<int>? RestartAttemptPeriod { get; set; }

        public TaskDefinitionRestartPolicyArgs()
        {
        }
        public static new TaskDefinitionRestartPolicyArgs Empty => new TaskDefinitionRestartPolicyArgs();
    }
}
