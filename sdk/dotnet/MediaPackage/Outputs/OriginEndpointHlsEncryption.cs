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
    public sealed class OriginEndpointHlsEncryption
    {
        /// <summary>
        /// A constant initialization vector for encryption (optional). When not specified the initialization vector will be periodically rotated.
        /// </summary>
        public readonly string? ConstantInitializationVector;
        /// <summary>
        /// The encryption method to use.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaPackage.OriginEndpointHlsEncryptionEncryptionMethod? EncryptionMethod;
        /// <summary>
        /// Interval (in seconds) between each encryption key rotation.
        /// </summary>
        public readonly int? KeyRotationIntervalSeconds;
        /// <summary>
        /// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
        /// </summary>
        public readonly bool? RepeatExtXKey;
        public readonly Outputs.OriginEndpointSpekeKeyProvider SpekeKeyProvider;

        [OutputConstructor]
        private OriginEndpointHlsEncryption(
            string? constantInitializationVector,

            Pulumi.AwsNative.MediaPackage.OriginEndpointHlsEncryptionEncryptionMethod? encryptionMethod,

            int? keyRotationIntervalSeconds,

            bool? repeatExtXKey,

            Outputs.OriginEndpointSpekeKeyProvider spekeKeyProvider)
        {
            ConstantInitializationVector = constantInitializationVector;
            EncryptionMethod = encryptionMethod;
            KeyRotationIntervalSeconds = keyRotationIntervalSeconds;
            RepeatExtXKey = repeatExtXKey;
            SpekeKeyProvider = spekeKeyProvider;
        }
    }
}