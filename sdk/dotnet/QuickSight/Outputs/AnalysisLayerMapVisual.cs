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
    public sealed class AnalysisLayerMapVisual
    {
        /// <summary>
        /// The configuration settings of the visual.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialLayerMapConfiguration? ChartConfiguration;
        /// <summary>
        /// The dataset that is used to create the layer map visual. You can't create a visual without a dataset.
        /// </summary>
        public readonly string DataSetIdentifier;
        public readonly Outputs.AnalysisVisualSubtitleLabelOptions? Subtitle;
        public readonly Outputs.AnalysisVisualTitleLabelOptions? Title;
        /// <summary>
        /// The alt text for the visual.
        /// </summary>
        public readonly string? VisualContentAltText;
        /// <summary>
        /// The ID of the visual.
        /// </summary>
        public readonly string VisualId;

        [OutputConstructor]
        private AnalysisLayerMapVisual(
            Outputs.AnalysisGeospatialLayerMapConfiguration? chartConfiguration,

            string dataSetIdentifier,

            Outputs.AnalysisVisualSubtitleLabelOptions? subtitle,

            Outputs.AnalysisVisualTitleLabelOptions? title,

            string? visualContentAltText,

            string visualId)
        {
            ChartConfiguration = chartConfiguration;
            DataSetIdentifier = dataSetIdentifier;
            Subtitle = subtitle;
            Title = title;
            VisualContentAltText = visualContentAltText;
            VisualId = visualId;
        }
    }
}
