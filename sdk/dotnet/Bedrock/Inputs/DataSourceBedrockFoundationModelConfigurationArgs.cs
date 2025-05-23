// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Settings for a foundation model used to parse documents for a data source.
    /// </summary>
    public sealed class DataSourceBedrockFoundationModelConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the foundation model to use for parsing.
        /// </summary>
        [Input("modelArn", required: true)]
        public Input<string> ModelArn { get; set; } = null!;

        /// <summary>
        /// Specifies whether to enable parsing of multimodal data, including both text and/or images.
        /// </summary>
        [Input("parsingModality")]
        public Input<Pulumi.AwsNative.Bedrock.DataSourceParsingModality>? ParsingModality { get; set; }

        /// <summary>
        /// Instructions for interpreting the contents of a document.
        /// </summary>
        [Input("parsingPrompt")]
        public Input<Inputs.DataSourceParsingPromptArgs>? ParsingPrompt { get; set; }

        public DataSourceBedrockFoundationModelConfigurationArgs()
        {
        }
        public static new DataSourceBedrockFoundationModelConfigurationArgs Empty => new DataSourceBedrockFoundationModelConfigurationArgs();
    }
}
