// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Inputs
{

    public sealed class ProjectS3DestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket in which Evidently stores evaluation events.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// The bucket prefix in which Evidently stores evaluation events.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public ProjectS3DestinationArgs()
        {
        }
        public static new ProjectS3DestinationArgs Empty => new ProjectS3DestinationArgs();
    }
}
