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

__all__ = ['ModelPackageGroupArgs', 'ModelPackageGroup']

@pulumi.input_type
class ModelPackageGroupArgs:
    def __init__(__self__, *,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 model_package_group_policy: Optional[Any] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ModelPackageGroupTagArgs']]]] = None):
        """
        The set of arguments for constructing a ModelPackageGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['ModelPackageGroupTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if model_package_group_description is not None:
            pulumi.set(__self__, "model_package_group_description", model_package_group_description)
        if model_package_group_name is not None:
            pulumi.set(__self__, "model_package_group_name", model_package_group_name)
        if model_package_group_policy is not None:
            pulumi.set(__self__, "model_package_group_policy", model_package_group_policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="modelPackageGroupDescription")
    def model_package_group_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "model_package_group_description")

    @model_package_group_description.setter
    def model_package_group_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_package_group_description", value)

    @property
    @pulumi.getter(name="modelPackageGroupName")
    def model_package_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "model_package_group_name")

    @model_package_group_name.setter
    def model_package_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "model_package_group_name", value)

    @property
    @pulumi.getter(name="modelPackageGroupPolicy")
    def model_package_group_policy(self) -> Optional[Any]:
        return pulumi.get(self, "model_package_group_policy")

    @model_package_group_policy.setter
    def model_package_group_policy(self, value: Optional[Any]):
        pulumi.set(self, "model_package_group_policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ModelPackageGroupTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ModelPackageGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)


class ModelPackageGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 model_package_group_policy: Optional[Any] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelPackageGroupTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SageMaker::ModelPackageGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelPackageGroupTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ModelPackageGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SageMaker::ModelPackageGroup

        :param str resource_name: The name of the resource.
        :param ModelPackageGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ModelPackageGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 model_package_group_description: Optional[pulumi.Input[str]] = None,
                 model_package_group_name: Optional[pulumi.Input[str]] = None,
                 model_package_group_policy: Optional[Any] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelPackageGroupTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModelPackageGroupArgs.__new__(ModelPackageGroupArgs)

            __props__.__dict__["model_package_group_description"] = model_package_group_description
            __props__.__dict__["model_package_group_name"] = model_package_group_name
            __props__.__dict__["model_package_group_policy"] = model_package_group_policy
            __props__.__dict__["tags"] = tags
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["model_package_group_arn"] = None
            __props__.__dict__["model_package_group_status"] = None
        super(ModelPackageGroup, __self__).__init__(
            'aws-native:sagemaker:ModelPackageGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ModelPackageGroup':
        """
        Get an existing ModelPackageGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ModelPackageGroupArgs.__new__(ModelPackageGroupArgs)

        __props__.__dict__["creation_time"] = None
        __props__.__dict__["model_package_group_arn"] = None
        __props__.__dict__["model_package_group_description"] = None
        __props__.__dict__["model_package_group_name"] = None
        __props__.__dict__["model_package_group_policy"] = None
        __props__.__dict__["model_package_group_status"] = None
        __props__.__dict__["tags"] = None
        return ModelPackageGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time at which the model package group was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="modelPackageGroupArn")
    def model_package_group_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "model_package_group_arn")

    @property
    @pulumi.getter(name="modelPackageGroupDescription")
    def model_package_group_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "model_package_group_description")

    @property
    @pulumi.getter(name="modelPackageGroupName")
    def model_package_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "model_package_group_name")

    @property
    @pulumi.getter(name="modelPackageGroupPolicy")
    def model_package_group_policy(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "model_package_group_policy")

    @property
    @pulumi.getter(name="modelPackageGroupStatus")
    def model_package_group_status(self) -> pulumi.Output['ModelPackageGroupStatus']:
        """
        The status of a modelpackage group job.
        """
        return pulumi.get(self, "model_package_group_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ModelPackageGroupTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")
