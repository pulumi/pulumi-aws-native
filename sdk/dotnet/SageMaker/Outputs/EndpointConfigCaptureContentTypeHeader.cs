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
    public sealed class EndpointConfigCaptureContentTypeHeader
    {
        public readonly ImmutableArray<string> CsvContentTypes;
        public readonly ImmutableArray<string> JsonContentTypes;

        [OutputConstructor]
        private EndpointConfigCaptureContentTypeHeader(
            ImmutableArray<string> csvContentTypes,

            ImmutableArray<string> jsonContentTypes)
        {
            CsvContentTypes = csvContentTypes;
            JsonContentTypes = jsonContentTypes;
        }
    }
}