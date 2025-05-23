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
    public sealed class TemplateVisualSubtitleLabelOptions
    {
        /// <summary>
        /// The long text format of the subtitle label, such as plain text or rich text.
        /// </summary>
        public readonly Outputs.TemplateLongFormatText? FormatText;
        /// <summary>
        /// The visibility of the subtitle label.
        /// </summary>
        public readonly object? Visibility;

        [OutputConstructor]
        private TemplateVisualSubtitleLabelOptions(
            Outputs.TemplateLongFormatText? formatText,

            object? visibility)
        {
            FormatText = formatText;
            Visibility = visibility;
        }
    }
}
