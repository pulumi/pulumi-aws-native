// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormFieldInputConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether a field has a default value.
        /// </summary>
        [Input("defaultChecked")]
        public Input<bool>? DefaultChecked { get; set; }

        /// <summary>
        /// The default country code for a phone number.
        /// </summary>
        [Input("defaultCountryCode")]
        public Input<string>? DefaultCountryCode { get; set; }

        /// <summary>
        /// The default value for the field.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// The text to display to describe the field.
        /// </summary>
        [Input("descriptiveText")]
        public Input<string>? DescriptiveText { get; set; }

        /// <summary>
        /// The configuration for the file uploader field.
        /// </summary>
        [Input("fileUploaderConfig")]
        public Input<Inputs.FormFileUploaderFieldConfigArgs>? FileUploaderConfig { get; set; }

        /// <summary>
        /// Specifies whether to render the field as an array. This property is ignored if the `dataSourceType` for the form is a Data Store.
        /// </summary>
        [Input("isArray")]
        public Input<bool>? IsArray { get; set; }

        /// <summary>
        /// The maximum value to display for the field.
        /// </summary>
        [Input("maxValue")]
        public Input<double>? MaxValue { get; set; }

        /// <summary>
        /// The minimum value to display for the field.
        /// </summary>
        [Input("minValue")]
        public Input<double>? MinValue { get; set; }

        /// <summary>
        /// The name of the field.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The text to display as a placeholder for the field.
        /// </summary>
        [Input("placeholder")]
        public Input<string>? Placeholder { get; set; }

        /// <summary>
        /// Specifies a read only field.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Specifies a field that requires input.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The stepping increment for a numeric value in a field.
        /// </summary>
        [Input("step")]
        public Input<double>? Step { get; set; }

        /// <summary>
        /// The input type for the field.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value for the field.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The information to use to customize the input fields with data at runtime.
        /// </summary>
        [Input("valueMappings")]
        public Input<Inputs.FormValueMappingsArgs>? ValueMappings { get; set; }

        public FormFieldInputConfigArgs()
        {
        }
        public static new FormFieldInputConfigArgs Empty => new FormFieldInputConfigArgs();
    }
}
