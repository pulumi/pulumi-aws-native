// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalogappregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
type ResourceAssociation struct {
	pulumi.CustomResourceState

	// The name or the Id of the Application.
	Application pulumi.StringOutput `pulumi:"application"`
	// The Amazon resource name (ARN) that specifies the application.
	ApplicationArn pulumi.StringOutput `pulumi:"applicationArn"`
	// The name or the Id of the Resource.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// The Amazon resource name (ARN) that specifies the resource.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// The type of the CFN Resource for now it's enum CFN_STACK.
	ResourceType ResourceAssociationResourceTypeOutput `pulumi:"resourceType"`
}

// NewResourceAssociation registers a new resource with the given unique name, arguments, and options.
func NewResourceAssociation(ctx *pulumi.Context,
	name string, args *ResourceAssociationArgs, opts ...pulumi.ResourceOption) (*ResourceAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"application",
		"resource",
		"resourceType",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceAssociation
	err := ctx.RegisterResource("aws-native:servicecatalogappregistry:ResourceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceAssociation gets an existing ResourceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceAssociationState, opts ...pulumi.ResourceOption) (*ResourceAssociation, error) {
	var resource ResourceAssociation
	err := ctx.ReadResource("aws-native:servicecatalogappregistry:ResourceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceAssociation resources.
type resourceAssociationState struct {
}

type ResourceAssociationState struct {
}

func (ResourceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAssociationState)(nil)).Elem()
}

type resourceAssociationArgs struct {
	// The name or the Id of the Application.
	Application string `pulumi:"application"`
	// The name or the Id of the Resource.
	Resource string `pulumi:"resource"`
	// The type of the CFN Resource for now it's enum CFN_STACK.
	ResourceType ResourceAssociationResourceType `pulumi:"resourceType"`
}

// The set of arguments for constructing a ResourceAssociation resource.
type ResourceAssociationArgs struct {
	// The name or the Id of the Application.
	Application pulumi.StringInput
	// The name or the Id of the Resource.
	Resource pulumi.StringInput
	// The type of the CFN Resource for now it's enum CFN_STACK.
	ResourceType ResourceAssociationResourceTypeInput
}

func (ResourceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAssociationArgs)(nil)).Elem()
}

type ResourceAssociationInput interface {
	pulumi.Input

	ToResourceAssociationOutput() ResourceAssociationOutput
	ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput
}

func (*ResourceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAssociation)(nil)).Elem()
}

func (i *ResourceAssociation) ToResourceAssociationOutput() ResourceAssociationOutput {
	return i.ToResourceAssociationOutputWithContext(context.Background())
}

func (i *ResourceAssociation) ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAssociationOutput)
}

type ResourceAssociationOutput struct{ *pulumi.OutputState }

func (ResourceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceAssociation)(nil)).Elem()
}

func (o ResourceAssociationOutput) ToResourceAssociationOutput() ResourceAssociationOutput {
	return o
}

func (o ResourceAssociationOutput) ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput {
	return o
}

// The name or the Id of the Application.
func (o ResourceAssociationOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAssociation) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// The Amazon resource name (ARN) that specifies the application.
func (o ResourceAssociationOutput) ApplicationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAssociation) pulumi.StringOutput { return v.ApplicationArn }).(pulumi.StringOutput)
}

// The name or the Id of the Resource.
func (o ResourceAssociationOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAssociation) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

// The Amazon resource name (ARN) that specifies the resource.
func (o ResourceAssociationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceAssociation) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

// The type of the CFN Resource for now it's enum CFN_STACK.
func (o ResourceAssociationOutput) ResourceType() ResourceAssociationResourceTypeOutput {
	return o.ApplyT(func(v *ResourceAssociation) ResourceAssociationResourceTypeOutput { return v.ResourceType }).(ResourceAssociationResourceTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceAssociationInput)(nil)).Elem(), &ResourceAssociation{})
	pulumi.RegisterOutputType(ResourceAssociationOutput{})
}
