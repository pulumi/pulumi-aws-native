// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorRdsSourceConfig
    {
        /// <summary>
        /// The host name of the database.
        /// </summary>
        public readonly string DatabaseHost;
        /// <summary>
        /// The name of the RDS database.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The port number where the database can be accessed.
        /// </summary>
        public readonly int DatabasePort;
        /// <summary>
        /// A string identifying the database instance.
        /// </summary>
        public readonly string DbInstanceIdentifier;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the role.
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
        /// </summary>
        public readonly string SecretManagerArn;
        /// <summary>
        /// The name of the table in the database.
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// An object containing information about the Amazon Virtual Private Cloud (VPC) configuration.
        /// </summary>
        public readonly Outputs.AnomalyDetectorVpcConfiguration VpcConfiguration;

        [OutputConstructor]
        private AnomalyDetectorRdsSourceConfig(
            string databaseHost,

            string databaseName,

            int databasePort,

            string dbInstanceIdentifier,

            string roleArn,

            string secretManagerArn,

            string tableName,

            Outputs.AnomalyDetectorVpcConfiguration vpcConfiguration)
        {
            DatabaseHost = databaseHost;
            DatabaseName = databaseName;
            DatabasePort = databasePort;
            DbInstanceIdentifier = dbInstanceIdentifier;
            RoleArn = roleArn;
            SecretManagerArn = secretManagerArn;
            TableName = tableName;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
