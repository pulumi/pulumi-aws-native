// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowErrorHandlingConfig
    {
        /// <summary>
        /// Specifies the name of the Amazon S3 bucket.
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// Specifies the Amazon S3 bucket prefix.
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// Specifies if the flow should fail after the first instance of a failure when attempting to place data in the destination.
        /// </summary>
        public readonly bool? FailOnFirstError;

        [OutputConstructor]
        private FlowErrorHandlingConfig(
            string? bucketName,

            string? bucketPrefix,

            bool? failOnFirstError)
        {
            BucketName = bucketName;
            BucketPrefix = bucketPrefix;
            FailOnFirstError = failOnFirstError;
        }
    }
}
