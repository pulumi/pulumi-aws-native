// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterTls
    {
        /// <summary>
        /// List of AWS Private CA ARNs.
        /// </summary>
        public readonly ImmutableArray<string> CertificateAuthorityArnList;
        /// <summary>
        /// TLS authentication is enabled or not.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private ClusterTls(
            ImmutableArray<string> certificateAuthorityArnList,

            bool? enabled)
        {
            CertificateAuthorityArnList = certificateAuthorityArnList;
            Enabled = enabled;
        }
    }
}
