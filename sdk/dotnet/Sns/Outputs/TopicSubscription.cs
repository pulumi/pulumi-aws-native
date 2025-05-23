// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sns.Outputs
{

    /// <summary>
    /// ``Subscription`` is an embedded property that describes the subscription endpoints of an SNS topic.
    ///   For full control over subscription behavior (for example, delivery policy, filtering, raw message delivery, and cross-region subscriptions), use the [AWS::SNS::Subscription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html) resource.
    /// </summary>
    [OutputType]
    public sealed class TopicSubscription
    {
        /// <summary>
        /// The endpoint that receives notifications from the SNS topic. The endpoint value depends on the protocol that you specify. For more information, see the ``Endpoint`` parameter of the ``Subscribe`` action in the *API Reference*.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The subscription's protocol. For more information, see the ``Protocol`` parameter of the ``Subscribe`` action in the *API Reference*.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private TopicSubscription(
            string endpoint,

            string protocol)
        {
            Endpoint = endpoint;
            Protocol = protocol;
        }
    }
}
