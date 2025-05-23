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
    public sealed class TemplateFunnelChartAggregatedFieldWells
    {
        /// <summary>
        /// The category field wells of a funnel chart. Values are grouped by category fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateDimensionField> Category;
        /// <summary>
        /// The value field wells of a funnel chart. Values are aggregated based on categories.
        /// </summary>
        public readonly ImmutableArray<Outputs.TemplateMeasureField> Values;

        [OutputConstructor]
        private TemplateFunnelChartAggregatedFieldWells(
            ImmutableArray<Outputs.TemplateDimensionField> category,

            ImmutableArray<Outputs.TemplateMeasureField> values)
        {
            Category = category;
            Values = values;
        }
    }
}
