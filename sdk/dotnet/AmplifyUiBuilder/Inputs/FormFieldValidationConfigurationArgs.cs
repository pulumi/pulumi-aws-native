// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormFieldValidationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("numValues")]
        private InputList<double>? _numValues;

        /// <summary>
        /// The validation to perform on a number value.
        /// </summary>
        public InputList<double> NumValues
        {
            get => _numValues ?? (_numValues = new InputList<double>());
            set => _numValues = value;
        }

        [Input("strValues")]
        private InputList<string>? _strValues;

        /// <summary>
        /// The validation to perform on a string value.
        /// </summary>
        public InputList<string> StrValues
        {
            get => _strValues ?? (_strValues = new InputList<string>());
            set => _strValues = value;
        }

        /// <summary>
        /// The validation to perform on an object type. ``
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The validation message to display.
        /// </summary>
        [Input("validationMessage")]
        public Input<string>? ValidationMessage { get; set; }

        public FormFieldValidationConfigurationArgs()
        {
        }
        public static new FormFieldValidationConfigurationArgs Empty => new FormFieldValidationConfigurationArgs();
    }
}