// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationSignals.Outputs
{

    /// <summary>
    /// This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.
    /// </summary>
    [OutputType]
    public sealed class ServiceLevelObjectiveMetric
    {
        /// <summary>
        /// An array of one or more dimensions to use to define the metric that you want to use.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceLevelObjectiveDimension> Dimensions;
        /// <summary>
        /// The name of the metric to use.
        /// </summary>
        public readonly string? MetricName;
        /// <summary>
        /// The namespace of the metric.
        /// </summary>
        public readonly string? Namespace;

        [OutputConstructor]
        private ServiceLevelObjectiveMetric(
            ImmutableArray<Outputs.ServiceLevelObjectiveDimension> dimensions,

            string? metricName,

            string? @namespace)
        {
            Dimensions = dimensions;
            MetricName = metricName;
            Namespace = @namespace;
        }
    }
}
