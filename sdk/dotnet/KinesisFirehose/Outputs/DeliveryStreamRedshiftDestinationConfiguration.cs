// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamRedshiftDestinationConfiguration
    {
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        public readonly string ClusterJDBCURL;
        public readonly Outputs.DeliveryStreamCopyCommand CopyCommand;
        public readonly string Password;
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        public readonly Outputs.DeliveryStreamRedshiftRetryOptions? RetryOptions;
        public readonly string RoleARN;
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration? S3BackupConfiguration;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamRedshiftDestinationConfigurationS3BackupMode? S3BackupMode;
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;
        public readonly string Username;

        [OutputConstructor]
        private DeliveryStreamRedshiftDestinationConfiguration(
            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            string clusterJDBCURL,

            Outputs.DeliveryStreamCopyCommand copyCommand,

            string password,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamRedshiftRetryOptions? retryOptions,

            string roleARN,

            Outputs.DeliveryStreamS3DestinationConfiguration? s3BackupConfiguration,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamRedshiftDestinationConfigurationS3BackupMode? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration,

            string username)
        {
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            ClusterJDBCURL = clusterJDBCURL;
            CopyCommand = copyCommand;
            Password = password;
            ProcessingConfiguration = processingConfiguration;
            RetryOptions = retryOptions;
            RoleARN = roleARN;
            S3BackupConfiguration = s3BackupConfiguration;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
            Username = username;
        }
    }
}