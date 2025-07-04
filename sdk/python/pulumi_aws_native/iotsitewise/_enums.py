# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'AssetModelDataType',
    'AssetModelDataTypeSpec',
    'AssetModelTypeName',
    'AssetPropertyNotificationState',
    'DatasetSourceSourceFormat',
    'DatasetSourceSourceType',
    'GatewayGreengrassV2CoreDeviceOperatingSystem',
    'PortalType',
]


@pulumi.type_token("aws-native:iotsitewise:AssetModelDataType")
class AssetModelDataType(builtins.str, Enum):
    STRING = "STRING"
    INTEGER = "INTEGER"
    DOUBLE = "DOUBLE"
    BOOLEAN = "BOOLEAN"
    STRUCT = "STRUCT"


@pulumi.type_token("aws-native:iotsitewise:AssetModelDataTypeSpec")
class AssetModelDataTypeSpec(builtins.str, Enum):
    AWSALARM_STATE = "AWS/ALARM_STATE"


@pulumi.type_token("aws-native:iotsitewise:AssetModelTypeName")
class AssetModelTypeName(builtins.str, Enum):
    MEASUREMENT = "Measurement"
    ATTRIBUTE = "Attribute"
    TRANSFORM = "Transform"
    METRIC = "Metric"


@pulumi.type_token("aws-native:iotsitewise:AssetPropertyNotificationState")
class AssetPropertyNotificationState(builtins.str, Enum):
    """
    The MQTT notification state (ENABLED or DISABLED) for this asset property.
    """
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


@pulumi.type_token("aws-native:iotsitewise:DatasetSourceSourceFormat")
class DatasetSourceSourceFormat(builtins.str, Enum):
    """
    The format of the dataset source associated with the dataset.
    """
    KNOWLEDGE_BASE = "KNOWLEDGE_BASE"


@pulumi.type_token("aws-native:iotsitewise:DatasetSourceSourceType")
class DatasetSourceSourceType(builtins.str, Enum):
    """
    The type of data source for the dataset.
    """
    KENDRA = "KENDRA"


@pulumi.type_token("aws-native:iotsitewise:GatewayGreengrassV2CoreDeviceOperatingSystem")
class GatewayGreengrassV2CoreDeviceOperatingSystem(builtins.str, Enum):
    """
    The operating system of the core device in AWS IoT Greengrass V2.
    """
    LINUX_AARCH64 = "LINUX_AARCH64"
    LINUX_AMD64 = "LINUX_AMD64"
    WINDOWS_AMD64 = "WINDOWS_AMD64"


@pulumi.type_token("aws-native:iotsitewise:PortalType")
class PortalType(builtins.str, Enum):
    """
    The type of portal
    """
    SITEWISE_PORTAL_V1 = "SITEWISE_PORTAL_V1"
    SITEWISE_PORTAL_V2 = "SITEWISE_PORTAL_V2"
