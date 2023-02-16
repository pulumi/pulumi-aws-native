// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// item in metric groups
    /// </summary>
    public sealed class ModelCardMetricGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricData", required: true)]
        private InputList<object>? _metricData;
        public InputList<object> MetricData
        {
            get => _metricData ?? (_metricData = new InputList<object>());
            set => _metricData = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ModelCardMetricGroupArgs()
        {
        }
        public static new ModelCardMetricGroupArgs Empty => new ModelCardMetricGroupArgs();
    }
}