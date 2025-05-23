// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardSheetImageTooltipConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The text that appears in the tooltip.
        /// </summary>
        [Input("tooltipText")]
        public Input<Inputs.DashboardSheetImageTooltipTextArgs>? TooltipText { get; set; }

        /// <summary>
        /// The visibility of the tooltip.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardSheetImageTooltipConfigurationArgs()
        {
        }
        public static new DashboardSheetImageTooltipConfigurationArgs Empty => new DashboardSheetImageTooltipConfigurationArgs();
    }
}
