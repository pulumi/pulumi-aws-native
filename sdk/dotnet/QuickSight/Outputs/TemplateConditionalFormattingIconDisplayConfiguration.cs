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
    public sealed class TemplateConditionalFormattingIconDisplayConfiguration
    {
        /// <summary>
        /// Determines the icon display configuration.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateConditionalFormattingIconDisplayOption? IconDisplayOption;

        [OutputConstructor]
        private TemplateConditionalFormattingIconDisplayConfiguration(Pulumi.AwsNative.QuickSight.TemplateConditionalFormattingIconDisplayOption? iconDisplayOption)
        {
            IconDisplayOption = iconDisplayOption;
        }
    }
}
