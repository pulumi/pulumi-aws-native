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

__all__ = ['CampaignArgs', 'Campaign']

@pulumi.input_type
class CampaignArgs:
    def __init__(__self__, *,
                 connect_instance_arn: pulumi.Input[builtins.str],
                 dialer_config: pulumi.Input['CampaignDialerConfigArgs'],
                 outbound_call_config: pulumi.Input['CampaignOutboundCallConfigArgs'],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Campaign resource.
        :param pulumi.Input[builtins.str] connect_instance_arn: Amazon Connect Instance Arn
        :param pulumi.Input['CampaignDialerConfigArgs'] dialer_config: Contains information about the dialer configuration.
        :param pulumi.Input['CampaignOutboundCallConfigArgs'] outbound_call_config: Contains information about the outbound call configuration.
        :param pulumi.Input[builtins.str] name: Amazon Connect Campaign Name
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: One or more tags.
        """
        pulumi.set(__self__, "connect_instance_arn", connect_instance_arn)
        pulumi.set(__self__, "dialer_config", dialer_config)
        pulumi.set(__self__, "outbound_call_config", outbound_call_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="connectInstanceArn")
    def connect_instance_arn(self) -> pulumi.Input[builtins.str]:
        """
        Amazon Connect Instance Arn
        """
        return pulumi.get(self, "connect_instance_arn")

    @connect_instance_arn.setter
    def connect_instance_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "connect_instance_arn", value)

    @property
    @pulumi.getter(name="dialerConfig")
    def dialer_config(self) -> pulumi.Input['CampaignDialerConfigArgs']:
        """
        Contains information about the dialer configuration.
        """
        return pulumi.get(self, "dialer_config")

    @dialer_config.setter
    def dialer_config(self, value: pulumi.Input['CampaignDialerConfigArgs']):
        pulumi.set(self, "dialer_config", value)

    @property
    @pulumi.getter(name="outboundCallConfig")
    def outbound_call_config(self) -> pulumi.Input['CampaignOutboundCallConfigArgs']:
        """
        Contains information about the outbound call configuration.
        """
        return pulumi.get(self, "outbound_call_config")

    @outbound_call_config.setter
    def outbound_call_config(self, value: pulumi.Input['CampaignOutboundCallConfigArgs']):
        pulumi.set(self, "outbound_call_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Amazon Connect Campaign Name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        One or more tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:connectcampaigns:Campaign")
class Campaign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connect_instance_arn: Optional[pulumi.Input[builtins.str]] = None,
                 dialer_config: Optional[pulumi.Input[Union['CampaignDialerConfigArgs', 'CampaignDialerConfigArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 outbound_call_config: Optional[pulumi.Input[Union['CampaignOutboundCallConfigArgs', 'CampaignOutboundCallConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::ConnectCampaigns::Campaign Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] connect_instance_arn: Amazon Connect Instance Arn
        :param pulumi.Input[Union['CampaignDialerConfigArgs', 'CampaignDialerConfigArgsDict']] dialer_config: Contains information about the dialer configuration.
        :param pulumi.Input[builtins.str] name: Amazon Connect Campaign Name
        :param pulumi.Input[Union['CampaignOutboundCallConfigArgs', 'CampaignOutboundCallConfigArgsDict']] outbound_call_config: Contains information about the outbound call configuration.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: One or more tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CampaignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::ConnectCampaigns::Campaign Resource Type

        :param str resource_name: The name of the resource.
        :param CampaignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CampaignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connect_instance_arn: Optional[pulumi.Input[builtins.str]] = None,
                 dialer_config: Optional[pulumi.Input[Union['CampaignDialerConfigArgs', 'CampaignDialerConfigArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 outbound_call_config: Optional[pulumi.Input[Union['CampaignOutboundCallConfigArgs', 'CampaignOutboundCallConfigArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CampaignArgs.__new__(CampaignArgs)

            if connect_instance_arn is None and not opts.urn:
                raise TypeError("Missing required property 'connect_instance_arn'")
            __props__.__dict__["connect_instance_arn"] = connect_instance_arn
            if dialer_config is None and not opts.urn:
                raise TypeError("Missing required property 'dialer_config'")
            __props__.__dict__["dialer_config"] = dialer_config
            __props__.__dict__["name"] = name
            if outbound_call_config is None and not opts.urn:
                raise TypeError("Missing required property 'outbound_call_config'")
            __props__.__dict__["outbound_call_config"] = outbound_call_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["connectInstanceArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Campaign, __self__).__init__(
            'aws-native:connectcampaigns:Campaign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Campaign':
        """
        Get an existing Campaign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CampaignArgs.__new__(CampaignArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["connect_instance_arn"] = None
        __props__.__dict__["dialer_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["outbound_call_config"] = None
        __props__.__dict__["tags"] = None
        return Campaign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        Amazon Connect Campaign Arn
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="connectInstanceArn")
    def connect_instance_arn(self) -> pulumi.Output[builtins.str]:
        """
        Amazon Connect Instance Arn
        """
        return pulumi.get(self, "connect_instance_arn")

    @property
    @pulumi.getter(name="dialerConfig")
    def dialer_config(self) -> pulumi.Output['outputs.CampaignDialerConfig']:
        """
        Contains information about the dialer configuration.
        """
        return pulumi.get(self, "dialer_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Amazon Connect Campaign Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundCallConfig")
    def outbound_call_config(self) -> pulumi.Output['outputs.CampaignOutboundCallConfig']:
        """
        Contains information about the outbound call configuration.
        """
        return pulumi.get(self, "outbound_call_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        One or more tags.
        """
        return pulumi.get(self, "tags")

