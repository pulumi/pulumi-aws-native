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

__all__ = [
    'GetVpcAttachmentResult',
    'AwaitableGetVpcAttachmentResult',
    'get_vpc_attachment',
    'get_vpc_attachment_output',
]

@pulumi.output_type
class GetVpcAttachmentResult:
    def __init__(__self__, attachment_id=None, attachment_policy_rule_number=None, attachment_type=None, core_network_arn=None, created_at=None, edge_location=None, options=None, owner_account_id=None, proposed_segment_change=None, resource_arn=None, segment_name=None, state=None, subnet_arns=None, tags=None, updated_at=None):
        if attachment_id and not isinstance(attachment_id, str):
            raise TypeError("Expected argument 'attachment_id' to be a str")
        pulumi.set(__self__, "attachment_id", attachment_id)
        if attachment_policy_rule_number and not isinstance(attachment_policy_rule_number, int):
            raise TypeError("Expected argument 'attachment_policy_rule_number' to be a int")
        pulumi.set(__self__, "attachment_policy_rule_number", attachment_policy_rule_number)
        if attachment_type and not isinstance(attachment_type, str):
            raise TypeError("Expected argument 'attachment_type' to be a str")
        pulumi.set(__self__, "attachment_type", attachment_type)
        if core_network_arn and not isinstance(core_network_arn, str):
            raise TypeError("Expected argument 'core_network_arn' to be a str")
        pulumi.set(__self__, "core_network_arn", core_network_arn)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if edge_location and not isinstance(edge_location, str):
            raise TypeError("Expected argument 'edge_location' to be a str")
        pulumi.set(__self__, "edge_location", edge_location)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if owner_account_id and not isinstance(owner_account_id, str):
            raise TypeError("Expected argument 'owner_account_id' to be a str")
        pulumi.set(__self__, "owner_account_id", owner_account_id)
        if proposed_segment_change and not isinstance(proposed_segment_change, dict):
            raise TypeError("Expected argument 'proposed_segment_change' to be a dict")
        pulumi.set(__self__, "proposed_segment_change", proposed_segment_change)
        if resource_arn and not isinstance(resource_arn, str):
            raise TypeError("Expected argument 'resource_arn' to be a str")
        pulumi.set(__self__, "resource_arn", resource_arn)
        if segment_name and not isinstance(segment_name, str):
            raise TypeError("Expected argument 'segment_name' to be a str")
        pulumi.set(__self__, "segment_name", segment_name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if subnet_arns and not isinstance(subnet_arns, list):
            raise TypeError("Expected argument 'subnet_arns' to be a list")
        pulumi.set(__self__, "subnet_arns", subnet_arns)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> Optional[str]:
        """
        Id of the attachment.
        """
        return pulumi.get(self, "attachment_id")

    @property
    @pulumi.getter(name="attachmentPolicyRuleNumber")
    def attachment_policy_rule_number(self) -> Optional[int]:
        """
        The policy rule number associated with the attachment.
        """
        return pulumi.get(self, "attachment_policy_rule_number")

    @property
    @pulumi.getter(name="attachmentType")
    def attachment_type(self) -> Optional[str]:
        """
        Attachment type.
        """
        return pulumi.get(self, "attachment_type")

    @property
    @pulumi.getter(name="coreNetworkArn")
    def core_network_arn(self) -> Optional[str]:
        """
        The ARN of a core network for the VPC attachment.
        """
        return pulumi.get(self, "core_network_arn")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        Creation time of the attachment.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="edgeLocation")
    def edge_location(self) -> Optional[str]:
        """
        The Region where the edge is located.
        """
        return pulumi.get(self, "edge_location")

    @property
    @pulumi.getter
    def options(self) -> Optional['outputs.VpcAttachmentVpcOptions']:
        """
        Vpc options of the attachment.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> Optional[str]:
        """
        Owner account of the attachment.
        """
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter(name="proposedSegmentChange")
    def proposed_segment_change(self) -> Optional['outputs.VpcAttachmentProposedSegmentChange']:
        """
        The attachment to move from one segment to another.
        """
        return pulumi.get(self, "proposed_segment_change")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> Optional[str]:
        """
        The ARN of the Resource.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="segmentName")
    def segment_name(self) -> Optional[str]:
        """
        The name of the segment attachment..
        """
        return pulumi.get(self, "segment_name")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        State of the attachment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subnetArns")
    def subnet_arns(self) -> Optional[Sequence[str]]:
        """
        Subnet Arn list
        """
        return pulumi.get(self, "subnet_arns")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.VpcAttachmentTag']]:
        """
        Tags for the attachment.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        """
        Last update time of the attachment.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetVpcAttachmentResult(GetVpcAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcAttachmentResult(
            attachment_id=self.attachment_id,
            attachment_policy_rule_number=self.attachment_policy_rule_number,
            attachment_type=self.attachment_type,
            core_network_arn=self.core_network_arn,
            created_at=self.created_at,
            edge_location=self.edge_location,
            options=self.options,
            owner_account_id=self.owner_account_id,
            proposed_segment_change=self.proposed_segment_change,
            resource_arn=self.resource_arn,
            segment_name=self.segment_name,
            state=self.state,
            subnet_arns=self.subnet_arns,
            tags=self.tags,
            updated_at=self.updated_at)


def get_vpc_attachment(attachment_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcAttachmentResult:
    """
    AWS::NetworkManager::VpcAttachment Resoruce Type


    :param str attachment_id: Id of the attachment.
    """
    __args__ = dict()
    __args__['attachmentId'] = attachment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:networkmanager:getVpcAttachment', __args__, opts=opts, typ=GetVpcAttachmentResult).value

    return AwaitableGetVpcAttachmentResult(
        attachment_id=__ret__.attachment_id,
        attachment_policy_rule_number=__ret__.attachment_policy_rule_number,
        attachment_type=__ret__.attachment_type,
        core_network_arn=__ret__.core_network_arn,
        created_at=__ret__.created_at,
        edge_location=__ret__.edge_location,
        options=__ret__.options,
        owner_account_id=__ret__.owner_account_id,
        proposed_segment_change=__ret__.proposed_segment_change,
        resource_arn=__ret__.resource_arn,
        segment_name=__ret__.segment_name,
        state=__ret__.state,
        subnet_arns=__ret__.subnet_arns,
        tags=__ret__.tags,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_vpc_attachment)
def get_vpc_attachment_output(attachment_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcAttachmentResult]:
    """
    AWS::NetworkManager::VpcAttachment Resoruce Type


    :param str attachment_id: Id of the attachment.
    """
    ...