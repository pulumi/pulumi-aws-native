// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class LogDeliveryConfigurationLogConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudWatchLogsConfiguration")]
        public Input<Inputs.LogDeliveryConfigurationCloudWatchLogsConfigurationArgs>? CloudWatchLogsConfiguration { get; set; }

        [Input("eventSource")]
        public Input<string>? EventSource { get; set; }

        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        public LogDeliveryConfigurationLogConfigurationArgs()
        {
        }
        public static new LogDeliveryConfigurationLogConfigurationArgs Empty => new LogDeliveryConfigurationLogConfigurationArgs();
    }
}