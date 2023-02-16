// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
func LookupWebACL(ctx *pulumi.Context, args *LookupWebACLArgs, opts ...pulumi.InvokeOption) (*LookupWebACLResult, error) {
	var rv LookupWebACLResult
	err := ctx.Invoke("aws-native:wafv2:getWebACL", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebACLArgs struct {
	Id    string      `pulumi:"id"`
	Name  string      `pulumi:"name"`
	Scope WebACLScope `pulumi:"scope"`
}

type LookupWebACLResult struct {
	Arn                  *string                     `pulumi:"arn"`
	Capacity             *int                        `pulumi:"capacity"`
	CaptchaConfig        *WebACLCaptchaConfig        `pulumi:"captchaConfig"`
	ChallengeConfig      *WebACLChallengeConfig      `pulumi:"challengeConfig"`
	CustomResponseBodies *WebACLCustomResponseBodies `pulumi:"customResponseBodies"`
	DefaultAction        *WebACLDefaultAction        `pulumi:"defaultAction"`
	Description          *string                     `pulumi:"description"`
	Id                   *string                     `pulumi:"id"`
	LabelNamespace       *string                     `pulumi:"labelNamespace"`
	// Collection of Rules.
	Rules            []WebACLRule            `pulumi:"rules"`
	Tags             []WebACLTag             `pulumi:"tags"`
	TokenDomains     []string                `pulumi:"tokenDomains"`
	VisibilityConfig *WebACLVisibilityConfig `pulumi:"visibilityConfig"`
}

func LookupWebACLOutput(ctx *pulumi.Context, args LookupWebACLOutputArgs, opts ...pulumi.InvokeOption) LookupWebACLResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebACLResult, error) {
			args := v.(LookupWebACLArgs)
			r, err := LookupWebACL(ctx, &args, opts...)
			var s LookupWebACLResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebACLResultOutput)
}

type LookupWebACLOutputArgs struct {
	Id    pulumi.StringInput `pulumi:"id"`
	Name  pulumi.StringInput `pulumi:"name"`
	Scope WebACLScopeInput   `pulumi:"scope"`
}

func (LookupWebACLOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebACLArgs)(nil)).Elem()
}

type LookupWebACLResultOutput struct{ *pulumi.OutputState }

func (LookupWebACLResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebACLResult)(nil)).Elem()
}

func (o LookupWebACLResultOutput) ToLookupWebACLResultOutput() LookupWebACLResultOutput {
	return o
}

func (o LookupWebACLResultOutput) ToLookupWebACLResultOutputWithContext(ctx context.Context) LookupWebACLResultOutput {
	return o
}

func (o LookupWebACLResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupWebACLResultOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o LookupWebACLResultOutput) CaptchaConfig() WebACLCaptchaConfigPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *WebACLCaptchaConfig { return v.CaptchaConfig }).(WebACLCaptchaConfigPtrOutput)
}

func (o LookupWebACLResultOutput) ChallengeConfig() WebACLChallengeConfigPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *WebACLChallengeConfig { return v.ChallengeConfig }).(WebACLChallengeConfigPtrOutput)
}

func (o LookupWebACLResultOutput) CustomResponseBodies() WebACLCustomResponseBodiesPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *WebACLCustomResponseBodies { return v.CustomResponseBodies }).(WebACLCustomResponseBodiesPtrOutput)
}

func (o LookupWebACLResultOutput) DefaultAction() WebACLDefaultActionPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *WebACLDefaultAction { return v.DefaultAction }).(WebACLDefaultActionPtrOutput)
}

func (o LookupWebACLResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebACLResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWebACLResultOutput) LabelNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *string { return v.LabelNamespace }).(pulumi.StringPtrOutput)
}

// Collection of Rules.
func (o LookupWebACLResultOutput) Rules() WebACLRuleArrayOutput {
	return o.ApplyT(func(v LookupWebACLResult) []WebACLRule { return v.Rules }).(WebACLRuleArrayOutput)
}

func (o LookupWebACLResultOutput) Tags() WebACLTagArrayOutput {
	return o.ApplyT(func(v LookupWebACLResult) []WebACLTag { return v.Tags }).(WebACLTagArrayOutput)
}

func (o LookupWebACLResultOutput) TokenDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebACLResult) []string { return v.TokenDomains }).(pulumi.StringArrayOutput)
}

func (o LookupWebACLResultOutput) VisibilityConfig() WebACLVisibilityConfigPtrOutput {
	return o.ApplyT(func(v LookupWebACLResult) *WebACLVisibilityConfig { return v.VisibilityConfig }).(WebACLVisibilityConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebACLResultOutput{})
}