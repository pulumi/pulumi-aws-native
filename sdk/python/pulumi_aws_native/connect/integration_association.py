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

__all__ = ['IntegrationAssociationArgs', 'IntegrationAssociation']

@pulumi.input_type
class IntegrationAssociationArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 integration_arn: pulumi.Input[str],
                 integration_type: pulumi.Input['IntegrationAssociationIntegrationType']):
        """
        The set of arguments for constructing a IntegrationAssociation resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "integration_arn", integration_arn)
        pulumi.set(__self__, "integration_type", integration_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="integrationArn")
    def integration_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "integration_arn")

    @integration_arn.setter
    def integration_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_arn", value)

    @property
    @pulumi.getter(name="integrationType")
    def integration_type(self) -> pulumi.Input['IntegrationAssociationIntegrationType']:
        return pulumi.get(self, "integration_type")

    @integration_type.setter
    def integration_type(self, value: pulumi.Input['IntegrationAssociationIntegrationType']):
        pulumi.set(self, "integration_type", value)


class IntegrationAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 integration_arn: Optional[pulumi.Input[str]] = None,
                 integration_type: Optional[pulumi.Input['IntegrationAssociationIntegrationType']] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Connect::IntegrationAssociation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Connect::IntegrationAssociation

        :param str resource_name: The name of the resource.
        :param IntegrationAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 integration_arn: Optional[pulumi.Input[str]] = None,
                 integration_type: Optional[pulumi.Input['IntegrationAssociationIntegrationType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationAssociationArgs.__new__(IntegrationAssociationArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if integration_arn is None and not opts.urn:
                raise TypeError("Missing required property 'integration_arn'")
            __props__.__dict__["integration_arn"] = integration_arn
            if integration_type is None and not opts.urn:
                raise TypeError("Missing required property 'integration_type'")
            __props__.__dict__["integration_type"] = integration_type
            __props__.__dict__["integration_association_id"] = None
        super(IntegrationAssociation, __self__).__init__(
            'aws-native:connect:IntegrationAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IntegrationAssociation':
        """
        Get an existing IntegrationAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IntegrationAssociationArgs.__new__(IntegrationAssociationArgs)

        __props__.__dict__["instance_id"] = None
        __props__.__dict__["integration_arn"] = None
        __props__.__dict__["integration_association_id"] = None
        __props__.__dict__["integration_type"] = None
        return IntegrationAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="integrationArn")
    def integration_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "integration_arn")

    @property
    @pulumi.getter(name="integrationAssociationId")
    def integration_association_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "integration_association_id")

    @property
    @pulumi.getter(name="integrationType")
    def integration_type(self) -> pulumi.Output['IntegrationAssociationIntegrationType']:
        return pulumi.get(self, "integration_type")
