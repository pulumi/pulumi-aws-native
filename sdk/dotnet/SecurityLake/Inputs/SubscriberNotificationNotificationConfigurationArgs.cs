// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    public sealed class SubscriberNotificationNotificationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("httpsNotificationConfiguration")]
        public Input<Inputs.SubscriberNotificationHttpsNotificationConfigurationArgs>? HttpsNotificationConfiguration { get; set; }

        [Input("sqsNotificationConfiguration")]
        public Input<Inputs.SubscriberNotificationSqsNotificationConfigurationArgs>? SqsNotificationConfiguration { get; set; }

        public SubscriberNotificationNotificationConfigurationArgs()
        {
        }
        public static new SubscriberNotificationNotificationConfigurationArgs Empty => new SubscriberNotificationNotificationConfigurationArgs();
    }
}