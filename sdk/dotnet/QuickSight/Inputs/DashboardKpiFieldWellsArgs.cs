// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardKpiFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetValues")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _targetValues;

        /// <summary>
        /// The target value field wells of a KPI visual.
        /// </summary>
        public InputList<Inputs.DashboardMeasureFieldArgs> TargetValues
        {
            get => _targetValues ?? (_targetValues = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _targetValues = value;
        }

        [Input("trendGroups")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _trendGroups;

        /// <summary>
        /// The trend group field wells of a KPI visual.
        /// </summary>
        public InputList<Inputs.DashboardDimensionFieldArgs> TrendGroups
        {
            get => _trendGroups ?? (_trendGroups = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _trendGroups = value;
        }

        [Input("values")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _values;

        /// <summary>
        /// The value field wells of a KPI visual.
        /// </summary>
        public InputList<Inputs.DashboardMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _values = value;
        }

        public DashboardKpiFieldWellsArgs()
        {
        }
        public static new DashboardKpiFieldWellsArgs Empty => new DashboardKpiFieldWellsArgs();
    }
}
