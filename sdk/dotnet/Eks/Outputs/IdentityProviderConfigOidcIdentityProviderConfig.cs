// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Outputs
{

    /// <summary>
    /// An object representing an OpenID Connect (OIDC) configuration.
    /// </summary>
    [OutputType]
    public sealed class IdentityProviderConfigOidcIdentityProviderConfig
    {
        /// <summary>
        /// This is also known as audience. The ID for the client application that makes authentication requests to the OpenID identity provider.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The JWT claim that the provider uses to return your groups.
        /// </summary>
        public readonly string? GroupsClaim;
        /// <summary>
        /// The prefix that is prepended to group claims to prevent clashes with existing names (such as system: groups).
        /// </summary>
        public readonly string? GroupsPrefix;
        /// <summary>
        /// The URL of the OpenID identity provider that allows the API server to discover public signing keys for verifying tokens.
        /// </summary>
        public readonly string IssuerUrl;
        public readonly ImmutableArray<Outputs.IdentityProviderConfigRequiredClaim> RequiredClaims;
        /// <summary>
        /// The JSON Web Token (JWT) claim to use as the username. The default is sub, which is expected to be a unique identifier of the end user. You can choose other claims, such as email or name, depending on the OpenID identity provider. Claims other than email are prefixed with the issuer URL to prevent naming clashes with other plug-ins.
        /// </summary>
        public readonly string? UsernameClaim;
        /// <summary>
        /// The prefix that is prepended to username claims to prevent clashes with existing names. If you do not provide this field, and username is a value other than email, the prefix defaults to issuerurl#. You can use the value - to disable all prefixing.
        /// </summary>
        public readonly string? UsernamePrefix;

        [OutputConstructor]
        private IdentityProviderConfigOidcIdentityProviderConfig(
            string clientId,

            string? groupsClaim,

            string? groupsPrefix,

            string issuerUrl,

            ImmutableArray<Outputs.IdentityProviderConfigRequiredClaim> requiredClaims,

            string? usernameClaim,

            string? usernamePrefix)
        {
            ClientId = clientId;
            GroupsClaim = groupsClaim;
            GroupsPrefix = groupsPrefix;
            IssuerUrl = issuerUrl;
            RequiredClaims = requiredClaims;
            UsernameClaim = usernameClaim;
            UsernamePrefix = usernamePrefix;
        }
    }
}