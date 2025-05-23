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
    /// The endpoint for a monitoring job.
    /// </summary>
    [OutputType]
    public sealed class DataQualityJobDefinitionEndpointInput
    {
        /// <summary>
        /// An endpoint in customer's account which has enabled `DataCaptureConfig` enabled.
        /// </summary>
        public readonly string EndpointName;
        /// <summary>
        /// Indexes or names of the features to be excluded from analysis
        /// </summary>
        public readonly string? ExcludeFeaturesAttribute;
        /// <summary>
        /// Path to the filesystem where the endpoint data is available to the container.
        /// </summary>
        public readonly string LocalPath;
        /// <summary>
        /// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.DataQualityJobDefinitionEndpointInputS3DataDistributionType? S3DataDistributionType;
        /// <summary>
        /// Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.DataQualityJobDefinitionEndpointInputS3InputMode? S3InputMode;

        [OutputConstructor]
        private DataQualityJobDefinitionEndpointInput(
            string endpointName,

            string? excludeFeaturesAttribute,

            string localPath,

            Pulumi.AwsNative.SageMaker.DataQualityJobDefinitionEndpointInputS3DataDistributionType? s3DataDistributionType,

            Pulumi.AwsNative.SageMaker.DataQualityJobDefinitionEndpointInputS3InputMode? s3InputMode)
        {
            EndpointName = endpointName;
            ExcludeFeaturesAttribute = excludeFeaturesAttribute;
            LocalPath = localPath;
            S3DataDistributionType = s3DataDistributionType;
            S3InputMode = s3InputMode;
        }
    }
}
