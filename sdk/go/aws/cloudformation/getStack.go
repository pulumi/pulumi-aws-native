// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFormation::Stack
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	var rv LookupStackResult
	err := ctx.Invoke("aws-native:cloudformation:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackArgs struct {
	Id string `pulumi:"id"`
}

type LookupStackResult struct {
	Id               *string     `pulumi:"id"`
	NotificationARNs []string    `pulumi:"notificationARNs"`
	Parameters       interface{} `pulumi:"parameters"`
	Tags             []StackTag  `pulumi:"tags"`
	TemplateURL      *string     `pulumi:"templateURL"`
	TimeoutInMinutes *int        `pulumi:"timeoutInMinutes"`
}

func LookupStackOutput(ctx *pulumi.Context, args LookupStackOutputArgs, opts ...pulumi.InvokeOption) LookupStackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackResult, error) {
			args := v.(LookupStackArgs)
			r, err := LookupStack(ctx, &args, opts...)
			var s LookupStackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackResultOutput)
}

type LookupStackOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackArgs)(nil)).Elem()
}

type LookupStackResultOutput struct{ *pulumi.OutputState }

func (LookupStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackResult)(nil)).Elem()
}

func (o LookupStackResultOutput) ToLookupStackResultOutput() LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToLookupStackResultOutputWithContext(ctx context.Context) LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) NotificationARNs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []string { return v.NotificationARNs }).(pulumi.StringArrayOutput)
}

func (o LookupStackResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupStackResultOutput) Tags() StackTagArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackTag { return v.Tags }).(StackTagArrayOutput)
}

func (o LookupStackResultOutput) TemplateURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.TemplateURL }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) TimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *int { return v.TimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackResultOutput{})
}