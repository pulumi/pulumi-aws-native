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
from ._inputs import *

__all__ = ['EndpointGroupArgs', 'EndpointGroup']

@pulumi.input_type
class EndpointGroupArgs:
    def __init__(__self__, *,
                 endpoint_group_region: pulumi.Input[str],
                 listener_arn: pulumi.Input[str],
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input['EndpointGroupHealthCheckProtocol']] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_dial_percentage: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a EndpointGroup resource.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener
        :param pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]] endpoint_configurations: The list of endpoint objects.
        :param pulumi.Input[int] health_check_interval_seconds: The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
        :param pulumi.Input[int] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        :param pulumi.Input['EndpointGroupHealthCheckProtocol'] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        :param pulumi.Input[int] threshold_count: The number of consecutive health checks required to set the state of the endpoint to unhealthy.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to sent to an AWS Region
        """
        pulumi.set(__self__, "endpoint_group_region", endpoint_group_region)
        pulumi.set(__self__, "listener_arn", listener_arn)
        if endpoint_configurations is not None:
            pulumi.set(__self__, "endpoint_configurations", endpoint_configurations)
        if health_check_interval_seconds is not None:
            pulumi.set(__self__, "health_check_interval_seconds", health_check_interval_seconds)
        if health_check_path is not None:
            pulumi.set(__self__, "health_check_path", health_check_path)
        if health_check_port is not None:
            pulumi.set(__self__, "health_check_port", health_check_port)
        if health_check_protocol is not None:
            pulumi.set(__self__, "health_check_protocol", health_check_protocol)
        if port_overrides is not None:
            pulumi.set(__self__, "port_overrides", port_overrides)
        if threshold_count is not None:
            pulumi.set(__self__, "threshold_count", threshold_count)
        if traffic_dial_percentage is not None:
            pulumi.set(__self__, "traffic_dial_percentage", traffic_dial_percentage)

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Input[str]:
        """
        The name of the AWS Region where the endpoint group is located
        """
        return pulumi.get(self, "endpoint_group_region")

    @endpoint_group_region.setter
    def endpoint_group_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_group_region", value)

    @property
    @pulumi.getter(name="listenerArn")
    def listener_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the listener
        """
        return pulumi.get(self, "listener_arn")

    @listener_arn.setter
    def listener_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_arn", value)

    @property
    @pulumi.getter(name="endpointConfigurations")
    def endpoint_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]]:
        """
        The list of endpoint objects.
        """
        return pulumi.get(self, "endpoint_configurations")

    @endpoint_configurations.setter
    def endpoint_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupEndpointConfigurationArgs']]]]):
        pulumi.set(self, "endpoint_configurations", value)

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
        """
        return pulumi.get(self, "health_check_interval_seconds")

    @health_check_interval_seconds.setter
    def health_check_interval_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_interval_seconds", value)

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_path")

    @health_check_path.setter
    def health_check_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_path", value)

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        """
        return pulumi.get(self, "health_check_port")

    @health_check_port.setter
    def health_check_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_port", value)

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> Optional[pulumi.Input['EndpointGroupHealthCheckProtocol']]:
        """
        The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        """
        return pulumi.get(self, "health_check_protocol")

    @health_check_protocol.setter
    def health_check_protocol(self, value: Optional[pulumi.Input['EndpointGroupHealthCheckProtocol']]):
        pulumi.set(self, "health_check_protocol", value)

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]]:
        return pulumi.get(self, "port_overrides")

    @port_overrides.setter
    def port_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointGroupPortOverrideArgs']]]]):
        pulumi.set(self, "port_overrides", value)

    @property
    @pulumi.getter(name="thresholdCount")
    def threshold_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of consecutive health checks required to set the state of the endpoint to unhealthy.
        """
        return pulumi.get(self, "threshold_count")

    @threshold_count.setter
    def threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "threshold_count", value)

    @property
    @pulumi.getter(name="trafficDialPercentage")
    def traffic_dial_percentage(self) -> Optional[pulumi.Input[float]]:
        """
        The percentage of traffic to sent to an AWS Region
        """
        return pulumi.get(self, "traffic_dial_percentage")

    @traffic_dial_percentage.setter
    def traffic_dial_percentage(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "traffic_dial_percentage", value)


class EndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input['EndpointGroupHealthCheckProtocol']] = None,
                 listener_arn: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_dial_percentage: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::GlobalAccelerator::EndpointGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]] endpoint_configurations: The list of endpoint objects.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located
        :param pulumi.Input[int] health_check_interval_seconds: The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
        :param pulumi.Input[int] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        :param pulumi.Input['EndpointGroupHealthCheckProtocol'] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener
        :param pulumi.Input[int] threshold_count: The number of consecutive health checks required to set the state of the endpoint to unhealthy.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to sent to an AWS Region
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::GlobalAccelerator::EndpointGroup

        :param str resource_name: The name of the resource.
        :param EndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupEndpointConfigurationArgs']]]]] = None,
                 endpoint_group_region: Optional[pulumi.Input[str]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[int]] = None,
                 health_check_protocol: Optional[pulumi.Input['EndpointGroupHealthCheckProtocol']] = None,
                 listener_arn: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointGroupPortOverrideArgs']]]]] = None,
                 threshold_count: Optional[pulumi.Input[int]] = None,
                 traffic_dial_percentage: Optional[pulumi.Input[float]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointGroupArgs.__new__(EndpointGroupArgs)

            __props__.__dict__["endpoint_configurations"] = endpoint_configurations
            if endpoint_group_region is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_group_region'")
            __props__.__dict__["endpoint_group_region"] = endpoint_group_region
            __props__.__dict__["health_check_interval_seconds"] = health_check_interval_seconds
            __props__.__dict__["health_check_path"] = health_check_path
            __props__.__dict__["health_check_port"] = health_check_port
            __props__.__dict__["health_check_protocol"] = health_check_protocol
            if listener_arn is None and not opts.urn:
                raise TypeError("Missing required property 'listener_arn'")
            __props__.__dict__["listener_arn"] = listener_arn
            __props__.__dict__["port_overrides"] = port_overrides
            __props__.__dict__["threshold_count"] = threshold_count
            __props__.__dict__["traffic_dial_percentage"] = traffic_dial_percentage
            __props__.__dict__["endpoint_group_arn"] = None
        super(EndpointGroup, __self__).__init__(
            'aws-native:globalaccelerator:EndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EndpointGroup':
        """
        Get an existing EndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointGroupArgs.__new__(EndpointGroupArgs)

        __props__.__dict__["endpoint_configurations"] = None
        __props__.__dict__["endpoint_group_arn"] = None
        __props__.__dict__["endpoint_group_region"] = None
        __props__.__dict__["health_check_interval_seconds"] = None
        __props__.__dict__["health_check_path"] = None
        __props__.__dict__["health_check_port"] = None
        __props__.__dict__["health_check_protocol"] = None
        __props__.__dict__["listener_arn"] = None
        __props__.__dict__["port_overrides"] = None
        __props__.__dict__["threshold_count"] = None
        __props__.__dict__["traffic_dial_percentage"] = None
        return EndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endpointConfigurations")
    def endpoint_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointGroupEndpointConfiguration']]]:
        """
        The list of endpoint objects.
        """
        return pulumi.get(self, "endpoint_configurations")

    @property
    @pulumi.getter(name="endpointGroupArn")
    def endpoint_group_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the endpoint group
        """
        return pulumi.get(self, "endpoint_group_arn")

    @property
    @pulumi.getter(name="endpointGroupRegion")
    def endpoint_group_region(self) -> pulumi.Output[str]:
        """
        The name of the AWS Region where the endpoint group is located
        """
        return pulumi.get(self, "endpoint_group_region")

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The time in seconds between each health check for an endpoint. Must be a value of 10 or 30
        """
        return pulumi.get(self, "health_check_interval_seconds")

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "health_check_path")

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> pulumi.Output[Optional[int]]:
        """
        The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        """
        return pulumi.get(self, "health_check_port")

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> pulumi.Output[Optional['EndpointGroupHealthCheckProtocol']]:
        """
        The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.
        """
        return pulumi.get(self, "health_check_protocol")

    @property
    @pulumi.getter(name="listenerArn")
    def listener_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the listener
        """
        return pulumi.get(self, "listener_arn")

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointGroupPortOverride']]]:
        return pulumi.get(self, "port_overrides")

    @property
    @pulumi.getter(name="thresholdCount")
    def threshold_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of consecutive health checks required to set the state of the endpoint to unhealthy.
        """
        return pulumi.get(self, "threshold_count")

    @property
    @pulumi.getter(name="trafficDialPercentage")
    def traffic_dial_percentage(self) -> pulumi.Output[Optional[float]]:
        """
        The percentage of traffic to sent to an AWS Region
        """
        return pulumi.get(self, "traffic_dial_percentage")
