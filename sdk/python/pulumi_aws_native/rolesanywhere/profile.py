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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 role_arns: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 accept_role_session_name: Optional[pulumi.Input[builtins.bool]] = None,
                 attribute_mappings: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileAttributeMappingArgs']]]] = None,
                 duration_seconds: Optional[pulumi.Input[builtins.float]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 require_instance_properties: Optional[pulumi.Input[builtins.bool]] = None,
                 session_policy: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_arns: A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
        :param pulumi.Input[builtins.bool] accept_role_session_name: Used to determine if a custom role session name will be accepted in a temporary credential request.
        :param pulumi.Input[Sequence[pulumi.Input['ProfileAttributeMappingArgs']]] attribute_mappings: A mapping applied to the authenticating end-entity certificate.
        :param pulumi.Input[builtins.float] duration_seconds: The number of seconds vended session credentials will be valid for
        :param pulumi.Input[builtins.bool] enabled: The enabled status of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] managed_policy_arns: A list of managed policy ARNs. Managed policies identified by this list will be applied to the vended session credentials.
        :param pulumi.Input[builtins.str] name: The customer specified name of the resource.
        :param pulumi.Input[builtins.bool] require_instance_properties: Specifies whether instance properties are required in CreateSession requests with this profile.
        :param pulumi.Input[builtins.str] session_policy: A session policy that will applied to the trust boundary of the vended session credentials.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A list of Tags.
        """
        pulumi.set(__self__, "role_arns", role_arns)
        if accept_role_session_name is not None:
            pulumi.set(__self__, "accept_role_session_name", accept_role_session_name)
        if attribute_mappings is not None:
            pulumi.set(__self__, "attribute_mappings", attribute_mappings)
        if duration_seconds is not None:
            pulumi.set(__self__, "duration_seconds", duration_seconds)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if managed_policy_arns is not None:
            pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if require_instance_properties is not None:
            pulumi.set(__self__, "require_instance_properties", require_instance_properties)
        if session_policy is not None:
            pulumi.set(__self__, "session_policy", session_policy)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="roleArns")
    def role_arns(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
        """
        return pulumi.get(self, "role_arns")

    @role_arns.setter
    def role_arns(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "role_arns", value)

    @property
    @pulumi.getter(name="acceptRoleSessionName")
    def accept_role_session_name(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Used to determine if a custom role session name will be accepted in a temporary credential request.
        """
        return pulumi.get(self, "accept_role_session_name")

    @accept_role_session_name.setter
    def accept_role_session_name(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "accept_role_session_name", value)

    @property
    @pulumi.getter(name="attributeMappings")
    def attribute_mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProfileAttributeMappingArgs']]]]:
        """
        A mapping applied to the authenticating end-entity certificate.
        """
        return pulumi.get(self, "attribute_mappings")

    @attribute_mappings.setter
    def attribute_mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProfileAttributeMappingArgs']]]]):
        pulumi.set(self, "attribute_mappings", value)

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> Optional[pulumi.Input[builtins.float]]:
        """
        The number of seconds vended session credentials will be valid for
        """
        return pulumi.get(self, "duration_seconds")

    @duration_seconds.setter
    def duration_seconds(self, value: Optional[pulumi.Input[builtins.float]]):
        pulumi.set(self, "duration_seconds", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        The enabled status of the resource.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of managed policy ARNs. Managed policies identified by this list will be applied to the vended session credentials.
        """
        return pulumi.get(self, "managed_policy_arns")

    @managed_policy_arns.setter
    def managed_policy_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "managed_policy_arns", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The customer specified name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requireInstanceProperties")
    def require_instance_properties(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether instance properties are required in CreateSession requests with this profile.
        """
        return pulumi.get(self, "require_instance_properties")

    @require_instance_properties.setter
    def require_instance_properties(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "require_instance_properties", value)

    @property
    @pulumi.getter(name="sessionPolicy")
    def session_policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A session policy that will applied to the trust boundary of the vended session credentials.
        """
        return pulumi.get(self, "session_policy")

    @session_policy.setter
    def session_policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "session_policy", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A list of Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:rolesanywhere:Profile")
class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_role_session_name: Optional[pulumi.Input[builtins.bool]] = None,
                 attribute_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProfileAttributeMappingArgs', 'ProfileAttributeMappingArgsDict']]]]] = None,
                 duration_seconds: Optional[pulumi.Input[builtins.float]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 require_instance_properties: Optional[pulumi.Input[builtins.bool]] = None,
                 role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 session_policy: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::RolesAnywhere::Profile Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] accept_role_session_name: Used to determine if a custom role session name will be accepted in a temporary credential request.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProfileAttributeMappingArgs', 'ProfileAttributeMappingArgsDict']]]] attribute_mappings: A mapping applied to the authenticating end-entity certificate.
        :param pulumi.Input[builtins.float] duration_seconds: The number of seconds vended session credentials will be valid for
        :param pulumi.Input[builtins.bool] enabled: The enabled status of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] managed_policy_arns: A list of managed policy ARNs. Managed policies identified by this list will be applied to the vended session credentials.
        :param pulumi.Input[builtins.str] name: The customer specified name of the resource.
        :param pulumi.Input[builtins.bool] require_instance_properties: Specifies whether instance properties are required in CreateSession requests with this profile.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] role_arns: A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
        :param pulumi.Input[builtins.str] session_policy: A session policy that will applied to the trust boundary of the vended session credentials.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: A list of Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::RolesAnywhere::Profile Resource Type

        :param str resource_name: The name of the resource.
        :param ProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_role_session_name: Optional[pulumi.Input[builtins.bool]] = None,
                 attribute_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProfileAttributeMappingArgs', 'ProfileAttributeMappingArgsDict']]]]] = None,
                 duration_seconds: Optional[pulumi.Input[builtins.float]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 require_instance_properties: Optional[pulumi.Input[builtins.bool]] = None,
                 role_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 session_policy: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            __props__.__dict__["accept_role_session_name"] = accept_role_session_name
            __props__.__dict__["attribute_mappings"] = attribute_mappings
            __props__.__dict__["duration_seconds"] = duration_seconds
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["managed_policy_arns"] = managed_policy_arns
            __props__.__dict__["name"] = name
            __props__.__dict__["require_instance_properties"] = require_instance_properties
            if role_arns is None and not opts.urn:
                raise TypeError("Missing required property 'role_arns'")
            __props__.__dict__["role_arns"] = role_arns
            __props__.__dict__["session_policy"] = session_policy
            __props__.__dict__["tags"] = tags
            __props__.__dict__["profile_arn"] = None
            __props__.__dict__["profile_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["requireInstanceProperties"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Profile, __self__).__init__(
            'aws-native:rolesanywhere:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ProfileArgs.__new__(ProfileArgs)

        __props__.__dict__["accept_role_session_name"] = None
        __props__.__dict__["attribute_mappings"] = None
        __props__.__dict__["duration_seconds"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["managed_policy_arns"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["profile_arn"] = None
        __props__.__dict__["profile_id"] = None
        __props__.__dict__["require_instance_properties"] = None
        __props__.__dict__["role_arns"] = None
        __props__.__dict__["session_policy"] = None
        __props__.__dict__["tags"] = None
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptRoleSessionName")
    def accept_role_session_name(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Used to determine if a custom role session name will be accepted in a temporary credential request.
        """
        return pulumi.get(self, "accept_role_session_name")

    @property
    @pulumi.getter(name="attributeMappings")
    def attribute_mappings(self) -> pulumi.Output[Optional[Sequence['outputs.ProfileAttributeMapping']]]:
        """
        A mapping applied to the authenticating end-entity certificate.
        """
        return pulumi.get(self, "attribute_mappings")

    @property
    @pulumi.getter(name="durationSeconds")
    def duration_seconds(self) -> pulumi.Output[Optional[builtins.float]]:
        """
        The number of seconds vended session credentials will be valid for
        """
        return pulumi.get(self, "duration_seconds")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        The enabled status of the resource.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A list of managed policy ARNs. Managed policies identified by this list will be applied to the vended session credentials.
        """
        return pulumi.get(self, "managed_policy_arns")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The customer specified name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileArn")
    def profile_arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the profile.
        """
        return pulumi.get(self, "profile_arn")

    @property
    @pulumi.getter(name="profileId")
    def profile_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique primary identifier of the Profile
        """
        return pulumi.get(self, "profile_id")

    @property
    @pulumi.getter(name="requireInstanceProperties")
    def require_instance_properties(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether instance properties are required in CreateSession requests with this profile.
        """
        return pulumi.get(self, "require_instance_properties")

    @property
    @pulumi.getter(name="roleArns")
    def role_arns(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.
        """
        return pulumi.get(self, "role_arns")

    @property
    @pulumi.getter(name="sessionPolicy")
    def session_policy(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A session policy that will applied to the trust boundary of the vended session credentials.
        """
        return pulumi.get(self, "session_policy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A list of Tags.
        """
        return pulumi.get(self, "tags")

