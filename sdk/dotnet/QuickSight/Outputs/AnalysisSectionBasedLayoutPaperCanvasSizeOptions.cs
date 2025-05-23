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
    public sealed class AnalysisSectionBasedLayoutPaperCanvasSizeOptions
    {
        /// <summary>
        /// Defines the spacing between the canvas content and the top, bottom, left, and right edges.
        /// </summary>
        public readonly Outputs.AnalysisSpacing? PaperMargin;
        /// <summary>
        /// The paper orientation that is used to define canvas dimensions. Choose one of the following options:
        /// 
        /// - PORTRAIT
        /// - LANDSCAPE
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisPaperOrientation? PaperOrientation;
        /// <summary>
        /// The paper size that is used to define canvas dimensions.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisPaperSize? PaperSize;

        [OutputConstructor]
        private AnalysisSectionBasedLayoutPaperCanvasSizeOptions(
            Outputs.AnalysisSpacing? paperMargin,

            Pulumi.AwsNative.QuickSight.AnalysisPaperOrientation? paperOrientation,

            Pulumi.AwsNative.QuickSight.AnalysisPaperSize? paperSize)
        {
            PaperMargin = paperMargin;
            PaperOrientation = paperOrientation;
            PaperSize = paperSize;
        }
    }
}
