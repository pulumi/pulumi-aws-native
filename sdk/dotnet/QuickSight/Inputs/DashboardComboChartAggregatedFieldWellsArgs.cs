// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardComboChartAggregatedFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("barValues")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _barValues;
        public InputList<Inputs.DashboardMeasureFieldArgs> BarValues
        {
            get => _barValues ?? (_barValues = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _barValues = value;
        }

        [Input("category")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _category;
        public InputList<Inputs.DashboardDimensionFieldArgs> Category
        {
            get => _category ?? (_category = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _category = value;
        }

        [Input("colors")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _colors;
        public InputList<Inputs.DashboardDimensionFieldArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _colors = value;
        }

        [Input("lineValues")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _lineValues;
        public InputList<Inputs.DashboardMeasureFieldArgs> LineValues
        {
            get => _lineValues ?? (_lineValues = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _lineValues = value;
        }

        public DashboardComboChartAggregatedFieldWellsArgs()
        {
        }
        public static new DashboardComboChartAggregatedFieldWellsArgs Empty => new DashboardComboChartAggregatedFieldWellsArgs();
    }
}