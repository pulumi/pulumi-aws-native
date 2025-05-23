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

// Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.
//
//	You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought to AWS for use with your AWS resources using bring your own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the *Amazon EC2 User Guide*.
//	For more information, see [Elastic IP Addresses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) in the *Amazon EC2 User Guide*.
func LookupEip(ctx *pulumi.Context, args *LookupEipArgs, opts ...pulumi.InvokeOption) (*LookupEipResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEipResult
	err := ctx.Invoke("aws-native:ec2:getEip", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEipArgs struct {
	// The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
	AllocationId string `pulumi:"allocationId"`
	// The Elastic IP address.
	PublicIp string `pulumi:"publicIp"`
}

type LookupEipResult struct {
	// The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
	AllocationId *string `pulumi:"allocationId"`
	// The network (``vpc``).
	//  If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
	Domain *string `pulumi:"domain"`
	// The ID of the instance.
	//   Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	InstanceId *string `pulumi:"instanceId"`
	// The Elastic IP address.
	PublicIp *string `pulumi:"publicIp"`
	// The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
	//   Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	PublicIpv4Pool *string `pulumi:"publicIpv4Pool"`
	// Any tags assigned to the Elastic IP address.
	//   Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupEipOutput(ctx *pulumi.Context, args LookupEipOutputArgs, opts ...pulumi.InvokeOption) LookupEipResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEipResultOutput, error) {
			args := v.(LookupEipArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getEip", args, LookupEipResultOutput{}, options).(LookupEipResultOutput), nil
		}).(LookupEipResultOutput)
}

type LookupEipOutputArgs struct {
	// The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
	AllocationId pulumi.StringInput `pulumi:"allocationId"`
	// The Elastic IP address.
	PublicIp pulumi.StringInput `pulumi:"publicIp"`
}

func (LookupEipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEipArgs)(nil)).Elem()
}

type LookupEipResultOutput struct{ *pulumi.OutputState }

func (LookupEipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEipResult)(nil)).Elem()
}

func (o LookupEipResultOutput) ToLookupEipResultOutput() LookupEipResultOutput {
	return o
}

func (o LookupEipResultOutput) ToLookupEipResultOutputWithContext(ctx context.Context) LookupEipResultOutput {
	return o
}

// The ID that AWS assigns to represent the allocation of the address for use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example, `eipalloc-5723d13e` .
func (o LookupEipResultOutput) AllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.AllocationId }).(pulumi.StringPtrOutput)
}

// The network (“vpc“).
//
//	If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
func (o LookupEipResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// The ID of the instance.
//
//	Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
func (o LookupEipResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// The Elastic IP address.
func (o LookupEipResultOutput) PublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.PublicIp }).(pulumi.StringPtrOutput)
}

// The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
//
//	Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
func (o LookupEipResultOutput) PublicIpv4Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEipResult) *string { return v.PublicIpv4Pool }).(pulumi.StringPtrOutput)
}

// Any tags assigned to the Elastic IP address.
//
//	Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
func (o LookupEipResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupEipResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEipResultOutput{})
}
