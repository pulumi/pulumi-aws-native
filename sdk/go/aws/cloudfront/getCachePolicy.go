// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::CachePolicy
func LookupCachePolicy(ctx *pulumi.Context, args *LookupCachePolicyArgs, opts ...pulumi.InvokeOption) (*LookupCachePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCachePolicyResult
	err := ctx.Invoke("aws-native:cloudfront:getCachePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCachePolicyArgs struct {
	// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
	Id string `pulumi:"id"`
}

type LookupCachePolicyResult struct {
	// The cache policy configuration.
	CachePolicyConfig *CachePolicyConfig `pulumi:"cachePolicyConfig"`
	// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
	Id *string `pulumi:"id"`
	// The date and time when the cache policy was last modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
}

func LookupCachePolicyOutput(ctx *pulumi.Context, args LookupCachePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupCachePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCachePolicyResultOutput, error) {
			args := v.(LookupCachePolicyArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupCachePolicyResult
			secret, err := ctx.InvokePackageRaw("aws-native:cloudfront:getCachePolicy", args, &rv, "", opts...)
			if err != nil {
				return LookupCachePolicyResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupCachePolicyResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupCachePolicyResultOutput), nil
			}
			return output, nil
		}).(LookupCachePolicyResultOutput)
}

type LookupCachePolicyOutputArgs struct {
	// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCachePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCachePolicyArgs)(nil)).Elem()
}

type LookupCachePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCachePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCachePolicyResult)(nil)).Elem()
}

func (o LookupCachePolicyResultOutput) ToLookupCachePolicyResultOutput() LookupCachePolicyResultOutput {
	return o
}

func (o LookupCachePolicyResultOutput) ToLookupCachePolicyResultOutputWithContext(ctx context.Context) LookupCachePolicyResultOutput {
	return o
}

// The cache policy configuration.
func (o LookupCachePolicyResultOutput) CachePolicyConfig() CachePolicyConfigPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) *CachePolicyConfig { return v.CachePolicyConfig }).(CachePolicyConfigPtrOutput)
}

// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
func (o LookupCachePolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The date and time when the cache policy was last modified.
func (o LookupCachePolicyResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCachePolicyResultOutput{})
}
