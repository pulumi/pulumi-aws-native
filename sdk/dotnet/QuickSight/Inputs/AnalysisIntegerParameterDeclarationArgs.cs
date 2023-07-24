// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisIntegerParameterDeclarationArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultValues")]
        public Input<Inputs.AnalysisIntegerDefaultValuesArgs>? DefaultValues { get; set; }

        [Input("mappedDataSetParameters")]
        private InputList<Inputs.AnalysisMappedDataSetParameterArgs>? _mappedDataSetParameters;
        public InputList<Inputs.AnalysisMappedDataSetParameterArgs> MappedDataSetParameters
        {
            get => _mappedDataSetParameters ?? (_mappedDataSetParameters = new InputList<Inputs.AnalysisMappedDataSetParameterArgs>());
            set => _mappedDataSetParameters = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("parameterValueType", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisParameterValueType> ParameterValueType { get; set; } = null!;

        [Input("valueWhenUnset")]
        public Input<Inputs.AnalysisIntegerValueWhenUnsetConfigurationArgs>? ValueWhenUnset { get; set; }

        public AnalysisIntegerParameterDeclarationArgs()
        {
        }
        public static new AnalysisIntegerParameterDeclarationArgs Empty => new AnalysisIntegerParameterDeclarationArgs();
    }
}