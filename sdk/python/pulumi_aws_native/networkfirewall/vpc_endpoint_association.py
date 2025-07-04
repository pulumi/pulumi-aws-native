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

__all__ = ['VpcEndpointAssociationArgs', 'VpcEndpointAssociation']

@pulumi.input_type
class VpcEndpointAssociationArgs:
    def __init__(__self__, *,
                 firewall_arn: pulumi.Input[builtins.str],
                 subnet_mapping: pulumi.Input['VpcEndpointAssociationSubnetMappingArgs'],
                 vpc_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a VpcEndpointAssociation resource.
        :param pulumi.Input[builtins.str] firewall_arn: The Amazon Resource Name (ARN) of the firewall.
        :param pulumi.Input['VpcEndpointAssociationSubnetMappingArgs'] subnet_mapping: The ID for a subnet that's used in an association with a firewall. This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
        :param pulumi.Input[builtins.str] vpc_id: The unique identifier of the VPC for the endpoint association.
        :param pulumi.Input[builtins.str] description: A description of the VPC endpoint association.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: The key:value pairs to associate with the resource.
        """
        pulumi.set(__self__, "firewall_arn", firewall_arn)
        pulumi.set(__self__, "subnet_mapping", subnet_mapping)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="firewallArn")
    def firewall_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the firewall.
        """
        return pulumi.get(self, "firewall_arn")

    @firewall_arn.setter
    def firewall_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "firewall_arn", value)

    @property
    @pulumi.getter(name="subnetMapping")
    def subnet_mapping(self) -> pulumi.Input['VpcEndpointAssociationSubnetMappingArgs']:
        """
        The ID for a subnet that's used in an association with a firewall. This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
        """
        return pulumi.get(self, "subnet_mapping")

    @subnet_mapping.setter
    def subnet_mapping(self, value: pulumi.Input['VpcEndpointAssociationSubnetMappingArgs']):
        pulumi.set(self, "subnet_mapping", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[builtins.str]:
        """
        The unique identifier of the VPC for the endpoint association.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the VPC endpoint association.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        The key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:networkfirewall:VpcEndpointAssociation")
class VpcEndpointAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 firewall_arn: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_mapping: Optional[pulumi.Input[Union['VpcEndpointAssociationSubnetMappingArgs', 'VpcEndpointAssociationSubnetMappingArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A description of the VPC endpoint association.
        :param pulumi.Input[builtins.str] firewall_arn: The Amazon Resource Name (ARN) of the firewall.
        :param pulumi.Input[Union['VpcEndpointAssociationSubnetMappingArgs', 'VpcEndpointAssociationSubnetMappingArgsDict']] subnet_mapping: The ID for a subnet that's used in an association with a firewall. This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: The key:value pairs to associate with the resource.
        :param pulumi.Input[builtins.str] vpc_id: The unique identifier of the VPC for the endpoint association.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation

        :param str resource_name: The name of the resource.
        :param VpcEndpointAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 firewall_arn: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_mapping: Optional[pulumi.Input[Union['VpcEndpointAssociationSubnetMappingArgs', 'VpcEndpointAssociationSubnetMappingArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 vpc_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcEndpointAssociationArgs.__new__(VpcEndpointAssociationArgs)

            __props__.__dict__["description"] = description
            if firewall_arn is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_arn'")
            __props__.__dict__["firewall_arn"] = firewall_arn
            if subnet_mapping is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_mapping'")
            __props__.__dict__["subnet_mapping"] = subnet_mapping
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["endpoint_id"] = None
            __props__.__dict__["vpc_endpoint_association_arn"] = None
            __props__.__dict__["vpc_endpoint_association_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["description", "firewallArn", "subnetMapping", "vpcId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(VpcEndpointAssociation, __self__).__init__(
            'aws-native:networkfirewall:VpcEndpointAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VpcEndpointAssociation':
        """
        Get an existing VpcEndpointAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VpcEndpointAssociationArgs.__new__(VpcEndpointAssociationArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["endpoint_id"] = None
        __props__.__dict__["firewall_arn"] = None
        __props__.__dict__["subnet_mapping"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_endpoint_association_arn"] = None
        __props__.__dict__["vpc_endpoint_association_id"] = None
        __props__.__dict__["vpc_id"] = None
        return VpcEndpointAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description of the VPC endpoint association.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="firewallArn")
    def firewall_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the firewall.
        """
        return pulumi.get(self, "firewall_arn")

    @property
    @pulumi.getter(name="subnetMapping")
    def subnet_mapping(self) -> pulumi.Output['outputs.VpcEndpointAssociationSubnetMapping']:
        """
        The ID for a subnet that's used in an association with a firewall. This is used in `CreateFirewall` , `AssociateSubnets` , and `CreateVpcEndpointAssociation` . AWS Network Firewall creates an instance of the associated firewall in each subnet that you specify, to filter traffic in the subnet's Availability Zone.
        """
        return pulumi.get(self, "subnet_mapping")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        The key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcEndpointAssociationArn")
    def vpc_endpoint_association_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of a VPC endpoint association.
        """
        return pulumi.get(self, "vpc_endpoint_association_arn")

    @property
    @pulumi.getter(name="vpcEndpointAssociationId")
    def vpc_endpoint_association_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique identifier of the VPC endpoint association.
        """
        return pulumi.get(self, "vpc_endpoint_association_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[builtins.str]:
        """
        The unique identifier of the VPC for the endpoint association.
        """
        return pulumi.get(self, "vpc_id")

