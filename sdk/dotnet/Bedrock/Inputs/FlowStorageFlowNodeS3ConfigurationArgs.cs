// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// s3 storage configuration for storage node
    /// </summary>
    public sealed class FlowStorageFlowNodeS3ConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bucket name of an s3 that will be used for storage flow node configuration
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        public FlowStorageFlowNodeS3ConfigurationArgs()
        {
        }
        public static new FlowStorageFlowNodeS3ConfigurationArgs Empty => new FlowStorageFlowNodeS3ConfigurationArgs();
    }
}