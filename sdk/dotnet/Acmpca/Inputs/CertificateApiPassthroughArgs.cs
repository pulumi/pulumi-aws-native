// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    /// <summary>
    /// Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
    /// </summary>
    public sealed class CertificateApiPassthroughArgs : global::Pulumi.ResourceArgs
    {
        [Input("extensions")]
        public Input<Inputs.CertificateExtensionsArgs>? Extensions { get; set; }

        /// <summary>
        /// Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
        /// </summary>
        [Input("subject")]
        public Input<Inputs.CertificateSubjectArgs>? Subject { get; set; }

        public CertificateApiPassthroughArgs()
        {
        }
        public static new CertificateApiPassthroughArgs Empty => new CertificateApiPassthroughArgs();
    }
}