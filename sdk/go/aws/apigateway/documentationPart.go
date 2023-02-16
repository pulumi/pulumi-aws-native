// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::DocumentationPart
type DocumentationPart struct {
	pulumi.CustomResourceState

	// The identifier of the documentation Part.
	DocumentationPartId pulumi.StringOutput `pulumi:"documentationPartId"`
	// The location of the API entity that the documentation applies to.
	Location DocumentationPartLocationOutput `pulumi:"location"`
	// The documentation content map of the targeted API entity.
	Properties pulumi.StringOutput `pulumi:"properties"`
	// Identifier of the targeted API entity
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewDocumentationPart registers a new resource with the given unique name, arguments, and options.
func NewDocumentationPart(ctx *pulumi.Context,
	name string, args *DocumentationPartArgs, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource DocumentationPart
	err := ctx.RegisterResource("aws-native:apigateway:DocumentationPart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentationPart gets an existing DocumentationPart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentationPart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentationPartState, opts ...pulumi.ResourceOption) (*DocumentationPart, error) {
	var resource DocumentationPart
	err := ctx.ReadResource("aws-native:apigateway:DocumentationPart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentationPart resources.
type documentationPartState struct {
}

type DocumentationPartState struct {
}

func (DocumentationPartState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartState)(nil)).Elem()
}

type documentationPartArgs struct {
	// The location of the API entity that the documentation applies to.
	Location DocumentationPartLocation `pulumi:"location"`
	// The documentation content map of the targeted API entity.
	Properties string `pulumi:"properties"`
	// Identifier of the targeted API entity
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a DocumentationPart resource.
type DocumentationPartArgs struct {
	// The location of the API entity that the documentation applies to.
	Location DocumentationPartLocationInput
	// The documentation content map of the targeted API entity.
	Properties pulumi.StringInput
	// Identifier of the targeted API entity
	RestApiId pulumi.StringInput
}

func (DocumentationPartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentationPartArgs)(nil)).Elem()
}

type DocumentationPartInput interface {
	pulumi.Input

	ToDocumentationPartOutput() DocumentationPartOutput
	ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput
}

func (*DocumentationPart) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentationPart)(nil)).Elem()
}

func (i *DocumentationPart) ToDocumentationPartOutput() DocumentationPartOutput {
	return i.ToDocumentationPartOutputWithContext(context.Background())
}

func (i *DocumentationPart) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentationPartOutput)
}

type DocumentationPartOutput struct{ *pulumi.OutputState }

func (DocumentationPartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentationPart)(nil)).Elem()
}

func (o DocumentationPartOutput) ToDocumentationPartOutput() DocumentationPartOutput {
	return o
}

func (o DocumentationPartOutput) ToDocumentationPartOutputWithContext(ctx context.Context) DocumentationPartOutput {
	return o
}

// The identifier of the documentation Part.
func (o DocumentationPartOutput) DocumentationPartId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentationPart) pulumi.StringOutput { return v.DocumentationPartId }).(pulumi.StringOutput)
}

// The location of the API entity that the documentation applies to.
func (o DocumentationPartOutput) Location() DocumentationPartLocationOutput {
	return o.ApplyT(func(v *DocumentationPart) DocumentationPartLocationOutput { return v.Location }).(DocumentationPartLocationOutput)
}

// The documentation content map of the targeted API entity.
func (o DocumentationPartOutput) Properties() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentationPart) pulumi.StringOutput { return v.Properties }).(pulumi.StringOutput)
}

// Identifier of the targeted API entity
func (o DocumentationPartOutput) RestApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentationPart) pulumi.StringOutput { return v.RestApiId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentationPartInput)(nil)).Elem(), &DocumentationPart{})
	pulumi.RegisterOutputType(DocumentationPartOutput{})
}