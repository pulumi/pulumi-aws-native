// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs.Inputs
{

    /// <summary>
    /// the key-value pairs that further define a metric.
    /// </summary>
    public sealed class MetricFilterDimensionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the dimension. Maximum length of 255.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value of the dimension. Maximum length of 255.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public MetricFilterDimensionArgs()
        {
        }
        public static new MetricFilterDimensionArgs Empty => new MetricFilterDimensionArgs();
    }
}