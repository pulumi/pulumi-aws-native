// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// s3 Retrieval configuration for Retrieval node
    /// </summary>
    [OutputType]
    public sealed class FlowRetrievalFlowNodeS3Configuration
    {
        /// <summary>
        /// bucket name of an s3 that will be used for Retrieval flow node configuration
        /// </summary>
        public readonly string BucketName;

        [OutputConstructor]
        private FlowRetrievalFlowNodeS3Configuration(string bucketName)
        {
            BucketName = bucketName;
        }
    }
}
