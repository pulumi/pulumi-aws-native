// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeGuruProfiler.Outputs
{

    /// <summary>
    /// Notification medium for users to get alerted for events that occur in application profile. We support SNS topic as a notification channel.
    /// </summary>
    [OutputType]
    public sealed class ProfilingGroupChannel
    {
        public readonly string? ChannelId;
        public readonly string ChannelUri;

        [OutputConstructor]
        private ProfilingGroupChannel(
            string? channelId,

            string channelUri)
        {
            ChannelId = channelId;
            ChannelUri = channelUri;
        }
    }
}