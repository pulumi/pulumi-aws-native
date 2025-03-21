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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = ['LocationObjectStorageArgs', 'LocationObjectStorage']

@pulumi.input_type
class LocationObjectStorageArgs:
    def __init__(__self__, *,
                 agent_arns: pulumi.Input[Sequence[pulumi.Input[str]]],
                 access_key: Optional[pulumi.Input[str]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_protocol: Optional[pulumi.Input['LocationObjectStorageServerProtocol']] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a LocationObjectStorage resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        :param pulumi.Input[str] access_key: Optional. The access key is used if credentials are required to access the self-managed object storage server.
        :param pulumi.Input[str] bucket_name: The name of the bucket on the self-managed object storage server.
        :param pulumi.Input[str] secret_key: Optional. The secret key is used if credentials are required to access the self-managed object storage server.
        :param pulumi.Input[str] server_certificate: X.509 PEM content containing a certificate authority or chain to trust.
        :param pulumi.Input[str] server_hostname: The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
        :param pulumi.Input[int] server_port: The port that your self-managed server accepts inbound network traffic on.
        :param pulumi.Input['LocationObjectStorageServerProtocol'] server_protocol: The protocol that the object storage server uses to communicate.
        :param pulumi.Input[str] subdirectory: The subdirectory in the self-managed object storage server that is used to read data from.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        pulumi.set(__self__, "agent_arns", agent_arns)
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if server_certificate is not None:
            pulumi.set(__self__, "server_certificate", server_certificate)
        if server_hostname is not None:
            pulumi.set(__self__, "server_hostname", server_hostname)
        if server_port is not None:
            pulumi.set(__self__, "server_port", server_port)
        if server_protocol is not None:
            pulumi.set(__self__, "server_protocol", server_protocol)
        if subdirectory is not None:
            pulumi.set(__self__, "subdirectory", subdirectory)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        """
        return pulumi.get(self, "agent_arns")

    @agent_arns.setter
    def agent_arns(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "agent_arns", value)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The access key is used if credentials are required to access the self-managed object storage server.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket on the self-managed object storage server.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The secret key is used if credentials are required to access the self-managed object storage server.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        X.509 PEM content containing a certificate authority or chain to trust.
        """
        return pulumi.get(self, "server_certificate")

    @server_certificate.setter
    def server_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_certificate", value)

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
        """
        return pulumi.get(self, "server_hostname")

    @server_hostname.setter
    def server_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_hostname", value)

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that your self-managed server accepts inbound network traffic on.
        """
        return pulumi.get(self, "server_port")

    @server_port.setter
    def server_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "server_port", value)

    @property
    @pulumi.getter(name="serverProtocol")
    def server_protocol(self) -> Optional[pulumi.Input['LocationObjectStorageServerProtocol']]:
        """
        The protocol that the object storage server uses to communicate.
        """
        return pulumi.get(self, "server_protocol")

    @server_protocol.setter
    def server_protocol(self, value: Optional[pulumi.Input['LocationObjectStorageServerProtocol']]):
        pulumi.set(self, "server_protocol", value)

    @property
    @pulumi.getter
    def subdirectory(self) -> Optional[pulumi.Input[str]]:
        """
        The subdirectory in the self-managed object storage server that is used to read data from.
        """
        return pulumi.get(self, "subdirectory")

    @subdirectory.setter
    def subdirectory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdirectory", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class LocationObjectStorage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_protocol: Optional[pulumi.Input['LocationObjectStorageServerProtocol']] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        """
        Resource schema for AWS::DataSync::LocationObjectStorage.

        ## Example Usage
        ### Example

        ```python
        import pulumi
        import pulumi_aws_native as aws_native

        location_object_storage = aws_native.datasync.LocationObjectStorage("locationObjectStorage",
            agent_arns=["arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs"],
            bucket_name="MyBucket",
            server_hostname="MyServer@example.com",
            server_protocol=aws_native.datasync.LocationObjectStorageServerProtocol.HTTPS,
            subdirectory="/MySubdirectory")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Optional. The access key is used if credentials are required to access the self-managed object storage server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] agent_arns: The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        :param pulumi.Input[str] bucket_name: The name of the bucket on the self-managed object storage server.
        :param pulumi.Input[str] secret_key: Optional. The secret key is used if credentials are required to access the self-managed object storage server.
        :param pulumi.Input[str] server_certificate: X.509 PEM content containing a certificate authority or chain to trust.
        :param pulumi.Input[str] server_hostname: The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
        :param pulumi.Input[int] server_port: The port that your self-managed server accepts inbound network traffic on.
        :param pulumi.Input['LocationObjectStorageServerProtocol'] server_protocol: The protocol that the object storage server uses to communicate.
        :param pulumi.Input[str] subdirectory: The subdirectory in the self-managed object storage server that is used to read data from.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocationObjectStorageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::DataSync::LocationObjectStorage.

        ## Example Usage
        ### Example

        ```python
        import pulumi
        import pulumi_aws_native as aws_native

        location_object_storage = aws_native.datasync.LocationObjectStorage("locationObjectStorage",
            agent_arns=["arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs"],
            bucket_name="MyBucket",
            server_hostname="MyServer@example.com",
            server_protocol=aws_native.datasync.LocationObjectStorageServerProtocol.HTTPS,
            subdirectory="/MySubdirectory")

        ```

        :param str resource_name: The name of the resource.
        :param LocationObjectStorageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocationObjectStorageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 agent_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 server_certificate: Optional[pulumi.Input[str]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 server_port: Optional[pulumi.Input[int]] = None,
                 server_protocol: Optional[pulumi.Input['LocationObjectStorageServerProtocol']] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocationObjectStorageArgs.__new__(LocationObjectStorageArgs)

            __props__.__dict__["access_key"] = access_key
            if agent_arns is None and not opts.urn:
                raise TypeError("Missing required property 'agent_arns'")
            __props__.__dict__["agent_arns"] = agent_arns
            __props__.__dict__["bucket_name"] = bucket_name
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["server_certificate"] = server_certificate
            __props__.__dict__["server_hostname"] = server_hostname
            __props__.__dict__["server_port"] = server_port
            __props__.__dict__["server_protocol"] = server_protocol
            __props__.__dict__["subdirectory"] = subdirectory
            __props__.__dict__["tags"] = tags
            __props__.__dict__["location_arn"] = None
            __props__.__dict__["location_uri"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["bucketName"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(LocationObjectStorage, __self__).__init__(
            'aws-native:datasync:LocationObjectStorage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LocationObjectStorage':
        """
        Get an existing LocationObjectStorage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LocationObjectStorageArgs.__new__(LocationObjectStorageArgs)

        __props__.__dict__["access_key"] = None
        __props__.__dict__["agent_arns"] = None
        __props__.__dict__["bucket_name"] = None
        __props__.__dict__["location_arn"] = None
        __props__.__dict__["location_uri"] = None
        __props__.__dict__["secret_key"] = None
        __props__.__dict__["server_certificate"] = None
        __props__.__dict__["server_hostname"] = None
        __props__.__dict__["server_port"] = None
        __props__.__dict__["server_protocol"] = None
        __props__.__dict__["subdirectory"] = None
        __props__.__dict__["tags"] = None
        return LocationObjectStorage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. The access key is used if credentials are required to access the self-managed object storage server.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        """
        return pulumi.get(self, "agent_arns")

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the bucket on the self-managed object storage server.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the location that is created.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> pulumi.Output[str]:
        """
        The URL of the object storage location that was described.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. The secret key is used if credentials are required to access the self-managed object storage server.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> pulumi.Output[Optional[str]]:
        """
        X.509 PEM content containing a certificate authority or chain to trust.
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.
        """
        return pulumi.get(self, "server_hostname")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> pulumi.Output[Optional[int]]:
        """
        The port that your self-managed server accepts inbound network traffic on.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter(name="serverProtocol")
    def server_protocol(self) -> pulumi.Output[Optional['LocationObjectStorageServerProtocol']]:
        """
        The protocol that the object storage server uses to communicate.
        """
        return pulumi.get(self, "server_protocol")

    @property
    @pulumi.getter
    def subdirectory(self) -> pulumi.Output[Optional[str]]:
        """
        The subdirectory in the self-managed object storage server that is used to read data from.
        """
        return pulumi.get(self, "subdirectory")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

