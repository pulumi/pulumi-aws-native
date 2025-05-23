// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormFieldConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to hide a field.
        /// </summary>
        [Input("excluded")]
        public Input<bool>? Excluded { get; set; }

        /// <summary>
        /// Describes the configuration for the default input value to display for a field.
        /// </summary>
        [Input("inputType")]
        public Input<Inputs.FormFieldInputConfigArgs>? InputType { get; set; }

        /// <summary>
        /// The label for the field.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Specifies the field position.
        /// </summary>
        [Input("position")]
        public object? Position { get; set; }

        [Input("validations")]
        private InputList<Inputs.FormFieldValidationConfigurationArgs>? _validations;

        /// <summary>
        /// The validations to perform on the value in the field.
        /// </summary>
        public InputList<Inputs.FormFieldValidationConfigurationArgs> Validations
        {
            get => _validations ?? (_validations = new InputList<Inputs.FormFieldValidationConfigurationArgs>());
            set => _validations = value;
        }

        public FormFieldConfigArgs()
        {
        }
        public static new FormFieldConfigArgs Empty => new FormFieldConfigArgs();
    }
}
