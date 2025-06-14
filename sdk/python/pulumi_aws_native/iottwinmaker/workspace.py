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

__all__ = ['WorkspaceArgs', 'Workspace']

@pulumi.input_type
class WorkspaceArgs:
    def __init__(__self__, *,
                 role: pulumi.Input[builtins.str],
                 s3_location: pulumi.Input[builtins.str],
                 workspace_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Workspace resource.
        :param pulumi.Input[builtins.str] role: The ARN of the execution role associated with the workspace.
        :param pulumi.Input[builtins.str] s3_location: The ARN of the S3 bucket where resources associated with the workspace are stored.
        :param pulumi.Input[builtins.str] workspace_id: The ID of the workspace.
        :param pulumi.Input[builtins.str] description: The description of the workspace.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A map of key-value pairs to associate with a resource.
        """
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "s3_location", s3_location)
        pulumi.set(__self__, "workspace_id", workspace_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[builtins.str]:
        """
        The ARN of the execution role associated with the workspace.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> pulumi.Input[builtins.str]:
        """
        The ARN of the S3 bucket where resources associated with the workspace are stored.
        """
        return pulumi.get(self, "s3_location")

    @s3_location.setter
    def s3_location(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "s3_location", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the workspace.
        """
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map of key-value pairs to associate with a resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:iottwinmaker:Workspace")
class Workspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 s3_location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::IoTTwinMaker::Workspace

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The description of the workspace.
        :param pulumi.Input[builtins.str] role: The ARN of the execution role associated with the workspace.
        :param pulumi.Input[builtins.str] s3_location: The ARN of the S3 bucket where resources associated with the workspace are stored.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: A map of key-value pairs to associate with a resource.
        :param pulumi.Input[builtins.str] workspace_id: The ID of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkspaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::IoTTwinMaker::Workspace

        :param str resource_name: The name of the resource.
        :param WorkspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 s3_location: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 workspace_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

            __props__.__dict__["description"] = description
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if s3_location is None and not opts.urn:
                raise TypeError("Missing required property 's3_location'")
            __props__.__dict__["s3_location"] = s3_location
            __props__.__dict__["tags"] = tags
            if workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_id'")
            __props__.__dict__["workspace_id"] = workspace_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_date_time"] = None
            __props__.__dict__["update_date_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["workspaceId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Workspace, __self__).__init__(
            'aws-native:iottwinmaker:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_date_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["s3_location"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["update_date_time"] = None
        __props__.__dict__["workspace_id"] = None
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the workspace.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationDateTime")
    def creation_date_time(self) -> pulumi.Output[builtins.str]:
        """
        The date and time when the workspace was created.
        """
        return pulumi.get(self, "creation_date_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the execution role associated with the workspace.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="s3Location")
    def s3_location(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the S3 bucket where resources associated with the workspace are stored.
        """
        return pulumi.get(self, "s3_location")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A map of key-value pairs to associate with a resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateDateTime")
    def update_date_time(self) -> pulumi.Output[builtins.str]:
        """
        The date and time of the current update.
        """
        return pulumi.get(self, "update_date_time")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the workspace.
        """
        return pulumi.get(self, "workspace_id")

