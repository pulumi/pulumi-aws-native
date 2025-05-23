// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisMappedDataSetParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name that identifies a dataset within the analysis or dashboard.
        /// </summary>
        [Input("dataSetIdentifier", required: true)]
        public Input<string> DataSetIdentifier { get; set; } = null!;

        /// <summary>
        /// The name of the dataset parameter.
        /// </summary>
        [Input("dataSetParameterName", required: true)]
        public Input<string> DataSetParameterName { get; set; } = null!;

        public AnalysisMappedDataSetParameterArgs()
        {
        }
        public static new AnalysisMappedDataSetParameterArgs Empty => new AnalysisMappedDataSetParameterArgs();
    }
}
