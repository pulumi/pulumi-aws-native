// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCCidrBlock
func LookupVPCCidrBlock(ctx *pulumi.Context, args *LookupVPCCidrBlockArgs, opts ...pulumi.InvokeOption) (*LookupVPCCidrBlockResult, error) {
	var rv LookupVPCCidrBlockResult
	err := ctx.Invoke("aws-native:ec2:getVPCCidrBlock", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCCidrBlockArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPCCidrBlockResult struct {
	Id *string `pulumi:"id"`
}

func LookupVPCCidrBlockOutput(ctx *pulumi.Context, args LookupVPCCidrBlockOutputArgs, opts ...pulumi.InvokeOption) LookupVPCCidrBlockResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCCidrBlockResult, error) {
			args := v.(LookupVPCCidrBlockArgs)
			r, err := LookupVPCCidrBlock(ctx, &args, opts...)
			var s LookupVPCCidrBlockResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPCCidrBlockResultOutput)
}

type LookupVPCCidrBlockOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPCCidrBlockOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCCidrBlockArgs)(nil)).Elem()
}

type LookupVPCCidrBlockResultOutput struct{ *pulumi.OutputState }

func (LookupVPCCidrBlockResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCCidrBlockResult)(nil)).Elem()
}

func (o LookupVPCCidrBlockResultOutput) ToLookupVPCCidrBlockResultOutput() LookupVPCCidrBlockResultOutput {
	return o
}

func (o LookupVPCCidrBlockResultOutput) ToLookupVPCCidrBlockResultOutputWithContext(ctx context.Context) LookupVPCCidrBlockResultOutput {
	return o
}

func (o LookupVPCCidrBlockResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCCidrBlockResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCCidrBlockResultOutput{})
}