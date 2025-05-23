// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardGradientStopArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the color.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Determines the data value.
        /// </summary>
        [Input("dataValue")]
        public Input<double>? DataValue { get; set; }

        /// <summary>
        /// Determines gradient offset value.
        /// </summary>
        [Input("gradientOffset", required: true)]
        public Input<double> GradientOffset { get; set; } = null!;

        public DashboardGradientStopArgs()
        {
        }
        public static new DashboardGradientStopArgs Empty => new DashboardGradientStopArgs();
    }
}
