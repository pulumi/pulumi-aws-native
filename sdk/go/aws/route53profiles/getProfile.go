// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53profiles

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Route53Profiles::Profile
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResult
	err := ctx.Invoke("aws-native:route53profiles:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	// The ID of the profile.
	Id string `pulumi:"id"`
}

type LookupProfileResult struct {
	// The Amazon Resource Name (ARN) of the resolver profile.
	Arn *string `pulumi:"arn"`
	// The id of the creator request
	ClientToken *string `pulumi:"clientToken"`
	// The ID of the profile.
	Id *string `pulumi:"id"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProfileResultOutput, error) {
			args := v.(LookupProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:route53profiles:getProfile", args, LookupProfileResultOutput{}, options).(LookupProfileResultOutput), nil
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	// The ID of the profile.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the resolver profile.
func (o LookupProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The id of the creator request
func (o LookupProfileResultOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ClientToken }).(pulumi.StringPtrOutput)
}

// The ID of the profile.
func (o LookupProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupProfileResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
