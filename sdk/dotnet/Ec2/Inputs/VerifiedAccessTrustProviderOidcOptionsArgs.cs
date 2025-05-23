// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The OpenID Connect details for an oidc -type, user-identity based trust provider.
    /// </summary>
    public sealed class VerifiedAccessTrustProviderOidcOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OIDC authorization endpoint.
        /// </summary>
        [Input("authorizationEndpoint")]
        public Input<string>? AuthorizationEndpoint { get; set; }

        /// <summary>
        /// The client identifier.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The OIDC issuer.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The OIDC token endpoint.
        /// </summary>
        [Input("tokenEndpoint")]
        public Input<string>? TokenEndpoint { get; set; }

        /// <summary>
        /// The OIDC user info endpoint.
        /// </summary>
        [Input("userInfoEndpoint")]
        public Input<string>? UserInfoEndpoint { get; set; }

        public VerifiedAccessTrustProviderOidcOptionsArgs()
        {
        }
        public static new VerifiedAccessTrustProviderOidcOptionsArgs Empty => new VerifiedAccessTrustProviderOidcOptionsArgs();
    }
}
