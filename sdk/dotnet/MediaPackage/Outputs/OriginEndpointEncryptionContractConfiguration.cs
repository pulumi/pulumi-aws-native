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
    /// The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointEncryptionContractConfiguration
    {
        /// <summary>
        /// A collection of audio encryption presets.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Audio PresetSpeke20Audio;
        /// <summary>
        /// A collection of video encryption presets.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Video PresetSpeke20Video;

        [OutputConstructor]
        private OriginEndpointEncryptionContractConfiguration(
            Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Audio presetSpeke20Audio,

            Pulumi.AwsNative.MediaPackage.OriginEndpointEncryptionContractConfigurationPresetSpeke20Video presetSpeke20Video)
        {
            PresetSpeke20Audio = presetSpeke20Audio;
            PresetSpeke20Video = presetSpeke20Video;
        }
    }
}
