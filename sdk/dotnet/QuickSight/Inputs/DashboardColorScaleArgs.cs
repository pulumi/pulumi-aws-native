// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardColorScaleArgs : global::Pulumi.ResourceArgs
    {
        [Input("colorFillType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DashboardColorFillType> ColorFillType { get; set; } = null!;

        [Input("colors", required: true)]
        private InputList<Inputs.DashboardDataColorArgs>? _colors;
        public InputList<Inputs.DashboardDataColorArgs> Colors
        {
            get => _colors ?? (_colors = new InputList<Inputs.DashboardDataColorArgs>());
            set => _colors = value;
        }

        [Input("nullValueColor")]
        public Input<Inputs.DashboardDataColorArgs>? NullValueColor { get; set; }

        public DashboardColorScaleArgs()
        {
        }
        public static new DashboardColorScaleArgs Empty => new DashboardColorScaleArgs();
    }
}