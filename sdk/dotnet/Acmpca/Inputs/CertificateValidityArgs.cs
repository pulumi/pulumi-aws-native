// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    /// <summary>
    /// Length of time for which the certificate issued by your private certificate authority (CA), or by the private CA itself, is valid in days, months, or years. You can issue a certificate by calling the ``IssueCertificate`` operation.
    /// </summary>
    public sealed class CertificateValidityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the ``Value`` parameter represents days, months, or years.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// A long integer interpreted according to the value of ``Type``, below.
        /// </summary>
        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public CertificateValidityArgs()
        {
        }
        public static new CertificateValidityArgs Empty => new CertificateValidityArgs();
    }
}
