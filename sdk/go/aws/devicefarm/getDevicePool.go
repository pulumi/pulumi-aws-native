// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::DevicePool creates a new Device Pool for a given DF Project
func LookupDevicePool(ctx *pulumi.Context, args *LookupDevicePoolArgs, opts ...pulumi.InvokeOption) (*LookupDevicePoolResult, error) {
	var rv LookupDevicePoolResult
	err := ctx.Invoke("aws-native:devicefarm:getDevicePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevicePoolArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupDevicePoolResult struct {
	Arn         *string          `pulumi:"arn"`
	Description *string          `pulumi:"description"`
	MaxDevices  *int             `pulumi:"maxDevices"`
	Name        *string          `pulumi:"name"`
	Rules       []DevicePoolRule `pulumi:"rules"`
	Tags        []DevicePoolTag  `pulumi:"tags"`
}

func LookupDevicePoolOutput(ctx *pulumi.Context, args LookupDevicePoolOutputArgs, opts ...pulumi.InvokeOption) LookupDevicePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevicePoolResult, error) {
			args := v.(LookupDevicePoolArgs)
			r, err := LookupDevicePool(ctx, &args, opts...)
			var s LookupDevicePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevicePoolResultOutput)
}

type LookupDevicePoolOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupDevicePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevicePoolArgs)(nil)).Elem()
}

type LookupDevicePoolResultOutput struct{ *pulumi.OutputState }

func (LookupDevicePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevicePoolResult)(nil)).Elem()
}

func (o LookupDevicePoolResultOutput) ToLookupDevicePoolResultOutput() LookupDevicePoolResultOutput {
	return o
}

func (o LookupDevicePoolResultOutput) ToLookupDevicePoolResultOutputWithContext(ctx context.Context) LookupDevicePoolResultOutput {
	return o
}

func (o LookupDevicePoolResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupDevicePoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDevicePoolResultOutput) MaxDevices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) *int { return v.MaxDevices }).(pulumi.IntPtrOutput)
}

func (o LookupDevicePoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupDevicePoolResultOutput) Rules() DevicePoolRuleArrayOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) []DevicePoolRule { return v.Rules }).(DevicePoolRuleArrayOutput)
}

func (o LookupDevicePoolResultOutput) Tags() DevicePoolTagArrayOutput {
	return o.ApplyT(func(v LookupDevicePoolResult) []DevicePoolTag { return v.Tags }).(DevicePoolTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevicePoolResultOutput{})
}