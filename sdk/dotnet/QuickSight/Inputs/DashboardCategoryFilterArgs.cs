// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardCategoryFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("configuration", required: true)]
        public Input<Inputs.DashboardCategoryFilterConfigurationArgs> Configuration { get; set; } = null!;

        [Input("filterId", required: true)]
        public Input<string> FilterId { get; set; } = null!;

        public DashboardCategoryFilterArgs()
        {
        }
        public static new DashboardCategoryFilterArgs Empty => new DashboardCategoryFilterArgs();
    }
}