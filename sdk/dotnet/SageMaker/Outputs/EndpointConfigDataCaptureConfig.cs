// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class EndpointConfigDataCaptureConfig
    {
        public readonly Outputs.EndpointConfigCaptureContentTypeHeader? CaptureContentTypeHeader;
        public readonly ImmutableArray<Outputs.EndpointConfigCaptureOption> CaptureOptions;
        public readonly string DestinationS3Uri;
        public readonly bool? EnableCapture;
        public readonly int InitialSamplingPercentage;
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private EndpointConfigDataCaptureConfig(
            Outputs.EndpointConfigCaptureContentTypeHeader? captureContentTypeHeader,

            ImmutableArray<Outputs.EndpointConfigCaptureOption> captureOptions,

            string destinationS3Uri,

            bool? enableCapture,

            int initialSamplingPercentage,

            string? kmsKeyId)
        {
            CaptureContentTypeHeader = captureContentTypeHeader;
            CaptureOptions = captureOptions;
            DestinationS3Uri = destinationS3Uri;
            EnableCapture = enableCapture;
            InitialSamplingPercentage = initialSamplingPercentage;
            KmsKeyId = kmsKeyId;
        }
    }
}