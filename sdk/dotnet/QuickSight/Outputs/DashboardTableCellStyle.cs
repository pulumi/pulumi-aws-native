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
    public sealed class DashboardTableCellStyle
    {
        public readonly string? BackgroundColor;
        public readonly Outputs.DashboardGlobalTableBorderOptions? Border;
        public readonly Outputs.DashboardFontConfiguration? FontConfiguration;
        public readonly double? Height;
        public readonly Pulumi.AwsNative.QuickSight.DashboardHorizontalTextAlignment? HorizontalTextAlignment;
        public readonly Pulumi.AwsNative.QuickSight.DashboardTextWrap? TextWrap;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVerticalTextAlignment? VerticalTextAlignment;
        public readonly Pulumi.AwsNative.QuickSight.DashboardVisibility? Visibility;

        [OutputConstructor]
        private DashboardTableCellStyle(
            string? backgroundColor,

            Outputs.DashboardGlobalTableBorderOptions? border,

            Outputs.DashboardFontConfiguration? fontConfiguration,

            double? height,

            Pulumi.AwsNative.QuickSight.DashboardHorizontalTextAlignment? horizontalTextAlignment,

            Pulumi.AwsNative.QuickSight.DashboardTextWrap? textWrap,

            Pulumi.AwsNative.QuickSight.DashboardVerticalTextAlignment? verticalTextAlignment,

            Pulumi.AwsNative.QuickSight.DashboardVisibility? visibility)
        {
            BackgroundColor = backgroundColor;
            Border = border;
            FontConfiguration = fontConfiguration;
            Height = height;
            HorizontalTextAlignment = horizontalTextAlignment;
            TextWrap = textWrap;
            VerticalTextAlignment = verticalTextAlignment;
            Visibility = visibility;
        }
    }
}