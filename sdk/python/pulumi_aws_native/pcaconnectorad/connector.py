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
from ._inputs import *

__all__ = ['ConnectorArgs', 'Connector']

@pulumi.input_type
class ConnectorArgs:
    def __init__(__self__, *,
                 certificate_authority_arn: pulumi.Input[builtins.str],
                 directory_id: pulumi.Input[builtins.str],
                 vpc_information: pulumi.Input['ConnectorVpcInformationArgs'],
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Connector resource.
        :param pulumi.Input[builtins.str] certificate_authority_arn: The Amazon Resource Name (ARN) of the certificate authority being used.
        :param pulumi.Input[builtins.str] directory_id: The identifier of the Active Directory.
        :param pulumi.Input['ConnectorVpcInformationArgs'] vpc_information: Information of the VPC and security group(s) used with the connector.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Metadata assigned to a connector consisting of a key-value pair.
        """
        pulumi.set(__self__, "certificate_authority_arn", certificate_authority_arn)
        pulumi.set(__self__, "directory_id", directory_id)
        pulumi.set(__self__, "vpc_information", vpc_information)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="certificateAuthorityArn")
    def certificate_authority_arn(self) -> pulumi.Input[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the certificate authority being used.
        """
        return pulumi.get(self, "certificate_authority_arn")

    @certificate_authority_arn.setter
    def certificate_authority_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "certificate_authority_arn", value)

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> pulumi.Input[builtins.str]:
        """
        The identifier of the Active Directory.
        """
        return pulumi.get(self, "directory_id")

    @directory_id.setter
    def directory_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "directory_id", value)

    @property
    @pulumi.getter(name="vpcInformation")
    def vpc_information(self) -> pulumi.Input['ConnectorVpcInformationArgs']:
        """
        Information of the VPC and security group(s) used with the connector.
        """
        return pulumi.get(self, "vpc_information")

    @vpc_information.setter
    def vpc_information(self, value: pulumi.Input['ConnectorVpcInformationArgs']):
        pulumi.set(self, "vpc_information", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Metadata assigned to a connector consisting of a key-value pair.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("aws-native:pcaconnectorad:Connector")
class Connector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_authority_arn: Optional[pulumi.Input[builtins.str]] = None,
                 directory_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vpc_information: Optional[pulumi.Input[Union['ConnectorVpcInformationArgs', 'ConnectorVpcInformationArgsDict']]] = None,
                 __props__=None):
        """
        Represents a Connector that connects AWS PrivateCA and your directory

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] certificate_authority_arn: The Amazon Resource Name (ARN) of the certificate authority being used.
        :param pulumi.Input[builtins.str] directory_id: The identifier of the Active Directory.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] tags: Metadata assigned to a connector consisting of a key-value pair.
        :param pulumi.Input[Union['ConnectorVpcInformationArgs', 'ConnectorVpcInformationArgsDict']] vpc_information: Information of the VPC and security group(s) used with the connector.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a Connector that connects AWS PrivateCA and your directory

        :param str resource_name: The name of the resource.
        :param ConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_authority_arn: Optional[pulumi.Input[builtins.str]] = None,
                 directory_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 vpc_information: Optional[pulumi.Input[Union['ConnectorVpcInformationArgs', 'ConnectorVpcInformationArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectorArgs.__new__(ConnectorArgs)

            if certificate_authority_arn is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_authority_arn'")
            __props__.__dict__["certificate_authority_arn"] = certificate_authority_arn
            if directory_id is None and not opts.urn:
                raise TypeError("Missing required property 'directory_id'")
            __props__.__dict__["directory_id"] = directory_id
            __props__.__dict__["tags"] = tags
            if vpc_information is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_information'")
            __props__.__dict__["vpc_information"] = vpc_information
            __props__.__dict__["connector_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["certificateAuthorityArn", "directoryId", "vpcInformation"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Connector, __self__).__init__(
            'aws-native:pcaconnectorad:Connector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connector':
        """
        Get an existing Connector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectorArgs.__new__(ConnectorArgs)

        __props__.__dict__["certificate_authority_arn"] = None
        __props__.__dict__["connector_arn"] = None
        __props__.__dict__["directory_id"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_information"] = None
        return Connector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateAuthorityArn")
    def certificate_authority_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the certificate authority being used.
        """
        return pulumi.get(self, "certificate_authority_arn")

    @property
    @pulumi.getter(name="connectorArn")
    def connector_arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html) .
        """
        return pulumi.get(self, "connector_arn")

    @property
    @pulumi.getter(name="directoryId")
    def directory_id(self) -> pulumi.Output[builtins.str]:
        """
        The identifier of the Active Directory.
        """
        return pulumi.get(self, "directory_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        Metadata assigned to a connector consisting of a key-value pair.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcInformation")
    def vpc_information(self) -> pulumi.Output['outputs.ConnectorVpcInformation']:
        """
        Information of the VPC and security group(s) used with the connector.
        """
        return pulumi.get(self, "vpc_information")

