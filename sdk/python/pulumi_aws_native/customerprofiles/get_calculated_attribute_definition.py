# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetCalculatedAttributeDefinitionResult',
    'AwaitableGetCalculatedAttributeDefinitionResult',
    'get_calculated_attribute_definition',
    'get_calculated_attribute_definition_output',
]

@pulumi.output_type
class GetCalculatedAttributeDefinitionResult:
    def __init__(__self__, attribute_details=None, conditions=None, created_at=None, description=None, display_name=None, last_updated_at=None, statistic=None, tags=None):
        if attribute_details and not isinstance(attribute_details, dict):
            raise TypeError("Expected argument 'attribute_details' to be a dict")
        pulumi.set(__self__, "attribute_details", attribute_details)
        if conditions and not isinstance(conditions, dict):
            raise TypeError("Expected argument 'conditions' to be a dict")
        pulumi.set(__self__, "conditions", conditions)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if last_updated_at and not isinstance(last_updated_at, str):
            raise TypeError("Expected argument 'last_updated_at' to be a str")
        pulumi.set(__self__, "last_updated_at", last_updated_at)
        if statistic and not isinstance(statistic, str):
            raise TypeError("Expected argument 'statistic' to be a str")
        pulumi.set(__self__, "statistic", statistic)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="attributeDetails")
    def attribute_details(self) -> Optional['outputs.CalculatedAttributeDefinitionAttributeDetails']:
        return pulumi.get(self, "attribute_details")

    @property
    @pulumi.getter
    def conditions(self) -> Optional['outputs.CalculatedAttributeDefinitionConditions']:
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of when the calculated attribute definition was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> Optional[str]:
        """
        The timestamp of when the calculated attribute definition was most recently edited.
        """
        return pulumi.get(self, "last_updated_at")

    @property
    @pulumi.getter
    def statistic(self) -> Optional['CalculatedAttributeDefinitionStatistic']:
        return pulumi.get(self, "statistic")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.CalculatedAttributeDefinitionTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetCalculatedAttributeDefinitionResult(GetCalculatedAttributeDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCalculatedAttributeDefinitionResult(
            attribute_details=self.attribute_details,
            conditions=self.conditions,
            created_at=self.created_at,
            description=self.description,
            display_name=self.display_name,
            last_updated_at=self.last_updated_at,
            statistic=self.statistic,
            tags=self.tags)


def get_calculated_attribute_definition(calculated_attribute_name: Optional[str] = None,
                                        domain_name: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCalculatedAttributeDefinitionResult:
    """
    A calculated attribute definition for Customer Profiles
    """
    __args__ = dict()
    __args__['calculatedAttributeName'] = calculated_attribute_name
    __args__['domainName'] = domain_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:customerprofiles:getCalculatedAttributeDefinition', __args__, opts=opts, typ=GetCalculatedAttributeDefinitionResult).value

    return AwaitableGetCalculatedAttributeDefinitionResult(
        attribute_details=pulumi.get(__ret__, 'attribute_details'),
        conditions=pulumi.get(__ret__, 'conditions'),
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        last_updated_at=pulumi.get(__ret__, 'last_updated_at'),
        statistic=pulumi.get(__ret__, 'statistic'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_calculated_attribute_definition)
def get_calculated_attribute_definition_output(calculated_attribute_name: Optional[pulumi.Input[str]] = None,
                                               domain_name: Optional[pulumi.Input[str]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCalculatedAttributeDefinitionResult]:
    """
    A calculated attribute definition for Customer Profiles
    """
    ...