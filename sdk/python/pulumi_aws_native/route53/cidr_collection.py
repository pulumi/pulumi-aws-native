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

__all__ = ['CidrCollectionArgs', 'CidrCollection']

@pulumi.input_type
class CidrCollectionArgs:
    def __init__(__self__, *,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input['CidrCollectionLocationArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a CidrCollection resource.
        :param pulumi.Input[Sequence[pulumi.Input['CidrCollectionLocationArgs']]] locations: A complex type that contains information about the list of CIDR locations.
        :param pulumi.Input[builtins.str] name: A unique name for the CIDR collection.
        """
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CidrCollectionLocationArgs']]]]:
        """
        A complex type that contains information about the list of CIDR locations.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CidrCollectionLocationArgs']]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique name for the CIDR collection.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("aws-native:route53:CidrCollection")
class CidrCollection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CidrCollectionLocationArgs', 'CidrCollectionLocationArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Route53::CidrCollection.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['CidrCollectionLocationArgs', 'CidrCollectionLocationArgsDict']]]] locations: A complex type that contains information about the list of CIDR locations.
        :param pulumi.Input[builtins.str] name: A unique name for the CIDR collection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CidrCollectionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Route53::CidrCollection.

        :param str resource_name: The name of the resource.
        :param CidrCollectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CidrCollectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CidrCollectionLocationArgs', 'CidrCollectionLocationArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CidrCollectionArgs.__new__(CidrCollectionArgs)

            __props__.__dict__["locations"] = locations
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(CidrCollection, __self__).__init__(
            'aws-native:route53:CidrCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CidrCollection':
        """
        Get an existing CidrCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CidrCollectionArgs.__new__(CidrCollectionArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["locations"] = None
        __props__.__dict__["name"] = None
        return CidrCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon resource name (ARN) to uniquely identify the AWS resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        UUID of the CIDR collection.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Optional[Sequence['outputs.CidrCollectionLocation']]]:
        """
        A complex type that contains information about the list of CIDR locations.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        A unique name for the CIDR collection.
        """
        return pulumi.get(self, "name")

