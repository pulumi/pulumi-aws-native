// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The Amazon S3 location and configuration for storing inference request and response data.
    /// </summary>
    [OutputType]
    public sealed class InferenceExperimentDataStorageConfig
    {
        /// <summary>
        /// Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.
        /// </summary>
        public readonly Outputs.InferenceExperimentCaptureContentTypeHeader? ContentType;
        /// <summary>
        /// The Amazon S3 bucket where the inference request and response data is stored.
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.
        /// </summary>
        public readonly string? KmsKey;

        [OutputConstructor]
        private InferenceExperimentDataStorageConfig(
            Outputs.InferenceExperimentCaptureContentTypeHeader? contentType,

            string destination,

            string? kmsKey)
        {
            ContentType = contentType;
            Destination = destination;
            KmsKey = kmsKey;
        }
    }
}
