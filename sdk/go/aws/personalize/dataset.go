// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package personalize

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Personalize::Dataset.
type Dataset struct {
	pulumi.CustomResourceState

	// The ARN of the dataset
	DatasetArn pulumi.StringOutput `pulumi:"datasetArn"`
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to
	DatasetGroupArn  pulumi.StringOutput       `pulumi:"datasetGroupArn"`
	DatasetImportJob DatasetImportJobPtrOutput `pulumi:"datasetImportJob"`
	// The type of dataset
	DatasetType DatasetTypeOutput `pulumi:"datasetType"`
	// The name for the dataset
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
	SchemaArn pulumi.StringOutput `pulumi:"schemaArn"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetGroupArn == nil {
		return nil, errors.New("invalid value for required argument 'DatasetGroupArn'")
	}
	if args.DatasetType == nil {
		return nil, errors.New("invalid value for required argument 'DatasetType'")
	}
	if args.SchemaArn == nil {
		return nil, errors.New("invalid value for required argument 'SchemaArn'")
	}
	var resource Dataset
	err := ctx.RegisterResource("aws-native:personalize:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("aws-native:personalize:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to
	DatasetGroupArn  string            `pulumi:"datasetGroupArn"`
	DatasetImportJob *DatasetImportJob `pulumi:"datasetImportJob"`
	// The type of dataset
	DatasetType DatasetType `pulumi:"datasetType"`
	// The name for the dataset
	Name *string `pulumi:"name"`
	// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
	SchemaArn string `pulumi:"schemaArn"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to
	DatasetGroupArn  pulumi.StringInput
	DatasetImportJob DatasetImportJobPtrInput
	// The type of dataset
	DatasetType DatasetTypeInput
	// The name for the dataset
	Name pulumi.StringPtrInput
	// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
	SchemaArn pulumi.StringInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

// The ARN of the dataset
func (o DatasetOutput) DatasetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.DatasetArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the dataset group to add the dataset to
func (o DatasetOutput) DatasetGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.DatasetGroupArn }).(pulumi.StringOutput)
}

func (o DatasetOutput) DatasetImportJob() DatasetImportJobPtrOutput {
	return o.ApplyT(func(v *Dataset) DatasetImportJobPtrOutput { return v.DatasetImportJob }).(DatasetImportJobPtrOutput)
}

// The type of dataset
func (o DatasetOutput) DatasetType() DatasetTypeOutput {
	return o.ApplyT(func(v *Dataset) DatasetTypeOutput { return v.DatasetType }).(DatasetTypeOutput)
}

// The name for the dataset
func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the schema to associate with the dataset. The schema defines the dataset fields.
func (o DatasetOutput) SchemaArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.SchemaArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetInput)(nil)).Elem(), &Dataset{})
	pulumi.RegisterOutputType(DatasetOutput{})
}