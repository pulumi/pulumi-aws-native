// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Represents a Metric Source Object.
    /// </summary>
    public sealed class ModelPackageMetricsSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The digest of the metric source.
        /// </summary>
        [Input("contentDigest")]
        public Input<string>? ContentDigest { get; set; }

        /// <summary>
        /// The type of content stored in the metric source.
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// The Amazon S3 URI for the metric source.
        /// </summary>
        [Input("s3Uri", required: true)]
        public Input<string> S3Uri { get; set; } = null!;

        public ModelPackageMetricsSourceArgs()
        {
        }
        public static new ModelPackageMetricsSourceArgs Empty => new ModelPackageMetricsSourceArgs();
    }
}
