// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an association between a local gateway route table and a VPC.
func LookupLocalGatewayRouteTableVpcAssociation(ctx *pulumi.Context, args *LookupLocalGatewayRouteTableVpcAssociationArgs, opts ...pulumi.InvokeOption) (*LookupLocalGatewayRouteTableVpcAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLocalGatewayRouteTableVpcAssociationResult
	err := ctx.Invoke("aws-native:ec2:getLocalGatewayRouteTableVpcAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocalGatewayRouteTableVpcAssociationArgs struct {
	// The ID of the association.
	LocalGatewayRouteTableVpcAssociationId string `pulumi:"localGatewayRouteTableVpcAssociationId"`
}

type LookupLocalGatewayRouteTableVpcAssociationResult struct {
	// The ID of the local gateway.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// The ID of the association.
	LocalGatewayRouteTableVpcAssociationId *string `pulumi:"localGatewayRouteTableVpcAssociationId"`
	// The state of the association.
	State *string `pulumi:"state"`
	// The tags for the association.
	Tags []LocalGatewayRouteTableVpcAssociationTag `pulumi:"tags"`
}

func LookupLocalGatewayRouteTableVpcAssociationOutput(ctx *pulumi.Context, args LookupLocalGatewayRouteTableVpcAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupLocalGatewayRouteTableVpcAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalGatewayRouteTableVpcAssociationResult, error) {
			args := v.(LookupLocalGatewayRouteTableVpcAssociationArgs)
			r, err := LookupLocalGatewayRouteTableVpcAssociation(ctx, &args, opts...)
			var s LookupLocalGatewayRouteTableVpcAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalGatewayRouteTableVpcAssociationResultOutput)
}

type LookupLocalGatewayRouteTableVpcAssociationOutputArgs struct {
	// The ID of the association.
	LocalGatewayRouteTableVpcAssociationId pulumi.StringInput `pulumi:"localGatewayRouteTableVpcAssociationId"`
}

func (LookupLocalGatewayRouteTableVpcAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGatewayRouteTableVpcAssociationArgs)(nil)).Elem()
}

type LookupLocalGatewayRouteTableVpcAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupLocalGatewayRouteTableVpcAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalGatewayRouteTableVpcAssociationResult)(nil)).Elem()
}

func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) ToLookupLocalGatewayRouteTableVpcAssociationResultOutput() LookupLocalGatewayRouteTableVpcAssociationResultOutput {
	return o
}

func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) ToLookupLocalGatewayRouteTableVpcAssociationResultOutputWithContext(ctx context.Context) LookupLocalGatewayRouteTableVpcAssociationResultOutput {
	return o
}

// The ID of the local gateway.
func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) LocalGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGatewayRouteTableVpcAssociationResult) *string { return v.LocalGatewayId }).(pulumi.StringPtrOutput)
}

// The ID of the association.
func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) LocalGatewayRouteTableVpcAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGatewayRouteTableVpcAssociationResult) *string {
		return v.LocalGatewayRouteTableVpcAssociationId
	}).(pulumi.StringPtrOutput)
}

// The state of the association.
func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalGatewayRouteTableVpcAssociationResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The tags for the association.
func (o LookupLocalGatewayRouteTableVpcAssociationResultOutput) Tags() LocalGatewayRouteTableVpcAssociationTagArrayOutput {
	return o.ApplyT(func(v LookupLocalGatewayRouteTableVpcAssociationResult) []LocalGatewayRouteTableVpcAssociationTag {
		return v.Tags
	}).(LocalGatewayRouteTableVpcAssociationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalGatewayRouteTableVpcAssociationResultOutput{})
}