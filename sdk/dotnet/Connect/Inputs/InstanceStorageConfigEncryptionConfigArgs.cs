// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    public sealed class InstanceStorageConfigEncryptionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of encryption.
        /// </summary>
        [Input("encryptionType", required: true)]
        public Input<Pulumi.AwsNative.Connect.InstanceStorageConfigEncryptionType> EncryptionType { get; set; } = null!;

        /// <summary>
        /// The full ARN of the encryption key.
        /// 
        /// &gt; Be sure to provide the full ARN of the encryption key, not just the ID.
        /// &gt; 
        /// &gt; Amazon Connect supports only KMS keys with the default key spec of [`SYMMETRIC_DEFAULT`](https://docs.aws.amazon.com/kms/latest/developerguide/asymmetric-key-specs.html#key-spec-symmetric-default) .
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public InstanceStorageConfigEncryptionConfigArgs()
        {
        }
        public static new InstanceStorageConfigEncryptionConfigArgs Empty => new InstanceStorageConfigEncryptionConfigArgs();
    }
}
