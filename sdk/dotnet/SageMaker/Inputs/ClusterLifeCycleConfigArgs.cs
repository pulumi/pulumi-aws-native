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
    /// The lifecycle configuration for a SageMaker HyperPod cluster.
    /// </summary>
    public sealed class ClusterLifeCycleConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.
        /// </summary>
        [Input("onCreate", required: true)]
        public Input<string> OnCreate { get; set; } = null!;

        /// <summary>
        /// An Amazon S3 bucket path where your lifecycle scripts are stored.
        /// </summary>
        [Input("sourceS3Uri", required: true)]
        public Input<string> SourceS3Uri { get; set; } = null!;

        public ClusterLifeCycleConfigArgs()
        {
        }
        public static new ClusterLifeCycleConfigArgs Empty => new ClusterLifeCycleConfigArgs();
    }
}
