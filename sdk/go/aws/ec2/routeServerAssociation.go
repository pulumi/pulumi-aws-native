// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VPC Route Server Association
type RouteServerAssociation struct {
	pulumi.CustomResourceState

	// Route Server ID
	RouteServerId pulumi.StringOutput `pulumi:"routeServerId"`
	// VPC ID
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteServerAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteServerAssociation(ctx *pulumi.Context,
	name string, args *RouteServerAssociationArgs, opts ...pulumi.ResourceOption) (*RouteServerAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteServerId == nil {
		return nil, errors.New("invalid value for required argument 'RouteServerId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"routeServerId",
		"vpcId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteServerAssociation
	err := ctx.RegisterResource("aws-native:ec2:RouteServerAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteServerAssociation gets an existing RouteServerAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteServerAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteServerAssociationState, opts ...pulumi.ResourceOption) (*RouteServerAssociation, error) {
	var resource RouteServerAssociation
	err := ctx.ReadResource("aws-native:ec2:RouteServerAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteServerAssociation resources.
type routeServerAssociationState struct {
}

type RouteServerAssociationState struct {
}

func (RouteServerAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeServerAssociationState)(nil)).Elem()
}

type routeServerAssociationArgs struct {
	// Route Server ID
	RouteServerId string `pulumi:"routeServerId"`
	// VPC ID
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteServerAssociation resource.
type RouteServerAssociationArgs struct {
	// Route Server ID
	RouteServerId pulumi.StringInput
	// VPC ID
	VpcId pulumi.StringInput
}

func (RouteServerAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeServerAssociationArgs)(nil)).Elem()
}

type RouteServerAssociationInput interface {
	pulumi.Input

	ToRouteServerAssociationOutput() RouteServerAssociationOutput
	ToRouteServerAssociationOutputWithContext(ctx context.Context) RouteServerAssociationOutput
}

func (*RouteServerAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteServerAssociation)(nil)).Elem()
}

func (i *RouteServerAssociation) ToRouteServerAssociationOutput() RouteServerAssociationOutput {
	return i.ToRouteServerAssociationOutputWithContext(context.Background())
}

func (i *RouteServerAssociation) ToRouteServerAssociationOutputWithContext(ctx context.Context) RouteServerAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteServerAssociationOutput)
}

type RouteServerAssociationOutput struct{ *pulumi.OutputState }

func (RouteServerAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteServerAssociation)(nil)).Elem()
}

func (o RouteServerAssociationOutput) ToRouteServerAssociationOutput() RouteServerAssociationOutput {
	return o
}

func (o RouteServerAssociationOutput) ToRouteServerAssociationOutputWithContext(ctx context.Context) RouteServerAssociationOutput {
	return o
}

// Route Server ID
func (o RouteServerAssociationOutput) RouteServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteServerAssociation) pulumi.StringOutput { return v.RouteServerId }).(pulumi.StringOutput)
}

// VPC ID
func (o RouteServerAssociationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteServerAssociation) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteServerAssociationInput)(nil)).Elem(), &RouteServerAssociation{})
	pulumi.RegisterOutputType(RouteServerAssociationOutput{})
}
