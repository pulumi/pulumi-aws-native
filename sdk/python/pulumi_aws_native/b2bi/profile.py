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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = ['ProfileArgs', 'Profile']

@pulumi.input_type
class ProfileArgs:
    def __init__(__self__, *,
                 business_name: pulumi.Input[builtins.str],
                 logging: pulumi.Input['ProfileLogging'],
                 phone: pulumi.Input[builtins.str],
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Profile resource.
        :param pulumi.Input[builtins.str] business_name: Returns the name for the business associated with this profile.
        :param pulumi.Input['ProfileLogging'] logging: Specifies whether or not logging is enabled for this profile.
        :param pulumi.Input[builtins.str] name: Returns the display name for profile.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A key-value pair for a specific profile. Tags are metadata that you can use to search for and group capabilities for various purposes.
        """
        pulumi.set(__self__, "business_name", business_name)
        pulumi.set(__self__, "logging", logging)
        pulumi.set(__self__, "phone", phone)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="businessName")
    def business_name(self) -> pulumi.Input[builtins.str]:
        """
        Returns the name for the business associated with this profile.
        """
        return pulumi.get(self, "business_name")

    @business_name.setter
    def business_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "business_name", value)

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Input['ProfileLogging']:
        """
        Specifies whether or not logging is enabled for this profile.
        """
        return pulumi.get(self, "logging")

    @logging.setter
    def logging(self, value: pulumi.Input['ProfileLogging']):
        pulumi.set(self, "logging", value)

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Returns the display name for profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A key-value pair for a specific profile. Tags are metadata that you can use to search for and group capabilities for various purposes.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:b2bi:Profile")
class Profile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 business_name: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 logging: Optional[pulumi.Input['ProfileLogging']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::B2BI::Profile Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] business_name: Returns the name for the business associated with this profile.
        :param pulumi.Input['ProfileLogging'] logging: Specifies whether or not logging is enabled for this profile.
        :param pulumi.Input[builtins.str] name: Returns the display name for profile.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: A key-value pair for a specific profile. Tags are metadata that you can use to search for and group capabilities for various purposes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::B2BI::Profile Resource Type

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
                 business_name: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 logging: Optional[pulumi.Input['ProfileLogging']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 phone: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileArgs.__new__(ProfileArgs)

            if business_name is None and not opts.urn:
                raise TypeError("Missing required property 'business_name'")
            __props__.__dict__["business_name"] = business_name
            __props__.__dict__["email"] = email
            if logging is None and not opts.urn:
                raise TypeError("Missing required property 'logging'")
            __props__.__dict__["logging"] = logging
            __props__.__dict__["name"] = name
            if phone is None and not opts.urn:
                raise TypeError("Missing required property 'phone'")
            __props__.__dict__["phone"] = phone
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["log_group_name"] = None
            __props__.__dict__["modified_at"] = None
            __props__.__dict__["profile_arn"] = None
            __props__.__dict__["profile_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["logging"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Profile, __self__).__init__(
            'aws-native:b2bi:Profile',
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

        __props__.__dict__["business_name"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["email"] = None
        __props__.__dict__["log_group_name"] = None
        __props__.__dict__["logging"] = None
        __props__.__dict__["modified_at"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["phone"] = None
        __props__.__dict__["profile_arn"] = None
        __props__.__dict__["profile_id"] = None
        __props__.__dict__["tags"] = None
        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="businessName")
    def business_name(self) -> pulumi.Output[builtins.str]:
        """
        Returns the name for the business associated with this profile.
        """
        return pulumi.get(self, "business_name")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Returns the timestamp for creation date and time of the profile.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[builtins.str]:
        """
        Returns the name of the logging group.
        """
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter
    def logging(self) -> pulumi.Output['ProfileLogging']:
        """
        Specifies whether or not logging is enabled for this profile.
        """
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter(name="modifiedAt")
    def modified_at(self) -> pulumi.Output[builtins.str]:
        """
        Returns the timestamp that identifies the most recent date and time that the profile was modified.
        """
        return pulumi.get(self, "modified_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Returns the display name for profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter(name="profileArn")
    def profile_arn(self) -> pulumi.Output[builtins.str]:
        """
        Returns an Amazon Resource Name (ARN) for the profile.
        """
        return pulumi.get(self, "profile_arn")

    @property
    @pulumi.getter(name="profileId")
    def profile_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "profile_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A key-value pair for a specific profile. Tags are metadata that you can use to search for and group capabilities for various purposes.
        """
        return pulumi.get(self, "tags")

