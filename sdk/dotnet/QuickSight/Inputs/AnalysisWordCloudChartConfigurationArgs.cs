// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisWordCloudChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryLabelOptions")]
        public Input<Inputs.AnalysisChartAxisLabelOptionsArgs>? CategoryLabelOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisWordCloudFieldWellsArgs>? FieldWells { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisWordCloudSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("wordCloudOptions")]
        public Input<Inputs.AnalysisWordCloudOptionsArgs>? WordCloudOptions { get; set; }

        public AnalysisWordCloudChartConfigurationArgs()
        {
        }
        public static new AnalysisWordCloudChartConfigurationArgs Empty => new AnalysisWordCloudChartConfigurationArgs();
    }
}