// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb
{
    public static class GetTable
    {
        /// <summary>
        /// The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
        ///  You should be aware of the following behaviors when working with DDB tables:
        ///   +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).
        ///   
        ///    Our guidance is to use the latest schema documented for your CFNlong templates. This schema supports the provisioning of all table settings below. When using this schema in your CFNlong templates, please ensure that your Identity and Access Management (IAM) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
        /// </summary>
        public static Task<GetTableResult> InvokeAsync(GetTableArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTableResult>("aws-native:dynamodb:getTable", args ?? new GetTableArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
        ///  You should be aware of the following behaviors when working with DDB tables:
        ///   +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).
        ///   
        ///    Our guidance is to use the latest schema documented for your CFNlong templates. This schema supports the provisioning of all table settings below. When using this schema in your CFNlong templates, please ensure that your Identity and Access Management (IAM) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("aws-native:dynamodb:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
        ///  You should be aware of the following behaviors when working with DDB tables:
        ///   +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).
        ///   
        ///    Our guidance is to use the latest schema documented for your CFNlong templates. This schema supports the provisioning of all table settings below. When using this schema in your CFNlong templates, please ensure that your Identity and Access Management (IAM) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
        /// </summary>
        public static Output<GetTableResult> Invoke(GetTableInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTableResult>("aws-native:dynamodb:getTable", args ?? new GetTableInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTableArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("tableName", required: true)]
        public string TableName { get; set; } = null!;

        public GetTableArgs()
        {
        }
        public static new GetTableArgs Empty => new GetTableArgs();
    }

    public sealed class GetTableInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public GetTableInvokeArgs()
        {
        }
        public static new GetTableInvokeArgs Empty => new GetTableInvokeArgs();
    }


    [OutputType]
    public sealed class GetTableResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` .
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// A list of attributes that describe the key schema for the table and indexes.
        ///  This property is required to create a DDB table.
        ///  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
        /// </summary>
        public readonly ImmutableArray<Outputs.TableAttributeDefinition> AttributeDefinitions;
        /// <summary>
        /// Specify how you are charged for read and write throughput and how you manage capacity.
        ///  Valid values include:
        ///   +  ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for most DynamoDB workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-demand capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html). 
        ///   +  ``PROVISIONED`` - We recommend using ``PROVISIONED`` for steady workloads with predictable growth where capacity requirements can be reliably forecasted. ``PROVISIONED`` sets the billing mode to [Provisioned capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html).
        ///   
        ///  If not specified, the default is ``PROVISIONED``.
        /// </summary>
        public readonly string? BillingMode;
        /// <summary>
        /// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
        /// </summary>
        public readonly Outputs.TableContributorInsightsSpecification? ContributorInsightsSpecification;
        /// <summary>
        /// Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
        /// </summary>
        public readonly bool? DeletionProtectionEnabled;
        /// <summary>
        /// Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
        ///   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
        ///  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
        ///  Updates are not supported. The following are exceptions:
        ///   +  If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.
        ///   +  You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.
        /// </summary>
        public readonly ImmutableArray<Outputs.TableGlobalSecondaryIndex> GlobalSecondaryIndexes;
        /// <summary>
        /// Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
        /// </summary>
        public readonly Union<ImmutableArray<Outputs.TableKeySchema>, object>? KeySchema;
        /// <summary>
        /// The Kinesis Data Streams configuration for the specified table.
        /// </summary>
        public readonly Outputs.TableKinesisStreamSpecification? KinesisStreamSpecification;
        /// <summary>
        /// Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
        /// </summary>
        public readonly ImmutableArray<Outputs.TableLocalSecondaryIndex> LocalSecondaryIndexes;
        /// <summary>
        /// Sets the maximum number of read and write units for the specified on-demand table. If you use this property, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both.
        /// </summary>
        public readonly Outputs.TableOnDemandThroughput? OnDemandThroughput;
        /// <summary>
        /// The settings used to enable point in time recovery.
        /// </summary>
        public readonly Outputs.TablePointInTimeRecoverySpecification? PointInTimeRecoverySpecification;
        /// <summary>
        /// Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
        ///  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
        /// </summary>
        public readonly Outputs.TableProvisionedThroughput? ProvisionedThroughput;
        /// <summary>
        /// An AWS resource-based policy document in JSON format that will be attached to the table.
        ///  When you attach a resource-based policy while creating a table, the policy application is *strongly consistent*.
        ///  The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this limit. For a full list of all considerations that apply for resource-based policies, see [Resource-based policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html).
        ///   You need to specify the ``CreateTable`` and ``PutResourcePolicy`` IAM actions for authorizing a user to create a table with a resource-based policy.
        /// </summary>
        public readonly Outputs.TableResourcePolicy? ResourcePolicy;
        /// <summary>
        /// Specifies the settings to enable server-side encryption.
        /// </summary>
        public readonly Outputs.TableSseSpecification? SseSpecification;
        /// <summary>
        /// The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` .
        /// 
        /// &gt; You must specify the `StreamSpecification` property to use this attribute.
        /// </summary>
        public readonly string? StreamArn;
        /// <summary>
        /// The settings for the DDB table stream, which capture changes to items stored in the table.
        /// </summary>
        public readonly Outputs.TableStreamSpecification? StreamSpecification;
        /// <summary>
        /// The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
        /// </summary>
        public readonly string? TableClass;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        ///  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// Specifies the Time to Live (TTL) settings for the table.
        ///   For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
        /// </summary>
        public readonly Outputs.TableTimeToLiveSpecification? TimeToLiveSpecification;
        /// <summary>
        /// Represents the warm throughput (in read units per second and write units per second) for creating a table.
        /// </summary>
        public readonly Outputs.TableWarmThroughput? WarmThroughput;

        [OutputConstructor]
        private GetTableResult(
            string? arn,

            ImmutableArray<Outputs.TableAttributeDefinition> attributeDefinitions,

            string? billingMode,

            Outputs.TableContributorInsightsSpecification? contributorInsightsSpecification,

            bool? deletionProtectionEnabled,

            ImmutableArray<Outputs.TableGlobalSecondaryIndex> globalSecondaryIndexes,

            Union<ImmutableArray<Outputs.TableKeySchema>, object>? keySchema,

            Outputs.TableKinesisStreamSpecification? kinesisStreamSpecification,

            ImmutableArray<Outputs.TableLocalSecondaryIndex> localSecondaryIndexes,

            Outputs.TableOnDemandThroughput? onDemandThroughput,

            Outputs.TablePointInTimeRecoverySpecification? pointInTimeRecoverySpecification,

            Outputs.TableProvisionedThroughput? provisionedThroughput,

            Outputs.TableResourcePolicy? resourcePolicy,

            Outputs.TableSseSpecification? sseSpecification,

            string? streamArn,

            Outputs.TableStreamSpecification? streamSpecification,

            string? tableClass,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            Outputs.TableTimeToLiveSpecification? timeToLiveSpecification,

            Outputs.TableWarmThroughput? warmThroughput)
        {
            Arn = arn;
            AttributeDefinitions = attributeDefinitions;
            BillingMode = billingMode;
            ContributorInsightsSpecification = contributorInsightsSpecification;
            DeletionProtectionEnabled = deletionProtectionEnabled;
            GlobalSecondaryIndexes = globalSecondaryIndexes;
            KeySchema = keySchema;
            KinesisStreamSpecification = kinesisStreamSpecification;
            LocalSecondaryIndexes = localSecondaryIndexes;
            OnDemandThroughput = onDemandThroughput;
            PointInTimeRecoverySpecification = pointInTimeRecoverySpecification;
            ProvisionedThroughput = provisionedThroughput;
            ResourcePolicy = resourcePolicy;
            SseSpecification = sseSpecification;
            StreamArn = streamArn;
            StreamSpecification = streamSpecification;
            TableClass = tableClass;
            Tags = tags;
            TimeToLiveSpecification = timeToLiveSpecification;
            WarmThroughput = warmThroughput;
        }
    }
}
