// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardLineChartSeriesSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("lineStyleSettings")]
        public Input<Inputs.DashboardLineChartLineStyleSettingsArgs>? LineStyleSettings { get; set; }

        [Input("markerStyleSettings")]
        public Input<Inputs.DashboardLineChartMarkerStyleSettingsArgs>? MarkerStyleSettings { get; set; }

        public DashboardLineChartSeriesSettingsArgs()
        {
        }
        public static new DashboardLineChartSeriesSettingsArgs Empty => new DashboardLineChartSeriesSettingsArgs();
    }
}