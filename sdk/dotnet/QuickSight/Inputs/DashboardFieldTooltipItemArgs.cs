// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFieldTooltipItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardFieldTooltipItemArgs()
        {
        }
        public static new DashboardFieldTooltipItemArgs Empty => new DashboardFieldTooltipItemArgs();
    }
}