// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleFirehoseActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("batchMode")]
        public Input<bool>? BatchMode { get; set; }

        [Input("deliveryStreamName", required: true)]
        public Input<string> DeliveryStreamName { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public TopicRuleFirehoseActionArgs()
        {
        }
        public static new TopicRuleFirehoseActionArgs Empty => new TopicRuleFirehoseActionArgs();
    }
}