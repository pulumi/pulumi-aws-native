// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::FunctionConfiguration
func LookupFunctionConfiguration(ctx *pulumi.Context, args *LookupFunctionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupFunctionConfigurationResult, error) {
	var rv LookupFunctionConfigurationResult
	err := ctx.Invoke("aws-native:appsync:getFunctionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFunctionConfigurationArgs struct {
	Id string `pulumi:"id"`
}

type LookupFunctionConfigurationResult struct {
	Code                              *string                              `pulumi:"code"`
	CodeS3Location                    *string                              `pulumi:"codeS3Location"`
	DataSourceName                    *string                              `pulumi:"dataSourceName"`
	Description                       *string                              `pulumi:"description"`
	FunctionArn                       *string                              `pulumi:"functionArn"`
	FunctionId                        *string                              `pulumi:"functionId"`
	FunctionVersion                   *string                              `pulumi:"functionVersion"`
	Id                                *string                              `pulumi:"id"`
	MaxBatchSize                      *int                                 `pulumi:"maxBatchSize"`
	Name                              *string                              `pulumi:"name"`
	RequestMappingTemplate            *string                              `pulumi:"requestMappingTemplate"`
	RequestMappingTemplateS3Location  *string                              `pulumi:"requestMappingTemplateS3Location"`
	ResponseMappingTemplate           *string                              `pulumi:"responseMappingTemplate"`
	ResponseMappingTemplateS3Location *string                              `pulumi:"responseMappingTemplateS3Location"`
	Runtime                           *FunctionConfigurationAppSyncRuntime `pulumi:"runtime"`
	SyncConfig                        *FunctionConfigurationSyncConfig     `pulumi:"syncConfig"`
}

func LookupFunctionConfigurationOutput(ctx *pulumi.Context, args LookupFunctionConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupFunctionConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFunctionConfigurationResult, error) {
			args := v.(LookupFunctionConfigurationArgs)
			r, err := LookupFunctionConfiguration(ctx, &args, opts...)
			var s LookupFunctionConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFunctionConfigurationResultOutput)
}

type LookupFunctionConfigurationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupFunctionConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionConfigurationArgs)(nil)).Elem()
}

type LookupFunctionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupFunctionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFunctionConfigurationResult)(nil)).Elem()
}

func (o LookupFunctionConfigurationResultOutput) ToLookupFunctionConfigurationResultOutput() LookupFunctionConfigurationResultOutput {
	return o
}

func (o LookupFunctionConfigurationResultOutput) ToLookupFunctionConfigurationResultOutputWithContext(ctx context.Context) LookupFunctionConfigurationResultOutput {
	return o
}

func (o LookupFunctionConfigurationResultOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) CodeS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.CodeS3Location }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) DataSourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.DataSourceName }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) FunctionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionArn }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) FunctionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionId }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) FunctionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.FunctionVersion }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) MaxBatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *int { return v.MaxBatchSize }).(pulumi.IntPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) RequestMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.RequestMappingTemplate }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) RequestMappingTemplateS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.RequestMappingTemplateS3Location }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) ResponseMappingTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.ResponseMappingTemplate }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) ResponseMappingTemplateS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *string { return v.ResponseMappingTemplateS3Location }).(pulumi.StringPtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) Runtime() FunctionConfigurationAppSyncRuntimePtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *FunctionConfigurationAppSyncRuntime { return v.Runtime }).(FunctionConfigurationAppSyncRuntimePtrOutput)
}

func (o LookupFunctionConfigurationResultOutput) SyncConfig() FunctionConfigurationSyncConfigPtrOutput {
	return o.ApplyT(func(v LookupFunctionConfigurationResult) *FunctionConfigurationSyncConfig { return v.SyncConfig }).(FunctionConfigurationSyncConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFunctionConfigurationResultOutput{})
}