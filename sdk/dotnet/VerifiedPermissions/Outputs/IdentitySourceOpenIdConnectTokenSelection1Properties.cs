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
    public sealed class IdentitySourceOpenIdConnectTokenSelection1Properties
    {
        public readonly Outputs.IdentitySourceOpenIdConnectIdentityTokenConfiguration IdentityTokenOnly;

        [OutputConstructor]
        private IdentitySourceOpenIdConnectTokenSelection1Properties(Outputs.IdentitySourceOpenIdConnectIdentityTokenConfiguration identityTokenOnly)
        {
            IdentityTokenOnly = identityTokenOnly;
        }
    }
}
