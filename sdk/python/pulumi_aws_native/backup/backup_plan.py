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

__all__ = ['BackupPlanArgs', 'BackupPlan']

@pulumi.input_type
class BackupPlanArgs:
    def __init__(__self__, *,
                 backup_plan: pulumi.Input['BackupPlanResourceTypeArgs'],
                 backup_plan_tags: Optional[Any] = None):
        """
        The set of arguments for constructing a BackupPlan resource.
        """
        pulumi.set(__self__, "backup_plan", backup_plan)
        if backup_plan_tags is not None:
            pulumi.set(__self__, "backup_plan_tags", backup_plan_tags)

    @property
    @pulumi.getter(name="backupPlan")
    def backup_plan(self) -> pulumi.Input['BackupPlanResourceTypeArgs']:
        return pulumi.get(self, "backup_plan")

    @backup_plan.setter
    def backup_plan(self, value: pulumi.Input['BackupPlanResourceTypeArgs']):
        pulumi.set(self, "backup_plan", value)

    @property
    @pulumi.getter(name="backupPlanTags")
    def backup_plan_tags(self) -> Optional[Any]:
        return pulumi.get(self, "backup_plan_tags")

    @backup_plan_tags.setter
    def backup_plan_tags(self, value: Optional[Any]):
        pulumi.set(self, "backup_plan_tags", value)


class BackupPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan: Optional[pulumi.Input[pulumi.InputType['BackupPlanResourceTypeArgs']]] = None,
                 backup_plan_tags: Optional[Any] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Backup::BackupPlan

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Backup::BackupPlan

        :param str resource_name: The name of the resource.
        :param BackupPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan: Optional[pulumi.Input[pulumi.InputType['BackupPlanResourceTypeArgs']]] = None,
                 backup_plan_tags: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupPlanArgs.__new__(BackupPlanArgs)

            if backup_plan is None and not opts.urn:
                raise TypeError("Missing required property 'backup_plan'")
            __props__.__dict__["backup_plan"] = backup_plan
            __props__.__dict__["backup_plan_tags"] = backup_plan_tags
            __props__.__dict__["backup_plan_arn"] = None
            __props__.__dict__["backup_plan_id"] = None
            __props__.__dict__["version_id"] = None
        super(BackupPlan, __self__).__init__(
            'aws-native:backup:BackupPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BackupPlan':
        """
        Get an existing BackupPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BackupPlanArgs.__new__(BackupPlanArgs)

        __props__.__dict__["backup_plan"] = None
        __props__.__dict__["backup_plan_arn"] = None
        __props__.__dict__["backup_plan_id"] = None
        __props__.__dict__["backup_plan_tags"] = None
        __props__.__dict__["version_id"] = None
        return BackupPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPlan")
    def backup_plan(self) -> pulumi.Output['outputs.BackupPlanResourceType']:
        return pulumi.get(self, "backup_plan")

    @property
    @pulumi.getter(name="backupPlanArn")
    def backup_plan_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "backup_plan_arn")

    @property
    @pulumi.getter(name="backupPlanId")
    def backup_plan_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "backup_plan_id")

    @property
    @pulumi.getter(name="backupPlanTags")
    def backup_plan_tags(self) -> pulumi.Output[Optional[Any]]:
        return pulumi.get(self, "backup_plan_tags")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "version_id")
