// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An OpenSearch Serverless index resource
type Index struct {
	pulumi.CustomResourceState

	// The endpoint for the collection.
	CollectionEndpoint pulumi.StringOutput `pulumi:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	IndexName pulumi.StringOutput `pulumi:"indexName"`
	// Index Mappings
	Mappings MappingsPropertiesPtrOutput `pulumi:"mappings"`
	// Index settings
	Settings IndexSettingsPtrOutput `pulumi:"settings"`
	// The unique identifier for the index.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectionEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'CollectionEndpoint'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"collectionEndpoint",
		"indexName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("aws-native:opensearchserverless:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("aws-native:opensearchserverless:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// The endpoint for the collection.
	CollectionEndpoint string `pulumi:"collectionEndpoint"`
	// The name of the OpenSearch Serverless index.
	IndexName *string `pulumi:"indexName"`
	// Index Mappings
	Mappings *MappingsProperties `pulumi:"mappings"`
	// Index settings
	Settings *IndexSettings `pulumi:"settings"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// The endpoint for the collection.
	CollectionEndpoint pulumi.StringInput
	// The name of the OpenSearch Serverless index.
	IndexName pulumi.StringPtrInput
	// Index Mappings
	Mappings MappingsPropertiesPtrInput
	// Index settings
	Settings IndexSettingsPtrInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

// The endpoint for the collection.
func (o IndexOutput) CollectionEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.CollectionEndpoint }).(pulumi.StringOutput)
}

// The name of the OpenSearch Serverless index.
func (o IndexOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.IndexName }).(pulumi.StringOutput)
}

// Index Mappings
func (o IndexOutput) Mappings() MappingsPropertiesPtrOutput {
	return o.ApplyT(func(v *Index) MappingsPropertiesPtrOutput { return v.Mappings }).(MappingsPropertiesPtrOutput)
}

// Index settings
func (o IndexOutput) Settings() IndexSettingsPtrOutput {
	return o.ApplyT(func(v *Index) IndexSettingsPtrOutput { return v.Settings }).(IndexSettingsPtrOutput)
}

// The unique identifier for the index.
func (o IndexOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterOutputType(IndexOutput{})
}
