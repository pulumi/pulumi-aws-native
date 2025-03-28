// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A resource schema representing a Lake Formation Data Cells Filter.
type DataCellsFilter struct {
	pulumi.CustomResourceState

	// A list of columns to be included in this Data Cells Filter.
	ColumnNames pulumi.StringArrayOutput `pulumi:"columnNames"`
	// An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
	ColumnWildcard DataCellsFilterColumnWildcardPtrOutput `pulumi:"columnWildcard"`
	// The name of the Database that the Table resides in.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The desired name of the Data Cells Filter.
	Name pulumi.StringOutput `pulumi:"name"`
	// An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
	RowFilter DataCellsFilterRowFilterPtrOutput `pulumi:"rowFilter"`
	// The Catalog Id of the Table on which to create a Data Cells Filter.
	TableCatalogId pulumi.StringOutput `pulumi:"tableCatalogId"`
	// The name of the Table to create a Data Cells Filter for.
	TableName pulumi.StringOutput `pulumi:"tableName"`
}

// NewDataCellsFilter registers a new resource with the given unique name, arguments, and options.
func NewDataCellsFilter(ctx *pulumi.Context,
	name string, args *DataCellsFilterArgs, opts ...pulumi.ResourceOption) (*DataCellsFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.TableCatalogId == nil {
		return nil, errors.New("invalid value for required argument 'TableCatalogId'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"columnNames[*]",
		"columnWildcard",
		"databaseName",
		"name",
		"rowFilter",
		"tableCatalogId",
		"tableName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataCellsFilter
	err := ctx.RegisterResource("aws-native:lakeformation:DataCellsFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCellsFilter gets an existing DataCellsFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCellsFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCellsFilterState, opts ...pulumi.ResourceOption) (*DataCellsFilter, error) {
	var resource DataCellsFilter
	err := ctx.ReadResource("aws-native:lakeformation:DataCellsFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCellsFilter resources.
type dataCellsFilterState struct {
}

type DataCellsFilterState struct {
}

func (DataCellsFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCellsFilterState)(nil)).Elem()
}

type dataCellsFilterArgs struct {
	// A list of columns to be included in this Data Cells Filter.
	ColumnNames []string `pulumi:"columnNames"`
	// An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
	ColumnWildcard *DataCellsFilterColumnWildcard `pulumi:"columnWildcard"`
	// The name of the Database that the Table resides in.
	DatabaseName string `pulumi:"databaseName"`
	// The desired name of the Data Cells Filter.
	Name *string `pulumi:"name"`
	// An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
	RowFilter *DataCellsFilterRowFilter `pulumi:"rowFilter"`
	// The Catalog Id of the Table on which to create a Data Cells Filter.
	TableCatalogId string `pulumi:"tableCatalogId"`
	// The name of the Table to create a Data Cells Filter for.
	TableName string `pulumi:"tableName"`
}

// The set of arguments for constructing a DataCellsFilter resource.
type DataCellsFilterArgs struct {
	// A list of columns to be included in this Data Cells Filter.
	ColumnNames pulumi.StringArrayInput
	// An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
	ColumnWildcard DataCellsFilterColumnWildcardPtrInput
	// The name of the Database that the Table resides in.
	DatabaseName pulumi.StringInput
	// The desired name of the Data Cells Filter.
	Name pulumi.StringPtrInput
	// An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
	RowFilter DataCellsFilterRowFilterPtrInput
	// The Catalog Id of the Table on which to create a Data Cells Filter.
	TableCatalogId pulumi.StringInput
	// The name of the Table to create a Data Cells Filter for.
	TableName pulumi.StringInput
}

func (DataCellsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCellsFilterArgs)(nil)).Elem()
}

type DataCellsFilterInput interface {
	pulumi.Input

	ToDataCellsFilterOutput() DataCellsFilterOutput
	ToDataCellsFilterOutputWithContext(ctx context.Context) DataCellsFilterOutput
}

func (*DataCellsFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCellsFilter)(nil)).Elem()
}

func (i *DataCellsFilter) ToDataCellsFilterOutput() DataCellsFilterOutput {
	return i.ToDataCellsFilterOutputWithContext(context.Background())
}

func (i *DataCellsFilter) ToDataCellsFilterOutputWithContext(ctx context.Context) DataCellsFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCellsFilterOutput)
}

type DataCellsFilterOutput struct{ *pulumi.OutputState }

func (DataCellsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCellsFilter)(nil)).Elem()
}

func (o DataCellsFilterOutput) ToDataCellsFilterOutput() DataCellsFilterOutput {
	return o
}

func (o DataCellsFilterOutput) ToDataCellsFilterOutputWithContext(ctx context.Context) DataCellsFilterOutput {
	return o
}

// A list of columns to be included in this Data Cells Filter.
func (o DataCellsFilterOutput) ColumnNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataCellsFilter) pulumi.StringArrayOutput { return v.ColumnNames }).(pulumi.StringArrayOutput)
}

// An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required
func (o DataCellsFilterOutput) ColumnWildcard() DataCellsFilterColumnWildcardPtrOutput {
	return o.ApplyT(func(v *DataCellsFilter) DataCellsFilterColumnWildcardPtrOutput { return v.ColumnWildcard }).(DataCellsFilterColumnWildcardPtrOutput)
}

// The name of the Database that the Table resides in.
func (o DataCellsFilterOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCellsFilter) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// The desired name of the Data Cells Filter.
func (o DataCellsFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCellsFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required
func (o DataCellsFilterOutput) RowFilter() DataCellsFilterRowFilterPtrOutput {
	return o.ApplyT(func(v *DataCellsFilter) DataCellsFilterRowFilterPtrOutput { return v.RowFilter }).(DataCellsFilterRowFilterPtrOutput)
}

// The Catalog Id of the Table on which to create a Data Cells Filter.
func (o DataCellsFilterOutput) TableCatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCellsFilter) pulumi.StringOutput { return v.TableCatalogId }).(pulumi.StringOutput)
}

// The name of the Table to create a Data Cells Filter for.
func (o DataCellsFilterOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCellsFilter) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCellsFilterInput)(nil)).Elem(), &DataCellsFilter{})
	pulumi.RegisterOutputType(DataCellsFilterOutput{})
}
