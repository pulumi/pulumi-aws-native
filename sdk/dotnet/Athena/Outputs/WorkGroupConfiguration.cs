// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Athena.Outputs
{

    [OutputType]
    public sealed class WorkGroupConfiguration
    {
        /// <summary>
        /// Specifies a user defined JSON string that is passed to the session engine.
        /// </summary>
        public readonly string? AdditionalConfiguration;
        /// <summary>
        /// The upper limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan. No default is defined.
        /// 
        /// &gt; This property currently supports integer types. Support for long values is planned.
        /// </summary>
        public readonly int? BytesScannedCutoffPerQuery;
        /// <summary>
        /// Specifies the KMS key that is used to encrypt the user's data stores in Athena. This setting does not apply to Athena SQL workgroups.
        /// </summary>
        public readonly Outputs.WorkGroupCustomerContentEncryptionConfiguration? CustomerContentEncryptionConfiguration;
        /// <summary>
        /// If set to "true", the settings for the workgroup override client-side settings. If set to "false", client-side settings are used. For more information, see [Override client-side settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html) .
        /// </summary>
        public readonly bool? EnforceWorkGroupConfiguration;
        /// <summary>
        /// The engine version that all queries running on the workgroup use.
        /// </summary>
        public readonly Outputs.WorkGroupEngineVersion? EngineVersion;
        /// <summary>
        /// Role used to access user resources in an Athena for Apache Spark session. This property applies only to Spark-enabled workgroups in Athena.
        /// </summary>
        public readonly string? ExecutionRole;
        /// <summary>
        /// The configuration for storing results in Athena owned storage, which includes whether this feature is enabled; whether encryption configuration, if any, is used for encrypting query results.
        /// </summary>
        public readonly Outputs.WorkGroupManagedQueryResultsConfiguration? ManagedQueryResultsConfiguration;
        /// <summary>
        /// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
        /// </summary>
        public readonly bool? PublishCloudWatchMetricsEnabled;
        /// <summary>
        /// If set to `true` , allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to `false` , workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. The default is `false` . For more information about Requester Pays buckets, see [Requester Pays Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html) in the *Amazon Simple Storage Service Developer Guide* .
        /// </summary>
        public readonly bool? RequesterPaysEnabled;
        /// <summary>
        /// Specifies the location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. For more information, see [Work with query results and recent queries](https://docs.aws.amazon.com/athena/latest/ug/querying.html) .
        /// </summary>
        public readonly Outputs.WorkGroupResultConfiguration? ResultConfiguration;

        [OutputConstructor]
        private WorkGroupConfiguration(
            string? additionalConfiguration,

            int? bytesScannedCutoffPerQuery,

            Outputs.WorkGroupCustomerContentEncryptionConfiguration? customerContentEncryptionConfiguration,

            bool? enforceWorkGroupConfiguration,

            Outputs.WorkGroupEngineVersion? engineVersion,

            string? executionRole,

            Outputs.WorkGroupManagedQueryResultsConfiguration? managedQueryResultsConfiguration,

            bool? publishCloudWatchMetricsEnabled,

            bool? requesterPaysEnabled,

            Outputs.WorkGroupResultConfiguration? resultConfiguration)
        {
            AdditionalConfiguration = additionalConfiguration;
            BytesScannedCutoffPerQuery = bytesScannedCutoffPerQuery;
            CustomerContentEncryptionConfiguration = customerContentEncryptionConfiguration;
            EnforceWorkGroupConfiguration = enforceWorkGroupConfiguration;
            EngineVersion = engineVersion;
            ExecutionRole = executionRole;
            ManagedQueryResultsConfiguration = managedQueryResultsConfiguration;
            PublishCloudWatchMetricsEnabled = publishCloudWatchMetricsEnabled;
            RequesterPaysEnabled = requesterPaysEnabled;
            ResultConfiguration = resultConfiguration;
        }
    }
}
