# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    def __init__(__self__, cluster_arn=None, cluster_status=None, creation_time=None, failure_message=None, instance_groups=None, node_recovery=None, tags=None):
        if cluster_arn and not isinstance(cluster_arn, str):
            raise TypeError("Expected argument 'cluster_arn' to be a str")
        pulumi.set(__self__, "cluster_arn", cluster_arn)
        if cluster_status and not isinstance(cluster_status, str):
            raise TypeError("Expected argument 'cluster_status' to be a str")
        pulumi.set(__self__, "cluster_status", cluster_status)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if failure_message and not isinstance(failure_message, str):
            raise TypeError("Expected argument 'failure_message' to be a str")
        pulumi.set(__self__, "failure_message", failure_message)
        if instance_groups and not isinstance(instance_groups, list):
            raise TypeError("Expected argument 'instance_groups' to be a list")
        pulumi.set(__self__, "instance_groups", instance_groups)
        if node_recovery and not isinstance(node_recovery, str):
            raise TypeError("Expected argument 'node_recovery' to be a str")
        pulumi.set(__self__, "node_recovery", node_recovery)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterArn")
    def cluster_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the HyperPod Cluster.
        """
        return pulumi.get(self, "cluster_arn")

    @property
    @pulumi.getter(name="clusterStatus")
    def cluster_status(self) -> Optional['ClusterStatus']:
        """
        The status of the HyperPod Cluster.
        """
        return pulumi.get(self, "cluster_status")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        The time at which the HyperPod cluster was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="failureMessage")
    def failure_message(self) -> Optional[str]:
        """
        The failure message of the HyperPod Cluster.
        """
        return pulumi.get(self, "failure_message")

    @property
    @pulumi.getter(name="instanceGroups")
    def instance_groups(self) -> Optional[Sequence['outputs.ClusterInstanceGroup']]:
        return pulumi.get(self, "instance_groups")

    @property
    @pulumi.getter(name="nodeRecovery")
    def node_recovery(self) -> Optional['ClusterNodeRecovery']:
        """
        If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected.
        """
        return pulumi.get(self, "node_recovery")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Custom tags for managing the SageMaker HyperPod cluster as an AWS resource. You can add tags to your cluster in the same way you add them in other AWS services that support tagging.
        """
        return pulumi.get(self, "tags")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            cluster_arn=self.cluster_arn,
            cluster_status=self.cluster_status,
            creation_time=self.creation_time,
            failure_message=self.failure_message,
            instance_groups=self.instance_groups,
            node_recovery=self.node_recovery,
            tags=self.tags)


def get_cluster(cluster_arn: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Resource Type definition for AWS::SageMaker::Cluster


    :param str cluster_arn: The Amazon Resource Name (ARN) of the HyperPod Cluster.
    """
    __args__ = dict()
    __args__['clusterArn'] = cluster_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        cluster_arn=pulumi.get(__ret__, 'cluster_arn'),
        cluster_status=pulumi.get(__ret__, 'cluster_status'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        failure_message=pulumi.get(__ret__, 'failure_message'),
        instance_groups=pulumi.get(__ret__, 'instance_groups'),
        node_recovery=pulumi.get(__ret__, 'node_recovery'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_cluster)
def get_cluster_output(cluster_arn: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterResult]:
    """
    Resource Type definition for AWS::SageMaker::Cluster


    :param str cluster_arn: The Amazon Resource Name (ARN) of the HyperPod Cluster.
    """
    ...