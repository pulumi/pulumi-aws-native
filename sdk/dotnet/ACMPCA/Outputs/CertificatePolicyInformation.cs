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
    /// Structure that contains X.509 Policy information.
    /// </summary>
    [OutputType]
    public sealed class CertificatePolicyInformation
    {
        public readonly string CertPolicyId;
        public readonly ImmutableArray<Outputs.CertificatePolicyQualifierInfo> PolicyQualifiers;

        [OutputConstructor]
        private CertificatePolicyInformation(
            string certPolicyId,

            ImmutableArray<Outputs.CertificatePolicyQualifierInfo> policyQualifiers)
        {
            CertPolicyId = certPolicyId;
            PolicyQualifiers = policyQualifiers;
        }
    }
}