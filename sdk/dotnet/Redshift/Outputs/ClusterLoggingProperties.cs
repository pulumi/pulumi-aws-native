// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Outputs
{

    [OutputType]
    public sealed class ClusterLoggingProperties
    {
        /// <summary>
        /// The name of an existing S3 bucket where the log files are to be stored.
        /// 
        /// Constraints:
        /// 
        /// - Must be in the same region as the cluster
        /// - The cluster must have read bucket and put object permissions
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// The log destination type. An enum with possible values of `s3` and `cloudwatch` .
        /// </summary>
        public readonly string? LogDestinationType;
        /// <summary>
        /// The collection of exported log types. Possible values are `connectionlog` , `useractivitylog` , and `userlog` .
        /// </summary>
        public readonly ImmutableArray<string> LogExports;
        /// <summary>
        /// The prefix applied to the log file names.
        /// 
        /// Valid characters are any letter from any language, any whitespace character, any numeric character, and the following characters: underscore ( `_` ), period ( `.` ), colon ( `:` ), slash ( `/` ), equal ( `=` ), plus ( `+` ), backslash ( `\` ), hyphen ( `-` ), at symbol ( `@` ).
        /// </summary>
        public readonly string? S3KeyPrefix;

        [OutputConstructor]
        private ClusterLoggingProperties(
            string? bucketName,

            string? logDestinationType,

            ImmutableArray<string> logExports,

            string? s3KeyPrefix)
        {
            BucketName = bucketName;
            LogDestinationType = logDestinationType;
            LogExports = logExports;
            S3KeyPrefix = s3KeyPrefix;
        }
    }
}
