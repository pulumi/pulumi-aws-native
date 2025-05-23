// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rekognition.Inputs
{

    /// <summary>
    /// The Amazon Kinesis Data Stream stream to which the Amazon Rekognition stream processor streams the analysis results, as part of face search feature.
    /// </summary>
    public sealed class StreamProcessorKinesisDataStreamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the Kinesis Data Stream stream.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public StreamProcessorKinesisDataStreamArgs()
        {
        }
        public static new StreamProcessorKinesisDataStreamArgs Empty => new StreamProcessorKinesisDataStreamArgs();
    }
}
