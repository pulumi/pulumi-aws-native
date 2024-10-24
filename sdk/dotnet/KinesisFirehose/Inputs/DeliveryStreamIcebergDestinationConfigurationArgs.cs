// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamIcebergDestinationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bufferingHints")]
        public Input<Inputs.DeliveryStreamBufferingHintsArgs>? BufferingHints { get; set; }

        /// <summary>
        /// Configuration describing where the destination Apache Iceberg Tables are persisted.
        /// </summary>
        [Input("catalogConfiguration", required: true)]
        public Input<Inputs.DeliveryStreamCatalogConfigurationArgs> CatalogConfiguration { get; set; } = null!;

        [Input("cloudWatchLoggingOptions")]
        public Input<Inputs.DeliveryStreamCloudWatchLoggingOptionsArgs>? CloudWatchLoggingOptions { get; set; }

        [Input("destinationTableConfigurationList")]
        private InputList<Inputs.DeliveryStreamDestinationTableConfigurationArgs>? _destinationTableConfigurationList;

        /// <summary>
        /// Provides a list of `DestinationTableConfigurations` which Firehose uses to deliver data to Apache Iceberg Tables. Firehose will write data with insert if table specific configuration is not provided here.
        /// </summary>
        public InputList<Inputs.DeliveryStreamDestinationTableConfigurationArgs> DestinationTableConfigurationList
        {
            get => _destinationTableConfigurationList ?? (_destinationTableConfigurationList = new InputList<Inputs.DeliveryStreamDestinationTableConfigurationArgs>());
            set => _destinationTableConfigurationList = value;
        }

        [Input("processingConfiguration")]
        public Input<Inputs.DeliveryStreamProcessingConfigurationArgs>? ProcessingConfiguration { get; set; }

        [Input("retryOptions")]
        public Input<Inputs.DeliveryStreamRetryOptionsArgs>? RetryOptions { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role to be assumed by Firehose for calling Apache Iceberg Tables.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// Describes how Firehose will backup records. Currently,S3 backup only supports `FailedDataOnly` .
        /// </summary>
        [Input("s3BackupMode")]
        public Input<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamIcebergDestinationConfigurations3BackupMode>? S3BackupMode { get; set; }

        [Input("s3Configuration", required: true)]
        public Input<Inputs.DeliveryStreamS3DestinationConfigurationArgs> S3Configuration { get; set; } = null!;

        public DeliveryStreamIcebergDestinationConfigurationArgs()
        {
        }
        public static new DeliveryStreamIcebergDestinationConfigurationArgs Empty => new DeliveryStreamIcebergDestinationConfigurationArgs();
    }
}
