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
    'GetSourceApiAssociationResult',
    'AwaitableGetSourceApiAssociationResult',
    'get_source_api_association',
    'get_source_api_association_output',
]

@pulumi.output_type
class GetSourceApiAssociationResult:
    def __init__(__self__, association_arn=None, association_id=None, description=None, last_successful_merge_date=None, merged_api_arn=None, merged_api_id=None, source_api_arn=None, source_api_association_config=None, source_api_association_status=None, source_api_association_status_detail=None, source_api_id=None):
        if association_arn and not isinstance(association_arn, str):
            raise TypeError("Expected argument 'association_arn' to be a str")
        pulumi.set(__self__, "association_arn", association_arn)
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if last_successful_merge_date and not isinstance(last_successful_merge_date, str):
            raise TypeError("Expected argument 'last_successful_merge_date' to be a str")
        pulumi.set(__self__, "last_successful_merge_date", last_successful_merge_date)
        if merged_api_arn and not isinstance(merged_api_arn, str):
            raise TypeError("Expected argument 'merged_api_arn' to be a str")
        pulumi.set(__self__, "merged_api_arn", merged_api_arn)
        if merged_api_id and not isinstance(merged_api_id, str):
            raise TypeError("Expected argument 'merged_api_id' to be a str")
        pulumi.set(__self__, "merged_api_id", merged_api_id)
        if source_api_arn and not isinstance(source_api_arn, str):
            raise TypeError("Expected argument 'source_api_arn' to be a str")
        pulumi.set(__self__, "source_api_arn", source_api_arn)
        if source_api_association_config and not isinstance(source_api_association_config, dict):
            raise TypeError("Expected argument 'source_api_association_config' to be a dict")
        pulumi.set(__self__, "source_api_association_config", source_api_association_config)
        if source_api_association_status and not isinstance(source_api_association_status, str):
            raise TypeError("Expected argument 'source_api_association_status' to be a str")
        pulumi.set(__self__, "source_api_association_status", source_api_association_status)
        if source_api_association_status_detail and not isinstance(source_api_association_status_detail, str):
            raise TypeError("Expected argument 'source_api_association_status_detail' to be a str")
        pulumi.set(__self__, "source_api_association_status_detail", source_api_association_status_detail)
        if source_api_id and not isinstance(source_api_id, str):
            raise TypeError("Expected argument 'source_api_id' to be a str")
        pulumi.set(__self__, "source_api_id", source_api_id)

    @property
    @pulumi.getter(name="associationArn")
    def association_arn(self) -> Optional[builtins.str]:
        """
        ARN of the SourceApiAssociation.
        """
        return pulumi.get(self, "association_arn")

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> Optional[builtins.str]:
        """
        Id of the SourceApiAssociation.
        """
        return pulumi.get(self, "association_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        Description of the SourceApiAssociation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastSuccessfulMergeDate")
    def last_successful_merge_date(self) -> Optional[builtins.str]:
        """
        Date of last schema successful merge.
        """
        return pulumi.get(self, "last_successful_merge_date")

    @property
    @pulumi.getter(name="mergedApiArn")
    def merged_api_arn(self) -> Optional[builtins.str]:
        """
        ARN of the Merged API in the association.
        """
        return pulumi.get(self, "merged_api_arn")

    @property
    @pulumi.getter(name="mergedApiId")
    def merged_api_id(self) -> Optional[builtins.str]:
        """
        GraphQLApiId of the Merged API in the association.
        """
        return pulumi.get(self, "merged_api_id")

    @property
    @pulumi.getter(name="sourceApiArn")
    def source_api_arn(self) -> Optional[builtins.str]:
        """
        ARN of the source API in the association.
        """
        return pulumi.get(self, "source_api_arn")

    @property
    @pulumi.getter(name="sourceApiAssociationConfig")
    def source_api_association_config(self) -> Optional['outputs.SourceApiAssociationConfig']:
        """
        Customized configuration for SourceApiAssociation.
        """
        return pulumi.get(self, "source_api_association_config")

    @property
    @pulumi.getter(name="sourceApiAssociationStatus")
    def source_api_association_status(self) -> Optional['SourceApiAssociationStatus']:
        """
        Current status of SourceApiAssociation.
        """
        return pulumi.get(self, "source_api_association_status")

    @property
    @pulumi.getter(name="sourceApiAssociationStatusDetail")
    def source_api_association_status_detail(self) -> Optional[builtins.str]:
        """
        Current SourceApiAssociation status details.
        """
        return pulumi.get(self, "source_api_association_status_detail")

    @property
    @pulumi.getter(name="sourceApiId")
    def source_api_id(self) -> Optional[builtins.str]:
        """
        GraphQLApiId of the source API in the association.
        """
        return pulumi.get(self, "source_api_id")


class AwaitableGetSourceApiAssociationResult(GetSourceApiAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourceApiAssociationResult(
            association_arn=self.association_arn,
            association_id=self.association_id,
            description=self.description,
            last_successful_merge_date=self.last_successful_merge_date,
            merged_api_arn=self.merged_api_arn,
            merged_api_id=self.merged_api_id,
            source_api_arn=self.source_api_arn,
            source_api_association_config=self.source_api_association_config,
            source_api_association_status=self.source_api_association_status,
            source_api_association_status_detail=self.source_api_association_status_detail,
            source_api_id=self.source_api_id)


def get_source_api_association(association_arn: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourceApiAssociationResult:
    """
    Resource Type definition for AWS::AppSync::SourceApiAssociation


    :param builtins.str association_arn: ARN of the SourceApiAssociation.
    """
    __args__ = dict()
    __args__['associationArn'] = association_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:appsync:getSourceApiAssociation', __args__, opts=opts, typ=GetSourceApiAssociationResult).value

    return AwaitableGetSourceApiAssociationResult(
        association_arn=pulumi.get(__ret__, 'association_arn'),
        association_id=pulumi.get(__ret__, 'association_id'),
        description=pulumi.get(__ret__, 'description'),
        last_successful_merge_date=pulumi.get(__ret__, 'last_successful_merge_date'),
        merged_api_arn=pulumi.get(__ret__, 'merged_api_arn'),
        merged_api_id=pulumi.get(__ret__, 'merged_api_id'),
        source_api_arn=pulumi.get(__ret__, 'source_api_arn'),
        source_api_association_config=pulumi.get(__ret__, 'source_api_association_config'),
        source_api_association_status=pulumi.get(__ret__, 'source_api_association_status'),
        source_api_association_status_detail=pulumi.get(__ret__, 'source_api_association_status_detail'),
        source_api_id=pulumi.get(__ret__, 'source_api_id'))
def get_source_api_association_output(association_arn: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSourceApiAssociationResult]:
    """
    Resource Type definition for AWS::AppSync::SourceApiAssociation


    :param builtins.str association_arn: ARN of the SourceApiAssociation.
    """
    __args__ = dict()
    __args__['associationArn'] = association_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:appsync:getSourceApiAssociation', __args__, opts=opts, typ=GetSourceApiAssociationResult)
    return __ret__.apply(lambda __response__: GetSourceApiAssociationResult(
        association_arn=pulumi.get(__response__, 'association_arn'),
        association_id=pulumi.get(__response__, 'association_id'),
        description=pulumi.get(__response__, 'description'),
        last_successful_merge_date=pulumi.get(__response__, 'last_successful_merge_date'),
        merged_api_arn=pulumi.get(__response__, 'merged_api_arn'),
        merged_api_id=pulumi.get(__response__, 'merged_api_id'),
        source_api_arn=pulumi.get(__response__, 'source_api_arn'),
        source_api_association_config=pulumi.get(__response__, 'source_api_association_config'),
        source_api_association_status=pulumi.get(__response__, 'source_api_association_status'),
        source_api_association_status_detail=pulumi.get(__response__, 'source_api_association_status_detail'),
        source_api_id=pulumi.get(__response__, 'source_api_id')))
