// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
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
	// The ID of the web ACL.
	Id string `pulumi:"id"`
	// The name of the web ACL. You cannot change the name of a web ACL after you create it.
	Name string `pulumi:"name"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	Scope WebAclScope `pulumi:"scope"`
}

type LookupWebAclResult struct {
	// The Amazon Resource Name (ARN) of the web ACL.
	Arn *string `pulumi:"arn"`
	// Specifies custom configurations for the associations between the web ACL and protected resources.
	//
	// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
	//
	// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
	//
	// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
	AssociationConfig *WebAclAssociationConfig `pulumi:"associationConfig"`
	// The web ACL capacity units (WCUs) currently being used by this web ACL.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	Capacity *int `pulumi:"capacity"`
	// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
	CaptchaConfig *WebAclCaptchaConfig `pulumi:"captchaConfig"`
	// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings. If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
	ChallengeConfig *WebAclChallengeConfig `pulumi:"challengeConfig"`
	// A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
	CustomResponseBodies map[string]WebAclCustomResponseBody `pulumi:"customResponseBodies"`
	// Collection of dataProtects.
	DataProtectionConfig *WebAclDataProtectionConfig `pulumi:"dataProtectionConfig"`
	// The action to perform if none of the `Rules` contained in the `WebACL` match.
	DefaultAction *WebAclDefaultAction `pulumi:"defaultAction"`
	// A description of the web ACL that helps with identification.
	Description *string `pulumi:"description"`
	// The ID of the web ACL.
	Id *string `pulumi:"id"`
	// The label namespace prefix for this web ACL. All labels added by rules in this web ACL have this prefix.
	//
	// The syntax for the label namespace prefix for a web ACL is the following: `awswaf:<account ID>:webacl:<web ACL name>:`
	//
	// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
	LabelNamespace *string `pulumi:"labelNamespace"`
	// Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
	OnSourceDDoSProtectionConfig *WebAclOnSourceDDoSProtectionConfig `pulumi:"onSourceDDoSProtectionConfig"`
	// Collection of Rules.
	Rules []WebAclRule `pulumi:"rules"`
	// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags []aws.Tag `pulumi:"tags"`
	// Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
	TokenDomains []string `pulumi:"tokenDomains"`
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig *WebAclVisibilityConfig `pulumi:"visibilityConfig"`
}

func LookupWebAclOutput(ctx *pulumi.Context, args LookupWebAclOutputArgs, opts ...pulumi.InvokeOption) LookupWebAclResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWebAclResultOutput, error) {
			args := v.(LookupWebAclArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:wafv2:getWebAcl", args, LookupWebAclResultOutput{}, options).(LookupWebAclResultOutput), nil
		}).(LookupWebAclResultOutput)
}

type LookupWebAclOutputArgs struct {
	// The ID of the web ACL.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the web ACL. You cannot change the name of a web ACL after you create it.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. For an AWS Amplify application, use `CLOUDFRONT` . A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	//
	// For information about how to define the association of the web ACL with your resource, see `WebACLAssociation` .
	Scope WebAclScopeInput `pulumi:"scope"`
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

// The Amazon Resource Name (ARN) of the web ACL.
func (o LookupWebAclResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Specifies custom configurations for the associations between the web ACL and protected resources.
//
// Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).
//
// > You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://docs.aws.amazon.com/waf/pricing/) .
//
// For Application Load Balancer and AWS AppSync , the limit is fixed at 8 KB (8,192 bytes).
func (o LookupWebAclResultOutput) AssociationConfig() WebAclAssociationConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclAssociationConfig { return v.AssociationConfig }).(WebAclAssociationConfigPtrOutput)
}

// The web ACL capacity units (WCUs) currently being used by this web ACL.
//
// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
func (o LookupWebAclResultOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig` .
func (o LookupWebAclResultOutput) CaptchaConfig() WebAclCaptchaConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclCaptchaConfig { return v.CaptchaConfig }).(WebAclCaptchaConfigPtrOutput)
}

// Specifies how AWS WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings. If you don't specify this, AWS WAF uses its default settings for `ChallengeConfig` .
func (o LookupWebAclResultOutput) ChallengeConfig() WebAclChallengeConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclChallengeConfig { return v.ChallengeConfig }).(WebAclChallengeConfigPtrOutput)
}

// A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.
//
// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the *AWS WAF Developer Guide* .
//
// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the *AWS WAF Developer Guide* .
func (o LookupWebAclResultOutput) CustomResponseBodies() WebAclCustomResponseBodyMapOutput {
	return o.ApplyT(func(v LookupWebAclResult) map[string]WebAclCustomResponseBody { return v.CustomResponseBodies }).(WebAclCustomResponseBodyMapOutput)
}

// Collection of dataProtects.
func (o LookupWebAclResultOutput) DataProtectionConfig() WebAclDataProtectionConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclDataProtectionConfig { return v.DataProtectionConfig }).(WebAclDataProtectionConfigPtrOutput)
}

// The action to perform if none of the `Rules` contained in the `WebACL` match.
func (o LookupWebAclResultOutput) DefaultAction() WebAclDefaultActionPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclDefaultAction { return v.DefaultAction }).(WebAclDefaultActionPtrOutput)
}

// A description of the web ACL that helps with identification.
func (o LookupWebAclResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the web ACL.
func (o LookupWebAclResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The label namespace prefix for this web ACL. All labels added by rules in this web ACL have this prefix.
//
// The syntax for the label namespace prefix for a web ACL is the following: `awswaf:<account ID>:webacl:<web ACL name>:`
//
// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
func (o LookupWebAclResultOutput) LabelNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *string { return v.LabelNamespace }).(pulumi.StringPtrOutput)
}

// Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.
func (o LookupWebAclResultOutput) OnSourceDDoSProtectionConfig() WebAclOnSourceDDoSProtectionConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclOnSourceDDoSProtectionConfig { return v.OnSourceDDoSProtectionConfig }).(WebAclOnSourceDDoSProtectionConfigPtrOutput)
}

// Collection of Rules.
func (o LookupWebAclResultOutput) Rules() WebAclRuleArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []WebAclRule { return v.Rules }).(WebAclRuleArrayOutput)
}

// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
//
// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
func (o LookupWebAclResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
func (o LookupWebAclResultOutput) TokenDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAclResult) []string { return v.TokenDomains }).(pulumi.StringArrayOutput)
}

// Defines and enables Amazon CloudWatch metrics and web request sample collection.
func (o LookupWebAclResultOutput) VisibilityConfig() WebAclVisibilityConfigPtrOutput {
	return o.ApplyT(func(v LookupWebAclResult) *WebAclVisibilityConfig { return v.VisibilityConfig }).(WebAclVisibilityConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAclResultOutput{})
}
