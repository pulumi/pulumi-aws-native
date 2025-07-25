# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'ChannelInputType',
    'OriginEndpointAdMarkerDash',
    'OriginEndpointAdMarkerHls',
    'OriginEndpointCmafEncryptionMethod',
    'OriginEndpointContainerType',
    'OriginEndpointDashCompactness',
    'OriginEndpointDashDrmSignaling',
    'OriginEndpointDashPeriodTrigger',
    'OriginEndpointDashProfile',
    'OriginEndpointDashSegmentTemplateFormat',
    'OriginEndpointDashTtmlProfile',
    'OriginEndpointDashUtcTimingMode',
    'OriginEndpointDrmSystem',
    'OriginEndpointEndpointErrorCondition',
    'OriginEndpointPresetSpeke20Audio',
    'OriginEndpointPresetSpeke20Video',
    'OriginEndpointScteFilter',
    'OriginEndpointTsEncryptionMethod',
]


@pulumi.type_token("aws-native:mediapackagev2:ChannelInputType")
class ChannelInputType(builtins.str, Enum):
    HLS = "HLS"
    CMAF = "CMAF"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointAdMarkerDash")
class OriginEndpointAdMarkerDash(builtins.str, Enum):
    BINARY = "BINARY"
    XML = "XML"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointAdMarkerHls")
class OriginEndpointAdMarkerHls(builtins.str, Enum):
    DATERANGE = "DATERANGE"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointCmafEncryptionMethod")
class OriginEndpointCmafEncryptionMethod(builtins.str, Enum):
    CENC = "CENC"
    CBCS = "CBCS"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointContainerType")
class OriginEndpointContainerType(builtins.str, Enum):
    TS = "TS"
    CMAF = "CMAF"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashCompactness")
class OriginEndpointDashCompactness(builtins.str, Enum):
    STANDARD = "STANDARD"
    NONE = "NONE"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashDrmSignaling")
class OriginEndpointDashDrmSignaling(builtins.str, Enum):
    INDIVIDUAL = "INDIVIDUAL"
    REFERENCED = "REFERENCED"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashPeriodTrigger")
class OriginEndpointDashPeriodTrigger(builtins.str, Enum):
    AVAILS = "AVAILS"
    DRM_KEY_ROTATION = "DRM_KEY_ROTATION"
    SOURCE_CHANGES = "SOURCE_CHANGES"
    SOURCE_DISRUPTIONS = "SOURCE_DISRUPTIONS"
    NONE = "NONE"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashProfile")
class OriginEndpointDashProfile(builtins.str, Enum):
    DVB_DASH = "DVB_DASH"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashSegmentTemplateFormat")
class OriginEndpointDashSegmentTemplateFormat(builtins.str, Enum):
    NUMBER_WITH_TIMELINE = "NUMBER_WITH_TIMELINE"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashTtmlProfile")
class OriginEndpointDashTtmlProfile(builtins.str, Enum):
    IMSC1 = "IMSC_1"
    EBU_TT_D101 = "EBU_TT_D_101"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDashUtcTimingMode")
class OriginEndpointDashUtcTimingMode(builtins.str, Enum):
    HTTP_HEAD = "HTTP_HEAD"
    HTTP_ISO = "HTTP_ISO"
    HTTP_XSDATE = "HTTP_XSDATE"
    UTC_DIRECT = "UTC_DIRECT"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointDrmSystem")
class OriginEndpointDrmSystem(builtins.str, Enum):
    CLEAR_KEY_AES128 = "CLEAR_KEY_AES_128"
    FAIRPLAY = "FAIRPLAY"
    PLAYREADY = "PLAYREADY"
    WIDEVINE = "WIDEVINE"
    IRDETO = "IRDETO"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointEndpointErrorCondition")
class OriginEndpointEndpointErrorCondition(builtins.str, Enum):
    STALE_MANIFEST = "STALE_MANIFEST"
    INCOMPLETE_MANIFEST = "INCOMPLETE_MANIFEST"
    MISSING_DRM_KEY = "MISSING_DRM_KEY"
    SLATE_INPUT = "SLATE_INPUT"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointPresetSpeke20Audio")
class OriginEndpointPresetSpeke20Audio(builtins.str, Enum):
    PRESET_AUDIO1 = "PRESET_AUDIO_1"
    PRESET_AUDIO2 = "PRESET_AUDIO_2"
    PRESET_AUDIO3 = "PRESET_AUDIO_3"
    SHARED = "SHARED"
    UNENCRYPTED = "UNENCRYPTED"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointPresetSpeke20Video")
class OriginEndpointPresetSpeke20Video(builtins.str, Enum):
    PRESET_VIDEO1 = "PRESET_VIDEO_1"
    PRESET_VIDEO2 = "PRESET_VIDEO_2"
    PRESET_VIDEO3 = "PRESET_VIDEO_3"
    PRESET_VIDEO4 = "PRESET_VIDEO_4"
    PRESET_VIDEO5 = "PRESET_VIDEO_5"
    PRESET_VIDEO6 = "PRESET_VIDEO_6"
    PRESET_VIDEO7 = "PRESET_VIDEO_7"
    PRESET_VIDEO8 = "PRESET_VIDEO_8"
    SHARED = "SHARED"
    UNENCRYPTED = "UNENCRYPTED"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointScteFilter")
class OriginEndpointScteFilter(builtins.str, Enum):
    SPLICE_INSERT = "SPLICE_INSERT"
    BREAK_ = "BREAK"
    PROVIDER_ADVERTISEMENT = "PROVIDER_ADVERTISEMENT"
    DISTRIBUTOR_ADVERTISEMENT = "DISTRIBUTOR_ADVERTISEMENT"
    PROVIDER_PLACEMENT_OPPORTUNITY = "PROVIDER_PLACEMENT_OPPORTUNITY"
    DISTRIBUTOR_PLACEMENT_OPPORTUNITY = "DISTRIBUTOR_PLACEMENT_OPPORTUNITY"
    PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY = "PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY"
    DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY = "DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY"
    PROGRAM = "PROGRAM"


@pulumi.type_token("aws-native:mediapackagev2:OriginEndpointTsEncryptionMethod")
class OriginEndpointTsEncryptionMethod(builtins.str, Enum):
    AES128 = "AES_128"
    SAMPLE_AES = "SAMPLE_AES"
