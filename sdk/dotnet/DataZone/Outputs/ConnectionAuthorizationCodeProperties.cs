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
    /// Authorization Code Properties
    /// </summary>
    [OutputType]
    public sealed class ConnectionAuthorizationCodeProperties
    {
        public readonly string? AuthorizationCode;
        public readonly string? RedirectUri;

        [OutputConstructor]
        private ConnectionAuthorizationCodeProperties(
            string? authorizationCode,

            string? redirectUri)
        {
            AuthorizationCode = authorizationCode;
            RedirectUri = redirectUri;
        }
    }
}
