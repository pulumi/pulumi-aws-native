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

__all__ = ['ComponentVersionArgs', 'ComponentVersion']

@pulumi.input_type
class ComponentVersionArgs:
    def __init__(__self__, *,
                 inline_recipe: Optional[pulumi.Input[builtins.str]] = None,
                 lambda_function: Optional[pulumi.Input['ComponentVersionLambdaFunctionRecipeSourceArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ComponentVersion resource.
        :param pulumi.Input[builtins.str] inline_recipe: The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
               
               You must specify either `InlineRecipe` or `LambdaFunction` .
        :param pulumi.Input['ComponentVersionLambdaFunctionRecipeSourceArgs'] lambda_function: The parameters to create a component from a Lambda function.
               
               You must specify either `InlineRecipe` or `LambdaFunction` .
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
               
               This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
               
               ```json
               "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
               }
               ```
        """
        if inline_recipe is not None:
            pulumi.set(__self__, "inline_recipe", inline_recipe)
        if lambda_function is not None:
            pulumi.set(__self__, "lambda_function", lambda_function)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="inlineRecipe")
    def inline_recipe(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.

        You must specify either `InlineRecipe` or `LambdaFunction` .
        """
        return pulumi.get(self, "inline_recipe")

    @inline_recipe.setter
    def inline_recipe(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "inline_recipe", value)

    @property
    @pulumi.getter(name="lambdaFunction")
    def lambda_function(self) -> Optional[pulumi.Input['ComponentVersionLambdaFunctionRecipeSourceArgs']]:
        """
        The parameters to create a component from a Lambda function.

        You must specify either `InlineRecipe` or `LambdaFunction` .
        """
        return pulumi.get(self, "lambda_function")

    @lambda_function.setter
    def lambda_function(self, value: Optional[pulumi.Input['ComponentVersionLambdaFunctionRecipeSourceArgs']]):
        pulumi.set(self, "lambda_function", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .

        This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.

        ```json
        "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
        }
        ```
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:greengrassv2:ComponentVersion")
class ComponentVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inline_recipe: Optional[pulumi.Input[builtins.str]] = None,
                 lambda_function: Optional[pulumi.Input[Union['ComponentVersionLambdaFunctionRecipeSourceArgs', 'ComponentVersionLambdaFunctionRecipeSourceArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Resource for Greengrass component version.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] inline_recipe: The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
               
               You must specify either `InlineRecipe` or `LambdaFunction` .
        :param pulumi.Input[Union['ComponentVersionLambdaFunctionRecipeSourceArgs', 'ComponentVersionLambdaFunctionRecipeSourceArgsDict']] lambda_function: The parameters to create a component from a Lambda function.
               
               You must specify either `InlineRecipe` or `LambdaFunction` .
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
               
               This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
               
               ```json
               "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
               }
               ```
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ComponentVersionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for Greengrass component version.

        :param str resource_name: The name of the resource.
        :param ComponentVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ComponentVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inline_recipe: Optional[pulumi.Input[builtins.str]] = None,
                 lambda_function: Optional[pulumi.Input[Union['ComponentVersionLambdaFunctionRecipeSourceArgs', 'ComponentVersionLambdaFunctionRecipeSourceArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ComponentVersionArgs.__new__(ComponentVersionArgs)

            __props__.__dict__["inline_recipe"] = inline_recipe
            __props__.__dict__["lambda_function"] = lambda_function
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["component_name"] = None
            __props__.__dict__["component_version"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["inlineRecipe", "lambdaFunction"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ComponentVersion, __self__).__init__(
            'aws-native:greengrassv2:ComponentVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ComponentVersion':
        """
        Get an existing ComponentVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ComponentVersionArgs.__new__(ComponentVersionArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["component_name"] = None
        __props__.__dict__["component_version"] = None
        __props__.__dict__["inline_recipe"] = None
        __props__.__dict__["lambda_function"] = None
        __props__.__dict__["tags"] = None
        return ComponentVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the component version.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="componentName")
    def component_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the component.
        """
        return pulumi.get(self, "component_name")

    @property
    @pulumi.getter(name="componentVersion")
    def component_version(self) -> pulumi.Output[builtins.str]:
        """
        The version of the component.
        """
        return pulumi.get(self, "component_version")

    @property
    @pulumi.getter(name="inlineRecipe")
    def inline_recipe(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The recipe to use to create the component. The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.

        You must specify either `InlineRecipe` or `LambdaFunction` .
        """
        return pulumi.get(self, "inline_recipe")

    @property
    @pulumi.getter(name="lambdaFunction")
    def lambda_function(self) -> pulumi.Output[Optional['outputs.ComponentVersionLambdaFunctionRecipeSource']]:
        """
        The parameters to create a component from a Lambda function.

        You must specify either `InlineRecipe` or `LambdaFunction` .
        """
        return pulumi.get(self, "lambda_function")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Application-specific metadata to attach to the component version. You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .

        This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.

        ```json
        "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
        }
        ```
        """
        return pulumi.get(self, "tags")

