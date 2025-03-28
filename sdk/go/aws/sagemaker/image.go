// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Image
type Image struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the image.
	//
	// *Type* : String
	//
	// *Length Constraints* : Maximum length of 256.
	//
	// *Pattern* : `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$`
	ImageArn pulumi.StringOutput `pulumi:"imageArn"`
	// The description of the image.
	ImageDescription pulumi.StringPtrOutput `pulumi:"imageDescription"`
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName pulumi.StringPtrOutput `pulumi:"imageDisplayName"`
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn pulumi.StringOutput `pulumi:"imageRoleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ImageRoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"imageName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("aws-native:sagemaker:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("aws-native:sagemaker:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The description of the image.
	ImageDescription *string `pulumi:"imageDescription"`
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName *string `pulumi:"imageDisplayName"`
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName *string `pulumi:"imageName"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn string `pulumi:"imageRoleArn"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The description of the image.
	ImageDescription pulumi.StringPtrInput
	// The display name of the image.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	//
	// *Pattern* : `^\S(.*\S)?$`
	ImageDisplayName pulumi.StringPtrInput
	// The name of the Image. Must be unique by region in your account.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 63.
	//
	// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
	ImageName pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	//
	// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
	//
	// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
	ImageRoleArn pulumi.StringInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The Amazon Resource Name (ARN) of the image.
//
// *Type* : String
//
// *Length Constraints* : Maximum length of 256.
//
// *Pattern* : `^arn:aws(-[\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$`
func (o ImageOutput) ImageArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageArn }).(pulumi.StringOutput)
}

// The description of the image.
func (o ImageOutput) ImageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageDescription }).(pulumi.StringPtrOutput)
}

// The display name of the image.
//
// *Length Constraints* : Minimum length of 1. Maximum length of 128.
//
// *Pattern* : `^\S(.*\S)?$`
func (o ImageOutput) ImageDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageDisplayName }).(pulumi.StringPtrOutput)
}

// The name of the Image. Must be unique by region in your account.
//
// *Length Constraints* : Minimum length of 1. Maximum length of 63.
//
// *Pattern* : `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`
func (o ImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
//
// *Length Constraints* : Minimum length of 20. Maximum length of 2048.
//
// *Pattern* : `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`
func (o ImageOutput) ImageRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageRoleArn }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ImageOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Image) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterOutputType(ImageOutput{})
}
