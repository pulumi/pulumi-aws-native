// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
type ResolverDnssecConfig struct {
	pulumi.CustomResourceState

	// Id
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// AccountId
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// ResourceId
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
	ValidationStatus ResolverDnssecConfigValidationStatusOutput `pulumi:"validationStatus"`
}

// NewResolverDnssecConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverDnssecConfig(ctx *pulumi.Context,
	name string, args *ResolverDnssecConfigArgs, opts ...pulumi.ResourceOption) (*ResolverDnssecConfig, error) {
	if args == nil {
		args = &ResolverDnssecConfigArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverDnssecConfig
	err := ctx.RegisterResource("aws-native:route53resolver:ResolverDnssecConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverDnssecConfig gets an existing ResolverDnssecConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverDnssecConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverDnssecConfigState, opts ...pulumi.ResourceOption) (*ResolverDnssecConfig, error) {
	var resource ResolverDnssecConfig
	err := ctx.ReadResource("aws-native:route53resolver:ResolverDnssecConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverDnssecConfig resources.
type resolverDnssecConfigState struct {
}

type ResolverDnssecConfigState struct {
}

func (ResolverDnssecConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDnssecConfigState)(nil)).Elem()
}

type resolverDnssecConfigArgs struct {
	// ResourceId
	ResourceId *string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverDnssecConfig resource.
type ResolverDnssecConfigArgs struct {
	// ResourceId
	ResourceId pulumi.StringPtrInput
}

func (ResolverDnssecConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverDnssecConfigArgs)(nil)).Elem()
}

type ResolverDnssecConfigInput interface {
	pulumi.Input

	ToResolverDnssecConfigOutput() ResolverDnssecConfigOutput
	ToResolverDnssecConfigOutputWithContext(ctx context.Context) ResolverDnssecConfigOutput
}

func (*ResolverDnssecConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDnssecConfig)(nil)).Elem()
}

func (i *ResolverDnssecConfig) ToResolverDnssecConfigOutput() ResolverDnssecConfigOutput {
	return i.ToResolverDnssecConfigOutputWithContext(context.Background())
}

func (i *ResolverDnssecConfig) ToResolverDnssecConfigOutputWithContext(ctx context.Context) ResolverDnssecConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverDnssecConfigOutput)
}

type ResolverDnssecConfigOutput struct{ *pulumi.OutputState }

func (ResolverDnssecConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverDnssecConfig)(nil)).Elem()
}

func (o ResolverDnssecConfigOutput) ToResolverDnssecConfigOutput() ResolverDnssecConfigOutput {
	return o
}

func (o ResolverDnssecConfigOutput) ToResolverDnssecConfigOutputWithContext(ctx context.Context) ResolverDnssecConfigOutput {
	return o
}

// Id
func (o ResolverDnssecConfigOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverDnssecConfig) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// AccountId
func (o ResolverDnssecConfigOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverDnssecConfig) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// ResourceId
func (o ResolverDnssecConfigOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResolverDnssecConfig) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
func (o ResolverDnssecConfigOutput) ValidationStatus() ResolverDnssecConfigValidationStatusOutput {
	return o.ApplyT(func(v *ResolverDnssecConfig) ResolverDnssecConfigValidationStatusOutput { return v.ValidationStatus }).(ResolverDnssecConfigValidationStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverDnssecConfigInput)(nil)).Elem(), &ResolverDnssecConfig{})
	pulumi.RegisterOutputType(ResolverDnssecConfigOutput{})
}
