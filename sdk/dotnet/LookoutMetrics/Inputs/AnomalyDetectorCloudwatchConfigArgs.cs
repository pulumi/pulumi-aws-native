// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorCloudwatchConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public AnomalyDetectorCloudwatchConfigArgs()
        {
        }
        public static new AnomalyDetectorCloudwatchConfigArgs Empty => new AnomalyDetectorCloudwatchConfigArgs();
    }
}