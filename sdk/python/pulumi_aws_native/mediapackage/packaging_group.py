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
from ._inputs import *

__all__ = ['PackagingGroupArgs', 'PackagingGroup']

@pulumi.input_type
class PackagingGroupArgs:
    def __init__(__self__, *,
                 aws_id: pulumi.Input[builtins.str],
                 authorization: Optional[pulumi.Input['PackagingGroupAuthorizationArgs']] = None,
                 egress_access_logs: Optional[pulumi.Input['PackagingGroupLogConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]] = None):
        """
        The set of arguments for constructing a PackagingGroup resource.
        :param pulumi.Input[builtins.str] aws_id: The ID of the PackagingGroup.
        :param pulumi.Input['PackagingGroupAuthorizationArgs'] authorization: CDN Authorization
        :param pulumi.Input['PackagingGroupLogConfigurationArgs'] egress_access_logs: The configuration parameters for egress access logging.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]] tags: A collection of tags associated with a resource
        """
        pulumi.set(__self__, "aws_id", aws_id)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if egress_access_logs is not None:
            pulumi.set(__self__, "egress_access_logs", egress_access_logs)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the PackagingGroup.
        """
        return pulumi.get(self, "aws_id")

    @aws_id.setter
    def aws_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "aws_id", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input['PackagingGroupAuthorizationArgs']]:
        """
        CDN Authorization
        """
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input['PackagingGroupAuthorizationArgs']]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="egressAccessLogs")
    def egress_access_logs(self) -> Optional[pulumi.Input['PackagingGroupLogConfigurationArgs']]:
        """
        The configuration parameters for egress access logging.
        """
        return pulumi.get(self, "egress_access_logs")

    @egress_access_logs.setter
    def egress_access_logs(self, value: Optional[pulumi.Input['PackagingGroupLogConfigurationArgs']]):
        pulumi.set(self, "egress_access_logs", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.CreateOnlyTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:mediapackage:PackagingGroup")
class PackagingGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Union['PackagingGroupAuthorizationArgs', 'PackagingGroupAuthorizationArgsDict']]] = None,
                 aws_id: Optional[pulumi.Input[builtins.str]] = None,
                 egress_access_logs: Optional[pulumi.Input[Union['PackagingGroupLogConfigurationArgs', 'PackagingGroupLogConfigurationArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::MediaPackage::PackagingGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['PackagingGroupAuthorizationArgs', 'PackagingGroupAuthorizationArgsDict']] authorization: CDN Authorization
        :param pulumi.Input[builtins.str] aws_id: The ID of the PackagingGroup.
        :param pulumi.Input[Union['PackagingGroupLogConfigurationArgs', 'PackagingGroupLogConfigurationArgsDict']] egress_access_logs: The configuration parameters for egress access logging.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]] tags: A collection of tags associated with a resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PackagingGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::MediaPackage::PackagingGroup

        :param str resource_name: The name of the resource.
        :param PackagingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PackagingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[Union['PackagingGroupAuthorizationArgs', 'PackagingGroupAuthorizationArgsDict']]] = None,
                 aws_id: Optional[pulumi.Input[builtins.str]] = None,
                 egress_access_logs: Optional[pulumi.Input[Union['PackagingGroupLogConfigurationArgs', 'PackagingGroupLogConfigurationArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.CreateOnlyTagArgs', '_root_inputs.CreateOnlyTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PackagingGroupArgs.__new__(PackagingGroupArgs)

            __props__.__dict__["authorization"] = authorization
            if aws_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_id'")
            __props__.__dict__["aws_id"] = aws_id
            __props__.__dict__["egress_access_logs"] = egress_access_logs
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["domain_name"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["tags[*]"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(PackagingGroup, __self__).__init__(
            'aws-native:mediapackage:PackagingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PackagingGroup':
        """
        Get an existing PackagingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PackagingGroupArgs.__new__(PackagingGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["authorization"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["domain_name"] = None
        __props__.__dict__["egress_access_logs"] = None
        __props__.__dict__["tags"] = None
        return PackagingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The ARN of the PackagingGroup.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Optional['outputs.PackagingGroupAuthorization']]:
        """
        CDN Authorization
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the PackagingGroup.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[builtins.str]:
        """
        The fully qualified domain name for Assets in the PackagingGroup.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="egressAccessLogs")
    def egress_access_logs(self) -> pulumi.Output[Optional['outputs.PackagingGroupLogConfiguration']]:
        """
        The configuration parameters for egress access logging.
        """
        return pulumi.get(self, "egress_access_logs")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.CreateOnlyTag']]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

