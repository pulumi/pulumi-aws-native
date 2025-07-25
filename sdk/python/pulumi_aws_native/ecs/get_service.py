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
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    def __init__(__self__, availability_zone_rebalancing=None, capacity_provider_strategy=None, deployment_configuration=None, deployment_controller=None, desired_count=None, enable_ecs_managed_tags=None, enable_execute_command=None, health_check_grace_period_seconds=None, load_balancers=None, name=None, network_configuration=None, placement_constraints=None, placement_strategies=None, platform_version=None, propagate_tags=None, service_arn=None, service_registries=None, tags=None, task_definition=None, vpc_lattice_configurations=None):
        if availability_zone_rebalancing and not isinstance(availability_zone_rebalancing, str):
            raise TypeError("Expected argument 'availability_zone_rebalancing' to be a str")
        pulumi.set(__self__, "availability_zone_rebalancing", availability_zone_rebalancing)
        if capacity_provider_strategy and not isinstance(capacity_provider_strategy, list):
            raise TypeError("Expected argument 'capacity_provider_strategy' to be a list")
        pulumi.set(__self__, "capacity_provider_strategy", capacity_provider_strategy)
        if deployment_configuration and not isinstance(deployment_configuration, dict):
            raise TypeError("Expected argument 'deployment_configuration' to be a dict")
        pulumi.set(__self__, "deployment_configuration", deployment_configuration)
        if deployment_controller and not isinstance(deployment_controller, dict):
            raise TypeError("Expected argument 'deployment_controller' to be a dict")
        pulumi.set(__self__, "deployment_controller", deployment_controller)
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
        if vpc_lattice_configurations and not isinstance(vpc_lattice_configurations, list):
            raise TypeError("Expected argument 'vpc_lattice_configurations' to be a list")
        pulumi.set(__self__, "vpc_lattice_configurations", vpc_lattice_configurations)

    @property
    @pulumi.getter(name="availabilityZoneRebalancing")
    def availability_zone_rebalancing(self) -> Optional['ServiceAvailabilityZoneRebalancing']:
        """
        Indicates whether to use Availability Zone rebalancing for the service.
         For more information, see [Balancing an Amazon ECS service across Availability Zones](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-rebalancing.html) in the *Amazon Elastic Container Service Developer Guide*.
        """
        return pulumi.get(self, "availability_zone_rebalancing")

    @property
    @pulumi.getter(name="capacityProviderStrategy")
    def capacity_provider_strategy(self) -> Optional[Sequence['outputs.ServiceCapacityProviderStrategyItem']]:
        """
        The capacity provider strategy to use for the service.
         If a ``capacityProviderStrategy`` is specified, the ``launchType`` parameter must be omitted. If no ``capacityProviderStrategy`` or ``launchType`` is specified, the ``defaultCapacityProviderStrategy`` for the cluster is used.
         A capacity provider strategy can contain a maximum of 20 capacity providers.
          To remove this property from your service resource, specify an empty ``CapacityProviderStrategyItem`` array.
        """
        return pulumi.get(self, "capacity_provider_strategy")

    @property
    @pulumi.getter(name="deploymentConfiguration")
    def deployment_configuration(self) -> Optional['outputs.ServiceDeploymentConfiguration']:
        """
        Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.
        """
        return pulumi.get(self, "deployment_configuration")

    @property
    @pulumi.getter(name="deploymentController")
    def deployment_controller(self) -> Optional['outputs.ServiceDeploymentController']:
        """
        The deployment controller to use for the service. If no deployment controller is specified, the default value of ``ECS`` is used.
        """
        return pulumi.get(self, "deployment_controller")

    @property
    @pulumi.getter(name="desiredCount")
    def desired_count(self) -> Optional[builtins.int]:
        """
        The number of instantiations of the specified task definition to place and keep running in your service.
         For new services, if a desired count is not specified, a default value of ``1`` is used. When using the ``DAEMON`` scheduling strategy, the desired count is not required.
         For existing services, if a desired count is not specified, it is omitted from the operation.
        """
        return pulumi.get(self, "desired_count")

    @property
    @pulumi.getter(name="enableEcsManagedTags")
    def enable_ecs_managed_tags(self) -> Optional[builtins.bool]:
        """
        Specifies whether to turn on Amazon ECS managed tags for the tasks within the service. For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide*.
         When you use Amazon ECS managed tags, you need to set the ``propagateTags`` request parameter.
        """
        return pulumi.get(self, "enable_ecs_managed_tags")

    @property
    @pulumi.getter(name="enableExecuteCommand")
    def enable_execute_command(self) -> Optional[builtins.bool]:
        """
        Determines whether the execute command functionality is turned on for the service. If ``true``, the execute command functionality is turned on for all containers in tasks as part of the service.
        """
        return pulumi.get(self, "enable_execute_command")

    @property
    @pulumi.getter(name="healthCheckGracePeriodSeconds")
    def health_check_grace_period_seconds(self) -> Optional[builtins.int]:
        """
        The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task has first started. If you don't specify a health check grace period value, the default value of ``0`` is used. If you don't use any of the health checks, then ``healthCheckGracePeriodSeconds`` is unused.
         If your service's tasks take a while to start and respond to health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.
        """
        return pulumi.get(self, "health_check_grace_period_seconds")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> Optional[Sequence['outputs.ServiceLoadBalancer']]:
        """
        A list of load balancer objects to associate with the service. If you specify the ``Role`` property, ``LoadBalancers`` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide*.
          To remove this property from your service resource, specify an empty ``LoadBalancer`` array.
        """
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the Amazon ECS service, such as `sample-webapp` .
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfiguration")
    def network_configuration(self) -> Optional['outputs.ServiceNetworkConfiguration']:
        """
        The network configuration for the service. This parameter is required for task definitions that use the ``awsvpc`` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide*.
        """
        return pulumi.get(self, "network_configuration")

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> Optional[Sequence['outputs.ServicePlacementConstraint']]:
        """
        An array of placement constraint objects to use for tasks in your service. You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.
          To remove this property from your service resource, specify an empty ``PlacementConstraint`` array.
        """
        return pulumi.get(self, "placement_constraints")

    @property
    @pulumi.getter(name="placementStrategies")
    def placement_strategies(self) -> Optional[Sequence['outputs.ServicePlacementStrategy']]:
        """
        The placement strategy objects to use for tasks in your service. You can specify a maximum of 5 strategy rules for each service.
          To remove this property from your service resource, specify an empty ``PlacementStrategy`` array.
        """
        return pulumi.get(self, "placement_strategies")

    @property
    @pulumi.getter(name="platformVersion")
    def platform_version(self) -> Optional[builtins.str]:
        """
        The platform version that your tasks in the service are running on. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the ``LATEST`` platform version is used. For more information, see [platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide*.
        """
        return pulumi.get(self, "platform_version")

    @property
    @pulumi.getter(name="propagateTags")
    def propagate_tags(self) -> Optional['ServicePropagateTags']:
        """
        Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags aren't propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.
         You must set this to a value other than ``NONE`` when you use Cost Explorer. For more information, see [Amazon ECS usage reports](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/usage-reports.html) in the *Amazon Elastic Container Service Developer Guide*.
         The default is ``NONE``.
        """
        return pulumi.get(self, "propagate_tags")

    @property
    @pulumi.getter(name="serviceArn")
    def service_arn(self) -> Optional[builtins.str]:
        """
        The ARN that identifies the service. For more information about the ARN format, see [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids) in the *Amazon ECS Developer Guide* .
        """
        return pulumi.get(self, "service_arn")

    @property
    @pulumi.getter(name="serviceRegistries")
    def service_registries(self) -> Optional[Sequence['outputs.ServiceRegistry']]:
        """
        The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).
          Each service may be associated with one service registry. Multiple service registries for each service isn't supported.
           To remove this property from your service resource, specify an empty ``ServiceRegistry`` array.
        """
        return pulumi.get(self, "service_registries")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        The metadata that you apply to the service to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well.
         The following basic restrictions apply to tags:
          +  Maximum number of tags per resource - 50
          +  For each resource, each tag key must be unique, and each tag key can have only one value.
          +  Maximum key length - 128 Unicode characters in UTF-8
          +  Maximum value length - 256 Unicode characters in UTF-8
          +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
          +  Tag keys and values are case-sensitive.
          +  Do not use ``aws:``, ``AWS:``, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskDefinition")
    def task_definition(self) -> Optional[builtins.str]:
        """
        The ``family`` and ``revision`` (``family:revision``) or full ARN of the task definition to run in your service. If a ``revision`` isn't specified, the latest ``ACTIVE`` revision is used.
         A task definition must be specified if the service uses either the ``ECS`` or ``CODE_DEPLOY`` deployment controllers.
         For more information about deployment types, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html).
        """
        return pulumi.get(self, "task_definition")

    @property
    @pulumi.getter(name="vpcLatticeConfigurations")
    def vpc_lattice_configurations(self) -> Optional[Sequence['outputs.ServiceVpcLatticeConfiguration']]:
        """
        The VPC Lattice configuration for the service being created.
        """
        return pulumi.get(self, "vpc_lattice_configurations")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            availability_zone_rebalancing=self.availability_zone_rebalancing,
            capacity_provider_strategy=self.capacity_provider_strategy,
            deployment_configuration=self.deployment_configuration,
            deployment_controller=self.deployment_controller,
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
            task_definition=self.task_definition,
            vpc_lattice_configurations=self.vpc_lattice_configurations)


def get_service(cluster: Optional[builtins.str] = None,
                service_arn: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    The ``AWS::ECS::Service`` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
      The stack update fails if you change any properties that require replacement and at least one ECS Service Connect ``ServiceConnectConfiguration`` property is configured. This is because AWS CloudFormation creates the replacement service first, but each ``ServiceConnectService`` must have a name that is unique in the namespace.
       Starting April 15, 2023, AWS; will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, ECS, or EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.
       On June 12, 2025, Amazon ECS launched support for updating capacity provider configuration for ECS services. With this launch, ECS also aligned the CFN update behavior for ``CapacityProviderStrategy`` parameter with the standard practice. For more information, see [adds support for updating capacity provider configuration for ECS services](https://docs.aws.amazon.com/about-aws/whats-new/2025/05/amazon-ecs-capacity-provider-configuration-ecs/). Previously ECS ignored the ``CapacityProviderStrategy`` property if it was set to an empty list for example, ``[]`` in CFN, because updating capacity provider configuration was not supported. Now, with support for capacity provider updates, customers can remove capacity providers from a service by passing an empty list. When you specify an empty list (``[]``) for the ``CapacityProviderStrategy`` property in your CFN template, ECS will remove any capacity providers associated with the service, as follows:
      +  For services created with a capacity provider strategy after the launch:
      +  If there's a cluster default strategy set, the service will revert to using that default strategy.
      +  If no cluster default strategy exists, you will receive the following error:
     No launch type to fall back to for empty capacity provider strategy. Your service was not created with a launch type.

      +  For services created with a capacity provider strategy prior to the launch:
      +  If ``CapacityProviderStrategy`` had ``FARGATE_SPOT`` or ``FARGATE`` capacity providers, the launch type will be updated to ``FARGATE`` and the capacity provider will be removed.
      +  If the strategy included Auto Scaling group capacity providers, the service will revert to EC2 launch type, and the Auto Scaling group capacity providers will not be used.

     Recommended Actions
     If you are currently using ``CapacityProviderStrategy: []`` in your CFN templates, you should take one of the following actions:
      +  If you do not intend to update the Capacity Provider Strategy:
      +  Remove the ``CapacityProviderStrategy`` property entirely from your CFN template
      +  Alternatively, use ``!Ref ::NoValue`` for the ``CapacityProviderStrategy`` property in your template

      +  If you intend to maintain or update the Capacity Provider Strategy, specify the actual Capacity Provider Strategy for the service in your CFN template.

     If your CFN template had an empty list ([]) for ``CapacityProviderStrategy`` prior to the aforementioned launch on June 12, and you are using the same template with ``CapacityProviderStrategy: []``, you might encounter the following error:
      Invalid request provided: When switching from launch type to capacity provider strategy on an existing service, or making a change to a capacity provider strategy on a service that is already using one, you must force a new deployment. (Service: Ecs, Status Code: 400, Request ID: xxx) (SDK Attempt Count: 1)" (RequestToken: xxx HandlerErrorCode: InvalidRequest)
     Note that CFN automatically initiates a new deployment when it detects a parameter change, but customers cannot choose to force a deployment through CFN. This is an invalid input scenario that requires one of the remediation actions listed above.
     If you are experiencing active production issues related to this change, contact AWS Support or your Technical Account Manager.


    :param builtins.str cluster: The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on. If you do not specify a cluster, the default cluster is assumed.
    :param builtins.str service_arn: The ARN that identifies the service. For more information about the ARN format, see [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids) in the *Amazon ECS Developer Guide* .
    """
    __args__ = dict()
    __args__['cluster'] = cluster
    __args__['serviceArn'] = service_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ecs:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        availability_zone_rebalancing=pulumi.get(__ret__, 'availability_zone_rebalancing'),
        capacity_provider_strategy=pulumi.get(__ret__, 'capacity_provider_strategy'),
        deployment_configuration=pulumi.get(__ret__, 'deployment_configuration'),
        deployment_controller=pulumi.get(__ret__, 'deployment_controller'),
        desired_count=pulumi.get(__ret__, 'desired_count'),
        enable_ecs_managed_tags=pulumi.get(__ret__, 'enable_ecs_managed_tags'),
        enable_execute_command=pulumi.get(__ret__, 'enable_execute_command'),
        health_check_grace_period_seconds=pulumi.get(__ret__, 'health_check_grace_period_seconds'),
        load_balancers=pulumi.get(__ret__, 'load_balancers'),
        name=pulumi.get(__ret__, 'name'),
        network_configuration=pulumi.get(__ret__, 'network_configuration'),
        placement_constraints=pulumi.get(__ret__, 'placement_constraints'),
        placement_strategies=pulumi.get(__ret__, 'placement_strategies'),
        platform_version=pulumi.get(__ret__, 'platform_version'),
        propagate_tags=pulumi.get(__ret__, 'propagate_tags'),
        service_arn=pulumi.get(__ret__, 'service_arn'),
        service_registries=pulumi.get(__ret__, 'service_registries'),
        tags=pulumi.get(__ret__, 'tags'),
        task_definition=pulumi.get(__ret__, 'task_definition'),
        vpc_lattice_configurations=pulumi.get(__ret__, 'vpc_lattice_configurations'))
def get_service_output(cluster: Optional[pulumi.Input[builtins.str]] = None,
                       service_arn: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceResult]:
    """
    The ``AWS::ECS::Service`` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
      The stack update fails if you change any properties that require replacement and at least one ECS Service Connect ``ServiceConnectConfiguration`` property is configured. This is because AWS CloudFormation creates the replacement service first, but each ``ServiceConnectService`` must have a name that is unique in the namespace.
       Starting April 15, 2023, AWS; will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, ECS, or EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.
       On June 12, 2025, Amazon ECS launched support for updating capacity provider configuration for ECS services. With this launch, ECS also aligned the CFN update behavior for ``CapacityProviderStrategy`` parameter with the standard practice. For more information, see [adds support for updating capacity provider configuration for ECS services](https://docs.aws.amazon.com/about-aws/whats-new/2025/05/amazon-ecs-capacity-provider-configuration-ecs/). Previously ECS ignored the ``CapacityProviderStrategy`` property if it was set to an empty list for example, ``[]`` in CFN, because updating capacity provider configuration was not supported. Now, with support for capacity provider updates, customers can remove capacity providers from a service by passing an empty list. When you specify an empty list (``[]``) for the ``CapacityProviderStrategy`` property in your CFN template, ECS will remove any capacity providers associated with the service, as follows:
      +  For services created with a capacity provider strategy after the launch:
      +  If there's a cluster default strategy set, the service will revert to using that default strategy.
      +  If no cluster default strategy exists, you will receive the following error:
     No launch type to fall back to for empty capacity provider strategy. Your service was not created with a launch type.

      +  For services created with a capacity provider strategy prior to the launch:
      +  If ``CapacityProviderStrategy`` had ``FARGATE_SPOT`` or ``FARGATE`` capacity providers, the launch type will be updated to ``FARGATE`` and the capacity provider will be removed.
      +  If the strategy included Auto Scaling group capacity providers, the service will revert to EC2 launch type, and the Auto Scaling group capacity providers will not be used.

     Recommended Actions
     If you are currently using ``CapacityProviderStrategy: []`` in your CFN templates, you should take one of the following actions:
      +  If you do not intend to update the Capacity Provider Strategy:
      +  Remove the ``CapacityProviderStrategy`` property entirely from your CFN template
      +  Alternatively, use ``!Ref ::NoValue`` for the ``CapacityProviderStrategy`` property in your template

      +  If you intend to maintain or update the Capacity Provider Strategy, specify the actual Capacity Provider Strategy for the service in your CFN template.

     If your CFN template had an empty list ([]) for ``CapacityProviderStrategy`` prior to the aforementioned launch on June 12, and you are using the same template with ``CapacityProviderStrategy: []``, you might encounter the following error:
      Invalid request provided: When switching from launch type to capacity provider strategy on an existing service, or making a change to a capacity provider strategy on a service that is already using one, you must force a new deployment. (Service: Ecs, Status Code: 400, Request ID: xxx) (SDK Attempt Count: 1)" (RequestToken: xxx HandlerErrorCode: InvalidRequest)
     Note that CFN automatically initiates a new deployment when it detects a parameter change, but customers cannot choose to force a deployment through CFN. This is an invalid input scenario that requires one of the remediation actions listed above.
     If you are experiencing active production issues related to this change, contact AWS Support or your Technical Account Manager.


    :param builtins.str cluster: The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on. If you do not specify a cluster, the default cluster is assumed.
    :param builtins.str service_arn: The ARN that identifies the service. For more information about the ARN format, see [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids) in the *Amazon ECS Developer Guide* .
    """
    __args__ = dict()
    __args__['cluster'] = cluster
    __args__['serviceArn'] = service_arn
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:ecs:getService', __args__, opts=opts, typ=GetServiceResult)
    return __ret__.apply(lambda __response__: GetServiceResult(
        availability_zone_rebalancing=pulumi.get(__response__, 'availability_zone_rebalancing'),
        capacity_provider_strategy=pulumi.get(__response__, 'capacity_provider_strategy'),
        deployment_configuration=pulumi.get(__response__, 'deployment_configuration'),
        deployment_controller=pulumi.get(__response__, 'deployment_controller'),
        desired_count=pulumi.get(__response__, 'desired_count'),
        enable_ecs_managed_tags=pulumi.get(__response__, 'enable_ecs_managed_tags'),
        enable_execute_command=pulumi.get(__response__, 'enable_execute_command'),
        health_check_grace_period_seconds=pulumi.get(__response__, 'health_check_grace_period_seconds'),
        load_balancers=pulumi.get(__response__, 'load_balancers'),
        name=pulumi.get(__response__, 'name'),
        network_configuration=pulumi.get(__response__, 'network_configuration'),
        placement_constraints=pulumi.get(__response__, 'placement_constraints'),
        placement_strategies=pulumi.get(__response__, 'placement_strategies'),
        platform_version=pulumi.get(__response__, 'platform_version'),
        propagate_tags=pulumi.get(__response__, 'propagate_tags'),
        service_arn=pulumi.get(__response__, 'service_arn'),
        service_registries=pulumi.get(__response__, 'service_registries'),
        tags=pulumi.get(__response__, 'tags'),
        task_definition=pulumi.get(__response__, 'task_definition'),
        vpc_lattice_configurations=pulumi.get(__response__, 'vpc_lattice_configurations')))
