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

__all__ = ['WorkflowArgs', 'Workflow']

@pulumi.input_type
class WorkflowArgs:
    def __init__(__self__, *,
                 definition_uri: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input['WorkflowEngine']] = None,
                 main: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_template: Optional[pulumi.Input['WorkflowParameterTemplateArgs']] = None,
                 storage_capacity: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input['WorkflowTagMapArgs']] = None):
        """
        The set of arguments for constructing a Workflow resource.
        """
        if definition_uri is not None:
            pulumi.set(__self__, "definition_uri", definition_uri)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if main is not None:
            pulumi.set(__self__, "main", main)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameter_template is not None:
            pulumi.set(__self__, "parameter_template", parameter_template)
        if storage_capacity is not None:
            pulumi.set(__self__, "storage_capacity", storage_capacity)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="definitionUri")
    def definition_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "definition_uri")

    @definition_uri.setter
    def definition_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "definition_uri", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input['WorkflowEngine']]:
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input['WorkflowEngine']]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def main(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "main")

    @main.setter
    def main(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "main", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parameterTemplate")
    def parameter_template(self) -> Optional[pulumi.Input['WorkflowParameterTemplateArgs']]:
        return pulumi.get(self, "parameter_template")

    @parameter_template.setter
    def parameter_template(self, value: Optional[pulumi.Input['WorkflowParameterTemplateArgs']]):
        pulumi.set(self, "parameter_template", value)

    @property
    @pulumi.getter(name="storageCapacity")
    def storage_capacity(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "storage_capacity")

    @storage_capacity.setter
    def storage_capacity(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "storage_capacity", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['WorkflowTagMapArgs']]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['WorkflowTagMapArgs']]):
        pulumi.set(self, "tags", value)


class Workflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition_uri: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input['WorkflowEngine']] = None,
                 main: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_template: Optional[pulumi.Input[pulumi.InputType['WorkflowParameterTemplateArgs']]] = None,
                 storage_capacity: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['WorkflowTagMapArgs']]] = None,
                 __props__=None):
        """
        Definition of AWS::Omics::Workflow Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WorkflowArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Omics::Workflow Resource Type

        :param str resource_name: The name of the resource.
        :param WorkflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition_uri: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input['WorkflowEngine']] = None,
                 main: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameter_template: Optional[pulumi.Input[pulumi.InputType['WorkflowParameterTemplateArgs']]] = None,
                 storage_capacity: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['WorkflowTagMapArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkflowArgs.__new__(WorkflowArgs)

            __props__.__dict__["definition_uri"] = definition_uri
            __props__.__dict__["description"] = description
            __props__.__dict__["engine"] = engine
            __props__.__dict__["main"] = main
            __props__.__dict__["name"] = name
            __props__.__dict__["parameter_template"] = parameter_template
            __props__.__dict__["storage_capacity"] = storage_capacity
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        super(Workflow, __self__).__init__(
            'aws-native:omics:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workflow':
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkflowArgs.__new__(WorkflowArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["definition_uri"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["engine"] = None
        __props__.__dict__["main"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameter_template"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["storage_capacity"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Workflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="definitionUri")
    def definition_uri(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "definition_uri")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional['WorkflowEngine']]:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def main(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "main")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parameterTemplate")
    def parameter_template(self) -> pulumi.Output[Optional['outputs.WorkflowParameterTemplate']]:
        return pulumi.get(self, "parameter_template")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['WorkflowStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageCapacity")
    def storage_capacity(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "storage_capacity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.WorkflowTagMap']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['WorkflowType']:
        return pulumi.get(self, "type")
