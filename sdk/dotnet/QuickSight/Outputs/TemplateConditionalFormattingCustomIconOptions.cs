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
    public sealed class TemplateConditionalFormattingCustomIconOptions
    {
        public readonly Pulumi.AwsNative.QuickSight.TemplateIcon? Icon;
        public readonly string? UnicodeIcon;

        [OutputConstructor]
        private TemplateConditionalFormattingCustomIconOptions(
            Pulumi.AwsNative.QuickSight.TemplateIcon? icon,

            string? unicodeIcon)
        {
            Icon = icon;
            UnicodeIcon = unicodeIcon;
        }
    }
}