// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Ground truth input provided in S3 
    /// </summary>
    [OutputType]
    public sealed class ModelQualityJobDefinitionMonitoringGroundTruthS3Input
    {
        /// <summary>
        /// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
        /// </summary>
        public readonly string S3Uri;

        [OutputConstructor]
        private ModelQualityJobDefinitionMonitoringGroundTruthS3Input(string s3Uri)
        {
            S3Uri = s3Uri;
        }
    }
}