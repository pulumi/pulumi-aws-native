// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceDatabaseConfiguration
    {
        /// <summary>
        /// Information about the database column that provides information for user context filtering.
        /// </summary>
        public readonly Outputs.DataSourceAclConfiguration? AclConfiguration;
        /// <summary>
        /// Information about where the index should get the document information from the database.
        /// </summary>
        public readonly Outputs.DataSourceColumnConfiguration ColumnConfiguration;
        /// <summary>
        /// Configuration information that's required to connect to a database.
        /// </summary>
        public readonly Outputs.DataSourceConnectionConfiguration ConnectionConfiguration;
        /// <summary>
        /// The type of database engine that runs the database.
        /// </summary>
        public readonly Pulumi.AwsNative.Kendra.DataSourceDatabaseEngineType DatabaseEngineType;
        /// <summary>
        /// Provides information about how Amazon Kendra uses quote marks around SQL identifiers when querying a database data source.
        /// </summary>
        public readonly Outputs.DataSourceSqlConfiguration? SqlConfiguration;
        /// <summary>
        /// Provides information for connecting to an Amazon VPC.
        /// </summary>
        public readonly Outputs.DataSourceVpcConfiguration? VpcConfiguration;

        [OutputConstructor]
        private DataSourceDatabaseConfiguration(
            Outputs.DataSourceAclConfiguration? aclConfiguration,

            Outputs.DataSourceColumnConfiguration columnConfiguration,

            Outputs.DataSourceConnectionConfiguration connectionConfiguration,

            Pulumi.AwsNative.Kendra.DataSourceDatabaseEngineType databaseEngineType,

            Outputs.DataSourceSqlConfiguration? sqlConfiguration,

            Outputs.DataSourceVpcConfiguration? vpcConfiguration)
        {
            AclConfiguration = aclConfiguration;
            ColumnConfiguration = columnConfiguration;
            ConnectionConfiguration = connectionConfiguration;
            DatabaseEngineType = databaseEngineType;
            SqlConfiguration = sqlConfiguration;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
