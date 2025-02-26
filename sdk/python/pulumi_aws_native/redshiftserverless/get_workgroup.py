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
    'GetWorkgroupResult',
    'AwaitableGetWorkgroupResult',
    'get_workgroup',
    'get_workgroup_output',
]

@pulumi.output_type
class GetWorkgroupResult:
    def __init__(__self__, enhanced_vpc_routing=None, port=None, price_performance_target=None, publicly_accessible=None, tags=None, workgroup=None):
        if enhanced_vpc_routing and not isinstance(enhanced_vpc_routing, bool):
            raise TypeError("Expected argument 'enhanced_vpc_routing' to be a bool")
        pulumi.set(__self__, "enhanced_vpc_routing", enhanced_vpc_routing)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if price_performance_target and not isinstance(price_performance_target, dict):
            raise TypeError("Expected argument 'price_performance_target' to be a dict")
        pulumi.set(__self__, "price_performance_target", price_performance_target)
        if publicly_accessible and not isinstance(publicly_accessible, bool):
            raise TypeError("Expected argument 'publicly_accessible' to be a bool")
        pulumi.set(__self__, "publicly_accessible", publicly_accessible)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if workgroup and not isinstance(workgroup, dict):
            raise TypeError("Expected argument 'workgroup' to be a dict")
        pulumi.set(__self__, "workgroup", workgroup)

    @property
    @pulumi.getter(name="enhancedVpcRouting")
    def enhanced_vpc_routing(self) -> Optional[bool]:
        """
        The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
        """
        return pulumi.get(self, "enhanced_vpc_routing")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        """
        The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="pricePerformanceTarget")
    def price_performance_target(self) -> Optional['outputs.WorkgroupPerformanceTarget']:
        """
        A property that represents the price performance target settings for the workgroup.
        """
        return pulumi.get(self, "price_performance_target")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> Optional[bool]:
        """
        A value that specifies whether the workgroup can be accessible from a public network.
        """
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The map of the key-value pairs used to tag the workgroup.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def workgroup(self) -> Optional['outputs.Workgroup']:
        """
        Definition for workgroup resource
        """
        return pulumi.get(self, "workgroup")


class AwaitableGetWorkgroupResult(GetWorkgroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkgroupResult(
            enhanced_vpc_routing=self.enhanced_vpc_routing,
            port=self.port,
            price_performance_target=self.price_performance_target,
            publicly_accessible=self.publicly_accessible,
            tags=self.tags,
            workgroup=self.workgroup)


def get_workgroup(workgroup_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkgroupResult:
    """
    Definition of AWS::RedshiftServerless::Workgroup Resource Type


    :param str workgroup_name: The name of the workgroup.
    """
    __args__ = dict()
    __args__['workgroupName'] = workgroup_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:redshiftserverless:getWorkgroup', __args__, opts=opts, typ=GetWorkgroupResult).value

    return AwaitableGetWorkgroupResult(
        enhanced_vpc_routing=pulumi.get(__ret__, 'enhanced_vpc_routing'),
        port=pulumi.get(__ret__, 'port'),
        price_performance_target=pulumi.get(__ret__, 'price_performance_target'),
        publicly_accessible=pulumi.get(__ret__, 'publicly_accessible'),
        tags=pulumi.get(__ret__, 'tags'),
        workgroup=pulumi.get(__ret__, 'workgroup'))
def get_workgroup_output(workgroup_name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWorkgroupResult]:
    """
    Definition of AWS::RedshiftServerless::Workgroup Resource Type


    :param str workgroup_name: The name of the workgroup.
    """
    __args__ = dict()
    __args__['workgroupName'] = workgroup_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:redshiftserverless:getWorkgroup', __args__, opts=opts, typ=GetWorkgroupResult)
    return __ret__.apply(lambda __response__: GetWorkgroupResult(
        enhanced_vpc_routing=pulumi.get(__response__, 'enhanced_vpc_routing'),
        port=pulumi.get(__response__, 'port'),
        price_performance_target=pulumi.get(__response__, 'price_performance_target'),
        publicly_accessible=pulumi.get(__response__, 'publicly_accessible'),
        tags=pulumi.get(__response__, 'tags'),
        workgroup=pulumi.get(__response__, 'workgroup')))
