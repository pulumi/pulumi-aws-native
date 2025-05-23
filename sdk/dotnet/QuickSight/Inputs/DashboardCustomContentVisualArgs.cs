// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardCustomContentVisualArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.DashboardVisualCustomActionArgs>? _actions;

        /// <summary>
        /// The list of custom actions that are configured for a visual.
        /// </summary>
        public InputList<Inputs.DashboardVisualCustomActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.DashboardVisualCustomActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The configuration of a `CustomContentVisual` .
        /// </summary>
        [Input("chartConfiguration")]
        public Input<Inputs.DashboardCustomContentConfigurationArgs>? ChartConfiguration { get; set; }

        /// <summary>
        /// The dataset that is used to create the custom content visual. You can't create a visual without a dataset.
        /// </summary>
        [Input("dataSetIdentifier", required: true)]
        public Input<string> DataSetIdentifier { get; set; } = null!;

        /// <summary>
        /// The subtitle that is displayed on the visual.
        /// </summary>
        [Input("subtitle")]
        public Input<Inputs.DashboardVisualSubtitleLabelOptionsArgs>? Subtitle { get; set; }

        /// <summary>
        /// The title that is displayed on the visual.
        /// </summary>
        [Input("title")]
        public Input<Inputs.DashboardVisualTitleLabelOptionsArgs>? Title { get; set; }

        /// <summary>
        /// The alt text for the visual.
        /// </summary>
        [Input("visualContentAltText")]
        public Input<string>? VisualContentAltText { get; set; }

        /// <summary>
        /// The unique identifier of a visual. This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
        /// </summary>
        [Input("visualId", required: true)]
        public Input<string> VisualId { get; set; } = null!;

        public DashboardCustomContentVisualArgs()
        {
        }
        public static new DashboardCustomContentVisualArgs Empty => new DashboardCustomContentVisualArgs();
    }
}
