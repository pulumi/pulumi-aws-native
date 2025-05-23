// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTableCellStyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The background color for the table cells.
        /// </summary>
        [Input("backgroundColor")]
        public Input<string>? BackgroundColor { get; set; }

        /// <summary>
        /// The borders for the table cells.
        /// </summary>
        [Input("border")]
        public Input<Inputs.DashboardGlobalTableBorderOptionsArgs>? Border { get; set; }

        /// <summary>
        /// The font configuration of the table cells.
        /// </summary>
        [Input("fontConfiguration")]
        public Input<Inputs.DashboardFontConfigurationArgs>? FontConfiguration { get; set; }

        /// <summary>
        /// The height color for the table cells.
        /// </summary>
        [Input("height")]
        public Input<double>? Height { get; set; }

        /// <summary>
        /// The horizontal text alignment (left, center, right, auto) for the table cells.
        /// </summary>
        [Input("horizontalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardHorizontalTextAlignment>? HorizontalTextAlignment { get; set; }

        /// <summary>
        /// The text wrap (none, wrap) for the table cells.
        /// </summary>
        [Input("textWrap")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardTextWrap>? TextWrap { get; set; }

        /// <summary>
        /// The vertical text alignment (top, middle, bottom) for the table cells.
        /// </summary>
        [Input("verticalTextAlignment")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVerticalTextAlignment>? VerticalTextAlignment { get; set; }

        /// <summary>
        /// The visibility of the table cells.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardTableCellStyleArgs()
        {
        }
        public static new DashboardTableCellStyleArgs Empty => new DashboardTableCellStyleArgs();
    }
}
