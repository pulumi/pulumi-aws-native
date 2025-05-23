// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    [OutputType]
    public sealed class DataAutomationProjectDocumentStandardExtraction
    {
        /// <summary>
        /// Whether to generate bounding boxes.
        /// </summary>
        public readonly Outputs.DataAutomationProjectDocumentBoundingBox BoundingBox;
        /// <summary>
        /// Which granularities to generate data for.
        /// </summary>
        public readonly Outputs.DataAutomationProjectDocumentExtractionGranularity Granularity;

        [OutputConstructor]
        private DataAutomationProjectDocumentStandardExtraction(
            Outputs.DataAutomationProjectDocumentBoundingBox boundingBox,

            Outputs.DataAutomationProjectDocumentExtractionGranularity granularity)
        {
            BoundingBox = boundingBox;
            Granularity = granularity;
        }
    }
}
