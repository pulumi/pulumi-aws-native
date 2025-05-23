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
    public sealed class AnalysisParameterControl
    {
        /// <summary>
        /// A control from a date parameter that specifies date and time.
        /// </summary>
        public readonly Outputs.AnalysisParameterDateTimePickerControl? DateTimePicker;
        /// <summary>
        /// A control to display a dropdown list with buttons that are used to select a single value.
        /// </summary>
        public readonly Outputs.AnalysisParameterDropDownControl? Dropdown;
        /// <summary>
        /// A control to display a list with buttons or boxes that are used to select either a single value or multiple values.
        /// </summary>
        public readonly Outputs.AnalysisParameterListControl? List;
        /// <summary>
        /// A control to display a horizontal toggle bar. This is used to change a value by sliding the toggle.
        /// </summary>
        public readonly Outputs.AnalysisParameterSliderControl? Slider;
        /// <summary>
        /// A control to display a text box that is used to enter multiple entries.
        /// </summary>
        public readonly Outputs.AnalysisParameterTextAreaControl? TextArea;
        /// <summary>
        /// A control to display a text box that is used to enter a single entry.
        /// </summary>
        public readonly Outputs.AnalysisParameterTextFieldControl? TextField;

        [OutputConstructor]
        private AnalysisParameterControl(
            Outputs.AnalysisParameterDateTimePickerControl? dateTimePicker,

            Outputs.AnalysisParameterDropDownControl? dropdown,

            Outputs.AnalysisParameterListControl? list,

            Outputs.AnalysisParameterSliderControl? slider,

            Outputs.AnalysisParameterTextAreaControl? textArea,

            Outputs.AnalysisParameterTextFieldControl? textField)
        {
            DateTimePicker = dateTimePicker;
            Dropdown = dropdown;
            List = list;
            Slider = slider;
            TextArea = textArea;
            TextField = textField;
        }
    }
}
