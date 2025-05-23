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
    public sealed class DeliveryStreamS3DestinationConfiguration
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon S3 bucket to send data to.
        /// </summary>
        public readonly string BucketArn;
        /// <summary>
        /// Configures how Kinesis Data Firehose buffers incoming data while delivering it to the Amazon S3 bucket.
        /// </summary>
        public readonly Outputs.DeliveryStreamBufferingHints? BufferingHints;
        /// <summary>
        /// The CloudWatch logging options for your Firehose stream.
        /// </summary>
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        /// <summary>
        /// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket. For valid values, see the `CompressionFormat` content for the [S3DestinationConfiguration](https://docs.aws.amazon.com/firehose/latest/APIReference/API_S3DestinationConfiguration.html) data type in the *Amazon Kinesis Data Firehose API Reference* .
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamS3DestinationConfigurationCompressionFormat? CompressionFormat;
        /// <summary>
        /// Configures Amazon Simple Storage Service (Amazon S3) server-side encryption. Kinesis Data Firehose uses AWS Key Management Service ( AWS KMS) to encrypt the data that it delivers to your Amazon S3 bucket.
        /// </summary>
        public readonly Outputs.DeliveryStreamEncryptionConfiguration? EncryptionConfiguration;
        /// <summary>
        /// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3. This prefix appears immediately following the bucket name. For information about how to specify this prefix, see [Custom Prefixes for Amazon S3 Objects](https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html) .
        /// </summary>
        public readonly string? ErrorOutputPrefix;
        /// <summary>
        /// A prefix that Kinesis Data Firehose adds to the files that it delivers to the Amazon S3 bucket. The prefix helps you identify the files that Kinesis Data Firehose delivered.
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// The ARN of an AWS Identity and Access Management (IAM) role that grants Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS (if you enable data encryption). For more information, see [Grant Kinesis Data Firehose Access to an Amazon S3 Destination](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3) in the *Amazon Kinesis Data Firehose Developer Guide* .
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private DeliveryStreamS3DestinationConfiguration(
            string bucketArn,

            Outputs.DeliveryStreamBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamS3DestinationConfigurationCompressionFormat? compressionFormat,

            Outputs.DeliveryStreamEncryptionConfiguration? encryptionConfiguration,

            string? errorOutputPrefix,

            string? prefix,

            string roleArn)
        {
            BucketArn = bucketArn;
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            CompressionFormat = compressionFormat;
            EncryptionConfiguration = encryptionConfiguration;
            ErrorOutputPrefix = errorOutputPrefix;
            Prefix = prefix;
            RoleArn = roleArn;
        }
    }
}
