// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateConditionalFormattingCustomIconConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the color of the icon.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Determines the icon display configuration.
        /// </summary>
        [Input("displayConfiguration")]
        public Input<Inputs.TemplateConditionalFormattingIconDisplayConfigurationArgs>? DisplayConfiguration { get; set; }

        /// <summary>
        /// The expression that determines the condition of the icon set.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// Custom icon options for an icon set.
        /// </summary>
        [Input("iconOptions", required: true)]
        public Input<Inputs.TemplateConditionalFormattingCustomIconOptionsArgs> IconOptions { get; set; } = null!;

        public TemplateConditionalFormattingCustomIconConditionArgs()
        {
        }
        public static new TemplateConditionalFormattingCustomIconConditionArgs Empty => new TemplateConditionalFormattingCustomIconConditionArgs();
    }
}
