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

__all__ = [
    'GetScalingPolicyResult',
    'AwaitableGetScalingPolicyResult',
    'get_scaling_policy',
    'get_scaling_policy_output',
]

@pulumi.output_type
class GetScalingPolicyResult:
    def __init__(__self__, arn=None, policy_type=None, predictive_scaling_policy_configuration=None, step_scaling_policy_configuration=None, target_tracking_scaling_policy_configuration=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)
        if predictive_scaling_policy_configuration and not isinstance(predictive_scaling_policy_configuration, dict):
            raise TypeError("Expected argument 'predictive_scaling_policy_configuration' to be a dict")
        pulumi.set(__self__, "predictive_scaling_policy_configuration", predictive_scaling_policy_configuration)
        if step_scaling_policy_configuration and not isinstance(step_scaling_policy_configuration, dict):
            raise TypeError("Expected argument 'step_scaling_policy_configuration' to be a dict")
        pulumi.set(__self__, "step_scaling_policy_configuration", step_scaling_policy_configuration)
        if target_tracking_scaling_policy_configuration and not isinstance(target_tracking_scaling_policy_configuration, dict):
            raise TypeError("Expected argument 'target_tracking_scaling_policy_configuration' to be a dict")
        pulumi.set(__self__, "target_tracking_scaling_policy_configuration", target_tracking_scaling_policy_configuration)

    @property
    @pulumi.getter
    def arn(self) -> Optional[builtins.str]:
        """
        Returns the ARN of a scaling policy.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[builtins.str]:
        """
        The scaling policy type.
         The following policy types are supported: 
          ``TargetTrackingScaling``—Not supported for Amazon EMR
          ``StepScaling``—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="predictiveScalingPolicyConfiguration")
    def predictive_scaling_policy_configuration(self) -> Optional['outputs.ScalingPolicyPredictiveScalingPolicyConfiguration']:
        """
        The predictive scaling policy configuration.
        """
        return pulumi.get(self, "predictive_scaling_policy_configuration")

    @property
    @pulumi.getter(name="stepScalingPolicyConfiguration")
    def step_scaling_policy_configuration(self) -> Optional['outputs.ScalingPolicyStepScalingPolicyConfiguration']:
        """
        A step scaling policy.
        """
        return pulumi.get(self, "step_scaling_policy_configuration")

    @property
    @pulumi.getter(name="targetTrackingScalingPolicyConfiguration")
    def target_tracking_scaling_policy_configuration(self) -> Optional['outputs.ScalingPolicyTargetTrackingScalingPolicyConfiguration']:
        """
        A target tracking scaling policy.
        """
        return pulumi.get(self, "target_tracking_scaling_policy_configuration")


class AwaitableGetScalingPolicyResult(GetScalingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingPolicyResult(
            arn=self.arn,
            policy_type=self.policy_type,
            predictive_scaling_policy_configuration=self.predictive_scaling_policy_configuration,
            step_scaling_policy_configuration=self.step_scaling_policy_configuration,
            target_tracking_scaling_policy_configuration=self.target_tracking_scaling_policy_configuration)


def get_scaling_policy(arn: Optional[builtins.str] = None,
                       scalable_dimension: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingPolicyResult:
    """
    The ``AWS::ApplicationAutoScaling::ScalingPolicy`` resource defines a scaling policy that Application Auto Scaling uses to adjust the capacity of a scalable target.
     For more information, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide*.


    :param builtins.str arn: Returns the ARN of a scaling policy.
    :param builtins.str scalable_dimension: The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
             +   ``ecs:service:DesiredCount`` - The task count of an ECS service.
             +   ``elasticmapreduce:instancegroup:InstanceCount`` - The instance count of an EMR Instance Group.
             +   ``ec2:spot-fleet-request:TargetCapacity`` - The target capacity of a Spot Fleet.
             +   ``appstream:fleet:DesiredCapacity`` - The capacity of an AppStream 2.0 fleet.
             +   ``dynamodb:table:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB table.
             +   ``dynamodb:table:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB table.
             +   ``dynamodb:index:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB global secondary index.
             +   ``dynamodb:index:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB global secondary index.
             +   ``rds:cluster:ReadReplicaCount`` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
             +   ``sagemaker:variant:DesiredInstanceCount`` - The number of EC2 instances for a SageMaker model endpoint variant.
             +   ``custom-resource:ResourceType:Property`` - The scalable dimension for a custom resource provided by your own application or service.
             +   ``comprehend:document-classifier-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend document classification endpoint.
             +   ``comprehend:entity-recognizer-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
             +   ``lambda:function:ProvisionedConcurrency`` - The provisioned concurrency for a Lambda function.
             +   ``cassandra:table:ReadCapacityUnits`` - The provisioned read capacity for an Amazon Keyspaces table.
             +   ``cassandra:table:WriteCapacityUnits`` - The provisioned write capacity for an Amazon Keyspaces table.
             +   ``kafka:broker-storage:VolumeSize`` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
             +   ``elasticache:replication-group:NodeGroups`` - The number of node groups for an Amazon ElastiCache replication group.
             +   ``elasticache:replication-group:Replicas`` - The number of replicas per node group for an Amazon ElastiCache replication group.
             +   ``neptune:cluster:ReadReplicaCount`` - The count of read replicas in an Amazon Neptune DB cluster.
             +   ``sagemaker:variant:DesiredProvisionedConcurrency`` - The provisioned concurrency for a SageMaker serverless endpoint.
             +   ``sagemaker:inference-component:DesiredCopyCount`` - The number of copies across an endpoint for a SageMaker inference component.
             +   ``workspaces:workspacespool:DesiredUserSessions`` - The number of user sessions for the WorkSpaces in the pool.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['scalableDimension'] = scalable_dimension
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:applicationautoscaling:getScalingPolicy', __args__, opts=opts, typ=GetScalingPolicyResult).value

    return AwaitableGetScalingPolicyResult(
        arn=pulumi.get(__ret__, 'arn'),
        policy_type=pulumi.get(__ret__, 'policy_type'),
        predictive_scaling_policy_configuration=pulumi.get(__ret__, 'predictive_scaling_policy_configuration'),
        step_scaling_policy_configuration=pulumi.get(__ret__, 'step_scaling_policy_configuration'),
        target_tracking_scaling_policy_configuration=pulumi.get(__ret__, 'target_tracking_scaling_policy_configuration'))
def get_scaling_policy_output(arn: Optional[pulumi.Input[builtins.str]] = None,
                              scalable_dimension: Optional[pulumi.Input[builtins.str]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetScalingPolicyResult]:
    """
    The ``AWS::ApplicationAutoScaling::ScalingPolicy`` resource defines a scaling policy that Application Auto Scaling uses to adjust the capacity of a scalable target.
     For more information, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide*.


    :param builtins.str arn: Returns the ARN of a scaling policy.
    :param builtins.str scalable_dimension: The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
             +   ``ecs:service:DesiredCount`` - The task count of an ECS service.
             +   ``elasticmapreduce:instancegroup:InstanceCount`` - The instance count of an EMR Instance Group.
             +   ``ec2:spot-fleet-request:TargetCapacity`` - The target capacity of a Spot Fleet.
             +   ``appstream:fleet:DesiredCapacity`` - The capacity of an AppStream 2.0 fleet.
             +   ``dynamodb:table:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB table.
             +   ``dynamodb:table:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB table.
             +   ``dynamodb:index:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB global secondary index.
             +   ``dynamodb:index:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB global secondary index.
             +   ``rds:cluster:ReadReplicaCount`` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
             +   ``sagemaker:variant:DesiredInstanceCount`` - The number of EC2 instances for a SageMaker model endpoint variant.
             +   ``custom-resource:ResourceType:Property`` - The scalable dimension for a custom resource provided by your own application or service.
             +   ``comprehend:document-classifier-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend document classification endpoint.
             +   ``comprehend:entity-recognizer-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
             +   ``lambda:function:ProvisionedConcurrency`` - The provisioned concurrency for a Lambda function.
             +   ``cassandra:table:ReadCapacityUnits`` - The provisioned read capacity for an Amazon Keyspaces table.
             +   ``cassandra:table:WriteCapacityUnits`` - The provisioned write capacity for an Amazon Keyspaces table.
             +   ``kafka:broker-storage:VolumeSize`` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
             +   ``elasticache:replication-group:NodeGroups`` - The number of node groups for an Amazon ElastiCache replication group.
             +   ``elasticache:replication-group:Replicas`` - The number of replicas per node group for an Amazon ElastiCache replication group.
             +   ``neptune:cluster:ReadReplicaCount`` - The count of read replicas in an Amazon Neptune DB cluster.
             +   ``sagemaker:variant:DesiredProvisionedConcurrency`` - The provisioned concurrency for a SageMaker serverless endpoint.
             +   ``sagemaker:inference-component:DesiredCopyCount`` - The number of copies across an endpoint for a SageMaker inference component.
             +   ``workspaces:workspacespool:DesiredUserSessions`` - The number of user sessions for the WorkSpaces in the pool.
    """
    __args__ = dict()
    __args__['arn'] = arn
    __args__['scalableDimension'] = scalable_dimension
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:applicationautoscaling:getScalingPolicy', __args__, opts=opts, typ=GetScalingPolicyResult)
    return __ret__.apply(lambda __response__: GetScalingPolicyResult(
        arn=pulumi.get(__response__, 'arn'),
        policy_type=pulumi.get(__response__, 'policy_type'),
        predictive_scaling_policy_configuration=pulumi.get(__response__, 'predictive_scaling_policy_configuration'),
        step_scaling_policy_configuration=pulumi.get(__response__, 'step_scaling_policy_configuration'),
        target_tracking_scaling_policy_configuration=pulumi.get(__response__, 'target_tracking_scaling_policy_configuration')))
