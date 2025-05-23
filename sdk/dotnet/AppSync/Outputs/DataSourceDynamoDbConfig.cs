// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class DataSourceDynamoDbConfig
    {
        /// <summary>
        /// The AWS Region.
        /// </summary>
        public readonly string AwsRegion;
        /// <summary>
        /// The DeltaSyncConfig for a versioned datasource.
        /// </summary>
        public readonly Outputs.DataSourceDeltaSyncConfig? DeltaSyncConfig;
        /// <summary>
        /// The table name.
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// Set to TRUE to use AWS Identity and Access Management with this data source.
        /// </summary>
        public readonly bool? UseCallerCredentials;
        /// <summary>
        /// Set to TRUE to use Conflict Detection and Resolution with this data source.
        /// </summary>
        public readonly bool? Versioned;

        [OutputConstructor]
        private DataSourceDynamoDbConfig(
            string awsRegion,

            Outputs.DataSourceDeltaSyncConfig? deltaSyncConfig,

            string tableName,

            bool? useCallerCredentials,

            bool? versioned)
        {
            AwsRegion = awsRegion;
            DeltaSyncConfig = deltaSyncConfig;
            TableName = tableName;
            UseCallerCredentials = useCallerCredentials;
            Versioned = versioned;
        }
    }
}
