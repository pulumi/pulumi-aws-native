// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDateTimeParameterDeclarationArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultValues")]
        public Input<Inputs.AnalysisDateTimeDefaultValuesArgs>? DefaultValues { get; set; }

        [Input("mappedDataSetParameters")]
        private InputList<Inputs.AnalysisMappedDataSetParameterArgs>? _mappedDataSetParameters;
        public InputList<Inputs.AnalysisMappedDataSetParameterArgs> MappedDataSetParameters
        {
            get => _mappedDataSetParameters ?? (_mappedDataSetParameters = new InputList<Inputs.AnalysisMappedDataSetParameterArgs>());
            set => _mappedDataSetParameters = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("timeGranularity")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity>? TimeGranularity { get; set; }

        [Input("valueWhenUnset")]
        public Input<Inputs.AnalysisDateTimeValueWhenUnsetConfigurationArgs>? ValueWhenUnset { get; set; }

        public AnalysisDateTimeParameterDeclarationArgs()
        {
        }
        public static new AnalysisDateTimeParameterDeclarationArgs Empty => new AnalysisDateTimeParameterDeclarationArgs();
    }
}