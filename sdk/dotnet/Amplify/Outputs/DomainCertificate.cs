// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify.Outputs
{

    [OutputType]
    public sealed class DomainCertificate
    {
        public readonly string? CertificateArn;
        public readonly Pulumi.AwsNative.Amplify.DomainCertificateCertificateType? CertificateType;
        public readonly string? CertificateVerificationDnsRecord;

        [OutputConstructor]
        private DomainCertificate(
            string? certificateArn,

            Pulumi.AwsNative.Amplify.DomainCertificateCertificateType? certificateType,

            string? certificateVerificationDnsRecord)
        {
            CertificateArn = certificateArn;
            CertificateType = certificateType;
            CertificateVerificationDnsRecord = certificateVerificationDnsRecord;
        }
    }
}