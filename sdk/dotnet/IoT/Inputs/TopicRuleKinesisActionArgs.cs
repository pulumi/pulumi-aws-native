// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleKinesisActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The partition key.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        /// <summary>
        /// The ARN of the IAM role that grants access to the Amazon Kinesis stream.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// The name of the Amazon Kinesis stream.
        /// </summary>
        [Input("streamName", required: true)]
        public Input<string> StreamName { get; set; } = null!;

        public TopicRuleKinesisActionArgs()
        {
        }
        public static new TopicRuleKinesisActionArgs Empty => new TopicRuleKinesisActionArgs();
    }
}
