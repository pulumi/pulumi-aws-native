// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Outputs
{

    /// <summary>
    /// Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityAccessMethod
    {
        public readonly string? AccessMethodType;
        public readonly string? CustomObjectIdentifier;

        [OutputConstructor]
        private CertificateAuthorityAccessMethod(
            string? accessMethodType,

            string? customObjectIdentifier)
        {
            AccessMethodType = accessMethodType;
            CustomObjectIdentifier = customObjectIdentifier;
        }
    }
}