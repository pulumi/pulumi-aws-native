// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class BucketInventoryTableConfiguration
    {
        /// <summary>
        /// Specifies whether inventory table configuration is enabled or disabled.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketInventoryTableConfigurationConfigurationState ConfigurationState;
        /// <summary>
        /// The encryption configuration for the inventory table.
        /// </summary>
        public readonly Outputs.BucketMetadataTableEncryptionConfiguration? EncryptionConfiguration;
        /// <summary>
        /// The ARN of the inventory table.
        /// </summary>
        public readonly string? TableArn;
        /// <summary>
        /// The name of the inventory table.
        /// </summary>
        public readonly string? TableName;

        [OutputConstructor]
        private BucketInventoryTableConfiguration(
            Pulumi.AwsNative.S3.BucketInventoryTableConfigurationConfigurationState configurationState,

            Outputs.BucketMetadataTableEncryptionConfiguration? encryptionConfiguration,

            string? tableArn,

            string? tableName)
        {
            ConfigurationState = configurationState;
            EncryptionConfiguration = encryptionConfiguration;
            TableArn = tableArn;
            TableName = tableName;
        }
    }
}
