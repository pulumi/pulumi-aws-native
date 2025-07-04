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

__all__ = ['TrustStoreRevocationArgs', 'TrustStoreRevocation']

@pulumi.input_type
class TrustStoreRevocationArgs:
    def __init__(__self__, *,
                 revocation_contents: Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreRevocationRevocationContentArgs']]]] = None,
                 trust_store_arn: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a TrustStoreRevocation resource.
        :param pulumi.Input[Sequence[pulumi.Input['TrustStoreRevocationRevocationContentArgs']]] revocation_contents: The attributes required to create a trust store revocation.
        :param pulumi.Input[builtins.str] trust_store_arn: The Amazon Resource Name (ARN) of the trust store.
        """
        if revocation_contents is not None:
            pulumi.set(__self__, "revocation_contents", revocation_contents)
        if trust_store_arn is not None:
            pulumi.set(__self__, "trust_store_arn", trust_store_arn)

    @property
    @pulumi.getter(name="revocationContents")
    def revocation_contents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreRevocationRevocationContentArgs']]]]:
        """
        The attributes required to create a trust store revocation.
        """
        return pulumi.get(self, "revocation_contents")

    @revocation_contents.setter
    def revocation_contents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrustStoreRevocationRevocationContentArgs']]]]):
        pulumi.set(self, "revocation_contents", value)

    @property
    @pulumi.getter(name="trustStoreArn")
    def trust_store_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Amazon Resource Name (ARN) of the trust store.
        """
        return pulumi.get(self, "trust_store_arn")

    @trust_store_arn.setter
    def trust_store_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "trust_store_arn", value)


@pulumi.type_token("aws-native:elasticloadbalancingv2:TrustStoreRevocation")
class TrustStoreRevocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 revocation_contents: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrustStoreRevocationRevocationContentArgs', 'TrustStoreRevocationRevocationContentArgsDict']]]]] = None,
                 trust_store_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStoreRevocation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['TrustStoreRevocationRevocationContentArgs', 'TrustStoreRevocationRevocationContentArgsDict']]]] revocation_contents: The attributes required to create a trust store revocation.
        :param pulumi.Input[builtins.str] trust_store_arn: The Amazon Resource Name (ARN) of the trust store.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrustStoreRevocationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStoreRevocation

        :param str resource_name: The name of the resource.
        :param TrustStoreRevocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrustStoreRevocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 revocation_contents: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TrustStoreRevocationRevocationContentArgs', 'TrustStoreRevocationRevocationContentArgsDict']]]]] = None,
                 trust_store_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrustStoreRevocationArgs.__new__(TrustStoreRevocationArgs)

            __props__.__dict__["revocation_contents"] = revocation_contents
            __props__.__dict__["trust_store_arn"] = trust_store_arn
            __props__.__dict__["revocation_id"] = None
            __props__.__dict__["trust_store_revocations"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["revocationContents[*]", "trustStoreArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(TrustStoreRevocation, __self__).__init__(
            'aws-native:elasticloadbalancingv2:TrustStoreRevocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TrustStoreRevocation':
        """
        Get an existing TrustStoreRevocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrustStoreRevocationArgs.__new__(TrustStoreRevocationArgs)

        __props__.__dict__["revocation_contents"] = None
        __props__.__dict__["revocation_id"] = None
        __props__.__dict__["trust_store_arn"] = None
        __props__.__dict__["trust_store_revocations"] = None
        return TrustStoreRevocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="revocationContents")
    def revocation_contents(self) -> pulumi.Output[Optional[Sequence['outputs.TrustStoreRevocationRevocationContent']]]:
        """
        The attributes required to create a trust store revocation.
        """
        return pulumi.get(self, "revocation_contents")

    @property
    @pulumi.getter(name="revocationId")
    def revocation_id(self) -> pulumi.Output[builtins.int]:
        """
        The ID associated with the revocation.
        """
        return pulumi.get(self, "revocation_id")

    @property
    @pulumi.getter(name="trustStoreArn")
    def trust_store_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Amazon Resource Name (ARN) of the trust store.
        """
        return pulumi.get(self, "trust_store_arn")

    @property
    @pulumi.getter(name="trustStoreRevocations")
    def trust_store_revocations(self) -> pulumi.Output[Sequence['outputs.TrustStoreRevocation']]:
        """
        The data associated with a trust store revocation
        """
        return pulumi.get(self, "trust_store_revocations")

