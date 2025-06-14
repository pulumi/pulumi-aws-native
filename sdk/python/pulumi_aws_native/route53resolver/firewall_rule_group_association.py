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

__all__ = ['FirewallRuleGroupAssociationArgs', 'FirewallRuleGroupAssociation']

@pulumi.input_type
class FirewallRuleGroupAssociationArgs:
    def __init__(__self__, *,
                 firewall_rule_group_id: pulumi.Input[builtins.str],
                 priority: pulumi.Input[builtins.int],
                 vpc_id: pulumi.Input[builtins.str],
                 mutation_protection: Optional[pulumi.Input['FirewallRuleGroupAssociationMutationProtection']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a FirewallRuleGroupAssociation resource.
        :param pulumi.Input[builtins.str] firewall_rule_group_id: FirewallRuleGroupId
        :param pulumi.Input[builtins.int] priority: Priority
        :param pulumi.Input[builtins.str] vpc_id: VpcId
        :param pulumi.Input['FirewallRuleGroupAssociationMutationProtection'] mutation_protection: MutationProtectionStatus
        :param pulumi.Input[builtins.str] name: FirewallRuleGroupAssociationName
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: Tags
        """
        pulumi.set(__self__, "firewall_rule_group_id", firewall_rule_group_id)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if mutation_protection is not None:
            pulumi.set(__self__, "mutation_protection", mutation_protection)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="firewallRuleGroupId")
    def firewall_rule_group_id(self) -> pulumi.Input[builtins.str]:
        """
        FirewallRuleGroupId
        """
        return pulumi.get(self, "firewall_rule_group_id")

    @firewall_rule_group_id.setter
    def firewall_rule_group_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "firewall_rule_group_id", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[builtins.int]:
        """
        Priority
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[builtins.str]:
        """
        VpcId
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="mutationProtection")
    def mutation_protection(self) -> Optional[pulumi.Input['FirewallRuleGroupAssociationMutationProtection']]:
        """
        MutationProtectionStatus
        """
        return pulumi.get(self, "mutation_protection")

    @mutation_protection.setter
    def mutation_protection(self, value: Optional[pulumi.Input['FirewallRuleGroupAssociationMutationProtection']]):
        pulumi.set(self, "mutation_protection", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        FirewallRuleGroupAssociationName
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        Tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:route53resolver:FirewallRuleGroupAssociation")
class FirewallRuleGroupAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_rule_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 mutation_protection: Optional[pulumi.Input['FirewallRuleGroupAssociationMutationProtection']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] firewall_rule_group_id: FirewallRuleGroupId
        :param pulumi.Input['FirewallRuleGroupAssociationMutationProtection'] mutation_protection: MutationProtectionStatus
        :param pulumi.Input[builtins.str] name: FirewallRuleGroupAssociationName
        :param pulumi.Input[builtins.int] priority: Priority
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: Tags
        :param pulumi.Input[builtins.str] vpc_id: VpcId
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallRuleGroupAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.

        :param str resource_name: The name of the resource.
        :param FirewallRuleGroupAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallRuleGroupAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 firewall_rule_group_id: Optional[pulumi.Input[builtins.str]] = None,
                 mutation_protection: Optional[pulumi.Input['FirewallRuleGroupAssociationMutationProtection']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallRuleGroupAssociationArgs.__new__(FirewallRuleGroupAssociationArgs)

            if firewall_rule_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_rule_group_id'")
            __props__.__dict__["firewall_rule_group_id"] = firewall_rule_group_id
            __props__.__dict__["mutation_protection"] = mutation_protection
            __props__.__dict__["name"] = name
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__.__dict__["priority"] = priority
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["creator_request_id"] = None
            __props__.__dict__["managed_owner_name"] = None
            __props__.__dict__["modification_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_message"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["firewallRuleGroupId", "vpcId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(FirewallRuleGroupAssociation, __self__).__init__(
            'aws-native:route53resolver:FirewallRuleGroupAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FirewallRuleGroupAssociation':
        """
        Get an existing FirewallRuleGroupAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FirewallRuleGroupAssociationArgs.__new__(FirewallRuleGroupAssociationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["creator_request_id"] = None
        __props__.__dict__["firewall_rule_group_id"] = None
        __props__.__dict__["managed_owner_name"] = None
        __props__.__dict__["modification_time"] = None
        __props__.__dict__["mutation_protection"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["priority"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_id"] = None
        return FirewallRuleGroupAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        Arn
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        Id
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[builtins.str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="creatorRequestId")
    def creator_request_id(self) -> pulumi.Output[builtins.str]:
        """
        The id of the creator request.
        """
        return pulumi.get(self, "creator_request_id")

    @property
    @pulumi.getter(name="firewallRuleGroupId")
    def firewall_rule_group_id(self) -> pulumi.Output[builtins.str]:
        """
        FirewallRuleGroupId
        """
        return pulumi.get(self, "firewall_rule_group_id")

    @property
    @pulumi.getter(name="managedOwnerName")
    def managed_owner_name(self) -> pulumi.Output[builtins.str]:
        """
        ServicePrincipal
        """
        return pulumi.get(self, "managed_owner_name")

    @property
    @pulumi.getter(name="modificationTime")
    def modification_time(self) -> pulumi.Output[builtins.str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "modification_time")

    @property
    @pulumi.getter(name="mutationProtection")
    def mutation_protection(self) -> pulumi.Output[Optional['FirewallRuleGroupAssociationMutationProtection']]:
        """
        MutationProtectionStatus
        """
        return pulumi.get(self, "mutation_protection")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        FirewallRuleGroupAssociationName
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[builtins.int]:
        """
        Priority
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['FirewallRuleGroupAssociationStatus']:
        """
        ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[builtins.str]:
        """
        FirewallDomainListAssociationStatus
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        Tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[builtins.str]:
        """
        VpcId
        """
        return pulumi.get(self, "vpc_id")

