// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class PipelineChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the channel from which the messages are processed.
        /// </summary>
        [Input("channelName", required: true)]
        public Input<string> ChannelName { get; set; } = null!;

        /// <summary>
        /// The name of the 'channel' activity.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The next activity in the pipeline.
        /// </summary>
        [Input("next")]
        public Input<string>? Next { get; set; }

        public PipelineChannelArgs()
        {
        }
        public static new PipelineChannelArgs Empty => new PipelineChannelArgs();
    }
}
