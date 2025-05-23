// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Inputs
{

    /// <summary>
    /// This structure defines the metrics that will be streamed.
    /// </summary>
    public sealed class MetricStreamFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricNames")]
        private InputList<string>? _metricNames;

        /// <summary>
        /// Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.
        /// </summary>
        public InputList<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new InputList<string>());
            set => _metricNames = value;
        }

        /// <summary>
        /// Only metrics with Namespace matching this value will be streamed.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public MetricStreamFilterArgs()
        {
        }
        public static new MetricStreamFilterArgs Empty => new MetricStreamFilterArgs();
    }
}
