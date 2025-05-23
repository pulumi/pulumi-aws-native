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
    public sealed class DataAutomationProjectVideoStandardExtraction
    {
        /// <summary>
        /// Settings for generating bounding boxes.
        /// </summary>
        public readonly Outputs.DataAutomationProjectVideoBoundingBox BoundingBox;
        /// <summary>
        /// Settings for generating categorical data.
        /// </summary>
        public readonly Outputs.DataAutomationProjectVideoExtractionCategory Category;

        [OutputConstructor]
        private DataAutomationProjectVideoStandardExtraction(
            Outputs.DataAutomationProjectVideoBoundingBox boundingBox,

            Outputs.DataAutomationProjectVideoExtractionCategory category)
        {
            BoundingBox = boundingBox;
            Category = category;
        }
    }
}
