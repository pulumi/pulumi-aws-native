// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    /// <summary>
    /// OAuth2 Properties
    /// </summary>
    [OutputType]
    public sealed class ConnectionOAuth2Properties
    {
        public readonly Outputs.ConnectionAuthorizationCodeProperties? AuthorizationCodeProperties;
        public readonly Outputs.ConnectionOAuth2ClientApplication? OAuth2ClientApplication;
        public readonly Outputs.ConnectionGlueOAuth2Credentials? OAuth2Credentials;
        public readonly Pulumi.AwsNative.DataZone.ConnectionOAuth2GrantType? OAuth2GrantType;
        public readonly string? TokenUrl;
        public readonly ImmutableDictionary<string, string>? TokenUrlParametersMap;

        [OutputConstructor]
        private ConnectionOAuth2Properties(
            Outputs.ConnectionAuthorizationCodeProperties? authorizationCodeProperties,

            Outputs.ConnectionOAuth2ClientApplication? oAuth2ClientApplication,

            Outputs.ConnectionGlueOAuth2Credentials? oAuth2Credentials,

            Pulumi.AwsNative.DataZone.ConnectionOAuth2GrantType? oAuth2GrantType,

            string? tokenUrl,

            ImmutableDictionary<string, string>? tokenUrlParametersMap)
        {
            AuthorizationCodeProperties = authorizationCodeProperties;
            OAuth2ClientApplication = oAuth2ClientApplication;
            OAuth2Credentials = oAuth2Credentials;
            OAuth2GrantType = oAuth2GrantType;
            TokenUrl = tokenUrl;
            TokenUrlParametersMap = tokenUrlParametersMap;
        }
    }
}
