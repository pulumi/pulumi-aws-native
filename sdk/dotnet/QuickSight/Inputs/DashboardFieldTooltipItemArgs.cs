// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// The unique ID of the field that is targeted by the tooltip.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        /// <summary>
        /// The label of the tooltip item.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Determines the target of the field tooltip item in a combo chart visual.
        /// </summary>
        [Input("tooltipTarget")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardTooltipTarget>? TooltipTarget { get; set; }

        /// <summary>
        /// The visibility of the tooltip item.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardFieldTooltipItemArgs()
        {
        }
        public static new DashboardFieldTooltipItemArgs Empty => new DashboardFieldTooltipItemArgs();
    }
}
