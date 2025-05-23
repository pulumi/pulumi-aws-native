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
    public sealed class AnalysisGaugeChartPrimaryValueConditionalFormatting
    {
        /// <summary>
        /// The conditional formatting of the primary value icon.
        /// </summary>
        public readonly Outputs.AnalysisConditionalFormattingIcon? Icon;
        /// <summary>
        /// The conditional formatting of the primary value text color.
        /// </summary>
        public readonly Outputs.AnalysisConditionalFormattingColor? TextColor;

        [OutputConstructor]
        private AnalysisGaugeChartPrimaryValueConditionalFormatting(
            Outputs.AnalysisConditionalFormattingIcon? icon,

            Outputs.AnalysisConditionalFormattingColor? textColor)
        {
            Icon = icon;
            TextColor = textColor;
        }
    }
}
