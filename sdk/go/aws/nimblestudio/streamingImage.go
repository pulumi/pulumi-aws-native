// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nimblestudio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a streaming session machine image that can be used to launch a streaming session
type StreamingImage struct {
	pulumi.CustomResourceState

	// <p>A human-readable description of the streaming image.</p>
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// <p>The ID of an EC2 machine image with which to create this streaming image.</p>
	Ec2ImageId              pulumi.StringOutput                         `pulumi:"ec2ImageId"`
	EncryptionConfiguration StreamingImageEncryptionConfigurationOutput `pulumi:"encryptionConfiguration"`
	// <p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>
	EulaIds pulumi.StringArrayOutput `pulumi:"eulaIds"`
	// <p>A friendly name for a streaming image resource.</p>
	Name pulumi.StringOutput `pulumi:"name"`
	// <p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>
	Owner pulumi.StringOutput `pulumi:"owner"`
	// <p>The platform of the streaming image, either WINDOWS or LINUX.</p>
	Platform         pulumi.StringOutput `pulumi:"platform"`
	StreamingImageId pulumi.StringOutput `pulumi:"streamingImageId"`
	// <p>The studioId. </p>
	StudioId pulumi.StringOutput         `pulumi:"studioId"`
	Tags     StreamingImageTagsPtrOutput `pulumi:"tags"`
}

// NewStreamingImage registers a new resource with the given unique name, arguments, and options.
func NewStreamingImage(ctx *pulumi.Context,
	name string, args *StreamingImageArgs, opts ...pulumi.ResourceOption) (*StreamingImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ec2ImageId == nil {
		return nil, errors.New("invalid value for required argument 'Ec2ImageId'")
	}
	if args.StudioId == nil {
		return nil, errors.New("invalid value for required argument 'StudioId'")
	}
	var resource StreamingImage
	err := ctx.RegisterResource("aws-native:nimblestudio:StreamingImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingImage gets an existing StreamingImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingImageState, opts ...pulumi.ResourceOption) (*StreamingImage, error) {
	var resource StreamingImage
	err := ctx.ReadResource("aws-native:nimblestudio:StreamingImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingImage resources.
type streamingImageState struct {
}

type StreamingImageState struct {
}

func (StreamingImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingImageState)(nil)).Elem()
}

type streamingImageArgs struct {
	// <p>A human-readable description of the streaming image.</p>
	Description *string `pulumi:"description"`
	// <p>The ID of an EC2 machine image with which to create this streaming image.</p>
	Ec2ImageId string `pulumi:"ec2ImageId"`
	// <p>A friendly name for a streaming image resource.</p>
	Name *string `pulumi:"name"`
	// <p>The studioId. </p>
	StudioId string              `pulumi:"studioId"`
	Tags     *StreamingImageTags `pulumi:"tags"`
}

// The set of arguments for constructing a StreamingImage resource.
type StreamingImageArgs struct {
	// <p>A human-readable description of the streaming image.</p>
	Description pulumi.StringPtrInput
	// <p>The ID of an EC2 machine image with which to create this streaming image.</p>
	Ec2ImageId pulumi.StringInput
	// <p>A friendly name for a streaming image resource.</p>
	Name pulumi.StringPtrInput
	// <p>The studioId. </p>
	StudioId pulumi.StringInput
	Tags     StreamingImageTagsPtrInput
}

func (StreamingImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingImageArgs)(nil)).Elem()
}

type StreamingImageInput interface {
	pulumi.Input

	ToStreamingImageOutput() StreamingImageOutput
	ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput
}

func (*StreamingImage) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingImage)(nil)).Elem()
}

func (i *StreamingImage) ToStreamingImageOutput() StreamingImageOutput {
	return i.ToStreamingImageOutputWithContext(context.Background())
}

func (i *StreamingImage) ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingImageOutput)
}

type StreamingImageOutput struct{ *pulumi.OutputState }

func (StreamingImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingImage)(nil)).Elem()
}

func (o StreamingImageOutput) ToStreamingImageOutput() StreamingImageOutput {
	return o
}

func (o StreamingImageOutput) ToStreamingImageOutputWithContext(ctx context.Context) StreamingImageOutput {
	return o
}

// <p>A human-readable description of the streaming image.</p>
func (o StreamingImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// <p>The ID of an EC2 machine image with which to create this streaming image.</p>
func (o StreamingImageOutput) Ec2ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.Ec2ImageId }).(pulumi.StringOutput)
}

func (o StreamingImageOutput) EncryptionConfiguration() StreamingImageEncryptionConfigurationOutput {
	return o.ApplyT(func(v *StreamingImage) StreamingImageEncryptionConfigurationOutput { return v.EncryptionConfiguration }).(StreamingImageEncryptionConfigurationOutput)
}

// <p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>
func (o StreamingImageOutput) EulaIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringArrayOutput { return v.EulaIds }).(pulumi.StringArrayOutput)
}

// <p>A friendly name for a streaming image resource.</p>
func (o StreamingImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// <p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>
func (o StreamingImageOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// <p>The platform of the streaming image, either WINDOWS or LINUX.</p>
func (o StreamingImageOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

func (o StreamingImageOutput) StreamingImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.StreamingImageId }).(pulumi.StringOutput)
}

// <p>The studioId. </p>
func (o StreamingImageOutput) StudioId() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingImage) pulumi.StringOutput { return v.StudioId }).(pulumi.StringOutput)
}

func (o StreamingImageOutput) Tags() StreamingImageTagsPtrOutput {
	return o.ApplyT(func(v *StreamingImage) StreamingImageTagsPtrOutput { return v.Tags }).(StreamingImageTagsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamingImageInput)(nil)).Elem(), &StreamingImage{})
	pulumi.RegisterOutputType(StreamingImageOutput{})
}