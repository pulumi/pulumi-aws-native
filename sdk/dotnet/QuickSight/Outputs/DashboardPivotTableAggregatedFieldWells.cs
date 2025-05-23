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
    public sealed class DashboardPivotTableAggregatedFieldWells
    {
        /// <summary>
        /// The columns field well for a pivot table. Values are grouped by columns fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Columns;
        /// <summary>
        /// The rows field well for a pivot table. Values are grouped by rows fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardDimensionField> Rows;
        /// <summary>
        /// The values field well for a pivot table. Values are aggregated based on rows and columns fields.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardMeasureField> Values;

        [OutputConstructor]
        private DashboardPivotTableAggregatedFieldWells(
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
