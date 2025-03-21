// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package qbusiness

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::QBusiness::Index Resource Type
type Index struct {
	pulumi.CustomResourceState

	// The identifier of the Amazon Q Business application using the index.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The capacity units you want to provision for your index. You can add and remove capacity to fit your usage needs.
	CapacityConfiguration IndexCapacityConfigurationPtrOutput `pulumi:"capacityConfiguration"`
	// The Unix timestamp when the index was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A description for the Amazon Q Business index.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the index.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configuration information for document attributes. Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
	//
	// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
	DocumentAttributeConfigurations IndexDocumentAttributeConfigurationArrayOutput `pulumi:"documentAttributeConfigurations"`
	// The Amazon Resource Name (ARN) of an Amazon Q Business index.
	IndexArn pulumi.StringOutput `pulumi:"indexArn"`
	// The identifier for the index.
	IndexId         pulumi.StringOutput   `pulumi:"indexId"`
	IndexStatistics IndexStatisticsOutput `pulumi:"indexStatistics"`
	// The current status of the index. When the status is `ACTIVE` , the index is ready.
	Status IndexStatusOutput `pulumi:"status"`
	// A list of key-value pairs that identify or categorize the index. You can also use tags to help control access to the index. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The index type that's suitable for your needs. For more information on what's included in each type of index, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#index-tiers) .
	Type IndexTypePtrOutput `pulumi:"type"`
	// The Unix timestamp when the index was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationId",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("aws-native:qbusiness:Index", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:qbusiness:Index", name, id, state, &resource, opts...)
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
	// The identifier of the Amazon Q Business application using the index.
	ApplicationId string `pulumi:"applicationId"`
	// The capacity units you want to provision for your index. You can add and remove capacity to fit your usage needs.
	CapacityConfiguration *IndexCapacityConfiguration `pulumi:"capacityConfiguration"`
	// A description for the Amazon Q Business index.
	Description *string `pulumi:"description"`
	// The name of the index.
	DisplayName string `pulumi:"displayName"`
	// Configuration information for document attributes. Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
	//
	// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
	DocumentAttributeConfigurations []IndexDocumentAttributeConfiguration `pulumi:"documentAttributeConfigurations"`
	// A list of key-value pairs that identify or categorize the index. You can also use tags to help control access to the index. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
	Tags []aws.Tag `pulumi:"tags"`
	// The index type that's suitable for your needs. For more information on what's included in each type of index, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#index-tiers) .
	Type *IndexType `pulumi:"type"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// The identifier of the Amazon Q Business application using the index.
	ApplicationId pulumi.StringInput
	// The capacity units you want to provision for your index. You can add and remove capacity to fit your usage needs.
	CapacityConfiguration IndexCapacityConfigurationPtrInput
	// A description for the Amazon Q Business index.
	Description pulumi.StringPtrInput
	// The name of the index.
	DisplayName pulumi.StringInput
	// Configuration information for document attributes. Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
	//
	// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
	DocumentAttributeConfigurations IndexDocumentAttributeConfigurationArrayInput
	// A list of key-value pairs that identify or categorize the index. You can also use tags to help control access to the index. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
	Tags aws.TagArrayInput
	// The index type that's suitable for your needs. For more information on what's included in each type of index, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#index-tiers) .
	Type IndexTypePtrInput
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

// The identifier of the Amazon Q Business application using the index.
func (o IndexOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The capacity units you want to provision for your index. You can add and remove capacity to fit your usage needs.
func (o IndexOutput) CapacityConfiguration() IndexCapacityConfigurationPtrOutput {
	return o.ApplyT(func(v *Index) IndexCapacityConfigurationPtrOutput { return v.CapacityConfiguration }).(IndexCapacityConfigurationPtrOutput)
}

// The Unix timestamp when the index was created.
func (o IndexOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A description for the Amazon Q Business index.
func (o IndexOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Index) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the index.
func (o IndexOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Configuration information for document attributes. Document attributes are metadata or fields associated with your documents. For example, the company department name associated with each document.
//
// For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html) .
func (o IndexOutput) DocumentAttributeConfigurations() IndexDocumentAttributeConfigurationArrayOutput {
	return o.ApplyT(func(v *Index) IndexDocumentAttributeConfigurationArrayOutput {
		return v.DocumentAttributeConfigurations
	}).(IndexDocumentAttributeConfigurationArrayOutput)
}

// The Amazon Resource Name (ARN) of an Amazon Q Business index.
func (o IndexOutput) IndexArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.IndexArn }).(pulumi.StringOutput)
}

// The identifier for the index.
func (o IndexOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.IndexId }).(pulumi.StringOutput)
}

func (o IndexOutput) IndexStatistics() IndexStatisticsOutput {
	return o.ApplyT(func(v *Index) IndexStatisticsOutput { return v.IndexStatistics }).(IndexStatisticsOutput)
}

// The current status of the index. When the status is `ACTIVE` , the index is ready.
func (o IndexOutput) Status() IndexStatusOutput {
	return o.ApplyT(func(v *Index) IndexStatusOutput { return v.Status }).(IndexStatusOutput)
}

// A list of key-value pairs that identify or categorize the index. You can also use tags to help control access to the index. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @.
func (o IndexOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Index) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The index type that's suitable for your needs. For more information on what's included in each type of index, see [Amazon Q Business tiers](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/tiers.html#index-tiers) .
func (o IndexOutput) Type() IndexTypePtrOutput {
	return o.ApplyT(func(v *Index) IndexTypePtrOutput { return v.Type }).(IndexTypePtrOutput)
}

// The Unix timestamp when the index was last updated.
func (o IndexOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterOutputType(IndexOutput{})
}
