# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPrivateGraphEndpointResult',
    'AwaitableGetPrivateGraphEndpointResult',
    'get_private_graph_endpoint',
    'get_private_graph_endpoint_output',
]

@pulumi.output_type
class GetPrivateGraphEndpointResult:
    def __init__(__self__, private_graph_endpoint_identifier=None, vpc_endpoint_id=None):
        if private_graph_endpoint_identifier and not isinstance(private_graph_endpoint_identifier, str):
            raise TypeError("Expected argument 'private_graph_endpoint_identifier' to be a str")
        pulumi.set(__self__, "private_graph_endpoint_identifier", private_graph_endpoint_identifier)
        if vpc_endpoint_id and not isinstance(vpc_endpoint_id, str):
            raise TypeError("Expected argument 'vpc_endpoint_id' to be a str")
        pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)

    @property
    @pulumi.getter(name="privateGraphEndpointIdentifier")
    def private_graph_endpoint_identifier(self) -> Optional[str]:
        """
        PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.

         For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
        """
        return pulumi.get(self, "private_graph_endpoint_identifier")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[str]:
        """
        VPC endpoint that provides a private connection between the Graph and specified VPC.
        """
        return pulumi.get(self, "vpc_endpoint_id")


class AwaitableGetPrivateGraphEndpointResult(GetPrivateGraphEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateGraphEndpointResult(
            private_graph_endpoint_identifier=self.private_graph_endpoint_identifier,
            vpc_endpoint_id=self.vpc_endpoint_id)


def get_private_graph_endpoint(private_graph_endpoint_identifier: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateGraphEndpointResult:
    """
    The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.


    :param str private_graph_endpoint_identifier: PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.
           
            For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
    """
    __args__ = dict()
    __args__['privateGraphEndpointIdentifier'] = private_graph_endpoint_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:neptunegraph:getPrivateGraphEndpoint', __args__, opts=opts, typ=GetPrivateGraphEndpointResult).value

    return AwaitableGetPrivateGraphEndpointResult(
        private_graph_endpoint_identifier=pulumi.get(__ret__, 'private_graph_endpoint_identifier'),
        vpc_endpoint_id=pulumi.get(__ret__, 'vpc_endpoint_id'))


@_utilities.lift_output_func(get_private_graph_endpoint)
def get_private_graph_endpoint_output(private_graph_endpoint_identifier: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateGraphEndpointResult]:
    """
    The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.


    :param str private_graph_endpoint_identifier: PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.
           
            For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
    """
    ...