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
    public sealed class DashboardConditionalFormattingColor
    {
        /// <summary>
        /// Formatting configuration for gradient color.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingGradientColor? Gradient;
        /// <summary>
        /// Formatting configuration for solid color.
        /// </summary>
        public readonly Outputs.DashboardConditionalFormattingSolidColor? Solid;

        [OutputConstructor]
        private DashboardConditionalFormattingColor(
            Outputs.DashboardConditionalFormattingGradientColor? gradient,

            Outputs.DashboardConditionalFormattingSolidColor? solid)
        {
            Gradient = gradient;
            Solid = solid;
        }
    }
}
