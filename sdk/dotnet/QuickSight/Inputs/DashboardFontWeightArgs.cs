// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFontWeightArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardFontWeightName>? Name { get; set; }

        public DashboardFontWeightArgs()
        {
        }
        public static new DashboardFontWeightArgs Empty => new DashboardFontWeightArgs();
    }
}