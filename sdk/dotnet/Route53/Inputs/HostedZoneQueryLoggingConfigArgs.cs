// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Inputs
{

    /// <summary>
    /// A complex type that contains information about a configuration for DNS query logging.
    /// </summary>
    public sealed class HostedZoneQueryLoggingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53 is publishing logs to.
        /// </summary>
        [Input("cloudWatchLogsLogGroupArn", required: true)]
        public Input<string> CloudWatchLogsLogGroupArn { get; set; } = null!;

        public HostedZoneQueryLoggingConfigArgs()
        {
        }
        public static new HostedZoneQueryLoggingConfigArgs Empty => new HostedZoneQueryLoggingConfigArgs();
    }
}