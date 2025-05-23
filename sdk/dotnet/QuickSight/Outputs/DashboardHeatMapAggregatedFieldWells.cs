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
    public sealed class DashboardHeatMapAggregatedFieldWells
    {
        /// <summary>
        /// The columns field well of a heat map.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Columns;
        /// <summary>
        /// The rows field well of a heat map.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Rows;
        /// <summary>
        /// The values field well of a heat map.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Values;

        [OutputConstructor]
        private DashboardHeatMapAggregatedFieldWells(
            ImmutableArray<Outputs.DashboardDimensionField> columns,

            ImmutableArray<Outputs.DashboardDimensionField> rows,

            ImmutableArray<Outputs.DashboardMeasureField> values)
        {
            Columns = columns;
            Rows = rows;
            Values = values;
        }
    }
}
