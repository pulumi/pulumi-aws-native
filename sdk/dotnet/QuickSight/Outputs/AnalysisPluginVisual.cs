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
    public sealed class AnalysisPluginVisual
    {
        /// <summary>
        /// A description of the plugin field wells and their persisted properties.
        /// </summary>
        public readonly Outputs.AnalysisPluginVisualConfiguration? ChartConfiguration;
        /// <summary>
        /// The Amazon Resource Name (ARN) that reflects the plugin and version.
        /// </summary>
        public readonly string PluginArn;
        public readonly Outputs.AnalysisVisualSubtitleLabelOptions? Subtitle;
        public readonly Outputs.AnalysisVisualTitleLabelOptions? Title;
        /// <summary>
        /// The alt text for the visual.
        /// </summary>
        public readonly string? VisualContentAltText;
        /// <summary>
        /// The ID of the visual that you want to use.
        /// </summary>
        public readonly string VisualId;

        [OutputConstructor]
        private AnalysisPluginVisual(
            Outputs.AnalysisPluginVisualConfiguration? chartConfiguration,

            string pluginArn,

            Outputs.AnalysisVisualSubtitleLabelOptions? subtitle,

            Outputs.AnalysisVisualTitleLabelOptions? title,

            string? visualContentAltText,

            string visualId)
        {
            ChartConfiguration = chartConfiguration;
            PluginArn = pluginArn;
            Subtitle = subtitle;
            Title = title;
            VisualContentAltText = visualContentAltText;
            VisualId = visualId;
        }
    }
}
