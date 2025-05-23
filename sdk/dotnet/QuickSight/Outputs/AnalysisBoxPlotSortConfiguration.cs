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
    public sealed class AnalysisBoxPlotSortConfiguration
    {
        /// <summary>
        /// The sort configuration of a group by fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisFieldSortOptions> CategorySort;
        /// <summary>
        /// The pagination configuration of a table visual or box plot.
        /// </summary>
        public readonly Outputs.AnalysisPaginationConfiguration? PaginationConfiguration;

        [OutputConstructor]
        private AnalysisBoxPlotSortConfiguration(
            ImmutableArray<Outputs.AnalysisFieldSortOptions> categorySort,

            Outputs.AnalysisPaginationConfiguration? paginationConfiguration)
        {
            CategorySort = categorySort;
            PaginationConfiguration = paginationConfiguration;
        }
    }
}
