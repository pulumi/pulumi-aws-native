// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::RestApi
func LookupRestApi(ctx *pulumi.Context, args *LookupRestApiArgs, opts ...pulumi.InvokeOption) (*LookupRestApiResult, error) {
	var rv LookupRestApiResult
	err := ctx.Invoke("aws-native:apigateway:getRestApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestApiArgs struct {
	RestApiId string `pulumi:"restApiId"`
}

type LookupRestApiResult struct {
	ApiKeySourceType          *string                       `pulumi:"apiKeySourceType"`
	BinaryMediaTypes          []string                      `pulumi:"binaryMediaTypes"`
	Body                      interface{}                   `pulumi:"body"`
	BodyS3Location            *RestApiS3Location            `pulumi:"bodyS3Location"`
	CloneFrom                 *string                       `pulumi:"cloneFrom"`
	Description               *string                       `pulumi:"description"`
	DisableExecuteApiEndpoint *bool                         `pulumi:"disableExecuteApiEndpoint"`
	EndpointConfiguration     *RestApiEndpointConfiguration `pulumi:"endpointConfiguration"`
	FailOnWarnings            *bool                         `pulumi:"failOnWarnings"`
	MinimumCompressionSize    *int                          `pulumi:"minimumCompressionSize"`
	Mode                      *string                       `pulumi:"mode"`
	Name                      *string                       `pulumi:"name"`
	Parameters                interface{}                   `pulumi:"parameters"`
	Policy                    interface{}                   `pulumi:"policy"`
	RestApiId                 *string                       `pulumi:"restApiId"`
	RootResourceId            *string                       `pulumi:"rootResourceId"`
	Tags                      []RestApiTag                  `pulumi:"tags"`
}

func LookupRestApiOutput(ctx *pulumi.Context, args LookupRestApiOutputArgs, opts ...pulumi.InvokeOption) LookupRestApiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestApiResult, error) {
			args := v.(LookupRestApiArgs)
			r, err := LookupRestApi(ctx, &args, opts...)
			var s LookupRestApiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestApiResultOutput)
}

type LookupRestApiOutputArgs struct {
	RestApiId pulumi.StringInput `pulumi:"restApiId"`
}

func (LookupRestApiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiArgs)(nil)).Elem()
}

type LookupRestApiResultOutput struct{ *pulumi.OutputState }

func (LookupRestApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiResult)(nil)).Elem()
}

func (o LookupRestApiResultOutput) ToLookupRestApiResultOutput() LookupRestApiResultOutput {
	return o
}

func (o LookupRestApiResultOutput) ToLookupRestApiResultOutputWithContext(ctx context.Context) LookupRestApiResultOutput {
	return o
}

func (o LookupRestApiResultOutput) ApiKeySourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.ApiKeySourceType }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) BinaryMediaTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRestApiResult) []string { return v.BinaryMediaTypes }).(pulumi.StringArrayOutput)
}

func (o LookupRestApiResultOutput) Body() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRestApiResult) interface{} { return v.Body }).(pulumi.AnyOutput)
}

func (o LookupRestApiResultOutput) BodyS3Location() RestApiS3LocationPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *RestApiS3Location { return v.BodyS3Location }).(RestApiS3LocationPtrOutput)
}

func (o LookupRestApiResultOutput) CloneFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.CloneFrom }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) DisableExecuteApiEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *bool { return v.DisableExecuteApiEndpoint }).(pulumi.BoolPtrOutput)
}

func (o LookupRestApiResultOutput) EndpointConfiguration() RestApiEndpointConfigurationPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *RestApiEndpointConfiguration { return v.EndpointConfiguration }).(RestApiEndpointConfigurationPtrOutput)
}

func (o LookupRestApiResultOutput) FailOnWarnings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *bool { return v.FailOnWarnings }).(pulumi.BoolPtrOutput)
}

func (o LookupRestApiResultOutput) MinimumCompressionSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *int { return v.MinimumCompressionSize }).(pulumi.IntPtrOutput)
}

func (o LookupRestApiResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRestApiResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupRestApiResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRestApiResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

func (o LookupRestApiResultOutput) RestApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.RestApiId }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) RootResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestApiResult) *string { return v.RootResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupRestApiResultOutput) Tags() RestApiTagArrayOutput {
	return o.ApplyT(func(v LookupRestApiResult) []RestApiTag { return v.Tags }).(RestApiTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestApiResultOutput{})
}