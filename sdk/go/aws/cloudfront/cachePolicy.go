// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A cache policy.
//
//	When it's attached to a cache behavior, the cache policy determines the following:
//	 +  The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
//	 +  The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
//
//	The headers, cookies, and query strings that are included in the cache key are also included in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a valid object in its cache that matches the request's cache key. If you want to send values to the origin but *not* include them in the cache key, use ``OriginRequestPolicy``.
type CachePolicy struct {
	pulumi.CustomResourceState

	// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The cache policy configuration.
	CachePolicyConfig CachePolicyConfigOutput `pulumi:"cachePolicyConfig"`
	// The date and time when the cache policy was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
}

// NewCachePolicy registers a new resource with the given unique name, arguments, and options.
func NewCachePolicy(ctx *pulumi.Context,
	name string, args *CachePolicyArgs, opts ...pulumi.ResourceOption) (*CachePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CachePolicyConfig == nil {
		return nil, errors.New("invalid value for required argument 'CachePolicyConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CachePolicy
	err := ctx.RegisterResource("aws-native:cloudfront:CachePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCachePolicy gets an existing CachePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCachePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CachePolicyState, opts ...pulumi.ResourceOption) (*CachePolicy, error) {
	var resource CachePolicy
	err := ctx.ReadResource("aws-native:cloudfront:CachePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CachePolicy resources.
type cachePolicyState struct {
}

type CachePolicyState struct {
}

func (CachePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cachePolicyState)(nil)).Elem()
}

type cachePolicyArgs struct {
	// The cache policy configuration.
	CachePolicyConfig CachePolicyConfig `pulumi:"cachePolicyConfig"`
}

// The set of arguments for constructing a CachePolicy resource.
type CachePolicyArgs struct {
	// The cache policy configuration.
	CachePolicyConfig CachePolicyConfigInput
}

func (CachePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cachePolicyArgs)(nil)).Elem()
}

type CachePolicyInput interface {
	pulumi.Input

	ToCachePolicyOutput() CachePolicyOutput
	ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput
}

func (*CachePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**CachePolicy)(nil)).Elem()
}

func (i *CachePolicy) ToCachePolicyOutput() CachePolicyOutput {
	return i.ToCachePolicyOutputWithContext(context.Background())
}

func (i *CachePolicy) ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CachePolicyOutput)
}

type CachePolicyOutput struct{ *pulumi.OutputState }

func (CachePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CachePolicy)(nil)).Elem()
}

func (o CachePolicyOutput) ToCachePolicyOutput() CachePolicyOutput {
	return o
}

func (o CachePolicyOutput) ToCachePolicyOutputWithContext(ctx context.Context) CachePolicyOutput {
	return o
}

// The unique identifier for the cache policy. For example: `2766f7b2-75c5-41c6-8f06-bf4303a2f2f5` .
func (o CachePolicyOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The cache policy configuration.
func (o CachePolicyOutput) CachePolicyConfig() CachePolicyConfigOutput {
	return o.ApplyT(func(v *CachePolicy) CachePolicyConfigOutput { return v.CachePolicyConfig }).(CachePolicyConfigOutput)
}

// The date and time when the cache policy was last modified.
func (o CachePolicyOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CachePolicy) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CachePolicyInput)(nil)).Elem(), &CachePolicy{})
	pulumi.RegisterOutputType(CachePolicyOutput{})
}
