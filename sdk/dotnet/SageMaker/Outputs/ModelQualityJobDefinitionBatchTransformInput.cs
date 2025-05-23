// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The batch transform input for a monitoring job.
    /// </summary>
    [OutputType]
    public sealed class ModelQualityJobDefinitionBatchTransformInput
    {
        /// <summary>
        /// A URI that identifies the Amazon S3 storage location where Batch Transform Job captures data.
        /// </summary>
        public readonly string DataCapturedDestinationS3Uri;
        /// <summary>
        /// The dataset format for your batch transform job.
        /// </summary>
        public readonly Outputs.ModelQualityJobDefinitionDatasetFormat DatasetFormat;
        /// <summary>
        /// Monitoring end time offset, e.g. PT0H
        /// </summary>
        public readonly string? EndTimeOffset;
        /// <summary>
        /// Index or JSONpath to locate predicted label(s)
        /// </summary>
        public readonly string? InferenceAttribute;
        /// <summary>
        /// Path to the filesystem where the endpoint data is available to the container.
        /// </summary>
        public readonly string LocalPath;
        /// <summary>
        /// Index or JSONpath to locate probabilities
        /// </summary>
        public readonly string? ProbabilityAttribute;
        /// <summary>
        /// The threshold for the class probability to be evaluated as a positive result.
        /// </summary>
        public readonly double? ProbabilityThresholdAttribute;
        /// <summary>
        /// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType? S3DataDistributionType;
        /// <summary>
        /// Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelQualityJobDefinitionBatchTransformInputS3InputMode? S3InputMode;
        /// <summary>
        /// Monitoring start time offset, e.g. -PT1H
        /// </summary>
        public readonly string? StartTimeOffset;

        [OutputConstructor]
        private ModelQualityJobDefinitionBatchTransformInput(
            string dataCapturedDestinationS3Uri,

            Outputs.ModelQualityJobDefinitionDatasetFormat datasetFormat,

            string? endTimeOffset,

            string? inferenceAttribute,

            string localPath,

            string? probabilityAttribute,

            double? probabilityThresholdAttribute,

            Pulumi.AwsNative.SageMaker.ModelQualityJobDefinitionBatchTransformInputS3DataDistributionType? s3DataDistributionType,

            Pulumi.AwsNative.SageMaker.ModelQualityJobDefinitionBatchTransformInputS3InputMode? s3InputMode,

            string? startTimeOffset)
        {
            DataCapturedDestinationS3Uri = dataCapturedDestinationS3Uri;
            DatasetFormat = datasetFormat;
            EndTimeOffset = endTimeOffset;
            InferenceAttribute = inferenceAttribute;
            LocalPath = localPath;
            ProbabilityAttribute = probabilityAttribute;
            ProbabilityThresholdAttribute = probabilityThresholdAttribute;
            S3DataDistributionType = s3DataDistributionType;
            S3InputMode = s3InputMode;
            StartTimeOffset = startTimeOffset;
        }
    }
}
