// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Outputs
{

    /// <summary>
    /// The chat channel configuration.
    /// </summary>
    [OutputType]
    public sealed class ResponsePlanChatChannel
    {
        public readonly ImmutableArray<string> ChatbotSns;

        [OutputConstructor]
        private ResponsePlanChatChannel(ImmutableArray<string> chatbotSns)
        {
            ChatbotSns = chatbotSns;
        }
    }
}