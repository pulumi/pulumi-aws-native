// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emrcontainers

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EMRContainers::VirtualCluster Type
type VirtualCluster struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Container provider of the virtual cluster.
	ContainerProvider VirtualClusterContainerProviderOutput `pulumi:"containerProvider"`
	// Name of the virtual cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of key-value pairs to apply to this virtual cluster.
	Tags VirtualClusterTagArrayOutput `pulumi:"tags"`
}

// NewVirtualCluster registers a new resource with the given unique name, arguments, and options.
func NewVirtualCluster(ctx *pulumi.Context,
	name string, args *VirtualClusterArgs, opts ...pulumi.ResourceOption) (*VirtualCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerProvider == nil {
		return nil, errors.New("invalid value for required argument 'ContainerProvider'")
	}
	var resource VirtualCluster
	err := ctx.RegisterResource("aws-native:emrcontainers:VirtualCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualCluster gets an existing VirtualCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualClusterState, opts ...pulumi.ResourceOption) (*VirtualCluster, error) {
	var resource VirtualCluster
	err := ctx.ReadResource("aws-native:emrcontainers:VirtualCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualCluster resources.
type virtualClusterState struct {
}

type VirtualClusterState struct {
}

func (VirtualClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualClusterState)(nil)).Elem()
}

type virtualClusterArgs struct {
	// Container provider of the virtual cluster.
	ContainerProvider VirtualClusterContainerProvider `pulumi:"containerProvider"`
	// Name of the virtual cluster.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this virtual cluster.
	Tags []VirtualClusterTag `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualCluster resource.
type VirtualClusterArgs struct {
	// Container provider of the virtual cluster.
	ContainerProvider VirtualClusterContainerProviderInput
	// Name of the virtual cluster.
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this virtual cluster.
	Tags VirtualClusterTagArrayInput
}

func (VirtualClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualClusterArgs)(nil)).Elem()
}

type VirtualClusterInput interface {
	pulumi.Input

	ToVirtualClusterOutput() VirtualClusterOutput
	ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput
}

func (*VirtualCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCluster)(nil)).Elem()
}

func (i *VirtualCluster) ToVirtualClusterOutput() VirtualClusterOutput {
	return i.ToVirtualClusterOutputWithContext(context.Background())
}

func (i *VirtualCluster) ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterOutput)
}

type VirtualClusterOutput struct{ *pulumi.OutputState }

func (VirtualClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCluster)(nil)).Elem()
}

func (o VirtualClusterOutput) ToVirtualClusterOutput() VirtualClusterOutput {
	return o
}

func (o VirtualClusterOutput) ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput {
	return o
}

func (o VirtualClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Container provider of the virtual cluster.
func (o VirtualClusterOutput) ContainerProvider() VirtualClusterContainerProviderOutput {
	return o.ApplyT(func(v *VirtualCluster) VirtualClusterContainerProviderOutput { return v.ContainerProvider }).(VirtualClusterContainerProviderOutput)
}

// Name of the virtual cluster.
func (o VirtualClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this virtual cluster.
func (o VirtualClusterOutput) Tags() VirtualClusterTagArrayOutput {
	return o.ApplyT(func(v *VirtualCluster) VirtualClusterTagArrayOutput { return v.Tags }).(VirtualClusterTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterInput)(nil)).Elem(), &VirtualCluster{})
	pulumi.RegisterOutputType(VirtualClusterOutput{})
}