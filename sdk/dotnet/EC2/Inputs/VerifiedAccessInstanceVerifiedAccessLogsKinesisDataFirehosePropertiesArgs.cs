// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// Sends Verified Access logs to Kinesis.
    /// </summary>
    public sealed class VerifiedAccessInstanceVerifiedAccessLogsKinesisDataFirehosePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the delivery stream.
        /// </summary>
        [Input("deliveryStream")]
        public Input<string>? DeliveryStream { get; set; }

        /// <summary>
        /// Indicates whether logging is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public VerifiedAccessInstanceVerifiedAccessLogsKinesisDataFirehosePropertiesArgs()
        {
        }
        public static new VerifiedAccessInstanceVerifiedAccessLogsKinesisDataFirehosePropertiesArgs Empty => new VerifiedAccessInstanceVerifiedAccessLogsKinesisDataFirehosePropertiesArgs();
    }
}