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
    /// Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years. You can issue a certificate by calling the ``IssueCertificate`` operation.
    /// </summary>
    [OutputType]
    public sealed class CertificateValidity
    {
        /// <summary>
        /// Specifies whether the ``Value`` parameter represents days, months, or years.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Time period.
        /// </summary>
        public readonly double Value;

        [OutputConstructor]
        private CertificateValidity(
            string type,

            double value)
        {
            Type = type;
            Value = value;
        }
    }
}