// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGaugeChartFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetValues")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _targetValues;
        public InputList<Inputs.DashboardMeasureFieldArgs> TargetValues
        {
            get => _targetValues ?? (_targetValues = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _targetValues = value;
        }

        [Input("values")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _values;
        public InputList<Inputs.DashboardMeasureFieldArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _values = value;
        }

        public DashboardGaugeChartFieldWellsArgs()
        {
        }
        public static new DashboardGaugeChartFieldWellsArgs Empty => new DashboardGaugeChartFieldWellsArgs();
    }
}