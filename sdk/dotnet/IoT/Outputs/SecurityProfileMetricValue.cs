// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// The value to be compared with the metric.
    /// </summary>
    [OutputType]
    public sealed class SecurityProfileMetricValue
    {
        /// <summary>
        /// If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.
        /// </summary>
        public readonly ImmutableArray<string> Cidrs;
        /// <summary>
        /// If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.
        /// </summary>
        public readonly string? Count;
        /// <summary>
        /// The numeral value of a metric.
        /// </summary>
        public readonly double? Number;
        /// <summary>
        /// The numeral values of a metric.
        /// </summary>
        public readonly ImmutableArray<double> Numbers;
        /// <summary>
        /// If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.
        /// </summary>
        public readonly ImmutableArray<int> Ports;
        /// <summary>
        /// The string values of a metric.
        /// </summary>
        public readonly ImmutableArray<string> Strings;

        [OutputConstructor]
        private SecurityProfileMetricValue(
            ImmutableArray<string> cidrs,

            string? count,

            double? number,

            ImmutableArray<double> numbers,

            ImmutableArray<int> ports,

            ImmutableArray<string> strings)
        {
            Cidrs = cidrs;
            Count = count;
            Number = number;
            Numbers = numbers;
            Ports = ports;
            Strings = strings;
        }
    }
}
