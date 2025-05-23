// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Outputs
{

    /// <summary>
    /// The metric to optimize during HPO.
    /// </summary>
    [OutputType]
    public sealed class SolutionConfigHpoConfigPropertiesHpoObjectiveProperties
    {
        /// <summary>
        /// The name of the metric
        /// </summary>
        public readonly string? MetricName;
        /// <summary>
        /// A regular expression for finding the metric in the training job logs.
        /// </summary>
        public readonly string? MetricRegex;
        /// <summary>
        /// The type of the metric. Valid values are Maximize and Minimize.
        /// </summary>
        public readonly Pulumi.AwsNative.Personalize.SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType? Type;

        [OutputConstructor]
        private SolutionConfigHpoConfigPropertiesHpoObjectiveProperties(
            string? metricName,

            string? metricRegex,

            Pulumi.AwsNative.Personalize.SolutionConfigHpoConfigPropertiesHpoObjectivePropertiesType? type)
        {
            MetricName = metricName;
            MetricRegex = metricRegex;
            Type = type;
        }
    }
}
