// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDecimalParameterDeclarationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default values of a parameter. If the parameter is a single-value parameter, a maximum of one default value can be provided.
        /// </summary>
        [Input("defaultValues")]
        public Input<Inputs.DashboardDecimalDefaultValuesArgs>? DefaultValues { get; set; }

        [Input("mappedDataSetParameters")]
        private InputList<Inputs.DashboardMappedDataSetParameterArgs>? _mappedDataSetParameters;
        public InputList<Inputs.DashboardMappedDataSetParameterArgs> MappedDataSetParameters
        {
            get => _mappedDataSetParameters ?? (_mappedDataSetParameters = new InputList<Inputs.DashboardMappedDataSetParameterArgs>());
            set => _mappedDataSetParameters = value;
        }

        /// <summary>
        /// The name of the parameter that is being declared.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value type determines whether the parameter is a single-value or multi-value parameter.
        /// </summary>
        [Input("parameterValueType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.DashboardParameterValueType> ParameterValueType { get; set; } = null!;

        /// <summary>
        /// The configuration that defines the default value of a `Decimal` parameter when a value has not been set.
        /// </summary>
        [Input("valueWhenUnset")]
        public Input<Inputs.DashboardDecimalValueWhenUnsetConfigurationArgs>? ValueWhenUnset { get; set; }

        public DashboardDecimalParameterDeclarationArgs()
        {
        }
        public static new DashboardDecimalParameterDeclarationArgs Empty => new DashboardDecimalParameterDeclarationArgs();
    }
}
