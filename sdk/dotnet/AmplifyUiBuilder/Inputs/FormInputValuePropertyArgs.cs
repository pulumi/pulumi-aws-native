// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormInputValuePropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The information to bind fields to data at runtime.
        /// </summary>
        [Input("bindingProperties")]
        public Input<Inputs.FormInputValuePropertyBindingPropertiesArgs>? BindingProperties { get; set; }

        [Input("concat")]
        private InputList<Inputs.FormInputValuePropertyArgs>? _concat;

        /// <summary>
        /// A list of form properties to concatenate to create the value to assign to this field property.
        /// </summary>
        public InputList<Inputs.FormInputValuePropertyArgs> Concat
        {
            get => _concat ?? (_concat = new InputList<Inputs.FormInputValuePropertyArgs>());
            set => _concat = value;
        }

        /// <summary>
        /// The value to assign to the input field.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public FormInputValuePropertyArgs()
        {
        }
        public static new FormInputValuePropertyArgs Empty => new FormInputValuePropertyArgs();
    }
}
