// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Outputs
{

    /// <summary>
    /// Contains information about the certificate subject. The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
    /// </summary>
    [OutputType]
    public sealed class CertificateGeneralName
    {
        /// <summary>
        /// Contains information about the certificate subject. The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity.
        /// </summary>
        public readonly Outputs.CertificateSubject? DirectoryName;
        public readonly string? DnsName;
        public readonly Outputs.CertificateEdiPartyName? EdiPartyName;
        public readonly string? IpAddress;
        public readonly Outputs.CertificateOtherName? OtherName;
        public readonly string? RegisteredId;
        public readonly string? Rfc822Name;
        public readonly string? UniformResourceIdentifier;

        [OutputConstructor]
        private CertificateGeneralName(
            Outputs.CertificateSubject? directoryName,

            string? dnsName,

            Outputs.CertificateEdiPartyName? ediPartyName,

            string? ipAddress,

            Outputs.CertificateOtherName? otherName,

            string? registeredId,

            string? rfc822Name,

            string? uniformResourceIdentifier)
        {
            DirectoryName = directoryName;
            DnsName = dnsName;
            EdiPartyName = ediPartyName;
            IpAddress = ipAddress;
            OtherName = otherName;
            RegisteredId = registeredId;
            Rfc822Name = rfc822Name;
            UniformResourceIdentifier = uniformResourceIdentifier;
        }
    }
}