// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
        [Input("dateTimeParameterDeclaration")]
        public Input<Inputs.TemplateDateTimeParameterDeclarationArgs>? DateTimeParameterDeclaration { get; set; }

        [Input("decimalParameterDeclaration")]
        public Input<Inputs.TemplateDecimalParameterDeclarationArgs>? DecimalParameterDeclaration { get; set; }

        [Input("integerParameterDeclaration")]
        public Input<Inputs.TemplateIntegerParameterDeclarationArgs>? IntegerParameterDeclaration { get; set; }

        [Input("stringParameterDeclaration")]
        public Input<Inputs.TemplateStringParameterDeclarationArgs>? StringParameterDeclaration { get; set; }

        public TemplateParameterDeclarationArgs()
        {
        }
        public static new TemplateParameterDeclarationArgs Empty => new TemplateParameterDeclarationArgs();
    }
}