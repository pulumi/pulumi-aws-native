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
    'GetTransitGatewayVpcAttachmentResult',
    'AwaitableGetTransitGatewayVpcAttachmentResult',
    'get_transit_gateway_vpc_attachment',
    'get_transit_gateway_vpc_attachment_output',
]

@pulumi.output_type
class GetTransitGatewayVpcAttachmentResult:
    def __init__(__self__, add_subnet_ids=None, id=None, options=None, remove_subnet_ids=None, tags=None):
        if add_subnet_ids and not isinstance(add_subnet_ids, list):
            raise TypeError("Expected argument 'add_subnet_ids' to be a list")
        pulumi.set(__self__, "add_subnet_ids", add_subnet_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if remove_subnet_ids and not isinstance(remove_subnet_ids, list):
            raise TypeError("Expected argument 'remove_subnet_ids' to be a list")
        pulumi.set(__self__, "remove_subnet_ids", remove_subnet_ids)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="addSubnetIds")
    def add_subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "add_subnet_ids")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def options(self) -> Optional['outputs.OptionsProperties']:
        """
        The options for the transit gateway vpc attachment.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="removeSubnetIds")
    def remove_subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "remove_subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TransitGatewayVpcAttachmentTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetTransitGatewayVpcAttachmentResult(GetTransitGatewayVpcAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitGatewayVpcAttachmentResult(
            add_subnet_ids=self.add_subnet_ids,
            id=self.id,
            options=self.options,
            remove_subnet_ids=self.remove_subnet_ids,
            tags=self.tags)


def get_transit_gateway_vpc_attachment(id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitGatewayVpcAttachmentResult:
    """
    Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getTransitGatewayVpcAttachment', __args__, opts=opts, typ=GetTransitGatewayVpcAttachmentResult).value

    return AwaitableGetTransitGatewayVpcAttachmentResult(
        add_subnet_ids=__ret__.add_subnet_ids,
        id=__ret__.id,
        options=__ret__.options,
        remove_subnet_ids=__ret__.remove_subnet_ids,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_transit_gateway_vpc_attachment)
def get_transit_gateway_vpc_attachment_output(id: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitGatewayVpcAttachmentResult]:
    """
    Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment
    """
    ...