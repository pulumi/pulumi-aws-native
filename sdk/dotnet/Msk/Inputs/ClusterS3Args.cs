// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterS3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the S3 bucket that is the destination for broker logs.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Specifies whether broker logs get sent to the specified Amazon S3 destination.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The S3 prefix that is the destination for broker logs.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public ClusterS3Args()
        {
        }
        public static new ClusterS3Args Empty => new ClusterS3Args();
    }
}
