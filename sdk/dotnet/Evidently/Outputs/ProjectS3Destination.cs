// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Outputs
{

    [OutputType]
    public sealed class ProjectS3Destination
    {
        public readonly string BucketName;
        public readonly string? Prefix;

        [OutputConstructor]
        private ProjectS3Destination(
            string bucketName,

            string? prefix)
        {
            BucketName = bucketName;
            Prefix = prefix;
        }
    }
}