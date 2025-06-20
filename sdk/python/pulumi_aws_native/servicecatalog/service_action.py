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
from ._enums import *
from ._inputs import *

__all__ = ['ServiceActionArgs', 'ServiceAction']

@pulumi.input_type
class ServiceActionArgs:
    def __init__(__self__, *,
                 definition: pulumi.Input[Sequence[pulumi.Input['ServiceActionDefinitionParameterArgs']]],
                 definition_type: pulumi.Input['ServiceActionDefinitionType'],
                 accept_language: Optional[pulumi.Input['ServiceActionAcceptLanguage']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ServiceAction resource.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceActionDefinitionParameterArgs']]] definition: A map that defines the self-service action.
        :param pulumi.Input['ServiceActionDefinitionType'] definition_type: The self-service action definition type. For example, `SSM_AUTOMATION` .
        :param pulumi.Input['ServiceActionAcceptLanguage'] accept_language: The language code.
               
               - `en` - English (default)
               - `jp` - Japanese
               - `zh` - Chinese
        :param pulumi.Input[builtins.str] description: The self-service action description.
        :param pulumi.Input[builtins.str] name: The self-service action name.
        """
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "definition_type", definition_type)
        if accept_language is not None:
            pulumi.set(__self__, "accept_language", accept_language)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Input[Sequence[pulumi.Input['ServiceActionDefinitionParameterArgs']]]:
        """
        A map that defines the self-service action.
        """
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: pulumi.Input[Sequence[pulumi.Input['ServiceActionDefinitionParameterArgs']]]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter(name="definitionType")
    def definition_type(self) -> pulumi.Input['ServiceActionDefinitionType']:
        """
        The self-service action definition type. For example, `SSM_AUTOMATION` .
        """
        return pulumi.get(self, "definition_type")

    @definition_type.setter
    def definition_type(self, value: pulumi.Input['ServiceActionDefinitionType']):
        pulumi.set(self, "definition_type", value)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[pulumi.Input['ServiceActionAcceptLanguage']]:
        """
        The language code.

        - `en` - English (default)
        - `jp` - Japanese
        - `zh` - Chinese
        """
        return pulumi.get(self, "accept_language")

    @accept_language.setter
    def accept_language(self, value: Optional[pulumi.Input['ServiceActionAcceptLanguage']]):
        pulumi.set(self, "accept_language", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The self-service action description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The self-service action name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("aws-native:servicecatalog:ServiceAction")
class ServiceAction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input['ServiceActionAcceptLanguage']] = None,
                 definition: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceActionDefinitionParameterArgs', 'ServiceActionDefinitionParameterArgsDict']]]]] = None,
                 definition_type: Optional[pulumi.Input['ServiceActionDefinitionType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource Schema for AWS::ServiceCatalog::ServiceAction

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['ServiceActionAcceptLanguage'] accept_language: The language code.
               
               - `en` - English (default)
               - `jp` - Japanese
               - `zh` - Chinese
        :param pulumi.Input[Sequence[pulumi.Input[Union['ServiceActionDefinitionParameterArgs', 'ServiceActionDefinitionParameterArgsDict']]]] definition: A map that defines the self-service action.
        :param pulumi.Input['ServiceActionDefinitionType'] definition_type: The self-service action definition type. For example, `SSM_AUTOMATION` .
        :param pulumi.Input[builtins.str] description: The self-service action description.
        :param pulumi.Input[builtins.str] name: The self-service action name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceActionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Schema for AWS::ServiceCatalog::ServiceAction

        :param str resource_name: The name of the resource.
        :param ServiceActionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceActionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_language: Optional[pulumi.Input['ServiceActionAcceptLanguage']] = None,
                 definition: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ServiceActionDefinitionParameterArgs', 'ServiceActionDefinitionParameterArgsDict']]]]] = None,
                 definition_type: Optional[pulumi.Input['ServiceActionDefinitionType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceActionArgs.__new__(ServiceActionArgs)

            __props__.__dict__["accept_language"] = accept_language
            if definition is None and not opts.urn:
                raise TypeError("Missing required property 'definition'")
            __props__.__dict__["definition"] = definition
            if definition_type is None and not opts.urn:
                raise TypeError("Missing required property 'definition_type'")
            __props__.__dict__["definition_type"] = definition_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["aws_id"] = None
        super(ServiceAction, __self__).__init__(
            'aws-native:servicecatalog:ServiceAction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceAction':
        """
        Get an existing ServiceAction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceActionArgs.__new__(ServiceActionArgs)

        __props__.__dict__["accept_language"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["definition_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        return ServiceAction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> pulumi.Output[Optional['ServiceActionAcceptLanguage']]:
        """
        The language code.

        - `en` - English (default)
        - `jp` - Japanese
        - `zh` - Chinese
        """
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The self-service action identifier. For example, `act-fs7abcd89wxyz` .
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[Sequence['outputs.ServiceActionDefinitionParameter']]:
        """
        A map that defines the self-service action.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="definitionType")
    def definition_type(self) -> pulumi.Output['ServiceActionDefinitionType']:
        """
        The self-service action definition type. For example, `SSM_AUTOMATION` .
        """
        return pulumi.get(self, "definition_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The self-service action description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The self-service action name.
        """
        return pulumi.get(self, "name")

