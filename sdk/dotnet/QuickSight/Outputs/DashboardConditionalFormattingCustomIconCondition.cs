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
    public sealed class DashboardConditionalFormattingCustomIconCondition
    {
        /// <summary>
        /// Determines the color of the icon.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Determines the icon display configuration.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingIconDisplayConfiguration? DisplayConfiguration;
        /// <summary>
        /// The expression that determines the condition of the icon set.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// Custom icon options for an icon set.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingCustomIconOptions IconOptions;

        [OutputConstructor]
        private DashboardConditionalFormattingCustomIconCondition(
            string? color,

            Outputs.DashboardConditionalFormattingIconDisplayConfiguration? displayConfiguration,

            string expression,

            Outputs.DashboardConditionalFormattingCustomIconOptions iconOptions)
        {
            Color = color;
            DisplayConfiguration = displayConfiguration;
            Expression = expression;
            IconOptions = iconOptions;
        }
    }
}
