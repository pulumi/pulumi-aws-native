// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Inputs
{

    /// <summary>
    /// The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.
    /// </summary>
    public sealed class LocationS3s3ConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the IAM role of the Amazon S3 bucket.
        /// </summary>
        [Input("bucketAccessRoleArn", required: true)]
        public Input<string> BucketAccessRoleArn { get; set; } = null!;

        public LocationS3s3ConfigArgs()
        {
        }
        public static new LocationS3s3ConfigArgs Empty => new LocationS3s3ConfigArgs();
    }
}
