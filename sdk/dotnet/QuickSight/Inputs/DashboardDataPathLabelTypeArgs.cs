// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDataPathLabelTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId")]
        public Input<string>? FieldId { get; set; }

        [Input("fieldValue")]
        public Input<string>? FieldValue { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardDataPathLabelTypeArgs()
        {
        }
        public static new DashboardDataPathLabelTypeArgs Empty => new DashboardDataPathLabelTypeArgs();
    }
}