// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.
func LookupFlowLog(ctx *pulumi.Context, args *LookupFlowLogArgs, opts ...pulumi.InvokeOption) (*LookupFlowLogResult, error) {
	var rv LookupFlowLogResult
	err := ctx.Invoke("aws-native:ec2:getFlowLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFlowLogArgs struct {
	// The Flow Log ID
	Id string `pulumi:"id"`
}

type LookupFlowLogResult struct {
	// The Flow Log ID
	Id *string `pulumi:"id"`
	// The tags to apply to the flow logs.
	Tags []FlowLogTag `pulumi:"tags"`
}

func LookupFlowLogOutput(ctx *pulumi.Context, args LookupFlowLogOutputArgs, opts ...pulumi.InvokeOption) LookupFlowLogResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlowLogResult, error) {
			args := v.(LookupFlowLogArgs)
			r, err := LookupFlowLog(ctx, &args, opts...)
			var s LookupFlowLogResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFlowLogResultOutput)
}

type LookupFlowLogOutputArgs struct {
	// The Flow Log ID
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupFlowLogOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowLogArgs)(nil)).Elem()
}

type LookupFlowLogResultOutput struct{ *pulumi.OutputState }

func (LookupFlowLogResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowLogResult)(nil)).Elem()
}

func (o LookupFlowLogResultOutput) ToLookupFlowLogResultOutput() LookupFlowLogResultOutput {
	return o
}

func (o LookupFlowLogResultOutput) ToLookupFlowLogResultOutputWithContext(ctx context.Context) LookupFlowLogResultOutput {
	return o
}

// The Flow Log ID
func (o LookupFlowLogResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowLogResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The tags to apply to the flow logs.
func (o LookupFlowLogResultOutput) Tags() FlowLogTagArrayOutput {
	return o.ApplyT(func(v LookupFlowLogResult) []FlowLogTag { return v.Tags }).(FlowLogTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlowLogResultOutput{})
}