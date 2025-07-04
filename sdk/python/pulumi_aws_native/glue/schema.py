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

__all__ = ['SchemaArgs', 'Schema']

@pulumi.input_type
class SchemaArgs:
    def __init__(__self__, *,
                 compatibility: pulumi.Input['SchemaCompatibility'],
                 data_format: pulumi.Input['SchemaDataFormat'],
                 checkpoint_version: Optional[pulumi.Input['SchemaVersionArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 registry: Optional[pulumi.Input['SchemaRegistryArgs']] = None,
                 schema_definition: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Schema resource.
        :param pulumi.Input['SchemaCompatibility'] compatibility: Compatibility setting for the schema.
        :param pulumi.Input['SchemaDataFormat'] data_format: Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'
        :param pulumi.Input['SchemaVersionArgs'] checkpoint_version: Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.
        :param pulumi.Input[builtins.str] description: A description of the schema. If description is not provided, there will not be any default value for this.
        :param pulumi.Input[builtins.str] name: Name of the schema.
        :param pulumi.Input['SchemaRegistryArgs'] registry: The registry where a schema is stored.
        :param pulumi.Input[builtins.str] schema_definition: Definition for the initial schema version in plain-text.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: List of tags to tag the schema
        """
        pulumi.set(__self__, "compatibility", compatibility)
        pulumi.set(__self__, "data_format", data_format)
        if checkpoint_version is not None:
            pulumi.set(__self__, "checkpoint_version", checkpoint_version)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if registry is not None:
            pulumi.set(__self__, "registry", registry)
        if schema_definition is not None:
            pulumi.set(__self__, "schema_definition", schema_definition)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def compatibility(self) -> pulumi.Input['SchemaCompatibility']:
        """
        Compatibility setting for the schema.
        """
        return pulumi.get(self, "compatibility")

    @compatibility.setter
    def compatibility(self, value: pulumi.Input['SchemaCompatibility']):
        pulumi.set(self, "compatibility", value)

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> pulumi.Input['SchemaDataFormat']:
        """
        Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'
        """
        return pulumi.get(self, "data_format")

    @data_format.setter
    def data_format(self, value: pulumi.Input['SchemaDataFormat']):
        pulumi.set(self, "data_format", value)

    @property
    @pulumi.getter(name="checkpointVersion")
    def checkpoint_version(self) -> Optional[pulumi.Input['SchemaVersionArgs']]:
        """
        Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.
        """
        return pulumi.get(self, "checkpoint_version")

    @checkpoint_version.setter
    def checkpoint_version(self, value: Optional[pulumi.Input['SchemaVersionArgs']]):
        pulumi.set(self, "checkpoint_version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the schema. If description is not provided, there will not be any default value for this.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the schema.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def registry(self) -> Optional[pulumi.Input['SchemaRegistryArgs']]:
        """
        The registry where a schema is stored.
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: Optional[pulumi.Input['SchemaRegistryArgs']]):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter(name="schemaDefinition")
    def schema_definition(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Definition for the initial schema version in plain-text.
        """
        return pulumi.get(self, "schema_definition")

    @schema_definition.setter
    def schema_definition(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schema_definition", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        List of tags to tag the schema
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:glue:Schema")
class Schema(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 checkpoint_version: Optional[pulumi.Input[Union['SchemaVersionArgs', 'SchemaVersionArgsDict']]] = None,
                 compatibility: Optional[pulumi.Input['SchemaCompatibility']] = None,
                 data_format: Optional[pulumi.Input['SchemaDataFormat']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 registry: Optional[pulumi.Input[Union['SchemaRegistryArgs', 'SchemaRegistryArgsDict']]] = None,
                 schema_definition: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        This resource represents a schema of Glue Schema Registry.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SchemaVersionArgs', 'SchemaVersionArgsDict']] checkpoint_version: Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.
        :param pulumi.Input['SchemaCompatibility'] compatibility: Compatibility setting for the schema.
        :param pulumi.Input['SchemaDataFormat'] data_format: Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'
        :param pulumi.Input[builtins.str] description: A description of the schema. If description is not provided, there will not be any default value for this.
        :param pulumi.Input[builtins.str] name: Name of the schema.
        :param pulumi.Input[Union['SchemaRegistryArgs', 'SchemaRegistryArgsDict']] registry: The registry where a schema is stored.
        :param pulumi.Input[builtins.str] schema_definition: Definition for the initial schema version in plain-text.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: List of tags to tag the schema
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SchemaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource represents a schema of Glue Schema Registry.

        :param str resource_name: The name of the resource.
        :param SchemaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SchemaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 checkpoint_version: Optional[pulumi.Input[Union['SchemaVersionArgs', 'SchemaVersionArgsDict']]] = None,
                 compatibility: Optional[pulumi.Input['SchemaCompatibility']] = None,
                 data_format: Optional[pulumi.Input['SchemaDataFormat']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 registry: Optional[pulumi.Input[Union['SchemaRegistryArgs', 'SchemaRegistryArgsDict']]] = None,
                 schema_definition: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SchemaArgs.__new__(SchemaArgs)

            __props__.__dict__["checkpoint_version"] = checkpoint_version
            if compatibility is None and not opts.urn:
                raise TypeError("Missing required property 'compatibility'")
            __props__.__dict__["compatibility"] = compatibility
            if data_format is None and not opts.urn:
                raise TypeError("Missing required property 'data_format'")
            __props__.__dict__["data_format"] = data_format
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["registry"] = registry
            __props__.__dict__["schema_definition"] = schema_definition
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["initial_schema_version_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["dataFormat", "name", "registry", "schemaDefinition"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Schema, __self__).__init__(
            'aws-native:glue:Schema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Schema':
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SchemaArgs.__new__(SchemaArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["checkpoint_version"] = None
        __props__.__dict__["compatibility"] = None
        __props__.__dict__["data_format"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["initial_schema_version_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["registry"] = None
        __props__.__dict__["schema_definition"] = None
        __props__.__dict__["tags"] = None
        return Schema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        Amazon Resource Name for the Schema.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="checkpointVersion")
    def checkpoint_version(self) -> pulumi.Output[Optional['outputs.SchemaVersion']]:
        """
        Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.
        """
        return pulumi.get(self, "checkpoint_version")

    @property
    @pulumi.getter
    def compatibility(self) -> pulumi.Output['SchemaCompatibility']:
        """
        Compatibility setting for the schema.
        """
        return pulumi.get(self, "compatibility")

    @property
    @pulumi.getter(name="dataFormat")
    def data_format(self) -> pulumi.Output['SchemaDataFormat']:
        """
        Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'
        """
        return pulumi.get(self, "data_format")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description of the schema. If description is not provided, there will not be any default value for this.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="initialSchemaVersionId")
    def initial_schema_version_id(self) -> pulumi.Output[builtins.str]:
        """
        Represents the version ID associated with the initial schema version.
        """
        return pulumi.get(self, "initial_schema_version_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the schema.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Output[Optional['outputs.SchemaRegistry']]:
        """
        The registry where a schema is stored.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter(name="schemaDefinition")
    def schema_definition(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Definition for the initial schema version in plain-text.
        """
        return pulumi.get(self, "schema_definition")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        List of tags to tag the schema
        """
        return pulumi.get(self, "tags")

