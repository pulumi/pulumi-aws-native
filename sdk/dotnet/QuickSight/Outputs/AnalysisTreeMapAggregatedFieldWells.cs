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
    public sealed class AnalysisTreeMapAggregatedFieldWells
    {
        /// <summary>
        /// The color field well of a tree map. Values are grouped by aggregations based on group by fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> Colors;
        /// <summary>
        /// The group by field well of a tree map. Values are grouped based on group by fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> Groups;
        /// <summary>
        /// The size field well of a tree map. Values are aggregated based on group by fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> Sizes;

        [OutputConstructor]
        private AnalysisTreeMapAggregatedFieldWells(
            ImmutableArray<Outputs.AnalysisMeasureField> colors,

            ImmutableArray<Outputs.AnalysisDimensionField> groups,

            ImmutableArray<Outputs.AnalysisMeasureField> sizes)
        {
            Colors = colors;
            Groups = groups;
            Sizes = sizes;
        }
    }
}
