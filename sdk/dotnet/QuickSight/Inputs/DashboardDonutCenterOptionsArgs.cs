// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDonutCenterOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("labelVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? LabelVisibility { get; set; }

        public DashboardDonutCenterOptionsArgs()
        {
        }
        public static new DashboardDonutCenterOptionsArgs Empty => new DashboardDonutCenterOptionsArgs();
    }
}