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
from ._enums import *
from ._inputs import *

__all__ = ['ProjectProfileArgs', 'ProjectProfile']

@pulumi.input_type
class ProjectProfileArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 domain_unit_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProfileEnvironmentConfigurationArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input['ProjectProfileStatus']] = None):
        """
        The set of arguments for constructing a ProjectProfile resource.
        :param pulumi.Input[builtins.str] description: The description of the project profile.
        :param pulumi.Input[builtins.str] name: The name of a project profile.
        :param pulumi.Input['ProjectProfileStatus'] status: The status of a project profile.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_identifier is not None:
            pulumi.set(__self__, "domain_identifier", domain_identifier)
        if domain_unit_identifier is not None:
            pulumi.set(__self__, "domain_unit_identifier", domain_unit_identifier)
        if environment_configurations is not None:
            pulumi.set(__self__, "environment_configurations", environment_configurations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the project profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "domain_identifier")

    @domain_identifier.setter
    def domain_identifier(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_identifier", value)

    @property
    @pulumi.getter(name="domainUnitIdentifier")
    def domain_unit_identifier(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "domain_unit_identifier")

    @domain_unit_identifier.setter
    def domain_unit_identifier(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_unit_identifier", value)

    @property
    @pulumi.getter(name="environmentConfigurations")
    def environment_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProfileEnvironmentConfigurationArgs']]]]:
        return pulumi.get(self, "environment_configurations")

    @environment_configurations.setter
    def environment_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectProfileEnvironmentConfigurationArgs']]]]):
        pulumi.set(self, "environment_configurations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of a project profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['ProjectProfileStatus']]:
        """
        The status of a project profile.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['ProjectProfileStatus']]):
        pulumi.set(self, "status", value)


@pulumi.type_token("aws-native:datazone:ProjectProfile")
class ProjectProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 domain_unit_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectProfileEnvironmentConfigurationArgs', 'ProjectProfileEnvironmentConfigurationArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input['ProjectProfileStatus']] = None,
                 __props__=None):
        """
        Definition of AWS::DataZone::ProjectProfile Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The description of the project profile.
        :param pulumi.Input[builtins.str] name: The name of a project profile.
        :param pulumi.Input['ProjectProfileStatus'] status: The status of a project profile.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::DataZone::ProjectProfile Resource Type

        :param str resource_name: The name of the resource.
        :param ProjectProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 domain_unit_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 environment_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectProfileEnvironmentConfigurationArgs', 'ProjectProfileEnvironmentConfigurationArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input['ProjectProfileStatus']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectProfileArgs.__new__(ProjectProfileArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["domain_identifier"] = domain_identifier
            __props__.__dict__["domain_unit_identifier"] = domain_unit_identifier
            __props__.__dict__["environment_configurations"] = environment_configurations
            __props__.__dict__["name"] = name
            __props__.__dict__["status"] = status
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["created_by"] = None
            __props__.__dict__["domain_id"] = None
            __props__.__dict__["domain_unit_id"] = None
            __props__.__dict__["identifier"] = None
            __props__.__dict__["last_updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domainIdentifier"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ProjectProfile, __self__).__init__(
            'aws-native:datazone:ProjectProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProjectProfile':
        """
        Get an existing ProjectProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProjectProfileArgs.__new__(ProjectProfileArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["created_by"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["domain_id"] = None
        __props__.__dict__["domain_identifier"] = None
        __props__.__dict__["domain_unit_id"] = None
        __props__.__dict__["domain_unit_identifier"] = None
        __props__.__dict__["environment_configurations"] = None
        __props__.__dict__["identifier"] = None
        __props__.__dict__["last_updated_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["status"] = None
        return ProjectProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the project profile.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The timestamp of when the project profile was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[builtins.str]:
        """
        The user who created the project profile.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the project profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[builtins.str]:
        """
        The domain ID of the project profile.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="domainIdentifier")
    def domain_identifier(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "domain_identifier")

    @property
    @pulumi.getter(name="domainUnitId")
    def domain_unit_id(self) -> pulumi.Output[builtins.str]:
        """
        The domain unit ID of the project profile.
        """
        return pulumi.get(self, "domain_unit_id")

    @property
    @pulumi.getter(name="domainUnitIdentifier")
    def domain_unit_identifier(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "domain_unit_identifier")

    @property
    @pulumi.getter(name="environmentConfigurations")
    def environment_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.ProjectProfileEnvironmentConfiguration']]]:
        return pulumi.get(self, "environment_configurations")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="lastUpdatedAt")
    def last_updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The timestamp at which a project profile was last updated.
        """
        return pulumi.get(self, "last_updated_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of a project profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['ProjectProfileStatus']]:
        """
        The status of a project profile.
        """
        return pulumi.get(self, "status")

