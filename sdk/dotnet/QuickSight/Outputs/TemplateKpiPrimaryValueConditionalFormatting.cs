// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateKpiPrimaryValueConditionalFormatting
    {
        public readonly Outputs.TemplateConditionalFormattingIcon? Icon;
        public readonly Outputs.TemplateConditionalFormattingColor? TextColor;

        [OutputConstructor]
        private TemplateKpiPrimaryValueConditionalFormatting(
            Outputs.TemplateConditionalFormattingIcon? icon,

            Outputs.TemplateConditionalFormattingColor? textColor)
        {
            Icon = icon;
            TextColor = textColor;
        }
    }
}