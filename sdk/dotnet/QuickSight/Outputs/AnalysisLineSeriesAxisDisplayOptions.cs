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
    public sealed class AnalysisLineSeriesAxisDisplayOptions
    {
        /// <summary>
        /// The options that determine the presentation of the line series axis.
        /// </summary>
        public readonly Outputs.AnalysisAxisDisplayOptions? AxisOptions;
        /// <summary>
        /// The configuration options that determine how missing data is treated during the rendering of a line chart.
        /// </summary>
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
