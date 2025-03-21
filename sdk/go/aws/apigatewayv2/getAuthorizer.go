// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::ApiGatewayV2::Authorizer“ resource creates an authorizer for a WebSocket API or an HTTP API. To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) and [Controlling and managing access to an HTTP API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) in the *API Gateway Developer Guide*.
func LookupAuthorizer(ctx *pulumi.Context, args *LookupAuthorizerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizerResult
	err := ctx.Invoke("aws-native:apigatewayv2:getAuthorizer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizerArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The authorizer ID.
	AuthorizerId string `pulumi:"authorizerId"`
}

type LookupAuthorizerResult struct {
	// Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers.
	AuthorizerCredentialsArn *string `pulumi:"authorizerCredentialsArn"`
	// The authorizer ID.
	AuthorizerId *string `pulumi:"authorizerId"`
	// Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	AuthorizerPayloadFormatVersion *string `pulumi:"authorizerPayloadFormatVersion"`
	// The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
	AuthorizerResultTtlInSeconds *int `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs).
	AuthorizerType *string `pulumi:"authorizerType"`
	// The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``.
	AuthorizerUri *string `pulumi:"authorizerUri"`
	// Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	EnableSimpleResponses *bool `pulumi:"enableSimpleResponses"`
	// The identity source for which authorization is requested.
	//  For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
	//  For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
	IdentitySource []string `pulumi:"identitySource"`
	// This parameter is not used.
	IdentityValidationExpression *string `pulumi:"identityValidationExpression"`
	// The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs.
	JwtConfiguration *AuthorizerJwtConfiguration `pulumi:"jwtConfiguration"`
	// The name of the authorizer.
	Name *string `pulumi:"name"`
}

func LookupAuthorizerOutput(ctx *pulumi.Context, args LookupAuthorizerOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAuthorizerResultOutput, error) {
			args := v.(LookupAuthorizerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:apigatewayv2:getAuthorizer", args, LookupAuthorizerResultOutput{}, options).(LookupAuthorizerResultOutput), nil
		}).(LookupAuthorizerResultOutput)
}

type LookupAuthorizerOutputArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The authorizer ID.
	AuthorizerId pulumi.StringInput `pulumi:"authorizerId"`
}

func (LookupAuthorizerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizerArgs)(nil)).Elem()
}

type LookupAuthorizerResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizerResult)(nil)).Elem()
}

func (o LookupAuthorizerResultOutput) ToLookupAuthorizerResultOutput() LookupAuthorizerResultOutput {
	return o
}

func (o LookupAuthorizerResultOutput) ToLookupAuthorizerResultOutputWithContext(ctx context.Context) LookupAuthorizerResultOutput {
	return o
}

// Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for “REQUEST“ authorizers.
func (o LookupAuthorizerResultOutput) AuthorizerCredentialsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerCredentialsArn }).(pulumi.StringPtrOutput)
}

// The authorizer ID.
func (o LookupAuthorizerResultOutput) AuthorizerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerId }).(pulumi.StringPtrOutput)
}

// Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are “1.0“ and “2.0“. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
func (o LookupAuthorizerResultOutput) AuthorizerPayloadFormatVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerPayloadFormatVersion }).(pulumi.StringPtrOutput)
}

// The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
func (o LookupAuthorizerResultOutput) AuthorizerResultTtlInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *int { return v.AuthorizerResultTtlInSeconds }).(pulumi.IntPtrOutput)
}

// The authorizer type. Specify “REQUEST“ for a Lambda function using incoming request parameters. Specify “JWT“ to use JSON Web Tokens (supported only for HTTP APIs).
func (o LookupAuthorizerResultOutput) AuthorizerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerType }).(pulumi.StringPtrOutput)
}

// The authorizer's Uniform Resource Identifier (URI). For “REQUEST“ authorizers, this must be a well-formed Lambda function URI, for example, “arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations“. In general, the URI has this form: “arn:aws:apigateway:{region}:lambda:path/{service_api}“, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial “/“. For Lambda functions, this is usually of the form “/2015-03-31/functions/[FunctionARN]/invocations“.
func (o LookupAuthorizerResultOutput) AuthorizerUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.AuthorizerUri }).(pulumi.StringPtrOutput)
}

// Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
func (o LookupAuthorizerResultOutput) EnableSimpleResponses() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *bool { return v.EnableSimpleResponses }).(pulumi.BoolPtrOutput)
}

// The identity source for which authorization is requested.
//
//	For a ``REQUEST`` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with ``$``, for example, ``$request.header.Auth``, ``$request.querystring.Name``. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).
//	For ``JWT``, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example ``$request.header.Authorization``.
func (o LookupAuthorizerResultOutput) IdentitySource() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) []string { return v.IdentitySource }).(pulumi.StringArrayOutput)
}

// This parameter is not used.
func (o LookupAuthorizerResultOutput) IdentityValidationExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.IdentityValidationExpression }).(pulumi.StringPtrOutput)
}

// The “JWTConfiguration“ property specifies the configuration of a JWT authorizer. Required for the “JWT“ authorizer type. Supported only for HTTP APIs.
func (o LookupAuthorizerResultOutput) JwtConfiguration() AuthorizerJwtConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *AuthorizerJwtConfiguration { return v.JwtConfiguration }).(AuthorizerJwtConfigurationPtrOutput)
}

// The name of the authorizer.
func (o LookupAuthorizerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizerResultOutput{})
}
