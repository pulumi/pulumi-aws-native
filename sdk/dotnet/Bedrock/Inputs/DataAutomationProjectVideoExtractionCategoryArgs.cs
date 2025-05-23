// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    public sealed class DataAutomationProjectVideoExtractionCategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether generating categorical data from video is enabled.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.DataAutomationProjectState> State { get; set; } = null!;

        [Input("types")]
        private InputList<Pulumi.AwsNative.Bedrock.DataAutomationProjectVideoExtractionCategoryType>? _types;

        /// <summary>
        /// The types of data to generate.
        /// </summary>
        public InputList<Pulumi.AwsNative.Bedrock.DataAutomationProjectVideoExtractionCategoryType> Types
        {
            get => _types ?? (_types = new InputList<Pulumi.AwsNative.Bedrock.DataAutomationProjectVideoExtractionCategoryType>());
            set => _types = value;
        }

        public DataAutomationProjectVideoExtractionCategoryArgs()
        {
        }
        public static new DataAutomationProjectVideoExtractionCategoryArgs Empty => new DataAutomationProjectVideoExtractionCategoryArgs();
    }
}
