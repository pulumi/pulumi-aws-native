// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs.Inputs
{

    /// <summary>
    /// Recording Thumbnail Configuration.
    /// </summary>
    public sealed class RecordingConfigurationThumbnailConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
        /// </summary>
        [Input("recordingMode")]
        public Input<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationRecordingMode>? RecordingMode { get; set; }

        /// <summary>
        /// Resolution indicates the desired resolution of recorded thumbnails.
        /// </summary>
        [Input("resolution")]
        public Input<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationResolution>? Resolution { get; set; }

        [Input("storage")]
        private InputList<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationStorageItem>? _storage;

        /// <summary>
        /// Storage indicates the format in which thumbnails are recorded.
        /// </summary>
        public InputList<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationStorageItem> Storage
        {
            get => _storage ?? (_storage = new InputList<Pulumi.AwsNative.Ivs.RecordingConfigurationThumbnailConfigurationStorageItem>());
            set => _storage = value;
        }

        /// <summary>
        /// Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.
        /// </summary>
        [Input("targetIntervalSeconds")]
        public Input<int>? TargetIntervalSeconds { get; set; }

        public RecordingConfigurationThumbnailConfigurationArgs()
        {
        }
        public static new RecordingConfigurationThumbnailConfigurationArgs Empty => new RecordingConfigurationThumbnailConfigurationArgs();
    }
}