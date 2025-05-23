// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamElasticsearchDestinationConfiguration
    {
        /// <summary>
        /// Configures how Kinesis Data Firehose buffers incoming data while delivering it to the Amazon ES domain.
        /// </summary>
        public readonly Outputs.DeliveryStreamElasticsearchBufferingHints? BufferingHints;
        /// <summary>
        /// The Amazon CloudWatch Logs logging options for the delivery stream.
        /// </summary>
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        /// <summary>
        /// The endpoint to use when communicating with the cluster. Specify either this `ClusterEndpoint` or the `DomainARN` field.
        /// </summary>
        public readonly string? ClusterEndpoint;
        /// <summary>
        /// Indicates the method for setting up document ID. The supported methods are Firehose generated document ID and OpenSearch Service generated document ID.
        /// </summary>
        public readonly Outputs.DeliveryStreamDocumentIdOptions? DocumentIdOptions;
        /// <summary>
        /// The ARN of the Amazon ES domain. The IAM role must have permissions for `DescribeElasticsearchDomain` , `DescribeElasticsearchDomains` , and `DescribeElasticsearchDomainConfig` after assuming the role specified in *RoleARN* .
        /// 
        /// Specify either `ClusterEndpoint` or `DomainARN` .
        /// </summary>
        public readonly string? DomainArn;
        /// <summary>
        /// The name of the Elasticsearch index to which Kinesis Data Firehose adds data for indexing.
        /// </summary>
        public readonly string IndexName;
        /// <summary>
        /// The frequency of Elasticsearch index rotation. If you enable index rotation, Kinesis Data Firehose appends a portion of the UTC arrival timestamp to the specified index name, and rotates the appended timestamp accordingly. For more information, see [Index Rotation for the Amazon ES Destination](https://docs.aws.amazon.com/firehose/latest/dev/basic-deliver.html#es-index-rotation) in the *Amazon Kinesis Data Firehose Developer Guide* .
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod? IndexRotationPeriod;
        /// <summary>
        /// The data processing configuration for the Kinesis Data Firehose delivery stream.
        /// </summary>
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        /// <summary>
        /// The retry behavior when Kinesis Data Firehose is unable to deliver data to Amazon ES.
        /// </summary>
        public readonly Outputs.DeliveryStreamElasticsearchRetryOptions? RetryOptions;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose for calling the Amazon ES Configuration API and for indexing documents. For more information, see [Controlling Access with Amazon Kinesis Data Firehose](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html) .
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The condition under which Kinesis Data Firehose delivers data to Amazon Simple Storage Service (Amazon S3). You can send Amazon S3 all documents (all data) or only the documents that Kinesis Data Firehose could not deliver to the Amazon ES destination. For more information and valid values, see the `S3BackupMode` content for the [ElasticsearchDestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ElasticsearchDestinationConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode? S3BackupMode;
        /// <summary>
        /// The S3 bucket where Kinesis Data Firehose backs up incoming data.
        /// </summary>
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;
        /// <summary>
        /// The Elasticsearch type name that Amazon ES adds to documents when indexing data.
        /// </summary>
        public readonly string? TypeName;
        /// <summary>
        /// The details of the VPC of the Amazon ES destination.
        /// </summary>
        public readonly Outputs.DeliveryStreamVpcConfiguration? VpcConfiguration;

        [OutputConstructor]
        private DeliveryStreamElasticsearchDestinationConfiguration(
            Outputs.DeliveryStreamElasticsearchBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            string? clusterEndpoint,

            Outputs.DeliveryStreamDocumentIdOptions? documentIdOptions,

            string? domainArn,

            string indexName,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod? indexRotationPeriod,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamElasticsearchRetryOptions? retryOptions,

            string roleArn,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration,

            string? typeName,

            Outputs.DeliveryStreamVpcConfiguration? vpcConfiguration)
        {
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            ClusterEndpoint = clusterEndpoint;
            DocumentIdOptions = documentIdOptions;
            DomainArn = domainArn;
            IndexName = indexName;
            IndexRotationPeriod = indexRotationPeriod;
            ProcessingConfiguration = processingConfiguration;
            RetryOptions = retryOptions;
            RoleArn = roleArn;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
            TypeName = typeName;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
