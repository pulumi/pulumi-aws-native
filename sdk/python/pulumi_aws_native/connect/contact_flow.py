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

__all__ = ['ContactFlowArgs', 'ContactFlow']

@pulumi.input_type
class ContactFlowArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 instance_arn: pulumi.Input[str],
                 type: pulumi.Input['ContactFlowType'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ContactFlowState']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ContactFlowTagArgs']]]] = None):
        """
        The set of arguments for constructing a ContactFlow resource.
        :param pulumi.Input[str] content: The content of the contact flow in JSON format.
        :param pulumi.Input[str] instance_arn: The identifier of the Amazon Connect instance (ARN).
        :param pulumi.Input['ContactFlowType'] type: The type of the contact flow.
        :param pulumi.Input[str] description: The description of the contact flow.
        :param pulumi.Input[str] name: The name of the contact flow.
        :param pulumi.Input['ContactFlowState'] state: The state of the contact flow.
        :param pulumi.Input[Sequence[pulumi.Input['ContactFlowTagArgs']]] tags: One or more tags.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "instance_arn", instance_arn)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The content of the contact flow in JSON format.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Input[str]:
        """
        The identifier of the Amazon Connect instance (ARN).
        """
        return pulumi.get(self, "instance_arn")

    @instance_arn.setter
    def instance_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_arn", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ContactFlowType']:
        """
        The type of the contact flow.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ContactFlowType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the contact flow.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the contact flow.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['ContactFlowState']]:
        """
        The state of the contact flow.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['ContactFlowState']]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContactFlowTagArgs']]]]:
        """
        One or more tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContactFlowTagArgs']]]]):
        pulumi.set(self, "tags", value)


class ContactFlow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ContactFlowState']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContactFlowTagArgs']]]]] = None,
                 type: Optional[pulumi.Input['ContactFlowType']] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Connect::ContactFlow

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: The content of the contact flow in JSON format.
        :param pulumi.Input[str] description: The description of the contact flow.
        :param pulumi.Input[str] instance_arn: The identifier of the Amazon Connect instance (ARN).
        :param pulumi.Input[str] name: The name of the contact flow.
        :param pulumi.Input['ContactFlowState'] state: The state of the contact flow.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContactFlowTagArgs']]]] tags: One or more tags.
        :param pulumi.Input['ContactFlowType'] type: The type of the contact flow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContactFlowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Connect::ContactFlow

        :param str resource_name: The name of the resource.
        :param ContactFlowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContactFlowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input['ContactFlowState']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContactFlowTagArgs']]]]] = None,
                 type: Optional[pulumi.Input['ContactFlowType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContactFlowArgs.__new__(ContactFlowArgs)

            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            __props__.__dict__["description"] = description
            if instance_arn is None and not opts.urn:
                raise TypeError("Missing required property 'instance_arn'")
            __props__.__dict__["instance_arn"] = instance_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["state"] = state
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["contact_flow_arn"] = None
        super(ContactFlow, __self__).__init__(
            'aws-native:connect:ContactFlow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ContactFlow':
        """
        Get an existing ContactFlow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ContactFlowArgs.__new__(ContactFlowArgs)

        __props__.__dict__["contact_flow_arn"] = None
        __props__.__dict__["content"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["instance_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return ContactFlow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contactFlowArn")
    def contact_flow_arn(self) -> pulumi.Output[str]:
        """
        The identifier of the contact flow (ARN).
        """
        return pulumi.get(self, "contact_flow_arn")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The content of the contact flow in JSON format.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the contact flow.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceArn")
    def instance_arn(self) -> pulumi.Output[str]:
        """
        The identifier of the Amazon Connect instance (ARN).
        """
        return pulumi.get(self, "instance_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the contact flow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional['ContactFlowState']]:
        """
        The state of the contact flow.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ContactFlowTag']]]:
        """
        One or more tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['ContactFlowType']:
        """
        The type of the contact flow.
        """
        return pulumi.get(self, "type")
