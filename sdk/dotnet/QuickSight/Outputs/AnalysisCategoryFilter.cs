// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisCategoryFilter
    {
        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        public readonly Outputs.AnalysisColumnIdentifier Column;
        /// <summary>
        /// The configuration for a `CategoryFilter` .
        /// </summary>
        public readonly Outputs.AnalysisCategoryFilterConfiguration Configuration;
        /// <summary>
        /// The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.
        /// </summary>
        public readonly Outputs.AnalysisDefaultFilterControlConfiguration? DefaultFilterControlConfiguration;
        /// <summary>
        /// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
        /// </summary>
        public readonly string FilterId;

        [OutputConstructor]
        private AnalysisCategoryFilter(
            Outputs.AnalysisColumnIdentifier column,

            Outputs.AnalysisCategoryFilterConfiguration configuration,

            Outputs.AnalysisDefaultFilterControlConfiguration? defaultFilterControlConfiguration,

            string filterId)
        {
            Column = column;
            Configuration = configuration;
            DefaultFilterControlConfiguration = defaultFilterControlConfiguration;
            FilterId = filterId;
        }
    }
}
