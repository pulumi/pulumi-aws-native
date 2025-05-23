// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecr.Outputs
{

    /// <summary>
    /// An array of objects representing the destination for a replication rule.
    /// </summary>
    [OutputType]
    public sealed class ReplicationConfigurationReplicationDestination
    {
        /// <summary>
        /// The Region to replicate to.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The AWS account ID of the Amazon ECR private registry to replicate to. When configuring cross-Region replication within your own registry, specify your own account ID.
        /// </summary>
        public readonly string RegistryId;

        [OutputConstructor]
        private ReplicationConfigurationReplicationDestination(
            string region,

            string registryId)
        {
            Region = region;
            RegistryId = registryId;
        }
    }
}
