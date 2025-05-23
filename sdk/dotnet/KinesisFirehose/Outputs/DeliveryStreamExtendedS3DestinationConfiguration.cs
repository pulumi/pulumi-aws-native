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
    public sealed class DeliveryStreamExtendedS3DestinationConfiguration
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon S3 bucket. For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        public readonly string BucketArn;
        /// <summary>
        /// The buffering option.
        /// </summary>
        public readonly Outputs.DeliveryStreamBufferingHints? BufferingHints;
        /// <summary>
        /// The Amazon CloudWatch logging options for your Firehose stream.
        /// </summary>
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        /// <summary>
        /// The compression format. If no value is specified, the default is `UNCOMPRESSED` .
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat? CompressionFormat;
        /// <summary>
        /// The time zone you prefer. UTC is the default.
        /// </summary>
        public readonly string? CustomTimeZone;
        /// <summary>
        /// The serializer, deserializer, and schema for converting data from the JSON format to the Parquet or ORC format before writing it to Amazon S3.
        /// </summary>
        public readonly Outputs.DeliveryStreamDataFormatConversionConfiguration? DataFormatConversionConfiguration;
        /// <summary>
        /// The configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
        /// </summary>
        public readonly Outputs.DeliveryStreamDynamicPartitioningConfiguration? DynamicPartitioningConfiguration;
        /// <summary>
        /// The encryption configuration for the Kinesis Data Firehose delivery stream. The default value is `NoEncryption` .
        /// </summary>
        public readonly Outputs.DeliveryStreamEncryptionConfiguration? EncryptionConfiguration;
        /// <summary>
        /// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
        /// </summary>
        public readonly string? ErrorOutputPrefix;
        /// <summary>
        /// Specify a file extension. It will override the default file extension
        /// </summary>
        public readonly string? FileExtension;
        /// <summary>
        /// The `YYYY/MM/DD/HH` time format prefix is automatically used for delivered Amazon S3 files. For more information, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// The data processing configuration for the Kinesis Data Firehose delivery stream.
        /// </summary>
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS credentials. For constraints, see [ExtendedS3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_ExtendedS3DestinationConfiguration.html) in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The configuration for backup in Amazon S3.
        /// </summary>
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration? S3BackupConfiguration;
        /// <summary>
        /// The Amazon S3 backup mode. After you create a Firehose stream, you can update it to enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the Firehose stream to disable it.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode? S3BackupMode;

        [OutputConstructor]
        private DeliveryStreamExtendedS3DestinationConfiguration(
            string bucketArn,

            Outputs.DeliveryStreamBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat? compressionFormat,

            string? customTimeZone,

            Outputs.DeliveryStreamDataFormatConversionConfiguration? dataFormatConversionConfiguration,

            Outputs.DeliveryStreamDynamicPartitioningConfiguration? dynamicPartitioningConfiguration,

            Outputs.DeliveryStreamEncryptionConfiguration? encryptionConfiguration,

            string? errorOutputPrefix,

            string? fileExtension,

            string? prefix,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            string roleArn,

            Outputs.DeliveryStreamS3DestinationConfiguration? s3BackupConfiguration,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode? s3BackupMode)
        {
            BucketArn = bucketArn;
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            CompressionFormat = compressionFormat;
            CustomTimeZone = customTimeZone;
            DataFormatConversionConfiguration = dataFormatConversionConfiguration;
            DynamicPartitioningConfiguration = dynamicPartitioningConfiguration;
            EncryptionConfiguration = encryptionConfiguration;
            ErrorOutputPrefix = errorOutputPrefix;
            FileExtension = fileExtension;
            Prefix = prefix;
            ProcessingConfiguration = processingConfiguration;
            RoleArn = roleArn;
            S3BackupConfiguration = s3BackupConfiguration;
            S3BackupMode = s3BackupMode;
        }
    }
}
