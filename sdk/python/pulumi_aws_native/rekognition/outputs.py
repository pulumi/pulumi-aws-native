# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'StreamProcessorBoundingBox',
    'StreamProcessorConnectedHomeSettings',
    'StreamProcessorDataSharingPreference',
    'StreamProcessorFaceSearchSettings',
    'StreamProcessorKinesisDataStream',
    'StreamProcessorKinesisVideoStream',
    'StreamProcessorNotificationChannel',
    'StreamProcessorPoint',
    'StreamProcessorS3Destination',
]

@pulumi.output_type
class StreamProcessorBoundingBox(dict):
    """
    A bounding box denoting a region of interest in the frame to be analyzed.
    """
    def __init__(__self__, *,
                 height: builtins.float,
                 left: builtins.float,
                 top: builtins.float,
                 width: builtins.float):
        """
        A bounding box denoting a region of interest in the frame to be analyzed.
        :param builtins.float height: Height of the bounding box as a ratio of the overall image height.
        :param builtins.float left: Left coordinate of the bounding box as a ratio of overall image width.
        :param builtins.float top: Top coordinate of the bounding box as a ratio of overall image height.
        :param builtins.float width: Width of the bounding box as a ratio of the overall image width.
        """
        pulumi.set(__self__, "height", height)
        pulumi.set(__self__, "left", left)
        pulumi.set(__self__, "top", top)
        pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def height(self) -> builtins.float:
        """
        Height of the bounding box as a ratio of the overall image height.
        """
        return pulumi.get(self, "height")

    @property
    @pulumi.getter
    def left(self) -> builtins.float:
        """
        Left coordinate of the bounding box as a ratio of overall image width.
        """
        return pulumi.get(self, "left")

    @property
    @pulumi.getter
    def top(self) -> builtins.float:
        """
        Top coordinate of the bounding box as a ratio of overall image height.
        """
        return pulumi.get(self, "top")

    @property
    @pulumi.getter
    def width(self) -> builtins.float:
        """
        Width of the bounding box as a ratio of the overall image width.
        """
        return pulumi.get(self, "width")


@pulumi.output_type
class StreamProcessorConnectedHomeSettings(dict):
    """
    Connected home settings to use on a streaming video. Note that either ConnectedHomeSettings or FaceSearchSettings should be set. Not both
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "minConfidence":
            suggest = "min_confidence"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamProcessorConnectedHomeSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamProcessorConnectedHomeSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamProcessorConnectedHomeSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 labels: Sequence[builtins.str],
                 min_confidence: Optional[builtins.float] = None):
        """
        Connected home settings to use on a streaming video. Note that either ConnectedHomeSettings or FaceSearchSettings should be set. Not both
        :param Sequence[builtins.str] labels: Specifies what you want to detect in the video, such as people, packages, or pets. The current valid labels you can include in this list are: "PERSON", "PET", "PACKAGE", and "ALL".
        :param builtins.float min_confidence: Minimum object class match confidence score that must be met to return a result for a recognized object.
        """
        pulumi.set(__self__, "labels", labels)
        if min_confidence is not None:
            pulumi.set(__self__, "min_confidence", min_confidence)

    @property
    @pulumi.getter
    def labels(self) -> Sequence[builtins.str]:
        """
        Specifies what you want to detect in the video, such as people, packages, or pets. The current valid labels you can include in this list are: "PERSON", "PET", "PACKAGE", and "ALL".
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="minConfidence")
    def min_confidence(self) -> Optional[builtins.float]:
        """
        Minimum object class match confidence score that must be met to return a result for a recognized object.
        """
        return pulumi.get(self, "min_confidence")


@pulumi.output_type
class StreamProcessorDataSharingPreference(dict):
    """
    Indicates whether Rekognition is allowed to store the video stream data for model-training.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "optIn":
            suggest = "opt_in"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamProcessorDataSharingPreference. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamProcessorDataSharingPreference.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamProcessorDataSharingPreference.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 opt_in: builtins.bool):
        """
        Indicates whether Rekognition is allowed to store the video stream data for model-training.
        :param builtins.bool opt_in: Flag to enable data-sharing
        """
        pulumi.set(__self__, "opt_in", opt_in)

    @property
    @pulumi.getter(name="optIn")
    def opt_in(self) -> builtins.bool:
        """
        Flag to enable data-sharing
        """
        return pulumi.get(self, "opt_in")


@pulumi.output_type
class StreamProcessorFaceSearchSettings(dict):
    """
    Face search settings to use on a streaming video. Note that either FaceSearchSettings or ConnectedHomeSettings should be set. Not both
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "collectionId":
            suggest = "collection_id"
        elif key == "faceMatchThreshold":
            suggest = "face_match_threshold"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamProcessorFaceSearchSettings. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamProcessorFaceSearchSettings.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamProcessorFaceSearchSettings.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 collection_id: builtins.str,
                 face_match_threshold: Optional[builtins.float] = None):
        """
        Face search settings to use on a streaming video. Note that either FaceSearchSettings or ConnectedHomeSettings should be set. Not both
        :param builtins.str collection_id: The ID of a collection that contains faces that you want to search for.
        :param builtins.float face_match_threshold: Minimum face match confidence score percentage that must be met to return a result for a recognized face. The default is 80. 0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted.
        """
        pulumi.set(__self__, "collection_id", collection_id)
        if face_match_threshold is not None:
            pulumi.set(__self__, "face_match_threshold", face_match_threshold)

    @property
    @pulumi.getter(name="collectionId")
    def collection_id(self) -> builtins.str:
        """
        The ID of a collection that contains faces that you want to search for.
        """
        return pulumi.get(self, "collection_id")

    @property
    @pulumi.getter(name="faceMatchThreshold")
    def face_match_threshold(self) -> Optional[builtins.float]:
        """
        Minimum face match confidence score percentage that must be met to return a result for a recognized face. The default is 80. 0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted.
        """
        return pulumi.get(self, "face_match_threshold")


@pulumi.output_type
class StreamProcessorKinesisDataStream(dict):
    """
    The Amazon Kinesis Data Stream stream to which the Amazon Rekognition stream processor streams the analysis results, as part of face search feature.
    """
    def __init__(__self__, *,
                 arn: builtins.str):
        """
        The Amazon Kinesis Data Stream stream to which the Amazon Rekognition stream processor streams the analysis results, as part of face search feature.
        :param builtins.str arn: ARN of the Kinesis Data Stream stream.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> builtins.str:
        """
        ARN of the Kinesis Data Stream stream.
        """
        return pulumi.get(self, "arn")


@pulumi.output_type
class StreamProcessorKinesisVideoStream(dict):
    """
    The Kinesis Video Stream that streams the source video.
    """
    def __init__(__self__, *,
                 arn: builtins.str):
        """
        The Kinesis Video Stream that streams the source video.
        :param builtins.str arn: ARN of the Kinesis Video Stream that streams the source video.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> builtins.str:
        """
        ARN of the Kinesis Video Stream that streams the source video.
        """
        return pulumi.get(self, "arn")


@pulumi.output_type
class StreamProcessorNotificationChannel(dict):
    """
    The ARN of the SNS notification channel where events of interests are published, as part of connected home feature.
    """
    def __init__(__self__, *,
                 arn: builtins.str):
        """
        The ARN of the SNS notification channel where events of interests are published, as part of connected home feature.
        :param builtins.str arn: ARN of the SNS topic.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> builtins.str:
        """
        ARN of the SNS topic.
        """
        return pulumi.get(self, "arn")


@pulumi.output_type
class StreamProcessorPoint(dict):
    """
    An (X, Y) cartesian coordinate denoting a point on the frame
    """
    def __init__(__self__, *,
                 x: builtins.float,
                 y: builtins.float):
        """
        An (X, Y) cartesian coordinate denoting a point on the frame
        :param builtins.float x: The X coordinate of the point.
        :param builtins.float y: The Y coordinate of the point.
        """
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def x(self) -> builtins.float:
        """
        The X coordinate of the point.
        """
        return pulumi.get(self, "x")

    @property
    @pulumi.getter
    def y(self) -> builtins.float:
        """
        The Y coordinate of the point.
        """
        return pulumi.get(self, "y")


@pulumi.output_type
class StreamProcessorS3Destination(dict):
    """
    The S3 location in customer's account where inference output & artifacts are stored, as part of connected home feature.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bucketName":
            suggest = "bucket_name"
        elif key == "objectKeyPrefix":
            suggest = "object_key_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StreamProcessorS3Destination. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StreamProcessorS3Destination.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StreamProcessorS3Destination.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 bucket_name: builtins.str,
                 object_key_prefix: Optional[builtins.str] = None):
        """
        The S3 location in customer's account where inference output & artifacts are stored, as part of connected home feature.
        :param builtins.str bucket_name: Name of the S3 bucket.
        :param builtins.str object_key_prefix: The object key prefix path where the results will be stored. Default is no prefix path
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        if object_key_prefix is not None:
            pulumi.set(__self__, "object_key_prefix", object_key_prefix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> builtins.str:
        """
        Name of the S3 bucket.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="objectKeyPrefix")
    def object_key_prefix(self) -> Optional[builtins.str]:
        """
        The object key prefix path where the results will be stored. Default is no prefix path
        """
        return pulumi.get(self, "object_key_prefix")


