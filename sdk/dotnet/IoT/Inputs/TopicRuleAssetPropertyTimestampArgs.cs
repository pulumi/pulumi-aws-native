// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleAssetPropertyTimestampArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A string that contains the nanosecond time offset. Accepts substitution templates.
        /// </summary>
        [Input("offsetInNanos")]
        public Input<string>? OffsetInNanos { get; set; }

        /// <summary>
        /// A string that contains the time in seconds since epoch. Accepts substitution templates.
        /// </summary>
        [Input("timeInSeconds", required: true)]
        public Input<string> TimeInSeconds { get; set; } = null!;

        public TopicRuleAssetPropertyTimestampArgs()
        {
        }
        public static new TopicRuleAssetPropertyTimestampArgs Empty => new TopicRuleAssetPropertyTimestampArgs();
    }
}
