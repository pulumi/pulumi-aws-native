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
from ._inputs import *

__all__ = ['AutoScalingGroupArgs', 'AutoScalingGroup']

@pulumi.input_type
class AutoScalingGroupArgs:
    def __init__(__self__, *,
                 max_size: pulumi.Input[str],
                 min_size: pulumi.Input[str],
                 auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 capacity_rebalance: Optional[pulumi.Input[bool]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 cooldown: Optional[pulumi.Input[str]] = None,
                 default_instance_warmup: Optional[pulumi.Input[int]] = None,
                 desired_capacity: Optional[pulumi.Input[str]] = None,
                 desired_capacity_type: Optional[pulumi.Input[str]] = None,
                 health_check_grace_period: Optional[pulumi.Input[int]] = None,
                 health_check_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 launch_configuration_name: Optional[pulumi.Input[str]] = None,
                 launch_template: Optional[pulumi.Input['AutoScalingGroupLaunchTemplateSpecificationArgs']] = None,
                 lifecycle_hook_specification_list: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupLifecycleHookSpecificationArgs']]]] = None,
                 load_balancer_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_instance_lifetime: Optional[pulumi.Input[int]] = None,
                 metrics_collection: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupMetricsCollectionArgs']]]] = None,
                 mixed_instances_policy: Optional[pulumi.Input['AutoScalingGroupMixedInstancesPolicyArgs']] = None,
                 new_instances_protected_from_scale_in: Optional[pulumi.Input[bool]] = None,
                 notification_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupNotificationConfigurationArgs']]]] = None,
                 placement_group: Optional[pulumi.Input[str]] = None,
                 service_linked_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupTagPropertyArgs']]]] = None,
                 target_group_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 termination_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v_pc_zone_identifier: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AutoScalingGroup resource.
        """
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)
        if auto_scaling_group_name is not None:
            pulumi.set(__self__, "auto_scaling_group_name", auto_scaling_group_name)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if capacity_rebalance is not None:
            pulumi.set(__self__, "capacity_rebalance", capacity_rebalance)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if cooldown is not None:
            pulumi.set(__self__, "cooldown", cooldown)
        if default_instance_warmup is not None:
            pulumi.set(__self__, "default_instance_warmup", default_instance_warmup)
        if desired_capacity is not None:
            pulumi.set(__self__, "desired_capacity", desired_capacity)
        if desired_capacity_type is not None:
            pulumi.set(__self__, "desired_capacity_type", desired_capacity_type)
        if health_check_grace_period is not None:
            pulumi.set(__self__, "health_check_grace_period", health_check_grace_period)
        if health_check_type is not None:
            pulumi.set(__self__, "health_check_type", health_check_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if launch_configuration_name is not None:
            pulumi.set(__self__, "launch_configuration_name", launch_configuration_name)
        if launch_template is not None:
            pulumi.set(__self__, "launch_template", launch_template)
        if lifecycle_hook_specification_list is not None:
            pulumi.set(__self__, "lifecycle_hook_specification_list", lifecycle_hook_specification_list)
        if load_balancer_names is not None:
            pulumi.set(__self__, "load_balancer_names", load_balancer_names)
        if max_instance_lifetime is not None:
            pulumi.set(__self__, "max_instance_lifetime", max_instance_lifetime)
        if metrics_collection is not None:
            pulumi.set(__self__, "metrics_collection", metrics_collection)
        if mixed_instances_policy is not None:
            pulumi.set(__self__, "mixed_instances_policy", mixed_instances_policy)
        if new_instances_protected_from_scale_in is not None:
            pulumi.set(__self__, "new_instances_protected_from_scale_in", new_instances_protected_from_scale_in)
        if notification_configurations is not None:
            pulumi.set(__self__, "notification_configurations", notification_configurations)
        if placement_group is not None:
            pulumi.set(__self__, "placement_group", placement_group)
        if service_linked_role_arn is not None:
            pulumi.set(__self__, "service_linked_role_arn", service_linked_role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_group_arns is not None:
            pulumi.set(__self__, "target_group_arns", target_group_arns)
        if termination_policies is not None:
            pulumi.set(__self__, "termination_policies", termination_policies)
        if v_pc_zone_identifier is not None:
            pulumi.set(__self__, "v_pc_zone_identifier", v_pc_zone_identifier)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Input[str]:
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: pulumi.Input[str]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Input[str]:
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: pulumi.Input[str]):
        pulumi.set(self, "min_size", value)

    @property
    @pulumi.getter(name="autoScalingGroupName")
    def auto_scaling_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_scaling_group_name")

    @auto_scaling_group_name.setter
    def auto_scaling_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_scaling_group_name", value)

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "availability_zones")

    @availability_zones.setter
    def availability_zones(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "availability_zones", value)

    @property
    @pulumi.getter(name="capacityRebalance")
    def capacity_rebalance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "capacity_rebalance")

    @capacity_rebalance.setter
    def capacity_rebalance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "capacity_rebalance", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter
    def cooldown(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cooldown")

    @cooldown.setter
    def cooldown(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cooldown", value)

    @property
    @pulumi.getter(name="defaultInstanceWarmup")
    def default_instance_warmup(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_instance_warmup")

    @default_instance_warmup.setter
    def default_instance_warmup(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_instance_warmup", value)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "desired_capacity")

    @desired_capacity.setter
    def desired_capacity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desired_capacity", value)

    @property
    @pulumi.getter(name="desiredCapacityType")
    def desired_capacity_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "desired_capacity_type")

    @desired_capacity_type.setter
    def desired_capacity_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desired_capacity_type", value)

    @property
    @pulumi.getter(name="healthCheckGracePeriod")
    def health_check_grace_period(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_grace_period")

    @health_check_grace_period.setter
    def health_check_grace_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_grace_period", value)

    @property
    @pulumi.getter(name="healthCheckType")
    def health_check_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_type")

    @health_check_type.setter
    def health_check_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="launchConfigurationName")
    def launch_configuration_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "launch_configuration_name")

    @launch_configuration_name.setter
    def launch_configuration_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch_configuration_name", value)

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> Optional[pulumi.Input['AutoScalingGroupLaunchTemplateSpecificationArgs']]:
        return pulumi.get(self, "launch_template")

    @launch_template.setter
    def launch_template(self, value: Optional[pulumi.Input['AutoScalingGroupLaunchTemplateSpecificationArgs']]):
        pulumi.set(self, "launch_template", value)

    @property
    @pulumi.getter(name="lifecycleHookSpecificationList")
    def lifecycle_hook_specification_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupLifecycleHookSpecificationArgs']]]]:
        return pulumi.get(self, "lifecycle_hook_specification_list")

    @lifecycle_hook_specification_list.setter
    def lifecycle_hook_specification_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupLifecycleHookSpecificationArgs']]]]):
        pulumi.set(self, "lifecycle_hook_specification_list", value)

    @property
    @pulumi.getter(name="loadBalancerNames")
    def load_balancer_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "load_balancer_names")

    @load_balancer_names.setter
    def load_balancer_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "load_balancer_names", value)

    @property
    @pulumi.getter(name="maxInstanceLifetime")
    def max_instance_lifetime(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_instance_lifetime")

    @max_instance_lifetime.setter
    def max_instance_lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_instance_lifetime", value)

    @property
    @pulumi.getter(name="metricsCollection")
    def metrics_collection(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupMetricsCollectionArgs']]]]:
        return pulumi.get(self, "metrics_collection")

    @metrics_collection.setter
    def metrics_collection(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupMetricsCollectionArgs']]]]):
        pulumi.set(self, "metrics_collection", value)

    @property
    @pulumi.getter(name="mixedInstancesPolicy")
    def mixed_instances_policy(self) -> Optional[pulumi.Input['AutoScalingGroupMixedInstancesPolicyArgs']]:
        return pulumi.get(self, "mixed_instances_policy")

    @mixed_instances_policy.setter
    def mixed_instances_policy(self, value: Optional[pulumi.Input['AutoScalingGroupMixedInstancesPolicyArgs']]):
        pulumi.set(self, "mixed_instances_policy", value)

    @property
    @pulumi.getter(name="newInstancesProtectedFromScaleIn")
    def new_instances_protected_from_scale_in(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "new_instances_protected_from_scale_in")

    @new_instances_protected_from_scale_in.setter
    def new_instances_protected_from_scale_in(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "new_instances_protected_from_scale_in", value)

    @property
    @pulumi.getter(name="notificationConfigurations")
    def notification_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupNotificationConfigurationArgs']]]]:
        return pulumi.get(self, "notification_configurations")

    @notification_configurations.setter
    def notification_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupNotificationConfigurationArgs']]]]):
        pulumi.set(self, "notification_configurations", value)

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "placement_group")

    @placement_group.setter
    def placement_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "placement_group", value)

    @property
    @pulumi.getter(name="serviceLinkedRoleARN")
    def service_linked_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_linked_role_arn")

    @service_linked_role_arn.setter
    def service_linked_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_linked_role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupTagPropertyArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoScalingGroupTagPropertyArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetGroupARNs")
    def target_group_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "target_group_arns")

    @target_group_arns.setter
    def target_group_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_group_arns", value)

    @property
    @pulumi.getter(name="terminationPolicies")
    def termination_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "termination_policies")

    @termination_policies.setter
    def termination_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "termination_policies", value)

    @property
    @pulumi.getter(name="vPCZoneIdentifier")
    def v_pc_zone_identifier(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "v_pc_zone_identifier")

    @v_pc_zone_identifier.setter
    def v_pc_zone_identifier(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "v_pc_zone_identifier", value)


warnings.warn("""AutoScalingGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class AutoScalingGroup(pulumi.CustomResource):
    warnings.warn("""AutoScalingGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 capacity_rebalance: Optional[pulumi.Input[bool]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 cooldown: Optional[pulumi.Input[str]] = None,
                 default_instance_warmup: Optional[pulumi.Input[int]] = None,
                 desired_capacity: Optional[pulumi.Input[str]] = None,
                 desired_capacity_type: Optional[pulumi.Input[str]] = None,
                 health_check_grace_period: Optional[pulumi.Input[int]] = None,
                 health_check_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 launch_configuration_name: Optional[pulumi.Input[str]] = None,
                 launch_template: Optional[pulumi.Input[pulumi.InputType['AutoScalingGroupLaunchTemplateSpecificationArgs']]] = None,
                 lifecycle_hook_specification_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupLifecycleHookSpecificationArgs']]]]] = None,
                 load_balancer_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_instance_lifetime: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[str]] = None,
                 metrics_collection: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupMetricsCollectionArgs']]]]] = None,
                 min_size: Optional[pulumi.Input[str]] = None,
                 mixed_instances_policy: Optional[pulumi.Input[pulumi.InputType['AutoScalingGroupMixedInstancesPolicyArgs']]] = None,
                 new_instances_protected_from_scale_in: Optional[pulumi.Input[bool]] = None,
                 notification_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupNotificationConfigurationArgs']]]]] = None,
                 placement_group: Optional[pulumi.Input[str]] = None,
                 service_linked_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupTagPropertyArgs']]]]] = None,
                 target_group_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 termination_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v_pc_zone_identifier: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AutoScaling::AutoScalingGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoScalingGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AutoScaling::AutoScalingGroup

        :param str resource_name: The name of the resource.
        :param AutoScalingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoScalingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_name: Optional[pulumi.Input[str]] = None,
                 availability_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 capacity_rebalance: Optional[pulumi.Input[bool]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 cooldown: Optional[pulumi.Input[str]] = None,
                 default_instance_warmup: Optional[pulumi.Input[int]] = None,
                 desired_capacity: Optional[pulumi.Input[str]] = None,
                 desired_capacity_type: Optional[pulumi.Input[str]] = None,
                 health_check_grace_period: Optional[pulumi.Input[int]] = None,
                 health_check_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 launch_configuration_name: Optional[pulumi.Input[str]] = None,
                 launch_template: Optional[pulumi.Input[pulumi.InputType['AutoScalingGroupLaunchTemplateSpecificationArgs']]] = None,
                 lifecycle_hook_specification_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupLifecycleHookSpecificationArgs']]]]] = None,
                 load_balancer_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_instance_lifetime: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[str]] = None,
                 metrics_collection: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupMetricsCollectionArgs']]]]] = None,
                 min_size: Optional[pulumi.Input[str]] = None,
                 mixed_instances_policy: Optional[pulumi.Input[pulumi.InputType['AutoScalingGroupMixedInstancesPolicyArgs']]] = None,
                 new_instances_protected_from_scale_in: Optional[pulumi.Input[bool]] = None,
                 notification_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupNotificationConfigurationArgs']]]]] = None,
                 placement_group: Optional[pulumi.Input[str]] = None,
                 service_linked_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoScalingGroupTagPropertyArgs']]]]] = None,
                 target_group_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 termination_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 v_pc_zone_identifier: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        pulumi.log.warn("""AutoScalingGroup is deprecated: AutoScalingGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoScalingGroupArgs.__new__(AutoScalingGroupArgs)

            __props__.__dict__["auto_scaling_group_name"] = auto_scaling_group_name
            __props__.__dict__["availability_zones"] = availability_zones
            __props__.__dict__["capacity_rebalance"] = capacity_rebalance
            __props__.__dict__["context"] = context
            __props__.__dict__["cooldown"] = cooldown
            __props__.__dict__["default_instance_warmup"] = default_instance_warmup
            __props__.__dict__["desired_capacity"] = desired_capacity
            __props__.__dict__["desired_capacity_type"] = desired_capacity_type
            __props__.__dict__["health_check_grace_period"] = health_check_grace_period
            __props__.__dict__["health_check_type"] = health_check_type
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["launch_configuration_name"] = launch_configuration_name
            __props__.__dict__["launch_template"] = launch_template
            __props__.__dict__["lifecycle_hook_specification_list"] = lifecycle_hook_specification_list
            __props__.__dict__["load_balancer_names"] = load_balancer_names
            __props__.__dict__["max_instance_lifetime"] = max_instance_lifetime
            if max_size is None and not opts.urn:
                raise TypeError("Missing required property 'max_size'")
            __props__.__dict__["max_size"] = max_size
            __props__.__dict__["metrics_collection"] = metrics_collection
            if min_size is None and not opts.urn:
                raise TypeError("Missing required property 'min_size'")
            __props__.__dict__["min_size"] = min_size
            __props__.__dict__["mixed_instances_policy"] = mixed_instances_policy
            __props__.__dict__["new_instances_protected_from_scale_in"] = new_instances_protected_from_scale_in
            __props__.__dict__["notification_configurations"] = notification_configurations
            __props__.__dict__["placement_group"] = placement_group
            __props__.__dict__["service_linked_role_arn"] = service_linked_role_arn
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_group_arns"] = target_group_arns
            __props__.__dict__["termination_policies"] = termination_policies
            __props__.__dict__["v_pc_zone_identifier"] = v_pc_zone_identifier
            __props__.__dict__["launch_template_specification"] = None
        super(AutoScalingGroup, __self__).__init__(
            'aws-native:autoscaling:AutoScalingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AutoScalingGroup':
        """
        Get an existing AutoScalingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AutoScalingGroupArgs.__new__(AutoScalingGroupArgs)

        __props__.__dict__["auto_scaling_group_name"] = None
        __props__.__dict__["availability_zones"] = None
        __props__.__dict__["capacity_rebalance"] = None
        __props__.__dict__["context"] = None
        __props__.__dict__["cooldown"] = None
        __props__.__dict__["default_instance_warmup"] = None
        __props__.__dict__["desired_capacity"] = None
        __props__.__dict__["desired_capacity_type"] = None
        __props__.__dict__["health_check_grace_period"] = None
        __props__.__dict__["health_check_type"] = None
        __props__.__dict__["instance_id"] = None
        __props__.__dict__["launch_configuration_name"] = None
        __props__.__dict__["launch_template"] = None
        __props__.__dict__["launch_template_specification"] = None
        __props__.__dict__["lifecycle_hook_specification_list"] = None
        __props__.__dict__["load_balancer_names"] = None
        __props__.__dict__["max_instance_lifetime"] = None
        __props__.__dict__["max_size"] = None
        __props__.__dict__["metrics_collection"] = None
        __props__.__dict__["min_size"] = None
        __props__.__dict__["mixed_instances_policy"] = None
        __props__.__dict__["new_instances_protected_from_scale_in"] = None
        __props__.__dict__["notification_configurations"] = None
        __props__.__dict__["placement_group"] = None
        __props__.__dict__["service_linked_role_arn"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_group_arns"] = None
        __props__.__dict__["termination_policies"] = None
        __props__.__dict__["v_pc_zone_identifier"] = None
        return AutoScalingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoScalingGroupName")
    def auto_scaling_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "auto_scaling_group_name")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="capacityRebalance")
    def capacity_rebalance(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "capacity_rebalance")

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def cooldown(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cooldown")

    @property
    @pulumi.getter(name="defaultInstanceWarmup")
    def default_instance_warmup(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "default_instance_warmup")

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="desiredCapacityType")
    def desired_capacity_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "desired_capacity_type")

    @property
    @pulumi.getter(name="healthCheckGracePeriod")
    def health_check_grace_period(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "health_check_grace_period")

    @property
    @pulumi.getter(name="healthCheckType")
    def health_check_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "health_check_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="launchConfigurationName")
    def launch_configuration_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "launch_configuration_name")

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> pulumi.Output[Optional['outputs.AutoScalingGroupLaunchTemplateSpecification']]:
        return pulumi.get(self, "launch_template")

    @property
    @pulumi.getter(name="launchTemplateSpecification")
    def launch_template_specification(self) -> pulumi.Output[str]:
        return pulumi.get(self, "launch_template_specification")

    @property
    @pulumi.getter(name="lifecycleHookSpecificationList")
    def lifecycle_hook_specification_list(self) -> pulumi.Output[Optional[Sequence['outputs.AutoScalingGroupLifecycleHookSpecification']]]:
        return pulumi.get(self, "lifecycle_hook_specification_list")

    @property
    @pulumi.getter(name="loadBalancerNames")
    def load_balancer_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "load_balancer_names")

    @property
    @pulumi.getter(name="maxInstanceLifetime")
    def max_instance_lifetime(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "max_instance_lifetime")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Output[str]:
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="metricsCollection")
    def metrics_collection(self) -> pulumi.Output[Optional[Sequence['outputs.AutoScalingGroupMetricsCollection']]]:
        return pulumi.get(self, "metrics_collection")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Output[str]:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter(name="mixedInstancesPolicy")
    def mixed_instances_policy(self) -> pulumi.Output[Optional['outputs.AutoScalingGroupMixedInstancesPolicy']]:
        return pulumi.get(self, "mixed_instances_policy")

    @property
    @pulumi.getter(name="newInstancesProtectedFromScaleIn")
    def new_instances_protected_from_scale_in(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "new_instances_protected_from_scale_in")

    @property
    @pulumi.getter(name="notificationConfigurations")
    def notification_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.AutoScalingGroupNotificationConfiguration']]]:
        return pulumi.get(self, "notification_configurations")

    @property
    @pulumi.getter(name="placementGroup")
    def placement_group(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "placement_group")

    @property
    @pulumi.getter(name="serviceLinkedRoleARN")
    def service_linked_role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "service_linked_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.AutoScalingGroupTagProperty']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetGroupARNs")
    def target_group_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "target_group_arns")

    @property
    @pulumi.getter(name="terminationPolicies")
    def termination_policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "termination_policies")

    @property
    @pulumi.getter(name="vPCZoneIdentifier")
    def v_pc_zone_identifier(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "v_pc_zone_identifier")
