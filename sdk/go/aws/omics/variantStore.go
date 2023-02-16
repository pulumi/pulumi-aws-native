// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::VariantStore Resource Type
type VariantStore struct {
	pulumi.CustomResourceState

	CreationTime   pulumi.StringOutput             `pulumi:"creationTime"`
	Description    pulumi.StringPtrOutput          `pulumi:"description"`
	Name           pulumi.StringOutput             `pulumi:"name"`
	Reference      VariantStoreReferenceItemOutput `pulumi:"reference"`
	SseConfig      VariantStoreSseConfigPtrOutput  `pulumi:"sseConfig"`
	Status         VariantStoreStoreStatusOutput   `pulumi:"status"`
	StatusMessage  pulumi.StringOutput             `pulumi:"statusMessage"`
	StoreArn       pulumi.StringOutput             `pulumi:"storeArn"`
	StoreSizeBytes pulumi.Float64Output            `pulumi:"storeSizeBytes"`
	Tags           VariantStoreTagMapPtrOutput     `pulumi:"tags"`
	UpdateTime     pulumi.StringOutput             `pulumi:"updateTime"`
}

// NewVariantStore registers a new resource with the given unique name, arguments, and options.
func NewVariantStore(ctx *pulumi.Context,
	name string, args *VariantStoreArgs, opts ...pulumi.ResourceOption) (*VariantStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Reference == nil {
		return nil, errors.New("invalid value for required argument 'Reference'")
	}
	var resource VariantStore
	err := ctx.RegisterResource("aws-native:omics:VariantStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariantStore gets an existing VariantStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariantStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariantStoreState, opts ...pulumi.ResourceOption) (*VariantStore, error) {
	var resource VariantStore
	err := ctx.ReadResource("aws-native:omics:VariantStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VariantStore resources.
type variantStoreState struct {
}

type VariantStoreState struct {
}

func (VariantStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*variantStoreState)(nil)).Elem()
}

type variantStoreArgs struct {
	Description *string                   `pulumi:"description"`
	Name        *string                   `pulumi:"name"`
	Reference   VariantStoreReferenceItem `pulumi:"reference"`
	SseConfig   *VariantStoreSseConfig    `pulumi:"sseConfig"`
	Tags        *VariantStoreTagMap       `pulumi:"tags"`
}

// The set of arguments for constructing a VariantStore resource.
type VariantStoreArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Reference   VariantStoreReferenceItemInput
	SseConfig   VariantStoreSseConfigPtrInput
	Tags        VariantStoreTagMapPtrInput
}

func (VariantStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variantStoreArgs)(nil)).Elem()
}

type VariantStoreInput interface {
	pulumi.Input

	ToVariantStoreOutput() VariantStoreOutput
	ToVariantStoreOutputWithContext(ctx context.Context) VariantStoreOutput
}

func (*VariantStore) ElementType() reflect.Type {
	return reflect.TypeOf((**VariantStore)(nil)).Elem()
}

func (i *VariantStore) ToVariantStoreOutput() VariantStoreOutput {
	return i.ToVariantStoreOutputWithContext(context.Background())
}

func (i *VariantStore) ToVariantStoreOutputWithContext(ctx context.Context) VariantStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariantStoreOutput)
}

type VariantStoreOutput struct{ *pulumi.OutputState }

func (VariantStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariantStore)(nil)).Elem()
}

func (o VariantStoreOutput) ToVariantStoreOutput() VariantStoreOutput {
	return o
}

func (o VariantStoreOutput) ToVariantStoreOutputWithContext(ctx context.Context) VariantStoreOutput {
	return o
}

func (o VariantStoreOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o VariantStoreOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VariantStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VariantStoreOutput) Reference() VariantStoreReferenceItemOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreReferenceItemOutput { return v.Reference }).(VariantStoreReferenceItemOutput)
}

func (o VariantStoreOutput) SseConfig() VariantStoreSseConfigPtrOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreSseConfigPtrOutput { return v.SseConfig }).(VariantStoreSseConfigPtrOutput)
}

func (o VariantStoreOutput) Status() VariantStoreStoreStatusOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreStoreStatusOutput { return v.Status }).(VariantStoreStoreStatusOutput)
}

func (o VariantStoreOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o VariantStoreOutput) StoreArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.StoreArn }).(pulumi.StringOutput)
}

func (o VariantStoreOutput) StoreSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *VariantStore) pulumi.Float64Output { return v.StoreSizeBytes }).(pulumi.Float64Output)
}

func (o VariantStoreOutput) Tags() VariantStoreTagMapPtrOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreTagMapPtrOutput { return v.Tags }).(VariantStoreTagMapPtrOutput)
}

func (o VariantStoreOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VariantStoreInput)(nil)).Elem(), &VariantStore{})
	pulumi.RegisterOutputType(VariantStoreOutput{})
}