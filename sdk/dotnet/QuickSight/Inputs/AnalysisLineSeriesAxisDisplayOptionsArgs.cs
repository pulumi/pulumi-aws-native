// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisLineSeriesAxisDisplayOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The options that determine the presentation of the line series axis.
        /// </summary>
        [Input("axisOptions")]
        public Input<Inputs.AnalysisAxisDisplayOptionsArgs>? AxisOptions { get; set; }

        [Input("missingDataConfigurations")]
        private InputList<Inputs.AnalysisMissingDataConfigurationArgs>? _missingDataConfigurations;

        /// <summary>
        /// The configuration options that determine how missing data is treated during the rendering of a line chart.
        /// </summary>
        public InputList<Inputs.AnalysisMissingDataConfigurationArgs> MissingDataConfigurations
        {
            get => _missingDataConfigurations ?? (_missingDataConfigurations = new InputList<Inputs.AnalysisMissingDataConfigurationArgs>());
            set => _missingDataConfigurations = value;
        }

        public AnalysisLineSeriesAxisDisplayOptionsArgs()
        {
        }
        public static new AnalysisLineSeriesAxisDisplayOptionsArgs Empty => new AnalysisLineSeriesAxisDisplayOptionsArgs();
    }
}
