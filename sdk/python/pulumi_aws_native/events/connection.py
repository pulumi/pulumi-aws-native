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

__all__ = ['ConnectionArgs', 'Connection']

@pulumi.input_type
class ConnectionArgs:
    def __init__(__self__, *,
                 auth_parameters: Optional[pulumi.Input['ConnectionAuthParametersArgs']] = None,
                 authorization_type: Optional[pulumi.Input['ConnectionAuthorizationType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 invocation_connectivity_parameters: Optional[pulumi.Input['InvocationConnectivityParametersPropertiesArgs']] = None,
                 kms_key_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Connection resource.
        :param pulumi.Input['ConnectionAuthParametersArgs'] auth_parameters: The authorization parameters to use to authorize with the endpoint.
               
               You must include only authorization parameters for the `AuthorizationType` you specify.
        :param pulumi.Input['ConnectionAuthorizationType'] authorization_type: The type of authorization to use for the connection.
               
               > OAUTH tokens are refreshed when a 401 or 407 response is returned.
        :param pulumi.Input[builtins.str] description: Description of the connection.
        :param pulumi.Input['InvocationConnectivityParametersPropertiesArgs'] invocation_connectivity_parameters: The private resource the HTTP request will be sent to.
        :param pulumi.Input[builtins.str] kms_key_identifier: The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
               
               If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the connection.
               
               For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
        :param pulumi.Input[builtins.str] name: Name of the connection.
        """
        if auth_parameters is not None:
            pulumi.set(__self__, "auth_parameters", auth_parameters)
        if authorization_type is not None:
            pulumi.set(__self__, "authorization_type", authorization_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if invocation_connectivity_parameters is not None:
            pulumi.set(__self__, "invocation_connectivity_parameters", invocation_connectivity_parameters)
        if kms_key_identifier is not None:
            pulumi.set(__self__, "kms_key_identifier", kms_key_identifier)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> Optional[pulumi.Input['ConnectionAuthParametersArgs']]:
        """
        The authorization parameters to use to authorize with the endpoint.

        You must include only authorization parameters for the `AuthorizationType` you specify.
        """
        return pulumi.get(self, "auth_parameters")

    @auth_parameters.setter
    def auth_parameters(self, value: Optional[pulumi.Input['ConnectionAuthParametersArgs']]):
        pulumi.set(self, "auth_parameters", value)

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional[pulumi.Input['ConnectionAuthorizationType']]:
        """
        The type of authorization to use for the connection.

        > OAUTH tokens are refreshed when a 401 or 407 response is returned.
        """
        return pulumi.get(self, "authorization_type")

    @authorization_type.setter
    def authorization_type(self, value: Optional[pulumi.Input['ConnectionAuthorizationType']]):
        pulumi.set(self, "authorization_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the connection.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="invocationConnectivityParameters")
    def invocation_connectivity_parameters(self) -> Optional[pulumi.Input['InvocationConnectivityParametersPropertiesArgs']]:
        """
        The private resource the HTTP request will be sent to.
        """
        return pulumi.get(self, "invocation_connectivity_parameters")

    @invocation_connectivity_parameters.setter
    def invocation_connectivity_parameters(self, value: Optional[pulumi.Input['InvocationConnectivityParametersPropertiesArgs']]):
        pulumi.set(self, "invocation_connectivity_parameters", value)

    @property
    @pulumi.getter(name="kmsKeyIdentifier")
    def kms_key_identifier(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

        If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the connection.

        For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
        """
        return pulumi.get(self, "kms_key_identifier")

    @kms_key_identifier.setter
    def kms_key_identifier(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kms_key_identifier", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the connection.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("aws-native:events:Connection")
class Connection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_parameters: Optional[pulumi.Input[Union['ConnectionAuthParametersArgs', 'ConnectionAuthParametersArgsDict']]] = None,
                 authorization_type: Optional[pulumi.Input['ConnectionAuthorizationType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 invocation_connectivity_parameters: Optional[pulumi.Input[Union['InvocationConnectivityParametersPropertiesArgs', 'InvocationConnectivityParametersPropertiesArgsDict']]] = None,
                 kms_key_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Events::Connection.

        ## Example Usage
        ### Example

        ```python
        import pulumi
        import pulumi_aws_native as aws_native

        config = pulumi.Config()
        pager_duty_api_key_param = config.require("pagerDutyAPIKeyParam")
        my_connection = aws_native.events.Connection("myConnection",
            authorization_type=aws_native.events.ConnectionAuthorizationType.API_KEY,
            description="Connection to PagerDuty API",
            auth_parameters={
                "api_key_auth_parameters": {
                    "api_key_name": "PagerDuty Authorization",
                    "api_key_value": pager_duty_api_key_param,
                },
            })
        my_api_destination = aws_native.events.ApiDestination("myApiDestination",
            connection_arn=my_connection.arn,
            description="API Destination to send events to PagerDuty",
            http_method=aws_native.events.ApiDestinationHttpMethod.POST,
            invocation_endpoint="https://events.pagerduty.com/v2/enqueue")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ConnectionAuthParametersArgs', 'ConnectionAuthParametersArgsDict']] auth_parameters: The authorization parameters to use to authorize with the endpoint.
               
               You must include only authorization parameters for the `AuthorizationType` you specify.
        :param pulumi.Input['ConnectionAuthorizationType'] authorization_type: The type of authorization to use for the connection.
               
               > OAUTH tokens are refreshed when a 401 or 407 response is returned.
        :param pulumi.Input[builtins.str] description: Description of the connection.
        :param pulumi.Input[Union['InvocationConnectivityParametersPropertiesArgs', 'InvocationConnectivityParametersPropertiesArgsDict']] invocation_connectivity_parameters: The private resource the HTTP request will be sent to.
        :param pulumi.Input[builtins.str] kms_key_identifier: The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
               
               If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the connection.
               
               For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
        :param pulumi.Input[builtins.str] name: Name of the connection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConnectionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Events::Connection.

        ## Example Usage
        ### Example

        ```python
        import pulumi
        import pulumi_aws_native as aws_native

        config = pulumi.Config()
        pager_duty_api_key_param = config.require("pagerDutyAPIKeyParam")
        my_connection = aws_native.events.Connection("myConnection",
            authorization_type=aws_native.events.ConnectionAuthorizationType.API_KEY,
            description="Connection to PagerDuty API",
            auth_parameters={
                "api_key_auth_parameters": {
                    "api_key_name": "PagerDuty Authorization",
                    "api_key_value": pager_duty_api_key_param,
                },
            })
        my_api_destination = aws_native.events.ApiDestination("myApiDestination",
            connection_arn=my_connection.arn,
            description="API Destination to send events to PagerDuty",
            http_method=aws_native.events.ApiDestinationHttpMethod.POST,
            invocation_endpoint="https://events.pagerduty.com/v2/enqueue")

        ```

        :param str resource_name: The name of the resource.
        :param ConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_parameters: Optional[pulumi.Input[Union['ConnectionAuthParametersArgs', 'ConnectionAuthParametersArgsDict']]] = None,
                 authorization_type: Optional[pulumi.Input['ConnectionAuthorizationType']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 invocation_connectivity_parameters: Optional[pulumi.Input[Union['InvocationConnectivityParametersPropertiesArgs', 'InvocationConnectivityParametersPropertiesArgsDict']]] = None,
                 kms_key_identifier: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectionArgs.__new__(ConnectionArgs)

            __props__.__dict__["auth_parameters"] = auth_parameters
            __props__.__dict__["authorization_type"] = authorization_type
            __props__.__dict__["description"] = description
            __props__.__dict__["invocation_connectivity_parameters"] = invocation_connectivity_parameters
            __props__.__dict__["kms_key_identifier"] = kms_key_identifier
            __props__.__dict__["name"] = name
            __props__.__dict__["arn"] = None
            __props__.__dict__["arn_for_policy"] = None
            __props__.__dict__["secret_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Connection, __self__).__init__(
            'aws-native:events:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Connection':
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ConnectionArgs.__new__(ConnectionArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["arn_for_policy"] = None
        __props__.__dict__["auth_parameters"] = None
        __props__.__dict__["authorization_type"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["invocation_connectivity_parameters"] = None
        __props__.__dict__["kms_key_identifier"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["secret_arn"] = None
        return Connection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[builtins.str]:
        """
        The arn of the connection resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="arnForPolicy")
    def arn_for_policy(self) -> pulumi.Output[builtins.str]:
        """
        The arn of the connection resource to be used in IAM policies.
        """
        return pulumi.get(self, "arn_for_policy")

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> pulumi.Output[Optional['outputs.ConnectionAuthParameters']]:
        """
        The authorization parameters to use to authorize with the endpoint.

        You must include only authorization parameters for the `AuthorizationType` you specify.
        """
        return pulumi.get(self, "auth_parameters")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> pulumi.Output[Optional['ConnectionAuthorizationType']]:
        """
        The type of authorization to use for the connection.

        > OAUTH tokens are refreshed when a 401 or 407 response is returned.
        """
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="invocationConnectivityParameters")
    def invocation_connectivity_parameters(self) -> pulumi.Output[Optional['outputs.InvocationConnectivityParametersProperties']]:
        """
        The private resource the HTTP request will be sent to.
        """
        return pulumi.get(self, "invocation_connectivity_parameters")

    @property
    @pulumi.getter(name="kmsKeyIdentifier")
    def kms_key_identifier(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

        If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt the connection.

        For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the *AWS Key Management Service Developer Guide* .
        """
        return pulumi.get(self, "kms_key_identifier")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Name of the connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretArn")
    def secret_arn(self) -> pulumi.Output[builtins.str]:
        """
        The arn of the secrets manager secret created in the customer account.
        """
        return pulumi.get(self, "secret_arn")

