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
    public sealed class AnalysisLineSeriesAxisDisplayOptions
    {
        public readonly Outputs.AnalysisAxisDisplayOptions? AxisOptions;
        public readonly ImmutableArray<Outputs.AnalysisMissingDataConfiguration> MissingDataConfigurations;

        [OutputConstructor]
        private AnalysisLineSeriesAxisDisplayOptions(
            Outputs.AnalysisAxisDisplayOptions? axisOptions,

            ImmutableArray<Outputs.AnalysisMissingDataConfiguration> missingDataConfigurations)
        {
            AxisOptions = axisOptions;
            MissingDataConfigurations = missingDataConfigurations;
        }
    }
}