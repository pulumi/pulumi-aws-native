// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Outputs
{

    [OutputType]
    public sealed class InstanceStorageConfigS3Config
    {
        public readonly string BucketName;
        public readonly string BucketPrefix;
        public readonly Outputs.InstanceStorageConfigEncryptionConfig? EncryptionConfig;

        [OutputConstructor]
        private InstanceStorageConfigS3Config(
            string bucketName,

            string bucketPrefix,

            Outputs.InstanceStorageConfigEncryptionConfig? encryptionConfig)
        {
            BucketName = bucketName;
            BucketPrefix = bucketPrefix;
            EncryptionConfig = encryptionConfig;
        }
    }
}