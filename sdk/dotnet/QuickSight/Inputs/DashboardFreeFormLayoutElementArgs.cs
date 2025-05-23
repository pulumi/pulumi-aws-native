// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFreeFormLayoutElementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The background style configuration of a free-form layout element.
        /// </summary>
        [Input("backgroundStyle")]
        public Input<Inputs.DashboardFreeFormLayoutElementBackgroundStyleArgs>? BackgroundStyle { get; set; }

        /// <summary>
        /// The border style configuration of a free-form layout element.
        /// </summary>
        [Input("borderStyle")]
        public Input<Inputs.DashboardFreeFormLayoutElementBorderStyleArgs>? BorderStyle { get; set; }

        /// <summary>
        /// A unique identifier for an element within a free-form layout.
        /// </summary>
        [Input("elementId", required: true)]
        public Input<string> ElementId { get; set; } = null!;

        /// <summary>
        /// The type of element.
        /// </summary>
        [Input("elementType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DashboardLayoutElementType> ElementType { get; set; } = null!;

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("height", required: true)]
        public Input<string> Height { get; set; } = null!;

        /// <summary>
        /// The loading animation configuration of a free-form layout element.
        /// </summary>
        [Input("loadingAnimation")]
        public Input<Inputs.DashboardLoadingAnimationArgs>? LoadingAnimation { get; set; }

        [Input("renderingRules")]
        private InputList<Inputs.DashboardSheetElementRenderingRuleArgs>? _renderingRules;

        /// <summary>
        /// The rendering rules that determine when an element should be displayed within a free-form layout.
        /// </summary>
        public InputList<Inputs.DashboardSheetElementRenderingRuleArgs> RenderingRules
        {
            get => _renderingRules ?? (_renderingRules = new InputList<Inputs.DashboardSheetElementRenderingRuleArgs>());
            set => _renderingRules = value;
        }

        /// <summary>
        /// The border style configuration of a free-form layout element. This border style is used when the element is selected.
        /// </summary>
        [Input("selectedBorderStyle")]
        public Input<Inputs.DashboardFreeFormLayoutElementBorderStyleArgs>? SelectedBorderStyle { get; set; }

        /// <summary>
        /// The visibility of an element within a free-form layout.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardVisibility>? Visibility { get; set; }

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("width", required: true)]
        public Input<string> Width { get; set; } = null!;

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("xAxisLocation", required: true)]
        public Input<string> XAxisLocation { get; set; } = null!;

        /// <summary>
        /// String based length that is composed of value and unit in px with Integer.MAX_VALUE as maximum value
        /// </summary>
        [Input("yAxisLocation", required: true)]
        public Input<string> YAxisLocation { get; set; } = null!;

        public DashboardFreeFormLayoutElementArgs()
        {
        }
        public static new DashboardFreeFormLayoutElementArgs Empty => new DashboardFreeFormLayoutElementArgs();
    }
}
