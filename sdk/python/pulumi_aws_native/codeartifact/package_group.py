# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['PackageGroupArgs', 'PackageGroup']

@pulumi.input_type
class PackageGroupArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 pattern: pulumi.Input[str],
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 origin_configuration: Optional[pulumi.Input['PackageGroupOriginConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a PackageGroup resource.
        :param pulumi.Input[str] domain_name: The name of the domain that contains the package group.
        :param pulumi.Input[str] pattern: The package group pattern that is used to gather packages.
        :param pulumi.Input[str] contact_info: The contact info of the package group.
        :param pulumi.Input[str] description: The text description of the package group.
        :param pulumi.Input[str] domain_owner: The 12-digit account ID of the AWS account that owns the domain.
        :param pulumi.Input['PackageGroupOriginConfigurationArgs'] origin_configuration: The package origin configuration of the package group.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to the package group.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "pattern", pattern)
        if contact_info is not None:
            pulumi.set(__self__, "contact_info", contact_info)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_owner is not None:
            pulumi.set(__self__, "domain_owner", domain_owner)
        if origin_configuration is not None:
            pulumi.set(__self__, "origin_configuration", origin_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The name of the domain that contains the package group.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Input[str]:
        """
        The package group pattern that is used to gather packages.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> Optional[pulumi.Input[str]]:
        """
        The contact info of the package group.
        """
        return pulumi.get(self, "contact_info")

    @contact_info.setter
    def contact_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "contact_info", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The text description of the package group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainOwner")
    def domain_owner(self) -> Optional[pulumi.Input[str]]:
        """
        The 12-digit account ID of the AWS account that owns the domain.
        """
        return pulumi.get(self, "domain_owner")

    @domain_owner.setter
    def domain_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_owner", value)

    @property
    @pulumi.getter(name="originConfiguration")
    def origin_configuration(self) -> Optional[pulumi.Input['PackageGroupOriginConfigurationArgs']]:
        """
        The package origin configuration of the package group.
        """
        return pulumi.get(self, "origin_configuration")

    @origin_configuration.setter
    def origin_configuration(self, value: Optional[pulumi.Input['PackageGroupOriginConfigurationArgs']]):
        pulumi.set(self, "origin_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to the package group.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class PackageGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 origin_configuration: Optional[pulumi.Input[pulumi.InputType['PackageGroupOriginConfigurationArgs']]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        The resource schema to create a CodeArtifact package group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] contact_info: The contact info of the package group.
        :param pulumi.Input[str] description: The text description of the package group.
        :param pulumi.Input[str] domain_name: The name of the domain that contains the package group.
        :param pulumi.Input[str] domain_owner: The 12-digit account ID of the AWS account that owns the domain.
        :param pulumi.Input[pulumi.InputType['PackageGroupOriginConfigurationArgs']] origin_configuration: The package origin configuration of the package group.
        :param pulumi.Input[str] pattern: The package group pattern that is used to gather packages.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: An array of key-value pairs to apply to the package group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PackageGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The resource schema to create a CodeArtifact package group.

        :param str resource_name: The name of the resource.
        :param PackageGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PackageGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 contact_info: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 domain_owner: Optional[pulumi.Input[str]] = None,
                 origin_configuration: Optional[pulumi.Input[pulumi.InputType['PackageGroupOriginConfigurationArgs']]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PackageGroupArgs.__new__(PackageGroupArgs)

            __props__.__dict__["contact_info"] = contact_info
            __props__.__dict__["description"] = description
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["domain_owner"] = domain_owner
            __props__.__dict__["origin_configuration"] = origin_configuration
            if pattern is None and not opts.urn:
                raise TypeError("Missing required property 'pattern'")
            __props__.__dict__["pattern"] = pattern
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domainName", "pattern"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(PackageGroup, __self__).__init__(
            'aws-native:codeartifact:PackageGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PackageGroup':
        """
        Get an existing PackageGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PackageGroupArgs.__new__(PackageGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["contact_info"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["domain_name"] = None
        __props__.__dict__["domain_owner"] = None
        __props__.__dict__["origin_configuration"] = None
        __props__.__dict__["pattern"] = None
        __props__.__dict__["tags"] = None
        return PackageGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the package group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="contactInfo")
    def contact_info(self) -> pulumi.Output[Optional[str]]:
        """
        The contact info of the package group.
        """
        return pulumi.get(self, "contact_info")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The text description of the package group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The name of the domain that contains the package group.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="domainOwner")
    def domain_owner(self) -> pulumi.Output[Optional[str]]:
        """
        The 12-digit account ID of the AWS account that owns the domain.
        """
        return pulumi.get(self, "domain_owner")

    @property
    @pulumi.getter(name="originConfiguration")
    def origin_configuration(self) -> pulumi.Output[Optional['outputs.PackageGroupOriginConfiguration']]:
        """
        The package origin configuration of the package group.
        """
        return pulumi.get(self, "origin_configuration")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        """
        The package group pattern that is used to gather packages.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to the package group.
        """
        return pulumi.get(self, "tags")
