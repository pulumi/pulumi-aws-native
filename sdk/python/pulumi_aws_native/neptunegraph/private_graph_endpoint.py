# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PrivateGraphEndpointArgs', 'PrivateGraphEndpoint']

@pulumi.input_type
class PrivateGraphEndpointArgs:
    def __init__(__self__, *,
                 graph_identifier: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a PrivateGraphEndpoint resource.
        :param pulumi.Input[str] graph_identifier: The auto-generated Graph Id assigned by the service.
        :param pulumi.Input[str] vpc_id: The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        pulumi.set(__self__, "graph_identifier", graph_identifier)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="graphIdentifier")
    def graph_identifier(self) -> pulumi.Input[str]:
        """
        The auto-generated Graph Id assigned by the service.
        """
        return pulumi.get(self, "graph_identifier")

    @graph_identifier.setter
    def graph_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "graph_identifier", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)


class PrivateGraphEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 graph_identifier: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] graph_identifier: The auto-generated Graph Id assigned by the service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        :param pulumi.Input[str] vpc_id: The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateGraphEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.

        :param str resource_name: The name of the resource.
        :param PrivateGraphEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateGraphEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 graph_identifier: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateGraphEndpointArgs.__new__(PrivateGraphEndpointArgs)

            if graph_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'graph_identifier'")
            __props__.__dict__["graph_identifier"] = graph_identifier
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["subnet_ids"] = subnet_ids
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["private_graph_endpoint_identifier"] = None
            __props__.__dict__["vpc_endpoint_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["graph_identifier", "security_group_ids[*]", "subnet_ids[*]", "vpc_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(PrivateGraphEndpoint, __self__).__init__(
            'aws-native:neptunegraph:PrivateGraphEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateGraphEndpoint':
        """
        Get an existing PrivateGraphEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateGraphEndpointArgs.__new__(PrivateGraphEndpointArgs)

        __props__.__dict__["graph_identifier"] = None
        __props__.__dict__["private_graph_endpoint_identifier"] = None
        __props__.__dict__["security_group_ids"] = None
        __props__.__dict__["subnet_ids"] = None
        __props__.__dict__["vpc_endpoint_id"] = None
        __props__.__dict__["vpc_id"] = None
        return PrivateGraphEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="graphIdentifier")
    def graph_identifier(self) -> pulumi.Output[str]:
        """
        The auto-generated Graph Id assigned by the service.
        """
        return pulumi.get(self, "graph_identifier")

    @property
    @pulumi.getter(name="privateGraphEndpointIdentifier")
    def private_graph_endpoint_identifier(self) -> pulumi.Output[str]:
        """
        PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.

         For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
        """
        return pulumi.get(self, "private_graph_endpoint_identifier")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Output[str]:
        """
        VPC endpoint that provides a private connection between the Graph and specified VPC.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
        """
        return pulumi.get(self, "vpc_id")
