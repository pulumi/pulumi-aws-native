// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TrustStoreRevocation
    {
        /// <summary>
        /// The number of revoked certificates.
        /// </summary>
        public readonly int? NumberOfRevokedEntries;
        /// <summary>
        /// The revocation ID of the revocation file.
        /// </summary>
        public readonly string? RevocationId;
        /// <summary>
        /// The type of revocation file.
        /// </summary>
        public readonly string? RevocationType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the trust store.
        /// </summary>
        public readonly string? TrustStoreArn;

        [OutputConstructor]
        private TrustStoreRevocation(
            int? numberOfRevokedEntries,

            string? revocationId,

            string? revocationType,

            string? trustStoreArn)
        {
            NumberOfRevokedEntries = numberOfRevokedEntries;
            RevocationId = revocationId;
            RevocationType = revocationType;
            TrustStoreArn = trustStoreArn;
        }
    }
}
