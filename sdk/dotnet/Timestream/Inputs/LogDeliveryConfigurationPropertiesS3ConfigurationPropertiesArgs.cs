// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Inputs
{

    /// <summary>
    /// S3 configuration for sending logs to customer account from the InfluxDB instance.
    /// </summary>
    public sealed class LogDeliveryConfigurationPropertiesS3ConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bucket name for logs to be sent from the InfluxDB instance
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Specifies whether logging to customer specified bucket is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public LogDeliveryConfigurationPropertiesS3ConfigurationPropertiesArgs()
        {
        }
        public static new LogDeliveryConfigurationPropertiesS3ConfigurationPropertiesArgs Empty => new LogDeliveryConfigurationPropertiesS3ConfigurationPropertiesArgs();
    }
}
