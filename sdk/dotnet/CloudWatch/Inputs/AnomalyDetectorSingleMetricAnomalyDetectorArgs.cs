// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Inputs
{

    public sealed class AnomalyDetectorSingleMetricAnomalyDetectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.AnomalyDetectorDimensionArgs>? _dimensions;
        public InputList<Inputs.AnomalyDetectorDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.AnomalyDetectorDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("stat")]
        public Input<string>? Stat { get; set; }

        public AnomalyDetectorSingleMetricAnomalyDetectorArgs()
        {
        }
        public static new AnomalyDetectorSingleMetricAnomalyDetectorArgs Empty => new AnomalyDetectorSingleMetricAnomalyDetectorArgs();
    }
}