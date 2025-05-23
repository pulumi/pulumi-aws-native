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
    public sealed class DashboardKpiProgressBarConditionalFormatting
    {
        /// <summary>
        /// The conditional formatting of the progress bar's foreground color.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingColor? ForegroundColor;

        [OutputConstructor]
        private DashboardKpiProgressBarConditionalFormatting(Outputs.DashboardConditionalFormattingColor? foregroundColor)
        {
            ForegroundColor = foregroundColor;
        }
    }
}
