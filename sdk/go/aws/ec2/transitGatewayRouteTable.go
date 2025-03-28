// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TransitGatewayRouteTable
type TransitGatewayRouteTable struct {
	pulumi.CustomResourceState

	// Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The ID of the transit gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
	// Transit Gateway Route Table primary identifier
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewTransitGatewayRouteTable registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayRouteTable(ctx *pulumi.Context,
	name string, args *TransitGatewayRouteTableArgs, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"transitGatewayId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransitGatewayRouteTable
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayRouteTable gets an existing TransitGatewayRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayRouteTableState, opts ...pulumi.ResourceOption) (*TransitGatewayRouteTable, error) {
	var resource TransitGatewayRouteTable
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayRouteTable resources.
type transitGatewayRouteTableState struct {
}

type TransitGatewayRouteTableState struct {
}

func (TransitGatewayRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableState)(nil)).Elem()
}

type transitGatewayRouteTableArgs struct {
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
	Tags []aws.Tag `pulumi:"tags"`
	// The ID of the transit gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a TransitGatewayRouteTable resource.
type TransitGatewayRouteTableArgs struct {
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
	Tags aws.TagArrayInput
	// The ID of the transit gateway.
	TransitGatewayId pulumi.StringInput
}

func (TransitGatewayRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayRouteTableArgs)(nil)).Elem()
}

type TransitGatewayRouteTableInput interface {
	pulumi.Input

	ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput
	ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput
}

func (*TransitGatewayRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayRouteTable)(nil)).Elem()
}

func (i *TransitGatewayRouteTable) ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput {
	return i.ToTransitGatewayRouteTableOutputWithContext(context.Background())
}

func (i *TransitGatewayRouteTable) ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayRouteTableOutput)
}

type TransitGatewayRouteTableOutput struct{ *pulumi.OutputState }

func (TransitGatewayRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransitGatewayRouteTable)(nil)).Elem()
}

func (o TransitGatewayRouteTableOutput) ToTransitGatewayRouteTableOutput() TransitGatewayRouteTableOutput {
	return o
}

func (o TransitGatewayRouteTableOutput) ToTransitGatewayRouteTableOutputWithContext(ctx context.Context) TransitGatewayRouteTableOutput {
	return o
}

// Tags are composed of a Key/Value pair. You can use tags to categorize and track each parameter group. The tag value null is permitted.
func (o TransitGatewayRouteTableOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *TransitGatewayRouteTable) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The ID of the transit gateway.
func (o TransitGatewayRouteTableOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayRouteTable) pulumi.StringOutput { return v.TransitGatewayId }).(pulumi.StringOutput)
}

// Transit Gateway Route Table primary identifier
func (o TransitGatewayRouteTableOutput) TransitGatewayRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransitGatewayRouteTable) pulumi.StringOutput { return v.TransitGatewayRouteTableId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayRouteTableInput)(nil)).Elem(), &TransitGatewayRouteTable{})
	pulumi.RegisterOutputType(TransitGatewayRouteTableOutput{})
}
