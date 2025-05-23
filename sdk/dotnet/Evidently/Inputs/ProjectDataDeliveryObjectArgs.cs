// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Inputs
{

    /// <summary>
    /// Destinations for data.
    /// </summary>
    public sealed class ProjectDataDeliveryObjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the project stores evaluation events in CloudWatch Logs , this structure stores the log group name.
        /// </summary>
        [Input("logGroup")]
        public Input<string>? LogGroup { get; set; }

        /// <summary>
        /// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
        /// </summary>
        [Input("s3")]
        public Input<Inputs.ProjectS3DestinationArgs>? S3 { get; set; }

        public ProjectDataDeliveryObjectArgs()
        {
        }
        public static new ProjectDataDeliveryObjectArgs Empty => new ProjectDataDeliveryObjectArgs();
    }
}
