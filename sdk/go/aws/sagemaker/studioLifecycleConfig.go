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

// Resource Type definition for AWS::SageMaker::StudioLifecycleConfig
type StudioLifecycleConfig struct {
	pulumi.CustomResourceState

	// The App type that the Lifecycle Configuration is attached to.
	StudioLifecycleConfigAppType StudioLifecycleConfigAppTypeOutput `pulumi:"studioLifecycleConfigAppType"`
	// The Amazon Resource Name (ARN) of the Lifecycle Configuration.
	StudioLifecycleConfigArn pulumi.StringOutput `pulumi:"studioLifecycleConfigArn"`
	// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent pulumi.StringOutput `pulumi:"studioLifecycleConfigContent"`
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	StudioLifecycleConfigName pulumi.StringOutput `pulumi:"studioLifecycleConfigName"`
	// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
}

// NewStudioLifecycleConfig registers a new resource with the given unique name, arguments, and options.
func NewStudioLifecycleConfig(ctx *pulumi.Context,
	name string, args *StudioLifecycleConfigArgs, opts ...pulumi.ResourceOption) (*StudioLifecycleConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StudioLifecycleConfigAppType == nil {
		return nil, errors.New("invalid value for required argument 'StudioLifecycleConfigAppType'")
	}
	if args.StudioLifecycleConfigContent == nil {
		return nil, errors.New("invalid value for required argument 'StudioLifecycleConfigContent'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"studioLifecycleConfigAppType",
		"studioLifecycleConfigContent",
		"studioLifecycleConfigName",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StudioLifecycleConfig
	err := ctx.RegisterResource("aws-native:sagemaker:StudioLifecycleConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudioLifecycleConfig gets an existing StudioLifecycleConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudioLifecycleConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudioLifecycleConfigState, opts ...pulumi.ResourceOption) (*StudioLifecycleConfig, error) {
	var resource StudioLifecycleConfig
	err := ctx.ReadResource("aws-native:sagemaker:StudioLifecycleConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StudioLifecycleConfig resources.
type studioLifecycleConfigState struct {
}

type StudioLifecycleConfigState struct {
}

func (StudioLifecycleConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*studioLifecycleConfigState)(nil)).Elem()
}

type studioLifecycleConfigArgs struct {
	// The App type that the Lifecycle Configuration is attached to.
	StudioLifecycleConfigAppType StudioLifecycleConfigAppType `pulumi:"studioLifecycleConfigAppType"`
	// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent string `pulumi:"studioLifecycleConfigContent"`
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	StudioLifecycleConfigName *string `pulumi:"studioLifecycleConfigName"`
	// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
}

// The set of arguments for constructing a StudioLifecycleConfig resource.
type StudioLifecycleConfigArgs struct {
	// The App type that the Lifecycle Configuration is attached to.
	StudioLifecycleConfigAppType StudioLifecycleConfigAppTypeInput
	// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
	StudioLifecycleConfigContent pulumi.StringInput
	// The name of the Amazon SageMaker Studio Lifecycle Configuration.
	StudioLifecycleConfigName pulumi.StringPtrInput
	// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	Tags aws.CreateOnlyTagArrayInput
}

func (StudioLifecycleConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studioLifecycleConfigArgs)(nil)).Elem()
}

type StudioLifecycleConfigInput interface {
	pulumi.Input

	ToStudioLifecycleConfigOutput() StudioLifecycleConfigOutput
	ToStudioLifecycleConfigOutputWithContext(ctx context.Context) StudioLifecycleConfigOutput
}

func (*StudioLifecycleConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioLifecycleConfig)(nil)).Elem()
}

func (i *StudioLifecycleConfig) ToStudioLifecycleConfigOutput() StudioLifecycleConfigOutput {
	return i.ToStudioLifecycleConfigOutputWithContext(context.Background())
}

func (i *StudioLifecycleConfig) ToStudioLifecycleConfigOutputWithContext(ctx context.Context) StudioLifecycleConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioLifecycleConfigOutput)
}

type StudioLifecycleConfigOutput struct{ *pulumi.OutputState }

func (StudioLifecycleConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioLifecycleConfig)(nil)).Elem()
}

func (o StudioLifecycleConfigOutput) ToStudioLifecycleConfigOutput() StudioLifecycleConfigOutput {
	return o
}

func (o StudioLifecycleConfigOutput) ToStudioLifecycleConfigOutputWithContext(ctx context.Context) StudioLifecycleConfigOutput {
	return o
}

// The App type that the Lifecycle Configuration is attached to.
func (o StudioLifecycleConfigOutput) StudioLifecycleConfigAppType() StudioLifecycleConfigAppTypeOutput {
	return o.ApplyT(func(v *StudioLifecycleConfig) StudioLifecycleConfigAppTypeOutput {
		return v.StudioLifecycleConfigAppType
	}).(StudioLifecycleConfigAppTypeOutput)
}

// The Amazon Resource Name (ARN) of the Lifecycle Configuration.
func (o StudioLifecycleConfigOutput) StudioLifecycleConfigArn() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioLifecycleConfig) pulumi.StringOutput { return v.StudioLifecycleConfigArn }).(pulumi.StringOutput)
}

// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
func (o StudioLifecycleConfigOutput) StudioLifecycleConfigContent() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioLifecycleConfig) pulumi.StringOutput { return v.StudioLifecycleConfigContent }).(pulumi.StringOutput)
}

// The name of the Amazon SageMaker Studio Lifecycle Configuration.
func (o StudioLifecycleConfigOutput) StudioLifecycleConfigName() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioLifecycleConfig) pulumi.StringOutput { return v.StudioLifecycleConfigName }).(pulumi.StringOutput)
}

// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
func (o StudioLifecycleConfigOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *StudioLifecycleConfig) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioLifecycleConfigInput)(nil)).Elem(), &StudioLifecycleConfig{})
	pulumi.RegisterOutputType(StudioLifecycleConfigOutput{})
}
