// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Outputs
{

    [OutputType]
    public sealed class FormFieldValidationConfiguration
    {
        /// <summary>
        /// The validation to perform on a number value.
        /// </summary>
        public readonly ImmutableArray<double> NumValues;
        /// <summary>
        /// The validation to perform on a string value.
        /// </summary>
        public readonly ImmutableArray<string> StrValues;
        /// <summary>
        /// The validation to perform on an object type. ``
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The validation message to display.
        /// </summary>
        public readonly string? ValidationMessage;

        [OutputConstructor]
        private FormFieldValidationConfiguration(
            ImmutableArray<double> numValues,

            ImmutableArray<string> strValues,

            string type,

            string? validationMessage)
        {
            NumValues = numValues;
            StrValues = strValues;
            Type = type;
            ValidationMessage = validationMessage;
        }
    }
}