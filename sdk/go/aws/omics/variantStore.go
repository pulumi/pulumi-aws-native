// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::VariantStore Resource Type
type VariantStore struct {
	pulumi.CustomResourceState

	// The store's ID.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// When the store was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A description for the store.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A name for the store.
	Name pulumi.StringOutput `pulumi:"name"`
	// The genome reference for the store's variants.
	Reference VariantStoreReferenceItemOutput `pulumi:"reference"`
	// Server-side encryption (SSE) settings for the store.
	SseConfig VariantStoreSseConfigPtrOutput `pulumi:"sseConfig"`
	// The store's status.
	Status VariantStoreStoreStatusOutput `pulumi:"status"`
	// The store's status message.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// The store's ARN.
	StoreArn pulumi.StringOutput `pulumi:"storeArn"`
	// The store's size in bytes.
	StoreSizeBytes pulumi.Float64Output `pulumi:"storeSizeBytes"`
	// Tags for the store.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// When the store was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
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
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"reference",
		"sseConfig",
		"tags.*",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// A description for the store.
	Description *string `pulumi:"description"`
	// A name for the store.
	Name *string `pulumi:"name"`
	// The genome reference for the store's variants.
	Reference VariantStoreReferenceItem `pulumi:"reference"`
	// Server-side encryption (SSE) settings for the store.
	SseConfig *VariantStoreSseConfig `pulumi:"sseConfig"`
	// Tags for the store.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VariantStore resource.
type VariantStoreArgs struct {
	// A description for the store.
	Description pulumi.StringPtrInput
	// A name for the store.
	Name pulumi.StringPtrInput
	// The genome reference for the store's variants.
	Reference VariantStoreReferenceItemInput
	// Server-side encryption (SSE) settings for the store.
	SseConfig VariantStoreSseConfigPtrInput
	// Tags for the store.
	Tags pulumi.StringMapInput
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

// The store's ID.
func (o VariantStoreOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// When the store was created.
func (o VariantStoreOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A description for the store.
func (o VariantStoreOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A name for the store.
func (o VariantStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The genome reference for the store's variants.
func (o VariantStoreOutput) Reference() VariantStoreReferenceItemOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreReferenceItemOutput { return v.Reference }).(VariantStoreReferenceItemOutput)
}

// Server-side encryption (SSE) settings for the store.
func (o VariantStoreOutput) SseConfig() VariantStoreSseConfigPtrOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreSseConfigPtrOutput { return v.SseConfig }).(VariantStoreSseConfigPtrOutput)
}

// The store's status.
func (o VariantStoreOutput) Status() VariantStoreStoreStatusOutput {
	return o.ApplyT(func(v *VariantStore) VariantStoreStoreStatusOutput { return v.Status }).(VariantStoreStoreStatusOutput)
}

// The store's status message.
func (o VariantStoreOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// The store's ARN.
func (o VariantStoreOutput) StoreArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.StoreArn }).(pulumi.StringOutput)
}

// The store's size in bytes.
func (o VariantStoreOutput) StoreSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *VariantStore) pulumi.Float64Output { return v.StoreSizeBytes }).(pulumi.Float64Output)
}

// Tags for the store.
func (o VariantStoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// When the store was updated.
func (o VariantStoreOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VariantStore) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VariantStoreInput)(nil)).Elem(), &VariantStore{})
	pulumi.RegisterOutputType(VariantStoreOutput{})
}
