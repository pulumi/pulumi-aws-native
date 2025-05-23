// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFilterListControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
        /// </summary>
        [Input("cascadingControlConfiguration")]
        public Input<Inputs.TemplateCascadingControlConfigurationArgs>? CascadingControlConfiguration { get; set; }

        /// <summary>
        /// The display options of a control.
        /// </summary>
        [Input("displayOptions")]
        public Input<Inputs.TemplateListControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        /// <summary>
        /// The ID of the `FilterListControl` .
        /// </summary>
        [Input("filterControlId", required: true)]
        public Input<string> FilterControlId { get; set; } = null!;

        /// <summary>
        /// A list of selectable values that are used in a control.
        /// </summary>
        [Input("selectableValues")]
        public Input<Inputs.TemplateFilterSelectableValuesArgs>? SelectableValues { get; set; }

        /// <summary>
        /// The source filter ID of the `FilterListControl` .
        /// </summary>
        [Input("sourceFilterId", required: true)]
        public Input<string> SourceFilterId { get; set; } = null!;

        /// <summary>
        /// The title of the `FilterListControl` .
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The type of the `FilterListControl` . Choose one of the following options:
        /// 
        /// - `MULTI_SELECT` : The user can select multiple entries from the list.
        /// - `SINGLE_SELECT` : The user can select a single entry from the list.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSheetControlListType>? Type { get; set; }

        public TemplateFilterListControlArgs()
        {
        }
        public static new TemplateFilterListControlArgs Empty => new TemplateFilterListControlArgs();
    }
}
