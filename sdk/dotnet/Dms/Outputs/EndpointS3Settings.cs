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
    public sealed class EndpointS3Settings
    {
        public readonly bool? AddColumnName;
        public readonly bool? AddTrailingPaddingCharacter;
        public readonly string? BucketFolder;
        public readonly string? BucketName;
        public readonly string? CannedAclForObjects;
        public readonly bool? CdcInsertsAndUpdates;
        public readonly bool? CdcInsertsOnly;
        public readonly int? CdcMaxBatchInterval;
        public readonly int? CdcMinFileSize;
        public readonly string? CdcPath;
        public readonly string? CompressionType;
        public readonly string? CsvDelimiter;
        public readonly string? CsvNoSupValue;
        public readonly string? CsvNullValue;
        public readonly string? CsvRowDelimiter;
        public readonly string? DataFormat;
        public readonly int? DataPageSize;
        public readonly string? DatePartitionDelimiter;
        public readonly bool? DatePartitionEnabled;
        public readonly string? DatePartitionSequence;
        public readonly string? DatePartitionTimezone;
        public readonly int? DictPageSizeLimit;
        public readonly bool? EnableStatistics;
        public readonly string? EncodingType;
        public readonly string? EncryptionMode;
        public readonly string? ExpectedBucketOwner;
        public readonly string? ExternalTableDefinition;
        public readonly bool? GlueCatalogGeneration;
        public readonly int? IgnoreHeaderRows;
        public readonly bool? IncludeOpForFullLoad;
        public readonly int? MaxFileSize;
        public readonly bool? ParquetTimestampInMillisecond;
        public readonly string? ParquetVersion;
        public readonly bool? PreserveTransactions;
        public readonly bool? Rfc4180;
        public readonly int? RowGroupLength;
        public readonly string? ServerSideEncryptionKmsKeyId;
        public readonly string? ServiceAccessRoleArn;
        public readonly string? TimestampColumnName;
        public readonly bool? UseCsvNoSupValue;
        public readonly bool? UseTaskStartTimeForFullLoadTimestamp;

        [OutputConstructor]
        private EndpointS3Settings(
            bool? addColumnName,

            bool? addTrailingPaddingCharacter,

            string? bucketFolder,

            string? bucketName,

            string? cannedAclForObjects,

            bool? cdcInsertsAndUpdates,

            bool? cdcInsertsOnly,

            int? cdcMaxBatchInterval,

            int? cdcMinFileSize,

            string? cdcPath,

            string? compressionType,

            string? csvDelimiter,

            string? csvNoSupValue,

            string? csvNullValue,

            string? csvRowDelimiter,

            string? dataFormat,

            int? dataPageSize,

            string? datePartitionDelimiter,

            bool? datePartitionEnabled,

            string? datePartitionSequence,

            string? datePartitionTimezone,

            int? dictPageSizeLimit,

            bool? enableStatistics,

            string? encodingType,

            string? encryptionMode,

            string? expectedBucketOwner,

            string? externalTableDefinition,

            bool? glueCatalogGeneration,

            int? ignoreHeaderRows,

            bool? includeOpForFullLoad,

            int? maxFileSize,

            bool? parquetTimestampInMillisecond,

            string? parquetVersion,

            bool? preserveTransactions,

            bool? rfc4180,

            int? rowGroupLength,

            string? serverSideEncryptionKmsKeyId,

            string? serviceAccessRoleArn,

            string? timestampColumnName,

            bool? useCsvNoSupValue,

            bool? useTaskStartTimeForFullLoadTimestamp)
        {
            AddColumnName = addColumnName;
            AddTrailingPaddingCharacter = addTrailingPaddingCharacter;
            BucketFolder = bucketFolder;
            BucketName = bucketName;
            CannedAclForObjects = cannedAclForObjects;
            CdcInsertsAndUpdates = cdcInsertsAndUpdates;
            CdcInsertsOnly = cdcInsertsOnly;
            CdcMaxBatchInterval = cdcMaxBatchInterval;
            CdcMinFileSize = cdcMinFileSize;
            CdcPath = cdcPath;
            CompressionType = compressionType;
            CsvDelimiter = csvDelimiter;
            CsvNoSupValue = csvNoSupValue;
            CsvNullValue = csvNullValue;
            CsvRowDelimiter = csvRowDelimiter;
            DataFormat = dataFormat;
            DataPageSize = dataPageSize;
            DatePartitionDelimiter = datePartitionDelimiter;
            DatePartitionEnabled = datePartitionEnabled;
            DatePartitionSequence = datePartitionSequence;
            DatePartitionTimezone = datePartitionTimezone;
            DictPageSizeLimit = dictPageSizeLimit;
            EnableStatistics = enableStatistics;
            EncodingType = encodingType;
            EncryptionMode = encryptionMode;
            ExpectedBucketOwner = expectedBucketOwner;
            ExternalTableDefinition = externalTableDefinition;
            GlueCatalogGeneration = glueCatalogGeneration;
            IgnoreHeaderRows = ignoreHeaderRows;
            IncludeOpForFullLoad = includeOpForFullLoad;
            MaxFileSize = maxFileSize;
            ParquetTimestampInMillisecond = parquetTimestampInMillisecond;
            ParquetVersion = parquetVersion;
            PreserveTransactions = preserveTransactions;
            Rfc4180 = rfc4180;
            RowGroupLength = rowGroupLength;
            ServerSideEncryptionKmsKeyId = serverSideEncryptionKmsKeyId;
            ServiceAccessRoleArn = serviceAccessRoleArn;
            TimestampColumnName = timestampColumnName;
            UseCsvNoSupValue = useCsvNoSupValue;
            UseTaskStartTimeForFullLoadTimestamp = useTaskStartTimeForFullLoadTimestamp;
        }
    }
}