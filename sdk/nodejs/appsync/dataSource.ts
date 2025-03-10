// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppSync::DataSource
 *
 * ## Example Usage
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const config = new pulumi.Config();
 * const graphQlApiId = config.require("graphQlApiId");
 * const dataSourceName = config.require("dataSourceName");
 * const dataSourceDescription = config.require("dataSourceDescription");
 * const serviceRoleArn = config.require("serviceRoleArn");
 * const lambdaFunctionArn = config.require("lambdaFunctionArn");
 * const dataSource = new aws_native.appsync.DataSource("dataSource", {
 *     apiId: graphQlApiId,
 *     name: dataSourceName,
 *     description: dataSourceDescription,
 *     type: "AWS_LAMBDA",
 *     serviceRoleArn: serviceRoleArn,
 *     lambdaConfig: {
 *         lambdaFunctionArn: lambdaFunctionArn,
 *     },
 * });
 *
 * ```
 * ### Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws_native from "@pulumi/aws-native";
 *
 * const config = new pulumi.Config();
 * const graphQlApiId = config.require("graphQlApiId");
 * const dataSourceName = config.require("dataSourceName");
 * const dataSourceDescription = config.require("dataSourceDescription");
 * const serviceRoleArn = config.require("serviceRoleArn");
 * const lambdaFunctionArn = config.require("lambdaFunctionArn");
 * const dataSource = new aws_native.appsync.DataSource("dataSource", {
 *     apiId: graphQlApiId,
 *     name: dataSourceName,
 *     description: dataSourceDescription,
 *     type: "AWS_LAMBDA",
 *     serviceRoleArn: serviceRoleArn,
 *     lambdaConfig: {
 *         lambdaFunctionArn: lambdaFunctionArn,
 *     },
 * });
 *
 * ```
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appsync:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * Unique AWS AppSync GraphQL API identifier where this data source will be created.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the API key, such as arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename.
     */
    public /*out*/ readonly dataSourceArn!: pulumi.Output<string>;
    /**
     * The description of the data source.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * AWS Region and TableName for an Amazon DynamoDB table in your account.
     */
    public readonly dynamoDbConfig!: pulumi.Output<outputs.appsync.DataSourceDynamoDbConfig | undefined>;
    /**
     * AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
     * As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.
     */
    public readonly elasticsearchConfig!: pulumi.Output<outputs.appsync.DataSourceElasticsearchConfig | undefined>;
    /**
     * ARN for the EventBridge bus.
     */
    public readonly eventBridgeConfig!: pulumi.Output<outputs.appsync.DataSourceEventBridgeConfig | undefined>;
    /**
     * Endpoints for an HTTP data source.
     */
    public readonly httpConfig!: pulumi.Output<outputs.appsync.DataSourceHttpConfig | undefined>;
    /**
     * An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account.
     */
    public readonly lambdaConfig!: pulumi.Output<outputs.appsync.DataSourceLambdaConfig | undefined>;
    /**
     * Enables or disables enhanced data source metrics for specified data sources. Note that `MetricsConfig` won't be used unless the `dataSourceLevelMetricsBehavior` value is set to `PER_DATA_SOURCE_METRICS` . If the `dataSourceLevelMetricsBehavior` is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig` will be ignored. However, you can still set its value.
     *
     * `MetricsConfig` can be `ENABLED` or `DISABLED` .
     */
    public readonly metricsConfig!: pulumi.Output<enums.appsync.DataSourceMetricsConfig | undefined>;
    /**
     * Friendly name for you to identify your AppSync data source after creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
     */
    public readonly openSearchServiceConfig!: pulumi.Output<outputs.appsync.DataSourceOpenSearchServiceConfig | undefined>;
    /**
     * Relational Database configuration of the relational database data source.
     */
    public readonly relationalDatabaseConfig!: pulumi.Output<outputs.appsync.DataSourceRelationalDatabaseConfig | undefined>;
    /**
     * The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source.
     */
    public readonly serviceRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The type of the data source.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamoDbConfig"] = args ? args.dynamoDbConfig : undefined;
            resourceInputs["elasticsearchConfig"] = args ? args.elasticsearchConfig : undefined;
            resourceInputs["eventBridgeConfig"] = args ? args.eventBridgeConfig : undefined;
            resourceInputs["httpConfig"] = args ? args.httpConfig : undefined;
            resourceInputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            resourceInputs["metricsConfig"] = args ? args.metricsConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openSearchServiceConfig"] = args ? args.openSearchServiceConfig : undefined;
            resourceInputs["relationalDatabaseConfig"] = args ? args.relationalDatabaseConfig : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["dataSourceArn"] = undefined /*out*/;
        } else {
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["dataSourceArn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dynamoDbConfig"] = undefined /*out*/;
            resourceInputs["elasticsearchConfig"] = undefined /*out*/;
            resourceInputs["eventBridgeConfig"] = undefined /*out*/;
            resourceInputs["httpConfig"] = undefined /*out*/;
            resourceInputs["lambdaConfig"] = undefined /*out*/;
            resourceInputs["metricsConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["openSearchServiceConfig"] = undefined /*out*/;
            resourceInputs["relationalDatabaseConfig"] = undefined /*out*/;
            resourceInputs["serviceRoleArn"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["apiId", "name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * Unique AWS AppSync GraphQL API identifier where this data source will be created.
     */
    apiId: pulumi.Input<string>;
    /**
     * The description of the data source.
     */
    description?: pulumi.Input<string>;
    /**
     * AWS Region and TableName for an Amazon DynamoDB table in your account.
     */
    dynamoDbConfig?: pulumi.Input<inputs.appsync.DataSourceDynamoDbConfigArgs>;
    /**
     * AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
     * As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.
     */
    elasticsearchConfig?: pulumi.Input<inputs.appsync.DataSourceElasticsearchConfigArgs>;
    /**
     * ARN for the EventBridge bus.
     */
    eventBridgeConfig?: pulumi.Input<inputs.appsync.DataSourceEventBridgeConfigArgs>;
    /**
     * Endpoints for an HTTP data source.
     */
    httpConfig?: pulumi.Input<inputs.appsync.DataSourceHttpConfigArgs>;
    /**
     * An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account.
     */
    lambdaConfig?: pulumi.Input<inputs.appsync.DataSourceLambdaConfigArgs>;
    /**
     * Enables or disables enhanced data source metrics for specified data sources. Note that `MetricsConfig` won't be used unless the `dataSourceLevelMetricsBehavior` value is set to `PER_DATA_SOURCE_METRICS` . If the `dataSourceLevelMetricsBehavior` is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig` will be ignored. However, you can still set its value.
     *
     * `MetricsConfig` can be `ENABLED` or `DISABLED` .
     */
    metricsConfig?: pulumi.Input<enums.appsync.DataSourceMetricsConfig>;
    /**
     * Friendly name for you to identify your AppSync data source after creation.
     */
    name?: pulumi.Input<string>;
    /**
     * AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
     */
    openSearchServiceConfig?: pulumi.Input<inputs.appsync.DataSourceOpenSearchServiceConfigArgs>;
    /**
     * Relational Database configuration of the relational database data source.
     */
    relationalDatabaseConfig?: pulumi.Input<inputs.appsync.DataSourceRelationalDatabaseConfigArgs>;
    /**
     * The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * The type of the data source.
     */
    type: pulumi.Input<string>;
}
