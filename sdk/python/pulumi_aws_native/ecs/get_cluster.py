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
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    def __init__(__self__, arn=None, capacity_providers=None, cluster_settings=None, configuration=None, default_capacity_provider_strategy=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if capacity_providers and not isinstance(capacity_providers, list):
            raise TypeError("Expected argument 'capacity_providers' to be a list")
        pulumi.set(__self__, "capacity_providers", capacity_providers)
        if cluster_settings and not isinstance(cluster_settings, list):
            raise TypeError("Expected argument 'cluster_settings' to be a list")
        pulumi.set(__self__, "cluster_settings", cluster_settings)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if default_capacity_provider_strategy and not isinstance(default_capacity_provider_strategy, list):
            raise TypeError("Expected argument 'default_capacity_provider_strategy' to be a list")
        pulumi.set(__self__, "default_capacity_provider_strategy", default_capacity_provider_strategy)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon ECS cluster, such as arn:aws:ecs:us-east-2:123456789012:cluster/MyECSCluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="capacityProviders")
    def capacity_providers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "capacity_providers")

    @property
    @pulumi.getter(name="clusterSettings")
    def cluster_settings(self) -> Optional[Sequence['outputs.ClusterSettings']]:
        return pulumi.get(self, "cluster_settings")

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.ClusterConfiguration']:
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="defaultCapacityProviderStrategy")
    def default_capacity_provider_strategy(self) -> Optional[Sequence['outputs.ClusterCapacityProviderStrategyItem']]:
        return pulumi.get(self, "default_capacity_provider_strategy")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ClusterTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            arn=self.arn,
            capacity_providers=self.capacity_providers,
            cluster_settings=self.cluster_settings,
            configuration=self.configuration,
            default_capacity_provider_strategy=self.default_capacity_provider_strategy,
            tags=self.tags)


def get_cluster(cluster_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Create an Elastic Container Service (ECS) cluster.


    :param str cluster_name: A user-generated string that you use to identify your cluster. If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ecs:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        arn=__ret__.arn,
        capacity_providers=__ret__.capacity_providers,
        cluster_settings=__ret__.cluster_settings,
        configuration=__ret__.configuration,
        default_capacity_provider_strategy=__ret__.default_capacity_provider_strategy,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_cluster)
def get_cluster_output(cluster_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterResult]:
    """
    Create an Elastic Container Service (ECS) cluster.


    :param str cluster_name: A user-generated string that you use to identify your cluster. If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
    """
    ...