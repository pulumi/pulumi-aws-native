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
from ._inputs import *

__all__ = ['FunctionArgs', 'Function']

@pulumi.input_type
class FunctionArgs:
    def __init__(__self__, *,
                 function_code: pulumi.Input[builtins.str],
                 function_config: pulumi.Input['FunctionConfigArgs'],
                 auto_publish: Optional[pulumi.Input[builtins.bool]] = None,
                 function_metadata: Optional[pulumi.Input['FunctionMetadataArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Function resource.
        :param pulumi.Input[builtins.str] function_code: The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.
        :param pulumi.Input['FunctionConfigArgs'] function_config: Contains configuration information about a CloudFront function.
        :param pulumi.Input[builtins.bool] auto_publish: A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``.
        :param pulumi.Input['FunctionMetadataArgs'] function_metadata: Contains metadata about a CloudFront function.
        :param pulumi.Input[builtins.str] name: A name to identify the function.
        """
        pulumi.set(__self__, "function_code", function_code)
        pulumi.set(__self__, "function_config", function_config)
        if auto_publish is not None:
            pulumi.set(__self__, "auto_publish", auto_publish)
        if function_metadata is not None:
            pulumi.set(__self__, "function_metadata", function_metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="functionCode")
    def function_code(self) -> pulumi.Input[builtins.str]:
        """
        The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.
        """
        return pulumi.get(self, "function_code")

    @function_code.setter
    def function_code(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "function_code", value)

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> pulumi.Input['FunctionConfigArgs']:
        """
        Contains configuration information about a CloudFront function.
        """
        return pulumi.get(self, "function_config")

    @function_config.setter
    def function_config(self, value: pulumi.Input['FunctionConfigArgs']):
        pulumi.set(self, "function_config", value)

    @property
    @pulumi.getter(name="autoPublish")
    def auto_publish(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``.
        """
        return pulumi.get(self, "auto_publish")

    @auto_publish.setter
    def auto_publish(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "auto_publish", value)

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> Optional[pulumi.Input['FunctionMetadataArgs']]:
        """
        Contains metadata about a CloudFront function.
        """
        return pulumi.get(self, "function_metadata")

    @function_metadata.setter
    def function_metadata(self, value: Optional[pulumi.Input['FunctionMetadataArgs']]):
        pulumi.set(self, "function_metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A name to identify the function.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("aws-native:cloudfront:Function")
class Function(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_publish: Optional[pulumi.Input[builtins.bool]] = None,
                 function_code: Optional[pulumi.Input[builtins.str]] = None,
                 function_config: Optional[pulumi.Input[Union['FunctionConfigArgs', 'FunctionConfigArgsDict']]] = None,
                 function_metadata: Optional[pulumi.Input[Union['FunctionMetadataArgs', 'FunctionMetadataArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a CF function.
         To create a function, you provide the function code and some configuration information about the function. The response contains an Amazon Resource Name (ARN) that uniquely identifies the function, and the function’s stage.
         By default, when you create a function, it’s in the ``DEVELOPMENT`` stage. In this stage, you can [test the function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/test-function.html) in the CF console (or with ``TestFunction`` in the CF API).
         When you’re ready to use your function with a CF distribution, publish the function to the ``LIVE`` stage. You can do this in the CF console, with ``PublishFunction`` in the CF API, or by updating the ``AWS::CloudFront::Function`` resource with the ``AutoPublish`` property set to ``true``. When the function is published to the ``LIVE`` stage, you can attach it to a distribution’s cache behavior, using the function’s ARN.
         To automatically publish the function to the ``LIVE`` stage when it’s created, set the ``AutoPublish`` property to ``true``.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] auto_publish: A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``.
        :param pulumi.Input[builtins.str] function_code: The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.
        :param pulumi.Input[Union['FunctionConfigArgs', 'FunctionConfigArgsDict']] function_config: Contains configuration information about a CloudFront function.
        :param pulumi.Input[Union['FunctionMetadataArgs', 'FunctionMetadataArgsDict']] function_metadata: Contains metadata about a CloudFront function.
        :param pulumi.Input[builtins.str] name: A name to identify the function.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a CF function.
         To create a function, you provide the function code and some configuration information about the function. The response contains an Amazon Resource Name (ARN) that uniquely identifies the function, and the function’s stage.
         By default, when you create a function, it’s in the ``DEVELOPMENT`` stage. In this stage, you can [test the function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/test-function.html) in the CF console (or with ``TestFunction`` in the CF API).
         When you’re ready to use your function with a CF distribution, publish the function to the ``LIVE`` stage. You can do this in the CF console, with ``PublishFunction`` in the CF API, or by updating the ``AWS::CloudFront::Function`` resource with the ``AutoPublish`` property set to ``true``. When the function is published to the ``LIVE`` stage, you can attach it to a distribution’s cache behavior, using the function’s ARN.
         To automatically publish the function to the ``LIVE`` stage when it’s created, set the ``AutoPublish`` property to ``true``.

        :param str resource_name: The name of the resource.
        :param FunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_publish: Optional[pulumi.Input[builtins.bool]] = None,
                 function_code: Optional[pulumi.Input[builtins.str]] = None,
                 function_config: Optional[pulumi.Input[Union['FunctionConfigArgs', 'FunctionConfigArgsDict']]] = None,
                 function_metadata: Optional[pulumi.Input[Union['FunctionMetadataArgs', 'FunctionMetadataArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionArgs.__new__(FunctionArgs)

            __props__.__dict__["auto_publish"] = auto_publish
            if function_code is None and not opts.urn:
                raise TypeError("Missing required property 'function_code'")
            __props__.__dict__["function_code"] = function_code
            if function_config is None and not opts.urn:
                raise TypeError("Missing required property 'function_config'")
            __props__.__dict__["function_config"] = function_config
            __props__.__dict__["function_metadata"] = function_metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["function_arn"] = None
            __props__.__dict__["stage"] = None
        super(Function, __self__).__init__(
            'aws-native:cloudfront:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FunctionArgs.__new__(FunctionArgs)

        __props__.__dict__["auto_publish"] = None
        __props__.__dict__["function_arn"] = None
        __props__.__dict__["function_code"] = None
        __props__.__dict__["function_config"] = None
        __props__.__dict__["function_metadata"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["stage"] = None
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoPublish")
    def auto_publish(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``.
        """
        return pulumi.get(self, "auto_publish")

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the function. For example:

        `arn:aws:cloudfront::123456789012:function/ExampleFunction` .

        To get the function ARN, use the following syntax:

        `!GetAtt *Function_Logical_ID* .FunctionMetadata.FunctionARN`
        """
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionCode")
    def function_code(self) -> pulumi.Output[builtins.str]:
        """
        The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.
        """
        return pulumi.get(self, "function_code")

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> pulumi.Output['outputs.FunctionConfig']:
        """
        Contains configuration information about a CloudFront function.
        """
        return pulumi.get(self, "function_config")

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> pulumi.Output[Optional['outputs.FunctionMetadata']]:
        """
        Contains metadata about a CloudFront function.
        """
        return pulumi.get(self, "function_metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A name to identify the function.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "stage")

