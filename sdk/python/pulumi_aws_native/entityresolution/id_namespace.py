# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['IdNamespaceArgs', 'IdNamespace']

@pulumi.input_type
class IdNamespaceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input['IdNamespaceType'],
                 description: Optional[pulumi.Input[str]] = None,
                 id_mapping_workflow_properties: Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceIdMappingWorkflowPropertiesArgs']]]] = None,
                 id_namespace_name: Optional[pulumi.Input[str]] = None,
                 input_source_config: Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceInputSourceArgs']]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a IdNamespace resource.
        :param pulumi.Input['IdNamespaceType'] type: The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
               
               The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
               
               The `TARGET` contains a configuration of `targetId` to which all `sourceIds` will resolve to.
        :param pulumi.Input[str] description: The description of the ID namespace.
        :param pulumi.Input[Sequence[pulumi.Input['IdNamespaceIdMappingWorkflowPropertiesArgs']]] id_mapping_workflow_properties: Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
        :param pulumi.Input[str] id_namespace_name: The name of the ID namespace.
        :param pulumi.Input[Sequence[pulumi.Input['IdNamespaceInputSourceArgs']]] input_source_config: A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The tags used to organize, track, or control access for this resource.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id_mapping_workflow_properties is not None:
            pulumi.set(__self__, "id_mapping_workflow_properties", id_mapping_workflow_properties)
        if id_namespace_name is not None:
            pulumi.set(__self__, "id_namespace_name", id_namespace_name)
        if input_source_config is not None:
            pulumi.set(__self__, "input_source_config", input_source_config)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['IdNamespaceType']:
        """
        The type of ID namespace. There are two types: `SOURCE` and `TARGET` .

        The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.

        The `TARGET` contains a configuration of `targetId` to which all `sourceIds` will resolve to.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['IdNamespaceType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the ID namespace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="idMappingWorkflowProperties")
    def id_mapping_workflow_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceIdMappingWorkflowPropertiesArgs']]]]:
        """
        Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
        """
        return pulumi.get(self, "id_mapping_workflow_properties")

    @id_mapping_workflow_properties.setter
    def id_mapping_workflow_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceIdMappingWorkflowPropertiesArgs']]]]):
        pulumi.set(self, "id_mapping_workflow_properties", value)

    @property
    @pulumi.getter(name="idNamespaceName")
    def id_namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ID namespace.
        """
        return pulumi.get(self, "id_namespace_name")

    @id_namespace_name.setter
    def id_namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id_namespace_name", value)

    @property
    @pulumi.getter(name="inputSourceConfig")
    def input_source_config(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceInputSourceArgs']]]]:
        """
        A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
        """
        return pulumi.get(self, "input_source_config")

    @input_source_config.setter
    def input_source_config(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdNamespaceInputSourceArgs']]]]):
        pulumi.set(self, "input_source_config", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The tags used to organize, track, or control access for this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class IdNamespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id_mapping_workflow_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceIdMappingWorkflowPropertiesArgs']]]]] = None,
                 id_namespace_name: Optional[pulumi.Input[str]] = None,
                 input_source_config: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceInputSourceArgs']]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input['IdNamespaceType']] = None,
                 __props__=None):
        """
        IdNamespace defined in AWS Entity Resolution service

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the ID namespace.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceIdMappingWorkflowPropertiesArgs']]]] id_mapping_workflow_properties: Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
        :param pulumi.Input[str] id_namespace_name: The name of the ID namespace.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceInputSourceArgs']]]] input_source_config: A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: The tags used to organize, track, or control access for this resource.
        :param pulumi.Input['IdNamespaceType'] type: The type of ID namespace. There are two types: `SOURCE` and `TARGET` .
               
               The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.
               
               The `TARGET` contains a configuration of `targetId` to which all `sourceIds` will resolve to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdNamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        IdNamespace defined in AWS Entity Resolution service

        :param str resource_name: The name of the resource.
        :param IdNamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdNamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id_mapping_workflow_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceIdMappingWorkflowPropertiesArgs']]]]] = None,
                 id_namespace_name: Optional[pulumi.Input[str]] = None,
                 input_source_config: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdNamespaceInputSourceArgs']]]]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input['IdNamespaceType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdNamespaceArgs.__new__(IdNamespaceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["id_mapping_workflow_properties"] = id_mapping_workflow_properties
            __props__.__dict__["id_namespace_name"] = id_namespace_name
            __props__.__dict__["input_source_config"] = input_source_config
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["id_namespace_arn"] = None
            __props__.__dict__["updated_at"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["idNamespaceName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(IdNamespace, __self__).__init__(
            'aws-native:entityresolution:IdNamespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IdNamespace':
        """
        Get an existing IdNamespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IdNamespaceArgs.__new__(IdNamespaceArgs)

        __props__.__dict__["created_at"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["id_mapping_workflow_properties"] = None
        __props__.__dict__["id_namespace_arn"] = None
        __props__.__dict__["id_namespace_name"] = None
        __props__.__dict__["input_source_config"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["updated_at"] = None
        return IdNamespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time when the IdNamespace was created
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the ID namespace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="idMappingWorkflowProperties")
    def id_mapping_workflow_properties(self) -> pulumi.Output[Optional[Sequence['outputs.IdNamespaceIdMappingWorkflowProperties']]]:
        """
        Determines the properties of `IdMappingWorflow` where this `IdNamespace` can be used as a `Source` or a `Target` .
        """
        return pulumi.get(self, "id_mapping_workflow_properties")

    @property
    @pulumi.getter(name="idNamespaceArn")
    def id_namespace_arn(self) -> pulumi.Output[str]:
        """
        The arn associated with the IdNamespace
        """
        return pulumi.get(self, "id_namespace_arn")

    @property
    @pulumi.getter(name="idNamespaceName")
    def id_namespace_name(self) -> pulumi.Output[str]:
        """
        The name of the ID namespace.
        """
        return pulumi.get(self, "id_namespace_name")

    @property
    @pulumi.getter(name="inputSourceConfig")
    def input_source_config(self) -> pulumi.Output[Optional[Sequence['outputs.IdNamespaceInputSource']]]:
        """
        A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName` .
        """
        return pulumi.get(self, "input_source_config")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes this role to access the resources defined in this `IdNamespace` on your behalf as part of the workflow run.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The tags used to organize, track, or control access for this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output['IdNamespaceType']:
        """
        The type of ID namespace. There are two types: `SOURCE` and `TARGET` .

        The `SOURCE` contains configurations for `sourceId` data that will be processed in an ID mapping workflow.

        The `TARGET` contains a configuration of `targetId` to which all `sourceIds` will resolve to.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time when the IdNamespace was updated
        """
        return pulumi.get(self, "updated_at")
