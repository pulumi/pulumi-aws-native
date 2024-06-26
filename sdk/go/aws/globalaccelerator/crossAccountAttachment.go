// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment
type CrossAccountAttachment struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the attachment.
	AttachmentArn pulumi.StringOutput `pulumi:"attachmentArn"`
	// The Friendly identifier of the attachment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Principals to share the resources with.
	Principals pulumi.StringArrayOutput `pulumi:"principals"`
	// Resources shared using the attachment.
	Resources CrossAccountAttachmentResourceArrayOutput `pulumi:"resources"`
	// Add tags for a cross-account attachment.
	//
	// For more information, see [Tagging in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewCrossAccountAttachment registers a new resource with the given unique name, arguments, and options.
func NewCrossAccountAttachment(ctx *pulumi.Context,
	name string, args *CrossAccountAttachmentArgs, opts ...pulumi.ResourceOption) (*CrossAccountAttachment, error) {
	if args == nil {
		args = &CrossAccountAttachmentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CrossAccountAttachment
	err := ctx.RegisterResource("aws-native:globalaccelerator:CrossAccountAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCrossAccountAttachment gets an existing CrossAccountAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCrossAccountAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CrossAccountAttachmentState, opts ...pulumi.ResourceOption) (*CrossAccountAttachment, error) {
	var resource CrossAccountAttachment
	err := ctx.ReadResource("aws-native:globalaccelerator:CrossAccountAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CrossAccountAttachment resources.
type crossAccountAttachmentState struct {
}

type CrossAccountAttachmentState struct {
}

func (CrossAccountAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*crossAccountAttachmentState)(nil)).Elem()
}

type crossAccountAttachmentArgs struct {
	// The Friendly identifier of the attachment.
	Name *string `pulumi:"name"`
	// Principals to share the resources with.
	Principals []string `pulumi:"principals"`
	// Resources shared using the attachment.
	Resources []CrossAccountAttachmentResource `pulumi:"resources"`
	// Add tags for a cross-account attachment.
	//
	// For more information, see [Tagging in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a CrossAccountAttachment resource.
type CrossAccountAttachmentArgs struct {
	// The Friendly identifier of the attachment.
	Name pulumi.StringPtrInput
	// Principals to share the resources with.
	Principals pulumi.StringArrayInput
	// Resources shared using the attachment.
	Resources CrossAccountAttachmentResourceArrayInput
	// Add tags for a cross-account attachment.
	//
	// For more information, see [Tagging in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
	Tags aws.TagArrayInput
}

func (CrossAccountAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*crossAccountAttachmentArgs)(nil)).Elem()
}

type CrossAccountAttachmentInput interface {
	pulumi.Input

	ToCrossAccountAttachmentOutput() CrossAccountAttachmentOutput
	ToCrossAccountAttachmentOutputWithContext(ctx context.Context) CrossAccountAttachmentOutput
}

func (*CrossAccountAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossAccountAttachment)(nil)).Elem()
}

func (i *CrossAccountAttachment) ToCrossAccountAttachmentOutput() CrossAccountAttachmentOutput {
	return i.ToCrossAccountAttachmentOutputWithContext(context.Background())
}

func (i *CrossAccountAttachment) ToCrossAccountAttachmentOutputWithContext(ctx context.Context) CrossAccountAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossAccountAttachmentOutput)
}

type CrossAccountAttachmentOutput struct{ *pulumi.OutputState }

func (CrossAccountAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossAccountAttachment)(nil)).Elem()
}

func (o CrossAccountAttachmentOutput) ToCrossAccountAttachmentOutput() CrossAccountAttachmentOutput {
	return o
}

func (o CrossAccountAttachmentOutput) ToCrossAccountAttachmentOutputWithContext(ctx context.Context) CrossAccountAttachmentOutput {
	return o
}

// The Amazon Resource Name (ARN) of the attachment.
func (o CrossAccountAttachmentOutput) AttachmentArn() pulumi.StringOutput {
	return o.ApplyT(func(v *CrossAccountAttachment) pulumi.StringOutput { return v.AttachmentArn }).(pulumi.StringOutput)
}

// The Friendly identifier of the attachment.
func (o CrossAccountAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CrossAccountAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Principals to share the resources with.
func (o CrossAccountAttachmentOutput) Principals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CrossAccountAttachment) pulumi.StringArrayOutput { return v.Principals }).(pulumi.StringArrayOutput)
}

// Resources shared using the attachment.
func (o CrossAccountAttachmentOutput) Resources() CrossAccountAttachmentResourceArrayOutput {
	return o.ApplyT(func(v *CrossAccountAttachment) CrossAccountAttachmentResourceArrayOutput { return v.Resources }).(CrossAccountAttachmentResourceArrayOutput)
}

// Add tags for a cross-account attachment.
//
// For more information, see [Tagging in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
func (o CrossAccountAttachmentOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *CrossAccountAttachment) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CrossAccountAttachmentInput)(nil)).Elem(), &CrossAccountAttachment{})
	pulumi.RegisterOutputType(CrossAccountAttachmentOutput{})
}
