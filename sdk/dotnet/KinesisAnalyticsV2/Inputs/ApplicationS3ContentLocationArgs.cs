// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    /// <summary>
    /// The location of an application or a custom artifact.
    /// </summary>
    public sealed class ApplicationS3ContentLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the S3 bucket containing the application code.
        /// </summary>
        [Input("bucketARN", required: true)]
        public Input<string> BucketARN { get; set; } = null!;

        /// <summary>
        /// The file key for the object containing the application code.
        /// </summary>
        [Input("fileKey", required: true)]
        public Input<string> FileKey { get; set; } = null!;

        /// <summary>
        /// The version of the object containing the application code.
        /// </summary>
        [Input("objectVersion")]
        public Input<string>? ObjectVersion { get; set; }

        public ApplicationS3ContentLocationArgs()
        {
        }
        public static new ApplicationS3ContentLocationArgs Empty => new ApplicationS3ContentLocationArgs();
    }
}