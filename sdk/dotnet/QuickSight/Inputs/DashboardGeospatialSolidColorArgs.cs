// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// Describes the properties for a solid color
    /// </summary>
    public sealed class DashboardGeospatialSolidColorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color and opacity values for the color.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Enables and disables the view state of the color.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardGeospatialColorState>? State { get; set; }

        public DashboardGeospatialSolidColorArgs()
        {
        }
        public static new DashboardGeospatialSolidColorArgs Empty => new DashboardGeospatialSolidColorArgs();
    }
}
