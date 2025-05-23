// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateWordCloudFieldWells
    {
        /// <summary>
        /// The aggregated field wells of a word cloud.
        /// </summary>
        public readonly Outputs.TemplateWordCloudAggregatedFieldWells? WordCloudAggregatedFieldWells;

        [OutputConstructor]
        private TemplateWordCloudFieldWells(Outputs.TemplateWordCloudAggregatedFieldWells? wordCloudAggregatedFieldWells)
        {
            WordCloudAggregatedFieldWells = wordCloudAggregatedFieldWells;
        }
    }
}
