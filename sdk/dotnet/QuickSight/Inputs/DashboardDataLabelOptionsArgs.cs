// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDataLabelOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the visibility of the category field labels.
        /// </summary>
        [Input("categoryLabelVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? CategoryLabelVisibility { get; set; }

        [Input("dataLabelTypes")]
        private InputList<Inputs.DashboardDataLabelTypeArgs>? _dataLabelTypes;

        /// <summary>
        /// The option that determines the data label type.
        /// </summary>
        public InputList<Inputs.DashboardDataLabelTypeArgs> DataLabelTypes
        {
            get => _dataLabelTypes ?? (_dataLabelTypes = new InputList<Inputs.DashboardDataLabelTypeArgs>());
            set => _dataLabelTypes = value;
        }

        /// <summary>
        /// Determines the color of the data labels.
        /// </summary>
        [Input("labelColor")]
        public Input<string>? LabelColor { get; set; }

        /// <summary>
        /// Determines the content of the data labels.
        /// </summary>
        [Input("labelContent")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardDataLabelContent>? LabelContent { get; set; }

        /// <summary>
        /// Determines the font configuration of the data labels.
        /// </summary>
        [Input("labelFontConfiguration")]
        public Input<Inputs.DashboardFontConfigurationArgs>? LabelFontConfiguration { get; set; }

        /// <summary>
        /// Determines the visibility of the measure field labels.
        /// </summary>
        [Input("measureLabelVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? MeasureLabelVisibility { get; set; }

        /// <summary>
        /// Determines whether overlap is enabled or disabled for the data labels.
        /// </summary>
        [Input("overlap")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardDataLabelOverlap>? Overlap { get; set; }

        /// <summary>
        /// Determines the position of the data labels.
        /// </summary>
        [Input("position")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardDataLabelPosition>? Position { get; set; }

        /// <summary>
        /// Determines the visibility of the total.
        /// </summary>
        [Input("totalsVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? TotalsVisibility { get; set; }

        /// <summary>
        /// Determines the visibility of the data labels.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        public DashboardDataLabelOptionsArgs()
        {
        }
        public static new DashboardDataLabelOptionsArgs Empty => new DashboardDataLabelOptionsArgs();
    }
}
