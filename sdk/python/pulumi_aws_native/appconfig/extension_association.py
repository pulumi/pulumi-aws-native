# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['ExtensionAssociationArgs', 'ExtensionAssociation']

@pulumi.input_type
class ExtensionAssociationArgs:
    def __init__(__self__, *,
                 extension_identifier: Optional[pulumi.Input[str]] = None,
                 extension_version_number: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]] = None):
        """
        The set of arguments for constructing a ExtensionAssociation resource.
        :param pulumi.Input[str] extension_identifier: The name, the ID, or the Amazon Resource Name (ARN) of the extension.
        :param pulumi.Input[int] extension_version_number: The version number of the extension. If not specified, AWS AppConfig uses the maximum version of the extension.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: The parameter names and values defined in the extensions. Extension parameters marked `Required` must be entered for this field.
        :param pulumi.Input[str] resource_identifier: The ARN of an application, configuration profile, or environment.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if extension_identifier is not None:
            pulumi.set(__self__, "extension_identifier", extension_identifier)
        if extension_version_number is not None:
            pulumi.set(__self__, "extension_version_number", extension_version_number)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if resource_identifier is not None:
            pulumi.set(__self__, "resource_identifier", resource_identifier)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="extensionIdentifier")
    def extension_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The name, the ID, or the Amazon Resource Name (ARN) of the extension.
        """
        return pulumi.get(self, "extension_identifier")

    @extension_identifier.setter
    def extension_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extension_identifier", value)

    @property
    @pulumi.getter(name="extensionVersionNumber")
    def extension_version_number(self) -> Optional[pulumi.Input[int]]:
        """
        The version number of the extension. If not specified, AWS AppConfig uses the maximum version of the extension.
        """
        return pulumi.get(self, "extension_version_number")

    @extension_version_number.setter
    def extension_version_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "extension_version_number", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The parameter names and values defined in the extensions. Extension parameters marked `Required` must be entered for this field.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="resourceIdentifier")
    def resource_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an application, configuration profile, or environment.
        """
        return pulumi.get(self, "resource_identifier")

    @resource_identifier.setter
    def resource_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_identifier", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]):
        pulumi.set(self, "tags", value)


class ExtensionAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extension_identifier: Optional[pulumi.Input[str]] = None,
                 extension_version_number: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] extension_identifier: The name, the ID, or the Amazon Resource Name (ARN) of the extension.
        :param pulumi.Input[int] extension_version_number: The version number of the extension. If not specified, AWS AppConfig uses the maximum version of the extension.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: The parameter names and values defined in the extensions. Extension parameters marked `Required` must be entered for this field.
        :param pulumi.Input[str] resource_identifier: The ARN of an application, configuration profile, or environment.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExtensionAssociationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param ExtensionAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtensionAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extension_identifier: Optional[pulumi.Input[str]] = None,
                 extension_version_number: Optional[pulumi.Input[int]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExtensionAssociationArgs.__new__(ExtensionAssociationArgs)

            __props__.__dict__["extension_identifier"] = extension_identifier
            __props__.__dict__["extension_version_number"] = extension_version_number
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["resource_identifier"] = resource_identifier
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["extension_arn"] = None
            __props__.__dict__["resource_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["extensionIdentifier", "extensionVersionNumber", "resourceIdentifier", "tags[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ExtensionAssociation, __self__).__init__(
            'aws-native:appconfig:ExtensionAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExtensionAssociation':
        """
        Get an existing ExtensionAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExtensionAssociationArgs.__new__(ExtensionAssociationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["extension_arn"] = None
        __props__.__dict__["extension_identifier"] = None
        __props__.__dict__["extension_version_number"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["resource_arn"] = None
        __props__.__dict__["resource_identifier"] = None
        __props__.__dict__["tags"] = None
        return ExtensionAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the extension defined in the association.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        """
        The system-generated ID for the association.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="extensionArn")
    def extension_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the extension defined in the association.
        """
        return pulumi.get(self, "extension_arn")

    @property
    @pulumi.getter(name="extensionIdentifier")
    def extension_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The name, the ID, or the Amazon Resource Name (ARN) of the extension.
        """
        return pulumi.get(self, "extension_identifier")

    @property
    @pulumi.getter(name="extensionVersionNumber")
    def extension_version_number(self) -> pulumi.Output[Optional[int]]:
        """
        The version number of the extension. If not specified, AWS AppConfig uses the maximum version of the extension.
        """
        return pulumi.get(self, "extension_version_number")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The parameter names and values defined in the extensions. Extension parameters marked `Required` must be entered for this field.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> pulumi.Output[str]:
        """
        The ARNs of applications, configuration profiles, or environments defined in the association.
        """
        return pulumi.get(self, "resource_arn")

    @property
    @pulumi.getter(name="resourceIdentifier")
    def resource_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of an application, configuration profile, or environment.
        """
        return pulumi.get(self, "resource_identifier")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.CreateOnlyTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

