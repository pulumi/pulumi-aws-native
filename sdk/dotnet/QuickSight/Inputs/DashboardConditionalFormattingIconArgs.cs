// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardConditionalFormattingIconArgs : global::Pulumi.ResourceArgs
    {
        [Input("customCondition")]
        public Input<Inputs.DashboardConditionalFormattingCustomIconConditionArgs>? CustomCondition { get; set; }

        [Input("iconSet")]
        public Input<Inputs.DashboardConditionalFormattingIconSetArgs>? IconSet { get; set; }

        public DashboardConditionalFormattingIconArgs()
        {
        }
        public static new DashboardConditionalFormattingIconArgs Empty => new DashboardConditionalFormattingIconArgs();
    }
}