// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.
func LookupGatewayRouteTableAssociation(ctx *pulumi.Context, args *LookupGatewayRouteTableAssociationArgs, opts ...pulumi.InvokeOption) (*LookupGatewayRouteTableAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayRouteTableAssociationResult
	err := ctx.Invoke("aws-native:ec2:getGatewayRouteTableAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayRouteTableAssociationArgs struct {
	// The ID of the gateway.
	GatewayId string `pulumi:"gatewayId"`
}

type LookupGatewayRouteTableAssociationResult struct {
	// The route table association ID.
	AssociationId *string `pulumi:"associationId"`
	// The ID of the route table.
	RouteTableId *string `pulumi:"routeTableId"`
}

func LookupGatewayRouteTableAssociationOutput(ctx *pulumi.Context, args LookupGatewayRouteTableAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayRouteTableAssociationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGatewayRouteTableAssociationResultOutput, error) {
			args := v.(LookupGatewayRouteTableAssociationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getGatewayRouteTableAssociation", args, LookupGatewayRouteTableAssociationResultOutput{}, options).(LookupGatewayRouteTableAssociationResultOutput), nil
		}).(LookupGatewayRouteTableAssociationResultOutput)
}

type LookupGatewayRouteTableAssociationOutputArgs struct {
	// The ID of the gateway.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
}

func (LookupGatewayRouteTableAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteTableAssociationArgs)(nil)).Elem()
}

type LookupGatewayRouteTableAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayRouteTableAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayRouteTableAssociationResult)(nil)).Elem()
}

func (o LookupGatewayRouteTableAssociationResultOutput) ToLookupGatewayRouteTableAssociationResultOutput() LookupGatewayRouteTableAssociationResultOutput {
	return o
}

func (o LookupGatewayRouteTableAssociationResultOutput) ToLookupGatewayRouteTableAssociationResultOutputWithContext(ctx context.Context) LookupGatewayRouteTableAssociationResultOutput {
	return o
}

// The route table association ID.
func (o LookupGatewayRouteTableAssociationResultOutput) AssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayRouteTableAssociationResult) *string { return v.AssociationId }).(pulumi.StringPtrOutput)
}

// The ID of the route table.
func (o LookupGatewayRouteTableAssociationResultOutput) RouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayRouteTableAssociationResult) *string { return v.RouteTableId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayRouteTableAssociationResultOutput{})
}
