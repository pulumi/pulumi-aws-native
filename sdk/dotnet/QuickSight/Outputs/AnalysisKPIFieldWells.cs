// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisKPIFieldWells
    {
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> TargetValues;
        public readonly ImmutableArray<Outputs.AnalysisDimensionField> TrendGroups;
        public readonly ImmutableArray<Outputs.AnalysisMeasureField> Values;

        [OutputConstructor]
        private AnalysisKPIFieldWells(
            ImmutableArray<Outputs.AnalysisMeasureField> targetValues,

            ImmutableArray<Outputs.AnalysisDimensionField> trendGroups,

            ImmutableArray<Outputs.AnalysisMeasureField> values)
        {
            TargetValues = targetValues;
            TrendGroups = trendGroups;
            Values = values;
        }
    }
}