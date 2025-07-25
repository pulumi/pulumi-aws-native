// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.
type HubV2 struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the Security Hub V2 resource.
	HubV2Arn pulumi.StringOutput `pulumi:"hubV2Arn"`
	// The date and time when the service was enabled in the account.
	SubscribedAt pulumi.StringOutput `pulumi:"subscribedAt"`
	// The tags to add to the hub V2 resource when you enable Security Hub.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewHubV2 registers a new resource with the given unique name, arguments, and options.
func NewHubV2(ctx *pulumi.Context,
	name string, args *HubV2Args, opts ...pulumi.ResourceOption) (*HubV2, error) {
	if args == nil {
		args = &HubV2Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HubV2
	err := ctx.RegisterResource("aws-native:securityhub:HubV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubV2 gets an existing HubV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubV2State, opts ...pulumi.ResourceOption) (*HubV2, error) {
	var resource HubV2
	err := ctx.ReadResource("aws-native:securityhub:HubV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HubV2 resources.
type hubV2State struct {
}

type HubV2State struct {
}

func (HubV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*hubV2State)(nil)).Elem()
}

type hubV2Args struct {
	// The tags to add to the hub V2 resource when you enable Security Hub.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a HubV2 resource.
type HubV2Args struct {
	// The tags to add to the hub V2 resource when you enable Security Hub.
	Tags pulumi.StringMapInput
}

func (HubV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*hubV2Args)(nil)).Elem()
}

type HubV2Input interface {
	pulumi.Input

	ToHubV2Output() HubV2Output
	ToHubV2OutputWithContext(ctx context.Context) HubV2Output
}

func (*HubV2) ElementType() reflect.Type {
	return reflect.TypeOf((**HubV2)(nil)).Elem()
}

func (i *HubV2) ToHubV2Output() HubV2Output {
	return i.ToHubV2OutputWithContext(context.Background())
}

func (i *HubV2) ToHubV2OutputWithContext(ctx context.Context) HubV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(HubV2Output)
}

type HubV2Output struct{ *pulumi.OutputState }

func (HubV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**HubV2)(nil)).Elem()
}

func (o HubV2Output) ToHubV2Output() HubV2Output {
	return o
}

func (o HubV2Output) ToHubV2OutputWithContext(ctx context.Context) HubV2Output {
	return o
}

// The Amazon Resource Name of the Security Hub V2 resource.
func (o HubV2Output) HubV2Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *HubV2) pulumi.StringOutput { return v.HubV2Arn }).(pulumi.StringOutput)
}

// The date and time when the service was enabled in the account.
func (o HubV2Output) SubscribedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *HubV2) pulumi.StringOutput { return v.SubscribedAt }).(pulumi.StringOutput)
}

// The tags to add to the hub V2 resource when you enable Security Hub.
func (o HubV2Output) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HubV2) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HubV2Input)(nil)).Elem(), &HubV2{})
	pulumi.RegisterOutputType(HubV2Output{})
}
