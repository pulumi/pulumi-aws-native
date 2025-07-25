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
from ._inputs import *

__all__ = ['LinkArgs', 'Link']

@pulumi.input_type
class LinkArgs:
    def __init__(__self__, *,
                 bandwidth: pulumi.Input['LinkBandwidthArgs'],
                 global_network_id: pulumi.Input[builtins.str],
                 site_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 provider: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Link resource.
        :param pulumi.Input['LinkBandwidthArgs'] bandwidth: The Bandwidth for the link.
        :param pulumi.Input[builtins.str] global_network_id: The ID of the global network.
        :param pulumi.Input[builtins.str] site_id: The ID of the site
        :param pulumi.Input[builtins.str] description: The description of the link.
        :param pulumi.Input[builtins.str] provider: The provider of the link.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The tags for the link.
        :param pulumi.Input[builtins.str] type: The type of the link.
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "global_network_id", global_network_id)
        pulumi.set(__self__, "site_id", site_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if provider is not None:
            pulumi.set(__self__, "provider", provider)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Input['LinkBandwidthArgs']:
        """
        The Bandwidth for the link.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: pulumi.Input['LinkBandwidthArgs']):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="globalNetworkId")
    def global_network_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the global network.
        """
        return pulumi.get(self, "global_network_id")

    @global_network_id.setter
    def global_network_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "global_network_id", value)

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the site
        """
        return pulumi.get(self, "site_id")

    @site_id.setter
    def site_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "site_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the link.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def provider(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The provider of the link.
        """
        return pulumi.get(self, "provider")

    @provider.setter
    def provider(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The tags for the link.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of the link.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("aws-native:networkmanager:Link")
class Link(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[Union['LinkBandwidthArgs', 'LinkBandwidthArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 global_network_id: Optional[pulumi.Input[builtins.str]] = None,
                 provider: Optional[pulumi.Input[builtins.str]] = None,
                 site_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The AWS::NetworkManager::Link type describes a link.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['LinkBandwidthArgs', 'LinkBandwidthArgsDict']] bandwidth: The Bandwidth for the link.
        :param pulumi.Input[builtins.str] description: The description of the link.
        :param pulumi.Input[builtins.str] global_network_id: The ID of the global network.
        :param pulumi.Input[builtins.str] provider: The provider of the link.
        :param pulumi.Input[builtins.str] site_id: The ID of the site
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: The tags for the link.
        :param pulumi.Input[builtins.str] type: The type of the link.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::NetworkManager::Link type describes a link.

        :param str resource_name: The name of the resource.
        :param LinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[Union['LinkBandwidthArgs', 'LinkBandwidthArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 global_network_id: Optional[pulumi.Input[builtins.str]] = None,
                 provider: Optional[pulumi.Input[builtins.str]] = None,
                 site_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LinkArgs.__new__(LinkArgs)

            if bandwidth is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth'")
            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["description"] = description
            if global_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'global_network_id'")
            __props__.__dict__["global_network_id"] = global_network_id
            __props__.__dict__["provider"] = provider
            if site_id is None and not opts.urn:
                raise TypeError("Missing required property 'site_id'")
            __props__.__dict__["site_id"] = site_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["link_arn"] = None
            __props__.__dict__["link_id"] = None
            __props__.__dict__["state"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["globalNetworkId", "siteId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Link, __self__).__init__(
            'aws-native:networkmanager:Link',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Link':
        """
        Get an existing Link resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LinkArgs.__new__(LinkArgs)

        __props__.__dict__["bandwidth"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["global_network_id"] = None
        __props__.__dict__["link_arn"] = None
        __props__.__dict__["link_id"] = None
        __props__.__dict__["provider"] = None
        __props__.__dict__["site_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Link(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output['outputs.LinkBandwidth']:
        """
        The Bandwidth for the link.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time that the device was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the link.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalNetworkId")
    def global_network_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the global network.
        """
        return pulumi.get(self, "global_network_id")

    @property
    @pulumi.getter(name="linkArn")
    def link_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the link.
        """
        return pulumi.get(self, "link_arn")

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the link.
        """
        return pulumi.get(self, "link_id")

    @property
    @pulumi.getter
    def provider(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The provider of the link.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter(name="siteId")
    def site_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the site
        """
        return pulumi.get(self, "site_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[builtins.str]:
        """
        The state of the link.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The tags for the link.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The type of the link.
        """
        return pulumi.get(self, "type")

