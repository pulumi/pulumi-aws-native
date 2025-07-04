// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dsql.Outputs
{

    /// <summary>
    /// The encryption configuration details for the cluster.
    /// </summary>
    [OutputType]
    public sealed class EncryptionDetailsProperties
    {
        /// <summary>
        /// The status of encryption for the cluster.
        /// </summary>
        public readonly string? EncryptionStatus;
        /// <summary>
        /// The type of encryption that protects data in the cluster.
        /// </summary>
        public readonly string? EncryptionType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.
        /// </summary>
        public readonly string? KmsKeyArn;

        [OutputConstructor]
        private EncryptionDetailsProperties(
            string? encryptionStatus,

            string? encryptionType,

            string? kmsKeyArn)
        {
            EncryptionStatus = encryptionStatus;
            EncryptionType = encryptionType;
            KmsKeyArn = kmsKeyArn;
        }
    }
}
