// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfilePardotConnectorProfileCredentials
    {
        /// <summary>
        /// The credentials used to access protected resources.
        /// </summary>
        public readonly string? AccessToken;
        /// <summary>
        /// The client credentials to fetch access token and refresh token.
        /// </summary>
        public readonly string? ClientCredentialsArn;
        /// <summary>
        /// The oauth needed to request security tokens from the connector endpoint.
        /// </summary>
        public readonly Outputs.ConnectorProfileConnectorOAuthRequest? ConnectorOAuthRequest;
        /// <summary>
        /// The credentials used to acquire new access tokens.
        /// </summary>
        public readonly string? RefreshToken;

        [OutputConstructor]
        private ConnectorProfilePardotConnectorProfileCredentials(
            string? accessToken,

            string? clientCredentialsArn,

            Outputs.ConnectorProfileConnectorOAuthRequest? connectorOAuthRequest,

            string? refreshToken)
        {
            AccessToken = accessToken;
            ClientCredentialsArn = clientCredentialsArn;
            ConnectorOAuthRequest = connectorOAuthRequest;
            RefreshToken = refreshToken;
        }
    }
}