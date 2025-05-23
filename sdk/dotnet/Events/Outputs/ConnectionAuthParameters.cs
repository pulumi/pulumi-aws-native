// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class ConnectionAuthParameters
    {
        /// <summary>
        /// The API Key parameters to use for authorization.
        /// </summary>
        public readonly Outputs.ConnectionApiKeyAuthParameters? ApiKeyAuthParameters;
        /// <summary>
        /// The authorization parameters for Basic authorization.
        /// </summary>
        public readonly Outputs.ConnectionBasicAuthParameters? BasicAuthParameters;
        /// <summary>
        /// For private OAuth authentication endpoints. The parameters EventBridge uses to authenticate against the endpoint.
        /// 
        /// For more information, see [Authorization methods for connections](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-target-connection-auth.html) in the **Amazon EventBridge User Guide** .
        /// </summary>
        public readonly Outputs.ConnectionConnectivityParameters? ConnectivityParameters;
        /// <summary>
        /// Additional parameters for the connection that are passed through with every invocation to the HTTP endpoint.
        /// </summary>
        public readonly Outputs.ConnectionHttpParameters? InvocationHttpParameters;
        /// <summary>
        /// The OAuth parameters to use for authorization.
        /// </summary>
        public readonly Outputs.ConnectionOAuthParameters? OAuthParameters;

        [OutputConstructor]
        private ConnectionAuthParameters(
            Outputs.ConnectionApiKeyAuthParameters? apiKeyAuthParameters,

            Outputs.ConnectionBasicAuthParameters? basicAuthParameters,

            Outputs.ConnectionConnectivityParameters? connectivityParameters,

            Outputs.ConnectionHttpParameters? invocationHttpParameters,

            Outputs.ConnectionOAuthParameters? oAuthParameters)
        {
            ApiKeyAuthParameters = apiKeyAuthParameters;
            BasicAuthParameters = basicAuthParameters;
            ConnectivityParameters = connectivityParameters;
            InvocationHttpParameters = invocationHttpParameters;
            OAuthParameters = oAuthParameters;
        }
    }
}
