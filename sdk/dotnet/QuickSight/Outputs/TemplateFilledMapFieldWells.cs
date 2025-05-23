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
    public sealed class TemplateFilledMapFieldWells
    {
        /// <summary>
        /// The aggregated field well of the filled map.
        /// </summary>
        public readonly Outputs.TemplateFilledMapAggregatedFieldWells? FilledMapAggregatedFieldWells;

        [OutputConstructor]
        private TemplateFilledMapFieldWells(Outputs.TemplateFilledMapAggregatedFieldWells? filledMapAggregatedFieldWells)
        {
            FilledMapAggregatedFieldWells = filledMapAggregatedFieldWells;
        }
    }
}
