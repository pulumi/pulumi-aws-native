// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package greengrassv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for Greengrass component version.
func LookupComponentVersion(ctx *pulumi.Context, args *LookupComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupComponentVersionResult, error) {
	var rv LookupComponentVersionResult
	err := ctx.Invoke("aws-native:greengrassv2:getComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentVersionArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupComponentVersionResult struct {
	Arn              *string     `pulumi:"arn"`
	ComponentName    *string     `pulumi:"componentName"`
	ComponentVersion *string     `pulumi:"componentVersion"`
	Tags             interface{} `pulumi:"tags"`
}

func LookupComponentVersionOutput(ctx *pulumi.Context, args LookupComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentVersionResult, error) {
			args := v.(LookupComponentVersionArgs)
			r, err := LookupComponentVersion(ctx, &args, opts...)
			var s LookupComponentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentVersionResultOutput)
}

type LookupComponentVersionOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionArgs)(nil)).Elem()
}

type LookupComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionResult)(nil)).Elem()
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutput() LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutputWithContext(ctx context.Context) LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupComponentVersionResultOutput) ComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.ComponentName }).(pulumi.StringPtrOutput)
}

func (o LookupComponentVersionResultOutput) ComponentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) *string { return v.ComponentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupComponentVersionResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentVersionResultOutput{})
}