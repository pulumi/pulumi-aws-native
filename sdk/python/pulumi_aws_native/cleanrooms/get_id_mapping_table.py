# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetIdMappingTableResult',
    'AwaitableGetIdMappingTableResult',
    'get_id_mapping_table',
    'get_id_mapping_table_output',
]

@pulumi.output_type
class GetIdMappingTableResult:
    def __init__(__self__, arn=None, collaboration_arn=None, collaboration_identifier=None, description=None, id_mapping_table_identifier=None, input_reference_properties=None, kms_key_arn=None, membership_arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if collaboration_arn and not isinstance(collaboration_arn, str):
            raise TypeError("Expected argument 'collaboration_arn' to be a str")
        pulumi.set(__self__, "collaboration_arn", collaboration_arn)
        if collaboration_identifier and not isinstance(collaboration_identifier, str):
            raise TypeError("Expected argument 'collaboration_identifier' to be a str")
        pulumi.set(__self__, "collaboration_identifier", collaboration_identifier)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id_mapping_table_identifier and not isinstance(id_mapping_table_identifier, str):
            raise TypeError("Expected argument 'id_mapping_table_identifier' to be a str")
        pulumi.set(__self__, "id_mapping_table_identifier", id_mapping_table_identifier)
        if input_reference_properties and not isinstance(input_reference_properties, dict):
            raise TypeError("Expected argument 'input_reference_properties' to be a dict")
        pulumi.set(__self__, "input_reference_properties", input_reference_properties)
        if kms_key_arn and not isinstance(kms_key_arn, str):
            raise TypeError("Expected argument 'kms_key_arn' to be a str")
        pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if membership_arn and not isinstance(membership_arn, str):
            raise TypeError("Expected argument 'membership_arn' to be a str")
        pulumi.set(__self__, "membership_arn", membership_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the ID mapping table.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="collaborationArn")
    def collaboration_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the collaboration that contains this ID mapping table.
        """
        return pulumi.get(self, "collaboration_arn")

    @property
    @pulumi.getter(name="collaborationIdentifier")
    def collaboration_identifier(self) -> Optional[str]:
        return pulumi.get(self, "collaboration_identifier")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the ID mapping table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="idMappingTableIdentifier")
    def id_mapping_table_identifier(self) -> Optional[str]:
        return pulumi.get(self, "id_mapping_table_identifier")

    @property
    @pulumi.getter(name="inputReferenceProperties")
    def input_reference_properties(self) -> Optional['outputs.IdMappingTableInputReferenceProperties']:
        return pulumi.get(self, "input_reference_properties")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the AWS KMS key.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter(name="membershipArn")
    def membership_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the membership resource for the ID mapping table.
        """
        return pulumi.get(self, "membership_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        An optional label that you can assign to a resource when you create it. Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetIdMappingTableResult(GetIdMappingTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdMappingTableResult(
            arn=self.arn,
            collaboration_arn=self.collaboration_arn,
            collaboration_identifier=self.collaboration_identifier,
            description=self.description,
            id_mapping_table_identifier=self.id_mapping_table_identifier,
            input_reference_properties=self.input_reference_properties,
            kms_key_arn=self.kms_key_arn,
            membership_arn=self.membership_arn,
            tags=self.tags)


def get_id_mapping_table(id_mapping_table_identifier: Optional[str] = None,
                         membership_identifier: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdMappingTableResult:
    """
    Represents an association between an ID mapping workflow and a collaboration


    :param str membership_identifier: The unique identifier of the membership resource for the ID mapping table.
    """
    __args__ = dict()
    __args__['idMappingTableIdentifier'] = id_mapping_table_identifier
    __args__['membershipIdentifier'] = membership_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cleanrooms:getIdMappingTable', __args__, opts=opts, typ=GetIdMappingTableResult).value

    return AwaitableGetIdMappingTableResult(
        arn=pulumi.get(__ret__, 'arn'),
        collaboration_arn=pulumi.get(__ret__, 'collaboration_arn'),
        collaboration_identifier=pulumi.get(__ret__, 'collaboration_identifier'),
        description=pulumi.get(__ret__, 'description'),
        id_mapping_table_identifier=pulumi.get(__ret__, 'id_mapping_table_identifier'),
        input_reference_properties=pulumi.get(__ret__, 'input_reference_properties'),
        kms_key_arn=pulumi.get(__ret__, 'kms_key_arn'),
        membership_arn=pulumi.get(__ret__, 'membership_arn'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_id_mapping_table)
def get_id_mapping_table_output(id_mapping_table_identifier: Optional[pulumi.Input[str]] = None,
                                membership_identifier: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdMappingTableResult]:
    """
    Represents an association between an ID mapping workflow and a collaboration


    :param str membership_identifier: The unique identifier of the membership resource for the ID mapping table.
    """
    ...