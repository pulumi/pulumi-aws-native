// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The configuration options for AWS Verified Access instances.
    /// </summary>
    public sealed class VerifiedAccessInstanceVerifiedAccessLogsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sends Verified Access logs to CloudWatch Logs.
        /// </summary>
        [Input("cloudWatchLogs")]
        public Input<Inputs.VerifiedAccessInstanceVerifiedAccessLogsCloudWatchLogsPropertiesArgs>? CloudWatchLogs { get; set; }

        /// <summary>
        /// Include claims from trust providers in Verified Access logs.
        /// </summary>
        [Input("includeTrustContext")]
        public Input<bool>? IncludeTrustContext { get; set; }

        /// <summary>
        /// Sends Verified Access logs to Kinesis.
        /// </summary>
        [Input("kinesisDataFirehose")]
        public Input<Inputs.VerifiedAccessInstanceVerifiedAccessLogsKinesisDataFirehosePropertiesArgs>? KinesisDataFirehose { get; set; }

        /// <summary>
        /// Select log version for Verified Access logs.
        /// </summary>
        [Input("logVersion")]
        public Input<string>? LogVersion { get; set; }

        /// <summary>
        /// Sends Verified Access logs to Amazon S3.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.VerifiedAccessInstanceVerifiedAccessLogsS3PropertiesArgs>? S3 { get; set; }

        public VerifiedAccessInstanceVerifiedAccessLogsArgs()
        {
        }
        public static new VerifiedAccessInstanceVerifiedAccessLogsArgs Empty => new VerifiedAccessInstanceVerifiedAccessLogsArgs();
    }
}
