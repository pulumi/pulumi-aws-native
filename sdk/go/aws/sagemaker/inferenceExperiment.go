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

// Resource Type definition for AWS::SageMaker::InferenceExperiment
type InferenceExperiment struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the inference experiment.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The timestamp at which you created the inference experiment.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The Amazon S3 location and configuration for storing inference request and response data.
	DataStorageConfig InferenceExperimentDataStorageConfigPtrOutput `pulumi:"dataStorageConfig"`
	// The description of the inference experiment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The desired state of the experiment after starting or stopping operation.
	DesiredState     InferenceExperimentDesiredStatePtrOutput  `pulumi:"desiredState"`
	EndpointMetadata InferenceExperimentEndpointMetadataOutput `pulumi:"endpointMetadata"`
	// The name of the endpoint.
	EndpointName pulumi.StringOutput `pulumi:"endpointName"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// The timestamp at which you last modified the inference experiment.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
	ModelVariants InferenceExperimentModelVariantConfigArrayOutput `pulumi:"modelVariants"`
	// The name for the inference experiment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The duration for which the inference experiment ran or will run.
	//
	// The maximum duration that you can set for an inference experiment is 30 days.
	Schedule InferenceExperimentSchedulePtrOutput `pulumi:"schedule"`
	// The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
	ShadowModeConfig InferenceExperimentShadowModeConfigPtrOutput `pulumi:"shadowModeConfig"`
	// The status of the inference experiment.
	Status InferenceExperimentStatusOutput `pulumi:"status"`
	// The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
	StatusReason pulumi.StringPtrOutput `pulumi:"statusReason"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The type of the inference experiment that you want to run.
	Type InferenceExperimentTypeOutput `pulumi:"type"`
}

// NewInferenceExperiment registers a new resource with the given unique name, arguments, and options.
func NewInferenceExperiment(ctx *pulumi.Context,
	name string, args *InferenceExperimentArgs, opts ...pulumi.ResourceOption) (*InferenceExperiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.ModelVariants == nil {
		return nil, errors.New("invalid value for required argument 'ModelVariants'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"endpointName",
		"kmsKey",
		"name",
		"roleArn",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InferenceExperiment
	err := ctx.RegisterResource("aws-native:sagemaker:InferenceExperiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInferenceExperiment gets an existing InferenceExperiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInferenceExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InferenceExperimentState, opts ...pulumi.ResourceOption) (*InferenceExperiment, error) {
	var resource InferenceExperiment
	err := ctx.ReadResource("aws-native:sagemaker:InferenceExperiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InferenceExperiment resources.
type inferenceExperimentState struct {
}

type InferenceExperimentState struct {
}

func (InferenceExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceExperimentState)(nil)).Elem()
}

type inferenceExperimentArgs struct {
	// The Amazon S3 location and configuration for storing inference request and response data.
	DataStorageConfig *InferenceExperimentDataStorageConfig `pulumi:"dataStorageConfig"`
	// The description of the inference experiment.
	Description *string `pulumi:"description"`
	// The desired state of the experiment after starting or stopping operation.
	DesiredState *InferenceExperimentDesiredState `pulumi:"desiredState"`
	// The name of the endpoint.
	EndpointName string `pulumi:"endpointName"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKey *string `pulumi:"kmsKey"`
	// An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
	ModelVariants []InferenceExperimentModelVariantConfig `pulumi:"modelVariants"`
	// The name for the inference experiment.
	Name *string `pulumi:"name"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
	RoleArn string `pulumi:"roleArn"`
	// The duration for which the inference experiment ran or will run.
	//
	// The maximum duration that you can set for an inference experiment is 30 days.
	Schedule *InferenceExperimentSchedule `pulumi:"schedule"`
	// The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
	ShadowModeConfig *InferenceExperimentShadowModeConfig `pulumi:"shadowModeConfig"`
	// The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
	StatusReason *string `pulumi:"statusReason"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The type of the inference experiment that you want to run.
	Type InferenceExperimentType `pulumi:"type"`
}

// The set of arguments for constructing a InferenceExperiment resource.
type InferenceExperimentArgs struct {
	// The Amazon S3 location and configuration for storing inference request and response data.
	DataStorageConfig InferenceExperimentDataStorageConfigPtrInput
	// The description of the inference experiment.
	Description pulumi.StringPtrInput
	// The desired state of the experiment after starting or stopping operation.
	DesiredState InferenceExperimentDesiredStatePtrInput
	// The name of the endpoint.
	EndpointName pulumi.StringInput
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	KmsKey pulumi.StringPtrInput
	// An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
	ModelVariants InferenceExperimentModelVariantConfigArrayInput
	// The name for the inference experiment.
	Name pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
	RoleArn pulumi.StringInput
	// The duration for which the inference experiment ran or will run.
	//
	// The maximum duration that you can set for an inference experiment is 30 days.
	Schedule InferenceExperimentSchedulePtrInput
	// The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
	ShadowModeConfig InferenceExperimentShadowModeConfigPtrInput
	// The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
	StatusReason pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
	// The type of the inference experiment that you want to run.
	Type InferenceExperimentTypeInput
}

func (InferenceExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceExperimentArgs)(nil)).Elem()
}

type InferenceExperimentInput interface {
	pulumi.Input

	ToInferenceExperimentOutput() InferenceExperimentOutput
	ToInferenceExperimentOutputWithContext(ctx context.Context) InferenceExperimentOutput
}

func (*InferenceExperiment) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceExperiment)(nil)).Elem()
}

func (i *InferenceExperiment) ToInferenceExperimentOutput() InferenceExperimentOutput {
	return i.ToInferenceExperimentOutputWithContext(context.Background())
}

func (i *InferenceExperiment) ToInferenceExperimentOutputWithContext(ctx context.Context) InferenceExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceExperimentOutput)
}

type InferenceExperimentOutput struct{ *pulumi.OutputState }

func (InferenceExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceExperiment)(nil)).Elem()
}

func (o InferenceExperimentOutput) ToInferenceExperimentOutput() InferenceExperimentOutput {
	return o
}

func (o InferenceExperimentOutput) ToInferenceExperimentOutputWithContext(ctx context.Context) InferenceExperimentOutput {
	return o
}

// The Amazon Resource Name (ARN) of the inference experiment.
func (o InferenceExperimentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The timestamp at which you created the inference experiment.
func (o InferenceExperimentOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The Amazon S3 location and configuration for storing inference request and response data.
func (o InferenceExperimentOutput) DataStorageConfig() InferenceExperimentDataStorageConfigPtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentDataStorageConfigPtrOutput { return v.DataStorageConfig }).(InferenceExperimentDataStorageConfigPtrOutput)
}

// The description of the inference experiment.
func (o InferenceExperimentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The desired state of the experiment after starting or stopping operation.
func (o InferenceExperimentOutput) DesiredState() InferenceExperimentDesiredStatePtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentDesiredStatePtrOutput { return v.DesiredState }).(InferenceExperimentDesiredStatePtrOutput)
}

func (o InferenceExperimentOutput) EndpointMetadata() InferenceExperimentEndpointMetadataOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentEndpointMetadataOutput { return v.EndpointMetadata }).(InferenceExperimentEndpointMetadataOutput)
}

// The name of the endpoint.
func (o InferenceExperimentOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.EndpointName }).(pulumi.StringOutput)
}

// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
func (o InferenceExperimentOutput) KmsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringPtrOutput { return v.KmsKey }).(pulumi.StringPtrOutput)
}

// The timestamp at which you last modified the inference experiment.
func (o InferenceExperimentOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.
func (o InferenceExperimentOutput) ModelVariants() InferenceExperimentModelVariantConfigArrayOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentModelVariantConfigArrayOutput { return v.ModelVariants }).(InferenceExperimentModelVariantConfigArrayOutput)
}

// The name for the inference experiment.
func (o InferenceExperimentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.
func (o InferenceExperimentOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The duration for which the inference experiment ran or will run.
//
// The maximum duration that you can set for an inference experiment is 30 days.
func (o InferenceExperimentOutput) Schedule() InferenceExperimentSchedulePtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentSchedulePtrOutput { return v.Schedule }).(InferenceExperimentSchedulePtrOutput)
}

// The configuration of `ShadowMode` inference experiment type, which shows the production variant that takes all the inference requests, and the shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant it also shows the percentage of requests that Amazon SageMaker replicates.
func (o InferenceExperimentOutput) ShadowModeConfig() InferenceExperimentShadowModeConfigPtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentShadowModeConfigPtrOutput { return v.ShadowModeConfig }).(InferenceExperimentShadowModeConfigPtrOutput)
}

// The status of the inference experiment.
func (o InferenceExperimentOutput) Status() InferenceExperimentStatusOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentStatusOutput { return v.Status }).(InferenceExperimentStatusOutput)
}

// The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.
func (o InferenceExperimentOutput) StatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceExperiment) pulumi.StringPtrOutput { return v.StatusReason }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o InferenceExperimentOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *InferenceExperiment) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The type of the inference experiment that you want to run.
func (o InferenceExperimentOutput) Type() InferenceExperimentTypeOutput {
	return o.ApplyT(func(v *InferenceExperiment) InferenceExperimentTypeOutput { return v.Type }).(InferenceExperimentTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InferenceExperimentInput)(nil)).Elem(), &InferenceExperiment{})
	pulumi.RegisterOutputType(InferenceExperimentOutput{})
}
