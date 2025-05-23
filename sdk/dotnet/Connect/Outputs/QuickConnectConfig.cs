// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    /// <summary>
    /// Configuration settings for the quick connect.
    /// </summary>
    [OutputType]
    public sealed class QuickConnectConfig
    {
        /// <summary>
        /// The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.
        /// </summary>
        public readonly Outputs.QuickConnectPhoneNumberQuickConnectConfig? PhoneConfig;
        /// <summary>
        /// The queue configuration. This is required only if QuickConnectType is QUEUE.
        /// </summary>
        public readonly Outputs.QuickConnectQueueQuickConnectConfig? QueueConfig;
        /// <summary>
        /// The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).
        /// </summary>
        public readonly Pulumi.AwsNative.Connect.QuickConnectType QuickConnectType;
        /// <summary>
        /// The user configuration. This is required only if QuickConnectType is USER.
        /// </summary>
        public readonly Outputs.QuickConnectUserQuickConnectConfig? UserConfig;

        [OutputConstructor]
        private QuickConnectConfig(
            Outputs.QuickConnectPhoneNumberQuickConnectConfig? phoneConfig,

            Outputs.QuickConnectQueueQuickConnectConfig? queueConfig,

            Pulumi.AwsNative.Connect.QuickConnectType quickConnectType,

            Outputs.QuickConnectUserQuickConnectConfig? userConfig)
        {
            PhoneConfig = phoneConfig;
            QueueConfig = queueConfig;
            QuickConnectType = quickConnectType;
            UserConfig = userConfig;
        }
    }
}
