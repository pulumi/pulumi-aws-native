// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Outputs
{

    [OutputType]
    public sealed class EndpointKafkaSettings
    {
        public readonly string? Broker;
        public readonly bool? IncludeControlDetails;
        public readonly bool? IncludeNullAndEmpty;
        public readonly bool? IncludePartitionValue;
        public readonly bool? IncludeTableAlterOperations;
        public readonly bool? IncludeTransactionDetails;
        public readonly string? MessageFormat;
        public readonly int? MessageMaxBytes;
        public readonly bool? NoHexPrefix;
        public readonly bool? PartitionIncludeSchemaTable;
        public readonly string? SaslPassword;
        public readonly string? SaslUserName;
        public readonly string? SecurityProtocol;
        public readonly string? SslCaCertificateArn;
        public readonly string? SslClientCertificateArn;
        public readonly string? SslClientKeyArn;
        public readonly string? SslClientKeyPassword;
        public readonly string? Topic;

        [OutputConstructor]
        private EndpointKafkaSettings(
            string? broker,

            bool? includeControlDetails,

            bool? includeNullAndEmpty,

            bool? includePartitionValue,

            bool? includeTableAlterOperations,

            bool? includeTransactionDetails,

            string? messageFormat,

            int? messageMaxBytes,

            bool? noHexPrefix,

            bool? partitionIncludeSchemaTable,

            string? saslPassword,

            string? saslUserName,

            string? securityProtocol,

            string? sslCaCertificateArn,

            string? sslClientCertificateArn,

            string? sslClientKeyArn,

            string? sslClientKeyPassword,

            string? topic)
        {
            Broker = broker;
            IncludeControlDetails = includeControlDetails;
            IncludeNullAndEmpty = includeNullAndEmpty;
            IncludePartitionValue = includePartitionValue;
            IncludeTableAlterOperations = includeTableAlterOperations;
            IncludeTransactionDetails = includeTransactionDetails;
            MessageFormat = messageFormat;
            MessageMaxBytes = messageMaxBytes;
            NoHexPrefix = noHexPrefix;
            PartitionIncludeSchemaTable = partitionIncludeSchemaTable;
            SaslPassword = saslPassword;
            SaslUserName = saslUserName;
            SecurityProtocol = securityProtocol;
            SslCaCertificateArn = sslCaCertificateArn;
            SslClientCertificateArn = sslClientCertificateArn;
            SslClientKeyArn = sslClientKeyArn;
            SslClientKeyPassword = sslClientKeyPassword;
            Topic = topic;
        }
    }
}