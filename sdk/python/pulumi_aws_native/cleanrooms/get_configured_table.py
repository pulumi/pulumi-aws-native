# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetConfiguredTableResult',
    'AwaitableGetConfiguredTableResult',
    'get_configured_table',
    'get_configured_table_output',
]

@pulumi.output_type
class GetConfiguredTableResult:
    def __init__(__self__, analysis_rules=None, arn=None, configured_table_identifier=None, description=None, name=None, tags=None):
        if analysis_rules and not isinstance(analysis_rules, list):
            raise TypeError("Expected argument 'analysis_rules' to be a list")
        pulumi.set(__self__, "analysis_rules", analysis_rules)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if configured_table_identifier and not isinstance(configured_table_identifier, str):
            raise TypeError("Expected argument 'configured_table_identifier' to be a str")
        pulumi.set(__self__, "configured_table_identifier", configured_table_identifier)
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
    @pulumi.getter(name="analysisRules")
    def analysis_rules(self) -> Optional[Sequence['outputs.ConfiguredTableAnalysisRule']]:
        """
        The analysis rule that was created for the configured table.
        """
        return pulumi.get(self, "analysis_rules")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Returns the Amazon Resource Name (ARN) of the specified configured table.

        Example: `arn:aws:cleanrooms:us-east-1:111122223333:configuredtable/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configuredTableIdentifier")
    def configured_table_identifier(self) -> Optional[str]:
        """
        Returns the unique identifier of the specified configured table.

        Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
        """
        return pulumi.get(self, "configured_table_identifier")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description for the configured table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A name for the configured table.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
        """
        return pulumi.get(self, "tags")


class AwaitableGetConfiguredTableResult(GetConfiguredTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfiguredTableResult(
            analysis_rules=self.analysis_rules,
            arn=self.arn,
            configured_table_identifier=self.configured_table_identifier,
            description=self.description,
            name=self.name,
            tags=self.tags)


def get_configured_table(configured_table_identifier: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfiguredTableResult:
    """
    Represents a table that can be associated with collaborations


    :param str configured_table_identifier: Returns the unique identifier of the specified configured table.
           
           Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
    """
    __args__ = dict()
    __args__['configuredTableIdentifier'] = configured_table_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cleanrooms:getConfiguredTable', __args__, opts=opts, typ=GetConfiguredTableResult).value

    return AwaitableGetConfiguredTableResult(
        analysis_rules=pulumi.get(__ret__, 'analysis_rules'),
        arn=pulumi.get(__ret__, 'arn'),
        configured_table_identifier=pulumi.get(__ret__, 'configured_table_identifier'),
        description=pulumi.get(__ret__, 'description'),
        name=pulumi.get(__ret__, 'name'),
        tags=pulumi.get(__ret__, 'tags'))
def get_configured_table_output(configured_table_identifier: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetConfiguredTableResult]:
    """
    Represents a table that can be associated with collaborations


    :param str configured_table_identifier: Returns the unique identifier of the specified configured table.
           
           Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`
    """
    __args__ = dict()
    __args__['configuredTableIdentifier'] = configured_table_identifier
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cleanrooms:getConfiguredTable', __args__, opts=opts, typ=GetConfiguredTableResult)
    return __ret__.apply(lambda __response__: GetConfiguredTableResult(
        analysis_rules=pulumi.get(__response__, 'analysis_rules'),
        arn=pulumi.get(__response__, 'arn'),
        configured_table_identifier=pulumi.get(__response__, 'configured_table_identifier'),
        description=pulumi.get(__response__, 'description'),
        name=pulumi.get(__response__, 'name'),
        tags=pulumi.get(__response__, 'tags')))
