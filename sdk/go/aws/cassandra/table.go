// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Cassandra::Table
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("id"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				ClusteringKeyColumns: cassandra.TableClusteringKeyColumnArray{
//					&cassandra.TableClusteringKeyColumnArgs{
//						Column: &cassandra.TableColumnArgs{
//							ColumnName: pulumi.String("division"),
//							ColumnType: pulumi.String("ASCII"),
//						},
//						OrderBy: cassandra.TableClusteringKeyColumnOrderByAsc,
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//				BillingMode: &cassandra.TableBillingModeArgs{
//					Mode: cassandra.TableModeProvisioned,
//					ProvisionedThroughput: &cassandra.TableProvisionedThroughputArgs{
//						ReadCapacityUnits:  pulumi.Int(5),
//						WriteCapacityUnits: pulumi.Int(5),
//					},
//				},
//				ClientSideTimestampsEnabled: pulumi.Bool(true),
//				DefaultTimeToLive:           pulumi.Int(63072000),
//				PointInTimeRecoveryEnabled:  pulumi.Bool(true),
//				Tags: aws.TagArray{
//					&aws.TagArgs{
//						Key:   pulumi.String("tag1"),
//						Value: pulumi.String("val1"),
//					},
//					&aws.TagArgs{
//						Key:   pulumi.String("tag2"),
//						Value: pulumi.String("val2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("id"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				ClusteringKeyColumns: cassandra.TableClusteringKeyColumnArray{
//					&cassandra.TableClusteringKeyColumnArgs{
//						Column: &cassandra.TableColumnArgs{
//							ColumnName: pulumi.String("division"),
//							ColumnType: pulumi.String("ASCII"),
//						},
//						OrderBy: cassandra.TableClusteringKeyColumnOrderByAsc,
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//				BillingMode: &cassandra.TableBillingModeArgs{
//					Mode: cassandra.TableModeProvisioned,
//					ProvisionedThroughput: &cassandra.TableProvisionedThroughputArgs{
//						ReadCapacityUnits:  pulumi.Int(5),
//						WriteCapacityUnits: pulumi.Int(5),
//					},
//				},
//				ClientSideTimestampsEnabled: pulumi.Bool(true),
//				DefaultTimeToLive:           pulumi.Int(63072000),
//				PointInTimeRecoveryEnabled:  pulumi.Bool(true),
//				Tags: aws.TagArray{
//					&aws.TagArgs{
//						Key:   pulumi.String("tag1"),
//						Value: pulumi.String("val1"),
//					},
//					&aws.TagArgs{
//						Key:   pulumi.String("tag2"),
//						Value: pulumi.String("val2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("id"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				ClusteringKeyColumns: cassandra.TableClusteringKeyColumnArray{
//					&cassandra.TableClusteringKeyColumnArgs{
//						Column: &cassandra.TableColumnArgs{
//							ColumnName: pulumi.String("division"),
//							ColumnType: pulumi.String("ASCII"),
//						},
//						OrderBy: cassandra.TableClusteringKeyColumnOrderByAsc,
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//				BillingMode: &cassandra.TableBillingModeArgs{
//					Mode: cassandra.TableModeProvisioned,
//					ProvisionedThroughput: &cassandra.TableProvisionedThroughputArgs{
//						ReadCapacityUnits:  pulumi.Int(5),
//						WriteCapacityUnits: pulumi.Int(5),
//					},
//				},
//				DefaultTimeToLive: pulumi.Int(63072000),
//				EncryptionSpecification: &cassandra.TableEncryptionSpecificationArgs{
//					EncryptionType:   cassandra.TableEncryptionTypeCustomerManagedKmsKey,
//					KmsKeyIdentifier: pulumi.String("arn:aws:kms:eu-west-1:5555555555555:key/11111111-1111-111-1111-111111111111"),
//				},
//				PointInTimeRecoveryEnabled: pulumi.Bool(true),
//				Tags: aws.TagArray{
//					&aws.TagArgs{
//						Key:   pulumi.String("tag1"),
//						Value: pulumi.String("val1"),
//					},
//					&aws.TagArgs{
//						Key:   pulumi.String("tag2"),
//						Value: pulumi.String("val2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myNewTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("id"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				ClusteringKeyColumns: cassandra.TableClusteringKeyColumnArray{
//					&cassandra.TableClusteringKeyColumnArgs{
//						Column: &cassandra.TableColumnArgs{
//							ColumnName: pulumi.String("division"),
//							ColumnType: pulumi.String("ASCII"),
//						},
//						OrderBy: cassandra.TableClusteringKeyColumnOrderByAsc,
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//				BillingMode: &cassandra.TableBillingModeArgs{
//					Mode: cassandra.TableModeProvisioned,
//					ProvisionedThroughput: &cassandra.TableProvisionedThroughputArgs{
//						ReadCapacityUnits:  pulumi.Int(5),
//						WriteCapacityUnits: pulumi.Int(5),
//					},
//				},
//				DefaultTimeToLive: pulumi.Int(63072000),
//				EncryptionSpecification: &cassandra.TableEncryptionSpecificationArgs{
//					EncryptionType:   cassandra.TableEncryptionTypeCustomerManagedKmsKey,
//					KmsKeyIdentifier: pulumi.String("arn:aws:kms:eu-west-1:5555555555555:key/11111111-1111-111-1111-111111111111"),
//				},
//				PointInTimeRecoveryEnabled: pulumi.Bool(true),
//				Tags: aws.TagArray{
//					&aws.TagArgs{
//						Key:   pulumi.String("tag1"),
//						Value: pulumi.String("val1"),
//					},
//					&aws.TagArgs{
//						Key:   pulumi.String("tag2"),
//						Value: pulumi.String("val2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/cassandra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cassandra.NewTable(ctx, "myTable", &cassandra.TableArgs{
//				KeyspaceName: pulumi.String("my_keyspace"),
//				TableName:    pulumi.String("my_table"),
//				PartitionKeyColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("Message"),
//						ColumnType: pulumi.String("ASCII"),
//					},
//				},
//				RegularColumns: cassandra.TableColumnArray{
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("name"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("region"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("project"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("role"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("pay_scale"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("vacation_hrs"),
//						ColumnType: pulumi.String("FLOAT"),
//					},
//					&cassandra.TableColumnArgs{
//						ColumnName: pulumi.String("manager_id"),
//						ColumnType: pulumi.String("TEXT"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Table struct {
	pulumi.CustomResourceState

	// The optional auto scaling capacity settings for a table in provisioned capacity mode.
	AutoScalingSpecifications TableAutoScalingSpecificationPtrOutput `pulumi:"autoScalingSpecifications"`
	// The billing mode for the table, which determines how you'll be charged for reads and writes:
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	BillingMode TableBillingModePtrOutput `pulumi:"billingMode"`
	// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
	CdcSpecification TableCdcSpecificationPtrOutput `pulumi:"cdcSpecification"`
	// Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.
	ClientSideTimestampsEnabled pulumi.BoolPtrOutput `pulumi:"clientSideTimestampsEnabled"`
	// Clustering key columns of the table
	ClusteringKeyColumns TableClusteringKeyColumnArrayOutput `pulumi:"clusteringKeyColumns"`
	// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
	DefaultTimeToLive pulumi.IntPtrOutput `pulumi:"defaultTimeToLive"`
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
	// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
	//
	// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
	//
	// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
	EncryptionSpecification TableEncryptionSpecificationPtrOutput `pulumi:"encryptionSpecification"`
	// Name for Cassandra keyspace
	KeyspaceName pulumi.StringOutput `pulumi:"keyspaceName"`
	// Partition key columns of the table
	PartitionKeyColumns TableColumnArrayOutput `pulumi:"partitionKeyColumns"`
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
	PointInTimeRecoveryEnabled pulumi.BoolPtrOutput `pulumi:"pointInTimeRecoveryEnabled"`
	// Non-key columns of the table
	RegularColumns TableColumnArrayOutput `pulumi:"regularColumns"`
	// The AWS Region specific settings of a multi-Region table.
	//
	// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
	//
	// - `region` : The Region where these settings are applied. (Required)
	// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
	// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
	ReplicaSpecifications TableReplicaSpecificationArrayOutput `pulumi:"replicaSpecifications"`
	// Name for Cassandra table
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// An array of key-value pairs to apply to this resource
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyspaceName == nil {
		return nil, errors.New("invalid value for required argument 'KeyspaceName'")
	}
	if args.PartitionKeyColumns == nil {
		return nil, errors.New("invalid value for required argument 'PartitionKeyColumns'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clientSideTimestampsEnabled",
		"clusteringKeyColumns[*]",
		"keyspaceName",
		"partitionKeyColumns[*]",
		"tableName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("aws-native:cassandra:Table", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:cassandra:Table", name, id, state, &resource, opts...)
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
	// The optional auto scaling capacity settings for a table in provisioned capacity mode.
	AutoScalingSpecifications *TableAutoScalingSpecification `pulumi:"autoScalingSpecifications"`
	// The billing mode for the table, which determines how you'll be charged for reads and writes:
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	BillingMode *TableBillingMode `pulumi:"billingMode"`
	// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
	CdcSpecification *TableCdcSpecification `pulumi:"cdcSpecification"`
	// Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.
	ClientSideTimestampsEnabled *bool `pulumi:"clientSideTimestampsEnabled"`
	// Clustering key columns of the table
	ClusteringKeyColumns []TableClusteringKeyColumn `pulumi:"clusteringKeyColumns"`
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
	// Name for Cassandra keyspace
	KeyspaceName string `pulumi:"keyspaceName"`
	// Partition key columns of the table
	PartitionKeyColumns []TableColumn `pulumi:"partitionKeyColumns"`
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
	PointInTimeRecoveryEnabled *bool `pulumi:"pointInTimeRecoveryEnabled"`
	// Non-key columns of the table
	RegularColumns []TableColumn `pulumi:"regularColumns"`
	// The AWS Region specific settings of a multi-Region table.
	//
	// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
	//
	// - `region` : The Region where these settings are applied. (Required)
	// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
	// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
	ReplicaSpecifications []TableReplicaSpecification `pulumi:"replicaSpecifications"`
	// Name for Cassandra table
	TableName *string `pulumi:"tableName"`
	// An array of key-value pairs to apply to this resource
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The optional auto scaling capacity settings for a table in provisioned capacity mode.
	AutoScalingSpecifications TableAutoScalingSpecificationPtrInput
	// The billing mode for the table, which determines how you'll be charged for reads and writes:
	//
	// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
	// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
	//
	// If you don't specify a value for this property, then the table will use on-demand mode.
	BillingMode TableBillingModePtrInput
	// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
	CdcSpecification TableCdcSpecificationPtrInput
	// Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.
	ClientSideTimestampsEnabled pulumi.BoolPtrInput
	// Clustering key columns of the table
	ClusteringKeyColumns TableClusteringKeyColumnArrayInput
	// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
	DefaultTimeToLive pulumi.IntPtrInput
	// The encryption at rest options for the table.
	//
	// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
	// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
	//
	// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
	//
	// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
	EncryptionSpecification TableEncryptionSpecificationPtrInput
	// Name for Cassandra keyspace
	KeyspaceName pulumi.StringInput
	// Partition key columns of the table
	PartitionKeyColumns TableColumnArrayInput
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
	PointInTimeRecoveryEnabled pulumi.BoolPtrInput
	// Non-key columns of the table
	RegularColumns TableColumnArrayInput
	// The AWS Region specific settings of a multi-Region table.
	//
	// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
	//
	// - `region` : The Region where these settings are applied. (Required)
	// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
	// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
	ReplicaSpecifications TableReplicaSpecificationArrayInput
	// Name for Cassandra table
	TableName pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource
	Tags aws.TagArrayInput
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

// The optional auto scaling capacity settings for a table in provisioned capacity mode.
func (o TableOutput) AutoScalingSpecifications() TableAutoScalingSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableAutoScalingSpecificationPtrOutput { return v.AutoScalingSpecifications }).(TableAutoScalingSpecificationPtrOutput)
}

// The billing mode for the table, which determines how you'll be charged for reads and writes:
//
// - *On-demand mode* (default) - You pay based on the actual reads and writes your application performs.
// - *Provisioned mode* - Lets you specify the number of reads and writes per second that you need for your application.
//
// If you don't specify a value for this property, then the table will use on-demand mode.
func (o TableOutput) BillingMode() TableBillingModePtrOutput {
	return o.ApplyT(func(v *Table) TableBillingModePtrOutput { return v.BillingMode }).(TableBillingModePtrOutput)
}

// The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the *Amazon Keyspaces Developer Guide* .
func (o TableOutput) CdcSpecification() TableCdcSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableCdcSpecificationPtrOutput { return v.CdcSpecification }).(TableCdcSpecificationPtrOutput)
}

// Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.
func (o TableOutput) ClientSideTimestampsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.BoolPtrOutput { return v.ClientSideTimestampsEnabled }).(pulumi.BoolPtrOutput)
}

// Clustering key columns of the table
func (o TableOutput) ClusteringKeyColumns() TableClusteringKeyColumnArrayOutput {
	return o.ApplyT(func(v *Table) TableClusteringKeyColumnArrayOutput { return v.ClusteringKeyColumns }).(TableClusteringKeyColumnArrayOutput)
}

// Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
func (o TableOutput) DefaultTimeToLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.IntPtrOutput { return v.DefaultTimeToLive }).(pulumi.IntPtrOutput)
}

// The encryption at rest options for the table.
//
// - *AWS owned key* (default) - The key is owned by Amazon Keyspaces .
// - *Customer managed key* - The key is stored in your account and is created, owned, and managed by you.
//
// > If you choose encryption with a customer managed key, you must specify a valid customer managed KMS key with permissions granted to Amazon Keyspaces.
//
// For more information, see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html) in the *Amazon Keyspaces Developer Guide* .
func (o TableOutput) EncryptionSpecification() TableEncryptionSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableEncryptionSpecificationPtrOutput { return v.EncryptionSpecification }).(TableEncryptionSpecificationPtrOutput)
}

// Name for Cassandra keyspace
func (o TableOutput) KeyspaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.KeyspaceName }).(pulumi.StringOutput)
}

// Partition key columns of the table
func (o TableOutput) PartitionKeyColumns() TableColumnArrayOutput {
	return o.ApplyT(func(v *Table) TableColumnArrayOutput { return v.PartitionKeyColumns }).(TableColumnArrayOutput)
}

// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table
func (o TableOutput) PointInTimeRecoveryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.BoolPtrOutput { return v.PointInTimeRecoveryEnabled }).(pulumi.BoolPtrOutput)
}

// Non-key columns of the table
func (o TableOutput) RegularColumns() TableColumnArrayOutput {
	return o.ApplyT(func(v *Table) TableColumnArrayOutput { return v.RegularColumns }).(TableColumnArrayOutput)
}

// The AWS Region specific settings of a multi-Region table.
//
// For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.
//
// - `region` : The Region where these settings are applied. (Required)
// - `readCapacityUnits` : The provisioned read capacity units. (Optional)
// - `readCapacityAutoScaling` : The read capacity auto scaling settings for the table. (Optional)
func (o TableOutput) ReplicaSpecifications() TableReplicaSpecificationArrayOutput {
	return o.ApplyT(func(v *Table) TableReplicaSpecificationArrayOutput { return v.ReplicaSpecifications }).(TableReplicaSpecificationArrayOutput)
}

// Name for Cassandra table
func (o TableOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource
func (o TableOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Table) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}
