// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    /// <summary>
    /// An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
    /// </summary>
    public sealed class ConfigurationSetEventDestinationKinesisFirehoseDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.
        /// </summary>
        [Input("deliveryStreamArn", required: true)]
        public Input<string> DeliveryStreamArn { get; set; } = null!;

        /// <summary>
        /// The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.
        /// </summary>
        [Input("iamRoleArn", required: true)]
        public Input<string> IamRoleArn { get; set; } = null!;

        public ConfigurationSetEventDestinationKinesisFirehoseDestinationArgs()
        {
        }
        public static new ConfigurationSetEventDestinationKinesisFirehoseDestinationArgs Empty => new ConfigurationSetEventDestinationKinesisFirehoseDestinationArgs();
    }
}