// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// Contains the channel and queue identifier for a routing profile.
    /// </summary>
    public sealed class RoutingProfileQueueReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.
        /// </summary>
        [Input("channel", required: true)]
        public Input<Pulumi.AwsNative.Connect.RoutingProfileChannel> Channel { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the queue.
        /// </summary>
        [Input("queueArn", required: true)]
        public Input<string> QueueArn { get; set; } = null!;

        public RoutingProfileQueueReferenceArgs()
        {
        }
        public static new RoutingProfileQueueReferenceArgs Empty => new RoutingProfileQueueReferenceArgs();
    }
}
