// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class ConnectorProfileOAuthPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization code url required to redirect to SAP Login Page to fetch authorization code for OAuth type authentication.
        /// </summary>
        [Input("authCodeUrl")]
        public Input<string>? AuthCodeUrl { get; set; }

        [Input("oAuthScopes")]
        private InputList<string>? _oAuthScopes;

        /// <summary>
        /// The OAuth scopes required for OAuth type authentication.
        /// </summary>
        public InputList<string> OAuthScopes
        {
            get => _oAuthScopes ?? (_oAuthScopes = new InputList<string>());
            set => _oAuthScopes = value;
        }

        /// <summary>
        /// The token url required to fetch access/refresh tokens using authorization code and also to refresh expired access token using refresh token.
        /// </summary>
        [Input("tokenUrl")]
        public Input<string>? TokenUrl { get; set; }

        public ConnectorProfileOAuthPropertiesArgs()
        {
        }
        public static new ConnectorProfileOAuthPropertiesArgs Empty => new ConnectorProfileOAuthPropertiesArgs();
    }
}
