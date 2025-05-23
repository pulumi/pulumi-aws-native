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
    'GetFlowResult',
    'AwaitableGetFlowResult',
    'get_flow',
    'get_flow_output',
]

@pulumi.output_type
class GetFlowResult:
    def __init__(__self__, egress_ip=None, flow_arn=None, flow_availability_zone=None, flow_ndi_machine_name=None, flow_size=None, maintenance=None, media_streams=None, ndi_config=None, source=None, source_failover_config=None, source_monitoring_config=None, vpc_interfaces=None):
        if egress_ip and not isinstance(egress_ip, str):
            raise TypeError("Expected argument 'egress_ip' to be a str")
        pulumi.set(__self__, "egress_ip", egress_ip)
        if flow_arn and not isinstance(flow_arn, str):
            raise TypeError("Expected argument 'flow_arn' to be a str")
        pulumi.set(__self__, "flow_arn", flow_arn)
        if flow_availability_zone and not isinstance(flow_availability_zone, str):
            raise TypeError("Expected argument 'flow_availability_zone' to be a str")
        pulumi.set(__self__, "flow_availability_zone", flow_availability_zone)
        if flow_ndi_machine_name and not isinstance(flow_ndi_machine_name, str):
            raise TypeError("Expected argument 'flow_ndi_machine_name' to be a str")
        pulumi.set(__self__, "flow_ndi_machine_name", flow_ndi_machine_name)
        if flow_size and not isinstance(flow_size, str):
            raise TypeError("Expected argument 'flow_size' to be a str")
        pulumi.set(__self__, "flow_size", flow_size)
        if maintenance and not isinstance(maintenance, dict):
            raise TypeError("Expected argument 'maintenance' to be a dict")
        pulumi.set(__self__, "maintenance", maintenance)
        if media_streams and not isinstance(media_streams, list):
            raise TypeError("Expected argument 'media_streams' to be a list")
        pulumi.set(__self__, "media_streams", media_streams)
        if ndi_config and not isinstance(ndi_config, dict):
            raise TypeError("Expected argument 'ndi_config' to be a dict")
        pulumi.set(__self__, "ndi_config", ndi_config)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)
        if source_failover_config and not isinstance(source_failover_config, dict):
            raise TypeError("Expected argument 'source_failover_config' to be a dict")
        pulumi.set(__self__, "source_failover_config", source_failover_config)
        if source_monitoring_config and not isinstance(source_monitoring_config, dict):
            raise TypeError("Expected argument 'source_monitoring_config' to be a dict")
        pulumi.set(__self__, "source_monitoring_config", source_monitoring_config)
        if vpc_interfaces and not isinstance(vpc_interfaces, list):
            raise TypeError("Expected argument 'vpc_interfaces' to be a list")
        pulumi.set(__self__, "vpc_interfaces", vpc_interfaces)

    @property
    @pulumi.getter(name="egressIp")
    def egress_ip(self) -> Optional[builtins.str]:
        """
        The IP address from which video will be sent to output destinations.
        """
        return pulumi.get(self, "egress_ip")

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> Optional[builtins.str]:
        """
        The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter(name="flowAvailabilityZone")
    def flow_availability_zone(self) -> Optional[builtins.str]:
        """
        The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)
        """
        return pulumi.get(self, "flow_availability_zone")

    @property
    @pulumi.getter(name="flowNdiMachineName")
    def flow_ndi_machine_name(self) -> Optional[builtins.str]:
        """
        A prefix for the names of the NDI sources that the flow creates.(ReadOnly)
        """
        return pulumi.get(self, "flow_ndi_machine_name")

    @property
    @pulumi.getter(name="flowSize")
    def flow_size(self) -> Optional['FlowSize']:
        """
        Determines the processing capacity and feature set of the flow. Set this optional parameter to LARGE if you want to enable NDI outputs on the flow.
        """
        return pulumi.get(self, "flow_size")

    @property
    @pulumi.getter
    def maintenance(self) -> Optional['outputs.FlowMaintenance']:
        """
        The maintenance settings you want to use for the flow.
        """
        return pulumi.get(self, "maintenance")

    @property
    @pulumi.getter(name="mediaStreams")
    def media_streams(self) -> Optional[Sequence['outputs.FlowMediaStream']]:
        """
        The media streams associated with the flow. You can associate any of these media streams with sources and outputs on the flow.
        """
        return pulumi.get(self, "media_streams")

    @property
    @pulumi.getter(name="ndiConfig")
    def ndi_config(self) -> Optional['outputs.FlowNdiConfig']:
        """
        Specifies the configuration settings for NDI outputs. Required when the flow includes NDI outputs.
        """
        return pulumi.get(self, "ndi_config")

    @property
    @pulumi.getter
    def source(self) -> Optional['outputs.FlowSource']:
        """
        The source of the flow.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceFailoverConfig")
    def source_failover_config(self) -> Optional['outputs.FlowFailoverConfig']:
        """
        The source failover config of the flow.
        """
        return pulumi.get(self, "source_failover_config")

    @property
    @pulumi.getter(name="sourceMonitoringConfig")
    def source_monitoring_config(self) -> Optional['outputs.FlowSourceMonitoringConfig']:
        """
        The source monitoring config of the flow.
        """
        return pulumi.get(self, "source_monitoring_config")

    @property
    @pulumi.getter(name="vpcInterfaces")
    def vpc_interfaces(self) -> Optional[Sequence['outputs.FlowVpcInterface']]:
        """
        The VPC interfaces that you added to this flow.
        """
        return pulumi.get(self, "vpc_interfaces")


class AwaitableGetFlowResult(GetFlowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlowResult(
            egress_ip=self.egress_ip,
            flow_arn=self.flow_arn,
            flow_availability_zone=self.flow_availability_zone,
            flow_ndi_machine_name=self.flow_ndi_machine_name,
            flow_size=self.flow_size,
            maintenance=self.maintenance,
            media_streams=self.media_streams,
            ndi_config=self.ndi_config,
            source=self.source,
            source_failover_config=self.source_failover_config,
            source_monitoring_config=self.source_monitoring_config,
            vpc_interfaces=self.vpc_interfaces)


def get_flow(flow_arn: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlowResult:
    """
    Resource schema for AWS::MediaConnect::Flow


    :param builtins.str flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
    """
    __args__ = dict()
    __args__['flowArn'] = flow_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:mediaconnect:getFlow', __args__, opts=opts, typ=GetFlowResult).value

    return AwaitableGetFlowResult(
        egress_ip=pulumi.get(__ret__, 'egress_ip'),
        flow_arn=pulumi.get(__ret__, 'flow_arn'),
        flow_availability_zone=pulumi.get(__ret__, 'flow_availability_zone'),
        flow_ndi_machine_name=pulumi.get(__ret__, 'flow_ndi_machine_name'),
        flow_size=pulumi.get(__ret__, 'flow_size'),
        maintenance=pulumi.get(__ret__, 'maintenance'),
        media_streams=pulumi.get(__ret__, 'media_streams'),
        ndi_config=pulumi.get(__ret__, 'ndi_config'),
        source=pulumi.get(__ret__, 'source'),
        source_failover_config=pulumi.get(__ret__, 'source_failover_config'),
        source_monitoring_config=pulumi.get(__ret__, 'source_monitoring_config'),
        vpc_interfaces=pulumi.get(__ret__, 'vpc_interfaces'))
def get_flow_output(flow_arn: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFlowResult]:
    """
    Resource schema for AWS::MediaConnect::Flow


    :param builtins.str flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
    """
    __args__ = dict()
    __args__['flowArn'] = flow_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:mediaconnect:getFlow', __args__, opts=opts, typ=GetFlowResult)
    return __ret__.apply(lambda __response__: GetFlowResult(
        egress_ip=pulumi.get(__response__, 'egress_ip'),
        flow_arn=pulumi.get(__response__, 'flow_arn'),
        flow_availability_zone=pulumi.get(__response__, 'flow_availability_zone'),
        flow_ndi_machine_name=pulumi.get(__response__, 'flow_ndi_machine_name'),
        flow_size=pulumi.get(__response__, 'flow_size'),
        maintenance=pulumi.get(__response__, 'maintenance'),
        media_streams=pulumi.get(__response__, 'media_streams'),
        ndi_config=pulumi.get(__response__, 'ndi_config'),
        source=pulumi.get(__response__, 'source'),
        source_failover_config=pulumi.get(__response__, 'source_failover_config'),
        source_monitoring_config=pulumi.get(__response__, 'source_monitoring_config'),
        vpc_interfaces=pulumi.get(__response__, 'vpc_interfaces')))
