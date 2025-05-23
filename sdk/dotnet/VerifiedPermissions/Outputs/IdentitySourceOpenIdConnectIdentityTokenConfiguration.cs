// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions.Outputs
{

    [OutputType]
    public sealed class IdentitySourceOpenIdConnectIdentityTokenConfiguration
    {
        public readonly ImmutableArray<string> ClientIds;
        public readonly string? PrincipalIdClaim;

        [OutputConstructor]
        private IdentitySourceOpenIdConnectIdentityTokenConfiguration(
            ImmutableArray<string> clientIds,

            string? principalIdClaim)
        {
            ClientIds = clientIds;
            PrincipalIdClaim = principalIdClaim;
        }
    }
}
