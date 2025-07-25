// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.
 *  You should be aware of the following behaviors when working with DDB tables:
 *   +  CFNlong typically creates DDB tables in parallel. However, if your template includes multiple DDB tables with indexes, you must declare dependencies so that the tables are created sequentially. DDBlong limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).
 *
 *    Our guidance is to use the latest schema documented for your CFNlong templates. This schema supports the provisioning of all table settings below. When using this schema in your CFNlong templates, please ensure that your Identity and Access Management (IAM) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myDynamoDBTable = new aws_native.dynamodb.Table("myDynamoDBTable", {
 *     attributeDefinitions: [
 *         {
 *             attributeName: "Album",
 *             attributeType: "S",
 *         },
 *         {
 *             attributeName: "Artist",
 *             attributeType: "S",
 *         },
 *         {
 *             attributeName: "Sales",
 *             attributeType: "N",
 *         },
 *         {
 *             attributeName: "NumberOfSongs",
 *             attributeType: "N",
 *         },
 *     ],
 *     keySchema: [
 *         {
 *             attributeName: "Album",
 *             keyType: "HASH",
 *         },
 *         {
 *             attributeName: "Artist",
 *             keyType: "RANGE",
 *         },
 *     ],
 *     provisionedThroughput: {
 *         readCapacityUnits: 5,
 *         writeCapacityUnits: 5,
 *     },
 *     tableName: "myTableName",
 *     globalSecondaryIndexes: [
 *         {
 *             indexName: "myGSI",
 *             keySchema: [
 *                 {
 *                     attributeName: "Sales",
 *                     keyType: "HASH",
 *                 },
 *                 {
 *                     attributeName: "Artist",
 *                     keyType: "RANGE",
 *                 },
 *             ],
 *             projection: {
 *                 nonKeyAttributes: [
 *                     "Album",
 *                     "NumberOfSongs",
 *                 ],
 *                 projectionType: "INCLUDE",
 *             },
 *             provisionedThroughput: {
 *                 readCapacityUnits: 5,
 *                 writeCapacityUnits: 5,
 *             },
 *         },
 *         {
 *             indexName: "myGSI2",
 *             keySchema: [
 *                 {
 *                     attributeName: "NumberOfSongs",
 *                     keyType: "HASH",
 *                 },
 *                 {
 *                     attributeName: "Sales",
 *                     keyType: "RANGE",
 *                 },
 *             ],
 *             projection: {
 *                 nonKeyAttributes: [
 *                     "Album",
 *                     "Artist",
 *                 ],
 *                 projectionType: "INCLUDE",
 *             },
 *             provisionedThroughput: {
 *                 readCapacityUnits: 5,
 *                 writeCapacityUnits: 5,
 *             },
 *         },
 *     ],
 *     localSecondaryIndexes: [{
 *         indexName: "myLSI",
 *         keySchema: [
 *             {
 *                 attributeName: "Album",
 *                 keyType: "HASH",
 *             },
 *             {
 *                 attributeName: "Sales",
 *                 keyType: "RANGE",
 *             },
 *         ],
 *         projection: {
 *             nonKeyAttributes: [
 *                 "Artist",
 *                 "NumberOfSongs",
 *             ],
 *             projectionType: "INCLUDE",
 *         },
 *     }],
 * });
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const myDynamoDBTable = new aws_native.dynamodb.Table("myDynamoDBTable", {
 *     attributeDefinitions: [
 *         {
 *             attributeName: "Album",
 *             attributeType: "S",
 *         },
 *         {
 *             attributeName: "Artist",
 *             attributeType: "S",
 *         },
 *         {
 *             attributeName: "Sales",
 *             attributeType: "N",
 *         },
 *         {
 *             attributeName: "NumberOfSongs",
 *             attributeType: "N",
 *         },
 *     ],
 *     keySchema: [
 *         {
 *             attributeName: "Album",
 *             keyType: "HASH",
 *         },
 *         {
 *             attributeName: "Artist",
 *             keyType: "RANGE",
 *         },
 *     ],
 *     provisionedThroughput: {
 *         readCapacityUnits: 5,
 *         writeCapacityUnits: 5,
 *     },
 *     tableName: "myTableName",
 *     globalSecondaryIndexes: [
 *         {
 *             indexName: "myGSI",
 *             keySchema: [
 *                 {
 *                     attributeName: "Sales",
 *                     keyType: "HASH",
 *                 },
 *                 {
 *                     attributeName: "Artist",
 *                     keyType: "RANGE",
 *                 },
 *             ],
 *             projection: {
 *                 nonKeyAttributes: [
 *                     "Album",
 *                     "NumberOfSongs",
 *                 ],
 *                 projectionType: "INCLUDE",
 *             },
 *             provisionedThroughput: {
 *                 readCapacityUnits: 5,
 *                 writeCapacityUnits: 5,
 *             },
 *         },
 *         {
 *             indexName: "myGSI2",
 *             keySchema: [
 *                 {
 *                     attributeName: "NumberOfSongs",
 *                     keyType: "HASH",
 *                 },
 *                 {
 *                     attributeName: "Sales",
 *                     keyType: "RANGE",
 *                 },
 *             ],
 *             projection: {
 *                 nonKeyAttributes: [
 *                     "Album",
 *                     "Artist",
 *                 ],
 *                 projectionType: "INCLUDE",
 *             },
 *             provisionedThroughput: {
 *                 readCapacityUnits: 5,
 *                 writeCapacityUnits: 5,
 *             },
 *         },
 *     ],
 *     localSecondaryIndexes: [{
 *         indexName: "myLSI",
 *         keySchema: [
 *             {
 *                 attributeName: "Album",
 *                 keyType: "HASH",
 *             },
 *             {
 *                 attributeName: "Sales",
 *                 keyType: "RANGE",
 *             },
 *         ],
 *         projection: {
 *             nonKeyAttributes: [
 *                 "Artist",
 *                 "NumberOfSongs",
 *             ],
 *             projectionType: "INCLUDE",
 *         },
 *     }],
 * });
 *
 * ```
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dynamodb:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` .
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of attributes that describe the key schema for the table and indexes.
     *  This property is required to create a DDB table.
     *  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
     */
    public readonly attributeDefinitions!: pulumi.Output<outputs.dynamodb.TableAttributeDefinition[] | undefined>;
    /**
     * Specify how you are charged for read and write throughput and how you manage capacity.
     *  Valid values include:
     *   +  ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for most DynamoDB workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-demand capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html). 
     *   +  ``PROVISIONED`` - We recommend using ``PROVISIONED`` for steady workloads with predictable growth where capacity requirements can be reliably forecasted. ``PROVISIONED`` sets the billing mode to [Provisioned capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html).
     *   
     *  If not specified, the default is ``PROVISIONED``.
     */
    public readonly billingMode!: pulumi.Output<string | undefined>;
    /**
     * The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
     */
    public readonly contributorInsightsSpecification!: pulumi.Output<outputs.dynamodb.TableContributorInsightsSpecification | undefined>;
    /**
     * Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
     */
    public readonly deletionProtectionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
     *   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
     *  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
     *  Updates are not supported. The following are exceptions:
     *   +  If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.
     *   +  You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.
     */
    public readonly globalSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableGlobalSecondaryIndex[] | undefined>;
    /**
     * Specifies the properties of data being imported from the S3 bucket source to the" table.
     *   If you specify the ``ImportSourceSpecification`` property, and also specify either the ``StreamSpecification``, the ``TableClass`` property, the ``DeletionProtectionEnabled`` property, or the ``WarmThroughput`` property, the IAM entity creating/updating stack must have ``UpdateTable`` permission.
     */
    public readonly importSourceSpecification!: pulumi.Output<outputs.dynamodb.TableImportSourceSpecification | undefined>;
    /**
     * Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
     */
    public readonly keySchema!: pulumi.Output<outputs.dynamodb.TableKeySchema[] | any>;
    /**
     * The Kinesis Data Streams configuration for the specified table.
     */
    public readonly kinesisStreamSpecification!: pulumi.Output<outputs.dynamodb.TableKinesisStreamSpecification | undefined>;
    /**
     * Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
     */
    public readonly localSecondaryIndexes!: pulumi.Output<outputs.dynamodb.TableLocalSecondaryIndex[] | undefined>;
    /**
     * Sets the maximum number of read and write units for the specified on-demand table. If you use this property, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both.
     */
    public readonly onDemandThroughput!: pulumi.Output<outputs.dynamodb.TableOnDemandThroughput | undefined>;
    /**
     * The settings used to enable point in time recovery.
     */
    public readonly pointInTimeRecoverySpecification!: pulumi.Output<outputs.dynamodb.TablePointInTimeRecoverySpecification | undefined>;
    /**
     * Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
     *  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
     */
    public readonly provisionedThroughput!: pulumi.Output<outputs.dynamodb.TableProvisionedThroughput | undefined>;
    /**
     * An AWS resource-based policy document in JSON format that will be attached to the table.
     *  When you attach a resource-based policy while creating a table, the policy application is *strongly consistent*.
     *  The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this limit. For a full list of all considerations that apply for resource-based policies, see [Resource-based policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html).
     *   You need to specify the ``CreateTable`` and ``PutResourcePolicy`` IAM actions for authorizing a user to create a table with a resource-based policy.
     */
    public readonly resourcePolicy!: pulumi.Output<outputs.dynamodb.TableResourcePolicy | undefined>;
    /**
     * Specifies the settings to enable server-side encryption.
     */
    public readonly sseSpecification!: pulumi.Output<outputs.dynamodb.TableSseSpecification | undefined>;
    /**
     * The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` .
     *
     * > You must specify the `StreamSpecification` property to use this attribute.
     */
    public /*out*/ readonly streamArn!: pulumi.Output<string>;
    /**
     * The settings for the DDB table stream, which capture changes to items stored in the table.
     */
    public readonly streamSpecification!: pulumi.Output<outputs.dynamodb.TableStreamSpecification | undefined>;
    /**
     * The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
     */
    public readonly tableClass!: pulumi.Output<string | undefined>;
    /**
     * A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
     *   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    public readonly tableName!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     *  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * Specifies the Time to Live (TTL) settings for the table.
     *   For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
     */
    public readonly timeToLiveSpecification!: pulumi.Output<outputs.dynamodb.TableTimeToLiveSpecification | undefined>;
    /**
     * Represents the warm throughput (in read units per second and write units per second) for creating a table.
     */
    public readonly warmThroughput!: pulumi.Output<outputs.dynamodb.TableWarmThroughput | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keySchema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keySchema'");
            }
            resourceInputs["attributeDefinitions"] = args ? args.attributeDefinitions : undefined;
            resourceInputs["billingMode"] = args ? args.billingMode : undefined;
            resourceInputs["contributorInsightsSpecification"] = args ? args.contributorInsightsSpecification : undefined;
            resourceInputs["deletionProtectionEnabled"] = args ? args.deletionProtectionEnabled : undefined;
            resourceInputs["globalSecondaryIndexes"] = args ? args.globalSecondaryIndexes : undefined;
            resourceInputs["importSourceSpecification"] = args ? args.importSourceSpecification : undefined;
            resourceInputs["keySchema"] = args ? args.keySchema : undefined;
            resourceInputs["kinesisStreamSpecification"] = args ? args.kinesisStreamSpecification : undefined;
            resourceInputs["localSecondaryIndexes"] = args ? args.localSecondaryIndexes : undefined;
            resourceInputs["onDemandThroughput"] = args ? args.onDemandThroughput : undefined;
            resourceInputs["pointInTimeRecoverySpecification"] = args ? args.pointInTimeRecoverySpecification : undefined;
            resourceInputs["provisionedThroughput"] = args ? args.provisionedThroughput : undefined;
            resourceInputs["resourcePolicy"] = args ? args.resourcePolicy : undefined;
            resourceInputs["sseSpecification"] = args ? args.sseSpecification : undefined;
            resourceInputs["streamSpecification"] = args ? args.streamSpecification : undefined;
            resourceInputs["tableClass"] = args ? args.tableClass : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeToLiveSpecification"] = args ? args.timeToLiveSpecification : undefined;
            resourceInputs["warmThroughput"] = args ? args.warmThroughput : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["streamArn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["attributeDefinitions"] = undefined /*out*/;
            resourceInputs["billingMode"] = undefined /*out*/;
            resourceInputs["contributorInsightsSpecification"] = undefined /*out*/;
            resourceInputs["deletionProtectionEnabled"] = undefined /*out*/;
            resourceInputs["globalSecondaryIndexes"] = undefined /*out*/;
            resourceInputs["importSourceSpecification"] = undefined /*out*/;
            resourceInputs["keySchema"] = undefined /*out*/;
            resourceInputs["kinesisStreamSpecification"] = undefined /*out*/;
            resourceInputs["localSecondaryIndexes"] = undefined /*out*/;
            resourceInputs["onDemandThroughput"] = undefined /*out*/;
            resourceInputs["pointInTimeRecoverySpecification"] = undefined /*out*/;
            resourceInputs["provisionedThroughput"] = undefined /*out*/;
            resourceInputs["resourcePolicy"] = undefined /*out*/;
            resourceInputs["sseSpecification"] = undefined /*out*/;
            resourceInputs["streamArn"] = undefined /*out*/;
            resourceInputs["streamSpecification"] = undefined /*out*/;
            resourceInputs["tableClass"] = undefined /*out*/;
            resourceInputs["tableName"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["timeToLiveSpecification"] = undefined /*out*/;
            resourceInputs["warmThroughput"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["importSourceSpecification", "tableName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * A list of attributes that describe the key schema for the table and indexes.
     *  This property is required to create a DDB table.
     *  Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt). Replacement if you edit an existing AttributeDefinition.
     */
    attributeDefinitions?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableAttributeDefinitionArgs>[]>;
    /**
     * Specify how you are charged for read and write throughput and how you manage capacity.
     *  Valid values include:
     *   +  ``PAY_PER_REQUEST`` - We recommend using ``PAY_PER_REQUEST`` for most DynamoDB workloads. ``PAY_PER_REQUEST`` sets the billing mode to [On-demand capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html). 
     *   +  ``PROVISIONED`` - We recommend using ``PROVISIONED`` for steady workloads with predictable growth where capacity requirements can be reliably forecasted. ``PROVISIONED`` sets the billing mode to [Provisioned capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html).
     *   
     *  If not specified, the default is ``PROVISIONED``.
     */
    billingMode?: pulumi.Input<string>;
    /**
     * The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
     */
    contributorInsightsSpecification?: pulumi.Input<inputs.dynamodb.TableContributorInsightsSpecificationArgs>;
    /**
     * Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*.
     */
    deletionProtectionEnabled?: pulumi.Input<boolean>;
    /**
     * Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
     *   If you update a table to include a new global secondary index, CFNlong initiates the index creation and then proceeds with the stack update. CFNlong doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is ``ACTIVE``. You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
     *  If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index. 
     *  Updates are not supported. The following are exceptions:
     *   +  If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.
     *   +  You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.
     */
    globalSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableGlobalSecondaryIndexArgs>[]>;
    /**
     * Specifies the properties of data being imported from the S3 bucket source to the" table.
     *   If you specify the ``ImportSourceSpecification`` property, and also specify either the ``StreamSpecification``, the ``TableClass`` property, the ``DeletionProtectionEnabled`` property, or the ``WarmThroughput`` property, the IAM entity creating/updating stack must have ``UpdateTable`` permission.
     */
    importSourceSpecification?: pulumi.Input<inputs.dynamodb.TableImportSourceSpecificationArgs>;
    /**
     * Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property.
     */
    keySchema: pulumi.Input<pulumi.Input<inputs.dynamodb.TableKeySchemaArgs>[] | any>;
    /**
     * The Kinesis Data Streams configuration for the specified table.
     */
    kinesisStreamSpecification?: pulumi.Input<inputs.dynamodb.TableKinesisStreamSpecificationArgs>;
    /**
     * Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
     */
    localSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.TableLocalSecondaryIndexArgs>[]>;
    /**
     * Sets the maximum number of read and write units for the specified on-demand table. If you use this property, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both.
     */
    onDemandThroughput?: pulumi.Input<inputs.dynamodb.TableOnDemandThroughputArgs>;
    /**
     * The settings used to enable point in time recovery.
     */
    pointInTimeRecoverySpecification?: pulumi.Input<inputs.dynamodb.TablePointInTimeRecoverySpecificationArgs>;
    /**
     * Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). 
     *  If you set ``BillingMode`` as ``PROVISIONED``, you must specify this property. If you set ``BillingMode`` as ``PAY_PER_REQUEST``, you cannot specify this property.
     */
    provisionedThroughput?: pulumi.Input<inputs.dynamodb.TableProvisionedThroughputArgs>;
    /**
     * An AWS resource-based policy document in JSON format that will be attached to the table.
     *  When you attach a resource-based policy while creating a table, the policy application is *strongly consistent*.
     *  The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this limit. For a full list of all considerations that apply for resource-based policies, see [Resource-based policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html).
     *   You need to specify the ``CreateTable`` and ``PutResourcePolicy`` IAM actions for authorizing a user to create a table with a resource-based policy.
     */
    resourcePolicy?: pulumi.Input<inputs.dynamodb.TableResourcePolicyArgs>;
    /**
     * Specifies the settings to enable server-side encryption.
     */
    sseSpecification?: pulumi.Input<inputs.dynamodb.TableSseSpecificationArgs>;
    /**
     * The settings for the DDB table stream, which capture changes to items stored in the table.
     */
    streamSpecification?: pulumi.Input<inputs.dynamodb.TableStreamSpecificationArgs>;
    /**
     * The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``.
     */
    tableClass?: pulumi.Input<string>;
    /**
     * A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
     *   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    tableName?: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     *  For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * Specifies the Time to Live (TTL) settings for the table.
     *   For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
     */
    timeToLiveSpecification?: pulumi.Input<inputs.dynamodb.TableTimeToLiveSpecificationArgs>;
    /**
     * Represents the warm throughput (in read units per second and write units per second) for creating a table.
     */
    warmThroughput?: pulumi.Input<inputs.dynamodb.TableWarmThroughputArgs>;
}
