// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardVisualTitleLabelOptions
    {
        /// <summary>
        /// The short text format of the title label, such as plain text or rich text.
        /// </summary>
        public readonly Outputs.DashboardShortFormatText? FormatText;
        /// <summary>
        /// The visibility of the title label.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? Visibility;

        [OutputConstructor]
        private DashboardVisualTitleLabelOptions(
            Outputs.DashboardShortFormatText? formatText,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? visibility)
        {
            FormatText = formatText;
            Visibility = visibility;
        }
    }
}
