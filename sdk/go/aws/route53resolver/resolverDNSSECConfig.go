// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
type ResolverDNSSECConfig struct {
	pulumi.CustomResourceState

	// AccountId
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// ResourceId
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
	ValidationStatus ResolverDNSSECConfigValidationStatusOutput `pulumi:"validationStatus"`
}

// NewResolverDNSSECConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverDNSSECConfig(ctx *pulumi.Context,
	name string, args *ResolverDNSSECConfigArgs, opts ...pulumi.ResourceOption) (*ResolverDNSSECConfig, error) {
	if args == nil {
		args = &ResolverDNSSECConfigArgs{}
	}

	var resource ResolverDNSSECConfig
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverDNSSECConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverDNSSECConfig gets an existing ResolverDNSSECConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverDNSSECConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverDNSSECConfigState, opts ...pulumi.ResourceOption) (*ResolverDNSSECConfig, error) {
	var resource ResolverDNSSECConfig
	err := ctx.ReadResource("aws-native:route53resolver:ResolverDNSSECConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverDNSSECConfig resources.
type resolverDNSSECConfigState struct {
}

type ResolverDNSSECConfigState struct {
}

func (ResolverDNSSECConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDNSSECConfigState)(nil)).Elem()
}

type resolverDNSSECConfigArgs struct {
	// ResourceId
	ResourceId *string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverDNSSECConfig resource.
type ResolverDNSSECConfigArgs struct {
	// ResourceId
	ResourceId pulumi.StringPtrInput
}

func (ResolverDNSSECConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDNSSECConfigArgs)(nil)).Elem()
}

type ResolverDNSSECConfigInput interface {
	pulumi.Input

	ToResolverDNSSECConfigOutput() ResolverDNSSECConfigOutput
	ToResolverDNSSECConfigOutputWithContext(ctx context.Context) ResolverDNSSECConfigOutput
}

func (*ResolverDNSSECConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDNSSECConfig)(nil)).Elem()
}

func (i *ResolverDNSSECConfig) ToResolverDNSSECConfigOutput() ResolverDNSSECConfigOutput {
	return i.ToResolverDNSSECConfigOutputWithContext(context.Background())
}

func (i *ResolverDNSSECConfig) ToResolverDNSSECConfigOutputWithContext(ctx context.Context) ResolverDNSSECConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDNSSECConfigOutput)
}

type ResolverDNSSECConfigOutput struct{ *pulumi.OutputState }

func (ResolverDNSSECConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDNSSECConfig)(nil)).Elem()
}

func (o ResolverDNSSECConfigOutput) ToResolverDNSSECConfigOutput() ResolverDNSSECConfigOutput {
	return o
}

func (o ResolverDNSSECConfigOutput) ToResolverDNSSECConfigOutputWithContext(ctx context.Context) ResolverDNSSECConfigOutput {
	return o
}

// AccountId
func (o ResolverDNSSECConfigOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverDNSSECConfig) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// ResourceId
func (o ResolverDNSSECConfigOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverDNSSECConfig) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
func (o ResolverDNSSECConfigOutput) ValidationStatus() ResolverDNSSECConfigValidationStatusOutput {
	return o.ApplyT(func(v *ResolverDNSSECConfig) ResolverDNSSECConfigValidationStatusOutput { return v.ValidationStatus }).(ResolverDNSSECConfigValidationStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDNSSECConfigInput)(nil)).Elem(), &ResolverDNSSECConfig{})
	pulumi.RegisterOutputType(ResolverDNSSECConfigOutput{})
}