// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationSpekeKeyProvider
    {
        /// <summary>
        /// Use `encryptionContractConfiguration` to configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.
        /// </summary>
        public readonly Outputs.PackagingConfigurationEncryptionContractConfiguration? EncryptionContractConfiguration;
        /// <summary>
        /// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API. Valid format: arn:aws:iam::{accountID}:role/{name}
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The system IDs to include in key requests.
        /// </summary>
        public readonly ImmutableArray<string> SystemIds;
        /// <summary>
        /// The URL of the external key provider service.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private PackagingConfigurationSpekeKeyProvider(
            Outputs.PackagingConfigurationEncryptionContractConfiguration? encryptionContractConfiguration,

            string roleArn,

            ImmutableArray<string> systemIds,

            string url)
        {
            EncryptionContractConfiguration = encryptionContractConfiguration;
            RoleArn = roleArn;
            SystemIds = systemIds;
            Url = url;
        }
    }
}
