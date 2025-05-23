// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    /// <summary>
    /// Specifies data stores to crawl.
    /// </summary>
    [OutputType]
    public sealed class CrawlerTargets
    {
        /// <summary>
        /// Specifies AWS Glue Data Catalog targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerCatalogTarget> CatalogTargets;
        /// <summary>
        /// Specifies an array of Delta data store targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerDeltaTarget> DeltaTargets;
        /// <summary>
        /// Specifies Amazon DynamoDB targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerDynamoDbTarget> DynamoDbTargets;
        /// <summary>
        /// Specifies Apache Hudi data store targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerHudiTarget> HudiTargets;
        /// <summary>
        /// Specifies Apache Iceberg data store targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerIcebergTarget> IcebergTargets;
        /// <summary>
        /// Specifies JDBC targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerJdbcTarget> JdbcTargets;
        /// <summary>
        /// A list of Mongo DB targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerMongoDbTarget> MongoDbTargets;
        /// <summary>
        /// Specifies Amazon Simple Storage Service (Amazon S3) targets.
        /// </summary>
        public readonly ImmutableArray<Outputs.CrawlerS3Target> S3Targets;

        [OutputConstructor]
        private CrawlerTargets(
            ImmutableArray<Outputs.CrawlerCatalogTarget> catalogTargets,

            ImmutableArray<Outputs.CrawlerDeltaTarget> deltaTargets,

            ImmutableArray<Outputs.CrawlerDynamoDbTarget> dynamoDbTargets,

            ImmutableArray<Outputs.CrawlerHudiTarget> hudiTargets,

            ImmutableArray<Outputs.CrawlerIcebergTarget> icebergTargets,

            ImmutableArray<Outputs.CrawlerJdbcTarget> jdbcTargets,

            ImmutableArray<Outputs.CrawlerMongoDbTarget> mongoDbTargets,

            ImmutableArray<Outputs.CrawlerS3Target> s3Targets)
        {
            CatalogTargets = catalogTargets;
            DeltaTargets = deltaTargets;
            DynamoDbTargets = dynamoDbTargets;
            HudiTargets = hudiTargets;
            IcebergTargets = icebergTargets;
            JdbcTargets = jdbcTargets;
            MongoDbTargets = mongoDbTargets;
            S3Targets = s3Targets;
        }
    }
}
