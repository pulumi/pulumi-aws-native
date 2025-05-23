// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Inputs
{

    /// <summary>
    /// The resource that receives events arriving from a channel.
    /// </summary>
    public sealed class ChannelDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of a resource that receives events from a channel.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The type of destination for events arriving from a channel.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.CloudTrail.ChannelDestinationType> Type { get; set; } = null!;

        public ChannelDestinationArgs()
        {
        }
        public static new ChannelDestinationArgs Empty => new ChannelDestinationArgs();
    }
}
