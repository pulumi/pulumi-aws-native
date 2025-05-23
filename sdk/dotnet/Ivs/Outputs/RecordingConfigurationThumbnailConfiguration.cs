// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs.Outputs
{

    /// <summary>
    /// Recording Thumbnail Configuration.
    /// </summary>
    [OutputType]
    public sealed class RecordingConfigurationThumbnailConfiguration
    {
        /// <summary>
        /// Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
        /// </summary>
        public readonly Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationRecordingMode? RecordingMode;
        /// <summary>
        /// Resolution indicates the desired resolution of recorded thumbnails.
        /// </summary>
        public readonly Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationResolution? Resolution;
        /// <summary>
        /// Storage indicates the format in which thumbnails are recorded.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationStorageItem> Storage;
        /// <summary>
        /// Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.
        /// </summary>
        public readonly int? TargetIntervalSeconds;

        [OutputConstructor]
        private RecordingConfigurationThumbnailConfiguration(
            Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationRecordingMode? recordingMode,

            Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationResolution? resolution,

            ImmutableArray<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationStorageItem> storage,

            int? targetIntervalSeconds)
        {
            RecordingMode = recordingMode;
            Resolution = resolution;
            Storage = storage;
            TargetIntervalSeconds = targetIntervalSeconds;
        }
    }
}
