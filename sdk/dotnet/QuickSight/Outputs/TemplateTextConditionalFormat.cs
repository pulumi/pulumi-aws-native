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
    public sealed class TemplateTextConditionalFormat
    {
        public readonly Outputs.TemplateConditionalFormattingColor? BackgroundColor;
        public readonly Outputs.TemplateConditionalFormattingIcon? Icon;
        public readonly Outputs.TemplateConditionalFormattingColor? TextColor;

        [OutputConstructor]
        private TemplateTextConditionalFormat(
            Outputs.TemplateConditionalFormattingColor? backgroundColor,

            Outputs.TemplateConditionalFormattingIcon? icon,

            Outputs.TemplateConditionalFormattingColor? textColor)
        {
            BackgroundColor = backgroundColor;
            Icon = icon;
            TextColor = textColor;
        }
    }
}