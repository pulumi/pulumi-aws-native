// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisFontConfiguration
    {
        public readonly string? FontColor;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisFontDecoration? FontDecoration;
        public readonly Outputs.AnalysisFontSize? FontSize;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisFontStyle? FontStyle;
        public readonly Outputs.AnalysisFontWeight? FontWeight;

        [OutputConstructor]
        private AnalysisFontConfiguration(
            string? fontColor,

            Pulumi.AwsNative.QuickSight.AnalysisFontDecoration? fontDecoration,

            Outputs.AnalysisFontSize? fontSize,

            Pulumi.AwsNative.QuickSight.AnalysisFontStyle? fontStyle,

            Outputs.AnalysisFontWeight? fontWeight)
        {
            FontColor = fontColor;
            FontDecoration = fontDecoration;
            FontSize = fontSize;
            FontStyle = fontStyle;
            FontWeight = fontWeight;
        }
    }
}