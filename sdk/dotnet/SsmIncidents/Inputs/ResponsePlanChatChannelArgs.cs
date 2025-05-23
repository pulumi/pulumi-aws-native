// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Inputs
{

    /// <summary>
    /// The chat channel configuration.
    /// </summary>
    public sealed class ResponsePlanChatChannelArgs : global::Pulumi.ResourceArgs
    {
        [Input("chatbotSns")]
        private InputList<string>? _chatbotSns;

        /// <summary>
        /// The Amazon SNS targets that  uses to notify the chat channel of updates to an incident. You can also make updates to the incident through the chat channel by using the Amazon SNS topics
        /// </summary>
        public InputList<string> ChatbotSns
        {
            get => _chatbotSns ?? (_chatbotSns = new InputList<string>());
            set => _chatbotSns = value;
        }

        public ResponsePlanChatChannelArgs()
        {
        }
        public static new ResponsePlanChatChannelArgs Empty => new ResponsePlanChatChannelArgs();
    }
}
