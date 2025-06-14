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

__all__ = ['AccessPointPolicyArgs', 'AccessPointPolicy']

@pulumi.input_type
class AccessPointPolicyArgs:
    def __init__(__self__, *,
                 object_lambda_access_point: pulumi.Input[builtins.str],
                 policy_document: Any):
        """
        The set of arguments for constructing a AccessPointPolicy resource.
        :param pulumi.Input[builtins.str] object_lambda_access_point: The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        :param Any policy_document: A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3ObjectLambda::AccessPointPolicy` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "object_lambda_access_point", object_lambda_access_point)
        pulumi.set(__self__, "policy_document", policy_document)

    @property
    @pulumi.getter(name="objectLambdaAccessPoint")
    def object_lambda_access_point(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        """
        return pulumi.get(self, "object_lambda_access_point")

    @object_lambda_access_point.setter
    def object_lambda_access_point(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "object_lambda_access_point", value)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> Any:
        """
        A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3ObjectLambda::AccessPointPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: Any):
        pulumi.set(self, "policy_document", value)


@pulumi.type_token("aws-native:s3objectlambda:AccessPointPolicy")
class AccessPointPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 object_lambda_access_point: Optional[pulumi.Input[builtins.str]] = None,
                 policy_document: Optional[Any] = None,
                 __props__=None):
        """
        AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] object_lambda_access_point: The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        :param Any policy_document: A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 
               
               Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3ObjectLambda::AccessPointPolicy` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPointPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda

        :param str resource_name: The name of the resource.
        :param AccessPointPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPointPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 object_lambda_access_point: Optional[pulumi.Input[builtins.str]] = None,
                 policy_document: Optional[Any] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPointPolicyArgs.__new__(AccessPointPolicyArgs)

            if object_lambda_access_point is None and not opts.urn:
                raise TypeError("Missing required property 'object_lambda_access_point'")
            __props__.__dict__["object_lambda_access_point"] = object_lambda_access_point
            if policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'policy_document'")
            __props__.__dict__["policy_document"] = policy_document
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["objectLambdaAccessPoint"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(AccessPointPolicy, __self__).__init__(
            'aws-native:s3objectlambda:AccessPointPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AccessPointPolicy':
        """
        Get an existing AccessPointPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AccessPointPolicyArgs.__new__(AccessPointPolicyArgs)

        __props__.__dict__["object_lambda_access_point"] = None
        __props__.__dict__["policy_document"] = None
        return AccessPointPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="objectLambdaAccessPoint")
    def object_lambda_access_point(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
        """
        return pulumi.get(self, "object_lambda_access_point")

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Output[Any]:
        """
        A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. 

        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3ObjectLambda::AccessPointPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "policy_document")

