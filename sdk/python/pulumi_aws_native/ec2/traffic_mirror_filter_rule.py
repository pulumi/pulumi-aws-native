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

__all__ = ['TrafficMirrorFilterRuleArgs', 'TrafficMirrorFilterRule']

@pulumi.input_type
class TrafficMirrorFilterRuleArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[builtins.str],
                 rule_action: pulumi.Input[builtins.str],
                 rule_number: pulumi.Input[builtins.int],
                 source_cidr_block: pulumi.Input[builtins.str],
                 traffic_direction: pulumi.Input[builtins.str],
                 traffic_mirror_filter_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 destination_port_range: Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']] = None,
                 protocol: Optional[pulumi.Input[builtins.int]] = None,
                 source_port_range: Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a TrafficMirrorFilterRule resource.
        :param pulumi.Input[builtins.str] destination_cidr_block: The destination CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[builtins.str] rule_action: The action to take on the filtered traffic (accept/reject).
        :param pulumi.Input[builtins.int] rule_number: The number of the Traffic Mirror rule.
        :param pulumi.Input[builtins.str] source_cidr_block: The source CIDR block to assign to the Traffic Mirror Filter rule.
        :param pulumi.Input[builtins.str] traffic_direction: The direction of traffic (ingress/egress).
        :param pulumi.Input[builtins.str] traffic_mirror_filter_id: The ID of the filter that this rule is associated with.
        :param pulumi.Input[builtins.str] description: The description of the Traffic Mirror Filter rule.
        :param pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs'] destination_port_range: The destination port range.
        :param pulumi.Input[builtins.int] protocol: The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.
        :param pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs'] source_port_range: The source port range.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: Any tags assigned to the Traffic Mirror Filter rule.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "rule_action", rule_action)
        pulumi.set(__self__, "rule_number", rule_number)
        pulumi.set(__self__, "source_cidr_block", source_cidr_block)
        pulumi.set(__self__, "traffic_direction", traffic_direction)
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_port_range is not None:
            pulumi.set(__self__, "destination_port_range", destination_port_range)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if source_port_range is not None:
            pulumi.set(__self__, "source_port_range", source_port_range)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[builtins.str]:
        """
        The destination CIDR block to assign to the Traffic Mirror rule.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Input[builtins.str]:
        """
        The action to take on the filtered traffic (accept/reject).
        """
        return pulumi.get(self, "rule_action")

    @rule_action.setter
    def rule_action(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "rule_action", value)

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> pulumi.Input[builtins.int]:
        """
        The number of the Traffic Mirror rule.
        """
        return pulumi.get(self, "rule_number")

    @rule_number.setter
    def rule_number(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "rule_number", value)

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> pulumi.Input[builtins.str]:
        """
        The source CIDR block to assign to the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "source_cidr_block")

    @source_cidr_block.setter
    def source_cidr_block(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "source_cidr_block", value)

    @property
    @pulumi.getter(name="trafficDirection")
    def traffic_direction(self) -> pulumi.Input[builtins.str]:
        """
        The direction of traffic (ingress/egress).
        """
        return pulumi.get(self, "traffic_direction")

    @traffic_direction.setter
    def traffic_direction(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "traffic_direction", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the filter that this rule is associated with.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']]:
        """
        The destination port range.
        """
        return pulumi.get(self, "destination_port_range")

    @destination_port_range.setter
    def destination_port_range(self, value: Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']]):
        pulumi.set(self, "destination_port_range", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']]:
        """
        The source port range.
        """
        return pulumi.get(self, "source_port_range")

    @source_port_range.setter
    def source_port_range(self, value: Optional[pulumi.Input['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs']]):
        pulumi.set(self, "source_port_range", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        Any tags assigned to the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:ec2:TrafficMirrorFilterRule")
class TrafficMirrorFilterRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[builtins.str]] = None,
                 destination_port_range: Optional[pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']]] = None,
                 protocol: Optional[pulumi.Input[builtins.int]] = None,
                 rule_action: Optional[pulumi.Input[builtins.str]] = None,
                 rule_number: Optional[pulumi.Input[builtins.int]] = None,
                 source_cidr_block: Optional[pulumi.Input[builtins.str]] = None,
                 source_port_range: Optional[pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 traffic_direction: Optional[pulumi.Input[builtins.str]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource Type definition for for AWS::EC2::TrafficMirrorFilterRule

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The description of the Traffic Mirror Filter rule.
        :param pulumi.Input[builtins.str] destination_cidr_block: The destination CIDR block to assign to the Traffic Mirror rule.
        :param pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']] destination_port_range: The destination port range.
        :param pulumi.Input[builtins.int] protocol: The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.
        :param pulumi.Input[builtins.str] rule_action: The action to take on the filtered traffic (accept/reject).
        :param pulumi.Input[builtins.int] rule_number: The number of the Traffic Mirror rule.
        :param pulumi.Input[builtins.str] source_cidr_block: The source CIDR block to assign to the Traffic Mirror Filter rule.
        :param pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']] source_port_range: The source port range.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: Any tags assigned to the Traffic Mirror Filter rule.
        :param pulumi.Input[builtins.str] traffic_direction: The direction of traffic (ingress/egress).
        :param pulumi.Input[builtins.str] traffic_mirror_filter_id: The ID of the filter that this rule is associated with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficMirrorFilterRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for for AWS::EC2::TrafficMirrorFilterRule

        :param str resource_name: The name of the resource.
        :param TrafficMirrorFilterRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorFilterRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[builtins.str]] = None,
                 destination_port_range: Optional[pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']]] = None,
                 protocol: Optional[pulumi.Input[builtins.int]] = None,
                 rule_action: Optional[pulumi.Input[builtins.str]] = None,
                 rule_number: Optional[pulumi.Input[builtins.int]] = None,
                 source_cidr_block: Optional[pulumi.Input[builtins.str]] = None,
                 source_port_range: Optional[pulumi.Input[Union['TrafficMirrorFilterRuleTrafficMirrorPortRangeArgs', 'TrafficMirrorFilterRuleTrafficMirrorPortRangeArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 traffic_direction: Optional[pulumi.Input[builtins.str]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficMirrorFilterRuleArgs.__new__(TrafficMirrorFilterRuleArgs)

            __props__.__dict__["description"] = description
            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["destination_port_range"] = destination_port_range
            __props__.__dict__["protocol"] = protocol
            if rule_action is None and not opts.urn:
                raise TypeError("Missing required property 'rule_action'")
            __props__.__dict__["rule_action"] = rule_action
            if rule_number is None and not opts.urn:
                raise TypeError("Missing required property 'rule_number'")
            __props__.__dict__["rule_number"] = rule_number
            if source_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'source_cidr_block'")
            __props__.__dict__["source_cidr_block"] = source_cidr_block
            __props__.__dict__["source_port_range"] = source_port_range
            __props__.__dict__["tags"] = tags
            if traffic_direction is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_direction'")
            __props__.__dict__["traffic_direction"] = traffic_direction
            if traffic_mirror_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
            __props__.__dict__["traffic_mirror_filter_rule_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["trafficMirrorFilterId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(TrafficMirrorFilterRule, __self__).__init__(
            'aws-native:ec2:TrafficMirrorFilterRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TrafficMirrorFilterRule':
        """
        Get an existing TrafficMirrorFilterRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrafficMirrorFilterRuleArgs.__new__(TrafficMirrorFilterRuleArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["destination_cidr_block"] = None
        __props__.__dict__["destination_port_range"] = None
        __props__.__dict__["protocol"] = None
        __props__.__dict__["rule_action"] = None
        __props__.__dict__["rule_number"] = None
        __props__.__dict__["source_cidr_block"] = None
        __props__.__dict__["source_port_range"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["traffic_direction"] = None
        __props__.__dict__["traffic_mirror_filter_id"] = None
        __props__.__dict__["traffic_mirror_filter_rule_id"] = None
        return TrafficMirrorFilterRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[builtins.str]:
        """
        The destination CIDR block to assign to the Traffic Mirror rule.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="destinationPortRange")
    def destination_port_range(self) -> pulumi.Output[Optional['outputs.TrafficMirrorFilterRuleTrafficMirrorPortRange']]:
        """
        The destination port range.
        """
        return pulumi.get(self, "destination_port_range")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="ruleAction")
    def rule_action(self) -> pulumi.Output[builtins.str]:
        """
        The action to take on the filtered traffic (accept/reject).
        """
        return pulumi.get(self, "rule_action")

    @property
    @pulumi.getter(name="ruleNumber")
    def rule_number(self) -> pulumi.Output[builtins.int]:
        """
        The number of the Traffic Mirror rule.
        """
        return pulumi.get(self, "rule_number")

    @property
    @pulumi.getter(name="sourceCidrBlock")
    def source_cidr_block(self) -> pulumi.Output[builtins.str]:
        """
        The source CIDR block to assign to the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "source_cidr_block")

    @property
    @pulumi.getter(name="sourcePortRange")
    def source_port_range(self) -> pulumi.Output[Optional['outputs.TrafficMirrorFilterRuleTrafficMirrorPortRange']]:
        """
        The source port range.
        """
        return pulumi.get(self, "source_port_range")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        Any tags assigned to the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficDirection")
    def traffic_direction(self) -> pulumi.Output[builtins.str]:
        """
        The direction of traffic (ingress/egress).
        """
        return pulumi.get(self, "traffic_direction")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the filter that this rule is associated with.
        """
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorFilterRuleId")
    def traffic_mirror_filter_rule_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Traffic Mirror Filter rule.
        """
        return pulumi.get(self, "traffic_mirror_filter_rule_id")

