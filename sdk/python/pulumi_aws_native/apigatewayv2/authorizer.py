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
from ._inputs import *

__all__ = ['AuthorizerArgs', 'Authorizer']

@pulumi.input_type
class AuthorizerArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[builtins.str],
                 authorizer_type: pulumi.Input[builtins.str],
                 authorizer_credentials_arn: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 authorizer_uri: Optional[pulumi.Input[builtins.str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[builtins.bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[builtins.str]] = None,
                 jwt_configuration: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Authorizer resource.
        :param pulumi.Input[builtins.str] api_id: The API identifier.
        :param pulumi.Input[builtins.str] authorizer_type: The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs).
        :param pulumi.Input[builtins.str] authorizer_credentials_arn: Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers.
        :param pulumi.Input[builtins.str] authorizer_payload_format_version: Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        :param pulumi.Input[builtins.int] authorizer_result_ttl_in_seconds: The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[builtins.str] authorizer_uri: The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
        :param pulumi.Input[builtins.bool] enable_simple_responses: Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] identity_source: The identity source for which authorization is requested.
                For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
                For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
        :param pulumi.Input[builtins.str] identity_validation_expression: This parameter is not used.
        :param pulumi.Input['AuthorizerJwtConfigurationArgs'] jwt_configuration: The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
        :param pulumi.Input[builtins.str] name: The name of the authorizer.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "authorizer_type", authorizer_type)
        if authorizer_credentials_arn is not None:
            pulumi.set(__self__, "authorizer_credentials_arn", authorizer_credentials_arn)
        if authorizer_payload_format_version is not None:
            pulumi.set(__self__, "authorizer_payload_format_version", authorizer_payload_format_version)
        if authorizer_result_ttl_in_seconds is not None:
            pulumi.set(__self__, "authorizer_result_ttl_in_seconds", authorizer_result_ttl_in_seconds)
        if authorizer_uri is not None:
            pulumi.set(__self__, "authorizer_uri", authorizer_uri)
        if enable_simple_responses is not None:
            pulumi.set(__self__, "enable_simple_responses", enable_simple_responses)
        if identity_source is not None:
            pulumi.set(__self__, "identity_source", identity_source)
        if identity_validation_expression is not None:
            pulumi.set(__self__, "identity_validation_expression", identity_validation_expression)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[builtins.str]:
        """
        The API identifier.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Input[builtins.str]:
        """
        The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs).
        """
        return pulumi.get(self, "authorizer_type")

    @authorizer_type.setter
    def authorizer_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "authorizer_type", value)

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers.
        """
        return pulumi.get(self, "authorizer_credentials_arn")

    @authorizer_credentials_arn.setter
    def authorizer_credentials_arn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "authorizer_credentials_arn", value)

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        """
        return pulumi.get(self, "authorizer_payload_format_version")

    @authorizer_payload_format_version.setter
    def authorizer_payload_format_version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "authorizer_payload_format_version", value)

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
        """
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @authorizer_result_ttl_in_seconds.setter
    def authorizer_result_ttl_in_seconds(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "authorizer_result_ttl_in_seconds", value)

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
        """
        return pulumi.get(self, "authorizer_uri")

    @authorizer_uri.setter
    def authorizer_uri(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "authorizer_uri", value)

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        """
        return pulumi.get(self, "enable_simple_responses")

    @enable_simple_responses.setter
    def enable_simple_responses(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_simple_responses", value)

    @property
    @pulumi.getter(name="identitySource")
    def identity_source(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The identity source for which authorization is requested.
         For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
         For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
        """
        return pulumi.get(self, "identity_source")

    @identity_source.setter
    def identity_source(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "identity_source", value)

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        This parameter is not used.
        """
        return pulumi.get(self, "identity_validation_expression")

    @identity_validation_expression.setter
    def identity_validation_expression(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "identity_validation_expression", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]:
        """
        The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
        """
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['AuthorizerJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the authorizer.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("aws-native:apigatewayv2:Authorizer")
class Authorizer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 authorizer_type: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_uri: Optional[pulumi.Input[builtins.str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[builtins.bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[builtins.str]] = None,
                 jwt_configuration: Optional[pulumi.Input[Union['AuthorizerJwtConfigurationArgs', 'AuthorizerJwtConfigurationArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The ``AWS::ApiGatewayV2::Authorizer`` resource creates an authorizer for a WebSocket API or an HTTP API. To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) and [Controlling and managing access to an HTTP API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) in the *API Gateway Developer Guide*.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_id: The API identifier.
        :param pulumi.Input[builtins.str] authorizer_credentials_arn: Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers.
        :param pulumi.Input[builtins.str] authorizer_payload_format_version: Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        :param pulumi.Input[builtins.int] authorizer_result_ttl_in_seconds: The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
        :param pulumi.Input[builtins.str] authorizer_type: The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs).
        :param pulumi.Input[builtins.str] authorizer_uri: The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
        :param pulumi.Input[builtins.bool] enable_simple_responses: Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] identity_source: The identity source for which authorization is requested.
                For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
                For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
        :param pulumi.Input[builtins.str] identity_validation_expression: This parameter is not used.
        :param pulumi.Input[Union['AuthorizerJwtConfigurationArgs', 'AuthorizerJwtConfigurationArgsDict']] jwt_configuration: The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
        :param pulumi.Input[builtins.str] name: The name of the authorizer.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``AWS::ApiGatewayV2::Authorizer`` resource creates an authorizer for a WebSocket API or an HTTP API. To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) and [Controlling and managing access to an HTTP API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) in the *API Gateway Developer Guide*.

        :param str resource_name: The name of the resource.
        :param AuthorizerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_credentials_arn: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_payload_format_version: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_result_ttl_in_seconds: Optional[pulumi.Input[builtins.int]] = None,
                 authorizer_type: Optional[pulumi.Input[builtins.str]] = None,
                 authorizer_uri: Optional[pulumi.Input[builtins.str]] = None,
                 enable_simple_responses: Optional[pulumi.Input[builtins.bool]] = None,
                 identity_source: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 identity_validation_expression: Optional[pulumi.Input[builtins.str]] = None,
                 jwt_configuration: Optional[pulumi.Input[Union['AuthorizerJwtConfigurationArgs', 'AuthorizerJwtConfigurationArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizerArgs.__new__(AuthorizerArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["authorizer_credentials_arn"] = authorizer_credentials_arn
            __props__.__dict__["authorizer_payload_format_version"] = authorizer_payload_format_version
            __props__.__dict__["authorizer_result_ttl_in_seconds"] = authorizer_result_ttl_in_seconds
            if authorizer_type is None and not opts.urn:
                raise TypeError("Missing required property 'authorizer_type'")
            __props__.__dict__["authorizer_type"] = authorizer_type
            __props__.__dict__["authorizer_uri"] = authorizer_uri
            __props__.__dict__["enable_simple_responses"] = enable_simple_responses
            __props__.__dict__["identity_source"] = identity_source
            __props__.__dict__["identity_validation_expression"] = identity_validation_expression
            __props__.__dict__["jwt_configuration"] = jwt_configuration
            __props__.__dict__["name"] = name
            __props__.__dict__["authorizer_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["apiId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Authorizer, __self__).__init__(
            'aws-native:apigatewayv2:Authorizer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Authorizer':
        """
        Get an existing Authorizer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AuthorizerArgs.__new__(AuthorizerArgs)

        __props__.__dict__["api_id"] = None
        __props__.__dict__["authorizer_credentials_arn"] = None
        __props__.__dict__["authorizer_id"] = None
        __props__.__dict__["authorizer_payload_format_version"] = None
        __props__.__dict__["authorizer_result_ttl_in_seconds"] = None
        __props__.__dict__["authorizer_type"] = None
        __props__.__dict__["authorizer_uri"] = None
        __props__.__dict__["enable_simple_responses"] = None
        __props__.__dict__["identity_source"] = None
        __props__.__dict__["identity_validation_expression"] = None
        __props__.__dict__["jwt_configuration"] = None
        __props__.__dict__["name"] = None
        return Authorizer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[builtins.str]:
        """
        The API identifier.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="authorizerCredentialsArn")
    def authorizer_credentials_arn(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers.
        """
        return pulumi.get(self, "authorizer_credentials_arn")

    @property
    @pulumi.getter(name="authorizerId")
    def authorizer_id(self) -> pulumi.Output[builtins.str]:
        """
        The authorizer ID.
        """
        return pulumi.get(self, "authorizer_id")

    @property
    @pulumi.getter(name="authorizerPayloadFormatVersion")
    def authorizer_payload_format_version(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        """
        return pulumi.get(self, "authorizer_payload_format_version")

    @property
    @pulumi.getter(name="authorizerResultTtlInSeconds")
    def authorizer_result_ttl_in_seconds(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
        """
        return pulumi.get(self, "authorizer_result_ttl_in_seconds")

    @property
    @pulumi.getter(name="authorizerType")
    def authorizer_type(self) -> pulumi.Output[builtins.str]:
        """
        The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs).
        """
        return pulumi.get(self, "authorizer_type")

    @property
    @pulumi.getter(name="authorizerUri")
    def authorizer_uri(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
        """
        return pulumi.get(self, "authorizer_uri")

    @property
    @pulumi.getter(name="enableSimpleResponses")
    def enable_simple_responses(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
        """
        return pulumi.get(self, "enable_simple_responses")

    @property
    @pulumi.getter(name="identitySource")
    def identity_source(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The identity source for which authorization is requested.
         For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
         For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
        """
        return pulumi.get(self, "identity_source")

    @property
    @pulumi.getter(name="identityValidationExpression")
    def identity_validation_expression(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        This parameter is not used.
        """
        return pulumi.get(self, "identity_validation_expression")

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> pulumi.Output[Optional['outputs.AuthorizerJwtConfiguration']]:
        """
        The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
        """
        return pulumi.get(self, "jwt_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the authorizer.
        """
        return pulumi.get(self, "name")

