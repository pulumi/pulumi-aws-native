// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for a Redshift-managed VPC endpoint.
func LookupEndpointAccess(ctx *pulumi.Context, args *LookupEndpointAccessArgs, opts ...pulumi.InvokeOption) (*LookupEndpointAccessResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointAccessResult
	err := ctx.Invoke("aws-native:redshift:getEndpointAccess", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointAccessArgs struct {
	// The name of the endpoint.
	EndpointName string `pulumi:"endpointName"`
}

type LookupEndpointAccessResult struct {
	// The DNS address of the endpoint.
	Address *string `pulumi:"address"`
	// The time (UTC) that the endpoint was created.
	EndpointCreateTime *string `pulumi:"endpointCreateTime"`
	// The status of the endpoint.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// The port number on which the cluster accepts incoming connections.
	Port *int `pulumi:"port"`
	// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
	VpcEndpoint *VpcEndpointProperties `pulumi:"vpcEndpoint"`
	// A list of vpc security group ids to apply to the created endpoint access.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.
	VpcSecurityGroups []EndpointAccessVpcSecurityGroup `pulumi:"vpcSecurityGroups"`
}

func LookupEndpointAccessOutput(ctx *pulumi.Context, args LookupEndpointAccessOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointAccessResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEndpointAccessResultOutput, error) {
			args := v.(LookupEndpointAccessArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:redshift:getEndpointAccess", args, LookupEndpointAccessResultOutput{}, options).(LookupEndpointAccessResultOutput), nil
		}).(LookupEndpointAccessResultOutput)
}

type LookupEndpointAccessOutputArgs struct {
	// The name of the endpoint.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
}

func (LookupEndpointAccessOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAccessArgs)(nil)).Elem()
}

type LookupEndpointAccessResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointAccessResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAccessResult)(nil)).Elem()
}

func (o LookupEndpointAccessResultOutput) ToLookupEndpointAccessResultOutput() LookupEndpointAccessResultOutput {
	return o
}

func (o LookupEndpointAccessResultOutput) ToLookupEndpointAccessResultOutputWithContext(ctx context.Context) LookupEndpointAccessResultOutput {
	return o
}

// The DNS address of the endpoint.
func (o LookupEndpointAccessResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The time (UTC) that the endpoint was created.
func (o LookupEndpointAccessResultOutput) EndpointCreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) *string { return v.EndpointCreateTime }).(pulumi.StringPtrOutput)
}

// The status of the endpoint.
func (o LookupEndpointAccessResultOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

// The port number on which the cluster accepts incoming connections.
func (o LookupEndpointAccessResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
func (o LookupEndpointAccessResultOutput) VpcEndpoint() VpcEndpointPropertiesPtrOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) *VpcEndpointProperties { return v.VpcEndpoint }).(VpcEndpointPropertiesPtrOutput)
}

// A list of vpc security group ids to apply to the created endpoint access.
func (o LookupEndpointAccessResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint.
func (o LookupEndpointAccessResultOutput) VpcSecurityGroups() EndpointAccessVpcSecurityGroupArrayOutput {
	return o.ApplyT(func(v LookupEndpointAccessResult) []EndpointAccessVpcSecurityGroup { return v.VpcSecurityGroups }).(EndpointAccessVpcSecurityGroupArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointAccessResultOutput{})
}
