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
    public sealed class AnalysisBoxPlotStyleOptions
    {
        /// <summary>
        /// The fill styles (solid, transparent) of the box plot.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisBoxPlotFillStyle? FillStyle;

        [OutputConstructor]
        private AnalysisBoxPlotStyleOptions(Pulumi.AwsNative.QuickSight.AnalysisBoxPlotFillStyle? fillStyle)
        {
            FillStyle = fillStyle;
        }
    }
}
