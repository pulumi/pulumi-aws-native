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
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    def __init__(__self__, capacity_provider_strategy=None, deployment_configuration=None, desired_count=None, enable_ecs_managed_tags=None, enable_execute_command=None, health_check_grace_period_seconds=None, load_balancers=None, name=None, network_configuration=None, placement_constraints=None, placement_strategies=None, platform_version=None, propagate_tags=None, service_arn=None, service_registries=None, tags=None, task_definition=None):
        if capacity_provider_strategy and not isinstance(capacity_provider_strategy, list):
            raise TypeError("Expected argument 'capacity_provider_strategy' to be a list")
        pulumi.set(__self__, "capacity_provider_strategy", capacity_provider_strategy)
        if deployment_configuration and not isinstance(deployment_configuration, dict):
            raise TypeError("Expected argument 'deployment_configuration' to be a dict")
        pulumi.set(__self__, "deployment_configuration", deployment_configuration)
        if desired_count and not isinstance(desired_count, int):
            raise TypeError("Expected argument 'desired_count' to be a int")
        pulumi.set(__self__, "desired_count", desired_count)
        if enable_ecs_managed_tags and not isinstance(enable_ecs_managed_tags, bool):
            raise TypeError("Expected argument 'enable_ecs_managed_tags' to be a bool")
        pulumi.set(__self__, "enable_ecs_managed_tags", enable_ecs_managed_tags)
        if enable_execute_command and not isinstance(enable_execute_command, bool):
            raise TypeError("Expected argument 'enable_execute_command' to be a bool")
        pulumi.set(__self__, "enable_execute_command", enable_execute_command)
        if health_check_grace_period_seconds and not isinstance(health_check_grace_period_seconds, int):
            raise TypeError("Expected argument 'health_check_grace_period_seconds' to be a int")
        pulumi.set(__self__, "health_check_grace_period_seconds", health_check_grace_period_seconds)
        if load_balancers and not isinstance(load_balancers, list):
            raise TypeError("Expected argument 'load_balancers' to be a list")
        pulumi.set(__self__, "load_balancers", load_balancers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_configuration and not isinstance(network_configuration, dict):
            raise TypeError("Expected argument 'network_configuration' to be a dict")
        pulumi.set(__self__, "network_configuration", network_configuration)
        if placement_constraints and not isinstance(placement_constraints, list):
            raise TypeError("Expected argument 'placement_constraints' to be a list")
        pulumi.set(__self__, "placement_constraints", placement_constraints)
        if placement_strategies and not isinstance(placement_strategies, list):
            raise TypeError("Expected argument 'placement_strategies' to be a list")
        pulumi.set(__self__, "placement_strategies", placement_strategies)
        if platform_version and not isinstance(platform_version, str):
            raise TypeError("Expected argument 'platform_version' to be a str")
        pulumi.set(__self__, "platform_version", platform_version)
        if propagate_tags and not isinstance(propagate_tags, str):
            raise TypeError("Expected argument 'propagate_tags' to be a str")
        pulumi.set(__self__, "propagate_tags", propagate_tags)
        if service_arn and not isinstance(service_arn, str):
            raise TypeError("Expected argument 'service_arn' to be a str")
        pulumi.set(__self__, "service_arn", service_arn)
        if service_registries and not isinstance(service_registries, list):
            raise TypeError("Expected argument 'service_registries' to be a list")
        pulumi.set(__self__, "service_registries", service_registries)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if task_definition and not isinstance(task_definition, str):
            raise TypeError("Expected argument 'task_definition' to be a str")
        pulumi.set(__self__, "task_definition", task_definition)

    @property
    @pulumi.getter(name="capacityProviderStrategy")
    def capacity_provider_strategy(self) -> Optional[Sequence['outputs.ServiceCapacityProviderStrategyItem']]:
        return pulumi.get(self, "capacity_provider_strategy")

    @property
    @pulumi.getter(name="deploymentConfiguration")
    def deployment_configuration(self) -> Optional['outputs.ServiceDeploymentConfiguration']:
        return pulumi.get(self, "deployment_configuration")

    @property
    @pulumi.getter(name="desiredCount")
    def desired_count(self) -> Optional[int]:
        return pulumi.get(self, "desired_count")

    @property
    @pulumi.getter(name="enableECSManagedTags")
    def enable_ecs_managed_tags(self) -> Optional[bool]:
        return pulumi.get(self, "enable_ecs_managed_tags")

    @property
    @pulumi.getter(name="enableExecuteCommand")
    def enable_execute_command(self) -> Optional[bool]:
        return pulumi.get(self, "enable_execute_command")

    @property
    @pulumi.getter(name="healthCheckGracePeriodSeconds")
    def health_check_grace_period_seconds(self) -> Optional[int]:
        return pulumi.get(self, "health_check_grace_period_seconds")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Optional[Sequence['outputs.ServiceLoadBalancer']]:
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> Optional['outputs.ServiceNetworkConfiguration']:
        return pulumi.get(self, "network_configuration")

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> Optional[Sequence['outputs.ServicePlacementConstraint']]:
        return pulumi.get(self, "placement_constraints")

    @property
    @pulumi.getter(name="placementStrategies")
    def placement_strategies(self) -> Optional[Sequence['outputs.ServicePlacementStrategy']]:
        return pulumi.get(self, "placement_strategies")

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> Optional[str]:
        return pulumi.get(self, "platform_version")

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> Optional['ServicePropagateTags']:
        return pulumi.get(self, "propagate_tags")

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> Optional[str]:
        return pulumi.get(self, "service_arn")

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> Optional[Sequence['outputs.ServiceRegistry']]:
        return pulumi.get(self, "service_registries")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ServiceTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> Optional[str]:
        return pulumi.get(self, "task_definition")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            capacity_provider_strategy=self.capacity_provider_strategy,
            deployment_configuration=self.deployment_configuration,
            desired_count=self.desired_count,
            enable_ecs_managed_tags=self.enable_ecs_managed_tags,
            enable_execute_command=self.enable_execute_command,
            health_check_grace_period_seconds=self.health_check_grace_period_seconds,
            load_balancers=self.load_balancers,
            name=self.name,
            network_configuration=self.network_configuration,
            placement_constraints=self.placement_constraints,
            placement_strategies=self.placement_strategies,
            platform_version=self.platform_version,
            propagate_tags=self.propagate_tags,
            service_arn=self.service_arn,
            service_registries=self.service_registries,
            tags=self.tags,
            task_definition=self.task_definition)


def get_service(cluster: Optional[str] = None,
                service_arn: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Resource Type definition for AWS::ECS::Service
    """
    __args__ = dict()
    __args__['cluster'] = cluster
    __args__['serviceArn'] = service_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ecs:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        capacity_provider_strategy=__ret__.capacity_provider_strategy,
        deployment_configuration=__ret__.deployment_configuration,
        desired_count=__ret__.desired_count,
        enable_ecs_managed_tags=__ret__.enable_ecs_managed_tags,
        enable_execute_command=__ret__.enable_execute_command,
        health_check_grace_period_seconds=__ret__.health_check_grace_period_seconds,
        load_balancers=__ret__.load_balancers,
        name=__ret__.name,
        network_configuration=__ret__.network_configuration,
        placement_constraints=__ret__.placement_constraints,
        placement_strategies=__ret__.placement_strategies,
        platform_version=__ret__.platform_version,
        propagate_tags=__ret__.propagate_tags,
        service_arn=__ret__.service_arn,
        service_registries=__ret__.service_registries,
        tags=__ret__.tags,
        task_definition=__ret__.task_definition)


@_utilities.lift_output_func(get_service)
def get_service_output(cluster: Optional[pulumi.Input[str]] = None,
                       service_arn: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    Resource Type definition for AWS::ECS::Service
    """
    ...