// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisSmallMultiplesOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxVisibleColumns")]
        public Input<double>? MaxVisibleColumns { get; set; }

        [Input("maxVisibleRows")]
        public Input<double>? MaxVisibleRows { get; set; }

        [Input("panelConfiguration")]
        public Input<Inputs.AnalysisPanelConfigurationArgs>? PanelConfiguration { get; set; }

        public AnalysisSmallMultiplesOptionsArgs()
        {
        }
        public static new AnalysisSmallMultiplesOptionsArgs Empty => new AnalysisSmallMultiplesOptionsArgs();
    }
}