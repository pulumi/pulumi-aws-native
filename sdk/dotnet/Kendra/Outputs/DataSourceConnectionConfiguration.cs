// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceConnectionConfiguration
    {
        public readonly string DatabaseHost;
        public readonly string DatabaseName;
        public readonly int DatabasePort;
        public readonly string SecretArn;
        public readonly string TableName;

        [OutputConstructor]
        private DataSourceConnectionConfiguration(
            string databaseHost,

            string databaseName,

            int databasePort,

            string secretArn,

            string tableName)
        {
            DatabaseHost = databaseHost;
            DatabaseName = databaseName;
            DatabasePort = databasePort;
            SecretArn = secretArn;
            TableName = tableName;
        }
    }
}