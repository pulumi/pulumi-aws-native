// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardBoxPlotAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupBy")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _groupBy;

        /// <summary>
        /// The group by field well of a box plot chart. Values are grouped based on group by fields.
        /// </summary>
        public InputList<Inputs.DashboardDimensionFieldArgs> GroupBy
        {
            get => _groupBy ?? (_groupBy = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _groupBy = value;
        }

        [Input("values")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _values;

        /// <summary>
        /// The value field well of a box plot chart. Values are aggregated based on group by fields.
        /// </summary>
        public InputList<Inputs.DashboardMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _values = value;
        }

        public DashboardBoxPlotAggregatedFieldWellsArgs()
        {
        }
        public static new DashboardBoxPlotAggregatedFieldWellsArgs Empty => new DashboardBoxPlotAggregatedFieldWellsArgs();
    }
}
