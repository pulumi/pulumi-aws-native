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
    'GetEntityResult',
    'AwaitableGetEntityResult',
    'get_entity',
    'get_entity_output',
]

@pulumi.output_type
class GetEntityResult:
    def __init__(__self__, arn=None, components=None, composite_components=None, creation_date_time=None, description=None, entity_name=None, has_child_entities=None, parent_entity_id=None, status=None, tags=None, update_date_time=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if components and not isinstance(components, dict):
            raise TypeError("Expected argument 'components' to be a dict")
        pulumi.set(__self__, "components", components)
        if composite_components and not isinstance(composite_components, dict):
            raise TypeError("Expected argument 'composite_components' to be a dict")
        pulumi.set(__self__, "composite_components", composite_components)
        if creation_date_time and not isinstance(creation_date_time, str):
            raise TypeError("Expected argument 'creation_date_time' to be a str")
        pulumi.set(__self__, "creation_date_time", creation_date_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if entity_name and not isinstance(entity_name, str):
            raise TypeError("Expected argument 'entity_name' to be a str")
        pulumi.set(__self__, "entity_name", entity_name)
        if has_child_entities and not isinstance(has_child_entities, bool):
            raise TypeError("Expected argument 'has_child_entities' to be a bool")
        pulumi.set(__self__, "has_child_entities", has_child_entities)
        if parent_entity_id and not isinstance(parent_entity_id, str):
            raise TypeError("Expected argument 'parent_entity_id' to be a str")
        pulumi.set(__self__, "parent_entity_id", parent_entity_id)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if update_date_time and not isinstance(update_date_time, str):
            raise TypeError("Expected argument 'update_date_time' to be a str")
        pulumi.set(__self__, "update_date_time", update_date_time)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        The ARN of the entity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def components(self) -> Optional[Mapping[str, 'outputs.EntityComponent']]:
        """
        A map that sets information about a component type.
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="compositeComponents")
    def composite_components(self) -> Optional[Mapping[str, 'outputs.EntityCompositeComponent']]:
        """
        A map that sets information about a composite component.
        """
        return pulumi.get(self, "composite_components")

    @property
    @pulumi.getter(name="creationDateTime")
    def creation_date_time(self) -> Optional[builtins.str]:
        """
        The date and time when the entity was created.
        """
        return pulumi.get(self, "creation_date_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        The description of the entity.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> Optional[builtins.str]:
        """
        The name of the entity.
        """
        return pulumi.get(self, "entity_name")

    @property
    @pulumi.getter(name="hasChildEntities")
    def has_child_entities(self) -> Optional[builtins.bool]:
        """
        A Boolean value that specifies whether the entity has child entities or not.
        """
        return pulumi.get(self, "has_child_entities")

    @property
    @pulumi.getter(name="parentEntityId")
    def parent_entity_id(self) -> Optional[builtins.str]:
        """
        The ID of the parent entity.
        """
        return pulumi.get(self, "parent_entity_id")

    @property
    @pulumi.getter
    def status(self) -> Optional['outputs.EntityStatus']:
        """
        The current status of the entity.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        A key-value pair to associate with a resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateDateTime")
    def update_date_time(self) -> Optional[builtins.str]:
        """
        The last date and time when the entity was updated.
        """
        return pulumi.get(self, "update_date_time")


class AwaitableGetEntityResult(GetEntityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntityResult(
            arn=self.arn,
            components=self.components,
            composite_components=self.composite_components,
            creation_date_time=self.creation_date_time,
            description=self.description,
            entity_name=self.entity_name,
            has_child_entities=self.has_child_entities,
            parent_entity_id=self.parent_entity_id,
            status=self.status,
            tags=self.tags,
            update_date_time=self.update_date_time)


def get_entity(entity_id: Optional[builtins.str] = None,
               workspace_id: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntityResult:
    """
    Resource schema for AWS::IoTTwinMaker::Entity


    :param builtins.str entity_id: The ID of the entity.
    :param builtins.str workspace_id: The ID of the workspace.
    """
    __args__ = dict()
    __args__['entityId'] = entity_id
    __args__['workspaceId'] = workspace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iottwinmaker:getEntity', __args__, opts=opts, typ=GetEntityResult).value

    return AwaitableGetEntityResult(
        arn=pulumi.get(__ret__, 'arn'),
        components=pulumi.get(__ret__, 'components'),
        composite_components=pulumi.get(__ret__, 'composite_components'),
        creation_date_time=pulumi.get(__ret__, 'creation_date_time'),
        description=pulumi.get(__ret__, 'description'),
        entity_name=pulumi.get(__ret__, 'entity_name'),
        has_child_entities=pulumi.get(__ret__, 'has_child_entities'),
        parent_entity_id=pulumi.get(__ret__, 'parent_entity_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        update_date_time=pulumi.get(__ret__, 'update_date_time'))
def get_entity_output(entity_id: Optional[pulumi.Input[builtins.str]] = None,
                      workspace_id: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEntityResult]:
    """
    Resource schema for AWS::IoTTwinMaker::Entity


    :param builtins.str entity_id: The ID of the entity.
    :param builtins.str workspace_id: The ID of the workspace.
    """
    __args__ = dict()
    __args__['entityId'] = entity_id
    __args__['workspaceId'] = workspace_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:iottwinmaker:getEntity', __args__, opts=opts, typ=GetEntityResult)
    return __ret__.apply(lambda __response__: GetEntityResult(
        arn=pulumi.get(__response__, 'arn'),
        components=pulumi.get(__response__, 'components'),
        composite_components=pulumi.get(__response__, 'composite_components'),
        creation_date_time=pulumi.get(__response__, 'creation_date_time'),
        description=pulumi.get(__response__, 'description'),
        entity_name=pulumi.get(__response__, 'entity_name'),
        has_child_entities=pulumi.get(__response__, 'has_child_entities'),
        parent_entity_id=pulumi.get(__response__, 'parent_entity_id'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        update_date_time=pulumi.get(__response__, 'update_date_time')))
