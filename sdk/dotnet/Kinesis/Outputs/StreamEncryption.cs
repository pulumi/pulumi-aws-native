// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kinesis.Outputs
{

    /// <summary>
    /// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.
    /// </summary>
    [OutputType]
    public sealed class StreamEncryption
    {
        /// <summary>
        /// The encryption type to use. The only valid value is KMS. 
        /// </summary>
        public readonly Pulumi.AwsNative.Kinesis.StreamEncryptionEncryptionType EncryptionType;
        /// <summary>
        /// The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
        /// </summary>
        public readonly string KeyId;

        [OutputConstructor]
        private StreamEncryption(
            Pulumi.AwsNative.Kinesis.StreamEncryptionEncryptionType encryptionType,

            string keyId)
        {
            EncryptionType = encryptionType;
            KeyId = keyId;
        }
    }
}