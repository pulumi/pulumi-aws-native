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

__all__ = ['ReferenceStoreArgs', 'ReferenceStore']

@pulumi.input_type
class ReferenceStoreArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sse_config: Optional[pulumi.Input['ReferenceStoreSseConfigArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ReferenceStore resource.
        :param pulumi.Input[builtins.str] description: A description for the store.
        :param pulumi.Input[builtins.str] name: A name for the store.
        :param pulumi.Input['ReferenceStoreSseConfigArgs'] sse_config: Server-side encryption (SSE) settings for the store.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags for the store.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sse_config is not None:
            pulumi.set(__self__, "sse_config", sse_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description for the store.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name for the store.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sseConfig")
    def sse_config(self) -> Optional[pulumi.Input['ReferenceStoreSseConfigArgs']]:
        """
        Server-side encryption (SSE) settings for the store.
        """
        return pulumi.get(self, "sse_config")

    @sse_config.setter
    def sse_config(self, value: Optional[pulumi.Input['ReferenceStoreSseConfigArgs']]):
        pulumi.set(self, "sse_config", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Tags for the store.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:omics:ReferenceStore")
class ReferenceStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sse_config: Optional[pulumi.Input[Union['ReferenceStoreSseConfigArgs', 'ReferenceStoreSseConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Definition of AWS::Omics::ReferenceStore Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A description for the store.
        :param pulumi.Input[builtins.str] name: A name for the store.
        :param pulumi.Input[Union['ReferenceStoreSseConfigArgs', 'ReferenceStoreSseConfigArgsDict']] sse_config: Server-side encryption (SSE) settings for the store.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Tags for the store.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ReferenceStoreArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Omics::ReferenceStore Resource Type

        :param str resource_name: The name of the resource.
        :param ReferenceStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReferenceStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 sse_config: Optional[pulumi.Input[Union['ReferenceStoreSseConfigArgs', 'ReferenceStoreSseConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReferenceStoreArgs.__new__(ReferenceStoreArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["sse_config"] = sse_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["reference_store_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "name", "sseConfig", "tags.*"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ReferenceStore, __self__).__init__(
            'aws-native:omics:ReferenceStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ReferenceStore':
        """
        Get an existing ReferenceStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReferenceStoreArgs.__new__(ReferenceStoreArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["reference_store_id"] = None
        __props__.__dict__["sse_config"] = None
        __props__.__dict__["tags"] = None
        return ReferenceStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The store's ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[builtins.str]:
        """
        When the store was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description for the store.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A name for the store.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="referenceStoreId")
    def reference_store_id(self) -> pulumi.Output[builtins.str]:
        """
        The store's ID.
        """
        return pulumi.get(self, "reference_store_id")

    @property
    @pulumi.getter(name="sseConfig")
    def sse_config(self) -> pulumi.Output[Optional['outputs.ReferenceStoreSseConfig']]:
        """
        Server-side encryption (SSE) settings for the store.
        """
        return pulumi.get(self, "sse_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Tags for the store.
        """
        return pulumi.get(self, "tags")

