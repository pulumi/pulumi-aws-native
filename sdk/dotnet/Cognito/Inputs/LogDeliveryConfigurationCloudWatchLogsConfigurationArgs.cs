// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class LogDeliveryConfigurationCloudWatchLogsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("logGroupArn")]
        public Input<string>? LogGroupArn { get; set; }

        public LogDeliveryConfigurationCloudWatchLogsConfigurationArgs()
        {
        }
        public static new LogDeliveryConfigurationCloudWatchLogsConfigurationArgs Empty => new LogDeliveryConfigurationCloudWatchLogsConfigurationArgs();
    }
}