// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Backup.Outputs
{

    /// <summary>
    /// A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports.
    /// </summary>
    [OutputType]
    public sealed class ReportDeliveryChannelProperties
    {
        /// <summary>
        /// A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
        /// </summary>
        public readonly ImmutableArray<string> Formats;
        /// <summary>
        /// The unique name of the S3 bucket that receives your reports.
        /// </summary>
        public readonly string S3BucketName;
        /// <summary>
        /// The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.
        /// </summary>
        public readonly string? S3KeyPrefix;

        [OutputConstructor]
        private ReportDeliveryChannelProperties(
            ImmutableArray<string> formats,

            string s3BucketName,

            string? s3KeyPrefix)
        {
            Formats = formats;
            S3BucketName = s3BucketName;
            S3KeyPrefix = s3KeyPrefix;
        }
    }
}