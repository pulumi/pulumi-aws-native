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
from . import outputs

__all__ = [
    'GetFunctionResult',
    'AwaitableGetFunctionResult',
    'get_function',
    'get_function_output',
]

@pulumi.output_type
class GetFunctionResult:
    def __init__(__self__, function_arn=None, function_code=None, function_config=None, function_metadata=None, name=None, stage=None):
        if function_arn and not isinstance(function_arn, str):
            raise TypeError("Expected argument 'function_arn' to be a str")
        pulumi.set(__self__, "function_arn", function_arn)
        if function_code and not isinstance(function_code, str):
            raise TypeError("Expected argument 'function_code' to be a str")
        pulumi.set(__self__, "function_code", function_code)
        if function_config and not isinstance(function_config, dict):
            raise TypeError("Expected argument 'function_config' to be a dict")
        pulumi.set(__self__, "function_config", function_config)
        if function_metadata and not isinstance(function_metadata, dict):
            raise TypeError("Expected argument 'function_metadata' to be a dict")
        pulumi.set(__self__, "function_metadata", function_metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if stage and not isinstance(stage, str):
            raise TypeError("Expected argument 'stage' to be a str")
        pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> Optional[str]:
        """
        The ARN of the function. For example:

        `arn:aws:cloudfront::123456789012:function/ExampleFunction` .

        To get the function ARN, use the following syntax:

        `!GetAtt *Function_Logical_ID* .FunctionMetadata.FunctionARN`
        """
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionCode")
    def function_code(self) -> Optional[str]:
        """
        The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide* .
        """
        return pulumi.get(self, "function_code")

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> Optional['outputs.FunctionConfig']:
        """
        Contains configuration information about a CloudFront function.
        """
        return pulumi.get(self, "function_config")

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> Optional['outputs.FunctionMetadata']:
        """
        Contains metadata about a CloudFront function.
        """
        return pulumi.get(self, "function_metadata")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        A name to identify the function.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def stage(self) -> Optional[str]:
        return pulumi.get(self, "stage")


class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            function_arn=self.function_arn,
            function_code=self.function_code,
            function_config=self.function_config,
            function_metadata=self.function_metadata,
            name=self.name,
            stage=self.stage)


def get_function(function_arn: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionResult:
    """
    Resource Type definition for AWS::CloudFront::Function


    :param str function_arn: The ARN of the function. For example:
           
           `arn:aws:cloudfront::123456789012:function/ExampleFunction` .
           
           To get the function ARN, use the following syntax:
           
           `!GetAtt *Function_Logical_ID* .FunctionMetadata.FunctionARN`
    """
    __args__ = dict()
    __args__['functionArn'] = function_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudfront:getFunction', __args__, opts=opts, typ=GetFunctionResult).value

    return AwaitableGetFunctionResult(
        function_arn=pulumi.get(__ret__, 'function_arn'),
        function_code=pulumi.get(__ret__, 'function_code'),
        function_config=pulumi.get(__ret__, 'function_config'),
        function_metadata=pulumi.get(__ret__, 'function_metadata'),
        name=pulumi.get(__ret__, 'name'),
        stage=pulumi.get(__ret__, 'stage'))
def get_function_output(function_arn: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionResult]:
    """
    Resource Type definition for AWS::CloudFront::Function


    :param str function_arn: The ARN of the function. For example:
           
           `arn:aws:cloudfront::123456789012:function/ExampleFunction` .
           
           To get the function ARN, use the following syntax:
           
           `!GetAtt *Function_Logical_ID* .FunctionMetadata.FunctionARN`
    """
    __args__ = dict()
    __args__['functionArn'] = function_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('aws-native:cloudfront:getFunction', __args__, opts=opts, typ=GetFunctionResult)
    return __ret__.apply(lambda __response__: GetFunctionResult(
        function_arn=pulumi.get(__response__, 'function_arn'),
        function_code=pulumi.get(__response__, 'function_code'),
        function_config=pulumi.get(__response__, 'function_config'),
        function_metadata=pulumi.get(__response__, 'function_metadata'),
        name=pulumi.get(__response__, 'name'),
        stage=pulumi.get(__response__, 'stage')))
