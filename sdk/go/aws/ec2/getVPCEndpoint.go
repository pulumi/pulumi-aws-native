// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPCEndpoint
func LookupVPCEndpoint(ctx *pulumi.Context, args *LookupVPCEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVPCEndpointResult, error) {
	var rv LookupVPCEndpointResult
	err := ctx.Invoke("aws-native:ec2:getVPCEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCEndpointArgs struct {
	Id string `pulumi:"id"`
}

type LookupVPCEndpointResult struct {
	CreationTimestamp   *string  `pulumi:"creationTimestamp"`
	DnsEntries          []string `pulumi:"dnsEntries"`
	Id                  *string  `pulumi:"id"`
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A policy to attach to the endpoint that controls access to the service.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Indicate whether to associate a private hosted zone with the specified VPC.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// One or more route table IDs.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the endpoint network interface.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The ID of one or more subnets in which to create an endpoint network interface.
	SubnetIds []string `pulumi:"subnetIds"`
}

func LookupVPCEndpointOutput(ctx *pulumi.Context, args LookupVPCEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVPCEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCEndpointResult, error) {
			args := v.(LookupVPCEndpointArgs)
			r, err := LookupVPCEndpoint(ctx, &args, opts...)
			var s LookupVPCEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPCEndpointResultOutput)
}

type LookupVPCEndpointOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupVPCEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEndpointArgs)(nil)).Elem()
}

type LookupVPCEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVPCEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEndpointResult)(nil)).Elem()
}

func (o LookupVPCEndpointResultOutput) ToLookupVPCEndpointResultOutput() LookupVPCEndpointResultOutput {
	return o
}

func (o LookupVPCEndpointResultOutput) ToLookupVPCEndpointResultOutputWithContext(ctx context.Context) LookupVPCEndpointResultOutput {
	return o
}

func (o LookupVPCEndpointResultOutput) CreationTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) *string { return v.CreationTimestamp }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEndpointResultOutput) DnsEntries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) []string { return v.DnsEntries }).(pulumi.StringArrayOutput)
}

func (o LookupVPCEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEndpointResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// A policy to attach to the endpoint that controls access to the service.
func (o LookupVPCEndpointResultOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) *string { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// Indicate whether to associate a private hosted zone with the specified VPC.
func (o LookupVPCEndpointResultOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) *bool { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// One or more route table IDs.
func (o LookupVPCEndpointResultOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) []string { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The ID of one or more security groups to associate with the endpoint network interface.
func (o LookupVPCEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The ID of one or more subnets in which to create an endpoint network interface.
func (o LookupVPCEndpointResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVPCEndpointResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCEndpointResultOutput{})
}