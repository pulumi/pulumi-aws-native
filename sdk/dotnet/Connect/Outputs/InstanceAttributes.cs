// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    [OutputType]
    public sealed class InstanceAttributes
    {
        public readonly bool? AutoResolveBestVoices;
        public readonly bool? ContactLens;
        public readonly bool? ContactflowLogs;
        public readonly bool? EarlyMedia;
        public readonly bool? EnhancedChatMonitoring;
        public readonly bool? EnhancedContactMonitoring;
        public readonly bool? HighVolumeOutBound;
        public readonly bool InboundCalls;
        public readonly bool? MultiPartyChatConference;
        public readonly bool? MultiPartyConference;
        public readonly bool OutboundCalls;
        public readonly bool? UseCustomTtsVoices;

        [OutputConstructor]
        private InstanceAttributes(
            bool? autoResolveBestVoices,

            bool? contactLens,

            bool? contactflowLogs,

            bool? earlyMedia,

            bool? enhancedChatMonitoring,

            bool? enhancedContactMonitoring,

            bool? highVolumeOutBound,

            bool inboundCalls,

            bool? multiPartyChatConference,

            bool? multiPartyConference,

            bool outboundCalls,

            bool? useCustomTtsVoices)
        {
            AutoResolveBestVoices = autoResolveBestVoices;
            ContactLens = contactLens;
            ContactflowLogs = contactflowLogs;
            EarlyMedia = earlyMedia;
            EnhancedChatMonitoring = enhancedChatMonitoring;
            EnhancedContactMonitoring = enhancedContactMonitoring;
            HighVolumeOutBound = highVolumeOutBound;
            InboundCalls = inboundCalls;
            MultiPartyChatConference = multiPartyChatConference;
            MultiPartyConference = multiPartyConference;
            OutboundCalls = outboundCalls;
            UseCustomTtsVoices = useCustomTtsVoices;
        }
    }
}
