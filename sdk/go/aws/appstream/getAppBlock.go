// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::AppBlock
func LookupAppBlock(ctx *pulumi.Context, args *LookupAppBlockArgs, opts ...pulumi.InvokeOption) (*LookupAppBlockResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppBlockResult
	err := ctx.Invoke("aws-native:appstream:getAppBlock", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppBlockArgs struct {
	// The ARN of the app block.
	Arn string `pulumi:"arn"`
}

type LookupAppBlockResult struct {
	// The ARN of the app block.
	Arn *string `pulumi:"arn"`
	// The time when the app block was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The tags of the app block.
	Tags []interface{} `pulumi:"tags"`
}

func LookupAppBlockOutput(ctx *pulumi.Context, args LookupAppBlockOutputArgs, opts ...pulumi.InvokeOption) LookupAppBlockResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAppBlockResultOutput, error) {
			args := v.(LookupAppBlockArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:appstream:getAppBlock", args, LookupAppBlockResultOutput{}, options).(LookupAppBlockResultOutput), nil
		}).(LookupAppBlockResultOutput)
}

type LookupAppBlockOutputArgs struct {
	// The ARN of the app block.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupAppBlockOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppBlockArgs)(nil)).Elem()
}

type LookupAppBlockResultOutput struct{ *pulumi.OutputState }

func (LookupAppBlockResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppBlockResult)(nil)).Elem()
}

func (o LookupAppBlockResultOutput) ToLookupAppBlockResultOutput() LookupAppBlockResultOutput {
	return o
}

func (o LookupAppBlockResultOutput) ToLookupAppBlockResultOutputWithContext(ctx context.Context) LookupAppBlockResultOutput {
	return o
}

// The ARN of the app block.
func (o LookupAppBlockResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppBlockResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The time when the app block was created.
func (o LookupAppBlockResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppBlockResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The tags of the app block.
func (o LookupAppBlockResultOutput) Tags() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupAppBlockResult) []interface{} { return v.Tags }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppBlockResultOutput{})
}
