// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    /// <summary>
    /// The configuration for HTTPS subscriber notification.
    /// </summary>
    public sealed class SubscriberNotificationHttpsNotificationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key name for the notification subscription.
        /// </summary>
        [Input("authorizationApiKeyName")]
        public Input<string>? AuthorizationApiKeyName { get; set; }

        /// <summary>
        /// The key value for the notification subscription.
        /// </summary>
        [Input("authorizationApiKeyValue")]
        public Input<string>? AuthorizationApiKeyValue { get; set; }

        /// <summary>
        /// The subscription endpoint in Security Lake.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// The HTTPS method used for the notification subscription.
        /// </summary>
        [Input("httpMethod")]
        public Input<Pulumi.AwsNative.SecurityLake.SubscriberNotificationHttpsNotificationConfigurationHttpMethod>? HttpMethod { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.
        /// </summary>
        [Input("targetRoleArn", required: true)]
        public Input<string> TargetRoleArn { get; set; } = null!;

        public SubscriberNotificationHttpsNotificationConfigurationArgs()
        {
        }
        public static new SubscriberNotificationHttpsNotificationConfigurationArgs Empty => new SubscriberNotificationHttpsNotificationConfigurationArgs();
    }
}
