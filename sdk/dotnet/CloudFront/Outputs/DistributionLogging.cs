// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// A complex type that specifies whether access logs are written for the distribution.
    ///   If you already enabled standard logging (legacy) and you want to enable standard logging (v2) to send your access logs to Amazon S3, we recommend that you specify a *different* Amazon S3 bucket or use a *separate path* in the same bucket (for example, use a log prefix or partitioning). This helps you keep track of which log files are associated with which logging subscription and prevents log files from overwriting each other. For more information, see [Standard logging (access logs)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide*.
    /// </summary>
    [OutputType]
    public sealed class DistributionLogging
    {
        /// <summary>
        /// The Amazon S3 bucket to store the access logs in, for example, ``amzn-s3-demo-bucket.s3.amazonaws.com``.
        /// </summary>
        public readonly string? Bucket;
        /// <summary>
        /// Specifies whether you want CloudFront to include cookies in access logs, specify ``true`` for ``IncludeCookies``. If you choose to include cookies in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for this distribution. If you don't want to include cookies when you create a distribution or if you want to disable include cookies for an existing distribution, specify ``false`` for ``IncludeCookies``.
        /// </summary>
        public readonly bool? IncludeCookies;
        /// <summary>
        /// An optional string that you want CloudFront to prefix to the access log ``filenames`` for this distribution, for example, ``myprefix/``. If you want to enable logging, but you don't want to specify a prefix, you still must include an empty ``Prefix`` element in the ``Logging`` element.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private DistributionLogging(
            string? bucket,

            bool? includeCookies,

            string? prefix)
        {
            Bucket = bucket;
            IncludeCookies = includeCookies;
            Prefix = prefix;
        }
    }
}
