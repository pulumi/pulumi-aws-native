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
from . import outputs
from ._enums import *

__all__ = [
    'ProfileAttributeMapping',
    'ProfileMappingRule',
    'TrustAnchorNotificationSetting',
    'TrustAnchorSource',
    'TrustAnchorSourceData0Properties',
    'TrustAnchorSourceData1Properties',
]

@pulumi.output_type
class ProfileAttributeMapping(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateField":
            suggest = "certificate_field"
        elif key == "mappingRules":
            suggest = "mapping_rules"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProfileAttributeMapping. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProfileAttributeMapping.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProfileAttributeMapping.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_field: 'ProfileCertificateField',
                 mapping_rules: Sequence['outputs.ProfileMappingRule']):
        """
        :param 'ProfileCertificateField' certificate_field: Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.
        :param Sequence['ProfileMappingRule'] mapping_rules: A list of mapping entries for every supported specifier or sub-field.
        """
        pulumi.set(__self__, "certificate_field", certificate_field)
        pulumi.set(__self__, "mapping_rules", mapping_rules)

    @property
    @pulumi.getter(name="certificateField")
    def certificate_field(self) -> 'ProfileCertificateField':
        """
        Fields (x509Subject, x509Issuer and x509SAN) within X.509 certificates.
        """
        return pulumi.get(self, "certificate_field")

    @property
    @pulumi.getter(name="mappingRules")
    def mapping_rules(self) -> Sequence['outputs.ProfileMappingRule']:
        """
        A list of mapping entries for every supported specifier or sub-field.
        """
        return pulumi.get(self, "mapping_rules")


@pulumi.output_type
class ProfileMappingRule(dict):
    def __init__(__self__, *,
                 specifier: builtins.str):
        """
        :param builtins.str specifier: Specifier within a certificate field, such as CN, OU, or UID from the Subject field.
        """
        pulumi.set(__self__, "specifier", specifier)

    @property
    @pulumi.getter
    def specifier(self) -> builtins.str:
        """
        Specifier within a certificate field, such as CN, OU, or UID from the Subject field.
        """
        return pulumi.get(self, "specifier")


@pulumi.output_type
class TrustAnchorNotificationSetting(dict):
    def __init__(__self__, *,
                 enabled: builtins.bool,
                 event: 'TrustAnchorNotificationEvent',
                 channel: Optional['TrustAnchorNotificationChannel'] = None,
                 threshold: Optional[builtins.float] = None):
        """
        :param builtins.bool enabled: Indicates whether the notification setting is enabled.
        :param 'TrustAnchorNotificationEvent' event: The event to which this notification setting is applied.
        :param 'TrustAnchorNotificationChannel' channel: The specified channel of notification. IAM Roles Anywhere uses CloudWatch metrics, EventBridge, and AWS Health Dashboard to notify for an event.
               
               > In the absence of a specific channel, IAM Roles Anywhere applies this setting to 'ALL' channels.
        :param builtins.float threshold: The number of days before a notification event. This value is required for a notification setting that is enabled.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "event", event)
        if channel is not None:
            pulumi.set(__self__, "channel", channel)
        if threshold is not None:
            pulumi.set(__self__, "threshold", threshold)

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        """
        Indicates whether the notification setting is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def event(self) -> 'TrustAnchorNotificationEvent':
        """
        The event to which this notification setting is applied.
        """
        return pulumi.get(self, "event")

    @property
    @pulumi.getter
    def channel(self) -> Optional['TrustAnchorNotificationChannel']:
        """
        The specified channel of notification. IAM Roles Anywhere uses CloudWatch metrics, EventBridge, and AWS Health Dashboard to notify for an event.

        > In the absence of a specific channel, IAM Roles Anywhere applies this setting to 'ALL' channels.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter
    def threshold(self) -> Optional[builtins.float]:
        """
        The number of days before a notification event. This value is required for a notification setting that is enabled.
        """
        return pulumi.get(self, "threshold")


@pulumi.output_type
class TrustAnchorSource(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceData":
            suggest = "source_data"
        elif key == "sourceType":
            suggest = "source_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrustAnchorSource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrustAnchorSource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrustAnchorSource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_data: Any,
                 source_type: 'TrustAnchorType'):
        """
        :param Union['TrustAnchorSourceData0Properties', 'TrustAnchorSourceData1Properties'] source_data: A union object representing the data field of the TrustAnchor depending on its type
        :param 'TrustAnchorType' source_type: The type of the TrustAnchor.
        """
        pulumi.set(__self__, "source_data", source_data)
        pulumi.set(__self__, "source_type", source_type)

    @property
    @pulumi.getter(name="sourceData")
    def source_data(self) -> Any:
        """
        A union object representing the data field of the TrustAnchor depending on its type
        """
        return pulumi.get(self, "source_data")

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> 'TrustAnchorType':
        """
        The type of the TrustAnchor.
        """
        return pulumi.get(self, "source_type")


@pulumi.output_type
class TrustAnchorSourceData0Properties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "x509CertificateData":
            suggest = "x509_certificate_data"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrustAnchorSourceData0Properties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrustAnchorSourceData0Properties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrustAnchorSourceData0Properties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 x509_certificate_data: builtins.str):
        pulumi.set(__self__, "x509_certificate_data", x509_certificate_data)

    @property
    @pulumi.getter(name="x509CertificateData")
    def x509_certificate_data(self) -> builtins.str:
        return pulumi.get(self, "x509_certificate_data")


@pulumi.output_type
class TrustAnchorSourceData1Properties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "acmPcaArn":
            suggest = "acm_pca_arn"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrustAnchorSourceData1Properties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrustAnchorSourceData1Properties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrustAnchorSourceData1Properties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 acm_pca_arn: builtins.str):
        pulumi.set(__self__, "acm_pca_arn", acm_pca_arn)

    @property
    @pulumi.getter(name="acmPcaArn")
    def acm_pca_arn(self) -> builtins.str:
        return pulumi.get(self, "acm_pca_arn")


