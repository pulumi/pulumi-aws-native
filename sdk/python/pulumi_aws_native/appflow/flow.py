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

__all__ = ['FlowArgs', 'Flow']

@pulumi.input_type
class FlowArgs:
    def __init__(__self__, *,
                 destination_flow_config_list: pulumi.Input[Sequence[pulumi.Input['FlowDestinationFlowConfigArgs']]],
                 source_flow_config: pulumi.Input['FlowSourceFlowConfigArgs'],
                 tasks: pulumi.Input[Sequence[pulumi.Input['FlowTaskArgs']]],
                 trigger_config: pulumi.Input['FlowTriggerConfigArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 k_ms_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FlowTagArgs']]]] = None):
        """
        The set of arguments for constructing a Flow resource.
        :param pulumi.Input[Sequence[pulumi.Input['FlowDestinationFlowConfigArgs']]] destination_flow_config_list: List of Destination connectors of the flow.
        :param pulumi.Input['FlowSourceFlowConfigArgs'] source_flow_config: Configurations of Source connector of the flow.
        :param pulumi.Input[Sequence[pulumi.Input['FlowTaskArgs']]] tasks: List of tasks for the flow.
        :param pulumi.Input['FlowTriggerConfigArgs'] trigger_config: Trigger settings of the flow.
        :param pulumi.Input[str] description: Description of the flow.
        :param pulumi.Input[str] flow_name: Name of the flow.
        :param pulumi.Input[str] k_ms_arn: The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
        :param pulumi.Input[Sequence[pulumi.Input['FlowTagArgs']]] tags: List of Tags.
        """
        pulumi.set(__self__, "destination_flow_config_list", destination_flow_config_list)
        pulumi.set(__self__, "source_flow_config", source_flow_config)
        pulumi.set(__self__, "tasks", tasks)
        pulumi.set(__self__, "trigger_config", trigger_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if flow_name is not None:
            pulumi.set(__self__, "flow_name", flow_name)
        if k_ms_arn is not None:
            pulumi.set(__self__, "k_ms_arn", k_ms_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="destinationFlowConfigList")
    def destination_flow_config_list(self) -> pulumi.Input[Sequence[pulumi.Input['FlowDestinationFlowConfigArgs']]]:
        """
        List of Destination connectors of the flow.
        """
        return pulumi.get(self, "destination_flow_config_list")

    @destination_flow_config_list.setter
    def destination_flow_config_list(self, value: pulumi.Input[Sequence[pulumi.Input['FlowDestinationFlowConfigArgs']]]):
        pulumi.set(self, "destination_flow_config_list", value)

    @property
    @pulumi.getter(name="sourceFlowConfig")
    def source_flow_config(self) -> pulumi.Input['FlowSourceFlowConfigArgs']:
        """
        Configurations of Source connector of the flow.
        """
        return pulumi.get(self, "source_flow_config")

    @source_flow_config.setter
    def source_flow_config(self, value: pulumi.Input['FlowSourceFlowConfigArgs']):
        pulumi.set(self, "source_flow_config", value)

    @property
    @pulumi.getter
    def tasks(self) -> pulumi.Input[Sequence[pulumi.Input['FlowTaskArgs']]]:
        """
        List of tasks for the flow.
        """
        return pulumi.get(self, "tasks")

    @tasks.setter
    def tasks(self, value: pulumi.Input[Sequence[pulumi.Input['FlowTaskArgs']]]):
        pulumi.set(self, "tasks", value)

    @property
    @pulumi.getter(name="triggerConfig")
    def trigger_config(self) -> pulumi.Input['FlowTriggerConfigArgs']:
        """
        Trigger settings of the flow.
        """
        return pulumi.get(self, "trigger_config")

    @trigger_config.setter
    def trigger_config(self, value: pulumi.Input['FlowTriggerConfigArgs']):
        pulumi.set(self, "trigger_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the flow.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="flowName")
    def flow_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the flow.
        """
        return pulumi.get(self, "flow_name")

    @flow_name.setter
    def flow_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_name", value)

    @property
    @pulumi.getter(name="kMSArn")
    def k_ms_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
        """
        return pulumi.get(self, "k_ms_arn")

    @k_ms_arn.setter
    def k_ms_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "k_ms_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FlowTagArgs']]]]:
        """
        List of Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FlowTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Flow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_flow_config_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowDestinationFlowConfigArgs']]]]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 k_ms_arn: Optional[pulumi.Input[str]] = None,
                 source_flow_config: Optional[pulumi.Input[pulumi.InputType['FlowSourceFlowConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTagArgs']]]]] = None,
                 tasks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTaskArgs']]]]] = None,
                 trigger_config: Optional[pulumi.Input[pulumi.InputType['FlowTriggerConfigArgs']]] = None,
                 __props__=None):
        """
        Resource schema for AWS::AppFlow::Flow.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the flow.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowDestinationFlowConfigArgs']]]] destination_flow_config_list: List of Destination connectors of the flow.
        :param pulumi.Input[str] flow_name: Name of the flow.
        :param pulumi.Input[str] k_ms_arn: The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
        :param pulumi.Input[pulumi.InputType['FlowSourceFlowConfigArgs']] source_flow_config: Configurations of Source connector of the flow.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTagArgs']]]] tags: List of Tags.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTaskArgs']]]] tasks: List of tasks for the flow.
        :param pulumi.Input[pulumi.InputType['FlowTriggerConfigArgs']] trigger_config: Trigger settings of the flow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::AppFlow::Flow.

        :param str resource_name: The name of the resource.
        :param FlowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_flow_config_list: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowDestinationFlowConfigArgs']]]]] = None,
                 flow_name: Optional[pulumi.Input[str]] = None,
                 k_ms_arn: Optional[pulumi.Input[str]] = None,
                 source_flow_config: Optional[pulumi.Input[pulumi.InputType['FlowSourceFlowConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTagArgs']]]]] = None,
                 tasks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FlowTaskArgs']]]]] = None,
                 trigger_config: Optional[pulumi.Input[pulumi.InputType['FlowTriggerConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlowArgs.__new__(FlowArgs)

            __props__.__dict__["description"] = description
            if destination_flow_config_list is None and not opts.urn:
                raise TypeError("Missing required property 'destination_flow_config_list'")
            __props__.__dict__["destination_flow_config_list"] = destination_flow_config_list
            __props__.__dict__["flow_name"] = flow_name
            __props__.__dict__["k_ms_arn"] = k_ms_arn
            if source_flow_config is None and not opts.urn:
                raise TypeError("Missing required property 'source_flow_config'")
            __props__.__dict__["source_flow_config"] = source_flow_config
            __props__.__dict__["tags"] = tags
            if tasks is None and not opts.urn:
                raise TypeError("Missing required property 'tasks'")
            __props__.__dict__["tasks"] = tasks
            if trigger_config is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_config'")
            __props__.__dict__["trigger_config"] = trigger_config
            __props__.__dict__["flow_arn"] = None
        super(Flow, __self__).__init__(
            'aws-native:appflow:Flow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Flow':
        """
        Get an existing Flow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FlowArgs.__new__(FlowArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["destination_flow_config_list"] = None
        __props__.__dict__["flow_arn"] = None
        __props__.__dict__["flow_name"] = None
        __props__.__dict__["k_ms_arn"] = None
        __props__.__dict__["source_flow_config"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tasks"] = None
        __props__.__dict__["trigger_config"] = None
        return Flow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the flow.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationFlowConfigList")
    def destination_flow_config_list(self) -> pulumi.Output[Sequence['outputs.FlowDestinationFlowConfig']]:
        """
        List of Destination connectors of the flow.
        """
        return pulumi.get(self, "destination_flow_config_list")

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Output[str]:
        """
        ARN identifier of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter(name="flowName")
    def flow_name(self) -> pulumi.Output[str]:
        """
        Name of the flow.
        """
        return pulumi.get(self, "flow_name")

    @property
    @pulumi.getter(name="kMSArn")
    def k_ms_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
        """
        return pulumi.get(self, "k_ms_arn")

    @property
    @pulumi.getter(name="sourceFlowConfig")
    def source_flow_config(self) -> pulumi.Output['outputs.FlowSourceFlowConfig']:
        """
        Configurations of Source connector of the flow.
        """
        return pulumi.get(self, "source_flow_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FlowTag']]]:
        """
        List of Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tasks(self) -> pulumi.Output[Sequence['outputs.FlowTask']]:
        """
        List of tasks for the flow.
        """
        return pulumi.get(self, "tasks")

    @property
    @pulumi.getter(name="triggerConfig")
    def trigger_config(self) -> pulumi.Output['outputs.FlowTriggerConfig']:
        """
        Trigger settings of the flow.
        """
        return pulumi.get(self, "trigger_config")
