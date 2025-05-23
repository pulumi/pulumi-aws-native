// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDefaultFilterControlOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default options that correspond to the filter control type of a `DateTimePicker` .
        /// </summary>
        [Input("defaultDateTimePickerOptions")]
        public Input<Inputs.AnalysisDefaultDateTimePickerControlOptionsArgs>? DefaultDateTimePickerOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `Dropdown` filter control type.
        /// </summary>
        [Input("defaultDropdownOptions")]
        public Input<Inputs.AnalysisDefaultFilterDropDownControlOptionsArgs>? DefaultDropdownOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `List` filter control type.
        /// </summary>
        [Input("defaultListOptions")]
        public Input<Inputs.AnalysisDefaultFilterListControlOptionsArgs>? DefaultListOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `RelativeDateTime` filter control type.
        /// </summary>
        [Input("defaultRelativeDateTimeOptions")]
        public Input<Inputs.AnalysisDefaultRelativeDateTimeControlOptionsArgs>? DefaultRelativeDateTimeOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `Slider` filter control type.
        /// </summary>
        [Input("defaultSliderOptions")]
        public Input<Inputs.AnalysisDefaultSliderControlOptionsArgs>? DefaultSliderOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `TextArea` filter control type.
        /// </summary>
        [Input("defaultTextAreaOptions")]
        public Input<Inputs.AnalysisDefaultTextAreaControlOptionsArgs>? DefaultTextAreaOptions { get; set; }

        /// <summary>
        /// The default options that correspond to the `TextField` filter control type.
        /// </summary>
        [Input("defaultTextFieldOptions")]
        public Input<Inputs.AnalysisDefaultTextFieldControlOptionsArgs>? DefaultTextFieldOptions { get; set; }

        public AnalysisDefaultFilterControlOptionsArgs()
        {
        }
        public static new AnalysisDefaultFilterControlOptionsArgs Empty => new AnalysisDefaultFilterControlOptionsArgs();
    }
}
