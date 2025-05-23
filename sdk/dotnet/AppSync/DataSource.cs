// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync
{
    /// <summary>
    /// Resource Type definition for AWS::AppSync::DataSource
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var graphQlApiId = config.Require("graphQlApiId");
    ///     var dataSourceName = config.Require("dataSourceName");
    ///     var dataSourceDescription = config.Require("dataSourceDescription");
    ///     var serviceRoleArn = config.Require("serviceRoleArn");
    ///     var lambdaFunctionArn = config.Require("lambdaFunctionArn");
    ///     var dataSource = new AwsNative.AppSync.DataSource("dataSource", new()
    ///     {
    ///         ApiId = graphQlApiId,
    ///         Name = dataSourceName,
    ///         Description = dataSourceDescription,
    ///         Type = "AWS_LAMBDA",
    ///         ServiceRoleArn = serviceRoleArn,
    ///         LambdaConfig = new AwsNative.AppSync.Inputs.DataSourceLambdaConfigArgs
    ///         {
    ///             LambdaFunctionArn = lambdaFunctionArn,
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var graphQlApiId = config.Require("graphQlApiId");
    ///     var dataSourceName = config.Require("dataSourceName");
    ///     var dataSourceDescription = config.Require("dataSourceDescription");
    ///     var serviceRoleArn = config.Require("serviceRoleArn");
    ///     var lambdaFunctionArn = config.Require("lambdaFunctionArn");
    ///     var dataSource = new AwsNative.AppSync.DataSource("dataSource", new()
    ///     {
    ///         ApiId = graphQlApiId,
    ///         Name = dataSourceName,
    ///         Description = dataSourceDescription,
    ///         Type = "AWS_LAMBDA",
    ///         ServiceRoleArn = serviceRoleArn,
    ///         LambdaConfig = new AwsNative.AppSync.Inputs.DataSourceLambdaConfigArgs
    ///         {
    ///             LambdaFunctionArn = lambdaFunctionArn,
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:appsync:DataSource")]
    public partial class DataSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique AWS AppSync GraphQL API identifier where this data source will be created.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the API key, such as arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename.
        /// </summary>
        [Output("dataSourceArn")]
        public Output<string> DataSourceArn { get; private set; } = null!;

        /// <summary>
        /// The description of the data source.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// AWS Region and TableName for an Amazon DynamoDB table in your account.
        /// </summary>
        [Output("dynamoDbConfig")]
        public Output<Outputs.DataSourceDynamoDbConfig?> DynamoDbConfig { get; private set; } = null!;

        /// <summary>
        /// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
        /// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.
        /// </summary>
        [Output("elasticsearchConfig")]
        public Output<Outputs.DataSourceElasticsearchConfig?> ElasticsearchConfig { get; private set; } = null!;

        /// <summary>
        /// ARN for the EventBridge bus.
        /// </summary>
        [Output("eventBridgeConfig")]
        public Output<Outputs.DataSourceEventBridgeConfig?> EventBridgeConfig { get; private set; } = null!;

        /// <summary>
        /// Endpoints for an HTTP data source.
        /// </summary>
        [Output("httpConfig")]
        public Output<Outputs.DataSourceHttpConfig?> HttpConfig { get; private set; } = null!;

        /// <summary>
        /// An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account.
        /// </summary>
        [Output("lambdaConfig")]
        public Output<Outputs.DataSourceLambdaConfig?> LambdaConfig { get; private set; } = null!;

        /// <summary>
        /// Enables or disables enhanced data source metrics for specified data sources. Note that `MetricsConfig` won't be used unless the `dataSourceLevelMetricsBehavior` value is set to `PER_DATA_SOURCE_METRICS` . If the `dataSourceLevelMetricsBehavior` is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig` will be ignored. However, you can still set its value.
        /// 
        /// `MetricsConfig` can be `ENABLED` or `DISABLED` .
        /// </summary>
        [Output("metricsConfig")]
        public Output<Pulumi.AwsNative.AppSync.DataSourceMetricsConfig?> MetricsConfig { get; private set; } = null!;

        /// <summary>
        /// Friendly name for you to identify your AppSync data source after creation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
        /// </summary>
        [Output("openSearchServiceConfig")]
        public Output<Outputs.DataSourceOpenSearchServiceConfig?> OpenSearchServiceConfig { get; private set; } = null!;

        /// <summary>
        /// Relational Database configuration of the relational database data source.
        /// </summary>
        [Output("relationalDatabaseConfig")]
        public Output<Outputs.DataSourceRelationalDatabaseConfig?> RelationalDatabaseConfig { get; private set; } = null!;

        /// <summary>
        /// The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source.
        /// </summary>
        [Output("serviceRoleArn")]
        public Output<string?> ServiceRoleArn { get; private set; } = null!;

        /// <summary>
        /// The type of the data source.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSource(string name, DataSourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DataSource", name, args ?? new DataSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appsync:DataSource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "apiId",
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataSource(name, id, options);
        }
    }

    public sealed class DataSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique AWS AppSync GraphQL API identifier where this data source will be created.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The description of the data source.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// AWS Region and TableName for an Amazon DynamoDB table in your account.
        /// </summary>
        [Input("dynamoDbConfig")]
        public Input<Inputs.DataSourceDynamoDbConfigArgs>? DynamoDbConfig { get; set; }

        /// <summary>
        /// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
        /// As of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.
        /// </summary>
        [Input("elasticsearchConfig")]
        public Input<Inputs.DataSourceElasticsearchConfigArgs>? ElasticsearchConfig { get; set; }

        /// <summary>
        /// ARN for the EventBridge bus.
        /// </summary>
        [Input("eventBridgeConfig")]
        public Input<Inputs.DataSourceEventBridgeConfigArgs>? EventBridgeConfig { get; set; }

        /// <summary>
        /// Endpoints for an HTTP data source.
        /// </summary>
        [Input("httpConfig")]
        public Input<Inputs.DataSourceHttpConfigArgs>? HttpConfig { get; set; }

        /// <summary>
        /// An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account.
        /// </summary>
        [Input("lambdaConfig")]
        public Input<Inputs.DataSourceLambdaConfigArgs>? LambdaConfig { get; set; }

        /// <summary>
        /// Enables or disables enhanced data source metrics for specified data sources. Note that `MetricsConfig` won't be used unless the `dataSourceLevelMetricsBehavior` value is set to `PER_DATA_SOURCE_METRICS` . If the `dataSourceLevelMetricsBehavior` is set to `FULL_REQUEST_DATA_SOURCE_METRICS` instead, `MetricsConfig` will be ignored. However, you can still set its value.
        /// 
        /// `MetricsConfig` can be `ENABLED` or `DISABLED` .
        /// </summary>
        [Input("metricsConfig")]
        public Input<Pulumi.AwsNative.AppSync.DataSourceMetricsConfig>? MetricsConfig { get; set; }

        /// <summary>
        /// Friendly name for you to identify your AppSync data source after creation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.
        /// </summary>
        [Input("openSearchServiceConfig")]
        public Input<Inputs.DataSourceOpenSearchServiceConfigArgs>? OpenSearchServiceConfig { get; set; }

        /// <summary>
        /// Relational Database configuration of the relational database data source.
        /// </summary>
        [Input("relationalDatabaseConfig")]
        public Input<Inputs.DataSourceRelationalDatabaseConfigArgs>? RelationalDatabaseConfig { get; set; }

        /// <summary>
        /// The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source.
        /// </summary>
        [Input("serviceRoleArn")]
        public Input<string>? ServiceRoleArn { get; set; }

        /// <summary>
        /// The type of the data source.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DataSourceArgs()
        {
        }
        public static new DataSourceArgs Empty => new DataSourceArgs();
    }
}
