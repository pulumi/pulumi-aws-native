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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetObjectTypeResult',
    'AwaitableGetObjectTypeResult',
    'get_object_type',
    'get_object_type_output',
]

@pulumi.output_type
class GetObjectTypeResult:
    def __init__(__self__, allow_profile_creation=None, created_at=None, description=None, encryption_key=None, expiration_days=None, fields=None, keys=None, last_updated_at=None, max_available_profile_object_count=None, max_profile_object_count=None, source_last_updated_timestamp_format=None, tags=None, template_id=None):
        if allow_profile_creation and not isinstance(allow_profile_creation, bool):
            raise TypeError("Expected argument 'allow_profile_creation' to be a bool")
        pulumi.set(__self__, "allow_profile_creation", allow_profile_creation)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if encryption_key and not isinstance(encryption_key, str):
            raise TypeError("Expected argument 'encryption_key' to be a str")
        pulumi.set(__self__, "encryption_key", encryption_key)
        if expiration_days and not isinstance(expiration_days, int):
            raise TypeError("Expected argument 'expiration_days' to be a int")
        pulumi.set(__self__, "expiration_days", expiration_days)
        if fields and not isinstance(fields, list):
            raise TypeError("Expected argument 'fields' to be a list")
        pulumi.set(__self__, "fields", fields)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if last_updated_at and not isinstance(last_updated_at, str):
            raise TypeError("Expected argument 'last_updated_at' to be a str")
        pulumi.set(__self__, "last_updated_at", last_updated_at)
        if max_available_profile_object_count and not isinstance(max_available_profile_object_count, int):
            raise TypeError("Expected argument 'max_available_profile_object_count' to be a int")
        pulumi.set(__self__, "max_available_profile_object_count", max_available_profile_object_count)
        if max_profile_object_count and not isinstance(max_profile_object_count, int):
            raise TypeError("Expected argument 'max_profile_object_count' to be a int")
        pulumi.set(__self__, "max_profile_object_count", max_profile_object_count)
        if source_last_updated_timestamp_format and not isinstance(source_last_updated_timestamp_format, str):
            raise TypeError("Expected argument 'source_last_updated_timestamp_format' to be a str")
        pulumi.set(__self__, "source_last_updated_timestamp_format", source_last_updated_timestamp_format)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if template_id and not isinstance(template_id, str):
            raise TypeError("Expected argument 'template_id' to be a str")
        pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="allowProfileCreation")
    def allow_profile_creation(self) -> Optional[builtins.bool]:
        """
        Indicates whether a profile should be created when data is received.
        """
        return pulumi.get(self, "allow_profile_creation")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        The time of this integration got created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Description of the profile object type.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional[builtins.str]:
        """
        The default encryption key
        """
        return pulumi.get(self, "encryption_key")

    @property
    @pulumi.getter(name="expirationDays")
    def expiration_days(self) -> Optional[builtins.int]:
        """
        The default number of days until the data within the domain expires.
        """
        return pulumi.get(self, "expiration_days")

    @property
    @pulumi.getter
    def fields(self) -> Optional[Sequence['outputs.ObjectTypeFieldMap']]:
        """
        A list of the name and ObjectType field.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def keys(self) -> Optional[Sequence['outputs.ObjectTypeKeyMap']]:
        """
        A list of unique keys that can be used to map data to the profile.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> Optional[builtins.str]:
        """
        The time of this integration got last updated at.
        """
        return pulumi.get(self, "last_updated_at")

    @property
    @pulumi.getter(name="maxAvailableProfileObjectCount")
    def max_available_profile_object_count(self) -> Optional[builtins.int]:
        """
        The maximum available number of profile objects
        """
        return pulumi.get(self, "max_available_profile_object_count")

    @property
    @pulumi.getter(name="maxProfileObjectCount")
    def max_profile_object_count(self) -> Optional[builtins.int]:
        """
        The maximum number of profile objects for this object type
        """
        return pulumi.get(self, "max_profile_object_count")

    @property
    @pulumi.getter(name="sourceLastUpdatedTimestampFormat")
    def source_last_updated_timestamp_format(self) -> Optional[builtins.str]:
        """
        The format of your sourceLastUpdatedTimestamp that was previously set up.
        """
        return pulumi.get(self, "source_last_updated_timestamp_format")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The tags (keys and values) associated with the integration.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[builtins.str]:
        """
        A unique identifier for the object template.
        """
        return pulumi.get(self, "template_id")


class AwaitableGetObjectTypeResult(GetObjectTypeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObjectTypeResult(
            allow_profile_creation=self.allow_profile_creation,
            created_at=self.created_at,
            description=self.description,
            encryption_key=self.encryption_key,
            expiration_days=self.expiration_days,
            fields=self.fields,
            keys=self.keys,
            last_updated_at=self.last_updated_at,
            max_available_profile_object_count=self.max_available_profile_object_count,
            max_profile_object_count=self.max_profile_object_count,
            source_last_updated_timestamp_format=self.source_last_updated_timestamp_format,
            tags=self.tags,
            template_id=self.template_id)


def get_object_type(domain_name: Optional[builtins.str] = None,
                    object_type_name: Optional[builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetObjectTypeResult:
    """
    An ObjectType resource of Amazon Connect Customer Profiles


    :param builtins.str domain_name: The unique name of the domain.
    :param builtins.str object_type_name: The name of the profile object type.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['objectTypeName'] = object_type_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:customerprofiles:getObjectType', __args__, opts=opts, typ=GetObjectTypeResult).value

    return AwaitableGetObjectTypeResult(
        allow_profile_creation=pulumi.get(__ret__, 'allow_profile_creation'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        encryption_key=pulumi.get(__ret__, 'encryption_key'),
        expiration_days=pulumi.get(__ret__, 'expiration_days'),
        fields=pulumi.get(__ret__, 'fields'),
        keys=pulumi.get(__ret__, 'keys'),
        last_updated_at=pulumi.get(__ret__, 'last_updated_at'),
        max_available_profile_object_count=pulumi.get(__ret__, 'max_available_profile_object_count'),
        max_profile_object_count=pulumi.get(__ret__, 'max_profile_object_count'),
        source_last_updated_timestamp_format=pulumi.get(__ret__, 'source_last_updated_timestamp_format'),
        tags=pulumi.get(__ret__, 'tags'),
        template_id=pulumi.get(__ret__, 'template_id'))
def get_object_type_output(domain_name: Optional[pulumi.Input[builtins.str]] = None,
                           object_type_name: Optional[pulumi.Input[builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetObjectTypeResult]:
    """
    An ObjectType resource of Amazon Connect Customer Profiles


    :param builtins.str domain_name: The unique name of the domain.
    :param builtins.str object_type_name: The name of the profile object type.
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    __args__['objectTypeName'] = object_type_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:customerprofiles:getObjectType', __args__, opts=opts, typ=GetObjectTypeResult)
    return __ret__.apply(lambda __response__: GetObjectTypeResult(
        allow_profile_creation=pulumi.get(__response__, 'allow_profile_creation'),
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        encryption_key=pulumi.get(__response__, 'encryption_key'),
        expiration_days=pulumi.get(__response__, 'expiration_days'),
        fields=pulumi.get(__response__, 'fields'),
        keys=pulumi.get(__response__, 'keys'),
        last_updated_at=pulumi.get(__response__, 'last_updated_at'),
        max_available_profile_object_count=pulumi.get(__response__, 'max_available_profile_object_count'),
        max_profile_object_count=pulumi.get(__response__, 'max_profile_object_count'),
        source_last_updated_timestamp_format=pulumi.get(__response__, 'source_last_updated_timestamp_format'),
        tags=pulumi.get(__response__, 'tags'),
        template_id=pulumi.get(__response__, 'template_id')))
