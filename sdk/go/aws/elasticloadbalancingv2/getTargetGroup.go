// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws-native:elasticloadbalancingv2:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetGroupArgs struct {
	// The ARN of the Target Group
	TargetGroupArn string `pulumi:"targetGroupArn"`
}

type LookupTargetGroupResult struct {
	// Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.
	HealthCheckEnabled *bool `pulumi:"healthCheckEnabled"`
	// The approximate amount of time, in seconds, between health checks of an individual target.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// [HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port the load balancer uses when performing health checks on targets.
	HealthCheckPort *string `pulumi:"healthCheckPort"`
	// The protocol the load balancer uses when performing health checks on targets.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The amount of time, in seconds, during which no response from a target means a failed health check.
	HealthCheckTimeoutSeconds *int `pulumi:"healthCheckTimeoutSeconds"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy.
	HealthyThresholdCount *int `pulumi:"healthyThresholdCount"`
	// The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.
	LoadBalancerArns []string `pulumi:"loadBalancerArns"`
	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
	Matcher *TargetGroupMatcher `pulumi:"matcher"`
	// The tags.
	Tags []aws.Tag `pulumi:"tags"`
	// The ARN of the Target Group
	TargetGroupArn *string `pulumi:"targetGroupArn"`
	// The attributes.
	TargetGroupAttributes []TargetGroupAttribute `pulumi:"targetGroupAttributes"`
	// The full name of the target group.
	TargetGroupFullName *string `pulumi:"targetGroupFullName"`
	// The name of the target group.
	TargetGroupName *string `pulumi:"targetGroupName"`
	// The targets.
	Targets []TargetGroupTargetDescription `pulumi:"targets"`
	// The number of consecutive health check failures required before considering a target unhealthy.
	UnhealthyThresholdCount *int `pulumi:"unhealthyThresholdCount"`
}

func LookupTargetGroupOutput(ctx *pulumi.Context, args LookupTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTargetGroupResultOutput, error) {
			args := v.(LookupTargetGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:elasticloadbalancingv2:getTargetGroup", args, LookupTargetGroupResultOutput{}, options).(LookupTargetGroupResultOutput), nil
		}).(LookupTargetGroupResultOutput)
}

type LookupTargetGroupOutputArgs struct {
	// The ARN of the Target Group
	TargetGroupArn pulumi.StringInput `pulumi:"targetGroupArn"`
}

func (LookupTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupArgs)(nil)).Elem()
}

type LookupTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupResult)(nil)).Elem()
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutput() LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutputWithContext(ctx context.Context) LookupTargetGroupResultOutput {
	return o
}

// Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.
func (o LookupTargetGroupResultOutput) HealthCheckEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *bool { return v.HealthCheckEnabled }).(pulumi.BoolPtrOutput)
}

// The approximate amount of time, in seconds, between health checks of an individual target.
func (o LookupTargetGroupResultOutput) HealthCheckIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthCheckIntervalSeconds }).(pulumi.IntPtrOutput)
}

// [HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.
func (o LookupTargetGroupResultOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

// The port the load balancer uses when performing health checks on targets.
func (o LookupTargetGroupResultOutput) HealthCheckPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckPort }).(pulumi.StringPtrOutput)
}

// The protocol the load balancer uses when performing health checks on targets.
func (o LookupTargetGroupResultOutput) HealthCheckProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.HealthCheckProtocol }).(pulumi.StringPtrOutput)
}

// The amount of time, in seconds, during which no response from a target means a failed health check.
func (o LookupTargetGroupResultOutput) HealthCheckTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthCheckTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// The number of consecutive health checks successes required before considering an unhealthy target healthy.
func (o LookupTargetGroupResultOutput) HealthyThresholdCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.HealthyThresholdCount }).(pulumi.IntPtrOutput)
}

// The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.
func (o LookupTargetGroupResultOutput) LoadBalancerArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []string { return v.LoadBalancerArns }).(pulumi.StringArrayOutput)
}

// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.
func (o LookupTargetGroupResultOutput) Matcher() TargetGroupMatcherPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *TargetGroupMatcher { return v.Matcher }).(TargetGroupMatcherPtrOutput)
}

// The tags.
func (o LookupTargetGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The ARN of the Target Group
func (o LookupTargetGroupResultOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

// The attributes.
func (o LookupTargetGroupResultOutput) TargetGroupAttributes() TargetGroupAttributeArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupAttribute { return v.TargetGroupAttributes }).(TargetGroupAttributeArrayOutput)
}

// The full name of the target group.
func (o LookupTargetGroupResultOutput) TargetGroupFullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupFullName }).(pulumi.StringPtrOutput)
}

// The name of the target group.
func (o LookupTargetGroupResultOutput) TargetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.TargetGroupName }).(pulumi.StringPtrOutput)
}

// The targets.
func (o LookupTargetGroupResultOutput) Targets() TargetGroupTargetDescriptionArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupTargetDescription { return v.Targets }).(TargetGroupTargetDescriptionArrayOutput)
}

// The number of consecutive health check failures required before considering a target unhealthy.
func (o LookupTargetGroupResultOutput) UnhealthyThresholdCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *int { return v.UnhealthyThresholdCount }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetGroupResultOutput{})
}
