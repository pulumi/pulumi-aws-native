// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePieChartFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field well configuration of a pie chart.
        /// </summary>
        [Input("pieChartAggregatedFieldWells")]
        public Input<Inputs.TemplatePieChartAggregatedFieldWellsArgs>? PieChartAggregatedFieldWells { get; set; }

        public TemplatePieChartFieldWellsArgs()
        {
        }
        public static new TemplatePieChartFieldWellsArgs Empty => new TemplatePieChartFieldWellsArgs();
    }
}
