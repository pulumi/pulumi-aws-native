// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateWaterfallChartFieldWellsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field well configuration of a waterfall visual.
        /// </summary>
        [Input("waterfallChartAggregatedFieldWells")]
        public Input<Inputs.TemplateWaterfallChartAggregatedFieldWellsArgs>? WaterfallChartAggregatedFieldWells { get; set; }

        public TemplateWaterfallChartFieldWellsArgs()
        {
        }
        public static new TemplateWaterfallChartFieldWellsArgs Empty => new TemplateWaterfallChartFieldWellsArgs();
    }
}
