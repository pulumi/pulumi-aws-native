// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableFieldCollapseStateOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("state")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardPivotTableFieldCollapseState>? State { get; set; }

        [Input("target", required: true)]
        public Input<Inputs.DashboardPivotTableFieldCollapseStateTargetArgs> Target { get; set; } = null!;

        public DashboardPivotTableFieldCollapseStateOptionArgs()
        {
        }
        public static new DashboardPivotTableFieldCollapseStateOptionArgs Empty => new DashboardPivotTableFieldCollapseStateOptionArgs();
    }
}