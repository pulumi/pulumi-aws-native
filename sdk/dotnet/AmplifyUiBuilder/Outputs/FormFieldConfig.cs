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
    public sealed class FormFieldConfig
    {
        public readonly bool? Excluded;
        public readonly Outputs.FormFieldInputConfig? InputType;
        public readonly string? Label;
        public readonly object? Position;
        public readonly ImmutableArray<Outputs.FormFieldValidationConfiguration> Validations;

        [OutputConstructor]
        private FormFieldConfig(
            bool? excluded,

            Outputs.FormFieldInputConfig? inputType,

            string? label,

            object? position,

            ImmutableArray<Outputs.FormFieldValidationConfiguration> validations)
        {
            Excluded = excluded;
            InputType = inputType;
            Label = label;
            Position = position;
            Validations = validations;
        }
    }
}