// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SsmIncidents.Outputs
{

    /// <summary>
    /// The ReplicationSet regional configuration.
    /// </summary>
    [OutputType]
    public sealed class ReplicationSetRegionConfiguration
    {
        /// <summary>
        /// The AWS Key Management Service key ID or Key Alias to use to encrypt your replication set.
        /// </summary>
        public readonly string SseKmsKeyId;

        [OutputConstructor]
        private ReplicationSetRegionConfiguration(string sseKmsKeyId)
        {
            SseKmsKeyId = sseKmsKeyId;
        }
    }
}
