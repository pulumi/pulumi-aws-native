// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Represents a File Source Object.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageFileSource
    {
        /// <summary>
        /// The digest of the file source.
        /// </summary>
        public readonly string? ContentDigest;
        /// <summary>
        /// The type of content stored in the file source.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// The Amazon S3 URI for the file source.
        /// </summary>
        public readonly string S3Uri;

        [OutputConstructor]
        private ModelPackageFileSource(
            string? contentDigest,

            string? contentType,

            string s3Uri)
        {
            ContentDigest = contentDigest;
            ContentType = contentType;
            S3Uri = s3Uri;
        }
    }
}