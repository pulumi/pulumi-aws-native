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
    /// Specifies default encryption for a bucket using server-side encryption with either Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS).
    /// </summary>
    [OutputType]
    public sealed class BucketEncryption
    {
        /// <summary>
        /// Specifies the default server-side-encryption configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.BucketServerSideEncryptionRule> ServerSideEncryptionConfiguration;

        [OutputConstructor]
        private BucketEncryption(ImmutableArray<Outputs.BucketServerSideEncryptionRule> serverSideEncryptionConfiguration)
        {
            ServerSideEncryptionConfiguration = serverSideEncryptionConfiguration;
        }
    }
}