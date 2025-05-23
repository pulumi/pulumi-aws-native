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
    public sealed class JobStatisticsConfiguration
    {
        /// <summary>
        /// List of included evaluations. When the list is undefined, all supported evaluations will be included.
        /// </summary>
        public readonly ImmutableArray<string> IncludedStatistics;
        /// <summary>
        /// List of overrides for evaluations.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobStatisticOverride> Overrides;

        [OutputConstructor]
        private JobStatisticsConfiguration(
            ImmutableArray<string> includedStatistics,

            ImmutableArray<Outputs.JobStatisticOverride> overrides)
        {
            IncludedStatistics = includedStatistics;
            Overrides = overrides;
        }
    }
}
