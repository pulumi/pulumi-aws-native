// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package timestream

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Timestream::Table resource creates a Timestream Table.
type Table struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name for the database which the table to be created belongs to.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The properties that determine whether magnetic store writes are enabled.
	MagneticStoreWriteProperties MagneticStoreWritePropertiesPropertiesPtrOutput `pulumi:"magneticStoreWriteProperties"`
	// The table name exposed as a read-only attribute.
	Name pulumi.StringOutput `pulumi:"name"`
	// The retention duration of the memory store and the magnetic store.
	RetentionProperties RetentionPropertiesPropertiesPtrOutput `pulumi:"retentionProperties"`
	// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// An array of key-value pairs to apply to this resource.
	Tags TableTagArrayOutput `pulumi:"tags"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	var resource Table
	err := ctx.RegisterResource("aws-native:timestream:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws-native:timestream:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The name for the database which the table to be created belongs to.
	DatabaseName string `pulumi:"databaseName"`
	// The properties that determine whether magnetic store writes are enabled.
	MagneticStoreWriteProperties *MagneticStoreWritePropertiesProperties `pulumi:"magneticStoreWriteProperties"`
	// The retention duration of the memory store and the magnetic store.
	RetentionProperties *RetentionPropertiesProperties `pulumi:"retentionProperties"`
	// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
	TableName *string `pulumi:"tableName"`
	// An array of key-value pairs to apply to this resource.
	Tags []TableTag `pulumi:"tags"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The name for the database which the table to be created belongs to.
	DatabaseName pulumi.StringInput
	// The properties that determine whether magnetic store writes are enabled.
	MagneticStoreWriteProperties MagneticStoreWritePropertiesPropertiesPtrInput
	// The retention duration of the memory store and the magnetic store.
	RetentionProperties RetentionPropertiesPropertiesPtrInput
	// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
	TableName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags TableTagArrayInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name for the database which the table to be created belongs to.
func (o TableOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// The properties that determine whether magnetic store writes are enabled.
func (o TableOutput) MagneticStoreWriteProperties() MagneticStoreWritePropertiesPropertiesPtrOutput {
	return o.ApplyT(func(v *Table) MagneticStoreWritePropertiesPropertiesPtrOutput { return v.MagneticStoreWriteProperties }).(MagneticStoreWritePropertiesPropertiesPtrOutput)
}

// The table name exposed as a read-only attribute.
func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The retention duration of the memory store and the magnetic store.
func (o TableOutput) RetentionProperties() RetentionPropertiesPropertiesPtrOutput {
	return o.ApplyT(func(v *Table) RetentionPropertiesPropertiesPtrOutput { return v.RetentionProperties }).(RetentionPropertiesPropertiesPtrOutput)
}

// The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.
func (o TableOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o TableOutput) Tags() TableTagArrayOutput {
	return o.ApplyT(func(v *Table) TableTagArrayOutput { return v.Tags }).(TableTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}