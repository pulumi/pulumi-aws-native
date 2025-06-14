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

__all__ = ['FormArgs', 'Form']

@pulumi.input_type
class FormArgs:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 cta: Optional[pulumi.Input['FormCtaArgs']] = None,
                 data_type: Optional[pulumi.Input['FormDataTypeConfigArgs']] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FormFieldConfigArgs']]]] = None,
                 form_action_type: Optional[pulumi.Input['FormActionType']] = None,
                 label_decorator: Optional[pulumi.Input['FormLabelDecorator']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 schema_version: Optional[pulumi.Input[builtins.str]] = None,
                 sectional_elements: Optional[pulumi.Input[Mapping[str, pulumi.Input['FormSectionalElementArgs']]]] = None,
                 style: Optional[pulumi.Input['FormStyleArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Form resource.
        :param pulumi.Input[builtins.str] app_id: The unique ID of the Amplify app associated with the form.
        :param pulumi.Input['FormCtaArgs'] cta: The `FormCTA` object that stores the call to action configuration for the form.
        :param pulumi.Input['FormDataTypeConfigArgs'] data_type: The type of data source to use to create the form.
        :param pulumi.Input[builtins.str] environment_name: The name of the backend environment that is a part of the Amplify app.
        :param pulumi.Input[Mapping[str, pulumi.Input['FormFieldConfigArgs']]] fields: The configuration information for the form's fields.
        :param pulumi.Input['FormActionType'] form_action_type: Specifies whether to perform a create or update action on the form.
        :param pulumi.Input['FormLabelDecorator'] label_decorator: Specifies an icon or decoration to display on the form.
        :param pulumi.Input[builtins.str] name: The name of the form.
        :param pulumi.Input[builtins.str] schema_version: The schema version of the form.
        :param pulumi.Input[Mapping[str, pulumi.Input['FormSectionalElementArgs']]] sectional_elements: The configuration information for the visual helper elements for the form. These elements are not associated with any data.
        :param pulumi.Input['FormStyleArgs'] style: The configuration for the form's style.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: One or more key-value pairs to use when tagging the form data.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if cta is not None:
            pulumi.set(__self__, "cta", cta)
        if data_type is not None:
            pulumi.set(__self__, "data_type", data_type)
        if environment_name is not None:
            pulumi.set(__self__, "environment_name", environment_name)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if form_action_type is not None:
            pulumi.set(__self__, "form_action_type", form_action_type)
        if label_decorator is not None:
            pulumi.set(__self__, "label_decorator", label_decorator)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema_version is not None:
            pulumi.set(__self__, "schema_version", schema_version)
        if sectional_elements is not None:
            pulumi.set(__self__, "sectional_elements", sectional_elements)
        if style is not None:
            pulumi.set(__self__, "style", style)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique ID of the Amplify app associated with the form.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def cta(self) -> Optional[pulumi.Input['FormCtaArgs']]:
        """
        The `FormCTA` object that stores the call to action configuration for the form.
        """
        return pulumi.get(self, "cta")

    @cta.setter
    def cta(self, value: Optional[pulumi.Input['FormCtaArgs']]):
        pulumi.set(self, "cta", value)

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[pulumi.Input['FormDataTypeConfigArgs']]:
        """
        The type of data source to use to create the form.
        """
        return pulumi.get(self, "data_type")

    @data_type.setter
    def data_type(self, value: Optional[pulumi.Input['FormDataTypeConfigArgs']]):
        pulumi.set(self, "data_type", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the backend environment that is a part of the Amplify app.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FormFieldConfigArgs']]]]:
        """
        The configuration information for the form's fields.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FormFieldConfigArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="formActionType")
    def form_action_type(self) -> Optional[pulumi.Input['FormActionType']]:
        """
        Specifies whether to perform a create or update action on the form.
        """
        return pulumi.get(self, "form_action_type")

    @form_action_type.setter
    def form_action_type(self, value: Optional[pulumi.Input['FormActionType']]):
        pulumi.set(self, "form_action_type", value)

    @property
    @pulumi.getter(name="labelDecorator")
    def label_decorator(self) -> Optional[pulumi.Input['FormLabelDecorator']]:
        """
        Specifies an icon or decoration to display on the form.
        """
        return pulumi.get(self, "label_decorator")

    @label_decorator.setter
    def label_decorator(self, value: Optional[pulumi.Input['FormLabelDecorator']]):
        pulumi.set(self, "label_decorator", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the form.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="schemaVersion")
    def schema_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The schema version of the form.
        """
        return pulumi.get(self, "schema_version")

    @schema_version.setter
    def schema_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schema_version", value)

    @property
    @pulumi.getter(name="sectionalElements")
    def sectional_elements(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FormSectionalElementArgs']]]]:
        """
        The configuration information for the visual helper elements for the form. These elements are not associated with any data.
        """
        return pulumi.get(self, "sectional_elements")

    @sectional_elements.setter
    def sectional_elements(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FormSectionalElementArgs']]]]):
        pulumi.set(self, "sectional_elements", value)

    @property
    @pulumi.getter
    def style(self) -> Optional[pulumi.Input['FormStyleArgs']]:
        """
        The configuration for the form's style.
        """
        return pulumi.get(self, "style")

    @style.setter
    def style(self, value: Optional[pulumi.Input['FormStyleArgs']]):
        pulumi.set(self, "style", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        One or more key-value pairs to use when tagging the form data.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:amplifyuibuilder:Form")
class Form(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 cta: Optional[pulumi.Input[Union['FormCtaArgs', 'FormCtaArgsDict']]] = None,
                 data_type: Optional[pulumi.Input[Union['FormDataTypeConfigArgs', 'FormDataTypeConfigArgsDict']]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FormFieldConfigArgs', 'FormFieldConfigArgsDict']]]]] = None,
                 form_action_type: Optional[pulumi.Input['FormActionType']] = None,
                 label_decorator: Optional[pulumi.Input['FormLabelDecorator']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 schema_version: Optional[pulumi.Input[builtins.str]] = None,
                 sectional_elements: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FormSectionalElementArgs', 'FormSectionalElementArgsDict']]]]] = None,
                 style: Optional[pulumi.Input[Union['FormStyleArgs', 'FormStyleArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Definition of AWS::AmplifyUIBuilder::Form Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] app_id: The unique ID of the Amplify app associated with the form.
        :param pulumi.Input[Union['FormCtaArgs', 'FormCtaArgsDict']] cta: The `FormCTA` object that stores the call to action configuration for the form.
        :param pulumi.Input[Union['FormDataTypeConfigArgs', 'FormDataTypeConfigArgsDict']] data_type: The type of data source to use to create the form.
        :param pulumi.Input[builtins.str] environment_name: The name of the backend environment that is a part of the Amplify app.
        :param pulumi.Input[Mapping[str, pulumi.Input[Union['FormFieldConfigArgs', 'FormFieldConfigArgsDict']]]] fields: The configuration information for the form's fields.
        :param pulumi.Input['FormActionType'] form_action_type: Specifies whether to perform a create or update action on the form.
        :param pulumi.Input['FormLabelDecorator'] label_decorator: Specifies an icon or decoration to display on the form.
        :param pulumi.Input[builtins.str] name: The name of the form.
        :param pulumi.Input[builtins.str] schema_version: The schema version of the form.
        :param pulumi.Input[Mapping[str, pulumi.Input[Union['FormSectionalElementArgs', 'FormSectionalElementArgsDict']]]] sectional_elements: The configuration information for the visual helper elements for the form. These elements are not associated with any data.
        :param pulumi.Input[Union['FormStyleArgs', 'FormStyleArgsDict']] style: The configuration for the form's style.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: One or more key-value pairs to use when tagging the form data.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FormArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::AmplifyUIBuilder::Form Resource Type

        :param str resource_name: The name of the resource.
        :param FormArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FormArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[builtins.str]] = None,
                 cta: Optional[pulumi.Input[Union['FormCtaArgs', 'FormCtaArgsDict']]] = None,
                 data_type: Optional[pulumi.Input[Union['FormDataTypeConfigArgs', 'FormDataTypeConfigArgsDict']]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FormFieldConfigArgs', 'FormFieldConfigArgsDict']]]]] = None,
                 form_action_type: Optional[pulumi.Input['FormActionType']] = None,
                 label_decorator: Optional[pulumi.Input['FormLabelDecorator']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 schema_version: Optional[pulumi.Input[builtins.str]] = None,
                 sectional_elements: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FormSectionalElementArgs', 'FormSectionalElementArgsDict']]]]] = None,
                 style: Optional[pulumi.Input[Union['FormStyleArgs', 'FormStyleArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FormArgs.__new__(FormArgs)

            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["cta"] = cta
            __props__.__dict__["data_type"] = data_type
            __props__.__dict__["environment_name"] = environment_name
            __props__.__dict__["fields"] = fields
            __props__.__dict__["form_action_type"] = form_action_type
            __props__.__dict__["label_decorator"] = label_decorator
            __props__.__dict__["name"] = name
            __props__.__dict__["schema_version"] = schema_version
            __props__.__dict__["sectional_elements"] = sectional_elements
            __props__.__dict__["style"] = style
            __props__.__dict__["tags"] = tags
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["appId", "environmentName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Form, __self__).__init__(
            'aws-native:amplifyuibuilder:Form',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Form':
        """
        Get an existing Form resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FormArgs.__new__(FormArgs)

        __props__.__dict__["app_id"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["cta"] = None
        __props__.__dict__["data_type"] = None
        __props__.__dict__["environment_name"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["form_action_type"] = None
        __props__.__dict__["label_decorator"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["schema_version"] = None
        __props__.__dict__["sectional_elements"] = None
        __props__.__dict__["style"] = None
        __props__.__dict__["tags"] = None
        return Form(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The unique ID of the Amplify app associated with the form.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID for the form.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def cta(self) -> pulumi.Output[Optional['outputs.FormCta']]:
        """
        The `FormCTA` object that stores the call to action configuration for the form.
        """
        return pulumi.get(self, "cta")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> pulumi.Output[Optional['outputs.FormDataTypeConfig']]:
        """
        The type of data source to use to create the form.
        """
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the backend environment that is a part of the Amplify app.
        """
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.FormFieldConfig']]]:
        """
        The configuration information for the form's fields.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="formActionType")
    def form_action_type(self) -> pulumi.Output[Optional['FormActionType']]:
        """
        Specifies whether to perform a create or update action on the form.
        """
        return pulumi.get(self, "form_action_type")

    @property
    @pulumi.getter(name="labelDecorator")
    def label_decorator(self) -> pulumi.Output[Optional['FormLabelDecorator']]:
        """
        Specifies an icon or decoration to display on the form.
        """
        return pulumi.get(self, "label_decorator")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the form.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="schemaVersion")
    def schema_version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The schema version of the form.
        """
        return pulumi.get(self, "schema_version")

    @property
    @pulumi.getter(name="sectionalElements")
    def sectional_elements(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.FormSectionalElement']]]:
        """
        The configuration information for the visual helper elements for the form. These elements are not associated with any data.
        """
        return pulumi.get(self, "sectional_elements")

    @property
    @pulumi.getter
    def style(self) -> pulumi.Output[Optional['outputs.FormStyle']]:
        """
        The configuration for the form's style.
        """
        return pulumi.get(self, "style")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        One or more key-value pairs to use when tagging the form data.
        """
        return pulumi.get(self, "tags")

