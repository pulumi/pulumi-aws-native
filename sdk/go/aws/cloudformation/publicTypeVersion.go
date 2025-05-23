// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Test and Publish a resource that has been registered in the CloudFormation Registry.
type PublicTypeVersion struct {
	pulumi.CustomResourceState

	// The Amazon Resource Number (ARN) of the extension.
	Arn pulumi.StringPtrOutput `pulumi:"arn"`
	// A url to the S3 bucket where logs for the testType run will be available
	LogDeliveryBucket pulumi.StringPtrOutput `pulumi:"logDeliveryBucket"`
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication
	PublicTypeArn pulumi.StringOutput `pulumi:"publicTypeArn"`
	// The version number of a public third-party extension
	PublicVersionNumber pulumi.StringPtrOutput `pulumi:"publicVersionNumber"`
	// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
	PublisherId pulumi.StringOutput `pulumi:"publisherId"`
	// The kind of extension
	Type PublicTypeVersionTypePtrOutput `pulumi:"type"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringPtrOutput `pulumi:"typeName"`
	// The Amazon Resource Number (ARN) of the extension with the versionId.
	TypeVersionArn pulumi.StringOutput `pulumi:"typeVersionArn"`
}

// NewPublicTypeVersion registers a new resource with the given unique name, arguments, and options.
func NewPublicTypeVersion(ctx *pulumi.Context,
	name string, args *PublicTypeVersionArgs, opts ...pulumi.ResourceOption) (*PublicTypeVersion, error) {
	if args == nil {
		args = &PublicTypeVersionArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"arn",
		"logDeliveryBucket",
		"publicVersionNumber",
		"type",
		"typeName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicTypeVersion
	err := ctx.RegisterResource("aws-native:cloudformation:PublicTypeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicTypeVersion gets an existing PublicTypeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicTypeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicTypeVersionState, opts ...pulumi.ResourceOption) (*PublicTypeVersion, error) {
	var resource PublicTypeVersion
	err := ctx.ReadResource("aws-native:cloudformation:PublicTypeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicTypeVersion resources.
type publicTypeVersionState struct {
}

type PublicTypeVersionState struct {
}

func (PublicTypeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicTypeVersionState)(nil)).Elem()
}

type publicTypeVersionArgs struct {
	// The Amazon Resource Number (ARN) of the extension.
	Arn *string `pulumi:"arn"`
	// A url to the S3 bucket where logs for the testType run will be available
	LogDeliveryBucket *string `pulumi:"logDeliveryBucket"`
	// The version number of a public third-party extension
	PublicVersionNumber *string `pulumi:"publicVersionNumber"`
	// The kind of extension
	Type *PublicTypeVersionType `pulumi:"type"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName *string `pulumi:"typeName"`
}

// The set of arguments for constructing a PublicTypeVersion resource.
type PublicTypeVersionArgs struct {
	// The Amazon Resource Number (ARN) of the extension.
	Arn pulumi.StringPtrInput
	// A url to the S3 bucket where logs for the testType run will be available
	LogDeliveryBucket pulumi.StringPtrInput
	// The version number of a public third-party extension
	PublicVersionNumber pulumi.StringPtrInput
	// The kind of extension
	Type PublicTypeVersionTypePtrInput
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	TypeName pulumi.StringPtrInput
}

func (PublicTypeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicTypeVersionArgs)(nil)).Elem()
}

type PublicTypeVersionInput interface {
	pulumi.Input

	ToPublicTypeVersionOutput() PublicTypeVersionOutput
	ToPublicTypeVersionOutputWithContext(ctx context.Context) PublicTypeVersionOutput
}

func (*PublicTypeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicTypeVersion)(nil)).Elem()
}

func (i *PublicTypeVersion) ToPublicTypeVersionOutput() PublicTypeVersionOutput {
	return i.ToPublicTypeVersionOutputWithContext(context.Background())
}

func (i *PublicTypeVersion) ToPublicTypeVersionOutputWithContext(ctx context.Context) PublicTypeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicTypeVersionOutput)
}

type PublicTypeVersionOutput struct{ *pulumi.OutputState }

func (PublicTypeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicTypeVersion)(nil)).Elem()
}

func (o PublicTypeVersionOutput) ToPublicTypeVersionOutput() PublicTypeVersionOutput {
	return o
}

func (o PublicTypeVersionOutput) ToPublicTypeVersionOutputWithContext(ctx context.Context) PublicTypeVersionOutput {
	return o
}

// The Amazon Resource Number (ARN) of the extension.
func (o PublicTypeVersionOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringPtrOutput { return v.Arn }).(pulumi.StringPtrOutput)
}

// A url to the S3 bucket where logs for the testType run will be available
func (o PublicTypeVersionOutput) LogDeliveryBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringPtrOutput { return v.LogDeliveryBucket }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Number (ARN) assigned to the public extension upon publication
func (o PublicTypeVersionOutput) PublicTypeArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringOutput { return v.PublicTypeArn }).(pulumi.StringOutput)
}

// The version number of a public third-party extension
func (o PublicTypeVersionOutput) PublicVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringPtrOutput { return v.PublicVersionNumber }).(pulumi.StringPtrOutput)
}

// The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region.
func (o PublicTypeVersionOutput) PublisherId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringOutput { return v.PublisherId }).(pulumi.StringOutput)
}

// The kind of extension
func (o PublicTypeVersionOutput) Type() PublicTypeVersionTypePtrOutput {
	return o.ApplyT(func(v *PublicTypeVersion) PublicTypeVersionTypePtrOutput { return v.Type }).(PublicTypeVersionTypePtrOutput)
}

// The name of the type being registered.
//
// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
func (o PublicTypeVersionOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringPtrOutput { return v.TypeName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Number (ARN) of the extension with the versionId.
func (o PublicTypeVersionOutput) TypeVersionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicTypeVersion) pulumi.StringOutput { return v.TypeVersionArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicTypeVersionInput)(nil)).Elem(), &PublicTypeVersion{})
	pulumi.RegisterOutputType(PublicTypeVersionOutput{})
}
