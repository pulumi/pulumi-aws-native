# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['DatasetGroupArgs', 'DatasetGroup']

@pulumi.input_type
class DatasetGroupArgs:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input['DatasetGroupDomain']] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatasetGroup resource.
        :param pulumi.Input['DatasetGroupDomain'] domain: The domain of a Domain dataset group.
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.
        :param pulumi.Input[str] name: The name for the new dataset group.
        :param pulumi.Input[str] role_arn: The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if kms_key_arn is not None:
            pulumi.set(__self__, "kms_key_arn", kms_key_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input['DatasetGroupDomain']]:
        """
        The domain of a Domain dataset group.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input['DatasetGroupDomain']]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.
        """
        return pulumi.get(self, "kms_key_arn")

    @kms_key_arn.setter
    def kms_key_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for the new dataset group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)


class DatasetGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input['DatasetGroupDomain']] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Schema for AWS::Personalize::DatasetGroup.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['DatasetGroupDomain'] domain: The domain of a Domain dataset group.
        :param pulumi.Input[str] kms_key_arn: The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.
        :param pulumi.Input[str] name: The name for the new dataset group.
        :param pulumi.Input[str] role_arn: The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatasetGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Schema for AWS::Personalize::DatasetGroup.

        :param str resource_name: The name of the resource.
        :param DatasetGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input['DatasetGroupDomain']] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatasetGroupArgs.__new__(DatasetGroupArgs)

            __props__.__dict__["domain"] = domain
            __props__.__dict__["kms_key_arn"] = kms_key_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["dataset_group_arn"] = None
        super(DatasetGroup, __self__).__init__(
            'aws-native:personalize:DatasetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatasetGroup':
        """
        Get an existing DatasetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatasetGroupArgs.__new__(DatasetGroupArgs)

        __props__.__dict__["dataset_group_arn"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["kms_key_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_arn"] = None
        return DatasetGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datasetGroupArn")
    def dataset_group_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the dataset group.
        """
        return pulumi.get(self, "dataset_group_arn")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[Optional['DatasetGroupDomain']]:
        """
        The domain of a Domain dataset group.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.
        """
        return pulumi.get(self, "kms_key_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the new dataset group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.
        """
        return pulumi.get(self, "role_arn")
