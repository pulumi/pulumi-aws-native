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

__all__ = ['MultiRegionAccessPointPolicyArgs', 'MultiRegionAccessPointPolicy']

@pulumi.input_type
class MultiRegionAccessPointPolicyArgs:
    def __init__(__self__, *,
                 mrap_name: pulumi.Input[builtins.str],
                 policy: Any):
        """
        The set of arguments for constructing a MultiRegionAccessPointPolicy resource.
        :param pulumi.Input[builtins.str] mrap_name: The name of the Multi Region Access Point to apply policy
        :param Any policy: Policy document to apply to a Multi Region Access Point
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "mrap_name", mrap_name)
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="mrapName")
    def mrap_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Multi Region Access Point to apply policy
        """
        return pulumi.get(self, "mrap_name")

    @mrap_name.setter
    def mrap_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "mrap_name", value)

    @property
    @pulumi.getter
    def policy(self) -> Any:
        """
        Policy document to apply to a Multi Region Access Point

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Any):
        pulumi.set(self, "policy", value)


@pulumi.type_token("aws-native:s3:MultiRegionAccessPointPolicy")
class MultiRegionAccessPointPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mrap_name: Optional[pulumi.Input[builtins.str]] = None,
                 policy: Optional[Any] = None,
                 __props__=None):
        """
        The policy to be attached to a Multi Region Access Point

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] mrap_name: The name of the Multi Region Access Point to apply policy
        :param Any policy: Policy document to apply to a Multi Region Access Point
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MultiRegionAccessPointPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The policy to be attached to a Multi Region Access Point

        :param str resource_name: The name of the resource.
        :param MultiRegionAccessPointPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MultiRegionAccessPointPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mrap_name: Optional[pulumi.Input[builtins.str]] = None,
                 policy: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MultiRegionAccessPointPolicyArgs.__new__(MultiRegionAccessPointPolicyArgs)

            if mrap_name is None and not opts.urn:
                raise TypeError("Missing required property 'mrap_name'")
            __props__.__dict__["mrap_name"] = mrap_name
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["policy_status"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["mrapName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(MultiRegionAccessPointPolicy, __self__).__init__(
            'aws-native:s3:MultiRegionAccessPointPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MultiRegionAccessPointPolicy':
        """
        Get an existing MultiRegionAccessPointPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MultiRegionAccessPointPolicyArgs.__new__(MultiRegionAccessPointPolicyArgs)

        __props__.__dict__["mrap_name"] = None
        __props__.__dict__["policy"] = None
        __props__.__dict__["policy_status"] = None
        return MultiRegionAccessPointPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="mrapName")
    def mrap_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Multi Region Access Point to apply policy
        """
        return pulumi.get(self, "mrap_name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Any]:
        """
        Policy document to apply to a Multi Region Access Point

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3::MultiRegionAccessPointPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="policyStatus")
    def policy_status(self) -> pulumi.Output['outputs.PolicyStatusProperties']:
        """
        The Policy Status associated with this Multi Region Access Point
        """
        return pulumi.get(self, "policy_status")

