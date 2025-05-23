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
    public sealed class AnalysisSankeyDiagramSortConfiguration
    {
        /// <summary>
        /// The limit on the number of destination nodes that are displayed in a sankey diagram.
        /// </summary>
        public readonly Outputs.AnalysisItemsLimitConfiguration? DestinationItemsLimit;
        /// <summary>
        /// The limit on the number of source nodes that are displayed in a sankey diagram.
        /// </summary>
        public readonly Outputs.AnalysisItemsLimitConfiguration? SourceItemsLimit;
        /// <summary>
        /// The sort configuration of the weight fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisFieldSortOptions> WeightSort;

        [OutputConstructor]
        private AnalysisSankeyDiagramSortConfiguration(
            Outputs.AnalysisItemsLimitConfiguration? destinationItemsLimit,

            Outputs.AnalysisItemsLimitConfiguration? sourceItemsLimit,

            ImmutableArray<Outputs.AnalysisFieldSortOptions> weightSort)
        {
            DestinationItemsLimit = destinationItemsLimit;
            SourceItemsLimit = sourceItemsLimit;
            WeightSort = weightSort;
        }
    }
}
