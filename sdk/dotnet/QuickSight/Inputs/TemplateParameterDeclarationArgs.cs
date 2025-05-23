// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateParameterDeclarationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A parameter declaration for the `DateTime` data type.
        /// </summary>
        [Input("dateTimeParameterDeclaration")]
        public Input<Inputs.TemplateDateTimeParameterDeclarationArgs>? DateTimeParameterDeclaration { get; set; }

        /// <summary>
        /// A parameter declaration for the `Decimal` data type.
        /// </summary>
        [Input("decimalParameterDeclaration")]
        public Input<Inputs.TemplateDecimalParameterDeclarationArgs>? DecimalParameterDeclaration { get; set; }

        /// <summary>
        /// A parameter declaration for the `Integer` data type.
        /// </summary>
        [Input("integerParameterDeclaration")]
        public Input<Inputs.TemplateIntegerParameterDeclarationArgs>? IntegerParameterDeclaration { get; set; }

        /// <summary>
        /// A parameter declaration for the `String` data type.
        /// </summary>
        [Input("stringParameterDeclaration")]
        public Input<Inputs.TemplateStringParameterDeclarationArgs>? StringParameterDeclaration { get; set; }

        public TemplateParameterDeclarationArgs()
        {
        }
        public static new TemplateParameterDeclarationArgs Empty => new TemplateParameterDeclarationArgs();
    }
}
