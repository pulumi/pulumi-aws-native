// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES.Inputs
{

    public sealed class ReceiptRuleLambdaActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("functionArn", required: true)]
        public Input<string> FunctionArn { get; set; } = null!;

        [Input("invocationType")]
        public Input<string>? InvocationType { get; set; }

        [Input("topicArn")]
        public Input<string>? TopicArn { get; set; }

        public ReceiptRuleLambdaActionArgs()
        {
        }
        public static new ReceiptRuleLambdaActionArgs Empty => new ReceiptRuleLambdaActionArgs();
    }
}