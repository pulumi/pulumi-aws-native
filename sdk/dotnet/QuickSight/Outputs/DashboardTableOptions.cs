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
    public sealed class DashboardTableOptions
    {
        /// <summary>
        /// The table cell style of table cells.
        /// </summary>
        public readonly Outputs.DashboardTableCellStyle? CellStyle;
        /// <summary>
        /// The table cell style of a table header.
        /// </summary>
        public readonly Outputs.DashboardTableCellStyle? HeaderStyle;
        /// <summary>
        /// The orientation (vertical, horizontal) for a table.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DashboardTableOrientation? Orientation;
        /// <summary>
        /// The row alternate color options (widget status, row alternate colors) for a table.
        /// </summary>
        public readonly Outputs.DashboardRowAlternateColorOptions? RowAlternateColorOptions;

        [OutputConstructor]
        private DashboardTableOptions(
            Outputs.DashboardTableCellStyle? cellStyle,

            Outputs.DashboardTableCellStyle? headerStyle,

            Pulumi.AwsNative.QuickSight.DashboardTableOrientation? orientation,

            Outputs.DashboardRowAlternateColorOptions? rowAlternateColorOptions)
        {
            CellStyle = cellStyle;
            HeaderStyle = headerStyle;
            Orientation = orientation;
            RowAlternateColorOptions = rowAlternateColorOptions;
        }
    }
}
