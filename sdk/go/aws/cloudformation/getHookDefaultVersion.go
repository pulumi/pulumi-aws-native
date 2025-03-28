// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Set a version as default version for a hook in CloudFormation Registry.
func LookupHookDefaultVersion(ctx *pulumi.Context, args *LookupHookDefaultVersionArgs, opts ...pulumi.InvokeOption) (*LookupHookDefaultVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHookDefaultVersionResult
	err := ctx.Invoke("aws-native:cloudformation:getHookDefaultVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHookDefaultVersionArgs struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
	Arn string `pulumi:"arn"`
}

type LookupHookDefaultVersionResult struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
	Arn *string `pulumi:"arn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
	// The Amazon Resource Name (ARN) of the type version.
	TypeVersionArn *string `pulumi:"typeVersionArn"`
	// The ID of an existing version of the hook to set as the default.
	VersionId *string `pulumi:"versionId"`
}

func LookupHookDefaultVersionOutput(ctx *pulumi.Context, args LookupHookDefaultVersionOutputArgs, opts ...pulumi.InvokeOption) LookupHookDefaultVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupHookDefaultVersionResultOutput, error) {
			args := v.(LookupHookDefaultVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudformation:getHookDefaultVersion", args, LookupHookDefaultVersionResultOutput{}, options).(LookupHookDefaultVersionResultOutput), nil
		}).(LookupHookDefaultVersionResultOutput)
}

type LookupHookDefaultVersionOutputArgs struct {
	// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupHookDefaultVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookDefaultVersionArgs)(nil)).Elem()
}

type LookupHookDefaultVersionResultOutput struct{ *pulumi.OutputState }

func (LookupHookDefaultVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHookDefaultVersionResult)(nil)).Elem()
}

func (o LookupHookDefaultVersionResultOutput) ToLookupHookDefaultVersionResultOutput() LookupHookDefaultVersionResultOutput {
	return o
}

func (o LookupHookDefaultVersionResultOutput) ToLookupHookDefaultVersionResultOutputWithContext(ctx context.Context) LookupHookDefaultVersionResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a HookDefaultVersion
func (o LookupHookDefaultVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookDefaultVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o LookupHookDefaultVersionResultOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookDefaultVersionResult) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the type version.
func (o LookupHookDefaultVersionResultOutput) TypeVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookDefaultVersionResult) *string { return v.TypeVersionArn }).(pulumi.StringPtrOutput)
}

// The ID of an existing version of the hook to set as the default.
func (o LookupHookDefaultVersionResultOutput) VersionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHookDefaultVersionResult) *string { return v.VersionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHookDefaultVersionResultOutput{})
}
