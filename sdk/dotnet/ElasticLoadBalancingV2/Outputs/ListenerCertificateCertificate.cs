// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class ListenerCertificateCertificate
    {
        public readonly string? CertificateArn;

        [OutputConstructor]
        private ListenerCertificateCertificate(string? certificateArn)
        {
            CertificateArn = certificateArn;
        }
    }
}