// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisListControlDisplayOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("searchOptions")]
        public Input<Inputs.AnalysisListControlSearchOptionsArgs>? SearchOptions { get; set; }

        [Input("selectAllOptions")]
        public Input<Inputs.AnalysisListControlSelectAllOptionsArgs>? SelectAllOptions { get; set; }

        [Input("titleOptions")]
        public Input<Inputs.AnalysisLabelOptionsArgs>? TitleOptions { get; set; }

        public AnalysisListControlDisplayOptionsArgs()
        {
        }
        public static new AnalysisListControlDisplayOptionsArgs Empty => new AnalysisListControlDisplayOptionsArgs();
    }
}