// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The lifecycle configuration for a SageMaker HyperPod cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterLifeCycleConfig
    {
        /// <summary>
        /// The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.
        /// </summary>
        public readonly string OnCreate;
        /// <summary>
        /// An Amazon S3 bucket path where your lifecycle scripts are stored.
        /// </summary>
        public readonly string SourceS3Uri;

        [OutputConstructor]
        private ClusterLifeCycleConfig(
            string onCreate,

            string sourceS3Uri)
        {
            OnCreate = onCreate;
            SourceS3Uri = sourceS3Uri;
        }
    }
}