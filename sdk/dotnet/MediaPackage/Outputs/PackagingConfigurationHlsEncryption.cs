// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// An HTTP Live Streaming (HLS) encryption configuration.
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationHlsEncryption
    {
        /// <summary>
        /// An HTTP Live Streaming (HLS) encryption configuration.
        /// </summary>
        public readonly string? ConstantInitializationVector;
        /// <summary>
        /// The encryption method to use.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaPackage.PackagingConfigurationHlsEncryptionEncryptionMethod? EncryptionMethod;
        public readonly Outputs.PackagingConfigurationSpekeKeyProvider SpekeKeyProvider;

        [OutputConstructor]
        private PackagingConfigurationHlsEncryption(
            string? constantInitializationVector,

            Pulumi.AwsNative.MediaPackage.PackagingConfigurationHlsEncryptionEncryptionMethod? encryptionMethod,

            Outputs.PackagingConfigurationSpekeKeyProvider spekeKeyProvider)
        {
            ConstantInitializationVector = constantInitializationVector;
            EncryptionMethod = encryptionMethod;
            SpekeKeyProvider = spekeKeyProvider;
        }
    }
}