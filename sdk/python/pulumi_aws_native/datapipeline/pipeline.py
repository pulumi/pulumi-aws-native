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

__all__ = ['PipelineArgs', 'Pipeline']

@pulumi.input_type
class PipelineArgs:
    def __init__(__self__, *,
                 activate: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameter_objects: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Pipeline resource.
        :param pulumi.Input[builtins.bool] activate: Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
        :param pulumi.Input[builtins.str] description: A description of the pipeline.
        :param pulumi.Input[builtins.str] name: The name of the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]] parameter_objects: The parameter objects used with the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]] parameter_values: The parameter values used with the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]] pipeline_objects: The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] pipeline_tags: A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
        """
        if activate is not None:
            pulumi.set(__self__, "activate", activate)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameter_objects is not None:
            pulumi.set(__self__, "parameter_objects", parameter_objects)
        if parameter_values is not None:
            pulumi.set(__self__, "parameter_values", parameter_values)
        if pipeline_objects is not None:
            pulumi.set(__self__, "pipeline_objects", pipeline_objects)
        if pipeline_tags is not None:
            pulumi.set(__self__, "pipeline_tags", pipeline_tags)

    @property
    @pulumi.getter
    def activate(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
        """
        return pulumi.get(self, "activate")

    @activate.setter
    def activate(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "activate", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the pipeline.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parameterObjects")
    def parameter_objects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]]]:
        """
        The parameter objects used with the pipeline.
        """
        return pulumi.get(self, "parameter_objects")

    @parameter_objects.setter
    def parameter_objects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterObjectArgs']]]]):
        pulumi.set(self, "parameter_objects", value)

    @property
    @pulumi.getter(name="parameterValues")
    def parameter_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]]:
        """
        The parameter values used with the pipeline.
        """
        return pulumi.get(self, "parameter_values")

    @parameter_values.setter
    def parameter_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineParameterValueArgs']]]]):
        pulumi.set(self, "parameter_values", value)

    @property
    @pulumi.getter(name="pipelineObjects")
    def pipeline_objects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]]:
        """
        The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
        """
        return pulumi.get(self, "pipeline_objects")

    @pipeline_objects.setter
    def pipeline_objects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PipelineObjectArgs']]]]):
        pulumi.set(self, "pipeline_objects", value)

    @property
    @pulumi.getter(name="pipelineTags")
    def pipeline_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
        """
        return pulumi.get(self, "pipeline_tags")

    @pipeline_tags.setter
    def pipeline_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "pipeline_tags", value)


@pulumi.type_token("aws-native:datapipeline:Pipeline")
class Pipeline(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameter_objects: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterObjectArgs', 'PipelineParameterObjectArgsDict']]]]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterValueArgs', 'PipelineParameterValueArgsDict']]]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineObjectArgs', 'PipelineObjectArgsDict']]]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] activate: Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
        :param pulumi.Input[builtins.str] description: A description of the pipeline.
        :param pulumi.Input[builtins.str] name: The name of the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterObjectArgs', 'PipelineParameterObjectArgsDict']]]] parameter_objects: The parameter objects used with the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterValueArgs', 'PipelineParameterValueArgsDict']]]] parameter_values: The parameter values used with the pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[Union['PipelineObjectArgs', 'PipelineObjectArgsDict']]]] pipeline_objects: The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] pipeline_tags: A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PipelineArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param PipelineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PipelineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activate: Optional[pulumi.Input[builtins.bool]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parameter_objects: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterObjectArgs', 'PipelineParameterObjectArgsDict']]]]] = None,
                 parameter_values: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineParameterValueArgs', 'PipelineParameterValueArgsDict']]]]] = None,
                 pipeline_objects: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PipelineObjectArgs', 'PipelineObjectArgsDict']]]]] = None,
                 pipeline_tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PipelineArgs.__new__(PipelineArgs)

            __props__.__dict__["activate"] = activate
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["parameter_objects"] = parameter_objects
            __props__.__dict__["parameter_values"] = parameter_values
            __props__.__dict__["pipeline_objects"] = pipeline_objects
            __props__.__dict__["pipeline_tags"] = pipeline_tags
            __props__.__dict__["pipeline_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Pipeline, __self__).__init__(
            'aws-native:datapipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PipelineArgs.__new__(PipelineArgs)

        __props__.__dict__["activate"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameter_objects"] = None
        __props__.__dict__["parameter_values"] = None
        __props__.__dict__["pipeline_id"] = None
        __props__.__dict__["pipeline_objects"] = None
        __props__.__dict__["pipeline_tags"] = None
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def activate(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.
        """
        return pulumi.get(self, "activate")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description of the pipeline.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the pipeline.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parameterObjects")
    def parameter_objects(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineParameterObject']]]:
        """
        The parameter objects used with the pipeline.
        """
        return pulumi.get(self, "parameter_objects")

    @property
    @pulumi.getter(name="parameterValues")
    def parameter_values(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineParameterValue']]]:
        """
        The parameter values used with the pipeline.
        """
        return pulumi.get(self, "parameter_values")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="pipelineObjects")
    def pipeline_objects(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineObject']]]:
        """
        The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.
        """
        return pulumi.get(self, "pipeline_objects")

    @property
    @pulumi.getter(name="pipelineTags")
    def pipeline_tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.
        """
        return pulumi.get(self, "pipeline_tags")

