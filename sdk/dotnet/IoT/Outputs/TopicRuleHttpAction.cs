// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleHttpAction
    {
        /// <summary>
        /// The authentication method to use when sending data to an HTTPS endpoint.
        /// </summary>
        public readonly Outputs.TopicRuleHttpAuthorization? Auth;
        /// <summary>
        /// The URL to which AWS IoT sends a confirmation message. The value of the confirmation URL must be a prefix of the endpoint URL. If you do not specify a confirmation URL AWS IoT uses the endpoint URL as the confirmation URL. If you use substitution templates in the confirmationUrl, you must create and enable topic rule destinations that match each possible value of the substitution template before traffic is allowed to your endpoint URL.
        /// </summary>
        public readonly string? ConfirmationUrl;
        /// <summary>
        /// The HTTP headers to send with the message data.
        /// </summary>
        public readonly ImmutableArray<Outputs.TopicRuleHttpActionHeader> Headers;
        /// <summary>
        /// The endpoint URL. If substitution templates are used in the URL, you must also specify a `confirmationUrl` . If this is a new destination, a new `TopicRuleDestination` is created if possible.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private TopicRuleHttpAction(
            Outputs.TopicRuleHttpAuthorization? auth,

            string? confirmationUrl,

            ImmutableArray<Outputs.TopicRuleHttpActionHeader> headers,

            string url)
        {
            Auth = auth;
            ConfirmationUrl = confirmationUrl;
            Headers = headers;
            Url = url;
        }
    }
}
