# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 data_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 federation_mode: Optional[pulumi.Input['EnvironmentFederationMode']] = None,
                 federation_parameters: Optional[pulumi.Input['EnvironmentFederationParametersArgs']] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 superuser_parameters: Optional[pulumi.Input['EnvironmentSuperuserParametersArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentTagArgs']]]] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_bundles: ARNs of FinSpace Data Bundles to install
        :param pulumi.Input[str] description: Description of the Environment
        :param pulumi.Input['EnvironmentFederationMode'] federation_mode: Federation mode used with the Environment
        :param pulumi.Input[str] kms_key_id: KMS key used to encrypt customer data within FinSpace Environment infrastructure
        :param pulumi.Input[str] name: Name of the Environment
        :param pulumi.Input[Sequence[pulumi.Input['EnvironmentTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if data_bundles is not None:
            pulumi.set(__self__, "data_bundles", data_bundles)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if federation_mode is not None:
            pulumi.set(__self__, "federation_mode", federation_mode)
        if federation_parameters is not None:
            pulumi.set(__self__, "federation_parameters", federation_parameters)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if superuser_parameters is not None:
            pulumi.set(__self__, "superuser_parameters", superuser_parameters)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataBundles")
    def data_bundles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ARNs of FinSpace Data Bundles to install
        """
        return pulumi.get(self, "data_bundles")

    @data_bundles.setter
    def data_bundles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "data_bundles", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the Environment
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="federationMode")
    def federation_mode(self) -> Optional[pulumi.Input['EnvironmentFederationMode']]:
        """
        Federation mode used with the Environment
        """
        return pulumi.get(self, "federation_mode")

    @federation_mode.setter
    def federation_mode(self, value: Optional[pulumi.Input['EnvironmentFederationMode']]):
        pulumi.set(self, "federation_mode", value)

    @property
    @pulumi.getter(name="federationParameters")
    def federation_parameters(self) -> Optional[pulumi.Input['EnvironmentFederationParametersArgs']]:
        return pulumi.get(self, "federation_parameters")

    @federation_parameters.setter
    def federation_parameters(self, value: Optional[pulumi.Input['EnvironmentFederationParametersArgs']]):
        pulumi.set(self, "federation_parameters", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        KMS key used to encrypt customer data within FinSpace Environment infrastructure
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Environment
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="superuserParameters")
    def superuser_parameters(self) -> Optional[pulumi.Input['EnvironmentSuperuserParametersArgs']]:
        return pulumi.get(self, "superuser_parameters")

    @superuser_parameters.setter
    def superuser_parameters(self, value: Optional[pulumi.Input['EnvironmentSuperuserParametersArgs']]):
        pulumi.set(self, "superuser_parameters", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EnvironmentTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 federation_mode: Optional[pulumi.Input['EnvironmentFederationMode']] = None,
                 federation_parameters: Optional[pulumi.Input[pulumi.InputType['EnvironmentFederationParametersArgs']]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 superuser_parameters: Optional[pulumi.Input[pulumi.InputType['EnvironmentSuperuserParametersArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentTagArgs']]]]] = None,
                 __props__=None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] data_bundles: ARNs of FinSpace Data Bundles to install
        :param pulumi.Input[str] description: Description of the Environment
        :param pulumi.Input['EnvironmentFederationMode'] federation_mode: Federation mode used with the Environment
        :param pulumi.Input[str] kms_key_id: KMS key used to encrypt customer data within FinSpace Environment infrastructure
        :param pulumi.Input[str] name: Name of the Environment
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EnvironmentArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An example resource schema demonstrating some basic constructs and validation rules.

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_bundles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 federation_mode: Optional[pulumi.Input['EnvironmentFederationMode']] = None,
                 federation_parameters: Optional[pulumi.Input[pulumi.InputType['EnvironmentFederationParametersArgs']]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 superuser_parameters: Optional[pulumi.Input[pulumi.InputType['EnvironmentSuperuserParametersArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EnvironmentTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["data_bundles"] = data_bundles
            __props__.__dict__["description"] = description
            __props__.__dict__["federation_mode"] = federation_mode
            __props__.__dict__["federation_parameters"] = federation_parameters
            __props__.__dict__["kms_key_id"] = kms_key_id
            __props__.__dict__["name"] = name
            __props__.__dict__["superuser_parameters"] = superuser_parameters
            __props__.__dict__["tags"] = tags
            __props__.__dict__["aws_account_id"] = None
            __props__.__dict__["dedicated_service_account_id"] = None
            __props__.__dict__["environment_arn"] = None
            __props__.__dict__["environment_id"] = None
            __props__.__dict__["environment_url"] = None
            __props__.__dict__["sage_maker_studio_domain_url"] = None
            __props__.__dict__["status"] = None
        super(Environment, __self__).__init__(
            'aws-native:finspace:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["data_bundles"] = None
        __props__.__dict__["dedicated_service_account_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["environment_arn"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["environment_url"] = None
        __props__.__dict__["federation_mode"] = None
        __props__.__dict__["federation_parameters"] = None
        __props__.__dict__["kms_key_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["sage_maker_studio_domain_url"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["superuser_parameters"] = None
        __props__.__dict__["tags"] = None
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID associated with the Environment
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="dataBundles")
    def data_bundles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        ARNs of FinSpace Data Bundles to install
        """
        return pulumi.get(self, "data_bundles")

    @property
    @pulumi.getter(name="dedicatedServiceAccountId")
    def dedicated_service_account_id(self) -> pulumi.Output[str]:
        """
        ID for FinSpace created account used to store Environment artifacts
        """
        return pulumi.get(self, "dedicated_service_account_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the Environment
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentArn")
    def environment_arn(self) -> pulumi.Output[str]:
        """
        ARN of the Environment
        """
        return pulumi.get(self, "environment_arn")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        """
        Unique identifier for representing FinSpace Environment
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="environmentUrl")
    def environment_url(self) -> pulumi.Output[str]:
        """
        URL used to login to the Environment
        """
        return pulumi.get(self, "environment_url")

    @property
    @pulumi.getter(name="federationMode")
    def federation_mode(self) -> pulumi.Output[Optional['EnvironmentFederationMode']]:
        """
        Federation mode used with the Environment
        """
        return pulumi.get(self, "federation_mode")

    @property
    @pulumi.getter(name="federationParameters")
    def federation_parameters(self) -> pulumi.Output[Optional['outputs.EnvironmentFederationParameters']]:
        return pulumi.get(self, "federation_parameters")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        KMS key used to encrypt customer data within FinSpace Environment infrastructure
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Environment
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sageMakerStudioDomainUrl")
    def sage_maker_studio_domain_url(self) -> pulumi.Output[str]:
        """
        SageMaker Studio Domain URL associated with the Environment
        """
        return pulumi.get(self, "sage_maker_studio_domain_url")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['EnvironmentStatus']:
        """
        State of the Environment
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="superuserParameters")
    def superuser_parameters(self) -> pulumi.Output[Optional['outputs.EnvironmentSuperuserParameters']]:
        return pulumi.get(self, "superuser_parameters")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.EnvironmentTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")
