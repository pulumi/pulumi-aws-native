// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Outputs
{

    [OutputType]
    public sealed class ApplicationS3Location
    {
        public readonly string S3Bucket;
        public readonly string S3Key;

        [OutputConstructor]
        private ApplicationS3Location(
            string s3Bucket,

            string s3Key)
        {
            S3Bucket = s3Bucket;
            S3Key = s3Key;
        }
    }
}