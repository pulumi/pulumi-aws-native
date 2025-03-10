// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::DedicatedIpPool
type DedicatedIpPool struct {
	pulumi.CustomResourceState

	// The name of the dedicated IP pool.
	PoolName pulumi.StringPtrOutput `pulumi:"poolName"`
	// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
	ScalingMode pulumi.StringPtrOutput `pulumi:"scalingMode"`
}

// NewDedicatedIpPool registers a new resource with the given unique name, arguments, and options.
func NewDedicatedIpPool(ctx *pulumi.Context,
	name string, args *DedicatedIpPoolArgs, opts ...pulumi.ResourceOption) (*DedicatedIpPool, error) {
	if args == nil {
		args = &DedicatedIpPoolArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"poolName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedIpPool
	err := ctx.RegisterResource("aws-native:ses:DedicatedIpPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedIpPool gets an existing DedicatedIpPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedIpPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedIpPoolState, opts ...pulumi.ResourceOption) (*DedicatedIpPool, error) {
	var resource DedicatedIpPool
	err := ctx.ReadResource("aws-native:ses:DedicatedIpPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedIpPool resources.
type dedicatedIpPoolState struct {
}

type DedicatedIpPoolState struct {
}

func (DedicatedIpPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedIpPoolState)(nil)).Elem()
}

type dedicatedIpPoolArgs struct {
	// The name of the dedicated IP pool.
	PoolName *string `pulumi:"poolName"`
	// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
	ScalingMode *string `pulumi:"scalingMode"`
}

// The set of arguments for constructing a DedicatedIpPool resource.
type DedicatedIpPoolArgs struct {
	// The name of the dedicated IP pool.
	PoolName pulumi.StringPtrInput
	// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
	ScalingMode pulumi.StringPtrInput
}

func (DedicatedIpPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedIpPoolArgs)(nil)).Elem()
}

type DedicatedIpPoolInput interface {
	pulumi.Input

	ToDedicatedIpPoolOutput() DedicatedIpPoolOutput
	ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput
}

func (*DedicatedIpPool) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedIpPool)(nil)).Elem()
}

func (i *DedicatedIpPool) ToDedicatedIpPoolOutput() DedicatedIpPoolOutput {
	return i.ToDedicatedIpPoolOutputWithContext(context.Background())
}

func (i *DedicatedIpPool) ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedIpPoolOutput)
}

type DedicatedIpPoolOutput struct{ *pulumi.OutputState }

func (DedicatedIpPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedIpPool)(nil)).Elem()
}

func (o DedicatedIpPoolOutput) ToDedicatedIpPoolOutput() DedicatedIpPoolOutput {
	return o
}

func (o DedicatedIpPoolOutput) ToDedicatedIpPoolOutputWithContext(ctx context.Context) DedicatedIpPoolOutput {
	return o
}

// The name of the dedicated IP pool.
func (o DedicatedIpPoolOutput) PoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringPtrOutput { return v.PoolName }).(pulumi.StringPtrOutput)
}

// Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD.
func (o DedicatedIpPoolOutput) ScalingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedIpPool) pulumi.StringPtrOutput { return v.ScalingMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedIpPoolInput)(nil)).Elem(), &DedicatedIpPool{})
	pulumi.RegisterOutputType(DedicatedIpPoolOutput{})
}
