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
    public sealed class DashboardDefaultFilterControlOptions
    {
        /// <summary>
        /// The default options that correspond to the filter control type of a `DateTimePicker` .
        /// </summary>
        public readonly Outputs.DashboardDefaultDateTimePickerControlOptions? DefaultDateTimePickerOptions;
        /// <summary>
        /// The default options that correspond to the `Dropdown` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultFilterDropDownControlOptions? DefaultDropdownOptions;
        /// <summary>
        /// The default options that correspond to the `List` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultFilterListControlOptions? DefaultListOptions;
        /// <summary>
        /// The default options that correspond to the `RelativeDateTime` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultRelativeDateTimeControlOptions? DefaultRelativeDateTimeOptions;
        /// <summary>
        /// The default options that correspond to the `Slider` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultSliderControlOptions? DefaultSliderOptions;
        /// <summary>
        /// The default options that correspond to the `TextArea` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultTextAreaControlOptions? DefaultTextAreaOptions;
        /// <summary>
        /// The default options that correspond to the `TextField` filter control type.
        /// </summary>
        public readonly Outputs.DashboardDefaultTextFieldControlOptions? DefaultTextFieldOptions;

        [OutputConstructor]
        private DashboardDefaultFilterControlOptions(
            Outputs.DashboardDefaultDateTimePickerControlOptions? defaultDateTimePickerOptions,

            Outputs.DashboardDefaultFilterDropDownControlOptions? defaultDropdownOptions,

            Outputs.DashboardDefaultFilterListControlOptions? defaultListOptions,

            Outputs.DashboardDefaultRelativeDateTimeControlOptions? defaultRelativeDateTimeOptions,

            Outputs.DashboardDefaultSliderControlOptions? defaultSliderOptions,

            Outputs.DashboardDefaultTextAreaControlOptions? defaultTextAreaOptions,

            Outputs.DashboardDefaultTextFieldControlOptions? defaultTextFieldOptions)
        {
            DefaultDateTimePickerOptions = defaultDateTimePickerOptions;
            DefaultDropdownOptions = defaultDropdownOptions;
            DefaultListOptions = defaultListOptions;
            DefaultRelativeDateTimeOptions = defaultRelativeDateTimeOptions;
            DefaultSliderOptions = defaultSliderOptions;
            DefaultTextAreaOptions = defaultTextAreaOptions;
            DefaultTextFieldOptions = defaultTextFieldOptions;
        }
    }
}
