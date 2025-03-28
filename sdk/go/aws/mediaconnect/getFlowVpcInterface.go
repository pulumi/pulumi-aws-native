// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaConnect::FlowVpcInterface
func LookupFlowVpcInterface(ctx *pulumi.Context, args *LookupFlowVpcInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupFlowVpcInterfaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFlowVpcInterfaceResult
	err := ctx.Invoke("aws-native:mediaconnect:getFlowVpcInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFlowVpcInterfaceArgs struct {
	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
	FlowArn string `pulumi:"flowArn"`
	// Immutable and has to be a unique against other VpcInterfaces in this Flow.
	Name string `pulumi:"name"`
}

type LookupFlowVpcInterfaceResult struct {
	// IDs of the network interfaces created in customer's account by MediaConnect.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// Role Arn MediaConnect can assume to create ENIs in customer's account.
	RoleArn *string `pulumi:"roleArn"`
	// Security Group IDs to be used on ENI.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Subnet must be in the AZ of the Flow
	SubnetId *string `pulumi:"subnetId"`
}

func LookupFlowVpcInterfaceOutput(ctx *pulumi.Context, args LookupFlowVpcInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupFlowVpcInterfaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFlowVpcInterfaceResultOutput, error) {
			args := v.(LookupFlowVpcInterfaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:mediaconnect:getFlowVpcInterface", args, LookupFlowVpcInterfaceResultOutput{}, options).(LookupFlowVpcInterfaceResultOutput), nil
		}).(LookupFlowVpcInterfaceResultOutput)
}

type LookupFlowVpcInterfaceOutputArgs struct {
	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
	FlowArn pulumi.StringInput `pulumi:"flowArn"`
	// Immutable and has to be a unique against other VpcInterfaces in this Flow.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupFlowVpcInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowVpcInterfaceArgs)(nil)).Elem()
}

type LookupFlowVpcInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupFlowVpcInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowVpcInterfaceResult)(nil)).Elem()
}

func (o LookupFlowVpcInterfaceResultOutput) ToLookupFlowVpcInterfaceResultOutput() LookupFlowVpcInterfaceResultOutput {
	return o
}

func (o LookupFlowVpcInterfaceResultOutput) ToLookupFlowVpcInterfaceResultOutputWithContext(ctx context.Context) LookupFlowVpcInterfaceResultOutput {
	return o
}

// IDs of the network interfaces created in customer's account by MediaConnect.
func (o LookupFlowVpcInterfaceResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFlowVpcInterfaceResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// Role Arn MediaConnect can assume to create ENIs in customer's account.
func (o LookupFlowVpcInterfaceResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowVpcInterfaceResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Security Group IDs to be used on ENI.
func (o LookupFlowVpcInterfaceResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFlowVpcInterfaceResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Subnet must be in the AZ of the Flow
func (o LookupFlowVpcInterfaceResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowVpcInterfaceResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlowVpcInterfaceResultOutput{})
}
