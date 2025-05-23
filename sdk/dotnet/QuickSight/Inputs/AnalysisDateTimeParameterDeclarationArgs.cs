// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// The default values of a parameter. If the parameter is a single-value parameter, a maximum of one default value can be provided.
        /// </summary>
        [Input("defaultValues")]
        public Input<Inputs.AnalysisDateTimeDefaultValuesArgs>? DefaultValues { get; set; }

        [Input("mappedDataSetParameters")]
        private InputList<Inputs.AnalysisMappedDataSetParameterArgs>? _mappedDataSetParameters;
        public InputList<Inputs.AnalysisMappedDataSetParameterArgs> MappedDataSetParameters
        {
            get => _mappedDataSetParameters ?? (_mappedDataSetParameters = new InputList<Inputs.AnalysisMappedDataSetParameterArgs>());
            set => _mappedDataSetParameters = value;
        }

        /// <summary>
        /// The name of the parameter that is being declared.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The level of time precision that is used to aggregate `DateTime` values.
        /// </summary>
        [Input("timeGranularity")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisTimeGranularity>? TimeGranularity { get; set; }

        /// <summary>
        /// The configuration that defines the default value of a `DateTime` parameter when a value has not been set.
        /// </summary>
        [Input("valueWhenUnset")]
        public Input<Inputs.AnalysisDateTimeValueWhenUnsetConfigurationArgs>? ValueWhenUnset { get; set; }

        public AnalysisDateTimeParameterDeclarationArgs()
        {
        }
        public static new AnalysisDateTimeParameterDeclarationArgs Empty => new AnalysisDateTimeParameterDeclarationArgs();
    }
}
