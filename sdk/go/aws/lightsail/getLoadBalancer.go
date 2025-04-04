// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::LoadBalancer
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("aws-native:lightsail:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	// The name of your load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
}

type LookupLoadBalancerResult struct {
	// The names of the instances attached to the load balancer.
	AttachedInstances []string `pulumi:"attachedInstances"`
	// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string `pulumi:"loadBalancerArn"`
	// Configuration option to enable session stickiness.
	SessionStickinessEnabled *bool `pulumi:"sessionStickinessEnabled"`
	// Configuration option to adjust session stickiness cookie duration parameter.
	SessionStickinessLbCookieDurationSeconds *string `pulumi:"sessionStickinessLbCookieDurationSeconds"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The name of the TLS policy to apply to the load balancer.
	TlsPolicyName *string `pulumi:"tlsPolicyName"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResultOutput, error) {
			args := v.(LookupLoadBalancerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:lightsail:getLoadBalancer", args, LookupLoadBalancerResultOutput{}, options).(LookupLoadBalancerResultOutput), nil
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	// The name of your load balancer.
	LoadBalancerName pulumi.StringInput `pulumi:"loadBalancerName"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

// The names of the instances attached to the load balancer.
func (o LookupLoadBalancerResultOutput) AttachedInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []string { return v.AttachedInstances }).(pulumi.StringArrayOutput)
}

// The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., "/").
func (o LookupLoadBalancerResultOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the load balancer.
func (o LookupLoadBalancerResultOutput) LoadBalancerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.LoadBalancerArn }).(pulumi.StringPtrOutput)
}

// Configuration option to enable session stickiness.
func (o LookupLoadBalancerResultOutput) SessionStickinessEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *bool { return v.SessionStickinessEnabled }).(pulumi.BoolPtrOutput)
}

// Configuration option to adjust session stickiness cookie duration parameter.
func (o LookupLoadBalancerResultOutput) SessionStickinessLbCookieDurationSeconds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.SessionStickinessLbCookieDurationSeconds }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLoadBalancerResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The name of the TLS policy to apply to the load balancer.
func (o LookupLoadBalancerResultOutput) TlsPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.TlsPolicyName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
