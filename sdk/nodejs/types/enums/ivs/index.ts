// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ChannelContainerFormat = {
    Ts: "TS",
    FragmentedMp4: "FRAGMENTED_MP4",
} as const;

/**
 * Indicates which content-packaging format is used (MPEG-TS or fMP4). If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.
 */
export type ChannelContainerFormat = (typeof ChannelContainerFormat)[keyof typeof ChannelContainerFormat];

export const ChannelLatencyMode = {
    Normal: "NORMAL",
    Low: "LOW",
} as const;

/**
 * Channel latency mode.
 */
export type ChannelLatencyMode = (typeof ChannelLatencyMode)[keyof typeof ChannelLatencyMode];

export const ChannelMultitrackInputConfigurationMaximumResolution = {
    Sd: "SD",
    Hd: "HD",
    FullHd: "FULL_HD",
} as const;

/**
 * Maximum resolution for multitrack input. Required if enabled is true.
 */
export type ChannelMultitrackInputConfigurationMaximumResolution = (typeof ChannelMultitrackInputConfigurationMaximumResolution)[keyof typeof ChannelMultitrackInputConfigurationMaximumResolution];

export const ChannelMultitrackInputConfigurationPolicy = {
    Allow: "ALLOW",
    Require: "REQUIRE",
} as const;

/**
 * Indicates whether multitrack input is allowed or required. Required if enabled is true.
 */
export type ChannelMultitrackInputConfigurationPolicy = (typeof ChannelMultitrackInputConfigurationPolicy)[keyof typeof ChannelMultitrackInputConfigurationPolicy];

export const ChannelPreset = {
    Empty: "",
    HigherBandwidthDelivery: "HIGHER_BANDWIDTH_DELIVERY",
    ConstrainedBandwidthDelivery: "CONSTRAINED_BANDWIDTH_DELIVERY",
} as const;

/**
 * Optional transcode preset for the channel. This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string ("").
 */
export type ChannelPreset = (typeof ChannelPreset)[keyof typeof ChannelPreset];

export const ChannelType = {
    Standard: "STANDARD",
    Basic: "BASIC",
    AdvancedSd: "ADVANCED_SD",
    AdvancedHd: "ADVANCED_HD",
} as const;

/**
 * Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
 */
export type ChannelType = (typeof ChannelType)[keyof typeof ChannelType];

export const IngestConfigurationIngestProtocol = {
    Rtmp: "RTMP",
    Rtmps: "RTMPS",
} as const;

/**
 * Ingest Protocol.
 */
export type IngestConfigurationIngestProtocol = (typeof IngestConfigurationIngestProtocol)[keyof typeof IngestConfigurationIngestProtocol];

export const IngestConfigurationState = {
    Active: "ACTIVE",
    Inactive: "INACTIVE",
} as const;

/**
 * State of IngestConfiguration which determines whether IngestConfiguration is in use or not.
 */
export type IngestConfigurationState = (typeof IngestConfigurationState)[keyof typeof IngestConfigurationState];

export const RecordingConfigurationRenditionConfigurationRenditionSelection = {
    All: "ALL",
    None: "NONE",
    Custom: "CUSTOM",
} as const;

/**
 * Resolution Selection indicates which set of renditions are recorded for a stream.
 */
export type RecordingConfigurationRenditionConfigurationRenditionSelection = (typeof RecordingConfigurationRenditionConfigurationRenditionSelection)[keyof typeof RecordingConfigurationRenditionConfigurationRenditionSelection];

export const RecordingConfigurationRenditionConfigurationRenditionsItem = {
    FullHd: "FULL_HD",
    Hd: "HD",
    Sd: "SD",
    LowestResolution: "LOWEST_RESOLUTION",
} as const;

export type RecordingConfigurationRenditionConfigurationRenditionsItem = (typeof RecordingConfigurationRenditionConfigurationRenditionsItem)[keyof typeof RecordingConfigurationRenditionConfigurationRenditionsItem];

export const RecordingConfigurationState = {
    Creating: "CREATING",
    CreateFailed: "CREATE_FAILED",
    Active: "ACTIVE",
} as const;

/**
 * Recording Configuration State.
 */
export type RecordingConfigurationState = (typeof RecordingConfigurationState)[keyof typeof RecordingConfigurationState];

export const RecordingConfigurationThumbnailConfigurationRecordingMode = {
    Interval: "INTERVAL",
    Disabled: "DISABLED",
} as const;

/**
 * Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
 */
export type RecordingConfigurationThumbnailConfigurationRecordingMode = (typeof RecordingConfigurationThumbnailConfigurationRecordingMode)[keyof typeof RecordingConfigurationThumbnailConfigurationRecordingMode];

export const RecordingConfigurationThumbnailConfigurationResolution = {
    FullHd: "FULL_HD",
    Hd: "HD",
    Sd: "SD",
    LowestResolution: "LOWEST_RESOLUTION",
} as const;

/**
 * Resolution indicates the desired resolution of recorded thumbnails.
 */
export type RecordingConfigurationThumbnailConfigurationResolution = (typeof RecordingConfigurationThumbnailConfigurationResolution)[keyof typeof RecordingConfigurationThumbnailConfigurationResolution];

export const RecordingConfigurationThumbnailConfigurationStorageItem = {
    Sequential: "SEQUENTIAL",
    Latest: "LATEST",
} as const;

export type RecordingConfigurationThumbnailConfigurationStorageItem = (typeof RecordingConfigurationThumbnailConfigurationStorageItem)[keyof typeof RecordingConfigurationThumbnailConfigurationStorageItem];

export const StageAutoParticipantRecordingConfigurationMediaTypesItem = {
    AudioVideo: "AUDIO_VIDEO",
    AudioOnly: "AUDIO_ONLY",
} as const;

export type StageAutoParticipantRecordingConfigurationMediaTypesItem = (typeof StageAutoParticipantRecordingConfigurationMediaTypesItem)[keyof typeof StageAutoParticipantRecordingConfigurationMediaTypesItem];
