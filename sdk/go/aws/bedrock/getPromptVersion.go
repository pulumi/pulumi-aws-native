// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::PromptVersion Resource Type
func LookupPromptVersion(ctx *pulumi.Context, args *LookupPromptVersionArgs, opts ...pulumi.InvokeOption) (*LookupPromptVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPromptVersionResult
	err := ctx.Invoke("aws-native:bedrock:getPromptVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPromptVersionArgs struct {
	// ARN of a prompt version resource
	Arn string `pulumi:"arn"`
}

type LookupPromptVersionResult struct {
	// ARN of a prompt version resource
	Arn *string `pulumi:"arn"`
	// Time Stamp.
	CreatedAt *string `pulumi:"createdAt"`
	// Name for a variant.
	DefaultVariant *string `pulumi:"defaultVariant"`
	// Name for a prompt resource.
	Name *string `pulumi:"name"`
	// Identifier for a Prompt
	PromptId *string `pulumi:"promptId"`
	// Time Stamp.
	UpdatedAt *string `pulumi:"updatedAt"`
	// List of prompt variants
	Variants []PromptVersionPromptVariant `pulumi:"variants"`
	// Version.
	Version *string `pulumi:"version"`
}

func LookupPromptVersionOutput(ctx *pulumi.Context, args LookupPromptVersionOutputArgs, opts ...pulumi.InvokeOption) LookupPromptVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPromptVersionResult, error) {
			args := v.(LookupPromptVersionArgs)
			r, err := LookupPromptVersion(ctx, &args, opts...)
			var s LookupPromptVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPromptVersionResultOutput)
}

type LookupPromptVersionOutputArgs struct {
	// ARN of a prompt version resource
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupPromptVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPromptVersionArgs)(nil)).Elem()
}

type LookupPromptVersionResultOutput struct{ *pulumi.OutputState }

func (LookupPromptVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPromptVersionResult)(nil)).Elem()
}

func (o LookupPromptVersionResultOutput) ToLookupPromptVersionResultOutput() LookupPromptVersionResultOutput {
	return o
}

func (o LookupPromptVersionResultOutput) ToLookupPromptVersionResultOutputWithContext(ctx context.Context) LookupPromptVersionResultOutput {
	return o
}

// ARN of a prompt version resource
func (o LookupPromptVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Time Stamp.
func (o LookupPromptVersionResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Name for a variant.
func (o LookupPromptVersionResultOutput) DefaultVariant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.DefaultVariant }).(pulumi.StringPtrOutput)
}

// Name for a prompt resource.
func (o LookupPromptVersionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Identifier for a Prompt
func (o LookupPromptVersionResultOutput) PromptId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.PromptId }).(pulumi.StringPtrOutput)
}

// Time Stamp.
func (o LookupPromptVersionResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

// List of prompt variants
func (o LookupPromptVersionResultOutput) Variants() PromptVersionPromptVariantArrayOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) []PromptVersionPromptVariant { return v.Variants }).(PromptVersionPromptVariantArrayOutput)
}

// Version.
func (o LookupPromptVersionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPromptVersionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPromptVersionResultOutput{})
}