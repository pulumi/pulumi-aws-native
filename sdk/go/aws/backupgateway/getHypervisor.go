// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package backupgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::BackupGateway::Hypervisor Resource Type
func LookupHypervisor(ctx *pulumi.Context, args *LookupHypervisorArgs, opts ...pulumi.InvokeOption) (*LookupHypervisorResult, error) {
	var rv LookupHypervisorResult
	err := ctx.Invoke("aws-native:backupgateway:getHypervisor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHypervisorArgs struct {
	HypervisorArn string `pulumi:"hypervisorArn"`
}

type LookupHypervisorResult struct {
	Host          *string `pulumi:"host"`
	HypervisorArn *string `pulumi:"hypervisorArn"`
}

func LookupHypervisorOutput(ctx *pulumi.Context, args LookupHypervisorOutputArgs, opts ...pulumi.InvokeOption) LookupHypervisorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHypervisorResult, error) {
			args := v.(LookupHypervisorArgs)
			r, err := LookupHypervisor(ctx, &args, opts...)
			var s LookupHypervisorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHypervisorResultOutput)
}

type LookupHypervisorOutputArgs struct {
	HypervisorArn pulumi.StringInput `pulumi:"hypervisorArn"`
}

func (LookupHypervisorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervisorArgs)(nil)).Elem()
}

type LookupHypervisorResultOutput struct{ *pulumi.OutputState }

func (LookupHypervisorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervisorResult)(nil)).Elem()
}

func (o LookupHypervisorResultOutput) ToLookupHypervisorResultOutput() LookupHypervisorResultOutput {
	return o
}

func (o LookupHypervisorResultOutput) ToLookupHypervisorResultOutputWithContext(ctx context.Context) LookupHypervisorResultOutput {
	return o
}

func (o LookupHypervisorResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervisorResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o LookupHypervisorResultOutput) HypervisorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervisorResult) *string { return v.HypervisorArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHypervisorResultOutput{})
}