// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDecimalDefaultValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("dynamicValue")]
        public Input<Inputs.TemplateDynamicDefaultValueArgs>? DynamicValue { get; set; }

        [Input("staticValues")]
        private InputList<double>? _staticValues;
        public InputList<double> StaticValues
        {
            get => _staticValues ?? (_staticValues = new InputList<double>());
            set => _staticValues = value;
        }

        public TemplateDecimalDefaultValuesArgs()
        {
        }
        public static new TemplateDecimalDefaultValuesArgs Empty => new TemplateDecimalDefaultValuesArgs();
    }
}