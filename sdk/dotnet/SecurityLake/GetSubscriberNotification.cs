// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake
{
    public static class GetSubscriberNotification
    {
        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::SubscriberNotification
        /// </summary>
        public static Task<GetSubscriberNotificationResult> InvokeAsync(GetSubscriberNotificationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriberNotificationResult>("aws-native:securitylake:getSubscriberNotification", args ?? new GetSubscriberNotificationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SecurityLake::SubscriberNotification
        /// </summary>
        public static Output<GetSubscriberNotificationResult> Invoke(GetSubscriberNotificationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriberNotificationResult>("aws-native:securitylake:getSubscriberNotification", args ?? new GetSubscriberNotificationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriberNotificationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN for the subscriber
        /// </summary>
        [Input("subscriberArn", required: true)]
        public string SubscriberArn { get; set; } = null!;

        public GetSubscriberNotificationArgs()
        {
        }
        public static new GetSubscriberNotificationArgs Empty => new GetSubscriberNotificationArgs();
    }

    public sealed class GetSubscriberNotificationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN for the subscriber
        /// </summary>
        [Input("subscriberArn", required: true)]
        public Input<string> SubscriberArn { get; set; } = null!;

        public GetSubscriberNotificationInvokeArgs()
        {
        }
        public static new GetSubscriberNotificationInvokeArgs Empty => new GetSubscriberNotificationInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriberNotificationResult
    {
        /// <summary>
        /// Specify the configurations you want to use for subscriber notification. The subscriber is notified when new data is written to the data lake for sources that the subscriber consumes in Security Lake .
        /// </summary>
        public readonly Outputs.SubscriberNotificationNotificationConfiguration? NotificationConfiguration;
        /// <summary>
        /// The endpoint the subscriber should listen to for notifications
        /// </summary>
        public readonly string? SubscriberEndpoint;

        [OutputConstructor]
        private GetSubscriberNotificationResult(
            Outputs.SubscriberNotificationNotificationConfiguration? notificationConfiguration,

            string? subscriberEndpoint)
        {
            NotificationConfiguration = notificationConfiguration;
            SubscriberEndpoint = subscriberEndpoint;
        }
    }
}