// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardTableConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldOptions")]
        public Input<Inputs.DashboardTableFieldOptionsArgs>? FieldOptions { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.DashboardTableFieldWellsArgs>? FieldWells { get; set; }

        [Input("paginatedReportOptions")]
        public Input<Inputs.DashboardTablePaginatedReportOptionsArgs>? PaginatedReportOptions { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.DashboardTableSortConfigurationArgs>? SortConfiguration { get; set; }

        [Input("tableInlineVisualizations")]
        private InputList<Inputs.DashboardTableInlineVisualizationArgs>? _tableInlineVisualizations;
        public InputList<Inputs.DashboardTableInlineVisualizationArgs> TableInlineVisualizations
        {
            get => _tableInlineVisualizations ?? (_tableInlineVisualizations = new InputList<Inputs.DashboardTableInlineVisualizationArgs>());
            set => _tableInlineVisualizations = value;
        }

        [Input("tableOptions")]
        public Input<Inputs.DashboardTableOptionsArgs>? TableOptions { get; set; }

        [Input("totalOptions")]
        public Input<Inputs.DashboardTotalOptionsArgs>? TotalOptions { get; set; }

        public DashboardTableConfigurationArgs()
        {
        }
        public static new DashboardTableConfigurationArgs Empty => new DashboardTableConfigurationArgs();
    }
}