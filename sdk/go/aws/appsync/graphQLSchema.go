// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::GraphQLSchema
//
// Deprecated: GraphQLSchema is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type GraphQLSchema struct {
	pulumi.CustomResourceState

	ApiId                pulumi.StringOutput    `pulumi:"apiId"`
	Definition           pulumi.StringPtrOutput `pulumi:"definition"`
	DefinitionS3Location pulumi.StringPtrOutput `pulumi:"definitionS3Location"`
}

// NewGraphQLSchema registers a new resource with the given unique name, arguments, and options.
func NewGraphQLSchema(ctx *pulumi.Context,
	name string, args *GraphQLSchemaArgs, opts ...pulumi.ResourceOption) (*GraphQLSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource GraphQLSchema
	err := ctx.RegisterResource("aws-native:appsync:GraphQLSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLSchema gets an existing GraphQLSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLSchemaState, opts ...pulumi.ResourceOption) (*GraphQLSchema, error) {
	var resource GraphQLSchema
	err := ctx.ReadResource("aws-native:appsync:GraphQLSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLSchema resources.
type graphQLSchemaState struct {
}

type GraphQLSchemaState struct {
}

func (GraphQLSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLSchemaState)(nil)).Elem()
}

type graphQLSchemaArgs struct {
	ApiId                string  `pulumi:"apiId"`
	Definition           *string `pulumi:"definition"`
	DefinitionS3Location *string `pulumi:"definitionS3Location"`
}

// The set of arguments for constructing a GraphQLSchema resource.
type GraphQLSchemaArgs struct {
	ApiId                pulumi.StringInput
	Definition           pulumi.StringPtrInput
	DefinitionS3Location pulumi.StringPtrInput
}

func (GraphQLSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLSchemaArgs)(nil)).Elem()
}

type GraphQLSchemaInput interface {
	pulumi.Input

	ToGraphQLSchemaOutput() GraphQLSchemaOutput
	ToGraphQLSchemaOutputWithContext(ctx context.Context) GraphQLSchemaOutput
}

func (*GraphQLSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLSchema)(nil)).Elem()
}

func (i *GraphQLSchema) ToGraphQLSchemaOutput() GraphQLSchemaOutput {
	return i.ToGraphQLSchemaOutputWithContext(context.Background())
}

func (i *GraphQLSchema) ToGraphQLSchemaOutputWithContext(ctx context.Context) GraphQLSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLSchemaOutput)
}

type GraphQLSchemaOutput struct{ *pulumi.OutputState }

func (GraphQLSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLSchema)(nil)).Elem()
}

func (o GraphQLSchemaOutput) ToGraphQLSchemaOutput() GraphQLSchemaOutput {
	return o
}

func (o GraphQLSchemaOutput) ToGraphQLSchemaOutputWithContext(ctx context.Context) GraphQLSchemaOutput {
	return o
}

func (o GraphQLSchemaOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLSchema) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

func (o GraphQLSchemaOutput) Definition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLSchema) pulumi.StringPtrOutput { return v.Definition }).(pulumi.StringPtrOutput)
}

func (o GraphQLSchemaOutput) DefinitionS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLSchema) pulumi.StringPtrOutput { return v.DefinitionS3Location }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GraphQLSchemaInput)(nil)).Elem(), &GraphQLSchema{})
	pulumi.RegisterOutputType(GraphQLSchemaOutput{})
}