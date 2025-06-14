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
from ._enums import *
from ._inputs import *

__all__ = ['DomainConfigurationArgs', 'DomainConfiguration']

@pulumi.input_type
class DomainConfigurationArgs:
    def __init__(__self__, *,
                 application_protocol: Optional[pulumi.Input['DomainConfigurationApplicationProtocol']] = None,
                 authentication_type: Optional[pulumi.Input['DomainConfigurationAuthenticationType']] = None,
                 authorizer_config: Optional[pulumi.Input['DomainConfigurationAuthorizerConfigArgs']] = None,
                 client_certificate_config: Optional[pulumi.Input['DomainConfigurationClientCertificateConfigArgs']] = None,
                 domain_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 domain_configuration_status: Optional[pulumi.Input['DomainConfigurationStatus']] = None,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_certificate_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 server_certificate_config: Optional[pulumi.Input['DomainConfigurationServerCertificateConfigArgs']] = None,
                 service_type: Optional[pulumi.Input['DomainConfigurationServiceType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 tls_config: Optional[pulumi.Input['DomainConfigurationTlsConfigArgs']] = None,
                 validation_certificate_arn: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DomainConfiguration resource.
        :param pulumi.Input['DomainConfigurationApplicationProtocol'] application_protocol: An enumerated string that speciﬁes the application-layer protocol.
        :param pulumi.Input['DomainConfigurationAuthenticationType'] authentication_type: An enumerated string that speciﬁes the authentication type.
        :param pulumi.Input['DomainConfigurationAuthorizerConfigArgs'] authorizer_config: An object that specifies the authorization service for a domain.
        :param pulumi.Input['DomainConfigurationClientCertificateConfigArgs'] client_certificate_config: An object that speciﬁes the client certificate conﬁguration for a domain.
        :param pulumi.Input[builtins.str] domain_configuration_name: The name of the domain configuration. This value must be unique to a region.
        :param pulumi.Input['DomainConfigurationStatus'] domain_configuration_status: The status to which the domain configuration should be updated.
               
               Valid values: `ENABLED` | `DISABLED`
        :param pulumi.Input[builtins.str] domain_name: The name of the domain.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] server_certificate_arns: The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
        :param pulumi.Input['DomainConfigurationServerCertificateConfigArgs'] server_certificate_config: The server certificate configuration.
               
               For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
        :param pulumi.Input['DomainConfigurationServiceType'] service_type: The type of service delivered by the endpoint.
               
               > AWS IoT Core currently supports only the `DATA` service type.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: Metadata which can be used to manage the domain configuration.
               
               > For URI Request parameters use format: ...key1=value1&key2=value2...
               > 
               > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
               > 
               > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
        :param pulumi.Input['DomainConfigurationTlsConfigArgs'] tls_config: An object that specifies the TLS configuration for a domain.
        :param pulumi.Input[builtins.str] validation_certificate_arn: The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
        """
        if application_protocol is not None:
            pulumi.set(__self__, "application_protocol", application_protocol)
        if authentication_type is not None:
            pulumi.set(__self__, "authentication_type", authentication_type)
        if authorizer_config is not None:
            pulumi.set(__self__, "authorizer_config", authorizer_config)
        if client_certificate_config is not None:
            pulumi.set(__self__, "client_certificate_config", client_certificate_config)
        if domain_configuration_name is not None:
            pulumi.set(__self__, "domain_configuration_name", domain_configuration_name)
        if domain_configuration_status is not None:
            pulumi.set(__self__, "domain_configuration_status", domain_configuration_status)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if server_certificate_arns is not None:
            pulumi.set(__self__, "server_certificate_arns", server_certificate_arns)
        if server_certificate_config is not None:
            pulumi.set(__self__, "server_certificate_config", server_certificate_config)
        if service_type is not None:
            pulumi.set(__self__, "service_type", service_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tls_config is not None:
            pulumi.set(__self__, "tls_config", tls_config)
        if validation_certificate_arn is not None:
            pulumi.set(__self__, "validation_certificate_arn", validation_certificate_arn)

    @property
    @pulumi.getter(name="applicationProtocol")
    def application_protocol(self) -> Optional[pulumi.Input['DomainConfigurationApplicationProtocol']]:
        """
        An enumerated string that speciﬁes the application-layer protocol.
        """
        return pulumi.get(self, "application_protocol")

    @application_protocol.setter
    def application_protocol(self, value: Optional[pulumi.Input['DomainConfigurationApplicationProtocol']]):
        pulumi.set(self, "application_protocol", value)

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> Optional[pulumi.Input['DomainConfigurationAuthenticationType']]:
        """
        An enumerated string that speciﬁes the authentication type.
        """
        return pulumi.get(self, "authentication_type")

    @authentication_type.setter
    def authentication_type(self, value: Optional[pulumi.Input['DomainConfigurationAuthenticationType']]):
        pulumi.set(self, "authentication_type", value)

    @property
    @pulumi.getter(name="authorizerConfig")
    def authorizer_config(self) -> Optional[pulumi.Input['DomainConfigurationAuthorizerConfigArgs']]:
        """
        An object that specifies the authorization service for a domain.
        """
        return pulumi.get(self, "authorizer_config")

    @authorizer_config.setter
    def authorizer_config(self, value: Optional[pulumi.Input['DomainConfigurationAuthorizerConfigArgs']]):
        pulumi.set(self, "authorizer_config", value)

    @property
    @pulumi.getter(name="clientCertificateConfig")
    def client_certificate_config(self) -> Optional[pulumi.Input['DomainConfigurationClientCertificateConfigArgs']]:
        """
        An object that speciﬁes the client certificate conﬁguration for a domain.
        """
        return pulumi.get(self, "client_certificate_config")

    @client_certificate_config.setter
    def client_certificate_config(self, value: Optional[pulumi.Input['DomainConfigurationClientCertificateConfigArgs']]):
        pulumi.set(self, "client_certificate_config", value)

    @property
    @pulumi.getter(name="domainConfigurationName")
    def domain_configuration_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the domain configuration. This value must be unique to a region.
        """
        return pulumi.get(self, "domain_configuration_name")

    @domain_configuration_name.setter
    def domain_configuration_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_configuration_name", value)

    @property
    @pulumi.getter(name="domainConfigurationStatus")
    def domain_configuration_status(self) -> Optional[pulumi.Input['DomainConfigurationStatus']]:
        """
        The status to which the domain configuration should be updated.

        Valid values: `ENABLED` | `DISABLED`
        """
        return pulumi.get(self, "domain_configuration_status")

    @domain_configuration_status.setter
    def domain_configuration_status(self, value: Optional[pulumi.Input['DomainConfigurationStatus']]):
        pulumi.set(self, "domain_configuration_status", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the domain.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="serverCertificateArns")
    def server_certificate_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
        """
        return pulumi.get(self, "server_certificate_arns")

    @server_certificate_arns.setter
    def server_certificate_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "server_certificate_arns", value)

    @property
    @pulumi.getter(name="serverCertificateConfig")
    def server_certificate_config(self) -> Optional[pulumi.Input['DomainConfigurationServerCertificateConfigArgs']]:
        """
        The server certificate configuration.

        For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
        """
        return pulumi.get(self, "server_certificate_config")

    @server_certificate_config.setter
    def server_certificate_config(self, value: Optional[pulumi.Input['DomainConfigurationServerCertificateConfigArgs']]):
        pulumi.set(self, "server_certificate_config", value)

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[pulumi.Input['DomainConfigurationServiceType']]:
        """
        The type of service delivered by the endpoint.

        > AWS IoT Core currently supports only the `DATA` service type.
        """
        return pulumi.get(self, "service_type")

    @service_type.setter
    def service_type(self, value: Optional[pulumi.Input['DomainConfigurationServiceType']]):
        pulumi.set(self, "service_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        Metadata which can be used to manage the domain configuration.

        > For URI Request parameters use format: ...key1=value1&key2=value2...
        > 
        > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
        > 
        > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tlsConfig")
    def tls_config(self) -> Optional[pulumi.Input['DomainConfigurationTlsConfigArgs']]:
        """
        An object that specifies the TLS configuration for a domain.
        """
        return pulumi.get(self, "tls_config")

    @tls_config.setter
    def tls_config(self, value: Optional[pulumi.Input['DomainConfigurationTlsConfigArgs']]):
        pulumi.set(self, "tls_config", value)

    @property
    @pulumi.getter(name="validationCertificateArn")
    def validation_certificate_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
        """
        return pulumi.get(self, "validation_certificate_arn")

    @validation_certificate_arn.setter
    def validation_certificate_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "validation_certificate_arn", value)


@pulumi.type_token("aws-native:iot:DomainConfiguration")
class DomainConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_protocol: Optional[pulumi.Input['DomainConfigurationApplicationProtocol']] = None,
                 authentication_type: Optional[pulumi.Input['DomainConfigurationAuthenticationType']] = None,
                 authorizer_config: Optional[pulumi.Input[Union['DomainConfigurationAuthorizerConfigArgs', 'DomainConfigurationAuthorizerConfigArgsDict']]] = None,
                 client_certificate_config: Optional[pulumi.Input[Union['DomainConfigurationClientCertificateConfigArgs', 'DomainConfigurationClientCertificateConfigArgsDict']]] = None,
                 domain_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 domain_configuration_status: Optional[pulumi.Input['DomainConfigurationStatus']] = None,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_certificate_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 server_certificate_config: Optional[pulumi.Input[Union['DomainConfigurationServerCertificateConfigArgs', 'DomainConfigurationServerCertificateConfigArgsDict']]] = None,
                 service_type: Optional[pulumi.Input['DomainConfigurationServiceType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 tls_config: Optional[pulumi.Input[Union['DomainConfigurationTlsConfigArgs', 'DomainConfigurationTlsConfigArgsDict']]] = None,
                 validation_certificate_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Create and manage a Domain Configuration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['DomainConfigurationApplicationProtocol'] application_protocol: An enumerated string that speciﬁes the application-layer protocol.
        :param pulumi.Input['DomainConfigurationAuthenticationType'] authentication_type: An enumerated string that speciﬁes the authentication type.
        :param pulumi.Input[Union['DomainConfigurationAuthorizerConfigArgs', 'DomainConfigurationAuthorizerConfigArgsDict']] authorizer_config: An object that specifies the authorization service for a domain.
        :param pulumi.Input[Union['DomainConfigurationClientCertificateConfigArgs', 'DomainConfigurationClientCertificateConfigArgsDict']] client_certificate_config: An object that speciﬁes the client certificate conﬁguration for a domain.
        :param pulumi.Input[builtins.str] domain_configuration_name: The name of the domain configuration. This value must be unique to a region.
        :param pulumi.Input['DomainConfigurationStatus'] domain_configuration_status: The status to which the domain configuration should be updated.
               
               Valid values: `ENABLED` | `DISABLED`
        :param pulumi.Input[builtins.str] domain_name: The name of the domain.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] server_certificate_arns: The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
        :param pulumi.Input[Union['DomainConfigurationServerCertificateConfigArgs', 'DomainConfigurationServerCertificateConfigArgsDict']] server_certificate_config: The server certificate configuration.
               
               For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
        :param pulumi.Input['DomainConfigurationServiceType'] service_type: The type of service delivered by the endpoint.
               
               > AWS IoT Core currently supports only the `DATA` service type.
        :param pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]] tags: Metadata which can be used to manage the domain configuration.
               
               > For URI Request parameters use format: ...key1=value1&key2=value2...
               > 
               > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
               > 
               > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
        :param pulumi.Input[Union['DomainConfigurationTlsConfigArgs', 'DomainConfigurationTlsConfigArgsDict']] tls_config: An object that specifies the TLS configuration for a domain.
        :param pulumi.Input[builtins.str] validation_certificate_arn: The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DomainConfigurationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create and manage a Domain Configuration

        :param str resource_name: The name of the resource.
        :param DomainConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_protocol: Optional[pulumi.Input['DomainConfigurationApplicationProtocol']] = None,
                 authentication_type: Optional[pulumi.Input['DomainConfigurationAuthenticationType']] = None,
                 authorizer_config: Optional[pulumi.Input[Union['DomainConfigurationAuthorizerConfigArgs', 'DomainConfigurationAuthorizerConfigArgsDict']]] = None,
                 client_certificate_config: Optional[pulumi.Input[Union['DomainConfigurationClientCertificateConfigArgs', 'DomainConfigurationClientCertificateConfigArgsDict']]] = None,
                 domain_configuration_name: Optional[pulumi.Input[builtins.str]] = None,
                 domain_configuration_status: Optional[pulumi.Input['DomainConfigurationStatus']] = None,
                 domain_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_certificate_arns: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 server_certificate_config: Optional[pulumi.Input[Union['DomainConfigurationServerCertificateConfigArgs', 'DomainConfigurationServerCertificateConfigArgsDict']]] = None,
                 service_type: Optional[pulumi.Input['DomainConfigurationServiceType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['_root_inputs.TagArgs', '_root_inputs.TagArgsDict']]]]] = None,
                 tls_config: Optional[pulumi.Input[Union['DomainConfigurationTlsConfigArgs', 'DomainConfigurationTlsConfigArgsDict']]] = None,
                 validation_certificate_arn: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainConfigurationArgs.__new__(DomainConfigurationArgs)

            __props__.__dict__["application_protocol"] = application_protocol
            __props__.__dict__["authentication_type"] = authentication_type
            __props__.__dict__["authorizer_config"] = authorizer_config
            __props__.__dict__["client_certificate_config"] = client_certificate_config
            __props__.__dict__["domain_configuration_name"] = domain_configuration_name
            __props__.__dict__["domain_configuration_status"] = domain_configuration_status
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["server_certificate_arns"] = server_certificate_arns
            __props__.__dict__["server_certificate_config"] = server_certificate_config
            __props__.__dict__["service_type"] = service_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tls_config"] = tls_config
            __props__.__dict__["validation_certificate_arn"] = validation_certificate_arn
            __props__.__dict__["arn"] = None
            __props__.__dict__["domain_type"] = None
            __props__.__dict__["server_certificates"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domainConfigurationName", "domainName", "serverCertificateArns[*]", "serviceType", "validationCertificateArn"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DomainConfiguration, __self__).__init__(
            'aws-native:iot:DomainConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DomainConfiguration':
        """
        Get an existing DomainConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainConfigurationArgs.__new__(DomainConfigurationArgs)

        __props__.__dict__["application_protocol"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["authentication_type"] = None
        __props__.__dict__["authorizer_config"] = None
        __props__.__dict__["client_certificate_config"] = None
        __props__.__dict__["domain_configuration_name"] = None
        __props__.__dict__["domain_configuration_status"] = None
        __props__.__dict__["domain_name"] = None
        __props__.__dict__["domain_type"] = None
        __props__.__dict__["server_certificate_arns"] = None
        __props__.__dict__["server_certificate_config"] = None
        __props__.__dict__["server_certificates"] = None
        __props__.__dict__["service_type"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tls_config"] = None
        __props__.__dict__["validation_certificate_arn"] = None
        return DomainConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationProtocol")
    def application_protocol(self) -> pulumi.Output[Optional['DomainConfigurationApplicationProtocol']]:
        """
        An enumerated string that speciﬁes the application-layer protocol.
        """
        return pulumi.get(self, "application_protocol")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The Amazon Resource Name (ARN) of the domain configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Output[Optional['DomainConfigurationAuthenticationType']]:
        """
        An enumerated string that speciﬁes the authentication type.
        """
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="authorizerConfig")
    def authorizer_config(self) -> pulumi.Output[Optional['outputs.DomainConfigurationAuthorizerConfig']]:
        """
        An object that specifies the authorization service for a domain.
        """
        return pulumi.get(self, "authorizer_config")

    @property
    @pulumi.getter(name="clientCertificateConfig")
    def client_certificate_config(self) -> pulumi.Output[Optional['outputs.DomainConfigurationClientCertificateConfig']]:
        """
        An object that speciﬁes the client certificate conﬁguration for a domain.
        """
        return pulumi.get(self, "client_certificate_config")

    @property
    @pulumi.getter(name="domainConfigurationName")
    def domain_configuration_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the domain configuration. This value must be unique to a region.
        """
        return pulumi.get(self, "domain_configuration_name")

    @property
    @pulumi.getter(name="domainConfigurationStatus")
    def domain_configuration_status(self) -> pulumi.Output[Optional['DomainConfigurationStatus']]:
        """
        The status to which the domain configuration should be updated.

        Valid values: `ENABLED` | `DISABLED`
        """
        return pulumi.get(self, "domain_configuration_status")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the domain.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="domainType")
    def domain_type(self) -> pulumi.Output['DomainConfigurationDomainType']:
        """
        The type of service delivered by the domain.
        """
        return pulumi.get(self, "domain_type")

    @property
    @pulumi.getter(name="serverCertificateArns")
    def server_certificate_arns(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
        """
        return pulumi.get(self, "server_certificate_arns")

    @property
    @pulumi.getter(name="serverCertificateConfig")
    def server_certificate_config(self) -> pulumi.Output[Optional['outputs.DomainConfigurationServerCertificateConfig']]:
        """
        The server certificate configuration.

        For more information, see [Configurable endpoints](https://docs.aws.amazon.com//iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.
        """
        return pulumi.get(self, "server_certificate_config")

    @property
    @pulumi.getter(name="serverCertificates")
    def server_certificates(self) -> pulumi.Output[Sequence['outputs.DomainConfigurationServerCertificateSummary']]:
        """
        The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN. This value is not required for AWS -managed domains.
        """
        return pulumi.get(self, "server_certificates")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> pulumi.Output[Optional['DomainConfigurationServiceType']]:
        """
        The type of service delivered by the endpoint.

        > AWS IoT Core currently supports only the `DATA` service type.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        Metadata which can be used to manage the domain configuration.

        > For URI Request parameters use format: ...key1=value1&key2=value2...
        > 
        > For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
        > 
        > For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsConfig")
    def tls_config(self) -> pulumi.Output[Optional['outputs.DomainConfigurationTlsConfig']]:
        """
        An object that specifies the TLS configuration for a domain.
        """
        return pulumi.get(self, "tls_config")

    @property
    @pulumi.getter(name="validationCertificateArn")
    def validation_certificate_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority. This value is not required for AWS -managed domains.
        """
        return pulumi.get(self, "validation_certificate_arn")

