// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler.Inputs
{

    /// <summary>
    /// The schedule target.
    /// </summary>
    public sealed class ScheduleTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the target.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule. If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.
        /// </summary>
        [Input("deadLetterConfig")]
        public Input<Inputs.ScheduleDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        /// <summary>
        /// The templated target type for the Amazon ECS [`RunTask`](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) API operation.
        /// </summary>
        [Input("ecsParameters")]
        public Input<Inputs.ScheduleEcsParametersArgs>? EcsParameters { get; set; }

        /// <summary>
        /// The templated target type for the EventBridge [`PutEvents`](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEvents.html) API operation.
        /// </summary>
        [Input("eventBridgeParameters")]
        public Input<Inputs.ScheduleEventBridgeParametersArgs>? EventBridgeParameters { get; set; }

        /// <summary>
        /// The text, or well-formed JSON, passed to the target. If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target, the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, EventBridge Scheduler delivers a default notification to the target.
        /// </summary>
        [Input("input")]
        public Input<string>? Input { get; set; }

        /// <summary>
        /// The templated target type for the Amazon Kinesis [`PutRecord`](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html) API operation.
        /// </summary>
        [Input("kinesisParameters")]
        public Input<Inputs.ScheduleKinesisParametersArgs>? KinesisParameters { get; set; }

        /// <summary>
        /// A `RetryPolicy` object that includes information about the retry policy settings, including the maximum age of an event, and the maximum number of times EventBridge Scheduler will try to deliver the event to a target.
        /// </summary>
        [Input("retryPolicy")]
        public Input<Inputs.ScheduleRetryPolicyArgs>? RetryPolicy { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the schedule is triggered.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The templated target type for the Amazon SageMaker [`StartPipelineExecution`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StartPipelineExecution.html) API operation.
        /// </summary>
        [Input("sageMakerPipelineParameters")]
        public Input<Inputs.ScheduleSageMakerPipelineParametersArgs>? SageMakerPipelineParameters { get; set; }

        /// <summary>
        /// The templated target type for the Amazon SQS [`SendMessage`](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessage.html) API operation. Contains the message group ID to use when the target is a FIFO queue. If you specify an Amazon SQS FIFO queue as a target, the queue must have content-based deduplication enabled. For more information, see [Using the Amazon SQS message deduplication ID](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html) in the *Amazon SQS Developer Guide* .
        /// </summary>
        [Input("sqsParameters")]
        public Input<Inputs.ScheduleSqsParametersArgs>? SqsParameters { get; set; }

        public ScheduleTargetArgs()
        {
        }
        public static new ScheduleTargetArgs Empty => new ScheduleTargetArgs();
    }
}
