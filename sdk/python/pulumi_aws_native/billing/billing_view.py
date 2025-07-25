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

__all__ = ['BillingViewArgs', 'BillingView']

@pulumi.input_type
class BillingViewArgs:
    def __init__(__self__, *,
                 source_views: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 data_filter_expression: Optional[pulumi.Input['DataFilterExpressionPropertiesArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a BillingView resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] source_views: An array of strings that define the billing view's source.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs associated to the billing view being created.
        """
        pulumi.set(__self__, "source_views", source_views)
        if data_filter_expression is not None:
            pulumi.set(__self__, "data_filter_expression", data_filter_expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="sourceViews")
    def source_views(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        An array of strings that define the billing view's source.
        """
        return pulumi.get(self, "source_views")

    @source_views.setter
    def source_views(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "source_views", value)

    @property
    @pulumi.getter(name="dataFilterExpression")
    def data_filter_expression(self) -> Optional[pulumi.Input['DataFilterExpressionPropertiesArgs']]:
        return pulumi.get(self, "data_filter_expression")

    @data_filter_expression.setter
    def data_filter_expression(self, value: Optional[pulumi.Input['DataFilterExpressionPropertiesArgs']]):
        pulumi.set(self, "data_filter_expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs associated to the billing view being created.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:billing:BillingView")
class BillingView(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_filter_expression: Optional[pulumi.Input[Union['DataFilterExpressionPropertiesArgs', 'DataFilterExpressionPropertiesArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 source_views: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        A billing view is a container of cost & usage metadata.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] source_views: An array of strings that define the billing view's source.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: An array of key-value pairs associated to the billing view being created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BillingViewArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A billing view is a container of cost & usage metadata.

        :param str resource_name: The name of the resource.
        :param BillingViewArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BillingViewArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_filter_expression: Optional[pulumi.Input[Union['DataFilterExpressionPropertiesArgs', 'DataFilterExpressionPropertiesArgsDict']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 source_views: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BillingViewArgs.__new__(BillingViewArgs)

            __props__.__dict__["data_filter_expression"] = data_filter_expression
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if source_views is None and not opts.urn:
                raise TypeError("Missing required property 'source_views'")
            __props__.__dict__["source_views"] = source_views
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["billing_view_type"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["owner_account_id"] = None
            __props__.__dict__["updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "name", "sourceViews[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(BillingView, __self__).__init__(
            'aws-native:billing:BillingView',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BillingView':
        """
        Get an existing BillingView resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BillingViewArgs.__new__(BillingViewArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["billing_view_type"] = None
        __props__.__dict__["created_at"] = None
        __props__.__dict__["data_filter_expression"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner_account_id"] = None
        __props__.__dict__["source_views"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["updated_at"] = None
        return BillingView(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="billingViewType")
    def billing_view_type(self) -> pulumi.Output['BillingViewType']:
        return pulumi.get(self, "billing_view_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.float]:
        """
        The time when the billing view was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dataFilterExpression")
    def data_filter_expression(self) -> pulumi.Output[Optional['outputs.DataFilterExpressionProperties']]:
        return pulumi.get(self, "data_filter_expression")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerAccountId")
    def owner_account_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "owner_account_id")

    @property
    @pulumi.getter(name="sourceViews")
    def source_views(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        An array of strings that define the billing view's source.
        """
        return pulumi.get(self, "source_views")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs associated to the billing view being created.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.float]:
        """
        The time when the billing view was last updated.
        """
        return pulumi.get(self, "updated_at")

