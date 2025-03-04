# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['FlowOutputArgs', 'FlowOutput']

@pulumi.input_type
class FlowOutputArgs:
    def __init__(__self__, *,
                 flow_arn: pulumi.Input[str],
                 protocol: pulumi.Input['FlowOutputProtocol'],
                 cidr_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input['FlowOutputEncryptionArgs']] = None,
                 max_latency: Optional[pulumi.Input[int]] = None,
                 media_stream_output_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['FlowOutputMediaStreamOutputConfigurationArgs']]]] = None,
                 min_latency: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_status: Optional[pulumi.Input['FlowOutputOutputStatus']] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 remote_id: Optional[pulumi.Input[str]] = None,
                 smoothing_latency: Optional[pulumi.Input[int]] = None,
                 stream_id: Optional[pulumi.Input[str]] = None,
                 vpc_interface_attachment: Optional[pulumi.Input['FlowOutputVpcInterfaceAttachmentArgs']] = None):
        """
        The set of arguments for constructing a FlowOutput resource.
        :param pulumi.Input[str] flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        :param pulumi.Input['FlowOutputProtocol'] protocol: The protocol that is used by the source or output.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_allow_list: The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        :param pulumi.Input[str] description: A description of the output.
        :param pulumi.Input[str] destination: The address where you want to send the output.
        :param pulumi.Input['FlowOutputEncryptionArgs'] encryption: The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        :param pulumi.Input[int] max_latency: The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
        :param pulumi.Input[Sequence[pulumi.Input['FlowOutputMediaStreamOutputConfigurationArgs']]] media_stream_output_configurations: The definition for each media stream that is associated with the output.
        :param pulumi.Input[int] min_latency: The minimum latency in milliseconds.
        :param pulumi.Input[str] name: The name of the output. This value must be unique within the current flow.
        :param pulumi.Input['FlowOutputOutputStatus'] output_status: An indication of whether the output should transmit data or not.
        :param pulumi.Input[int] port: The port to use when content is distributed to this output.
        :param pulumi.Input[str] remote_id: The remote ID for the Zixi-pull stream.
        :param pulumi.Input[int] smoothing_latency: The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
        :param pulumi.Input[str] stream_id: The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
        :param pulumi.Input['FlowOutputVpcInterfaceAttachmentArgs'] vpc_interface_attachment: The name of the VPC interface attachment to use for this output.
        """
        pulumi.set(__self__, "flow_arn", flow_arn)
        pulumi.set(__self__, "protocol", protocol)
        if cidr_allow_list is not None:
            pulumi.set(__self__, "cidr_allow_list", cidr_allow_list)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if max_latency is not None:
            pulumi.set(__self__, "max_latency", max_latency)
        if media_stream_output_configurations is not None:
            pulumi.set(__self__, "media_stream_output_configurations", media_stream_output_configurations)
        if min_latency is not None:
            pulumi.set(__self__, "min_latency", min_latency)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_status is not None:
            pulumi.set(__self__, "output_status", output_status)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if remote_id is not None:
            pulumi.set(__self__, "remote_id", remote_id)
        if smoothing_latency is not None:
            pulumi.set(__self__, "smoothing_latency", smoothing_latency)
        if stream_id is not None:
            pulumi.set(__self__, "stream_id", stream_id)
        if vpc_interface_attachment is not None:
            pulumi.set(__self__, "vpc_interface_attachment", vpc_interface_attachment)

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @flow_arn.setter
    def flow_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_arn", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input['FlowOutputProtocol']:
        """
        The protocol that is used by the source or output.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input['FlowOutputProtocol']):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="cidrAllowList")
    def cidr_allow_list(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        """
        return pulumi.get(self, "cidr_allow_list")

    @cidr_allow_list.setter
    def cidr_allow_list(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_allow_list", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the output.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        """
        The address where you want to send the output.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input['FlowOutputEncryptionArgs']]:
        """
        The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input['FlowOutputEncryptionArgs']]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="maxLatency")
    def max_latency(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
        """
        return pulumi.get(self, "max_latency")

    @max_latency.setter
    def max_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_latency", value)

    @property
    @pulumi.getter(name="mediaStreamOutputConfigurations")
    def media_stream_output_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FlowOutputMediaStreamOutputConfigurationArgs']]]]:
        """
        The definition for each media stream that is associated with the output.
        """
        return pulumi.get(self, "media_stream_output_configurations")

    @media_stream_output_configurations.setter
    def media_stream_output_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FlowOutputMediaStreamOutputConfigurationArgs']]]]):
        pulumi.set(self, "media_stream_output_configurations", value)

    @property
    @pulumi.getter(name="minLatency")
    def min_latency(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum latency in milliseconds.
        """
        return pulumi.get(self, "min_latency")

    @min_latency.setter
    def min_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_latency", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the output. This value must be unique within the current flow.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputStatus")
    def output_status(self) -> Optional[pulumi.Input['FlowOutputOutputStatus']]:
        """
        An indication of whether the output should transmit data or not.
        """
        return pulumi.get(self, "output_status")

    @output_status.setter
    def output_status(self, value: Optional[pulumi.Input['FlowOutputOutputStatus']]):
        pulumi.set(self, "output_status", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port to use when content is distributed to this output.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="remoteId")
    def remote_id(self) -> Optional[pulumi.Input[str]]:
        """
        The remote ID for the Zixi-pull stream.
        """
        return pulumi.get(self, "remote_id")

    @remote_id.setter
    def remote_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_id", value)

    @property
    @pulumi.getter(name="smoothingLatency")
    def smoothing_latency(self) -> Optional[pulumi.Input[int]]:
        """
        The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
        """
        return pulumi.get(self, "smoothing_latency")

    @smoothing_latency.setter
    def smoothing_latency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smoothing_latency", value)

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> Optional[pulumi.Input[str]]:
        """
        The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
        """
        return pulumi.get(self, "stream_id")

    @stream_id.setter
    def stream_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_id", value)

    @property
    @pulumi.getter(name="vpcInterfaceAttachment")
    def vpc_interface_attachment(self) -> Optional[pulumi.Input['FlowOutputVpcInterfaceAttachmentArgs']]:
        """
        The name of the VPC interface attachment to use for this output.
        """
        return pulumi.get(self, "vpc_interface_attachment")

    @vpc_interface_attachment.setter
    def vpc_interface_attachment(self, value: Optional[pulumi.Input['FlowOutputVpcInterfaceAttachmentArgs']]):
        pulumi.set(self, "vpc_interface_attachment", value)


class FlowOutput(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[Union['FlowOutputEncryptionArgs', 'FlowOutputEncryptionArgsDict']]] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 max_latency: Optional[pulumi.Input[int]] = None,
                 media_stream_output_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FlowOutputMediaStreamOutputConfigurationArgs', 'FlowOutputMediaStreamOutputConfigurationArgsDict']]]]] = None,
                 min_latency: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_status: Optional[pulumi.Input['FlowOutputOutputStatus']] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input['FlowOutputProtocol']] = None,
                 remote_id: Optional[pulumi.Input[str]] = None,
                 smoothing_latency: Optional[pulumi.Input[int]] = None,
                 stream_id: Optional[pulumi.Input[str]] = None,
                 vpc_interface_attachment: Optional[pulumi.Input[Union['FlowOutputVpcInterfaceAttachmentArgs', 'FlowOutputVpcInterfaceAttachmentArgsDict']]] = None,
                 __props__=None):
        """
        Resource schema for AWS::MediaConnect::FlowOutput

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_allow_list: The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        :param pulumi.Input[str] description: A description of the output.
        :param pulumi.Input[str] destination: The address where you want to send the output.
        :param pulumi.Input[Union['FlowOutputEncryptionArgs', 'FlowOutputEncryptionArgsDict']] encryption: The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        :param pulumi.Input[str] flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        :param pulumi.Input[int] max_latency: The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
        :param pulumi.Input[Sequence[pulumi.Input[Union['FlowOutputMediaStreamOutputConfigurationArgs', 'FlowOutputMediaStreamOutputConfigurationArgsDict']]]] media_stream_output_configurations: The definition for each media stream that is associated with the output.
        :param pulumi.Input[int] min_latency: The minimum latency in milliseconds.
        :param pulumi.Input[str] name: The name of the output. This value must be unique within the current flow.
        :param pulumi.Input['FlowOutputOutputStatus'] output_status: An indication of whether the output should transmit data or not.
        :param pulumi.Input[int] port: The port to use when content is distributed to this output.
        :param pulumi.Input['FlowOutputProtocol'] protocol: The protocol that is used by the source or output.
        :param pulumi.Input[str] remote_id: The remote ID for the Zixi-pull stream.
        :param pulumi.Input[int] smoothing_latency: The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
        :param pulumi.Input[str] stream_id: The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
        :param pulumi.Input[Union['FlowOutputVpcInterfaceAttachmentArgs', 'FlowOutputVpcInterfaceAttachmentArgsDict']] vpc_interface_attachment: The name of the VPC interface attachment to use for this output.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowOutputArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::MediaConnect::FlowOutput

        :param str resource_name: The name of the resource.
        :param FlowOutputArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowOutputArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_allow_list: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[Union['FlowOutputEncryptionArgs', 'FlowOutputEncryptionArgsDict']]] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 max_latency: Optional[pulumi.Input[int]] = None,
                 media_stream_output_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FlowOutputMediaStreamOutputConfigurationArgs', 'FlowOutputMediaStreamOutputConfigurationArgsDict']]]]] = None,
                 min_latency: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_status: Optional[pulumi.Input['FlowOutputOutputStatus']] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input['FlowOutputProtocol']] = None,
                 remote_id: Optional[pulumi.Input[str]] = None,
                 smoothing_latency: Optional[pulumi.Input[int]] = None,
                 stream_id: Optional[pulumi.Input[str]] = None,
                 vpc_interface_attachment: Optional[pulumi.Input[Union['FlowOutputVpcInterfaceAttachmentArgs', 'FlowOutputVpcInterfaceAttachmentArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlowOutputArgs.__new__(FlowOutputArgs)

            __props__.__dict__["cidr_allow_list"] = cidr_allow_list
            __props__.__dict__["description"] = description
            __props__.__dict__["destination"] = destination
            __props__.__dict__["encryption"] = encryption
            if flow_arn is None and not opts.urn:
                raise TypeError("Missing required property 'flow_arn'")
            __props__.__dict__["flow_arn"] = flow_arn
            __props__.__dict__["max_latency"] = max_latency
            __props__.__dict__["media_stream_output_configurations"] = media_stream_output_configurations
            __props__.__dict__["min_latency"] = min_latency
            __props__.__dict__["name"] = name
            __props__.__dict__["output_status"] = output_status
            __props__.__dict__["port"] = port
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["remote_id"] = remote_id
            __props__.__dict__["smoothing_latency"] = smoothing_latency
            __props__.__dict__["stream_id"] = stream_id
            __props__.__dict__["vpc_interface_attachment"] = vpc_interface_attachment
            __props__.__dict__["output_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(FlowOutput, __self__).__init__(
            'aws-native:mediaconnect:FlowOutput',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FlowOutput':
        """
        Get an existing FlowOutput resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FlowOutputArgs.__new__(FlowOutputArgs)

        __props__.__dict__["cidr_allow_list"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["destination"] = None
        __props__.__dict__["encryption"] = None
        __props__.__dict__["flow_arn"] = None
        __props__.__dict__["max_latency"] = None
        __props__.__dict__["media_stream_output_configurations"] = None
        __props__.__dict__["min_latency"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["output_arn"] = None
        __props__.__dict__["output_status"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["protocol"] = None
        __props__.__dict__["remote_id"] = None
        __props__.__dict__["smoothing_latency"] = None
        __props__.__dict__["stream_id"] = None
        __props__.__dict__["vpc_interface_attachment"] = None
        return FlowOutput(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cidrAllowList")
    def cidr_allow_list(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        """
        return pulumi.get(self, "cidr_allow_list")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the output.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[Optional[str]]:
        """
        The address where you want to send the output.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[Optional['outputs.FlowOutputEncryption']]:
        """
        The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter(name="maxLatency")
    def max_latency(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
        """
        return pulumi.get(self, "max_latency")

    @property
    @pulumi.getter(name="mediaStreamOutputConfigurations")
    def media_stream_output_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.FlowOutputMediaStreamOutputConfiguration']]]:
        """
        The definition for each media stream that is associated with the output.
        """
        return pulumi.get(self, "media_stream_output_configurations")

    @property
    @pulumi.getter(name="minLatency")
    def min_latency(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum latency in milliseconds.
        """
        return pulumi.get(self, "min_latency")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the output. This value must be unique within the current flow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputArn")
    def output_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the output.
        """
        return pulumi.get(self, "output_arn")

    @property
    @pulumi.getter(name="outputStatus")
    def output_status(self) -> pulumi.Output[Optional['FlowOutputOutputStatus']]:
        """
        An indication of whether the output should transmit data or not.
        """
        return pulumi.get(self, "output_status")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port to use when content is distributed to this output.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output['FlowOutputProtocol']:
        """
        The protocol that is used by the source or output.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="remoteId")
    def remote_id(self) -> pulumi.Output[Optional[str]]:
        """
        The remote ID for the Zixi-pull stream.
        """
        return pulumi.get(self, "remote_id")

    @property
    @pulumi.getter(name="smoothingLatency")
    def smoothing_latency(self) -> pulumi.Output[Optional[int]]:
        """
        The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
        """
        return pulumi.get(self, "smoothing_latency")

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> pulumi.Output[Optional[str]]:
        """
        The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
        """
        return pulumi.get(self, "stream_id")

    @property
    @pulumi.getter(name="vpcInterfaceAttachment")
    def vpc_interface_attachment(self) -> pulumi.Output[Optional['outputs.FlowOutputVpcInterfaceAttachment']]:
        """
        The name of the VPC interface attachment to use for this output.
        """
        return pulumi.get(self, "vpc_interface_attachment")

