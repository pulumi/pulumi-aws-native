// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Retrieval service configuration for Retrieval node
    /// </summary>
    [OutputType]
    public sealed class FlowVersionRetrievalFlowNodeServiceConfigurationProperties
    {
        public readonly Outputs.FlowVersionRetrievalFlowNodeS3Configuration? S3;

        [OutputConstructor]
        private FlowVersionRetrievalFlowNodeServiceConfigurationProperties(Outputs.FlowVersionRetrievalFlowNodeS3Configuration? s3)
        {
            S3 = s3;
        }
    }
}