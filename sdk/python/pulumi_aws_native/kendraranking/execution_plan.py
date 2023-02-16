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

__all__ = ['ExecutionPlanArgs', 'ExecutionPlan']

@pulumi.input_type
class ExecutionPlanArgs:
    def __init__(__self__, *,
                 capacity_units: Optional[pulumi.Input['ExecutionPlanCapacityUnitsConfigurationArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ExecutionPlanTagArgs']]]] = None):
        """
        The set of arguments for constructing a ExecutionPlan resource.
        :param pulumi.Input['ExecutionPlanCapacityUnitsConfigurationArgs'] capacity_units: Capacity units
        :param pulumi.Input[str] description: A description for the execution plan
        :param pulumi.Input[Sequence[pulumi.Input['ExecutionPlanTagArgs']]] tags: Tags for labeling the execution plan
        """
        if capacity_units is not None:
            pulumi.set(__self__, "capacity_units", capacity_units)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="capacityUnits")
    def capacity_units(self) -> Optional[pulumi.Input['ExecutionPlanCapacityUnitsConfigurationArgs']]:
        """
        Capacity units
        """
        return pulumi.get(self, "capacity_units")

    @capacity_units.setter
    def capacity_units(self, value: Optional[pulumi.Input['ExecutionPlanCapacityUnitsConfigurationArgs']]):
        pulumi.set(self, "capacity_units", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the execution plan
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExecutionPlanTagArgs']]]]:
        """
        Tags for labeling the execution plan
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExecutionPlanTagArgs']]]]):
        pulumi.set(self, "tags", value)


class ExecutionPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_units: Optional[pulumi.Input[pulumi.InputType['ExecutionPlanCapacityUnitsConfigurationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExecutionPlanTagArgs']]]]] = None,
                 __props__=None):
        """
        A KendraRanking Rescore execution plan

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ExecutionPlanCapacityUnitsConfigurationArgs']] capacity_units: Capacity units
        :param pulumi.Input[str] description: A description for the execution plan
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExecutionPlanTagArgs']]]] tags: Tags for labeling the execution plan
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExecutionPlanArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A KendraRanking Rescore execution plan

        :param str resource_name: The name of the resource.
        :param ExecutionPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExecutionPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_units: Optional[pulumi.Input[pulumi.InputType['ExecutionPlanCapacityUnitsConfigurationArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExecutionPlanTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExecutionPlanArgs.__new__(ExecutionPlanArgs)

            __props__.__dict__["capacity_units"] = capacity_units
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(ExecutionPlan, __self__).__init__(
            'aws-native:kendraranking:ExecutionPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExecutionPlan':
        """
        Get an existing ExecutionPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExecutionPlanArgs.__new__(ExecutionPlanArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["capacity_units"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return ExecutionPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="capacityUnits")
    def capacity_units(self) -> pulumi.Output[Optional['outputs.ExecutionPlanCapacityUnitsConfiguration']]:
        """
        Capacity units
        """
        return pulumi.get(self, "capacity_units")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the execution plan
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ExecutionPlanTag']]]:
        """
        Tags for labeling the execution plan
        """
        return pulumi.get(self, "tags")
