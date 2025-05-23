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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetCollaborationResult',
    'AwaitableGetCollaborationResult',
    'get_collaboration',
    'get_collaboration_output',
]

@pulumi.output_type
class GetCollaborationResult:
    def __init__(__self__, analytics_engine=None, arn=None, collaboration_identifier=None, description=None, name=None, tags=None):
        if analytics_engine and not isinstance(analytics_engine, str):
            raise TypeError("Expected argument 'analytics_engine' to be a str")
        pulumi.set(__self__, "analytics_engine", analytics_engine)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if collaboration_identifier and not isinstance(collaboration_identifier, str):
            raise TypeError("Expected argument 'collaboration_identifier' to be a str")
        pulumi.set(__self__, "collaboration_identifier", collaboration_identifier)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="analyticsEngine")
    def analytics_engine(self) -> Optional['CollaborationAnalyticsEngine']:
        """
        The analytics engine for the collaboration.
        """
        return pulumi.get(self, "analytics_engine")

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Returns the Amazon Resource Name (ARN) of the specified collaboration.

        Example: `arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="collaborationIdentifier")
    def collaboration_identifier(self) -> Optional[builtins.str]:
        """
        Returns the unique identifier of the specified collaboration.

        Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        """
        return pulumi.get(self, "collaboration_identifier")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        A description of the collaboration provided by the collaboration owner.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        A human-readable identifier provided by the collaboration owner. Display names are not unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
        """
        return pulumi.get(self, "tags")


class AwaitableGetCollaborationResult(GetCollaborationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCollaborationResult(
            analytics_engine=self.analytics_engine,
            arn=self.arn,
            collaboration_identifier=self.collaboration_identifier,
            description=self.description,
            name=self.name,
            tags=self.tags)


def get_collaboration(collaboration_identifier: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCollaborationResult:
    """
    Represents a collaboration between AWS accounts that allows for secure data collaboration


    :param builtins.str collaboration_identifier: Returns the unique identifier of the specified collaboration.
           
           Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
    """
    __args__ = dict()
    __args__['collaborationIdentifier'] = collaboration_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cleanrooms:getCollaboration', __args__, opts=opts, typ=GetCollaborationResult).value

    return AwaitableGetCollaborationResult(
        analytics_engine=pulumi.get(__ret__, 'analytics_engine'),
        arn=pulumi.get(__ret__, 'arn'),
        collaboration_identifier=pulumi.get(__ret__, 'collaboration_identifier'),
        description=pulumi.get(__ret__, 'description'),
        name=pulumi.get(__ret__, 'name'),
        tags=pulumi.get(__ret__, 'tags'))
def get_collaboration_output(collaboration_identifier: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCollaborationResult]:
    """
    Represents a collaboration between AWS accounts that allows for secure data collaboration


    :param builtins.str collaboration_identifier: Returns the unique identifier of the specified collaboration.
           
           Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
    """
    __args__ = dict()
    __args__['collaborationIdentifier'] = collaboration_identifier
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cleanrooms:getCollaboration', __args__, opts=opts, typ=GetCollaborationResult)
    return __ret__.apply(lambda __response__: GetCollaborationResult(
        analytics_engine=pulumi.get(__response__, 'analytics_engine'),
        arn=pulumi.get(__response__, 'arn'),
        collaboration_identifier=pulumi.get(__response__, 'collaboration_identifier'),
        description=pulumi.get(__response__, 'description'),
        name=pulumi.get(__response__, 'name'),
        tags=pulumi.get(__response__, 'tags')))
