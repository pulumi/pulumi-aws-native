// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// An object that contains the Amazon Resource Name (ARN) of the Amazon Lambda function that is used to preprocess records in the stream in a SQL-based Kinesis Data Analytics application.
    /// </summary>
    public sealed class ApplicationInputLambdaProcessorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Amazon Lambda function that operates on records in the stream.
        /// </summary>
        [Input("resourceARN", required: true)]
        public Input<string> ResourceARN { get; set; } = null!;

        public ApplicationInputLambdaProcessorArgs()
        {
        }
        public static new ApplicationInputLambdaProcessorArgs Empty => new ApplicationInputLambdaProcessorArgs();
    }
}