// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::Route
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("aws-native:apigatewayv2:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	ApiId   string `pulumi:"apiId"`
	RouteId string `pulumi:"routeId"`
}

type LookupRouteResult struct {
	ApiKeyRequired                   *bool       `pulumi:"apiKeyRequired"`
	AuthorizationScopes              []string    `pulumi:"authorizationScopes"`
	AuthorizationType                *string     `pulumi:"authorizationType"`
	ModelSelectionExpression         *string     `pulumi:"modelSelectionExpression"`
	OperationName                    *string     `pulumi:"operationName"`
	RequestModels                    interface{} `pulumi:"requestModels"`
	RouteId                          *string     `pulumi:"routeId"`
	RouteKey                         *string     `pulumi:"routeKey"`
	RouteResponseSelectionExpression *string     `pulumi:"routeResponseSelectionExpression"`
	Target                           *string     `pulumi:"target"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	ApiId   pulumi.StringInput `pulumi:"apiId"`
	RouteId pulumi.StringInput `pulumi:"routeId"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ApiKeyRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *bool { return v.ApiKeyRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupRouteResultOutput) AuthorizationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.AuthorizationScopes }).(pulumi.StringArrayOutput)
}

func (o LookupRouteResultOutput) AuthorizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.AuthorizationType }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) ModelSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.ModelSelectionExpression }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) RequestModels() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRouteResult) interface{} { return v.RequestModels }).(pulumi.AnyOutput)
}

func (o LookupRouteResultOutput) RouteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.RouteId }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) RouteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.RouteKey }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) RouteResponseSelectionExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.RouteResponseSelectionExpression }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}