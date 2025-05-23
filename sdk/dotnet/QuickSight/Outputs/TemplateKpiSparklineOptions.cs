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
    public sealed class TemplateKpiSparklineOptions
    {
        /// <summary>
        /// The color of the sparkline.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// The tooltip visibility of the sparkline.
        /// </summary>
        public readonly object? TooltipVisibility;
        /// <summary>
        /// The type of the sparkline.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateKpiSparklineType Type;
        /// <summary>
        /// The visibility of the sparkline.
        /// </summary>
        public readonly object? Visibility;

        [OutputConstructor]
        private TemplateKpiSparklineOptions(
            string? color,

            object? tooltipVisibility,

            Pulumi.AwsNative.QuickSight.TemplateKpiSparklineType type,

            object? visibility)
        {
            Color = color;
            TooltipVisibility = tooltipVisibility;
            Type = type;
            Visibility = visibility;
        }
    }
}
