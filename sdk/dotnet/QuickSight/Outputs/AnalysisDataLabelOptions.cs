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
    public sealed class AnalysisDataLabelOptions
    {
        /// <summary>
        /// Determines the visibility of the category field labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? CategoryLabelVisibility;
        /// <summary>
        /// The option that determines the data label type.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisDataLabelType> DataLabelTypes;
        /// <summary>
        /// Determines the color of the data labels.
        /// </summary>
        public readonly string? LabelColor;
        /// <summary>
        /// Determines the content of the data labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisDataLabelContent? LabelContent;
        /// <summary>
        /// Determines the font configuration of the data labels.
        /// </summary>
        public readonly Outputs.AnalysisFontConfiguration? LabelFontConfiguration;
        /// <summary>
        /// Determines the visibility of the measure field labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? MeasureLabelVisibility;
        /// <summary>
        /// Determines whether overlap is enabled or disabled for the data labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisDataLabelOverlap? Overlap;
        /// <summary>
        /// Determines the position of the data labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisDataLabelPosition? Position;
        /// <summary>
        /// Determines the visibility of the total.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? TotalsVisibility;
        /// <summary>
        /// Determines the visibility of the data labels.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? Visibility;

        [OutputConstructor]
        private AnalysisDataLabelOptions(
            Pulumi.AwsNative.QuickSight.AnalysisVisibility? categoryLabelVisibility,

            ImmutableArray<Outputs.AnalysisDataLabelType> dataLabelTypes,

            string? labelColor,

            Pulumi.AwsNative.QuickSight.AnalysisDataLabelContent? labelContent,

            Outputs.AnalysisFontConfiguration? labelFontConfiguration,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? measureLabelVisibility,

            Pulumi.AwsNative.QuickSight.AnalysisDataLabelOverlap? overlap,

            Pulumi.AwsNative.QuickSight.AnalysisDataLabelPosition? position,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? totalsVisibility,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? visibility)
        {
            CategoryLabelVisibility = categoryLabelVisibility;
            DataLabelTypes = dataLabelTypes;
            LabelColor = labelColor;
            LabelContent = labelContent;
            LabelFontConfiguration = labelFontConfiguration;
            MeasureLabelVisibility = measureLabelVisibility;
            Overlap = overlap;
            Position = position;
            TotalsVisibility = totalsVisibility;
            Visibility = visibility;
        }
    }
}
