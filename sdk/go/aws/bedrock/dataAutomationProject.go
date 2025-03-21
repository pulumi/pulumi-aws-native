// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bedrock

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Bedrock::DataAutomationProject Resource Type
type DataAutomationProject struct {
	pulumi.CustomResourceState

	// Time Stamp
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Blueprints to apply to objects processed by the project.
	CustomOutputConfiguration DataAutomationProjectCustomOutputConfigurationPtrOutput `pulumi:"customOutputConfiguration"`
	// KMS encryption context
	KmsEncryptionContext pulumi.StringMapOutput `pulumi:"kmsEncryptionContext"`
	// KMS key identifier
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Time Stamp
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// Additional settings for the project.
	OverrideConfiguration DataAutomationProjectOverrideConfigurationPtrOutput `pulumi:"overrideConfiguration"`
	// ARN of a DataAutomationProject
	ProjectArn pulumi.StringOutput `pulumi:"projectArn"`
	// Description of the DataAutomationProject
	ProjectDescription pulumi.StringPtrOutput `pulumi:"projectDescription"`
	// Name of the DataAutomationProject
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The project's stage.
	ProjectStage DataAutomationProjectStageOutput `pulumi:"projectStage"`
	// The project's standard output configuration.
	StandardOutputConfiguration DataAutomationProjectStandardOutputConfigurationPtrOutput `pulumi:"standardOutputConfiguration"`
	// The project's status.
	Status DataAutomationProjectStatusOutput `pulumi:"status"`
	// List of Tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewDataAutomationProject registers a new resource with the given unique name, arguments, and options.
func NewDataAutomationProject(ctx *pulumi.Context,
	name string, args *DataAutomationProjectArgs, opts ...pulumi.ResourceOption) (*DataAutomationProject, error) {
	if args == nil {
		args = &DataAutomationProjectArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"projectName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataAutomationProject
	err := ctx.RegisterResource("aws-native:bedrock:DataAutomationProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataAutomationProject gets an existing DataAutomationProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataAutomationProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataAutomationProjectState, opts ...pulumi.ResourceOption) (*DataAutomationProject, error) {
	var resource DataAutomationProject
	err := ctx.ReadResource("aws-native:bedrock:DataAutomationProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataAutomationProject resources.
type dataAutomationProjectState struct {
}

type DataAutomationProjectState struct {
}

func (DataAutomationProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAutomationProjectState)(nil)).Elem()
}

type dataAutomationProjectArgs struct {
	// Blueprints to apply to objects processed by the project.
	CustomOutputConfiguration *DataAutomationProjectCustomOutputConfiguration `pulumi:"customOutputConfiguration"`
	// KMS encryption context
	KmsEncryptionContext map[string]string `pulumi:"kmsEncryptionContext"`
	// KMS key identifier
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Additional settings for the project.
	OverrideConfiguration *DataAutomationProjectOverrideConfiguration `pulumi:"overrideConfiguration"`
	// Description of the DataAutomationProject
	ProjectDescription *string `pulumi:"projectDescription"`
	// Name of the DataAutomationProject
	ProjectName *string `pulumi:"projectName"`
	// The project's standard output configuration.
	StandardOutputConfiguration *DataAutomationProjectStandardOutputConfiguration `pulumi:"standardOutputConfiguration"`
	// List of Tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DataAutomationProject resource.
type DataAutomationProjectArgs struct {
	// Blueprints to apply to objects processed by the project.
	CustomOutputConfiguration DataAutomationProjectCustomOutputConfigurationPtrInput
	// KMS encryption context
	KmsEncryptionContext pulumi.StringMapInput
	// KMS key identifier
	KmsKeyId pulumi.StringPtrInput
	// Additional settings for the project.
	OverrideConfiguration DataAutomationProjectOverrideConfigurationPtrInput
	// Description of the DataAutomationProject
	ProjectDescription pulumi.StringPtrInput
	// Name of the DataAutomationProject
	ProjectName pulumi.StringPtrInput
	// The project's standard output configuration.
	StandardOutputConfiguration DataAutomationProjectStandardOutputConfigurationPtrInput
	// List of Tags
	Tags aws.TagArrayInput
}

func (DataAutomationProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataAutomationProjectArgs)(nil)).Elem()
}

type DataAutomationProjectInput interface {
	pulumi.Input

	ToDataAutomationProjectOutput() DataAutomationProjectOutput
	ToDataAutomationProjectOutputWithContext(ctx context.Context) DataAutomationProjectOutput
}

func (*DataAutomationProject) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAutomationProject)(nil)).Elem()
}

func (i *DataAutomationProject) ToDataAutomationProjectOutput() DataAutomationProjectOutput {
	return i.ToDataAutomationProjectOutputWithContext(context.Background())
}

func (i *DataAutomationProject) ToDataAutomationProjectOutputWithContext(ctx context.Context) DataAutomationProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataAutomationProjectOutput)
}

type DataAutomationProjectOutput struct{ *pulumi.OutputState }

func (DataAutomationProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataAutomationProject)(nil)).Elem()
}

func (o DataAutomationProjectOutput) ToDataAutomationProjectOutput() DataAutomationProjectOutput {
	return o
}

func (o DataAutomationProjectOutput) ToDataAutomationProjectOutputWithContext(ctx context.Context) DataAutomationProjectOutput {
	return o
}

// Time Stamp
func (o DataAutomationProjectOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Blueprints to apply to objects processed by the project.
func (o DataAutomationProjectOutput) CustomOutputConfiguration() DataAutomationProjectCustomOutputConfigurationPtrOutput {
	return o.ApplyT(func(v *DataAutomationProject) DataAutomationProjectCustomOutputConfigurationPtrOutput {
		return v.CustomOutputConfiguration
	}).(DataAutomationProjectCustomOutputConfigurationPtrOutput)
}

// KMS encryption context
func (o DataAutomationProjectOutput) KmsEncryptionContext() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringMapOutput { return v.KmsEncryptionContext }).(pulumi.StringMapOutput)
}

// KMS key identifier
func (o DataAutomationProjectOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Time Stamp
func (o DataAutomationProjectOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Additional settings for the project.
func (o DataAutomationProjectOutput) OverrideConfiguration() DataAutomationProjectOverrideConfigurationPtrOutput {
	return o.ApplyT(func(v *DataAutomationProject) DataAutomationProjectOverrideConfigurationPtrOutput {
		return v.OverrideConfiguration
	}).(DataAutomationProjectOverrideConfigurationPtrOutput)
}

// ARN of a DataAutomationProject
func (o DataAutomationProjectOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

// Description of the DataAutomationProject
func (o DataAutomationProjectOutput) ProjectDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringPtrOutput { return v.ProjectDescription }).(pulumi.StringPtrOutput)
}

// Name of the DataAutomationProject
func (o DataAutomationProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataAutomationProject) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The project's stage.
func (o DataAutomationProjectOutput) ProjectStage() DataAutomationProjectStageOutput {
	return o.ApplyT(func(v *DataAutomationProject) DataAutomationProjectStageOutput { return v.ProjectStage }).(DataAutomationProjectStageOutput)
}

// The project's standard output configuration.
func (o DataAutomationProjectOutput) StandardOutputConfiguration() DataAutomationProjectStandardOutputConfigurationPtrOutput {
	return o.ApplyT(func(v *DataAutomationProject) DataAutomationProjectStandardOutputConfigurationPtrOutput {
		return v.StandardOutputConfiguration
	}).(DataAutomationProjectStandardOutputConfigurationPtrOutput)
}

// The project's status.
func (o DataAutomationProjectOutput) Status() DataAutomationProjectStatusOutput {
	return o.ApplyT(func(v *DataAutomationProject) DataAutomationProjectStatusOutput { return v.Status }).(DataAutomationProjectStatusOutput)
}

// List of Tags
func (o DataAutomationProjectOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DataAutomationProject) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataAutomationProjectInput)(nil)).Elem(), &DataAutomationProject{})
	pulumi.RegisterOutputType(DataAutomationProjectOutput{})
}
