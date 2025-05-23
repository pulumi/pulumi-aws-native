// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Extension
func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupExtensionResult
	err := ctx.Invoke("aws-native:appconfig:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	// The system-generated ID of the extension.
	Id string `pulumi:"id"`
}

type LookupExtensionResult struct {
	// The actions defined in the extension.
	Actions map[string][]ExtensionAction `pulumi:"actions"`
	// The system-generated Amazon Resource Name (ARN) for the extension.
	Arn *string `pulumi:"arn"`
	// Description of the extension.
	Description *string `pulumi:"description"`
	// The system-generated ID of the extension.
	Id *string `pulumi:"id"`
	// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AWS AppConfig resource by using the `CreateExtensionAssociation` API action. For AWS Lambda extension actions, these parameters are included in the Lambda request object.
	Parameters map[string]ExtensionParameter `pulumi:"parameters"`
	// An array of key-value tags to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The extension version number.
	VersionNumber *int `pulumi:"versionNumber"`
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExtensionResultOutput, error) {
			args := v.(LookupExtensionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:appconfig:getExtension", args, LookupExtensionResultOutput{}, options).(LookupExtensionResultOutput), nil
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	// The system-generated ID of the extension.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}

type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

// The actions defined in the extension.
func (o LookupExtensionResultOutput) Actions() ExtensionActionArrayMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string][]ExtensionAction { return v.Actions }).(ExtensionActionArrayMapOutput)
}

// The system-generated Amazon Resource Name (ARN) for the extension.
func (o LookupExtensionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Description of the extension.
func (o LookupExtensionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The system-generated ID of the extension.
func (o LookupExtensionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AWS AppConfig resource by using the `CreateExtensionAssociation` API action. For AWS Lambda extension actions, these parameters are included in the Lambda request object.
func (o LookupExtensionResultOutput) Parameters() ExtensionParameterMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]ExtensionParameter { return v.Parameters }).(ExtensionParameterMapOutput)
}

// An array of key-value tags to apply to this resource.
func (o LookupExtensionResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupExtensionResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The extension version number.
func (o LookupExtensionResultOutput) VersionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *int { return v.VersionNumber }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
