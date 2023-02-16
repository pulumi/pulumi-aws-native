// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.
    /// </summary>
    [OutputType]
    public sealed class BucketReplicationConfiguration
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// A container for one or more replication rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketReplicationRule> Rules;

        [OutputConstructor]
        private BucketReplicationConfiguration(
            string role,

            ImmutableArray<Outputs.BucketReplicationRule> rules)
        {
            Role = role;
            Rules = rules;
        }
    }
}