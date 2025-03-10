// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package entityresolution

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SchemaMapping defined in AWS Entity Resolution service
type SchemaMapping struct {
	pulumi.CustomResourceState

	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the SchemaMapping
	Description  pulumi.StringPtrOutput `pulumi:"description"`
	HasWorkflows pulumi.BoolOutput      `pulumi:"hasWorkflows"`
	// The SchemaMapping attributes input
	MappedInputFields SchemaMappingSchemaInputAttributeArrayOutput `pulumi:"mappedInputFields"`
	SchemaArn         pulumi.StringOutput                          `pulumi:"schemaArn"`
	// The name of the SchemaMapping
	SchemaName pulumi.StringOutput `pulumi:"schemaName"`
	// The tags used to organize, track, or control access for this resource.
	Tags      aws.TagArrayOutput  `pulumi:"tags"`
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSchemaMapping registers a new resource with the given unique name, arguments, and options.
func NewSchemaMapping(ctx *pulumi.Context,
	name string, args *SchemaMappingArgs, opts ...pulumi.ResourceOption) (*SchemaMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MappedInputFields == nil {
		return nil, errors.New("invalid value for required argument 'MappedInputFields'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"schemaName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SchemaMapping
	err := ctx.RegisterResource("aws-native:entityresolution:SchemaMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchemaMapping gets an existing SchemaMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchemaMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaMappingState, opts ...pulumi.ResourceOption) (*SchemaMapping, error) {
	var resource SchemaMapping
	err := ctx.ReadResource("aws-native:entityresolution:SchemaMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchemaMapping resources.
type schemaMappingState struct {
}

type SchemaMappingState struct {
}

func (SchemaMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaMappingState)(nil)).Elem()
}

type schemaMappingArgs struct {
	// The description of the SchemaMapping
	Description *string `pulumi:"description"`
	// The SchemaMapping attributes input
	MappedInputFields []SchemaMappingSchemaInputAttribute `pulumi:"mappedInputFields"`
	// The name of the SchemaMapping
	SchemaName string `pulumi:"schemaName"`
	// The tags used to organize, track, or control access for this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a SchemaMapping resource.
type SchemaMappingArgs struct {
	// The description of the SchemaMapping
	Description pulumi.StringPtrInput
	// The SchemaMapping attributes input
	MappedInputFields SchemaMappingSchemaInputAttributeArrayInput
	// The name of the SchemaMapping
	SchemaName pulumi.StringInput
	// The tags used to organize, track, or control access for this resource.
	Tags aws.TagArrayInput
}

func (SchemaMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaMappingArgs)(nil)).Elem()
}

type SchemaMappingInput interface {
	pulumi.Input

	ToSchemaMappingOutput() SchemaMappingOutput
	ToSchemaMappingOutputWithContext(ctx context.Context) SchemaMappingOutput
}

func (*SchemaMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaMapping)(nil)).Elem()
}

func (i *SchemaMapping) ToSchemaMappingOutput() SchemaMappingOutput {
	return i.ToSchemaMappingOutputWithContext(context.Background())
}

func (i *SchemaMapping) ToSchemaMappingOutputWithContext(ctx context.Context) SchemaMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaMappingOutput)
}

type SchemaMappingOutput struct{ *pulumi.OutputState }

func (SchemaMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaMapping)(nil)).Elem()
}

func (o SchemaMappingOutput) ToSchemaMappingOutput() SchemaMappingOutput {
	return o
}

func (o SchemaMappingOutput) ToSchemaMappingOutputWithContext(ctx context.Context) SchemaMappingOutput {
	return o
}

func (o SchemaMappingOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the SchemaMapping
func (o SchemaMappingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchemaMappingOutput) HasWorkflows() pulumi.BoolOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.BoolOutput { return v.HasWorkflows }).(pulumi.BoolOutput)
}

// The SchemaMapping attributes input
func (o SchemaMappingOutput) MappedInputFields() SchemaMappingSchemaInputAttributeArrayOutput {
	return o.ApplyT(func(v *SchemaMapping) SchemaMappingSchemaInputAttributeArrayOutput { return v.MappedInputFields }).(SchemaMappingSchemaInputAttributeArrayOutput)
}

func (o SchemaMappingOutput) SchemaArn() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.StringOutput { return v.SchemaArn }).(pulumi.StringOutput)
}

// The name of the SchemaMapping
func (o SchemaMappingOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

// The tags used to organize, track, or control access for this resource.
func (o SchemaMappingOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *SchemaMapping) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o SchemaMappingOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaMapping) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaMappingInput)(nil)).Elem(), &SchemaMapping{})
	pulumi.RegisterOutputType(SchemaMappingOutput{})
}
