// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A response headers policy.
//
//	A response headers policy contains information about a set of HTTP response headers.
//	After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
//	For more information, see [Adding or removing HTTP headers in CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html) in the *Amazon CloudFront Developer Guide*.
func LookupResponseHeadersPolicy(ctx *pulumi.Context, args *LookupResponseHeadersPolicyArgs, opts ...pulumi.InvokeOption) (*LookupResponseHeadersPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResponseHeadersPolicyResult
	err := ctx.Invoke("aws-native:cloudfront:getResponseHeadersPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResponseHeadersPolicyArgs struct {
	// The unique identifier for the response headers policy. For example: `57f99797-3b20-4e1b-a728-27972a74082a` .
	Id string `pulumi:"id"`
}

type LookupResponseHeadersPolicyResult struct {
	// The unique identifier for the response headers policy. For example: `57f99797-3b20-4e1b-a728-27972a74082a` .
	Id *string `pulumi:"id"`
	// The date and time when the response headers policy was last modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// A response headers policy configuration.
	ResponseHeadersPolicyConfig *ResponseHeadersPolicyConfig `pulumi:"responseHeadersPolicyConfig"`
}

func LookupResponseHeadersPolicyOutput(ctx *pulumi.Context, args LookupResponseHeadersPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupResponseHeadersPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResponseHeadersPolicyResultOutput, error) {
			args := v.(LookupResponseHeadersPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudfront:getResponseHeadersPolicy", args, LookupResponseHeadersPolicyResultOutput{}, options).(LookupResponseHeadersPolicyResultOutput), nil
		}).(LookupResponseHeadersPolicyResultOutput)
}

type LookupResponseHeadersPolicyOutputArgs struct {
	// The unique identifier for the response headers policy. For example: `57f99797-3b20-4e1b-a728-27972a74082a` .
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResponseHeadersPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponseHeadersPolicyArgs)(nil)).Elem()
}

type LookupResponseHeadersPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupResponseHeadersPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponseHeadersPolicyResult)(nil)).Elem()
}

func (o LookupResponseHeadersPolicyResultOutput) ToLookupResponseHeadersPolicyResultOutput() LookupResponseHeadersPolicyResultOutput {
	return o
}

func (o LookupResponseHeadersPolicyResultOutput) ToLookupResponseHeadersPolicyResultOutputWithContext(ctx context.Context) LookupResponseHeadersPolicyResultOutput {
	return o
}

// The unique identifier for the response headers policy. For example: `57f99797-3b20-4e1b-a728-27972a74082a` .
func (o LookupResponseHeadersPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The date and time when the response headers policy was last modified.
func (o LookupResponseHeadersPolicyResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// A response headers policy configuration.
func (o LookupResponseHeadersPolicyResultOutput) ResponseHeadersPolicyConfig() ResponseHeadersPolicyConfigPtrOutput {
	return o.ApplyT(func(v LookupResponseHeadersPolicyResult) *ResponseHeadersPolicyConfig {
		return v.ResponseHeadersPolicyConfig
	}).(ResponseHeadersPolicyConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResponseHeadersPolicyResultOutput{})
}
