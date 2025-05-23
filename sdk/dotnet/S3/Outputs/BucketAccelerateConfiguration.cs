// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    /// <summary>
    /// Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.
    /// </summary>
    [OutputType]
    public sealed class BucketAccelerateConfiguration
    {
        /// <summary>
        /// Specifies the transfer acceleration status of the bucket.
        /// </summary>
        public readonly Pulumi.AwsNative.S3.BucketAccelerateConfigurationAccelerationStatus AccelerationStatus;

        [OutputConstructor]
        private BucketAccelerateConfiguration(Pulumi.AwsNative.S3.BucketAccelerateConfigurationAccelerationStatus accelerationStatus)
        {
            AccelerationStatus = accelerationStatus;
        }
    }
}
