# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'AuthenticationModePropertiesArgs',
    'GlobalReplicationGroupMemberArgs',
    'GlobalReplicationGroupRegionalConfigurationArgs',
    'GlobalReplicationGroupReshardingConfigurationArgs',
    'ServerlessCacheCacheUsageLimitsArgs',
    'ServerlessCacheDataStorageArgs',
    'ServerlessCacheEcpuPerSecondArgs',
    'ServerlessCacheEndpointArgs',
]

@pulumi.input_type
class AuthenticationModePropertiesArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['UserAuthenticationModePropertiesType'],
                 passwords: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input['UserAuthenticationModePropertiesType'] type: Authentication Type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passwords: Passwords used for this user account. You can create up to two passwords for each user.
        """
        pulumi.set(__self__, "type", type)
        if passwords is not None:
            pulumi.set(__self__, "passwords", passwords)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['UserAuthenticationModePropertiesType']:
        """
        Authentication Type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['UserAuthenticationModePropertiesType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def passwords(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Passwords used for this user account. You can create up to two passwords for each user.
        """
        return pulumi.get(self, "passwords")

    @passwords.setter
    def passwords(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "passwords", value)


@pulumi.input_type
class GlobalReplicationGroupMemberArgs:
    def __init__(__self__, *,
                 replication_group_id: Optional[pulumi.Input[str]] = None,
                 replication_group_region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input['GlobalReplicationGroupMemberRole']] = None):
        """
        :param pulumi.Input[str] replication_group_id: Regionally unique identifier for the member i.e. ReplicationGroupId.
        :param pulumi.Input[str] replication_group_region: The AWS region of the Global Datastore member.
        :param pulumi.Input['GlobalReplicationGroupMemberRole'] role: Indicates the role of the member, primary or secondary.
        """
        if replication_group_id is not None:
            pulumi.set(__self__, "replication_group_id", replication_group_id)
        if replication_group_region is not None:
            pulumi.set(__self__, "replication_group_region", replication_group_region)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="replicationGroupId")
    def replication_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Regionally unique identifier for the member i.e. ReplicationGroupId.
        """
        return pulumi.get(self, "replication_group_id")

    @replication_group_id.setter
    def replication_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_group_id", value)

    @property
    @pulumi.getter(name="replicationGroupRegion")
    def replication_group_region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region of the Global Datastore member.
        """
        return pulumi.get(self, "replication_group_region")

    @replication_group_region.setter
    def replication_group_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_group_region", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input['GlobalReplicationGroupMemberRole']]:
        """
        Indicates the role of the member, primary or secondary.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input['GlobalReplicationGroupMemberRole']]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class GlobalReplicationGroupRegionalConfigurationArgs:
    def __init__(__self__, *,
                 replication_group_id: Optional[pulumi.Input[str]] = None,
                 replication_group_region: Optional[pulumi.Input[str]] = None,
                 resharding_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalReplicationGroupReshardingConfigurationArgs']]]] = None):
        """
        :param pulumi.Input[str] replication_group_id: The replication group id of the Global Datastore member.
        :param pulumi.Input[str] replication_group_region: The AWS region of the Global Datastore member.
        :param pulumi.Input[Sequence[pulumi.Input['GlobalReplicationGroupReshardingConfigurationArgs']]] resharding_configurations: A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. 
        """
        if replication_group_id is not None:
            pulumi.set(__self__, "replication_group_id", replication_group_id)
        if replication_group_region is not None:
            pulumi.set(__self__, "replication_group_region", replication_group_region)
        if resharding_configurations is not None:
            pulumi.set(__self__, "resharding_configurations", resharding_configurations)

    @property
    @pulumi.getter(name="replicationGroupId")
    def replication_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The replication group id of the Global Datastore member.
        """
        return pulumi.get(self, "replication_group_id")

    @replication_group_id.setter
    def replication_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_group_id", value)

    @property
    @pulumi.getter(name="replicationGroupRegion")
    def replication_group_region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region of the Global Datastore member.
        """
        return pulumi.get(self, "replication_group_region")

    @replication_group_region.setter
    def replication_group_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replication_group_region", value)

    @property
    @pulumi.getter(name="reshardingConfigurations")
    def resharding_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GlobalReplicationGroupReshardingConfigurationArgs']]]]:
        """
        A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. 
        """
        return pulumi.get(self, "resharding_configurations")

    @resharding_configurations.setter
    def resharding_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GlobalReplicationGroupReshardingConfigurationArgs']]]]):
        pulumi.set(self, "resharding_configurations", value)


@pulumi.input_type
class GlobalReplicationGroupReshardingConfigurationArgs:
    def __init__(__self__, *,
                 node_group_id: Optional[pulumi.Input[str]] = None,
                 preferred_availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] node_group_id: Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] preferred_availability_zones: A list of preferred availability zones for the nodes of new node groups.
        """
        if node_group_id is not None:
            pulumi.set(__self__, "node_group_id", node_group_id)
        if preferred_availability_zones is not None:
            pulumi.set(__self__, "preferred_availability_zones", preferred_availability_zones)

    @property
    @pulumi.getter(name="nodeGroupId")
    def node_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.
        """
        return pulumi.get(self, "node_group_id")

    @node_group_id.setter
    def node_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_group_id", value)

    @property
    @pulumi.getter(name="preferredAvailabilityZones")
    def preferred_availability_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of preferred availability zones for the nodes of new node groups.
        """
        return pulumi.get(self, "preferred_availability_zones")

    @preferred_availability_zones.setter
    def preferred_availability_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "preferred_availability_zones", value)


@pulumi.input_type
class ServerlessCacheCacheUsageLimitsArgs:
    def __init__(__self__, *,
                 data_storage: Optional[pulumi.Input['ServerlessCacheDataStorageArgs']] = None,
                 ecpu_per_second: Optional[pulumi.Input['ServerlessCacheEcpuPerSecondArgs']] = None):
        """
        The cache capacity limit of the Serverless Cache.
        """
        if data_storage is not None:
            pulumi.set(__self__, "data_storage", data_storage)
        if ecpu_per_second is not None:
            pulumi.set(__self__, "ecpu_per_second", ecpu_per_second)

    @property
    @pulumi.getter(name="dataStorage")
    def data_storage(self) -> Optional[pulumi.Input['ServerlessCacheDataStorageArgs']]:
        return pulumi.get(self, "data_storage")

    @data_storage.setter
    def data_storage(self, value: Optional[pulumi.Input['ServerlessCacheDataStorageArgs']]):
        pulumi.set(self, "data_storage", value)

    @property
    @pulumi.getter(name="ecpuPerSecond")
    def ecpu_per_second(self) -> Optional[pulumi.Input['ServerlessCacheEcpuPerSecondArgs']]:
        return pulumi.get(self, "ecpu_per_second")

    @ecpu_per_second.setter
    def ecpu_per_second(self, value: Optional[pulumi.Input['ServerlessCacheEcpuPerSecondArgs']]):
        pulumi.set(self, "ecpu_per_second", value)


@pulumi.input_type
class ServerlessCacheDataStorageArgs:
    def __init__(__self__, *,
                 maximum: pulumi.Input[int],
                 unit: pulumi.Input['ServerlessCacheDataStorageUnit']):
        """
        The cached data capacity of the Serverless Cache.
        :param pulumi.Input[int] maximum: The maximum cached data capacity of the Serverless Cache.
        :param pulumi.Input['ServerlessCacheDataStorageUnit'] unit: The unix of cached data capacity of the Serverless Cache.
        """
        pulumi.set(__self__, "maximum", maximum)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter
    def maximum(self) -> pulumi.Input[int]:
        """
        The maximum cached data capacity of the Serverless Cache.
        """
        return pulumi.get(self, "maximum")

    @maximum.setter
    def maximum(self, value: pulumi.Input[int]):
        pulumi.set(self, "maximum", value)

    @property
    @pulumi.getter
    def unit(self) -> pulumi.Input['ServerlessCacheDataStorageUnit']:
        """
        The unix of cached data capacity of the Serverless Cache.
        """
        return pulumi.get(self, "unit")

    @unit.setter
    def unit(self, value: pulumi.Input['ServerlessCacheDataStorageUnit']):
        pulumi.set(self, "unit", value)


@pulumi.input_type
class ServerlessCacheEcpuPerSecondArgs:
    def __init__(__self__, *,
                 maximum: pulumi.Input[int]):
        """
        The ECPU per second of the Serverless Cache.
        :param pulumi.Input[int] maximum: The maximum ECPU per second of the Serverless Cache.
        """
        pulumi.set(__self__, "maximum", maximum)

    @property
    @pulumi.getter
    def maximum(self) -> pulumi.Input[int]:
        """
        The maximum ECPU per second of the Serverless Cache.
        """
        return pulumi.get(self, "maximum")

    @maximum.setter
    def maximum(self, value: pulumi.Input[int]):
        pulumi.set(self, "maximum", value)


@pulumi.input_type
class ServerlessCacheEndpointArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None):
        """
        The address and the port.
        :param pulumi.Input[str] address: Endpoint address.
        :param pulumi.Input[str] port: Endpoint port.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)


