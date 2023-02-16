// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    /// <summary>
    /// Configuration options for an SNS alert action.
    /// </summary>
    public sealed class AlertSNSConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// ARN of an SNS topic to send alert notifications to.
        /// </summary>
        [Input("snsTopicArn", required: true)]
        public Input<string> SnsTopicArn { get; set; } = null!;

        public AlertSNSConfigurationArgs()
        {
        }
        public static new AlertSNSConfigurationArgs Empty => new AlertSNSConfigurationArgs();
    }
}