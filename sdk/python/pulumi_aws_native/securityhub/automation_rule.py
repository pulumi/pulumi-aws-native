# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AutomationRuleArgs', 'AutomationRule']

@pulumi.input_type
class AutomationRuleArgs:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRulesActionArgs']]]] = None,
                 criteria: Optional[pulumi.Input['AutomationRulesFindingFiltersArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_terminal: Optional[pulumi.Input[bool]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rule_order: Optional[pulumi.Input[int]] = None,
                 rule_status: Optional[pulumi.Input['AutomationRuleRuleStatus']] = None,
                 tags: Optional[pulumi.Input['AutomationRuleTagsArgs']] = None):
        """
        The set of arguments for constructing a AutomationRule resource.
        :param pulumi.Input['AutomationRulesFindingFiltersArgs'] criteria: The rule criteria for evaluating findings
        :param pulumi.Input[str] description: Rule description
        :param pulumi.Input[bool] is_terminal: If Rule is a terminal rule
        :param pulumi.Input[str] rule_name: Rule name
        :param pulumi.Input[int] rule_order: Rule order value
        :param pulumi.Input['AutomationRuleRuleStatus'] rule_status: Status of the Rule upon creation
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_terminal is not None:
            pulumi.set(__self__, "is_terminal", is_terminal)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if rule_order is not None:
            pulumi.set(__self__, "rule_order", rule_order)
        if rule_status is not None:
            pulumi.set(__self__, "rule_status", rule_status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRulesActionArgs']]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutomationRulesActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def criteria(self) -> Optional[pulumi.Input['AutomationRulesFindingFiltersArgs']]:
        """
        The rule criteria for evaluating findings
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input['AutomationRulesFindingFiltersArgs']]):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Rule description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isTerminal")
    def is_terminal(self) -> Optional[pulumi.Input[bool]]:
        """
        If Rule is a terminal rule
        """
        return pulumi.get(self, "is_terminal")

    @is_terminal.setter
    def is_terminal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_terminal", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        Rule name
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="ruleOrder")
    def rule_order(self) -> Optional[pulumi.Input[int]]:
        """
        Rule order value
        """
        return pulumi.get(self, "rule_order")

    @rule_order.setter
    def rule_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_order", value)

    @property
    @pulumi.getter(name="ruleStatus")
    def rule_status(self) -> Optional[pulumi.Input['AutomationRuleRuleStatus']]:
        """
        Status of the Rule upon creation
        """
        return pulumi.get(self, "rule_status")

    @rule_status.setter
    def rule_status(self, value: Optional[pulumi.Input['AutomationRuleRuleStatus']]):
        pulumi.set(self, "rule_status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['AutomationRuleTagsArgs']]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['AutomationRuleTagsArgs']]):
        pulumi.set(self, "tags", value)


class AutomationRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRulesActionArgs']]]]] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['AutomationRulesFindingFiltersArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_terminal: Optional[pulumi.Input[bool]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rule_order: Optional[pulumi.Input[int]] = None,
                 rule_status: Optional[pulumi.Input['AutomationRuleRuleStatus']] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['AutomationRuleTagsArgs']]] = None,
                 __props__=None):
        """
        The AWS::SecurityHub::AutomationRule resource represents the Automation Rule in your account. One rule resource is created for each Automation Rule in which you configure rule criteria and actions.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutomationRulesFindingFiltersArgs']] criteria: The rule criteria for evaluating findings
        :param pulumi.Input[str] description: Rule description
        :param pulumi.Input[bool] is_terminal: If Rule is a terminal rule
        :param pulumi.Input[str] rule_name: Rule name
        :param pulumi.Input[int] rule_order: Rule order value
        :param pulumi.Input['AutomationRuleRuleStatus'] rule_status: Status of the Rule upon creation
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AutomationRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::SecurityHub::AutomationRule resource represents the Automation Rule in your account. One rule resource is created for each Automation Rule in which you configure rule criteria and actions.

        :param str resource_name: The name of the resource.
        :param AutomationRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutomationRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutomationRulesActionArgs']]]]] = None,
                 criteria: Optional[pulumi.Input[pulumi.InputType['AutomationRulesFindingFiltersArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_terminal: Optional[pulumi.Input[bool]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rule_order: Optional[pulumi.Input[int]] = None,
                 rule_status: Optional[pulumi.Input['AutomationRuleRuleStatus']] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['AutomationRuleTagsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutomationRuleArgs.__new__(AutomationRuleArgs)

            __props__.__dict__["actions"] = actions
            __props__.__dict__["criteria"] = criteria
            __props__.__dict__["description"] = description
            __props__.__dict__["is_terminal"] = is_terminal
            __props__.__dict__["rule_name"] = rule_name
            __props__.__dict__["rule_order"] = rule_order
            __props__.__dict__["rule_status"] = rule_status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["created_by"] = None
            __props__.__dict__["rule_arn"] = None
            __props__.__dict__["updated_at"] = None
        super(AutomationRule, __self__).__init__(
            'aws-native:securityhub:AutomationRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AutomationRule':
        """
        Get an existing AutomationRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AutomationRuleArgs.__new__(AutomationRuleArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["created_by"] = None
        __props__.__dict__["criteria"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["is_terminal"] = None
        __props__.__dict__["rule_arn"] = None
        __props__.__dict__["rule_name"] = None
        __props__.__dict__["rule_order"] = None
        __props__.__dict__["rule_status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["updated_at"] = None
        return AutomationRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[Sequence['outputs.AutomationRulesAction']]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time when Automation Rule was created
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[str]:
        """
        The identifier by which created the rule
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output[Optional['outputs.AutomationRulesFindingFilters']]:
        """
        The rule criteria for evaluating findings
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Rule description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isTerminal")
    def is_terminal(self) -> pulumi.Output[Optional[bool]]:
        """
        If Rule is a terminal rule
        """
        return pulumi.get(self, "is_terminal")

    @property
    @pulumi.getter(name="ruleArn")
    def rule_arn(self) -> pulumi.Output[str]:
        """
        An Automation Rule Arn is automatically created
        """
        return pulumi.get(self, "rule_arn")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[Optional[str]]:
        """
        Rule name
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="ruleOrder")
    def rule_order(self) -> pulumi.Output[Optional[int]]:
        """
        Rule order value
        """
        return pulumi.get(self, "rule_order")

    @property
    @pulumi.getter(name="ruleStatus")
    def rule_status(self) -> pulumi.Output[Optional['AutomationRuleRuleStatus']]:
        """
        Status of the Rule upon creation
        """
        return pulumi.get(self, "rule_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.AutomationRuleTags']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time when Automation Rule was last updated
        """
        return pulumi.get(self, "updated_at")
