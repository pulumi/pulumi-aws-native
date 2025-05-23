// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelPackage
type ModelPackage struct {
	pulumi.CustomResourceState

	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput `pulumi:"additionalInferenceSpecifications"`
	// An array of additional Inference Specification objects to be added to the existing array. The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
	AdditionalInferenceSpecificationsToAdd ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput `pulumi:"additionalInferenceSpecificationsToAdd"`
	// A description provided when the model approval is set.
	ApprovalDescription pulumi.StringPtrOutput `pulumi:"approvalDescription"`
	// Whether the model package is to be certified to be listed on AWS Marketplace. For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	CertifyForMarketplace pulumi.BoolPtrOutput `pulumi:"certifyForMarketplace"`
	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken pulumi.StringPtrOutput `pulumi:"clientToken"`
	// The time that the model package was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The metadata properties for the model package.
	CustomerMetadataProperties ModelPackageCustomerMetadataPropertiesPtrOutput `pulumi:"customerMetadataProperties"`
	// The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	DriftCheckBaselines ModelPackageDriftCheckBaselinesPtrOutput `pulumi:"driftCheckBaselines"`
	// Defines how to perform inference generation after a training job is run.
	InferenceSpecification ModelPackageInferenceSpecificationPtrOutput `pulumi:"inferenceSpecification"`
	// The last time the model package was modified.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties ModelPackageMetadataPropertiesPtrOutput `pulumi:"metadataProperties"`
	// The approval status of the model. This can be one of the following values.
	//
	// - `APPROVED` - The model is approved
	// - `REJECTED` - The model is rejected.
	// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
	ModelApprovalStatus ModelPackageModelApprovalStatusPtrOutput `pulumi:"modelApprovalStatus"`
	// An Amazon SageMaker Model Card.
	ModelCard ModelPackageModelCardPtrOutput `pulumi:"modelCard"`
	// Metrics for the model.
	ModelMetrics ModelPackageModelMetricsPtrOutput `pulumi:"modelMetrics"`
	// The Amazon Resource Name (ARN) of the model package.
	ModelPackageArn pulumi.StringOutput `pulumi:"modelPackageArn"`
	// The description of the model package.
	ModelPackageDescription pulumi.StringPtrOutput `pulumi:"modelPackageDescription"`
	// The model group to which the model belongs.
	ModelPackageGroupName pulumi.StringPtrOutput `pulumi:"modelPackageGroupName"`
	// The name of the model package. The name can be as follows:
	//
	// - For a versioned model, the name is automatically generated by SageMaker Model Registry and follows the format ' `ModelPackageGroupName/ModelPackageVersion` '.
	// - For an unversioned model, you must provide the name.
	ModelPackageName pulumi.StringPtrOutput `pulumi:"modelPackageName"`
	// The status of the model package. This can be one of the following values.
	//
	// - `PENDING` - The model package creation is pending.
	// - `IN_PROGRESS` - The model package is in the process of being created.
	// - `COMPLETED` - The model package was successfully created.
	// - `FAILED` - The model package creation failed.
	// - `DELETING` - The model package is in the process of being deleted.
	ModelPackageStatus ModelPackageStatusOutput `pulumi:"modelPackageStatus"`
	// Specifies the validation and image scan statuses of the model package.
	ModelPackageStatusDetails ModelPackageStatusDetailsPtrOutput `pulumi:"modelPackageStatusDetails"`
	// The version number of a versioned model.
	ModelPackageVersion pulumi.IntPtrOutput `pulumi:"modelPackageVersion"`
	// The Amazon Simple Storage Service path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	SamplePayloadUrl pulumi.StringPtrOutput              `pulumi:"samplePayloadUrl"`
	SecurityConfig   ModelPackageSecurityConfigPtrOutput `pulumi:"securityConfig"`
	// Indicates if you want to skip model validation.
	SkipModelValidation ModelPackageSkipModelValidationPtrOutput `pulumi:"skipModelValidation"`
	// A list of algorithms that were used to create a model package.
	SourceAlgorithmSpecification ModelPackageSourceAlgorithmSpecificationPtrOutput `pulumi:"sourceAlgorithmSpecification"`
	// The URI of the source for the model package.
	SourceUri pulumi.StringPtrOutput `pulumi:"sourceUri"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification.
	Task pulumi.StringPtrOutput `pulumi:"task"`
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	ValidationSpecification ModelPackageValidationSpecificationPtrOutput `pulumi:"validationSpecification"`
}

// NewModelPackage registers a new resource with the given unique name, arguments, and options.
func NewModelPackage(ctx *pulumi.Context,
	name string, args *ModelPackageArgs, opts ...pulumi.ResourceOption) (*ModelPackage, error) {
	if args == nil {
		args = &ModelPackageArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clientToken",
		"domain",
		"driftCheckBaselines",
		"inferenceSpecification",
		"metadataProperties",
		"modelMetrics",
		"modelPackageDescription",
		"modelPackageGroupName",
		"samplePayloadUrl",
		"securityConfig",
		"sourceAlgorithmSpecification",
		"task",
		"validationSpecification",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ModelPackage
	err := ctx.RegisterResource("aws-native:sagemaker:ModelPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelPackage gets an existing ModelPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelPackageState, opts ...pulumi.ResourceOption) (*ModelPackage, error) {
	var resource ModelPackage
	err := ctx.ReadResource("aws-native:sagemaker:ModelPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelPackage resources.
type modelPackageState struct {
}

type ModelPackageState struct {
}

func (ModelPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelPackageState)(nil)).Elem()
}

type modelPackageArgs struct {
	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications []ModelPackageAdditionalInferenceSpecificationDefinition `pulumi:"additionalInferenceSpecifications"`
	// An array of additional Inference Specification objects to be added to the existing array. The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
	AdditionalInferenceSpecificationsToAdd []ModelPackageAdditionalInferenceSpecificationDefinition `pulumi:"additionalInferenceSpecificationsToAdd"`
	// A description provided when the model approval is set.
	ApprovalDescription *string `pulumi:"approvalDescription"`
	// Whether the model package is to be certified to be listed on AWS Marketplace. For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	CertifyForMarketplace *bool `pulumi:"certifyForMarketplace"`
	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string `pulumi:"clientToken"`
	// The metadata properties for the model package.
	CustomerMetadataProperties *ModelPackageCustomerMetadataProperties `pulumi:"customerMetadataProperties"`
	// The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.
	Domain *string `pulumi:"domain"`
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	DriftCheckBaselines *ModelPackageDriftCheckBaselines `pulumi:"driftCheckBaselines"`
	// Defines how to perform inference generation after a training job is run.
	InferenceSpecification *ModelPackageInferenceSpecification `pulumi:"inferenceSpecification"`
	// The last time the model package was modified.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *ModelPackageMetadataProperties `pulumi:"metadataProperties"`
	// The approval status of the model. This can be one of the following values.
	//
	// - `APPROVED` - The model is approved
	// - `REJECTED` - The model is rejected.
	// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
	ModelApprovalStatus *ModelPackageModelApprovalStatus `pulumi:"modelApprovalStatus"`
	// An Amazon SageMaker Model Card.
	ModelCard *ModelPackageModelCard `pulumi:"modelCard"`
	// Metrics for the model.
	ModelMetrics *ModelPackageModelMetrics `pulumi:"modelMetrics"`
	// The description of the model package.
	ModelPackageDescription *string `pulumi:"modelPackageDescription"`
	// The model group to which the model belongs.
	ModelPackageGroupName *string `pulumi:"modelPackageGroupName"`
	// The name of the model package. The name can be as follows:
	//
	// - For a versioned model, the name is automatically generated by SageMaker Model Registry and follows the format ' `ModelPackageGroupName/ModelPackageVersion` '.
	// - For an unversioned model, you must provide the name.
	ModelPackageName *string `pulumi:"modelPackageName"`
	// Specifies the validation and image scan statuses of the model package.
	ModelPackageStatusDetails *ModelPackageStatusDetails `pulumi:"modelPackageStatusDetails"`
	// The version number of a versioned model.
	ModelPackageVersion *int `pulumi:"modelPackageVersion"`
	// The Amazon Simple Storage Service path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	SamplePayloadUrl *string                     `pulumi:"samplePayloadUrl"`
	SecurityConfig   *ModelPackageSecurityConfig `pulumi:"securityConfig"`
	// Indicates if you want to skip model validation.
	SkipModelValidation *ModelPackageSkipModelValidation `pulumi:"skipModelValidation"`
	// A list of algorithms that were used to create a model package.
	SourceAlgorithmSpecification *ModelPackageSourceAlgorithmSpecification `pulumi:"sourceAlgorithmSpecification"`
	// The URI of the source for the model package.
	SourceUri *string `pulumi:"sourceUri"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification.
	Task *string `pulumi:"task"`
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	ValidationSpecification *ModelPackageValidationSpecification `pulumi:"validationSpecification"`
}

// The set of arguments for constructing a ModelPackage resource.
type ModelPackageArgs struct {
	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications ModelPackageAdditionalInferenceSpecificationDefinitionArrayInput
	// An array of additional Inference Specification objects to be added to the existing array. The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
	AdditionalInferenceSpecificationsToAdd ModelPackageAdditionalInferenceSpecificationDefinitionArrayInput
	// A description provided when the model approval is set.
	ApprovalDescription pulumi.StringPtrInput
	// Whether the model package is to be certified to be listed on AWS Marketplace. For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	CertifyForMarketplace pulumi.BoolPtrInput
	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken pulumi.StringPtrInput
	// The metadata properties for the model package.
	CustomerMetadataProperties ModelPackageCustomerMetadataPropertiesPtrInput
	// The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.
	Domain pulumi.StringPtrInput
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	DriftCheckBaselines ModelPackageDriftCheckBaselinesPtrInput
	// Defines how to perform inference generation after a training job is run.
	InferenceSpecification ModelPackageInferenceSpecificationPtrInput
	// The last time the model package was modified.
	LastModifiedTime pulumi.StringPtrInput
	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties ModelPackageMetadataPropertiesPtrInput
	// The approval status of the model. This can be one of the following values.
	//
	// - `APPROVED` - The model is approved
	// - `REJECTED` - The model is rejected.
	// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
	ModelApprovalStatus ModelPackageModelApprovalStatusPtrInput
	// An Amazon SageMaker Model Card.
	ModelCard ModelPackageModelCardPtrInput
	// Metrics for the model.
	ModelMetrics ModelPackageModelMetricsPtrInput
	// The description of the model package.
	ModelPackageDescription pulumi.StringPtrInput
	// The model group to which the model belongs.
	ModelPackageGroupName pulumi.StringPtrInput
	// The name of the model package. The name can be as follows:
	//
	// - For a versioned model, the name is automatically generated by SageMaker Model Registry and follows the format ' `ModelPackageGroupName/ModelPackageVersion` '.
	// - For an unversioned model, you must provide the name.
	ModelPackageName pulumi.StringPtrInput
	// Specifies the validation and image scan statuses of the model package.
	ModelPackageStatusDetails ModelPackageStatusDetailsPtrInput
	// The version number of a versioned model.
	ModelPackageVersion pulumi.IntPtrInput
	// The Amazon Simple Storage Service path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	SamplePayloadUrl pulumi.StringPtrInput
	SecurityConfig   ModelPackageSecurityConfigPtrInput
	// Indicates if you want to skip model validation.
	SkipModelValidation ModelPackageSkipModelValidationPtrInput
	// A list of algorithms that were used to create a model package.
	SourceAlgorithmSpecification ModelPackageSourceAlgorithmSpecificationPtrInput
	// The URI of the source for the model package.
	SourceUri pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
	// The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification.
	Task pulumi.StringPtrInput
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	ValidationSpecification ModelPackageValidationSpecificationPtrInput
}

func (ModelPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelPackageArgs)(nil)).Elem()
}

type ModelPackageInput interface {
	pulumi.Input

	ToModelPackageOutput() ModelPackageOutput
	ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput
}

func (*ModelPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPackage)(nil)).Elem()
}

func (i *ModelPackage) ToModelPackageOutput() ModelPackageOutput {
	return i.ToModelPackageOutputWithContext(context.Background())
}

func (i *ModelPackage) ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPackageOutput)
}

type ModelPackageOutput struct{ *pulumi.OutputState }

func (ModelPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPackage)(nil)).Elem()
}

func (o ModelPackageOutput) ToModelPackageOutput() ModelPackageOutput {
	return o
}

func (o ModelPackageOutput) ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput {
	return o
}

// An array of additional Inference Specification objects.
func (o ModelPackageOutput) AdditionalInferenceSpecifications() ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
		return v.AdditionalInferenceSpecifications
	}).(ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput)
}

// An array of additional Inference Specification objects to be added to the existing array. The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
func (o ModelPackageOutput) AdditionalInferenceSpecificationsToAdd() ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
		return v.AdditionalInferenceSpecificationsToAdd
	}).(ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput)
}

// A description provided when the model approval is set.
func (o ModelPackageOutput) ApprovalDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ApprovalDescription }).(pulumi.StringPtrOutput)
}

// Whether the model package is to be certified to be listed on AWS Marketplace. For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
func (o ModelPackageOutput) CertifyForMarketplace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.BoolPtrOutput { return v.CertifyForMarketplace }).(pulumi.BoolPtrOutput)
}

// A unique token that guarantees that the call to this API is idempotent.
func (o ModelPackageOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

// The time that the model package was created.
func (o ModelPackageOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The metadata properties for the model package.
func (o ModelPackageOutput) CustomerMetadataProperties() ModelPackageCustomerMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageCustomerMetadataPropertiesPtrOutput {
		return v.CustomerMetadataProperties
	}).(ModelPackageCustomerMetadataPropertiesPtrOutput)
}

// The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.
func (o ModelPackageOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// Represents the drift check baselines that can be used when the model monitor is set using the model package.
func (o ModelPackageOutput) DriftCheckBaselines() ModelPackageDriftCheckBaselinesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageDriftCheckBaselinesPtrOutput { return v.DriftCheckBaselines }).(ModelPackageDriftCheckBaselinesPtrOutput)
}

// Defines how to perform inference generation after a training job is run.
func (o ModelPackageOutput) InferenceSpecification() ModelPackageInferenceSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageInferenceSpecificationPtrOutput { return v.InferenceSpecification }).(ModelPackageInferenceSpecificationPtrOutput)
}

// The last time the model package was modified.
func (o ModelPackageOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Metadata properties of the tracking entity, trial, or trial component.
func (o ModelPackageOutput) MetadataProperties() ModelPackageMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageMetadataPropertiesPtrOutput { return v.MetadataProperties }).(ModelPackageMetadataPropertiesPtrOutput)
}

// The approval status of the model. This can be one of the following values.
//
// - `APPROVED` - The model is approved
// - `REJECTED` - The model is rejected.
// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
func (o ModelPackageOutput) ModelApprovalStatus() ModelPackageModelApprovalStatusPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageModelApprovalStatusPtrOutput { return v.ModelApprovalStatus }).(ModelPackageModelApprovalStatusPtrOutput)
}

// An Amazon SageMaker Model Card.
func (o ModelPackageOutput) ModelCard() ModelPackageModelCardPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageModelCardPtrOutput { return v.ModelCard }).(ModelPackageModelCardPtrOutput)
}

// Metrics for the model.
func (o ModelPackageOutput) ModelMetrics() ModelPackageModelMetricsPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageModelMetricsPtrOutput { return v.ModelMetrics }).(ModelPackageModelMetricsPtrOutput)
}

// The Amazon Resource Name (ARN) of the model package.
func (o ModelPackageOutput) ModelPackageArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringOutput { return v.ModelPackageArn }).(pulumi.StringOutput)
}

// The description of the model package.
func (o ModelPackageOutput) ModelPackageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageDescription }).(pulumi.StringPtrOutput)
}

// The model group to which the model belongs.
func (o ModelPackageOutput) ModelPackageGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageGroupName }).(pulumi.StringPtrOutput)
}

// The name of the model package. The name can be as follows:
//
// - For a versioned model, the name is automatically generated by SageMaker Model Registry and follows the format ' `ModelPackageGroupName/ModelPackageVersion` '.
// - For an unversioned model, you must provide the name.
func (o ModelPackageOutput) ModelPackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageName }).(pulumi.StringPtrOutput)
}

// The status of the model package. This can be one of the following values.
//
// - `PENDING` - The model package creation is pending.
// - `IN_PROGRESS` - The model package is in the process of being created.
// - `COMPLETED` - The model package was successfully created.
// - `FAILED` - The model package creation failed.
// - `DELETING` - The model package is in the process of being deleted.
func (o ModelPackageOutput) ModelPackageStatus() ModelPackageStatusOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageStatusOutput { return v.ModelPackageStatus }).(ModelPackageStatusOutput)
}

// Specifies the validation and image scan statuses of the model package.
func (o ModelPackageOutput) ModelPackageStatusDetails() ModelPackageStatusDetailsPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageStatusDetailsPtrOutput { return v.ModelPackageStatusDetails }).(ModelPackageStatusDetailsPtrOutput)
}

// The version number of a versioned model.
func (o ModelPackageOutput) ModelPackageVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.IntPtrOutput { return v.ModelPackageVersion }).(pulumi.IntPtrOutput)
}

// The Amazon Simple Storage Service path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).
func (o ModelPackageOutput) SamplePayloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.SamplePayloadUrl }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) SecurityConfig() ModelPackageSecurityConfigPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageSecurityConfigPtrOutput { return v.SecurityConfig }).(ModelPackageSecurityConfigPtrOutput)
}

// Indicates if you want to skip model validation.
func (o ModelPackageOutput) SkipModelValidation() ModelPackageSkipModelValidationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageSkipModelValidationPtrOutput { return v.SkipModelValidation }).(ModelPackageSkipModelValidationPtrOutput)
}

// A list of algorithms that were used to create a model package.
func (o ModelPackageOutput) SourceAlgorithmSpecification() ModelPackageSourceAlgorithmSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageSourceAlgorithmSpecificationPtrOutput {
		return v.SourceAlgorithmSpecification
	}).(ModelPackageSourceAlgorithmSpecificationPtrOutput)
}

// The URI of the source for the model package.
func (o ModelPackageOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.SourceUri }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ModelPackageOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ModelPackage) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification.
func (o ModelPackageOutput) Task() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.Task }).(pulumi.StringPtrOutput)
}

// Specifies batch transform jobs that SageMaker runs to validate your model package.
func (o ModelPackageOutput) ValidationSpecification() ModelPackageValidationSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageValidationSpecificationPtrOutput { return v.ValidationSpecification }).(ModelPackageValidationSpecificationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelPackageInput)(nil)).Elem(), &ModelPackage{})
	pulumi.RegisterOutputType(ModelPackageOutput{})
}
