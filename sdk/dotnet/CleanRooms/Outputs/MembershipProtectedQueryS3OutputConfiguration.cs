// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Outputs
{

    [OutputType]
    public sealed class MembershipProtectedQueryS3OutputConfiguration
    {
        /// <summary>
        /// The S3 bucket to unload the protected query results.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The S3 prefix to unload the protected query results.
        /// </summary>
        public readonly string? KeyPrefix;
        /// <summary>
        /// Intended file format of the result.
        /// </summary>
        public readonly Pulumi.AwsNative.CleanRooms.MembershipResultFormat ResultFormat;
        /// <summary>
        /// Indicates whether files should be output as a single file ( `TRUE` ) or output as multiple files ( `FALSE` ). This parameter is only supported for analyses with the Spark analytics engine.
        /// </summary>
        public readonly bool? SingleFileOutput;

        [OutputConstructor]
        private MembershipProtectedQueryS3OutputConfiguration(
            string bucket,

            string? keyPrefix,

            Pulumi.AwsNative.CleanRooms.MembershipResultFormat resultFormat,

            bool? singleFileOutput)
        {
            Bucket = bucket;
            KeyPrefix = keyPrefix;
            ResultFormat = resultFormat;
            SingleFileOutput = singleFileOutput;
        }
    }
}
