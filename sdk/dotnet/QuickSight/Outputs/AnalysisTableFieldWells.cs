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
    public sealed class AnalysisTableFieldWells
    {
        /// <summary>
        /// The aggregated field well for the table.
        /// </summary>
        public readonly Outputs.AnalysisTableAggregatedFieldWells? TableAggregatedFieldWells;
        /// <summary>
        /// The unaggregated field well for the table.
        /// </summary>
        public readonly Outputs.AnalysisTableUnaggregatedFieldWells? TableUnaggregatedFieldWells;

        [OutputConstructor]
        private AnalysisTableFieldWells(
            Outputs.AnalysisTableAggregatedFieldWells? tableAggregatedFieldWells,

            Outputs.AnalysisTableUnaggregatedFieldWells? tableUnaggregatedFieldWells)
        {
            TableAggregatedFieldWells = tableAggregatedFieldWells;
            TableUnaggregatedFieldWells = tableUnaggregatedFieldWells;
        }
    }
}
