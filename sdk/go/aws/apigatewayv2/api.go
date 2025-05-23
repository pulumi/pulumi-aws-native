// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGatewayV2::Api“ resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the *API Gateway Developer Guide*. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the *API Gateway Developer Guide.*
type Api struct {
	pulumi.CustomResourceState

	// The default endpoint for an API. For example: `https://abcdef.execute-api.us-west-2.amazonaws.com` .
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression pulumi.StringPtrOutput `pulumi:"apiKeySelectionExpression"`
	// Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
	BasePath pulumi.StringPtrOutput `pulumi:"basePath"`
	// The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::Api` for more information about the expected schema for this property.
	Body pulumi.AnyOutput `pulumi:"body"`
	// The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	BodyS3Location ApiBodyS3LocationPtrOutput `pulumi:"bodyS3Location"`
	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	CorsConfiguration ApiCorsPtrOutput `pulumi:"corsConfiguration"`
	// This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
	CredentialsArn pulumi.StringPtrOutput `pulumi:"credentialsArn"`
	// The description of the API.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint pulumi.BoolPtrOutput `pulumi:"disableExecuteApiEndpoint"`
	// Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
	DisableSchemaValidation pulumi.BoolPtrOutput `pulumi:"disableSchemaValidation"`
	// Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.
	FailOnWarnings pulumi.BoolPtrOutput `pulumi:"failOnWarnings"`
	// The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.
	//
	// Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
	IpAddressType pulumi.StringPtrOutput `pulumi:"ipAddressType"`
	// The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	ProtocolType pulumi.StringPtrOutput `pulumi:"protocolType"`
	// This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
	RouteKey pulumi.StringPtrOutput `pulumi:"routeKey"`
	// The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	RouteSelectionExpression pulumi.StringPtrOutput `pulumi:"routeSelectionExpression"`
	// The collection of tags. Each tag element is associated with a given resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// A version identifier for the API.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		args = &ApiArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"protocolType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("aws-native:apigatewayv2:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("aws-native:apigatewayv2:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string `pulumi:"apiKeySelectionExpression"`
	// Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
	BasePath *string `pulumi:"basePath"`
	// The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::Api` for more information about the expected schema for this property.
	Body interface{} `pulumi:"body"`
	// The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	BodyS3Location *ApiBodyS3Location `pulumi:"bodyS3Location"`
	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	CorsConfiguration *ApiCors `pulumi:"corsConfiguration"`
	// This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
	CredentialsArn *string `pulumi:"credentialsArn"`
	// The description of the API.
	Description *string `pulumi:"description"`
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint *bool `pulumi:"disableExecuteApiEndpoint"`
	// Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
	DisableSchemaValidation *bool `pulumi:"disableSchemaValidation"`
	// Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.
	FailOnWarnings *bool `pulumi:"failOnWarnings"`
	// The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.
	//
	// Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
	IpAddressType *string `pulumi:"ipAddressType"`
	// The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	Name *string `pulumi:"name"`
	// The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	ProtocolType *string `pulumi:"protocolType"`
	// This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
	RouteKey *string `pulumi:"routeKey"`
	// The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	RouteSelectionExpression *string `pulumi:"routeSelectionExpression"`
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `pulumi:"tags"`
	// This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	Target *string `pulumi:"target"`
	// A version identifier for the API.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression pulumi.StringPtrInput
	// Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
	BasePath pulumi.StringPtrInput
	// The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::Api` for more information about the expected schema for this property.
	Body pulumi.Input
	// The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	BodyS3Location ApiBodyS3LocationPtrInput
	// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	CorsConfiguration ApiCorsPtrInput
	// This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
	CredentialsArn pulumi.StringPtrInput
	// The description of the API.
	Description pulumi.StringPtrInput
	// Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	DisableExecuteApiEndpoint pulumi.BoolPtrInput
	// Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
	DisableSchemaValidation pulumi.BoolPtrInput
	// Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.
	FailOnWarnings pulumi.BoolPtrInput
	// The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.
	//
	// Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
	IpAddressType pulumi.StringPtrInput
	// The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	Name pulumi.StringPtrInput
	// The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
	ProtocolType pulumi.StringPtrInput
	// This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
	RouteKey pulumi.StringPtrInput
	// The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	RouteSelectionExpression pulumi.StringPtrInput
	// The collection of tags. Each tag element is associated with a given resource.
	Tags pulumi.StringMapInput
	// This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	Target pulumi.StringPtrInput
	// A version identifier for the API.
	Version pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// The default endpoint for an API. For example: `https://abcdef.execute-api.us-west-2.amazonaws.com` .
func (o ApiOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// The API identifier.
func (o ApiOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
func (o ApiOutput) ApiKeySelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiKeySelectionExpression }).(pulumi.StringPtrOutput)
}

// Specifies how to interpret the base path of the API during import. Valid values are “ignore“, “prepend“, and “split“. The default value is “ignore“. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
func (o ApiOutput) BasePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.BasePath }).(pulumi.StringPtrOutput)
}

// The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a “Body“ or “BodyS3Location“. If you specify a “Body“ or “BodyS3Location“, don't specify CloudFormation resources such as “AWS::ApiGatewayV2::Authorizer“ or “AWS::ApiGatewayV2::Route“. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ApiGatewayV2::Api` for more information about the expected schema for this property.
func (o ApiOutput) Body() pulumi.AnyOutput {
	return o.ApplyT(func(v *Api) pulumi.AnyOutput { return v.Body }).(pulumi.AnyOutput)
}

// The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a “Body“ or “BodyS3Location“. If you specify a “Body“ or “BodyS3Location“, don't specify CloudFormation resources such as “AWS::ApiGatewayV2::Authorizer“ or “AWS::ApiGatewayV2::Route“. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
func (o ApiOutput) BodyS3Location() ApiBodyS3LocationPtrOutput {
	return o.ApplyT(func(v *Api) ApiBodyS3LocationPtrOutput { return v.BodyS3Location }).(ApiBodyS3LocationPtrOutput)
}

// A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
func (o ApiOutput) CorsConfiguration() ApiCorsPtrOutput {
	return o.ApplyT(func(v *Api) ApiCorsPtrOutput { return v.CorsConfiguration }).(ApiCorsPtrOutput)
}

// This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify “arn:aws:iam::*:user/*“. To use resource-based permissions on supported AWS services, specify “null“. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
func (o ApiOutput) CredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.CredentialsArn }).(pulumi.StringPtrOutput)
}

// The description of the API.
func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether clients can invoke your API by using the default “execute-api“ endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
func (o ApiOutput) DisableExecuteApiEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.DisableExecuteApiEndpoint }).(pulumi.BoolPtrOutput)
}

// Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
func (o ApiOutput) DisableSchemaValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.DisableSchemaValidation }).(pulumi.BoolPtrOutput)
}

// Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.
func (o ApiOutput) FailOnWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.FailOnWarnings }).(pulumi.BoolPtrOutput)
}

// The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.
//
// Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
func (o ApiOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

// The name of the API. Required unless you specify an OpenAPI definition for “Body“ or “S3BodyLocation“.
func (o ApiOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The API protocol. Valid values are “WEBSOCKET“ or “HTTP“. Required unless you specify an OpenAPI definition for “Body“ or “S3BodyLocation“.
func (o ApiOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

// This property is part of quick create. If you don't specify a “routeKey“, a default route of “$default“ is created. The “$default“ route acts as a catch-all for any request made to your API, for a particular stage. The “$default“ route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
func (o ApiOutput) RouteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RouteKey }).(pulumi.StringPtrOutput)
}

// The route selection expression for the API. For HTTP APIs, the “routeSelectionExpression“ must be “${request.method} ${request.path}“. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
func (o ApiOutput) RouteSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.RouteSelectionExpression }).(pulumi.StringPtrOutput)
}

// The collection of tags. Each tag element is associated with a given resource.
func (o ApiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
func (o ApiOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

// A version identifier for the API.
func (o ApiOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterOutputType(ApiOutput{})
}
