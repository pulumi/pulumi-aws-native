// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// The Amazon S3 location
    /// </summary>
    public sealed class SoftwarePackageVersionS3LocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 bucket
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The S3 key
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The S3 version
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public SoftwarePackageVersionS3LocationArgs()
        {
        }
        public static new SoftwarePackageVersionS3LocationArgs Empty => new SoftwarePackageVersionS3LocationArgs();
    }
}
