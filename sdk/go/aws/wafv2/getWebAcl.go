// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
func LookupWebAcl(ctx *pulumi.Context, args *LookupWebAclArgs, opts ...pulumi.InvokeOption) (*LookupWebAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAclResult
	err := ctx.Invoke("aws-native:wafv2:getWebAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAclArgs struct {
	Id    string      `pulumi:"id"`
	Name  string      `pulumi:"name"`
	Scope WebAclScope `pulumi:"scope"`
}

type LookupWebAclResult struct {
	Arn                  *string                     `pulumi:"arn"`
	AssociationConfig    *WebAclAssociationConfig    `pulumi:"associationConfig"`
	Capacity             *int                        `pulumi:"capacity"`
	CaptchaConfig        *WebAclCaptchaConfig        `pulumi:"captchaConfig"`
	ChallengeConfig      *WebAclChallengeConfig      `pulumi:"challengeConfig"`
	CustomResponseBodies *WebAclCustomResponseBodies `pulumi:"customResponseBodies"`
	DefaultAction        *WebAclDefaultAction        `pulumi:"defaultAction"`
	Description          *string                     `pulumi:"description"`
	Id                   *string                     `pulumi:"id"`
	LabelNamespace       *string                     `pulumi:"labelNamespace"`
	// Collection of Rules.
	Rules            []WebAclRule            `pulumi:"rules"`
	Tags             []WebAclTag             `pulumi:"tags"`
	TokenDomains     []string                `pulumi:"tokenDomains"`
	VisibilityConfig *WebAclVisibilityConfig `pulumi:"visibilityConfig"`
}

func LookupWebAclOutput(ctx *pulumi.Context, args LookupWebAclOutputArgs, opts ...pulumi.InvokeOption) LookupWebAclResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAclResult, error) {
			args := v.(LookupWebAclArgs)
			r, err := LookupWebAcl(ctx, &args, opts...)
			var s LookupWebAclResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAclResultOutput)
}

type LookupWebAclOutputArgs struct {
	Id    pulumi.StringInput `pulumi:"id"`
	Name  pulumi.StringInput `pulumi:"name"`
	Scope WebAclScopeInput   `pulumi:"scope"`
}

func (LookupWebAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclArgs)(nil)).Elem()
}

type LookupWebAclResultOutput struct{ *pulumi.OutputState }

func (LookupWebAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclResult)(nil)).Elem()
}

func (o LookupWebAclResultOutput) ToLookupWebAclResultOutput() LookupWebAclResultOutput {
	return o
}

func (o LookupWebAclResultOutput) ToLookupWebAclResultOutputWithContext(ctx context.Context) LookupWebAclResultOutput {
	return o
}

func (o LookupWebAclResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupWebAclResultOutput) AssociationConfig() WebAclAssociationConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclAssociationConfig { return v.AssociationConfig }).(WebAclAssociationConfigPtrOutput)
}

func (o LookupWebAclResultOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o LookupWebAclResultOutput) CaptchaConfig() WebAclCaptchaConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclCaptchaConfig { return v.CaptchaConfig }).(WebAclCaptchaConfigPtrOutput)
}

func (o LookupWebAclResultOutput) ChallengeConfig() WebAclChallengeConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclChallengeConfig { return v.ChallengeConfig }).(WebAclChallengeConfigPtrOutput)
}

func (o LookupWebAclResultOutput) CustomResponseBodies() WebAclCustomResponseBodiesPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclCustomResponseBodies { return v.CustomResponseBodies }).(WebAclCustomResponseBodiesPtrOutput)
}

func (o LookupWebAclResultOutput) DefaultAction() WebAclDefaultActionPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclDefaultAction { return v.DefaultAction }).(WebAclDefaultActionPtrOutput)
}

func (o LookupWebAclResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupWebAclResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWebAclResultOutput) LabelNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.LabelNamespace }).(pulumi.StringPtrOutput)
}

// Collection of Rules.
func (o LookupWebAclResultOutput) Rules() WebAclRuleArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []WebAclRule { return v.Rules }).(WebAclRuleArrayOutput)
}

func (o LookupWebAclResultOutput) Tags() WebAclTagArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []WebAclTag { return v.Tags }).(WebAclTagArrayOutput)
}

func (o LookupWebAclResultOutput) TokenDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []string { return v.TokenDomains }).(pulumi.StringArrayOutput)
}

func (o LookupWebAclResultOutput) VisibilityConfig() WebAclVisibilityConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclVisibilityConfig { return v.VisibilityConfig }).(WebAclVisibilityConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAclResultOutput{})
}