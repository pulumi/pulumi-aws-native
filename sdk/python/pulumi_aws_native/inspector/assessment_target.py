# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AssessmentTargetArgs', 'AssessmentTarget']

@pulumi.input_type
class AssessmentTargetArgs:
    def __init__(__self__, *,
                 assessment_target_name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AssessmentTarget resource.
        """
        if assessment_target_name is not None:
            pulumi.set(__self__, "assessment_target_name", assessment_target_name)
        if resource_group_arn is not None:
            pulumi.set(__self__, "resource_group_arn", resource_group_arn)

    @property
    @pulumi.getter(name="assessmentTargetName")
    def assessment_target_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "assessment_target_name")

    @assessment_target_name.setter
    def assessment_target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "assessment_target_name", value)

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_group_arn")

    @resource_group_arn.setter
    def resource_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_arn", value)


class AssessmentTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assessment_target_name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Inspector::AssessmentTarget

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AssessmentTargetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Inspector::AssessmentTarget

        :param str resource_name: The name of the resource.
        :param AssessmentTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssessmentTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assessment_target_name: Optional[pulumi.Input[str]] = None,
                 resource_group_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssessmentTargetArgs.__new__(AssessmentTargetArgs)

            __props__.__dict__["assessment_target_name"] = assessment_target_name
            __props__.__dict__["resource_group_arn"] = resource_group_arn
            __props__.__dict__["arn"] = None
        super(AssessmentTarget, __self__).__init__(
            'aws-native:inspector:AssessmentTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AssessmentTarget':
        """
        Get an existing AssessmentTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AssessmentTargetArgs.__new__(AssessmentTargetArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["assessment_target_name"] = None
        __props__.__dict__["resource_group_arn"] = None
        return AssessmentTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assessmentTargetName")
    def assessment_target_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "assessment_target_name")

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "resource_group_arn")
