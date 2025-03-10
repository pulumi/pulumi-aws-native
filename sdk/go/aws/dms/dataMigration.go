// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DMS::DataMigration.
type DataMigration struct {
	pulumi.CustomResourceState

	// The property describes an ARN of the data migration.
	DataMigrationArn pulumi.StringOutput `pulumi:"dataMigrationArn"`
	// The property describes the create time of the data migration.
	DataMigrationCreateTime pulumi.StringOutput `pulumi:"dataMigrationCreateTime"`
	// The property describes an ARN of the data migration.
	DataMigrationIdentifier pulumi.StringPtrOutput `pulumi:"dataMigrationIdentifier"`
	// The property describes a name to identify the data migration.
	DataMigrationName pulumi.StringPtrOutput `pulumi:"dataMigrationName"`
	// The property describes the settings for the data migration.
	DataMigrationSettings DataMigrationSettingsPtrOutput `pulumi:"dataMigrationSettings"`
	// The property describes the type of migration.
	DataMigrationType DataMigrationTypeOutput `pulumi:"dataMigrationType"`
	// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
	MigrationProjectIdentifier pulumi.StringOutput `pulumi:"migrationProjectIdentifier"`
	// The property describes Amazon Resource Name (ARN) of the service access role.
	ServiceAccessRoleArn pulumi.StringOutput `pulumi:"serviceAccessRoleArn"`
	// The property describes the settings for the data migration.
	SourceDataSettings DataMigrationSourceDataSettingsArrayOutput `pulumi:"sourceDataSettings"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewDataMigration registers a new resource with the given unique name, arguments, and options.
func NewDataMigration(ctx *pulumi.Context,
	name string, args *DataMigrationArgs, opts ...pulumi.ResourceOption) (*DataMigration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataMigrationType == nil {
		return nil, errors.New("invalid value for required argument 'DataMigrationType'")
	}
	if args.MigrationProjectIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'MigrationProjectIdentifier'")
	}
	if args.ServiceAccessRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccessRoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataMigration
	err := ctx.RegisterResource("aws-native:dms:DataMigration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataMigration gets an existing DataMigration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataMigrationState, opts ...pulumi.ResourceOption) (*DataMigration, error) {
	var resource DataMigration
	err := ctx.ReadResource("aws-native:dms:DataMigration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataMigration resources.
type dataMigrationState struct {
}

type DataMigrationState struct {
}

func (DataMigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMigrationState)(nil)).Elem()
}

type dataMigrationArgs struct {
	// The property describes an ARN of the data migration.
	DataMigrationIdentifier *string `pulumi:"dataMigrationIdentifier"`
	// The property describes a name to identify the data migration.
	DataMigrationName *string `pulumi:"dataMigrationName"`
	// The property describes the settings for the data migration.
	DataMigrationSettings *DataMigrationSettings `pulumi:"dataMigrationSettings"`
	// The property describes the type of migration.
	DataMigrationType DataMigrationType `pulumi:"dataMigrationType"`
	// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
	MigrationProjectIdentifier string `pulumi:"migrationProjectIdentifier"`
	// The property describes Amazon Resource Name (ARN) of the service access role.
	ServiceAccessRoleArn string `pulumi:"serviceAccessRoleArn"`
	// The property describes the settings for the data migration.
	SourceDataSettings []DataMigrationSourceDataSettings `pulumi:"sourceDataSettings"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a DataMigration resource.
type DataMigrationArgs struct {
	// The property describes an ARN of the data migration.
	DataMigrationIdentifier pulumi.StringPtrInput
	// The property describes a name to identify the data migration.
	DataMigrationName pulumi.StringPtrInput
	// The property describes the settings for the data migration.
	DataMigrationSettings DataMigrationSettingsPtrInput
	// The property describes the type of migration.
	DataMigrationType DataMigrationTypeInput
	// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
	MigrationProjectIdentifier pulumi.StringInput
	// The property describes Amazon Resource Name (ARN) of the service access role.
	ServiceAccessRoleArn pulumi.StringInput
	// The property describes the settings for the data migration.
	SourceDataSettings DataMigrationSourceDataSettingsArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
}

func (DataMigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMigrationArgs)(nil)).Elem()
}

type DataMigrationInput interface {
	pulumi.Input

	ToDataMigrationOutput() DataMigrationOutput
	ToDataMigrationOutputWithContext(ctx context.Context) DataMigrationOutput
}

func (*DataMigration) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMigration)(nil)).Elem()
}

func (i *DataMigration) ToDataMigrationOutput() DataMigrationOutput {
	return i.ToDataMigrationOutputWithContext(context.Background())
}

func (i *DataMigration) ToDataMigrationOutputWithContext(ctx context.Context) DataMigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMigrationOutput)
}

type DataMigrationOutput struct{ *pulumi.OutputState }

func (DataMigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMigration)(nil)).Elem()
}

func (o DataMigrationOutput) ToDataMigrationOutput() DataMigrationOutput {
	return o
}

func (o DataMigrationOutput) ToDataMigrationOutputWithContext(ctx context.Context) DataMigrationOutput {
	return o
}

// The property describes an ARN of the data migration.
func (o DataMigrationOutput) DataMigrationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringOutput { return v.DataMigrationArn }).(pulumi.StringOutput)
}

// The property describes the create time of the data migration.
func (o DataMigrationOutput) DataMigrationCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringOutput { return v.DataMigrationCreateTime }).(pulumi.StringOutput)
}

// The property describes an ARN of the data migration.
func (o DataMigrationOutput) DataMigrationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringPtrOutput { return v.DataMigrationIdentifier }).(pulumi.StringPtrOutput)
}

// The property describes a name to identify the data migration.
func (o DataMigrationOutput) DataMigrationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringPtrOutput { return v.DataMigrationName }).(pulumi.StringPtrOutput)
}

// The property describes the settings for the data migration.
func (o DataMigrationOutput) DataMigrationSettings() DataMigrationSettingsPtrOutput {
	return o.ApplyT(func(v *DataMigration) DataMigrationSettingsPtrOutput { return v.DataMigrationSettings }).(DataMigrationSettingsPtrOutput)
}

// The property describes the type of migration.
func (o DataMigrationOutput) DataMigrationType() DataMigrationTypeOutput {
	return o.ApplyT(func(v *DataMigration) DataMigrationTypeOutput { return v.DataMigrationType }).(DataMigrationTypeOutput)
}

// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
func (o DataMigrationOutput) MigrationProjectIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringOutput { return v.MigrationProjectIdentifier }).(pulumi.StringOutput)
}

// The property describes Amazon Resource Name (ARN) of the service access role.
func (o DataMigrationOutput) ServiceAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMigration) pulumi.StringOutput { return v.ServiceAccessRoleArn }).(pulumi.StringOutput)
}

// The property describes the settings for the data migration.
func (o DataMigrationOutput) SourceDataSettings() DataMigrationSourceDataSettingsArrayOutput {
	return o.ApplyT(func(v *DataMigration) DataMigrationSourceDataSettingsArrayOutput { return v.SourceDataSettings }).(DataMigrationSourceDataSettingsArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o DataMigrationOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *DataMigration) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataMigrationInput)(nil)).Elem(), &DataMigration{})
	pulumi.RegisterOutputType(DataMigrationOutput{})
}
