// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardNumericEqualityDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public DashboardNumericEqualityDrillDownFilterArgs()
        {
        }
        public static new DashboardNumericEqualityDrillDownFilterArgs Empty => new DashboardNumericEqualityDrillDownFilterArgs();
    }
}