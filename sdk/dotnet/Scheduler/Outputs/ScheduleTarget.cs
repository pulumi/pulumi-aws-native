// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Outputs
{

    /// <summary>
    /// The schedule target.
    /// </summary>
    [OutputType]
    public sealed class ScheduleTarget
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the target.
        /// </summary>
        public readonly string Arn;
        public readonly Outputs.ScheduleDeadLetterConfig? DeadLetterConfig;
        public readonly Outputs.ScheduleEcsParameters? EcsParameters;
        public readonly Outputs.ScheduleEventBridgeParameters? EventBridgeParameters;
        /// <summary>
        /// The text, or well-formed JSON, passed to the target. If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target, the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, EventBridge Scheduler delivers a default notification to the target.
        /// </summary>
        public readonly string? Input;
        public readonly Outputs.ScheduleKinesisParameters? KinesisParameters;
        public readonly Outputs.ScheduleRetryPolicy? RetryPolicy;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the schedule is triggered.
        /// </summary>
        public readonly string RoleArn;
        public readonly Outputs.ScheduleSageMakerPipelineParameters? SageMakerPipelineParameters;
        public readonly Outputs.ScheduleSqsParameters? SqsParameters;

        [OutputConstructor]
        private ScheduleTarget(
            string arn,

            Outputs.ScheduleDeadLetterConfig? deadLetterConfig,

            Outputs.ScheduleEcsParameters? ecsParameters,

            Outputs.ScheduleEventBridgeParameters? eventBridgeParameters,

            string? input,

            Outputs.ScheduleKinesisParameters? kinesisParameters,

            Outputs.ScheduleRetryPolicy? retryPolicy,

            string roleArn,

            Outputs.ScheduleSageMakerPipelineParameters? sageMakerPipelineParameters,

            Outputs.ScheduleSqsParameters? sqsParameters)
        {
            Arn = arn;
            DeadLetterConfig = deadLetterConfig;
            EcsParameters = ecsParameters;
            EventBridgeParameters = eventBridgeParameters;
            Input = input;
            KinesisParameters = kinesisParameters;
            RetryPolicy = retryPolicy;
            RoleArn = roleArn;
            SageMakerPipelineParameters = sageMakerPipelineParameters;
            SqsParameters = sqsParameters;
        }
    }
}