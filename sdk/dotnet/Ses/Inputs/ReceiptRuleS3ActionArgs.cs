// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class ReceiptRuleS3ActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("objectKeyPrefix")]
        public Input<string>? ObjectKeyPrefix { get; set; }

        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleS3ActionArgs()
        {
        }
        public static new ReceiptRuleS3ActionArgs Empty => new ReceiptRuleS3ActionArgs();
    }
}