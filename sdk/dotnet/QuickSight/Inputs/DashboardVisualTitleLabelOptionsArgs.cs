// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardVisualTitleLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The short text format of the title label, such as plain text or rich text.
        /// </summary>
        [Input("formatText")]
        public Input<Inputs.DashboardShortFormatTextArgs>? FormatText { get; set; }

        /// <summary>
        /// The visibility of the title label.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardVisualTitleLabelOptionsArgs()
        {
        }
        public static new DashboardVisualTitleLabelOptionsArgs Empty => new DashboardVisualTitleLabelOptionsArgs();
    }
}
