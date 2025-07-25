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

__all__ = ['SubnetRouteTableAssociationArgs', 'SubnetRouteTableAssociation']

@pulumi.input_type
class SubnetRouteTableAssociationArgs:
    def __init__(__self__, *,
                 route_table_id: pulumi.Input[builtins.str],
                 subnet_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a SubnetRouteTableAssociation resource.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table.
                The physical ID changes when the route table ID is changed.
        :param pulumi.Input[builtins.str] subnet_id: The ID of the subnet.
        """
        pulumi.set(__self__, "route_table_id", route_table_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the route table.
         The physical ID changes when the route table ID is changed.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.type_token("aws-native:ec2:SubnetRouteTableAssociation")
class SubnetRouteTableAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Associates a subnet with a route table. The subnet and route table must be in the same VPC. This association causes traffic originating from the subnet to be routed according to the routes in the route table. A route table can be associated with multiple subnets. To create a route table, see [AWS::EC2::RouteTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] route_table_id: The ID of the route table.
                The physical ID changes when the route table ID is changed.
        :param pulumi.Input[builtins.str] subnet_id: The ID of the subnet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetRouteTableAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Associates a subnet with a route table. The subnet and route table must be in the same VPC. This association causes traffic originating from the subnet to be routed according to the routes in the route table. A route table can be associated with multiple subnets. To create a route table, see [AWS::EC2::RouteTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html).

        :param str resource_name: The name of the resource.
        :param SubnetRouteTableAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetRouteTableAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetRouteTableAssociationArgs.__new__(SubnetRouteTableAssociationArgs)

            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__.__dict__["route_table_id"] = route_table_id
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["routeTableId", "subnetId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(SubnetRouteTableAssociation, __self__).__init__(
            'aws-native:ec2:SubnetRouteTableAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SubnetRouteTableAssociation':
        """
        Get an existing SubnetRouteTableAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SubnetRouteTableAssociationArgs.__new__(SubnetRouteTableAssociationArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["route_table_id"] = None
        __props__.__dict__["subnet_id"] = None
        return SubnetRouteTableAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the subnet route table association.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the route table.
         The physical ID changes when the route table ID is changed.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

