// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityLake.Inputs
{

    /// <summary>
    /// Provides encryption details of Amazon Security Lake object.
    /// </summary>
    public sealed class DataLakeEncryptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        public DataLakeEncryptionConfigurationArgs()
        {
        }
        public static new DataLakeEncryptionConfigurationArgs Empty => new DataLakeEncryptionConfigurationArgs();
    }
}