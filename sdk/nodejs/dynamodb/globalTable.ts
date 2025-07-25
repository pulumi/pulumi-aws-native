// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable
 */
export class GlobalTable extends pulumi.CustomResource {
    /**
     * Get an existing GlobalTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GlobalTable {
        return new GlobalTable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:dynamodb:GlobalTable';

    /**
     * Returns true if the given object is an instance of GlobalTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalTable.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` . The ARN returned is that of the replica in the region the stack is deployed to.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of attributes that describe the key schema for the global table and indexes.
     */
    public readonly attributeDefinitions!: pulumi.Output<outputs.dynamodb.GlobalTableAttributeDefinition[]>;
    /**
     * Specifies how you are charged for read and write throughput and how you manage capacity. Valid values are:
     *
     * - `PAY_PER_REQUEST`
     * - `PROVISIONED`
     *
     * All replicas in your global table will have the same billing mode. If you use `PROVISIONED` billing mode, you must provide an auto scaling configuration via the `WriteProvisionedThroughputSettings` property. The default value of this property is `PROVISIONED` .
     */
    public readonly billingMode!: pulumi.Output<string | undefined>;
    /**
     * Global secondary indexes to be created on the global table. You can create up to 20 global secondary indexes. Each replica in your global table will have the same global secondary index settings. You can only create or delete one global secondary index in a single stack operation.
     *
     * Since the backfilling of an index could take a long time, CloudFormation does not wait for the index to become active. If a stack operation rolls back, CloudFormation might not delete an index that has been added. In that case, you will need to delete the index manually.
     */
    public readonly globalSecondaryIndexes!: pulumi.Output<outputs.dynamodb.GlobalTableGlobalSecondaryIndex[] | undefined>;
    /**
     * The list of witnesses of the MRSC global table. Only one witness Region can be configured per MRSC global table.
     */
    public readonly globalTableWitnesses!: pulumi.Output<outputs.dynamodb.GlobalTableWitness[] | undefined>;
    /**
     * Specifies the attributes that make up the primary key for the table. The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
     */
    public readonly keySchema!: pulumi.Output<outputs.dynamodb.GlobalTableKeySchema[]>;
    /**
     * Local secondary indexes to be created on the table. You can create up to five local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes. Each replica in your global table will have the same local secondary index settings.
     */
    public readonly localSecondaryIndexes!: pulumi.Output<outputs.dynamodb.GlobalTableLocalSecondaryIndex[] | undefined>;
    /**
     * Specifies the consistency mode for a new global table.
     *
     * You can specify one of the following consistency modes:
     *
     * - `EVENTUAL` : Configures a new global table for multi-Region eventual consistency (MREC).
     * - `STRONG` : Configures a new global table for multi-Region strong consistency (MRSC).
     *
     * If you don't specify this field, the global table consistency mode defaults to `EVENTUAL` . For more information about global tables consistency modes, see [Consistency modes](https://docs.aws.amazon.com/V2globaltables_HowItWorks.html#V2globaltables_HowItWorks.consistency-modes) in DynamoDB developer guide.
     */
    public readonly multiRegionConsistency!: pulumi.Output<enums.dynamodb.GlobalTableMultiRegionConsistency | undefined>;
    /**
     * Specifies the list of replicas for your global table. The list must contain at least one element, the region where the stack defining the global table is deployed. For example, if you define your table in a stack deployed to us-east-1, you must have an entry in `Replicas` with the region us-east-1. You cannot remove the replica in the stack region.
     *
     * > Adding a replica might take a few minutes for an empty table, or up to several hours for large tables. If you want to add or remove a replica, we recommend submitting an `UpdateStack` operation containing only that change.
     * > 
     * > If you add or delete a replica during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new replica, you might need to manually delete the replica. 
     *
     * You can create a new global table with as many replicas as needed. You can add or remove replicas after table creation, but you can only add or remove a single replica in each update. For Multi-Region Strong Consistency (MRSC), you can add or remove up to 3 replicas, or 2 replicas plus a witness Region.
     */
    public readonly replicas!: pulumi.Output<outputs.dynamodb.GlobalTableReplicaSpecification[]>;
    /**
     * Specifies the settings to enable server-side encryption. These settings will be applied to all replicas. If you plan to use customer-managed KMS keys, you must provide a key for each replica using the `ReplicaSpecification.ReplicaSSESpecification` property.
     */
    public readonly sseSpecification!: pulumi.Output<outputs.dynamodb.GlobalTableSseSpecification | undefined>;
    /**
     * The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` . The `StreamArn` returned is that of the replica in the region the stack is deployed to.
     *
     * > You must specify the `StreamSpecification` property to use this attribute.
     */
    public /*out*/ readonly streamArn!: pulumi.Output<string>;
    /**
     * Specifies the streams settings on your global table. You must provide a value for this property if your global table contains more than one replica. You can only change the streams settings if your global table has only one replica. For Multi-Region Strong Consistency (MRSC), you do not need to provide a value for this property and can change the settings at any time.
     */
    public readonly streamSpecification!: pulumi.Output<outputs.dynamodb.GlobalTableStreamSpecification | undefined>;
    /**
     * Unique identifier for the table, such as `a123b456-01ab-23cd-123a-111222aaabbb` . The `TableId` returned is that of the replica in the region the stack is deployed to.
     */
    public /*out*/ readonly tableId!: pulumi.Output<string>;
    /**
     * A name for the global table. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
     *
     * > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    public readonly tableName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the time to live (TTL) settings for the table. This setting will be applied to all replicas.
     */
    public readonly timeToLiveSpecification!: pulumi.Output<outputs.dynamodb.GlobalTableTimeToLiveSpecification | undefined>;
    /**
     * Provides visibility into the number of read and write operations your table or secondary index can instantaneously support. The settings can be modified using the `UpdateTable` operation to meet the throughput requirements of an upcoming peak event.
     */
    public readonly warmThroughput!: pulumi.Output<outputs.dynamodb.GlobalTableWarmThroughput | undefined>;
    /**
     * Sets the write request settings for a global table or a global secondary index. You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
     */
    public readonly writeOnDemandThroughputSettings!: pulumi.Output<outputs.dynamodb.GlobalTableWriteOnDemandThroughputSettings | undefined>;
    /**
     * Specifies an auto scaling policy for write capacity. This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
     */
    public readonly writeProvisionedThroughputSettings!: pulumi.Output<outputs.dynamodb.GlobalTableWriteProvisionedThroughputSettings | undefined>;

    /**
     * Create a GlobalTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalTableArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.attributeDefinitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attributeDefinitions'");
            }
            if ((!args || args.keySchema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keySchema'");
            }
            if ((!args || args.replicas === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicas'");
            }
            resourceInputs["attributeDefinitions"] = args ? args.attributeDefinitions : undefined;
            resourceInputs["billingMode"] = args ? args.billingMode : undefined;
            resourceInputs["globalSecondaryIndexes"] = args ? args.globalSecondaryIndexes : undefined;
            resourceInputs["globalTableWitnesses"] = args ? args.globalTableWitnesses : undefined;
            resourceInputs["keySchema"] = args ? args.keySchema : undefined;
            resourceInputs["localSecondaryIndexes"] = args ? args.localSecondaryIndexes : undefined;
            resourceInputs["multiRegionConsistency"] = args ? args.multiRegionConsistency : undefined;
            resourceInputs["replicas"] = args ? args.replicas : undefined;
            resourceInputs["sseSpecification"] = args ? args.sseSpecification : undefined;
            resourceInputs["streamSpecification"] = args ? args.streamSpecification : undefined;
            resourceInputs["tableName"] = args ? args.tableName : undefined;
            resourceInputs["timeToLiveSpecification"] = args ? args.timeToLiveSpecification : undefined;
            resourceInputs["warmThroughput"] = args ? args.warmThroughput : undefined;
            resourceInputs["writeOnDemandThroughputSettings"] = args ? args.writeOnDemandThroughputSettings : undefined;
            resourceInputs["writeProvisionedThroughputSettings"] = args ? args.writeProvisionedThroughputSettings : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["streamArn"] = undefined /*out*/;
            resourceInputs["tableId"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["attributeDefinitions"] = undefined /*out*/;
            resourceInputs["billingMode"] = undefined /*out*/;
            resourceInputs["globalSecondaryIndexes"] = undefined /*out*/;
            resourceInputs["globalTableWitnesses"] = undefined /*out*/;
            resourceInputs["keySchema"] = undefined /*out*/;
            resourceInputs["localSecondaryIndexes"] = undefined /*out*/;
            resourceInputs["multiRegionConsistency"] = undefined /*out*/;
            resourceInputs["replicas"] = undefined /*out*/;
            resourceInputs["sseSpecification"] = undefined /*out*/;
            resourceInputs["streamArn"] = undefined /*out*/;
            resourceInputs["streamSpecification"] = undefined /*out*/;
            resourceInputs["tableId"] = undefined /*out*/;
            resourceInputs["tableName"] = undefined /*out*/;
            resourceInputs["timeToLiveSpecification"] = undefined /*out*/;
            resourceInputs["warmThroughput"] = undefined /*out*/;
            resourceInputs["writeOnDemandThroughputSettings"] = undefined /*out*/;
            resourceInputs["writeProvisionedThroughputSettings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["keySchema[*]", "localSecondaryIndexes[*]", "tableName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(GlobalTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a GlobalTable resource.
 */
export interface GlobalTableArgs {
    /**
     * A list of attributes that describe the key schema for the global table and indexes.
     */
    attributeDefinitions: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableAttributeDefinitionArgs>[]>;
    /**
     * Specifies how you are charged for read and write throughput and how you manage capacity. Valid values are:
     *
     * - `PAY_PER_REQUEST`
     * - `PROVISIONED`
     *
     * All replicas in your global table will have the same billing mode. If you use `PROVISIONED` billing mode, you must provide an auto scaling configuration via the `WriteProvisionedThroughputSettings` property. The default value of this property is `PROVISIONED` .
     */
    billingMode?: pulumi.Input<string>;
    /**
     * Global secondary indexes to be created on the global table. You can create up to 20 global secondary indexes. Each replica in your global table will have the same global secondary index settings. You can only create or delete one global secondary index in a single stack operation.
     *
     * Since the backfilling of an index could take a long time, CloudFormation does not wait for the index to become active. If a stack operation rolls back, CloudFormation might not delete an index that has been added. In that case, you will need to delete the index manually.
     */
    globalSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableGlobalSecondaryIndexArgs>[]>;
    /**
     * The list of witnesses of the MRSC global table. Only one witness Region can be configured per MRSC global table.
     */
    globalTableWitnesses?: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableWitnessArgs>[]>;
    /**
     * Specifies the attributes that make up the primary key for the table. The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
     */
    keySchema: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableKeySchemaArgs>[]>;
    /**
     * Local secondary indexes to be created on the table. You can create up to five local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes. Each replica in your global table will have the same local secondary index settings.
     */
    localSecondaryIndexes?: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableLocalSecondaryIndexArgs>[]>;
    /**
     * Specifies the consistency mode for a new global table.
     *
     * You can specify one of the following consistency modes:
     *
     * - `EVENTUAL` : Configures a new global table for multi-Region eventual consistency (MREC).
     * - `STRONG` : Configures a new global table for multi-Region strong consistency (MRSC).
     *
     * If you don't specify this field, the global table consistency mode defaults to `EVENTUAL` . For more information about global tables consistency modes, see [Consistency modes](https://docs.aws.amazon.com/V2globaltables_HowItWorks.html#V2globaltables_HowItWorks.consistency-modes) in DynamoDB developer guide.
     */
    multiRegionConsistency?: pulumi.Input<enums.dynamodb.GlobalTableMultiRegionConsistency>;
    /**
     * Specifies the list of replicas for your global table. The list must contain at least one element, the region where the stack defining the global table is deployed. For example, if you define your table in a stack deployed to us-east-1, you must have an entry in `Replicas` with the region us-east-1. You cannot remove the replica in the stack region.
     *
     * > Adding a replica might take a few minutes for an empty table, or up to several hours for large tables. If you want to add or remove a replica, we recommend submitting an `UpdateStack` operation containing only that change.
     * > 
     * > If you add or delete a replica during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new replica, you might need to manually delete the replica. 
     *
     * You can create a new global table with as many replicas as needed. You can add or remove replicas after table creation, but you can only add or remove a single replica in each update. For Multi-Region Strong Consistency (MRSC), you can add or remove up to 3 replicas, or 2 replicas plus a witness Region.
     */
    replicas: pulumi.Input<pulumi.Input<inputs.dynamodb.GlobalTableReplicaSpecificationArgs>[]>;
    /**
     * Specifies the settings to enable server-side encryption. These settings will be applied to all replicas. If you plan to use customer-managed KMS keys, you must provide a key for each replica using the `ReplicaSpecification.ReplicaSSESpecification` property.
     */
    sseSpecification?: pulumi.Input<inputs.dynamodb.GlobalTableSseSpecificationArgs>;
    /**
     * Specifies the streams settings on your global table. You must provide a value for this property if your global table contains more than one replica. You can only change the streams settings if your global table has only one replica. For Multi-Region Strong Consistency (MRSC), you do not need to provide a value for this property and can change the settings at any time.
     */
    streamSpecification?: pulumi.Input<inputs.dynamodb.GlobalTableStreamSpecificationArgs>;
    /**
     * A name for the global table. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
     *
     * > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
     */
    tableName?: pulumi.Input<string>;
    /**
     * Specifies the time to live (TTL) settings for the table. This setting will be applied to all replicas.
     */
    timeToLiveSpecification?: pulumi.Input<inputs.dynamodb.GlobalTableTimeToLiveSpecificationArgs>;
    /**
     * Provides visibility into the number of read and write operations your table or secondary index can instantaneously support. The settings can be modified using the `UpdateTable` operation to meet the throughput requirements of an upcoming peak event.
     */
    warmThroughput?: pulumi.Input<inputs.dynamodb.GlobalTableWarmThroughputArgs>;
    /**
     * Sets the write request settings for a global table or a global secondary index. You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
     */
    writeOnDemandThroughputSettings?: pulumi.Input<inputs.dynamodb.GlobalTableWriteOnDemandThroughputSettingsArgs>;
    /**
     * Specifies an auto scaling policy for write capacity. This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
     */
    writeProvisionedThroughputSettings?: pulumi.Input<inputs.dynamodb.GlobalTableWriteProvisionedThroughputSettingsArgs>;
}
