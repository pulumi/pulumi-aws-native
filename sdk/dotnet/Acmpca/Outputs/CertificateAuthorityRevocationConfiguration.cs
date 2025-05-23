// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Outputs
{

    /// <summary>
    /// Certificate Authority revocation information.
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityRevocationConfiguration
    {
        /// <summary>
        /// Configuration of the certificate revocation list (CRL), if any, maintained by your private CA.
        /// </summary>
        public readonly Outputs.CertificateAuthorityCrlConfiguration? CrlConfiguration;
        /// <summary>
        /// Configuration of Online Certificate Status Protocol (OCSP) support, if any, maintained by your private CA.
        /// </summary>
        public readonly Outputs.CertificateAuthorityOcspConfiguration? OcspConfiguration;

        [OutputConstructor]
        private CertificateAuthorityRevocationConfiguration(
            Outputs.CertificateAuthorityCrlConfiguration? crlConfiguration,

            Outputs.CertificateAuthorityOcspConfiguration? ocspConfiguration)
        {
            CrlConfiguration = crlConfiguration;
            OcspConfiguration = ocspConfiguration;
        }
    }
}
