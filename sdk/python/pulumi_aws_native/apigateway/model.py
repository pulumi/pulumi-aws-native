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

__all__ = ['ModelArgs', 'Model']

@pulumi.input_type
class ModelArgs:
    def __init__(__self__, *,
                 rest_api_id: pulumi.Input[builtins.str],
                 content_type: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[Any] = None):
        """
        The set of arguments for constructing a Model resource.
        :param pulumi.Input[builtins.str] rest_api_id: The string identifier of the associated RestApi.
        :param pulumi.Input[builtins.str] content_type: The content-type for the model.
        :param pulumi.Input[builtins.str] description: The description of the model.
        :param pulumi.Input[builtins.str] name: A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
                 If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        :param Any schema: The schema for the model. For `application/json` models, this should be JSON schema draft 4 model. Do not include "\\*/" characters in the description of any properties because such "\\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGateway::Model` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "rest_api_id", rest_api_id)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Input[builtins.str]:
        """
        The string identifier of the associated RestApi.
        """
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "rest_api_id", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The content-type for the model.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The description of the model.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
          If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[Any]:
        """
        The schema for the model. For `application/json` models, this should be JSON schema draft 4 model. Do not include "\\*/" characters in the description of any properties because such "\\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGateway::Model` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[Any]):
        pulumi.set(self, "schema", value)


@pulumi.type_token("aws-native:apigateway:Model")
class Model(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 rest_api_id: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[Any] = None,
                 __props__=None):
        """
        The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] content_type: The content-type for the model.
        :param pulumi.Input[builtins.str] description: The description of the model.
        :param pulumi.Input[builtins.str] name: A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
                 If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        :param pulumi.Input[builtins.str] rest_api_id: The string identifier of the associated RestApi.
        :param Any schema: The schema for the model. For `application/json` models, this should be JSON schema draft 4 model. Do not include "\\*/" characters in the description of any properties because such "\\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGateway::Model` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ModelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.

        :param str resource_name: The name of the resource.
        :param ModelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ModelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 rest_api_id: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModelArgs.__new__(ModelArgs)

            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if rest_api_id is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__.__dict__["rest_api_id"] = rest_api_id
            __props__.__dict__["schema"] = schema
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["contentType", "name", "restApiId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Model, __self__).__init__(
            'aws-native:apigateway:Model',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Model':
        """
        Get an existing Model resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ModelArgs.__new__(ModelArgs)

        __props__.__dict__["content_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["rest_api_id"] = None
        __props__.__dict__["schema"] = None
        return Model(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The content-type for the model.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The description of the model.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
          If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Output[builtins.str]:
        """
        The string identifier of the associated RestApi.
        """
        return pulumi.get(self, "rest_api_id")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[Optional[Any]]:
        """
        The schema for the model. For `application/json` models, this should be JSON schema draft 4 model. Do not include "\\*/" characters in the description of any properties because such "\\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGateway::Model` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "schema")

