// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VPC Route Server
func LookupRouteServer(ctx *pulumi.Context, args *LookupRouteServerArgs, opts ...pulumi.InvokeOption) (*LookupRouteServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteServerResult
	err := ctx.Invoke("aws-native:ec2:getRouteServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteServerArgs struct {
	// The ID of the Route Server.
	Id string `pulumi:"id"`
}

type LookupRouteServerResult struct {
	// The Amazon Resource Name (ARN) of the Route Server.
	Arn *string `pulumi:"arn"`
	// The ID of the Route Server.
	Id *string `pulumi:"id"`
	// Whether to enable persistent routes
	PersistRoutes *RouteServerPersistRoutes `pulumi:"persistRoutes"`
	// Whether to enable SNS notifications
	SnsNotificationsEnabled *bool `pulumi:"snsNotificationsEnabled"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupRouteServerOutput(ctx *pulumi.Context, args LookupRouteServerOutputArgs, opts ...pulumi.InvokeOption) LookupRouteServerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRouteServerResultOutput, error) {
			args := v.(LookupRouteServerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getRouteServer", args, LookupRouteServerResultOutput{}, options).(LookupRouteServerResultOutput), nil
		}).(LookupRouteServerResultOutput)
}

type LookupRouteServerOutputArgs struct {
	// The ID of the Route Server.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupRouteServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteServerArgs)(nil)).Elem()
}

type LookupRouteServerResultOutput struct{ *pulumi.OutputState }

func (LookupRouteServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteServerResult)(nil)).Elem()
}

func (o LookupRouteServerResultOutput) ToLookupRouteServerResultOutput() LookupRouteServerResultOutput {
	return o
}

func (o LookupRouteServerResultOutput) ToLookupRouteServerResultOutputWithContext(ctx context.Context) LookupRouteServerResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Route Server.
func (o LookupRouteServerResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteServerResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The ID of the Route Server.
func (o LookupRouteServerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteServerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Whether to enable persistent routes
func (o LookupRouteServerResultOutput) PersistRoutes() RouteServerPersistRoutesPtrOutput {
	return o.ApplyT(func(v LookupRouteServerResult) *RouteServerPersistRoutes { return v.PersistRoutes }).(RouteServerPersistRoutesPtrOutput)
}

// Whether to enable SNS notifications
func (o LookupRouteServerResultOutput) SnsNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteServerResult) *bool { return v.SnsNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupRouteServerResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupRouteServerResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteServerResultOutput{})
}
