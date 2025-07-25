// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AiOps.Inputs
{

    public sealed class InvestigationGroupChatbotNotificationChannelArgs : global::Pulumi.ResourceArgs
    {
        [Input("chatConfigurationArns")]
        private InputList<string>? _chatConfigurationArns;

        /// <summary>
        /// Returns the Amazon Resource Name (ARN) of any third-party chat integrations configured for the account.
        /// </summary>
        public InputList<string> ChatConfigurationArns
        {
            get => _chatConfigurationArns ?? (_chatConfigurationArns = new InputList<string>());
            set => _chatConfigurationArns = value;
        }

        /// <summary>
        /// Returns the ARN of an Amazon SNS topic used for third-party chat integrations.
        /// </summary>
        [Input("snsTopicArn")]
        public Input<string>? SnsTopicArn { get; set; }

        public InvestigationGroupChatbotNotificationChannelArgs()
        {
        }
        public static new InvestigationGroupChatbotNotificationChannelArgs Empty => new InvestigationGroupChatbotNotificationChannelArgs();
    }
}
