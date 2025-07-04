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

__all__ = ['ResponseHeadersPolicyArgs', 'ResponseHeadersPolicy']

@pulumi.input_type
class ResponseHeadersPolicyArgs:
    def __init__(__self__, *,
                 response_headers_policy_config: pulumi.Input['ResponseHeadersPolicyConfigArgs']):
        """
        The set of arguments for constructing a ResponseHeadersPolicy resource.
        :param pulumi.Input['ResponseHeadersPolicyConfigArgs'] response_headers_policy_config: A response headers policy configuration.
        """
        pulumi.set(__self__, "response_headers_policy_config", response_headers_policy_config)

    @property
    @pulumi.getter(name="responseHeadersPolicyConfig")
    def response_headers_policy_config(self) -> pulumi.Input['ResponseHeadersPolicyConfigArgs']:
        """
        A response headers policy configuration.
        """
        return pulumi.get(self, "response_headers_policy_config")

    @response_headers_policy_config.setter
    def response_headers_policy_config(self, value: pulumi.Input['ResponseHeadersPolicyConfigArgs']):
        pulumi.set(self, "response_headers_policy_config", value)


@pulumi.type_token("aws-native:cloudfront:ResponseHeadersPolicy")
class ResponseHeadersPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 response_headers_policy_config: Optional[pulumi.Input[Union['ResponseHeadersPolicyConfigArgs', 'ResponseHeadersPolicyConfigArgsDict']]] = None,
                 __props__=None):
        """
        A response headers policy.
         A response headers policy contains information about a set of HTTP response headers.
         After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
         For more information, see [Adding or removing HTTP headers in CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html) in the *Amazon CloudFront Developer Guide*.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ResponseHeadersPolicyConfigArgs', 'ResponseHeadersPolicyConfigArgsDict']] response_headers_policy_config: A response headers policy configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResponseHeadersPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A response headers policy.
         A response headers policy contains information about a set of HTTP response headers.
         After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
         For more information, see [Adding or removing HTTP headers in CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html) in the *Amazon CloudFront Developer Guide*.

        :param str resource_name: The name of the resource.
        :param ResponseHeadersPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResponseHeadersPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 response_headers_policy_config: Optional[pulumi.Input[Union['ResponseHeadersPolicyConfigArgs', 'ResponseHeadersPolicyConfigArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResponseHeadersPolicyArgs.__new__(ResponseHeadersPolicyArgs)

            if response_headers_policy_config is None and not opts.urn:
                raise TypeError("Missing required property 'response_headers_policy_config'")
            __props__.__dict__["response_headers_policy_config"] = response_headers_policy_config
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["last_modified_time"] = None
        super(ResponseHeadersPolicy, __self__).__init__(
            'aws-native:cloudfront:ResponseHeadersPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResponseHeadersPolicy':
        """
        Get an existing ResponseHeadersPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResponseHeadersPolicyArgs.__new__(ResponseHeadersPolicyArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["last_modified_time"] = None
        __props__.__dict__["response_headers_policy_config"] = None
        return ResponseHeadersPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique identifier for the response headers policy. For example: `57f99797-3b20-4e1b-a728-27972a74082a` .
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[builtins.str]:
        """
        The date and time when the response headers policy was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="responseHeadersPolicyConfig")
    def response_headers_policy_config(self) -> pulumi.Output['outputs.ResponseHeadersPolicyConfig']:
        """
        A response headers policy configuration.
        """
        return pulumi.get(self, "response_headers_policy_config")

