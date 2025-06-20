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

__all__ = ['IntelligentPromptRouterArgs', 'IntelligentPromptRouter']

@pulumi.input_type
class IntelligentPromptRouterArgs:
    def __init__(__self__, *,
                 fallback_model: pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs'],
                 models: pulumi.Input[Sequence[pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']]],
                 routing_criteria: pulumi.Input['IntelligentPromptRouterRoutingCriteriaArgs'],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 prompt_router_name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a IntelligentPromptRouter resource.
        :param pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs'] fallback_model: The default model to use when the routing criteria is not met.
        :param pulumi.Input[Sequence[pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']]] models: List of model configuration
        :param pulumi.Input['IntelligentPromptRouterRoutingCriteriaArgs'] routing_criteria: Routing criteria for a prompt router.
        :param pulumi.Input[builtins.str] description: Description of the Prompt Router.
        :param pulumi.Input[builtins.str] prompt_router_name: Name of the Prompt Router.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: List of Tags
        """
        pulumi.set(__self__, "fallback_model", fallback_model)
        pulumi.set(__self__, "models", models)
        pulumi.set(__self__, "routing_criteria", routing_criteria)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if prompt_router_name is not None:
            pulumi.set(__self__, "prompt_router_name", prompt_router_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="fallbackModel")
    def fallback_model(self) -> pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']:
        """
        The default model to use when the routing criteria is not met.
        """
        return pulumi.get(self, "fallback_model")

    @fallback_model.setter
    def fallback_model(self, value: pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']):
        pulumi.set(self, "fallback_model", value)

    @property
    @pulumi.getter
    def models(self) -> pulumi.Input[Sequence[pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']]]:
        """
        List of model configuration
        """
        return pulumi.get(self, "models")

    @models.setter
    def models(self, value: pulumi.Input[Sequence[pulumi.Input['IntelligentPromptRouterPromptRouterTargetModelArgs']]]):
        pulumi.set(self, "models", value)

    @property
    @pulumi.getter(name="routingCriteria")
    def routing_criteria(self) -> pulumi.Input['IntelligentPromptRouterRoutingCriteriaArgs']:
        """
        Routing criteria for a prompt router.
        """
        return pulumi.get(self, "routing_criteria")

    @routing_criteria.setter
    def routing_criteria(self, value: pulumi.Input['IntelligentPromptRouterRoutingCriteriaArgs']):
        pulumi.set(self, "routing_criteria", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the Prompt Router.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="promptRouterName")
    def prompt_router_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Prompt Router.
        """
        return pulumi.get(self, "prompt_router_name")

    @prompt_router_name.setter
    def prompt_router_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "prompt_router_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        List of Tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:bedrock:IntelligentPromptRouter")
class IntelligentPromptRouter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 fallback_model: Optional[pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']]] = None,
                 models: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']]]]] = None,
                 prompt_router_name: Optional[pulumi.Input[builtins.str]] = None,
                 routing_criteria: Optional[pulumi.Input[Union['IntelligentPromptRouterRoutingCriteriaArgs', 'IntelligentPromptRouterRoutingCriteriaArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::Bedrock::IntelligentPromptRouter Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the Prompt Router.
        :param pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']] fallback_model: The default model to use when the routing criteria is not met.
        :param pulumi.Input[Sequence[pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']]]] models: List of model configuration
        :param pulumi.Input[builtins.str] prompt_router_name: Name of the Prompt Router.
        :param pulumi.Input[Union['IntelligentPromptRouterRoutingCriteriaArgs', 'IntelligentPromptRouterRoutingCriteriaArgsDict']] routing_criteria: Routing criteria for a prompt router.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: List of Tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntelligentPromptRouterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Bedrock::IntelligentPromptRouter Resource Type

        :param str resource_name: The name of the resource.
        :param IntelligentPromptRouterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntelligentPromptRouterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 fallback_model: Optional[pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']]] = None,
                 models: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IntelligentPromptRouterPromptRouterTargetModelArgs', 'IntelligentPromptRouterPromptRouterTargetModelArgsDict']]]]] = None,
                 prompt_router_name: Optional[pulumi.Input[builtins.str]] = None,
                 routing_criteria: Optional[pulumi.Input[Union['IntelligentPromptRouterRoutingCriteriaArgs', 'IntelligentPromptRouterRoutingCriteriaArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntelligentPromptRouterArgs.__new__(IntelligentPromptRouterArgs)

            __props__.__dict__["description"] = description
            if fallback_model is None and not opts.urn:
                raise TypeError("Missing required property 'fallback_model'")
            __props__.__dict__["fallback_model"] = fallback_model
            if models is None and not opts.urn:
                raise TypeError("Missing required property 'models'")
            __props__.__dict__["models"] = models
            __props__.__dict__["prompt_router_name"] = prompt_router_name
            if routing_criteria is None and not opts.urn:
                raise TypeError("Missing required property 'routing_criteria'")
            __props__.__dict__["routing_criteria"] = routing_criteria
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["prompt_router_arn"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "fallbackModel", "models[*]", "promptRouterName", "routingCriteria"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(IntelligentPromptRouter, __self__).__init__(
            'aws-native:bedrock:IntelligentPromptRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IntelligentPromptRouter':
        """
        Get an existing IntelligentPromptRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IntelligentPromptRouterArgs.__new__(IntelligentPromptRouterArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["fallback_model"] = None
        __props__.__dict__["models"] = None
        __props__.__dict__["prompt_router_arn"] = None
        __props__.__dict__["prompt_router_name"] = None
        __props__.__dict__["routing_criteria"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return IntelligentPromptRouter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the Prompt Router.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fallbackModel")
    def fallback_model(self) -> pulumi.Output['outputs.IntelligentPromptRouterPromptRouterTargetModel']:
        """
        The default model to use when the routing criteria is not met.
        """
        return pulumi.get(self, "fallback_model")

    @property
    @pulumi.getter
    def models(self) -> pulumi.Output[Sequence['outputs.IntelligentPromptRouterPromptRouterTargetModel']]:
        """
        List of model configuration
        """
        return pulumi.get(self, "models")

    @property
    @pulumi.getter(name="promptRouterArn")
    def prompt_router_arn(self) -> pulumi.Output[builtins.str]:
        """
        Arn of the Prompt Router.
        """
        return pulumi.get(self, "prompt_router_arn")

    @property
    @pulumi.getter(name="promptRouterName")
    def prompt_router_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the Prompt Router.
        """
        return pulumi.get(self, "prompt_router_name")

    @property
    @pulumi.getter(name="routingCriteria")
    def routing_criteria(self) -> pulumi.Output['outputs.IntelligentPromptRouterRoutingCriteria']:
        """
        Routing criteria for a prompt router.
        """
        return pulumi.get(self, "routing_criteria")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['IntelligentPromptRouterPromptRouterStatus']:
        """
        The router's status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        List of Tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['IntelligentPromptRouterPromptRouterType']:
        """
        The router's type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Time Stamp
        """
        return pulumi.get(self, "updated_at")

