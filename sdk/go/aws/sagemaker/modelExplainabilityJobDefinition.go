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

// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
type ModelExplainabilityJobDefinition struct {
	pulumi.CustomResourceState

	// The time at which the job definition was created.
	CreationTime pulumi.StringOutput    `pulumi:"creationTime"`
	EndpointName pulumi.StringPtrOutput `pulumi:"endpointName"`
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn pulumi.StringOutput `pulumi:"jobDefinitionArn"`
	// The name of the model explainability job definition. The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName pulumi.StringPtrOutput `pulumi:"jobDefinitionName"`
	// Identifies the resources to deploy for a monitoring job.
	JobResources ModelExplainabilityJobDefinitionMonitoringResourcesOutput `pulumi:"jobResources"`
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput `pulumi:"modelExplainabilityAppSpecification"`
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput `pulumi:"modelExplainabilityBaselineConfig"`
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput `pulumi:"modelExplainabilityJobInput"`
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput `pulumi:"modelExplainabilityJobOutputConfig"`
	// Networking options for a model explainability job.
	NetworkConfig ModelExplainabilityJobDefinitionNetworkConfigPtrOutput `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition ModelExplainabilityJobDefinitionStoppingConditionPtrOutput `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayOutput `pulumi:"tags"`
}

// NewModelExplainabilityJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewModelExplainabilityJobDefinition(ctx *pulumi.Context,
	name string, args *ModelExplainabilityJobDefinitionArgs, opts ...pulumi.ResourceOption) (*ModelExplainabilityJobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobResources == nil {
		return nil, errors.New("invalid value for required argument 'JobResources'")
	}
	if args.ModelExplainabilityAppSpecification == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityAppSpecification'")
	}
	if args.ModelExplainabilityJobInput == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityJobInput'")
	}
	if args.ModelExplainabilityJobOutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'ModelExplainabilityJobOutputConfig'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"endpointName",
		"jobDefinitionName",
		"jobResources",
		"modelExplainabilityAppSpecification",
		"modelExplainabilityBaselineConfig",
		"modelExplainabilityJobInput",
		"modelExplainabilityJobOutputConfig",
		"networkConfig",
		"roleArn",
		"stoppingCondition",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ModelExplainabilityJobDefinition
	err := ctx.RegisterResource("aws-native:sagemaker:ModelExplainabilityJobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelExplainabilityJobDefinition gets an existing ModelExplainabilityJobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelExplainabilityJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelExplainabilityJobDefinitionState, opts ...pulumi.ResourceOption) (*ModelExplainabilityJobDefinition, error) {
	var resource ModelExplainabilityJobDefinition
	err := ctx.ReadResource("aws-native:sagemaker:ModelExplainabilityJobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelExplainabilityJobDefinition resources.
type modelExplainabilityJobDefinitionState struct {
}

type ModelExplainabilityJobDefinitionState struct {
}

func (ModelExplainabilityJobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelExplainabilityJobDefinitionState)(nil)).Elem()
}

type modelExplainabilityJobDefinitionArgs struct {
	EndpointName *string `pulumi:"endpointName"`
	// The name of the model explainability job definition. The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName *string `pulumi:"jobDefinitionName"`
	// Identifies the resources to deploy for a monitoring job.
	JobResources ModelExplainabilityJobDefinitionMonitoringResources `pulumi:"jobResources"`
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecification `pulumi:"modelExplainabilityAppSpecification"`
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig *ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfig `pulumi:"modelExplainabilityBaselineConfig"`
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput ModelExplainabilityJobDefinitionModelExplainabilityJobInput `pulumi:"modelExplainabilityJobInput"`
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig ModelExplainabilityJobDefinitionMonitoringOutputConfig `pulumi:"modelExplainabilityJobOutputConfig"`
	// Networking options for a model explainability job.
	NetworkConfig *ModelExplainabilityJobDefinitionNetworkConfig `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn string `pulumi:"roleArn"`
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition *ModelExplainabilityJobDefinitionStoppingCondition `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.CreateOnlyTag `pulumi:"tags"`
}

// The set of arguments for constructing a ModelExplainabilityJobDefinition resource.
type ModelExplainabilityJobDefinitionArgs struct {
	EndpointName pulumi.StringPtrInput
	// The name of the model explainability job definition. The name must be unique within an AWS Region in the AWS account.
	JobDefinitionName pulumi.StringPtrInput
	// Identifies the resources to deploy for a monitoring job.
	JobResources ModelExplainabilityJobDefinitionMonitoringResourcesInput
	// Configures the model explainability job to run a specified Docker container image.
	ModelExplainabilityAppSpecification ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationInput
	// The baseline configuration for a model explainability job.
	ModelExplainabilityBaselineConfig ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrInput
	// Inputs for the model explainability job.
	ModelExplainabilityJobInput ModelExplainabilityJobDefinitionModelExplainabilityJobInputInput
	// The output configuration for monitoring jobs.
	ModelExplainabilityJobOutputConfig ModelExplainabilityJobDefinitionMonitoringOutputConfigInput
	// Networking options for a model explainability job.
	NetworkConfig ModelExplainabilityJobDefinitionNetworkConfigPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn pulumi.StringInput
	// A time limit for how long the monitoring job is allowed to run before stopping.
	StoppingCondition ModelExplainabilityJobDefinitionStoppingConditionPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.CreateOnlyTagArrayInput
}

func (ModelExplainabilityJobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelExplainabilityJobDefinitionArgs)(nil)).Elem()
}

type ModelExplainabilityJobDefinitionInput interface {
	pulumi.Input

	ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput
	ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput
}

func (*ModelExplainabilityJobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelExplainabilityJobDefinition)(nil)).Elem()
}

func (i *ModelExplainabilityJobDefinition) ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput {
	return i.ToModelExplainabilityJobDefinitionOutputWithContext(context.Background())
}

func (i *ModelExplainabilityJobDefinition) ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelExplainabilityJobDefinitionOutput)
}

type ModelExplainabilityJobDefinitionOutput struct{ *pulumi.OutputState }

func (ModelExplainabilityJobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelExplainabilityJobDefinition)(nil)).Elem()
}

func (o ModelExplainabilityJobDefinitionOutput) ToModelExplainabilityJobDefinitionOutput() ModelExplainabilityJobDefinitionOutput {
	return o
}

func (o ModelExplainabilityJobDefinitionOutput) ToModelExplainabilityJobDefinitionOutputWithContext(ctx context.Context) ModelExplainabilityJobDefinitionOutput {
	return o
}

// The time at which the job definition was created.
func (o ModelExplainabilityJobDefinitionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ModelExplainabilityJobDefinitionOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringPtrOutput { return v.EndpointName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of job definition.
func (o ModelExplainabilityJobDefinitionOutput) JobDefinitionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.JobDefinitionArn }).(pulumi.StringOutput)
}

// The name of the model explainability job definition. The name must be unique within an AWS Region in the AWS account.
func (o ModelExplainabilityJobDefinitionOutput) JobDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringPtrOutput { return v.JobDefinitionName }).(pulumi.StringPtrOutput)
}

// Identifies the resources to deploy for a monitoring job.
func (o ModelExplainabilityJobDefinitionOutput) JobResources() ModelExplainabilityJobDefinitionMonitoringResourcesOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionMonitoringResourcesOutput {
		return v.JobResources
	}).(ModelExplainabilityJobDefinitionMonitoringResourcesOutput)
}

// Configures the model explainability job to run a specified Docker container image.
func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityAppSpecification() ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput {
		return v.ModelExplainabilityAppSpecification
	}).(ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationOutput)
}

// The baseline configuration for a model explainability job.
func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityBaselineConfig() ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput {
		return v.ModelExplainabilityBaselineConfig
	}).(ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigPtrOutput)
}

// Inputs for the model explainability job.
func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityJobInput() ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput {
		return v.ModelExplainabilityJobInput
	}).(ModelExplainabilityJobDefinitionModelExplainabilityJobInputOutput)
}

// The output configuration for monitoring jobs.
func (o ModelExplainabilityJobDefinitionOutput) ModelExplainabilityJobOutputConfig() ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput {
		return v.ModelExplainabilityJobOutputConfig
	}).(ModelExplainabilityJobDefinitionMonitoringOutputConfigOutput)
}

// Networking options for a model explainability job.
func (o ModelExplainabilityJobDefinitionOutput) NetworkConfig() ModelExplainabilityJobDefinitionNetworkConfigPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionNetworkConfigPtrOutput {
		return v.NetworkConfig
	}).(ModelExplainabilityJobDefinitionNetworkConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
func (o ModelExplainabilityJobDefinitionOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// A time limit for how long the monitoring job is allowed to run before stopping.
func (o ModelExplainabilityJobDefinitionOutput) StoppingCondition() ModelExplainabilityJobDefinitionStoppingConditionPtrOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) ModelExplainabilityJobDefinitionStoppingConditionPtrOutput {
		return v.StoppingCondition
	}).(ModelExplainabilityJobDefinitionStoppingConditionPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ModelExplainabilityJobDefinitionOutput) Tags() aws.CreateOnlyTagArrayOutput {
	return o.ApplyT(func(v *ModelExplainabilityJobDefinition) aws.CreateOnlyTagArrayOutput { return v.Tags }).(aws.CreateOnlyTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelExplainabilityJobDefinitionInput)(nil)).Elem(), &ModelExplainabilityJobDefinition{})
	pulumi.RegisterOutputType(ModelExplainabilityJobDefinitionOutput{})
}
