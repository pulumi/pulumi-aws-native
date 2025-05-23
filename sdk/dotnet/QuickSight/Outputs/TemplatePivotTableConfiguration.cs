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
    public sealed class TemplatePivotTableConfiguration
    {
        /// <summary>
        /// The field options for a pivot table visual.
        /// </summary>
        public readonly Outputs.TemplatePivotTableFieldOptions? FieldOptions;
        /// <summary>
        /// The field wells of the visual.
        /// </summary>
        public readonly Outputs.TemplatePivotTableFieldWells? FieldWells;
        /// <summary>
        /// The general visual interactions setup for a visual.
        /// </summary>
        public readonly Outputs.TemplateVisualInteractionOptions? Interactions;
        /// <summary>
        /// The paginated report options for a pivot table visual.
        /// </summary>
        public readonly Outputs.TemplatePivotTablePaginatedReportOptions? PaginatedReportOptions;
        /// <summary>
        /// The sort configuration for a `PivotTableVisual` .
        /// </summary>
        public readonly Outputs.TemplatePivotTableSortConfiguration? SortConfiguration;
        /// <summary>
        /// The table options for a pivot table visual.
        /// </summary>
        public readonly Outputs.TemplatePivotTableOptions? TableOptions;
        /// <summary>
        /// The total options for a pivot table visual.
        /// </summary>
        public readonly Outputs.TemplatePivotTableTotalOptions? TotalOptions;

        [OutputConstructor]
        private TemplatePivotTableConfiguration(
            Outputs.TemplatePivotTableFieldOptions? fieldOptions,

            Outputs.TemplatePivotTableFieldWells? fieldWells,

            Outputs.TemplateVisualInteractionOptions? interactions,

            Outputs.TemplatePivotTablePaginatedReportOptions? paginatedReportOptions,

            Outputs.TemplatePivotTableSortConfiguration? sortConfiguration,

            Outputs.TemplatePivotTableOptions? tableOptions,

            Outputs.TemplatePivotTableTotalOptions? totalOptions)
        {
            FieldOptions = fieldOptions;
            FieldWells = fieldWells;
            Interactions = interactions;
            PaginatedReportOptions = paginatedReportOptions;
            SortConfiguration = sortConfiguration;
            TableOptions = tableOptions;
            TotalOptions = totalOptions;
        }
    }
}
