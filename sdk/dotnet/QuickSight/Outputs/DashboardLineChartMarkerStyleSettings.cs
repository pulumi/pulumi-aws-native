// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardLineChartMarkerStyleSettings
    {
        public readonly string? MarkerColor;
        public readonly Pulumi.AwsNative.QuickSight.DashboardLineChartMarkerShape? MarkerShape;
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? MarkerSize;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? MarkerVisibility;

        [OutputConstructor]
        private DashboardLineChartMarkerStyleSettings(
            string? markerColor,

            Pulumi.AwsNative.QuickSight.DashboardLineChartMarkerShape? markerShape,

            string? markerSize,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? markerVisibility)
        {
            MarkerColor = markerColor;
            MarkerShape = markerShape;
            MarkerSize = markerSize;
            MarkerVisibility = markerVisibility;
        }
    }
}