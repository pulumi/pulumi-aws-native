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
    /// The options for device identity based trust providers.
    /// </summary>
    public sealed class VerifiedAccessTrustProviderDeviceOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL Verified Access will use to verify authenticity of the device tokens.
        /// </summary>
        [Input("publicSigningKeyUrl")]
        public Input<string>? PublicSigningKeyUrl { get; set; }

        /// <summary>
        /// The ID of the tenant application with the device-identity provider.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public VerifiedAccessTrustProviderDeviceOptionsArgs()
        {
        }
        public static new VerifiedAccessTrustProviderDeviceOptionsArgs Empty => new VerifiedAccessTrustProviderDeviceOptionsArgs();
    }
}
