// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Inputs
{

    public sealed class AnomalyDetectorConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedTimeRanges")]
        private InputList<Inputs.AnomalyDetectorRangeArgs>? _excludedTimeRanges;
        public InputList<Inputs.AnomalyDetectorRangeArgs> ExcludedTimeRanges
        {
            get => _excludedTimeRanges ?? (_excludedTimeRanges = new InputList<Inputs.AnomalyDetectorRangeArgs>());
            set => _excludedTimeRanges = value;
        }

        [Input("metricTimeZone")]
        public Input<string>? MetricTimeZone { get; set; }

        public AnomalyDetectorConfigurationArgs()
        {
        }
        public static new AnomalyDetectorConfigurationArgs Empty => new AnomalyDetectorConfigurationArgs();
    }
}