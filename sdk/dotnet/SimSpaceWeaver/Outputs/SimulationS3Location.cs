// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SimSpaceWeaver.Outputs
{

    [OutputType]
    public sealed class SimulationS3Location
    {
        /// <summary>
        /// The Schema S3 bucket name.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// This is the schema S3 object key, which includes the full path of "folders" from the bucket root to the schema.
        /// </summary>
        public readonly string ObjectKey;

        [OutputConstructor]
        private SimulationS3Location(
            string bucketName,

            string objectKey)
        {
            BucketName = bucketName;
            ObjectKey = objectKey;
        }
    }
}