// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qbusiness

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::QBusiness::WebExperience Resource Type
func LookupWebExperience(ctx *pulumi.Context, args *LookupWebExperienceArgs, opts ...pulumi.InvokeOption) (*LookupWebExperienceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebExperienceResult
	err := ctx.Invoke("aws-native:qbusiness:getWebExperience", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebExperienceArgs struct {
	ApplicationId   string `pulumi:"applicationId"`
	WebExperienceId string `pulumi:"webExperienceId"`
}

type LookupWebExperienceResult struct {
	CreatedAt                *string                                `pulumi:"createdAt"`
	DefaultEndpoint          *string                                `pulumi:"defaultEndpoint"`
	RoleArn                  *string                                `pulumi:"roleArn"`
	SamplePromptsControlMode *WebExperienceSamplePromptsControlMode `pulumi:"samplePromptsControlMode"`
	Status                   *WebExperienceStatus                   `pulumi:"status"`
	Subtitle                 *string                                `pulumi:"subtitle"`
	Tags                     []aws.Tag                              `pulumi:"tags"`
	Title                    *string                                `pulumi:"title"`
	UpdatedAt                *string                                `pulumi:"updatedAt"`
	WebExperienceArn         *string                                `pulumi:"webExperienceArn"`
	WebExperienceId          *string                                `pulumi:"webExperienceId"`
	WelcomeMessage           *string                                `pulumi:"welcomeMessage"`
}

func LookupWebExperienceOutput(ctx *pulumi.Context, args LookupWebExperienceOutputArgs, opts ...pulumi.InvokeOption) LookupWebExperienceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebExperienceResult, error) {
			args := v.(LookupWebExperienceArgs)
			r, err := LookupWebExperience(ctx, &args, opts...)
			var s LookupWebExperienceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebExperienceResultOutput)
}

type LookupWebExperienceOutputArgs struct {
	ApplicationId   pulumi.StringInput `pulumi:"applicationId"`
	WebExperienceId pulumi.StringInput `pulumi:"webExperienceId"`
}

func (LookupWebExperienceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebExperienceArgs)(nil)).Elem()
}

type LookupWebExperienceResultOutput struct{ *pulumi.OutputState }

func (LookupWebExperienceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebExperienceResult)(nil)).Elem()
}

func (o LookupWebExperienceResultOutput) ToLookupWebExperienceResultOutput() LookupWebExperienceResultOutput {
	return o
}

func (o LookupWebExperienceResultOutput) ToLookupWebExperienceResultOutputWithContext(ctx context.Context) LookupWebExperienceResultOutput {
	return o
}

func (o LookupWebExperienceResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) DefaultEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.DefaultEndpoint }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) SamplePromptsControlMode() WebExperienceSamplePromptsControlModePtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *WebExperienceSamplePromptsControlMode {
		return v.SamplePromptsControlMode
	}).(WebExperienceSamplePromptsControlModePtrOutput)
}

func (o LookupWebExperienceResultOutput) Status() WebExperienceStatusPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *WebExperienceStatus { return v.Status }).(WebExperienceStatusPtrOutput)
}

func (o LookupWebExperienceResultOutput) Subtitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.Subtitle }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupWebExperienceResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) WebExperienceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.WebExperienceArn }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) WebExperienceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.WebExperienceId }).(pulumi.StringPtrOutput)
}

func (o LookupWebExperienceResultOutput) WelcomeMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebExperienceResult) *string { return v.WelcomeMessage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebExperienceResultOutput{})
}