// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation
func LookupVpcEndpointAssociation(ctx *pulumi.Context, args *LookupVpcEndpointAssociationArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointAssociationResult
	err := ctx.Invoke("aws-native:networkfirewall:getVpcEndpointAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpcEndpointAssociationArgs struct {
	// The Amazon Resource Name (ARN) of a VPC endpoint association.
	VpcEndpointAssociationArn string `pulumi:"vpcEndpointAssociationArn"`
}

type LookupVpcEndpointAssociationResult struct {
	EndpointId *string `pulumi:"endpointId"`
	// The key:value pairs to associate with the resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The Amazon Resource Name (ARN) of a VPC endpoint association.
	VpcEndpointAssociationArn *string `pulumi:"vpcEndpointAssociationArn"`
	// The unique identifier of the VPC endpoint association.
	VpcEndpointAssociationId *string `pulumi:"vpcEndpointAssociationId"`
}

func LookupVpcEndpointAssociationOutput(ctx *pulumi.Context, args LookupVpcEndpointAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointAssociationResultOutput, error) {
			args := v.(LookupVpcEndpointAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:networkfirewall:getVpcEndpointAssociation", args, LookupVpcEndpointAssociationResultOutput{}, options).(LookupVpcEndpointAssociationResultOutput), nil
		}).(LookupVpcEndpointAssociationResultOutput)
}

type LookupVpcEndpointAssociationOutputArgs struct {
	// The Amazon Resource Name (ARN) of a VPC endpoint association.
	VpcEndpointAssociationArn pulumi.StringInput `pulumi:"vpcEndpointAssociationArn"`
}

func (LookupVpcEndpointAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointAssociationArgs)(nil)).Elem()
}

type LookupVpcEndpointAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointAssociationResult)(nil)).Elem()
}

func (o LookupVpcEndpointAssociationResultOutput) ToLookupVpcEndpointAssociationResultOutput() LookupVpcEndpointAssociationResultOutput {
	return o
}

func (o LookupVpcEndpointAssociationResultOutput) ToLookupVpcEndpointAssociationResultOutputWithContext(ctx context.Context) LookupVpcEndpointAssociationResultOutput {
	return o
}

func (o LookupVpcEndpointAssociationResultOutput) EndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointAssociationResult) *string { return v.EndpointId }).(pulumi.StringPtrOutput)
}

// The key:value pairs to associate with the resource.
func (o LookupVpcEndpointAssociationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointAssociationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The Amazon Resource Name (ARN) of a VPC endpoint association.
func (o LookupVpcEndpointAssociationResultOutput) VpcEndpointAssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointAssociationResult) *string { return v.VpcEndpointAssociationArn }).(pulumi.StringPtrOutput)
}

// The unique identifier of the VPC endpoint association.
func (o LookupVpcEndpointAssociationResultOutput) VpcEndpointAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointAssociationResult) *string { return v.VpcEndpointAssociationId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointAssociationResultOutput{})
}
