// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeSourceDynamoDBStreamParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("batchSize")]
        public Input<int>? BatchSize { get; set; }

        [Input("deadLetterConfig")]
        public Input<Inputs.PipeDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        [Input("maximumBatchingWindowInSeconds")]
        public Input<int>? MaximumBatchingWindowInSeconds { get; set; }

        [Input("maximumRecordAgeInSeconds")]
        public Input<int>? MaximumRecordAgeInSeconds { get; set; }

        [Input("maximumRetryAttempts")]
        public Input<int>? MaximumRetryAttempts { get; set; }

        [Input("onPartialBatchItemFailure")]
        public Input<Pulumi.AwsNative.Pipes.PipeOnPartialBatchItemFailureStreams>? OnPartialBatchItemFailure { get; set; }

        [Input("parallelizationFactor")]
        public Input<int>? ParallelizationFactor { get; set; }

        [Input("startingPosition", required: true)]
        public Input<Pulumi.AwsNative.Pipes.PipeDynamoDBStreamStartPosition> StartingPosition { get; set; } = null!;

        public PipeSourceDynamoDBStreamParametersArgs()
        {
        }
        public static new PipeSourceDynamoDBStreamParametersArgs Empty => new PipeSourceDynamoDBStreamParametersArgs();
    }
}