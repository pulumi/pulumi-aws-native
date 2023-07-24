// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTableConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldOptions")]
        public Input<Inputs.AnalysisPivotTableFieldOptionsArgs>? FieldOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.AnalysisPivotTableFieldWellsArgs>? FieldWells { get; set; }

        [Input("paginatedReportOptions")]
        public Input<Inputs.AnalysisPivotTablePaginatedReportOptionsArgs>? PaginatedReportOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.AnalysisPivotTableSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tableOptions")]
        public Input<Inputs.AnalysisPivotTableOptionsArgs>? TableOptions { get; set; }

        [Input("totalOptions")]
        public Input<Inputs.AnalysisPivotTableTotalOptionsArgs>? TotalOptions { get; set; }

        public AnalysisPivotTableConfigurationArgs()
        {
        }
        public static new AnalysisPivotTableConfigurationArgs Empty => new AnalysisPivotTableConfigurationArgs();
    }
}