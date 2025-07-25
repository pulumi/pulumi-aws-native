// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Cassandra::Table
func LookupTable(ctx *pulumi.Context, args *LookupTableArgs, opts ...pulumi.InvokeOption) (*LookupTableResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTableResult
	err := ctx.Invoke("aws-native:cassandra:getTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTableArgs struct {
	// Name for Cassandra keyspace
	KeyspaceName string `pulumi:"keyspaceName"`
	// Name for Cassandra table
	TableName string `pulumi:"tableName"`
}

type LookupTableResult struct {
	// The billing mode for the table, which determines how you'll be charged for reads and writes:
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	BillingMode *TableBillingMode `pulumi:"billingMode"`
	// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
	CdcSpecification *TableCdcSpecification `pulumi:"cdcSpecification"`
	// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
	DefaultTimeToLive *int `pulumi:"defaultTimeToLive"`
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
	// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
	//
	// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
	//
	// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
	EncryptionSpecification *TableEncryptionSpecification `pulumi:"encryptionSpecification"`
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
	PointInTimeRecoveryEnabled *bool `pulumi:"pointInTimeRecoveryEnabled"`
	// Non-key columns of the table
	RegularColumns []TableColumn `pulumi:"regularColumns"`
	// An array of key-value pairs to apply to this resource
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupTableOutput(ctx *pulumi.Context, args LookupTableOutputArgs, opts ...pulumi.InvokeOption) LookupTableResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTableResultOutput, error) {
			args := v.(LookupTableArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cassandra:getTable", args, LookupTableResultOutput{}, options).(LookupTableResultOutput), nil
		}).(LookupTableResultOutput)
}

type LookupTableOutputArgs struct {
	// Name for Cassandra keyspace
	KeyspaceName pulumi.StringInput `pulumi:"keyspaceName"`
	// Name for Cassandra table
	TableName pulumi.StringInput `pulumi:"tableName"`
}

func (LookupTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableArgs)(nil)).Elem()
}

type LookupTableResultOutput struct{ *pulumi.OutputState }

func (LookupTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTableResult)(nil)).Elem()
}

func (o LookupTableResultOutput) ToLookupTableResultOutput() LookupTableResultOutput {
	return o
}

func (o LookupTableResultOutput) ToLookupTableResultOutputWithContext(ctx context.Context) LookupTableResultOutput {
	return o
}

// The billing mode for the table, which determines how you'll be charged for reads and writes:
//
// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
//
// If you don't specify a value for this property, then the table will use on-demand mode.
func (o LookupTableResultOutput) BillingMode() TableBillingModePtrOutput {
	return o.ApplyT(func(v LookupTableResult) *TableBillingMode { return v.BillingMode }).(TableBillingModePtrOutput)
}

// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
func (o LookupTableResultOutput) CdcSpecification() TableCdcSpecificationPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *TableCdcSpecification { return v.CdcSpecification }).(TableCdcSpecificationPtrOutput)
}

// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
func (o LookupTableResultOutput) DefaultTimeToLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *int { return v.DefaultTimeToLive }).(pulumi.IntPtrOutput)
}

// The encryption at rest options for the table.
//
// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
//
// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
//
// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
func (o LookupTableResultOutput) EncryptionSpecification() TableEncryptionSpecificationPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *TableEncryptionSpecification { return v.EncryptionSpecification }).(TableEncryptionSpecificationPtrOutput)
}

// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
func (o LookupTableResultOutput) PointInTimeRecoveryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTableResult) *bool { return v.PointInTimeRecoveryEnabled }).(pulumi.BoolPtrOutput)
}

// Non-key columns of the table
func (o LookupTableResultOutput) RegularColumns() TableColumnArrayOutput {
	return o.ApplyT(func(v LookupTableResult) []TableColumn { return v.RegularColumns }).(TableColumnArrayOutput)
}

// An array of key-value pairs to apply to this resource
func (o LookupTableResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupTableResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTableResultOutput{})
}
