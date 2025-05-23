// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch.Outputs
{

    /// <summary>
    /// This structure defines the metrics that will be streamed.
    /// </summary>
    [OutputType]
    public sealed class MetricStreamFilter
    {
        /// <summary>
        /// Only metrics with MetricNames matching these values will be streamed. Must be set together with Namespace.
        /// </summary>
        public readonly ImmutableArray<string> MetricNames;
        /// <summary>
        /// Only metrics with Namespace matching this value will be streamed.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private MetricStreamFilter(
            ImmutableArray<string> metricNames,

            string @namespace)
        {
            MetricNames = metricNames;
            Namespace = @namespace;
        }
    }
}
