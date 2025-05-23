// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobColumnStatisticsConfiguration
    {
        /// <summary>
        /// List of column selectors. Selectors can be used to select columns from the dataset. When selectors are undefined, configuration will be applied to all supported columns.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobColumnSelector> Selectors;
        /// <summary>
        /// Configuration for evaluations. Statistics can be used to select evaluations and override parameters of evaluations.
        /// </summary>
        public readonly Outputs.JobStatisticsConfiguration Statistics;

        [OutputConstructor]
        private JobColumnStatisticsConfiguration(
            ImmutableArray<Outputs.JobColumnSelector> selectors,

            Outputs.JobStatisticsConfiguration statistics)
        {
            Selectors = selectors;
            Statistics = statistics;
        }
    }
}
