// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Express.Inputs
{

    /// <summary>
    /// Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).
    /// </summary>
    public sealed class DirectoryBucketBucketEncryptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("serverSideEncryptionConfiguration", required: true)]
        private InputList<Inputs.DirectoryBucketServerSideEncryptionRuleArgs>? _serverSideEncryptionConfiguration;

        /// <summary>
        /// Specifies the default server-side-encryption configuration.
        /// </summary>
        public InputList<Inputs.DirectoryBucketServerSideEncryptionRuleArgs> ServerSideEncryptionConfiguration
        {
            get => _serverSideEncryptionConfiguration ?? (_serverSideEncryptionConfiguration = new InputList<Inputs.DirectoryBucketServerSideEncryptionRuleArgs>());
            set => _serverSideEncryptionConfiguration = value;
        }

        public DirectoryBucketBucketEncryptionArgs()
        {
        }
        public static new DirectoryBucketBucketEncryptionArgs Empty => new DirectoryBucketBucketEncryptionArgs();
    }
}