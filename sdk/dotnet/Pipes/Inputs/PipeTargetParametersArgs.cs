// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeTargetParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameters for using an AWS Batch job as a target.
        /// </summary>
        [Input("batchJobParameters")]
        public Input<Inputs.PipeTargetBatchJobParametersArgs>? BatchJobParameters { get; set; }

        /// <summary>
        /// The parameters for using an CloudWatch Logs log stream as a target.
        /// </summary>
        [Input("cloudWatchLogsParameters")]
        public Input<Inputs.PipeTargetCloudWatchLogsParametersArgs>? CloudWatchLogsParameters { get; set; }

        /// <summary>
        /// The parameters for using an Amazon ECS task as a target.
        /// </summary>
        [Input("ecsTaskParameters")]
        public Input<Inputs.PipeTargetEcsTaskParametersArgs>? EcsTaskParameters { get; set; }

        /// <summary>
        /// The parameters for using an EventBridge event bus as a target.
        /// </summary>
        [Input("eventBridgeEventBusParameters")]
        public Input<Inputs.PipeTargetEventBridgeEventBusParametersArgs>? EventBridgeEventBusParameters { get; set; }

        /// <summary>
        /// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
        /// </summary>
        [Input("httpParameters")]
        public Input<Inputs.PipeTargetHttpParametersArgs>? HttpParameters { get; set; }

        /// <summary>
        /// Valid JSON text passed to the target. In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
        /// 
        /// To remove an input template, specify an empty string.
        /// </summary>
        [Input("inputTemplate")]
        public Input<string>? InputTemplate { get; set; }

        /// <summary>
        /// The parameters for using a Kinesis stream as a target.
        /// </summary>
        [Input("kinesisStreamParameters")]
        public Input<Inputs.PipeTargetKinesisStreamParametersArgs>? KinesisStreamParameters { get; set; }

        /// <summary>
        /// The parameters for using a Lambda function as a target.
        /// </summary>
        [Input("lambdaFunctionParameters")]
        public Input<Inputs.PipeTargetLambdaFunctionParametersArgs>? LambdaFunctionParameters { get; set; }

        /// <summary>
        /// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement.
        /// </summary>
        [Input("redshiftDataParameters")]
        public Input<Inputs.PipeTargetRedshiftDataParametersArgs>? RedshiftDataParameters { get; set; }

        /// <summary>
        /// The parameters for using a SageMaker AI pipeline as a target.
        /// </summary>
        [Input("sageMakerPipelineParameters")]
        public Input<Inputs.PipeTargetSageMakerPipelineParametersArgs>? SageMakerPipelineParameters { get; set; }

        /// <summary>
        /// The parameters for using a Amazon SQS stream as a target.
        /// </summary>
        [Input("sqsQueueParameters")]
        public Input<Inputs.PipeTargetSqsQueueParametersArgs>? SqsQueueParameters { get; set; }

        /// <summary>
        /// The parameters for using a Step Functions state machine as a target.
        /// </summary>
        [Input("stepFunctionStateMachineParameters")]
        public Input<Inputs.PipeTargetStateMachineParametersArgs>? StepFunctionStateMachineParameters { get; set; }

        /// <summary>
        /// The parameters for using a Timestream for LiveAnalytics table as a target.
        /// </summary>
        [Input("timestreamParameters")]
        public Input<Inputs.PipeTargetTimestreamParametersArgs>? TimestreamParameters { get; set; }

        public PipeTargetParametersArgs()
        {
        }
        public static new PipeTargetParametersArgs Empty => new PipeTargetParametersArgs();
    }
}
