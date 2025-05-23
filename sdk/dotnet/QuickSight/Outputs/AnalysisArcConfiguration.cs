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
    public sealed class AnalysisArcConfiguration
    {
        /// <summary>
        /// The option that determines the arc angle of a `GaugeChartVisual` .
        /// </summary>
        public readonly double? ArcAngle;
        /// <summary>
        /// The options that determine the arc thickness of a `GaugeChartVisual` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisArcThicknessOptions? ArcThickness;

        [OutputConstructor]
        private AnalysisArcConfiguration(
            double? arcAngle,

            Pulumi.AwsNative.QuickSight.AnalysisArcThicknessOptions? arcThickness)
        {
            ArcAngle = arcAngle;
            ArcThickness = arcThickness;
        }
    }
}
