// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Inputs
{

    public sealed class AnomalyDetectorMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.AnomalyDetectorDimensionArgs>? _dimensions;
        public InputList<Inputs.AnomalyDetectorDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.AnomalyDetectorDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public AnomalyDetectorMetricArgs()
        {
        }
        public static new AnomalyDetectorMetricArgs Empty => new AnomalyDetectorMetricArgs();
    }
}