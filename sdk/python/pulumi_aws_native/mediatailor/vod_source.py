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

__all__ = ['VodSourceArgs', 'VodSource']

@pulumi.input_type
class VodSourceArgs:
    def __init__(__self__, *,
                 http_package_configurations: pulumi.Input[Sequence[pulumi.Input['VodSourceHttpPackageConfigurationArgs']]],
                 source_location_name: pulumi.Input[builtins.str],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 vod_source_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a VodSource resource.
        :param pulumi.Input[Sequence[pulumi.Input['VodSourceHttpPackageConfigurationArgs']]] http_package_configurations: <p>A list of HTTP package configuration parameters for this VOD source.</p>
        :param pulumi.Input[builtins.str] source_location_name: The name of the source location that the VOD source is associated with.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The tags to assign to the VOD source.
        :param pulumi.Input[builtins.str] vod_source_name: The name of the VOD source.
        """
        pulumi.set(__self__, "http_package_configurations", http_package_configurations)
        pulumi.set(__self__, "source_location_name", source_location_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vod_source_name is not None:
            pulumi.set(__self__, "vod_source_name", vod_source_name)

    @property
    @pulumi.getter(name="httpPackageConfigurations")
    def http_package_configurations(self) -> pulumi.Input[Sequence[pulumi.Input['VodSourceHttpPackageConfigurationArgs']]]:
        """
        <p>A list of HTTP package configuration parameters for this VOD source.</p>
        """
        return pulumi.get(self, "http_package_configurations")

    @http_package_configurations.setter
    def http_package_configurations(self, value: pulumi.Input[Sequence[pulumi.Input['VodSourceHttpPackageConfigurationArgs']]]):
        pulumi.set(self, "http_package_configurations", value)

    @property
    @pulumi.getter(name="sourceLocationName")
    def source_location_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the source location that the VOD source is associated with.
        """
        return pulumi.get(self, "source_location_name")

    @source_location_name.setter
    def source_location_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "source_location_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The tags to assign to the VOD source.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vodSourceName")
    def vod_source_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the VOD source.
        """
        return pulumi.get(self, "vod_source_name")

    @vod_source_name.setter
    def vod_source_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vod_source_name", value)


@pulumi.type_token("aws-native:mediatailor:VodSource")
class VodSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_package_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VodSourceHttpPackageConfigurationArgs', 'VodSourceHttpPackageConfigurationArgsDict']]]]] = None,
                 source_location_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vod_source_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Definition of AWS::MediaTailor::VodSource Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VodSourceHttpPackageConfigurationArgs', 'VodSourceHttpPackageConfigurationArgsDict']]]] http_package_configurations: <p>A list of HTTP package configuration parameters for this VOD source.</p>
        :param pulumi.Input[builtins.str] source_location_name: The name of the source location that the VOD source is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: The tags to assign to the VOD source.
        :param pulumi.Input[builtins.str] vod_source_name: The name of the VOD source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VodSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::MediaTailor::VodSource Resource Type

        :param str resource_name: The name of the resource.
        :param VodSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VodSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_package_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VodSourceHttpPackageConfigurationArgs', 'VodSourceHttpPackageConfigurationArgsDict']]]]] = None,
                 source_location_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vod_source_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VodSourceArgs.__new__(VodSourceArgs)

            if http_package_configurations is None and not opts.urn:
                raise TypeError("Missing required property 'http_package_configurations'")
            __props__.__dict__["http_package_configurations"] = http_package_configurations
            if source_location_name is None and not opts.urn:
                raise TypeError("Missing required property 'source_location_name'")
            __props__.__dict__["source_location_name"] = source_location_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vod_source_name"] = vod_source_name
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["sourceLocationName", "vodSourceName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(VodSource, __self__).__init__(
            'aws-native:mediatailor:VodSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VodSource':
        """
        Get an existing VodSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VodSourceArgs.__new__(VodSourceArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["http_package_configurations"] = None
        __props__.__dict__["source_location_name"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vod_source_name"] = None
        return VodSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        <p>The ARN of the VOD source.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="httpPackageConfigurations")
    def http_package_configurations(self) -> pulumi.Output[Sequence['outputs.VodSourceHttpPackageConfiguration']]:
        """
        <p>A list of HTTP package configuration parameters for this VOD source.</p>
        """
        return pulumi.get(self, "http_package_configurations")

    @property
    @pulumi.getter(name="sourceLocationName")
    def source_location_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the source location that the VOD source is associated with.
        """
        return pulumi.get(self, "source_location_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The tags to assign to the VOD source.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vodSourceName")
    def vod_source_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the VOD source.
        """
        return pulumi.get(self, "vod_source_name")

