// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    /// <summary>
    /// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
    /// </summary>
    public sealed class MonitoringSubscriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A subscription configuration for additional CloudWatch metrics.
        /// </summary>
        [Input("realtimeMetricsSubscriptionConfig")]
        public Input<Inputs.MonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs>? RealtimeMetricsSubscriptionConfig { get; set; }

        public MonitoringSubscriptionArgs()
        {
        }
        public static new MonitoringSubscriptionArgs Empty => new MonitoringSubscriptionArgs();
    }
}
